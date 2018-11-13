package blobcache

import (
	"context"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/containers/image/docker/reference"
	"github.com/containers/image/image"
	"github.com/containers/image/manifest"
	"github.com/containers/image/transports"
	"github.com/containers/image/transports/alltransports"
	"github.com/containers/image/types"
	"github.com/containers/storage/pkg/ioutils"
	digest "github.com/opencontainers/go-digest"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var (
	// Transport caches written blobs and offers to return some from its cache.  Its
	// name is "blob-cache", and references which are passed to it to be parsed
	// should take the form "blob-cache:directory[,...]:reference", where _directory_
	// is an optional directory list, and _reference_ is another reference.
	// Importing this package registers Transport with the image library.
	Transport                        = &blobCacheTransport{}
	_         types.ImageReference   = &blobCacheReference{}
	_         types.ImageSource      = &blobCacheSource{}
	_         types.ImageDestination = &blobCacheDestination{}
)

type blobCacheTransport struct {
}

func init() {
	transports.Register(Transport)
}

// BlobCache is an object which saves copies of blobs that are written to it while passing them
// through to some real destination, and which can be queried directly in order to read them
// back.
type BlobCache interface {
	types.ImageReference
	// HasBlob checks if a blob that matches the passed-in digest (and size, if not -1), is
	// present in the cache.
	HasBlob(types.BlobInfo) (bool, int64, error)
	// Directories returns the list of cache directories.
	Directories() []string
	// ClearCache() clears the contents of the cache directories.  Note that this also clears
	// content which was not placed there by this cache implementation.
	ClearCache() error
}

type blobCacheReference struct {
	transport   *blobCacheTransport
	reference   types.ImageReference
	directories []string
}

type blobCacheSource struct {
	reference *blobCacheReference
	source    types.ImageSource
}

type blobCacheDestination struct {
	reference   *blobCacheReference
	destination types.ImageDestination
}

func makeFilename(blobSum digest.Digest, isConfig bool) string {
	if isConfig {
		return blobSum.String() + ".config"
	}
	return blobSum.String()
}

func (t *blobCacheTransport) Name() string {
	return "blob-cache"
}

func (t *blobCacheTransport) ParseReference(reference string) (types.ImageReference, error) {
	components := strings.SplitN(reference, ":", 2)
	realRef := ""
	directories := []string{}
	switch len(components) {
	case 2:
		directories = strings.Split(components[0], ",")
		if len(directories) == 1 && directories[0] == "" {
			directories = []string{}
		}
		realRef = components[1]
	default:
		return nil, errors.Errorf("error parsing reference for %q image %q: bad syntax", t.Name(), reference)
	}
	ref, err := alltransports.ParseImageName(realRef)
	if err != nil {
		return nil, errors.Wrapf(err, "error parsing reference %q", realRef)
	}
	return NewBlobCache(ref, directories)
}

func (t *blobCacheTransport) ValidatePolicyConfigurationScope(scope string) error {
	return nil
}

// NewBlobCache creates a new blob cache that wraps an image reference.  Any blobs which are
// written to the destination image created from the resulting reference will also be stored
// as-is to the specifed directory or a temporary directory.  The cache directory's contents
// can be cleared by calling the returned BlobCache()'s ClearCache() method.
func NewBlobCache(ref types.ImageReference, directories []string) (BlobCache, error) {
	if len(directories) == 0 {
		return nil, errors.Errorf("error building cache: no cache directory specified for %q", transports.ImageName(ref))
	}
	return &blobCacheReference{
		transport:   Transport,
		reference:   ref,
		directories: append([]string{}, directories...),
	}, nil
}

func (r *blobCacheReference) Transport() types.ImageTransport {
	return r.transport
}

func (r *blobCacheReference) StringWithinTransport() string {
	return strings.Join(r.directories, ",") + ":" + transports.ImageName(r.reference)
}

func (r *blobCacheReference) DockerReference() reference.Named {
	return r.reference.DockerReference()
}

func (r *blobCacheReference) PolicyConfigurationIdentity() string {
	return r.reference.PolicyConfigurationIdentity()
}

func (r *blobCacheReference) PolicyConfigurationNamespaces() []string {
	return r.reference.PolicyConfigurationNamespaces()
}

func (r *blobCacheReference) DeleteImage(ctx context.Context, sys *types.SystemContext) error {
	return r.reference.DeleteImage(ctx, sys)
}

func (r *blobCacheReference) HasBlob(blobinfo types.BlobInfo) (bool, int64, error) {
	if blobinfo.Digest == "" {
		return false, -1, nil
	}

	for _, directory := range r.directories {
		for _, isConfig := range []bool{false, true} {
			filename := filepath.Join(directory, makeFilename(blobinfo.Digest, isConfig))
			fileInfo, err := os.Stat(filename)
			if err == nil && (blobinfo.Size == -1 || blobinfo.Size == fileInfo.Size()) {
				return true, fileInfo.Size(), nil
			}
			if !os.IsNotExist(err) {
				return false, -1, errors.Wrapf(err, "error checking size of %q", filename)
			}
		}
	}

	return false, -1, nil
}

func (r *blobCacheReference) Directories() []string {
	return append([]string{}, r.directories...)
}

func (r *blobCacheReference) ClearCache() error {
	for _, directory := range r.directories {
		// Clear the directory's contents.
		f, err := os.Open(directory)
		if err != nil {
			return errors.Wrapf(err, "error opening directory %q", directory)
		}
		defer f.Close()
		names, err := f.Readdirnames(-1)
		if err != nil {
			return errors.Wrapf(err, "error reading directory %q", directory)
		}
		for _, name := range names {
			pathname := filepath.Join(directory, name)
			if err = os.RemoveAll(pathname); err != nil {
				return errors.Wrapf(err, "error removing %q while clearing cache %q", pathname, transports.ImageName(r))
			}
		}
	}
	return nil
}

func (r *blobCacheReference) NewImage(ctx context.Context, sys *types.SystemContext) (types.ImageCloser, error) {
	src, err := r.NewImageSource(ctx, sys)
	if err != nil {
		return nil, errors.Wrapf(err, "error creating new image %q", transports.ImageName(r.reference))
	}
	return image.FromSource(ctx, sys, src)
}

func (r *blobCacheReference) NewImageSource(ctx context.Context, sys *types.SystemContext) (types.ImageSource, error) {
	src, err := r.reference.NewImageSource(ctx, sys)
	if err != nil {
		return nil, errors.Wrapf(err, "error creating new image source %q", transports.ImageName(r.reference))
	}
	return &blobCacheSource{reference: r, source: src}, nil
}

func (r *blobCacheReference) NewImageDestination(ctx context.Context, sys *types.SystemContext) (types.ImageDestination, error) {
	dest, err := r.reference.NewImageDestination(ctx, sys)
	if err != nil {
		return nil, errors.Wrapf(err, "error creating new image destination %q", transports.ImageName(r.reference))
	}
	return &blobCacheDestination{reference: r, destination: dest}, nil
}

func (s *blobCacheSource) Reference() types.ImageReference {
	return s.reference
}

func (s *blobCacheSource) Close() error {
	return s.source.Close()
}

func (s *blobCacheSource) GetManifest(ctx context.Context, instanceDigest *digest.Digest) ([]byte, string, error) {
	if instanceDigest != nil {
		for _, directory := range s.reference.directories {
			filename := filepath.Join(directory, makeFilename(*instanceDigest, false))
			manifestBytes, err := ioutil.ReadFile(filename)
			if err == nil {
				return manifestBytes, manifest.GuessMIMEType(manifestBytes), nil
			}
			if !os.IsNotExist(err) {
				return nil, "", errors.Wrapf(err, "error checking for manifest file %q", filename)
			}
		}
	}
	return s.source.GetManifest(ctx, instanceDigest)
}

func (s *blobCacheSource) GetBlob(ctx context.Context, blobinfo types.BlobInfo) (io.ReadCloser, int64, error) {
	present, size, err := s.reference.HasBlob(blobinfo)
	if err != nil {
		return nil, -1, err
	}
	if present {
		for _, directory := range s.reference.directories {
			for _, isConfig := range []bool{false, true} {
				filename := filepath.Join(directory, makeFilename(blobinfo.Digest, isConfig))
				f, err := os.Open(filename)
				if err == nil {
					return f, size, nil
				}
				if !os.IsNotExist(err) {
					return nil, -1, errors.Wrapf(err, "error checking for cache file %q", filepath.Join(directory, filename))
				}
			}
		}
	}
	rc, size, err := s.source.GetBlob(ctx, blobinfo)
	if err != nil {
		return rc, size, errors.Wrapf(err, "error reading blob from source image %q", transports.ImageName(s.reference))
	}
	return rc, size, nil
}

func (s *blobCacheSource) GetSignatures(ctx context.Context, instanceDigest *digest.Digest) ([][]byte, error) {
	return s.source.GetSignatures(ctx, instanceDigest)
}

func (s *blobCacheSource) LayerInfosForCopy(ctx context.Context) ([]types.BlobInfo, error) {
	return nil, nil
}

func (d *blobCacheDestination) Reference() types.ImageReference {
	return d.reference
}

func (d *blobCacheDestination) Close() error {
	return d.destination.Close()
}

func (d *blobCacheDestination) SupportedManifestMIMETypes() []string {
	return d.destination.SupportedManifestMIMETypes()
}

func (d *blobCacheDestination) SupportsSignatures(ctx context.Context) error {
	return d.destination.SupportsSignatures(ctx)
}

func (d *blobCacheDestination) DesiredLayerCompression() types.LayerCompression {
	return d.destination.DesiredLayerCompression()
}

func (d *blobCacheDestination) AcceptsForeignLayerURLs() bool {
	return d.destination.AcceptsForeignLayerURLs()
}

func (d *blobCacheDestination) MustMatchRuntimeOS() bool {
	return d.destination.MustMatchRuntimeOS()
}

func (d *blobCacheDestination) IgnoresEmbeddedDockerReference() bool {
	return d.destination.IgnoresEmbeddedDockerReference()
}

func (d *blobCacheDestination) PutBlob(ctx context.Context, stream io.Reader, inputInfo types.BlobInfo, isConfig bool) (types.BlobInfo, error) {
	var tempfile *os.File
	var err error
	if inputInfo.Digest != "" {
		directory := d.reference.directories[0]
		if directory == "" {
			directory = "."
		}
		filename := filepath.Join(directory, makeFilename(inputInfo.Digest, isConfig))
		tempfile, err = ioutil.TempFile(directory, makeFilename(inputInfo.Digest, isConfig))
		if err == nil {
			stream = io.TeeReader(stream, tempfile)
			defer func() {
				if err == nil {
					if err = os.Rename(tempfile.Name(), filename); err != nil {
						err = errors.Wrapf(err, "error renaming new layer for blob %q into place at %q", inputInfo.Digest.String(), filename)
					}
				} else {
					if err2 := os.Remove(tempfile.Name()); err2 != nil {
						logrus.Debugf("error cleaning up temporary file %q for blob %q: %v", tempfile.Name(), inputInfo.Digest.String(), err2)
					}
				}
				tempfile.Close()
			}()
		} else {
			logrus.Debugf("error while creating a temporary file under %q to hold blob %q: %v", d.reference.directories[0], inputInfo.Digest.String(), err)
		}
	}
	newBlobInfo, err := d.destination.PutBlob(ctx, stream, inputInfo, isConfig)
	if err != nil {
		return newBlobInfo, errors.Wrapf(err, "error storing blob to image destination for cache %q", transports.ImageName(d.reference))
	}
	return newBlobInfo, nil
}

func (d *blobCacheDestination) HasBlob(ctx context.Context, info types.BlobInfo) (bool, int64, error) {
	present, size, err := d.reference.HasBlob(info)
	if err != nil {
		return false, -1, errors.Wrapf(err, "error checking for blob %q in cache %#v", info.Digest.String(), d.reference.directories)
	}
	if present {
		return present, size, nil
	}
	return d.destination.HasBlob(ctx, info)
}

func (d *blobCacheDestination) ReapplyBlob(ctx context.Context, info types.BlobInfo) (types.BlobInfo, error) {
	present, _, err := d.destination.HasBlob(ctx, info)
	if err != nil {
		return types.BlobInfo{}, errors.Wrapf(err, "error checking for blob %q in cache %#v", info.Digest.String(), d.reference.directories)
	}
	if !present {
		for _, directory := range d.reference.directories {
			for _, isConfig := range []bool{false, true} {
				filename := filepath.Join(directory, makeFilename(info.Digest, isConfig))
				f, err := os.Open(filename)
				if err == nil {
					defer f.Close()
					return d.destination.PutBlob(ctx, f, info, isConfig)
				}
			}
		}
	}
	return d.destination.ReapplyBlob(ctx, info)
}

func (d *blobCacheDestination) PutManifest(ctx context.Context, manifestBytes []byte) error {
	if len(d.reference.directories) > 0 {
		manifestDigest, err := manifest.Digest(manifestBytes)
		if err != nil {
			logrus.Warnf("error digesting manifest %q: %v", string(manifestBytes), err)
		} else {
			filename := filepath.Join(d.reference.directories[0], makeFilename(manifestDigest, false))
			if err = ioutils.AtomicWriteFile(filename, manifestBytes, 0600); err != nil {
				logrus.Warnf("error saving manifest as %q: %v", filename, err)
			}
		}
	}
	return d.destination.PutManifest(ctx, manifestBytes)
}

func (d *blobCacheDestination) PutSignatures(ctx context.Context, signatures [][]byte) error {
	return d.destination.PutSignatures(ctx, signatures)
}

func (d *blobCacheDestination) Commit(ctx context.Context) error {
	return d.destination.Commit(ctx)
}
