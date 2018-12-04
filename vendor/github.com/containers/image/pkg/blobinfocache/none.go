package blobinfocache

import (
	"github.com/containers/image/types"
	"github.com/opencontainers/go-digest"
)

// noCache implements a dummy BlobInfoCache which records no data.
type noCache struct {
}

// NoCache implements BlobInfoCache by not recording any data.
//
// This exists primarily for implementations of configGetter for Manifest.Inspect,
// because configs only have one representation.
// Any use of BlobInfoCache with blobs should usually use at least a short-lived cache.
var NoCache types.BlobInfoCache = noCache{}

func (noCache) UncompressedDigest(anyDigest digest.Digest) digest.Digest {
	return ""
}

// Records that the uncompressed version of anyDigest is uncompressed.
// It’s allowed for anyDigest == uncompressed (this is useful to avoid computing a DiffID when converting from schema1).
// WARNING: Only call this for LOCALLY VERIFIED data; don’t record a pair just because some remote author claims so (e.g. a manifest/config pair exists); otherwise the cache could be poisoned and allow substituting unexpected blobs.
// (Eventually, the DiffIDs in image config could detect the substitution, but that may be too late, and not all image formats contain that data.)
func (noCache) RecordDigestUncompressedPair(anyDigest digest.Digest, uncompressed digest.Digest) {
}

func (noCache) RecordKnownLocation(transport types.ImageTransport, scope types.BICTransportScope, blobDigest digest.Digest, location types.BICLocationReference) {
}

func (noCache) CandidateLocations(transport types.ImageTransport, scope types.BICTransportScope, digest digest.Digest, canSubstitute bool) []types.BICReplacementCandidate {
	return nil
}
