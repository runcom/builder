#!/bin/bash

# This script is called by packer on the subject CentOS VM, to setup the podman
# build/test environment.  It's not intended to be used outside of this context.

set -e

# Load in library (copied by packer, before this script was run)
source /tmp/libpod/$SCRIPT_BASE/lib.sh

req_env_var "
SCRIPT_BASE $SCRIPT_BASE
CNI_COMMIT $CNI_COMMIT
CRIO_COMMIT $CRIO_COMMIT
CRIU_COMMIT $CRIU_COMMIT
RHSM_COMMAND $RHSM_COMMAND
"

install_ooe

rhsm_enable

ooe.sh sudo yum -y erase "rh-amazon-rhui-client*"
ooe.sh sudo subscription-manager repos "--disable=*"
ooe.sh sudo subscription-manager repos \
    --enable=rhel-7-server-rpms \
    --enable=rhel-7-server-optional-rpms \
    --enable=rhel-7-server-extras-rpms \
    --enable=rhel-server-rhscl-7-rpms

ooe.sh sudo yum -y update

ooe.sh sudo yum -y install \
    atomic-registries \
    btrfs-progs-devel \
    bzip2 \
    device-mapper-devel \
    emacs-nox \
    findutils \
    glib2-devel \
    glibc-static \
    gnupg \
    golang \
    golang-github-cpuguy83-go-md2man \
    golang-github-cpuguy83-go-md2man \
    gpgme-devel \
    iptables \
    libassuan-devel \
    libcap-devel \
    libnet \
    libnet-devel \
    libnl3-devel \
    libseccomp-devel \
    libselinux-devel \
    lsof \
    make \
    nmap-ncat \
    ostree-devel \
    protobuf \
    protobuf-c \
    protobuf-c-devel \
    protobuf-compiler \
    protobuf-devel \
    protobuf-python \
    python \
    python2-future \
    python34-dateutil \
    python34-psutil \
    python34-pytoml \
    runc \
    skopeo-containers \
    unzip \
    vim \
    which \
    xz

install_scl_git

install_cni_plugins

install_buildah

install_conmon

install_criu

install_packer_copied_files

rhel_exit_handler  # release subscription!

rh_finalize

echo "SUCCESS!"
