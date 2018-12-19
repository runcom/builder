![PODMAN logo](../../logo/podman-logo-source.svg)

# Cirrus-CI

Similar to other integrated github CI/CD services, Cirrus utilizes a simple
YAML-based configuration/description file: ``.cirrus.yml``.  Ref: https://cirrus-ci.org/


## Workflow

All tasks execute in parallel, unless there are conditions or dependencies
which alter this behavior.  Within each task, each script executes in sequence,
so long as any previous script exited successfully.  The overall state of each
task (pass or fail) is set based on the exit status of the last script to execute.


### ``gating`` Task

***N/B: Steps below are performed by automation***

1. Launch a purpose-built container in Cirrus's community cluster.
   For container image details, please see
   [the contributors guide](https://github.com/containers/libpod/blob/master/CONTRIBUTING.md#go-format-and-lint).

3. ``validate``: Perform standard `make validate` source verification,
   Should run for less than a minute or two.

4. ``lint``: Execute regular `make lint` to check for any code cruft.
   Should also run for less than a few minutes.


### ``testing`` Task

***N/B: Steps below are performed by automation***

1. After `gating` passes, spin up one VM per
   `matrix: image_name` item. Once accessible, ``ssh``
   into each VM as the `root` user.

2. ``setup_environment.sh``: Configure root's `.bash_profile`
    for all subsequent scripts (each run in a new shell).  Any
    distribution-specific environment variables are also defined
    here.  For example, setting tags/flags to use compiling.

5. ``integration_test.sh``: Execute integration-testing.  This is
   much more involved, and relies on access to external
   resources like container images and code from other repositories.
   Total execution time is capped at 2-hours (includes all the above)
   but this script normally completes in less than an hour.


### ``optional_testing`` Task

***N/B: Steps below are performed by automation***

1. Optionally executes in parallel with ``testing``.  Requires
    **prior** to job-start, the magic string ``***CIRRUS: SYSTEM TEST***``
   is found in the pull-request *description*.  The *description* is the first
   text-box under the main *summary* line in the github WebUI.

2. ``setup_environment.sh``: Same as for other tasks.

3. ``system_test.sh``: Build both dependencies and libpod, install them,
   then execute `make localsystem` from the repository root.


### ``cache_images`` Task

Modifying the contents of cache-images is done by making changes to
one or more of the ``./contrib/cirrus/packer/*_setup.sh`` files.  Testing
those changes currently requires adding a temporary commit to a PR that
updates ``.cirrus.yml``:

* Remove all task sections except ``cache_images_task``.
* Remove the ``only_if`` condition and ``depends_on`` dependencies

The new image names will be displayed at the end of output, assuming the build
is successful, at that point the temporary commit may be removed.  Finally,
the new names may be used as ``image_name`` values in ``.cirrus.yml``.

***N/B: Steps below are performed by automation***

1. When a PR is merged (``$CIRRUS_BRANCH`` == ``master``), run another
   round of the ``gating`` and ``testing`` tasks (above).

2. Assuming tests pass, if the commit message contains the magic string
   ``***CIRRUS: REBUILD IMAGES***``, then this task continues.  Otherwise
   simply mark the master branch as 'passed'.

3. ``setup_environment.sh``: Same as for other tasks.

4. ``build_vm_images.sh``: Utilize [the packer tool](http://packer.io/docs/)
   to produce new VM images.  Create a new VM from each base-image, connect
   to them with ``ssh``, and perform the steps as defined by the
   ``$PACKER_BASE/libpod_images.json`` file:

    1. On a base-image VM, as root, copy the current state of the repository
       into ``/tmp/libpod``.
    2. Execute distribution-specific scripts to prepare the image for
       use by the ``integration_testing`` task (above).  For example,
       ``fedora_setup.sh``.
    3. If successful, shut down each VM and create a new GCE Image
       named with the base image, and the commit sha of the merge.

### Base-images

Base-images are VM disk-images specially prepared for executing as GCE VMs.
In particular, they run services on startup similar in purpose/function
as the standard 'cloud-init' services.

*  The google services are required for full support of ssh-key management
   and GCE OAuth capabilities.  Google provides native images in GCE
   with services pre-installed, for many platforms. For example,
   RHEL, CentOS, and Ubuntu.

*  Google does ***not*** provide any images for Fedora or Fedora Atomic
   Host (as of 11/2018), nor do they provide a base-image prepared to
   run packer for creating other images in the ``build_vm_images`` Task
   (above).

*  Base images do not need to be produced often, but doing so completely
   manually would be time-consuming and error-prone.  Therefor a special
   semi-automatic *Makefile* target is provided to assist with producing
   all the base-images: ``libpod_base_images``

To produce new base-images, including an `image-builder-image` (used by
the ``cache_images`` Task) some input parameters are required:

* ``GCP_PROJECT_ID``: The complete GCP project ID string e.g. foobar-12345
  identifying where the images will be stored.

* ``GOOGLE_APPLICATION_CREDENTIALS``: A *JSON* file containing
  credentials for a GCE service account.  This can be [a service
  account](https://cloud.google.com/docs/authentication/production#obtaining_and_providing_service_account_credentials_manually)
  or [end-user
  credentials](https://cloud.google.com/docs/authentication/end-user#creating_your_client_credentials)

* ``RHEL_IMAGE_FILE`` and ``RHEL_CSUM_FILE`` complete paths
  to a `rhel-server-ec2-*.raw.xz` and it's cooresponding
  checksum file.  These must be supplied manually because
  they're not available directly via URL like other images.

* ``RHSM_COMMAND`` contains the complete string needed to register
  the VM for installing package dependencies.  The VM will be de-registered
  upon completion.

*  Optionally, CSV's may be specified to ``PACKER_BUILDS``
   to limit the base-images produced.  For example,
   ``PACKER_BUILDS=fedora,image-builder-image``.

If there is an existing 'image-builder-image' within GCE, it may be utilized
to produce base-images (in addition to cache-images).  However it must be
created with support for nested-virtualization, and with elevated cloud
privileges (to access GCE, from within the GCE VM).  For example:

```
$ alias pgcloud='sudo podman run -it --rm -e AS_ID=$UID
    -e AS_USER=$USER -v $HOME:$HOME:z quay.io/cevich/gcloud_centos:latest'

$ URL=https://www.googleapis.com/auth
$ SCOPES=$URL/userinfo.email,$URL/compute,$URL/devstorage.full_control

# The --min-cpu-platform is critical for nested-virt.
$ pgcloud compute instances create $USER-making-images \
    --image-family image-builder-image \
    --boot-disk-size "200GB" \
    --min-cpu-platform "Intel Haswell" \
    --machine-type n1-standard-2 \
    --scopes $SCOPES
```

Alternatively, if there is no image-builder-image available yet, a bare-metal
CentOS 7 machine with network access to GCE is required.  Software dependencies
can be obtained from the ``packer/image-builder-image_base_setup.sh`` script.

In both cases, the following can be used to setup and build base-images.

```
$ IP_ADDRESS=1.2.3.4  # EXTERNAL_IP from command output above
$ rsync -av $PWD centos@$IP_ADDRESS:.
$ scp $GOOGLE_APPLICATION_CREDENTIALS centos@$IP_ADDRESS:.
$ ssh centos@$IP_ADDRESS
...
```

When ready, change to the ``packer`` sub-directory, and build the images:

```
$ cd libpod/contrib/cirrus/packer
$ make libpod_base_images GCP_PROJECT_ID=<VALUE> \
    GOOGLE_APPLICATION_CREDENTIALS=<VALUE> \
    RHEL_IMAGE_FILE=<VALUE> \
    RHEL_CSUM_FILE=<VALUE> \
    RHSM_COMMAND=<VALUE> \
    PACKER_BUILDS=<OPTIONAL>
```

Assuming this is successful (hence the semi-automatic part), packer will
produce a ``packer-manifest.json`` output file.  This contains the base-image
names suitable for updating in ``.cirrus.yml``, `env` keys ``*_BASE_IMAGE``.

On failure, it should be possible to determine the problem from the packer
output.  Sometimes that means setting `PACKER_LOG=1` and troubleshooting
the nested virt calls.  It's also possible to observe the (nested) qemu-kvm
console output.  Simply set the ``TTYDEV`` parameter, for example:

```
$ make libpod_base_images ... TTYDEV=$(tty)
  ...
```
