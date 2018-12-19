# io.podman
Podman Service Interface and API description.  The master version of this document can be found
in the [API.md](https://github.com/containers/libpod/blob/master/API.md) file in the upstream libpod repository.
## Index

[func AttachToContainer() NotImplemented](#AttachToContainer)

[func BuildImage(build: BuildInfo) BuildResponse](#BuildImage)

[func Commit(name: string, image_name: string, changes: []string, author: string, message: string, pause: bool, manifestType: string) string](#Commit)

[func ContainerCheckpoint(name: string, keep: bool, leaveRunning: bool, tcpEstablished: bool) string](#ContainerCheckpoint)

[func ContainerExists(name: string) int](#ContainerExists)

[func ContainerRestore(name: string, keep: bool, tcpEstablished: bool) string](#ContainerRestore)

[func ContainerRunlabel(runlabel: Runlabel) ](#ContainerRunlabel)

[func CreateContainer(create: Create) string](#CreateContainer)

[func CreateImage() NotImplemented](#CreateImage)

[func CreatePod(create: PodCreate) string](#CreatePod)

[func DeleteStoppedContainers() []string](#DeleteStoppedContainers)

[func DeleteUnusedImages() []string](#DeleteUnusedImages)

[func ExportContainer(name: string, path: string) string](#ExportContainer)

[func ExportImage(name: string, destination: string, compress: bool, tags: []string) string](#ExportImage)

[func GetAttachSockets(name: string) Sockets](#GetAttachSockets)

[func GetContainer(name: string) ListContainerData](#GetContainer)

[func GetContainerLogs(name: string) []string](#GetContainerLogs)

[func GetContainerStats(name: string) ContainerStats](#GetContainerStats)

[func GetImage(name: string) ImageInList](#GetImage)

[func GetInfo() PodmanInfo](#GetInfo)

[func GetPod(name: string) ListPodData](#GetPod)

[func GetPodStats(name: string) string, ContainerStats](#GetPodStats)

[func GetVersion() Version](#GetVersion)

[func HistoryImage(name: string) ImageHistory](#HistoryImage)

[func ImageExists(name: string) int](#ImageExists)

[func ImportImage(source: string, reference: string, message: string, changes: []string) string](#ImportImage)

[func InspectContainer(name: string) string](#InspectContainer)

[func InspectImage(name: string) string](#InspectImage)

[func InspectPod(name: string) string](#InspectPod)

[func KillContainer(name: string, signal: int) string](#KillContainer)

[func KillPod(name: string, signal: int) string](#KillPod)

[func ListContainerChanges(name: string) ContainerChanges](#ListContainerChanges)

[func ListContainerMounts() []string](#ListContainerMounts)

[func ListContainerPorts(name: string) NotImplemented](#ListContainerPorts)

[func ListContainerProcesses(name: string, opts: []string) []string](#ListContainerProcesses)

[func ListContainers() ListContainerData](#ListContainers)

[func ListImages() ImageInList](#ListImages)

[func ListPods() ListPodData](#ListPods)

[func MountContainer(name: string) string](#MountContainer)

[func PauseContainer(name: string) string](#PauseContainer)

[func PausePod(name: string) string](#PausePod)

[func Ping() StringResponse](#Ping)

[func PullImage(name: string, certDir: string, creds: string, signaturePolicy: string, tlsVerify: bool) string](#PullImage)

[func PushImage(name: string, tag: string, tlsverify: bool, signaturePolicy: string, creds: string, certDir: string, compress: bool, format: string, removeSignatures: bool, signBy: string) string](#PushImage)

[func RemoveContainer(name: string, force: bool) string](#RemoveContainer)

[func RemoveImage(name: string, force: bool) string](#RemoveImage)

[func RemovePod(name: string, force: bool) string](#RemovePod)

[func RenameContainer() NotImplemented](#RenameContainer)

[func ResizeContainerTty() NotImplemented](#ResizeContainerTty)

[func RestartContainer(name: string, timeout: int) string](#RestartContainer)

[func RestartPod(name: string) string](#RestartPod)

[func SearchImage(name: string, limit: int) ImageSearch](#SearchImage)

[func StartContainer(name: string) string](#StartContainer)

[func StartPod(name: string) string](#StartPod)

[func StopContainer(name: string, timeout: int) string](#StopContainer)

[func StopPod(name: string, timeout: int) string](#StopPod)

[func TagImage(name: string, tagged: string) string](#TagImage)

[func TopPod() NotImplemented](#TopPod)

[func UnmountContainer(name: string, force: bool) ](#UnmountContainer)

[func UnpauseContainer(name: string) string](#UnpauseContainer)

[func UnpausePod(name: string) string](#UnpausePod)

[func UpdateContainer() NotImplemented](#UpdateContainer)

[func WaitContainer(name: string) int](#WaitContainer)

[func WaitPod() NotImplemented](#WaitPod)

[type BuildInfo](#BuildInfo)

[type BuildResponse](#BuildResponse)

[type ContainerChanges](#ContainerChanges)

[type ContainerMount](#ContainerMount)

[type ContainerNameSpace](#ContainerNameSpace)

[type ContainerPortMappings](#ContainerPortMappings)

[type ContainerStats](#ContainerStats)

[type Create](#Create)

[type CreateResourceConfig](#CreateResourceConfig)

[type IDMap](#IDMap)

[type IDMappingOptions](#IDMappingOptions)

[type ImageHistory](#ImageHistory)

[type ImageInList](#ImageInList)

[type ImageSearch](#ImageSearch)

[type InfoDistribution](#InfoDistribution)

[type InfoGraphStatus](#InfoGraphStatus)

[type InfoHost](#InfoHost)

[type InfoPodmanBinary](#InfoPodmanBinary)

[type InfoStore](#InfoStore)

[type ListContainerData](#ListContainerData)

[type ListPodContainerInfo](#ListPodContainerInfo)

[type ListPodData](#ListPodData)

[type NotImplemented](#NotImplemented)

[type PodContainerErrorData](#PodContainerErrorData)

[type PodCreate](#PodCreate)

[type PodmanInfo](#PodmanInfo)

[type Runlabel](#Runlabel)

[type Sockets](#Sockets)

[type StringResponse](#StringResponse)

[type Version](#Version)

[error ContainerNotFound](#ContainerNotFound)

[error ErrorOccurred](#ErrorOccurred)

[error ImageNotFound](#ImageNotFound)

[error NoContainerRunning](#NoContainerRunning)

[error NoContainersInPod](#NoContainersInPod)

[error PodContainerError](#PodContainerError)

[error PodNotFound](#PodNotFound)

[error RuntimeError](#RuntimeError)

## Methods
### <a name="AttachToContainer"></a>func AttachToContainer
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method AttachToContainer() [NotImplemented](#NotImplemented)</div>
This method has not be implemented yet.
### <a name="BuildImage"></a>func BuildImage
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method BuildImage(build: [BuildInfo](#BuildInfo)) [BuildResponse](#BuildResponse)</div>
BuildImage takes a [BuildInfo](#BuildInfo) structure and builds an image.  At a minimum, you must provide the
'dockerfile' and 'tags' options in the BuildInfo structure. It will return a [BuildResponse](#BuildResponse) structure
that contains the build logs and resulting image ID.
### <a name="Commit"></a>func Commit
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method Commit(name: [string](https://godoc.org/builtin#string), image_name: [string](https://godoc.org/builtin#string), changes: [[]string](#[]string), author: [string](https://godoc.org/builtin#string), message: [string](https://godoc.org/builtin#string), pause: [bool](https://godoc.org/builtin#bool), manifestType: [string](https://godoc.org/builtin#string)) [string](https://godoc.org/builtin#string)</div>
Commit, creates an image from an existing container. It requires the name or
ID of the container as well as the resulting image name.  Optionally, you can define an author and message
to be added to the resulting image.  You can also define changes to the resulting image for the following
attributes: _CMD, ENTRYPOINT, ENV, EXPOSE, LABEL, ONBUILD, STOPSIGNAL, USER, VOLUME, and WORKDIR_.  To pause the
container while it is being committed, pass a _true_ bool for the pause argument.  If the container cannot
be found by the ID or name provided, a (ContainerNotFound)[#ContainerNotFound] error will be returned; otherwise,
the resulting image's ID will be returned as a string.
### <a name="ContainerCheckpoint"></a>func ContainerCheckpoint
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method ContainerCheckpoint(name: [string](https://godoc.org/builtin#string), keep: [bool](https://godoc.org/builtin#bool), leaveRunning: [bool](https://godoc.org/builtin#bool), tcpEstablished: [bool](https://godoc.org/builtin#bool)) [string](https://godoc.org/builtin#string)</div>
ContainerCheckPoint performs a checkpopint on a container by its name or full/partial container
ID.  On successful checkpoint, the id of the checkpointed container is returned.
### <a name="ContainerExists"></a>func ContainerExists
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method ContainerExists(name: [string](https://godoc.org/builtin#string)) [int](https://godoc.org/builtin#int)</div>
ContainerExists takes a full or partial container ID or name and returns an int as to
whether the container exists in local storage.  A result of 0 means the container does
exists; whereas a result of 1 means it could not be found.
#### Example
~~~
$ varlink call -m unix:/run/podman/io.podman/io.podman.ContainerExists '{"name": "flamboyant_payne"}'{
  "exists": 0
}
~~~
### <a name="ContainerRestore"></a>func ContainerRestore
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method ContainerRestore(name: [string](https://godoc.org/builtin#string), keep: [bool](https://godoc.org/builtin#bool), tcpEstablished: [bool](https://godoc.org/builtin#bool)) [string](https://godoc.org/builtin#string)</div>
ContainerRestore restores a container that has been checkpointed.  The container to be restored can
be identified by its name or full/partial container ID.  A successful restore will result in the return
of the container's ID.
### <a name="ContainerRunlabel"></a>func ContainerRunlabel
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method ContainerRunlabel(runlabel: [Runlabel](#Runlabel)) </div>
ContainerRunlabel runs executes a command as described by a given container image label.
### <a name="CreateContainer"></a>func CreateContainer
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method CreateContainer(create: [Create](#Create)) [string](https://godoc.org/builtin#string)</div>
CreateContainer creates a new container from an image.  It uses a [Create](#Create) type for input. The minimum
input required for CreateContainer is an image name.  If the image name is not found, an [ImageNotFound](#ImageNotFound)
error will be returned.  Otherwise, the ID of the newly created container will be returned.
#### Example
~~~
$ varlink call unix:/run/podman/io.podman/io.podman.CreateContainer '{"create": {"image": "alpine"}}'
{
  "container": "8759dafbc0a4dc3bcfb57eeb72e4331eb73c5cc09ab968e65ce45b9ad5c4b6bb"
}
~~~
### <a name="CreateImage"></a>func CreateImage
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method CreateImage() [NotImplemented](#NotImplemented)</div>
This function is not implemented yet.
### <a name="CreatePod"></a>func CreatePod
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method CreatePod(create: [PodCreate](#PodCreate)) [string](https://godoc.org/builtin#string)</div>
CreatePod creates a new empty pod.  It uses a [PodCreate](#PodCreate) type for input.
On success, the ID of the newly created pod will be returned.
#### Example
~~~
$ varlink call unix:/run/podman/io.podman/io.podman.CreatePod '{"create": {"name": "test"}}'
{
  "pod": "b05dee7bd4ccfee688099fe1588a7a898d6ddd6897de9251d4671c9b0feacb2a"
}
# $ varlink call unix:/run/podman/io.podman/io.podman.CreatePod '{"create": {"infra": true, "share": ["ipc", "net", "uts"]}}'
{
  "pod": "d7697449a8035f613c1a8891286502aca68fff7d5d49a85279b3bda229af3b28"
}
~~~
### <a name="DeleteStoppedContainers"></a>func DeleteStoppedContainers
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method DeleteStoppedContainers() [[]string](#[]string)</div>
DeleteStoppedContainers will delete all containers that are not running. It will return a list the deleted
container IDs.  See also [RemoveContainer](RemoveContainer).
#### Example
~~~
$ varlink call -m unix:/run/podman/io.podman/io.podman.DeleteStoppedContainers
{
  "containers": [
    "451410b931d00def8aa9b4f8084e4d4a39e5e04ea61f358cf53a5cf95afcdcee",
    "8b60f754a3e01389494a9581ade97d35c2765b6e2f19acd2d3040c82a32d1bc0",
    "cf2e99d4d3cad6073df199ed32bbe64b124f3e1aba6d78821aa8460e70d30084",
    "db901a329587312366e5ecff583d08f0875b4b79294322df67d90fc6eed08fc1"
  ]
}
~~~
### <a name="DeleteUnusedImages"></a>func DeleteUnusedImages
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method DeleteUnusedImages() [[]string](#[]string)</div>
DeleteUnusedImages deletes any images not associated with a container.  The IDs of the deleted images are returned
in a string array.
#### Example
~~~
$ varlink call -m unix:/run/podman/io.podman/io.podman.DeleteUnusedImages
{
  "images": [
    "166ea6588079559c724c15223f52927f514f73dd5c5cf2ae2d143e3b2e6e9b52",
    "da86e6ba6ca197bf6bc5e9d900febd906b133eaa4750e6bed647b0fbe50ed43e",
    "3ef70f7291f47dfe2b82931a993e16f5a44a0e7a68034c3e0e086d77f5829adc",
    "59788edf1f3e78cd0ebe6ce1446e9d10788225db3dedcfd1a59f764bad2b2690"
  ]
}
~~~
### <a name="ExportContainer"></a>func ExportContainer
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method ExportContainer(name: [string](https://godoc.org/builtin#string), path: [string](https://godoc.org/builtin#string)) [string](https://godoc.org/builtin#string)</div>
ExportContainer creates an image from a container.  It takes the name or ID of a container and a
path representing the target tarfile.  If the container cannot be found, a [ContainerNotFound](#ContainerNotFound)
error will be returned.
The return value is the written tarfile.
#### Example
~~~
$ varlink call -m unix:/run/podman/io.podman/io.podman.ExportContainer '{"name": "flamboyant_payne", "path": "/tmp/payne.tar" }'
{
  "tarfile": "/tmp/payne.tar"
}
~~~
### <a name="ExportImage"></a>func ExportImage
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method ExportImage(name: [string](https://godoc.org/builtin#string), destination: [string](https://godoc.org/builtin#string), compress: [bool](https://godoc.org/builtin#bool), tags: [[]string](#[]string)) [string](https://godoc.org/builtin#string)</div>
ExportImage takes the name or ID of an image and exports it to a destination like a tarball.  There is also
a booleon option to force compression.  It also takes in a string array of tags to be able to save multiple
tags of the same image to a tarball (each tag should be of the form <image>:<tag>).  Upon completion, the ID
of the image is returned. If the image cannot be found in local storage, an [ImageNotFound](#ImageNotFound)
error will be returned. See also [ImportImage](ImportImage).
### <a name="GetAttachSockets"></a>func GetAttachSockets
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method GetAttachSockets(name: [string](https://godoc.org/builtin#string)) [Sockets](#Sockets)</div>
GetAttachSockets takes the name or ID of an existing container.  It returns file paths for two sockets needed
to properly communicate with a container.  The first is the actual I/O socket that the container uses.  The
second is a "control" socket where things like resizing the TTY events are sent. If the container cannot be
found, a [ContainerNotFound](#ContainerNotFound) error will be returned.
#### Example
~~~
$ varlink call -m unix:/run/io.podman/io.podman.GetAttachSockets '{"name": "b7624e775431219161"}'
{
  "sockets": {
    "container_id": "b7624e7754312191613245ce1a46844abee60025818fe3c3f3203435623a1eca",
    "control_socket": "/var/lib/containers/storage/overlay-containers/b7624e7754312191613245ce1a46844abee60025818fe3c3f3203435623a1eca/userdata/ctl",
    "io_socket": "/var/run/libpod/socket/b7624e7754312191613245ce1a46844abee60025818fe3c3f3203435623a1eca/attach"
  }
}
~~~
### <a name="GetContainer"></a>func GetContainer
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method GetContainer(name: [string](https://godoc.org/builtin#string)) [ListContainerData](#ListContainerData)</div>
GetContainer takes a name or ID of a container and returns single ListContainerData
structure.  A [ContainerNotFound](#ContainerNotFound) error will be returned if the container cannot be found.
See also [ListContainers](ListContainers) and [InspectContainer](#InspectContainer).
### <a name="GetContainerLogs"></a>func GetContainerLogs
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method GetContainerLogs(name: [string](https://godoc.org/builtin#string)) [[]string](#[]string)</div>
GetContainerLogs takes a name or ID of a container and returns the logs of that container.
If the container cannot be found, a [ContainerNotFound](#ContainerNotFound) error will be returned.
The container logs are returned as an array of strings.  GetContainerLogs will honor the streaming
capability of varlink if the client invokes it.
### <a name="GetContainerStats"></a>func GetContainerStats
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method GetContainerStats(name: [string](https://godoc.org/builtin#string)) [ContainerStats](#ContainerStats)</div>
GetContainerStats takes the name or ID of a container and returns a single ContainerStats structure which
contains attributes like memory and cpu usage.  If the container cannot be found, a
[ContainerNotFound](#ContainerNotFound) error will be returned. If the container is not running, a [NoContainerRunning](#NoContainerRunning)
error will be returned
#### Example
~~~
$ varlink call -m unix:/run/podman/io.podman/io.podman.GetContainerStats '{"name": "c33e4164f384"}'
{
  "container": {
    "block_input": 0,
    "block_output": 0,
    "cpu": 2.571123918839990154678e-08,
    "cpu_nano": 49037378,
    "id": "c33e4164f384aa9d979072a63319d66b74fd7a128be71fa68ede24f33ec6cfee",
    "mem_limit": 33080606720,
    "mem_perc": 2.166828456524753747370e-03,
    "mem_usage": 716800,
    "name": "competent_wozniak",
    "net_input": 768,
    "net_output": 5910,
    "pids": 1,
    "system_nano": 10000000
  }
}
~~~
### <a name="GetImage"></a>func GetImage
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method GetImage(name: [string](https://godoc.org/builtin#string)) [ImageInList](#ImageInList)</div>
GetImage returns a single image in an [ImageInList](#ImageInList) struct.  You must supply an image name as a string.
If the image cannot be found, an [ImageNotFound](#ImageNotFound) error will be returned.
### <a name="GetInfo"></a>func GetInfo
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method GetInfo() [PodmanInfo](#PodmanInfo)</div>
GetInfo returns a [PodmanInfo](#PodmanInfo) struct that describes podman and its host such as storage stats,
build information of Podman, and system-wide registries.
### <a name="GetPod"></a>func GetPod
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method GetPod(name: [string](https://godoc.org/builtin#string)) [ListPodData](#ListPodData)</div>
GetPod takes a name or ID of a pod and returns single [ListPodData](#ListPodData)
structure.  A [PodNotFound](#PodNotFound) error will be returned if the pod cannot be found.
See also [ListPods](ListPods).
#### Example
~~~
$ varlink call -m unix:/run/podman/io.podman/io.podman.GetPod '{"name": "foobar"}'
{
  "pod": {
    "cgroup": "machine.slice",
    "containersinfo": [
      {
        "id": "00c130a45de0411f109f1a0cfea2e298df71db20fa939de5cab8b2160a36be45",
        "name": "1840835294cf-infra",
        "status": "running"
      },
      {
        "id": "49a5cce72093a5ca47c6de86f10ad7bb36391e2d89cef765f807e460865a0ec6",
        "name": "upbeat_murdock",
        "status": "running"
      }
    ],
    "createdat": "2018-12-07 13:10:15.014139258 -0600 CST",
    "id": "1840835294cf076a822e4e12ba4152411f131bd869e7f6a4e8b16df9b0ea5c7f",
    "name": "foobar",
    "numberofcontainers": "2",
    "status": "Running"
  }
}
~~~
### <a name="GetPodStats"></a>func GetPodStats
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method GetPodStats(name: [string](https://godoc.org/builtin#string)) [string](https://godoc.org/builtin#string), [ContainerStats](#ContainerStats)</div>
GetPodStats takes the name or ID of a pod and returns a pod name and slice of ContainerStats structure which
contains attributes like memory and cpu usage.  If the pod cannot be found, a [PodNotFound](#PodNotFound)
error will be returned.  If the pod has no running containers associated with it, a [NoContainerRunning](#NoContainerRunning)
error will be returned.
#### Example
~~~
$ varlink call unix:/run/podman/io.podman/io.podman.GetPodStats '{"name": "7f62b508b6f12b11d8fe02e"}'
{
  "containers": [
    {
      "block_input": 0,
      "block_output": 0,
      "cpu": 2.833470544016107524276e-08,
      "cpu_nano": 54363072,
      "id": "a64b51f805121fe2c5a3dc5112eb61d6ed139e3d1c99110360d08b58d48e4a93",
      "mem_limit": 12276146176,
      "mem_perc": 7.974359265237864966003e-03,
      "mem_usage": 978944,
      "name": "quirky_heisenberg",
      "net_input": 866,
      "net_output": 7388,
      "pids": 1,
      "system_nano": 20000000
    }
  ],
  "pod": "7f62b508b6f12b11d8fe02e0db4de6b9e43a7d7699b33a4fc0d574f6e82b4ebd"
}
~~~
### <a name="GetVersion"></a>func GetVersion
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method GetVersion() [Version](#Version)</div>
GetVersion returns a Version structure describing the libpod setup on their
system.
### <a name="HistoryImage"></a>func HistoryImage
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method HistoryImage(name: [string](https://godoc.org/builtin#string)) [ImageHistory](#ImageHistory)</div>
HistoryImage takes the name or ID of an image and returns information about its history and layers.  The returned
history is in the form of an array of ImageHistory structures.  If the image cannot be found, an
[ImageNotFound](#ImageNotFound) error is returned.
### <a name="ImageExists"></a>func ImageExists
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method ImageExists(name: [string](https://godoc.org/builtin#string)) [int](https://godoc.org/builtin#int)</div>
ImageExists talks a full or partial image ID or name and returns an int as to whether
the image exists in local storage. An int result of 0 means the image does exist in
local storage; whereas 1 indicates the image does not exists in local storage.
#### Example
~~~
$ varlink call -m unix:/run/podman/io.podman/io.podman.ImageExists '{"name": "imageddoesntexist"}'
{
  "exists": 1
}
~~~
### <a name="ImportImage"></a>func ImportImage
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method ImportImage(source: [string](https://godoc.org/builtin#string), reference: [string](https://godoc.org/builtin#string), message: [string](https://godoc.org/builtin#string), changes: [[]string](#[]string)) [string](https://godoc.org/builtin#string)</div>
ImportImage imports an image from a source (like tarball) into local storage.  The image can have additional
descriptions added to it using the message and changes options. See also [ExportImage](ExportImage).
### <a name="InspectContainer"></a>func InspectContainer
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method InspectContainer(name: [string](https://godoc.org/builtin#string)) [string](https://godoc.org/builtin#string)</div>
InspectContainer data takes a name or ID of a container returns the inspection
data in string format.  You can then serialize the string into JSON.  A [ContainerNotFound](#ContainerNotFound)
error will be returned if the container cannot be found. See also [InspectImage](#InspectImage).
### <a name="InspectImage"></a>func InspectImage
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method InspectImage(name: [string](https://godoc.org/builtin#string)) [string](https://godoc.org/builtin#string)</div>
InspectImage takes the name or ID of an image and returns a string respresentation of data associated with the
mage.  You must serialize the string into JSON to use it further.  An [ImageNotFound](#ImageNotFound) error will
be returned if the image cannot be found.
### <a name="InspectPod"></a>func InspectPod
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method InspectPod(name: [string](https://godoc.org/builtin#string)) [string](https://godoc.org/builtin#string)</div>
InspectPod takes the name or ID of an image and returns a string respresentation of data associated with the
pod.  You must serialize the string into JSON to use it further.  A [PodNotFound](#PodNotFound) error will
be returned if the pod cannot be found.
### <a name="KillContainer"></a>func KillContainer
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method KillContainer(name: [string](https://godoc.org/builtin#string), signal: [int](https://godoc.org/builtin#int)) [string](https://godoc.org/builtin#string)</div>
KillContainer takes the name or ID of a container as well as a signal to be applied to the container.  Once the
container has been killed, the container's ID is returned.  If the container cannot be found, a
[ContainerNotFound](#ContainerNotFound) error is returned. See also [StopContainer](StopContainer).
### <a name="KillPod"></a>func KillPod
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method KillPod(name: [string](https://godoc.org/builtin#string), signal: [int](https://godoc.org/builtin#int)) [string](https://godoc.org/builtin#string)</div>
KillPod takes the name or ID of a pod as well as a signal to be applied to the pod.  If the pod cannot be found, a
[PodNotFound](#PodNotFound) error is returned.
Containers in a pod are killed independently. If there is an error killing one container, the ID of those containers
will be returned in a list, along with the ID of the pod in a [PodContainerError](#PodContainerError).
If the pod was killed with no errors, the pod ID is returned.
See also [StopPod](StopPod).
#### Example
~~~
$ varlink call -m unix:/run/podman/io.podman/io.podman.KillPod '{"name": "foobar", "signal": 15}'
{
  "pod": "1840835294cf076a822e4e12ba4152411f131bd869e7f6a4e8b16df9b0ea5c7f"
}
~~~
### <a name="ListContainerChanges"></a>func ListContainerChanges
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method ListContainerChanges(name: [string](https://godoc.org/builtin#string)) [ContainerChanges](#ContainerChanges)</div>
ListContainerChanges takes a name or ID of a container and returns changes between the container and
its base image. It returns a struct of changed, deleted, and added path names.
### <a name="ListContainerMounts"></a>func ListContainerMounts
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method ListContainerMounts() [[]string](#[]string)</div>
ListContainerMounts gathers all the mounted container mount points and returns them as an array
of strings
#### Example
~~~
$ varlink call -m unix:/run/podman/io.podman/io.podman.ListContainerMounts
{
  "mounts": [
    "/var/lib/containers/storage/overlay/b215fb622c65ba3b06c6d2341be80b76a9de7ae415ce419e65228873d4f0dcc8/merged",
    "/var/lib/containers/storage/overlay/5eaf806073f79c0ed9a695180ad598e34f963f7407da1d2ccf3560bdab49b26f/merged",
    "/var/lib/containers/storage/overlay/1ecb6b1dbb251737c7a24a31869096839c3719d8b250bf075f75172ddcc701e1/merged",
    "/var/lib/containers/storage/overlay/7137b28a3c422165fe920cba851f2f8da271c6b5908672c451ebda03ad3919e2/merged"
  ]
}
~~~
### <a name="ListContainerPorts"></a>func ListContainerPorts
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method ListContainerPorts(name: [string](https://godoc.org/builtin#string)) [NotImplemented](#NotImplemented)</div>
This function is not implemented yet.
### <a name="ListContainerProcesses"></a>func ListContainerProcesses
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method ListContainerProcesses(name: [string](https://godoc.org/builtin#string), opts: [[]string](#[]string)) [[]string](#[]string)</div>
ListContainerProcesses takes a name or ID of a container and returns the processes
running inside the container as array of strings.  It will accept an array of string
arguments that represent ps options.  If the container cannot be found, a [ContainerNotFound](#ContainerNotFound)
error will be returned.
#### Example
~~~
$ varlink call -m unix:/run/podman/io.podman/io.podman.ListContainerProcesses '{"name": "135d71b9495f", "opts": []}'
{
  "container": [
    "  UID   PID  PPID  C STIME TTY          TIME CMD",
    "    0 21220 21210  0 09:05 pts/0    00:00:00 /bin/sh",
    "    0 21232 21220  0 09:05 pts/0    00:00:00 top",
    "    0 21284 21220  0 09:05 pts/0    00:00:00 vi /etc/hosts"
  ]
}
~~~
### <a name="ListContainers"></a>func ListContainers
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method ListContainers() [ListContainerData](#ListContainerData)</div>
ListContainers returns a list of containers in no particular order.  There are
returned as an array of ListContainerData structs.  See also [GetContainer](#GetContainer).
### <a name="ListImages"></a>func ListImages
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method ListImages() [ImageInList](#ImageInList)</div>
ListImages returns an array of ImageInList structures which provide basic information about
an image currently in storage.  See also [InspectImage](InspectImage).
### <a name="ListPods"></a>func ListPods
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method ListPods() [ListPodData](#ListPodData)</div>
ListPods returns a list of pods in no particular order.  They are
returned as an array of ListPodData structs.  See also [GetPod](#GetPod).
#### Example
~~~
$ varlink call -m unix:/run/podman/io.podman/io.podman.ListPods
{
  "pods": [
    {
      "cgroup": "machine.slice",
      "containersinfo": [
        {
          "id": "00c130a45de0411f109f1a0cfea2e298df71db20fa939de5cab8b2160a36be45",
          "name": "1840835294cf-infra",
          "status": "running"
        },
        {
          "id": "49a5cce72093a5ca47c6de86f10ad7bb36391e2d89cef765f807e460865a0ec6",
          "name": "upbeat_murdock",
          "status": "running"
        }
      ],
      "createdat": "2018-12-07 13:10:15.014139258 -0600 CST",
      "id": "1840835294cf076a822e4e12ba4152411f131bd869e7f6a4e8b16df9b0ea5c7f",
      "name": "foobar",
      "numberofcontainers": "2",
      "status": "Running"
    },
    {
      "cgroup": "machine.slice",
      "containersinfo": [
        {
          "id": "1ca4b7bbba14a75ba00072d4b705c77f3df87db0109afaa44d50cb37c04a477e",
          "name": "784306f655c6-infra",
          "status": "running"
        }
      ],
      "createdat": "2018-12-07 13:09:57.105112457 -0600 CST",
      "id": "784306f655c6200aea321dd430ba685e9b2cc1f7d7528a72f3ff74ffb29485a2",
      "name": "nostalgic_pike",
      "numberofcontainers": "1",
      "status": "Running"
    }
  ]
}
~~~
### <a name="MountContainer"></a>func MountContainer
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method MountContainer(name: [string](https://godoc.org/builtin#string)) [string](https://godoc.org/builtin#string)</div>
MountContainer mounts a container by name or full/partial ID.  Upon a successful mount, the destination
mount is returned as a string.
#### Example
~~~
$ varlink call -m unix:/run/podman/io.podman/io.podman.MountContainer '{"name": "jolly_shannon"}'{
  "path": "/var/lib/containers/storage/overlay/419eeb04e783ea159149ced67d9fcfc15211084d65e894792a96bedfae0470ca/merged"
}
~~~
### <a name="PauseContainer"></a>func PauseContainer
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method PauseContainer(name: [string](https://godoc.org/builtin#string)) [string](https://godoc.org/builtin#string)</div>
PauseContainer takes the name or ID of container and pauses it.  If the container cannot be found,
a [ContainerNotFound](#ContainerNotFound) error will be returned; otherwise the ID of the container is returned.
See also [UnpauseContainer](#UnpauseContainer).
### <a name="PausePod"></a>func PausePod
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method PausePod(name: [string](https://godoc.org/builtin#string)) [string](https://godoc.org/builtin#string)</div>
PausePod takes the name or ID of a pod and pauses the running containers associated with it.  If the pod cannot be found,
a [PodNotFound](#PodNotFound) error will be returned.
Containers in a pod are paused independently. If there is an error pausing one container, the ID of those containers
will be returned in a list, along with the ID of the pod in a [PodContainerError](#PodContainerError).
If the pod was paused with no errors, the pod ID is returned.
See also [UnpausePod](#UnpausePod).
#### Example
~~~
$ varlink call -m unix:/run/podman/io.podman/io.podman.PausePod '{"name": "foobar"}'
{
  "pod": "1840835294cf076a822e4e12ba4152411f131bd869e7f6a4e8b16df9b0ea5c7f"
}
~~~
### <a name="Ping"></a>func Ping
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method Ping() [StringResponse](#StringResponse)</div>
Ping provides a response for developers to ensure their varlink setup is working.
#### Example
~~~
$ varlink call -m unix:/run/podman/io.podman/io.podman.Ping
{
  "ping": {
    "message": "OK"
  }
}
~~~
### <a name="PullImage"></a>func PullImage
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method PullImage(name: [string](https://godoc.org/builtin#string), certDir: [string](https://godoc.org/builtin#string), creds: [string](https://godoc.org/builtin#string), signaturePolicy: [string](https://godoc.org/builtin#string), tlsVerify: [bool](https://godoc.org/builtin#bool)) [string](https://godoc.org/builtin#string)</div>
PullImage pulls an image from a repository to local storage.  After the pull is successful, the ID of the image
is returned.
#### Example
~~~
$ varlink call -m unix:/run/podman/io.podman/io.podman.PullImage '{"name": "registry.fedoraproject.org/fedora"}'
{
  "id": "426866d6fa419873f97e5cbd320eeb22778244c1dfffa01c944db3114f55772e"
}
~~~
### <a name="PushImage"></a>func PushImage
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method PushImage(name: [string](https://godoc.org/builtin#string), tag: [string](https://godoc.org/builtin#string), tlsverify: [bool](https://godoc.org/builtin#bool), signaturePolicy: [string](https://godoc.org/builtin#string), creds: [string](https://godoc.org/builtin#string), certDir: [string](https://godoc.org/builtin#string), compress: [bool](https://godoc.org/builtin#bool), format: [string](https://godoc.org/builtin#string), removeSignatures: [bool](https://godoc.org/builtin#bool), signBy: [string](https://godoc.org/builtin#string)) [string](https://godoc.org/builtin#string)</div>
PushImage takes three input arguments: the name or ID of an image, the fully-qualified destination name of the image,
and a boolean as to whether tls-verify should be used (with false disabling TLS, not affecting the default behavior).
It will return an [ImageNotFound](#ImageNotFound) error if
the image cannot be found in local storage; otherwise the ID of the image will be returned on success.
### <a name="RemoveContainer"></a>func RemoveContainer
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method RemoveContainer(name: [string](https://godoc.org/builtin#string), force: [bool](https://godoc.org/builtin#bool)) [string](https://godoc.org/builtin#string)</div>
RemoveContainer takes requires the name or ID of container as well a boolean representing whether a running
container can be stopped and removed.  Upon successful removal of the container, its ID is returned.  If the
container cannot be found by name or ID, a [ContainerNotFound](#ContainerNotFound) error will be returned.
#### Example
~~~
$ varlink call -m unix:/run/podman/io.podman/io.podman.RemoveContainer '{"name": "62f4fd98cb57"}'
{
  "container": "62f4fd98cb57f529831e8f90610e54bba74bd6f02920ffb485e15376ed365c20"
}
~~~
### <a name="RemoveImage"></a>func RemoveImage
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method RemoveImage(name: [string](https://godoc.org/builtin#string), force: [bool](https://godoc.org/builtin#bool)) [string](https://godoc.org/builtin#string)</div>
RemoveImage takes the name or ID of an image as well as a boolean that determines if containers using that image
should be deleted.  If the image cannot be found, an [ImageNotFound](#ImageNotFound) error will be returned.  The
ID of the removed image is returned when complete.  See also [DeleteUnusedImages](DeleteUnusedImages).
#### Example
~~~
varlink call -m unix:/run/podman/io.podman/io.podman.RemoveImage '{"name": "registry.fedoraproject.org/fedora", "force": true}'
{
  "image": "426866d6fa419873f97e5cbd320eeb22778244c1dfffa01c944db3114f55772e"
}
~~~
### <a name="RemovePod"></a>func RemovePod
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method RemovePod(name: [string](https://godoc.org/builtin#string), force: [bool](https://godoc.org/builtin#bool)) [string](https://godoc.org/builtin#string)</div>
RemovePod takes the name or ID of a pod as well a boolean representing whether a running
container in the pod can be stopped and removed.  If a pod has containers associated with it, and force is not true,
an error will occur.
If the pod cannot be found by name or ID, a [PodNotFound](#PodNotFound) error will be returned.
Containers in a pod are removed independently. If there is an error removing any container, the ID of those containers
will be returned in a list, along with the ID of the pod in a [PodContainerError](#PodContainerError).
If the pod was removed with no errors, the pod ID is returned.
#### Example
~~~
$ varlink call -m unix:/run/podman/io.podman/io.podman.RemovePod '{"name": "62f4fd98cb57", "force": "true"}'
{
  "pod": "62f4fd98cb57f529831e8f90610e54bba74bd6f02920ffb485e15376ed365c20"
}
~~~
### <a name="RenameContainer"></a>func RenameContainer
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method RenameContainer() [NotImplemented](#NotImplemented)</div>
This method has not be implemented yet.
### <a name="ResizeContainerTty"></a>func ResizeContainerTty
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method ResizeContainerTty() [NotImplemented](#NotImplemented)</div>
This method has not be implemented yet.
### <a name="RestartContainer"></a>func RestartContainer
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method RestartContainer(name: [string](https://godoc.org/builtin#string), timeout: [int](https://godoc.org/builtin#int)) [string](https://godoc.org/builtin#string)</div>
RestartContainer will restart a running container given a container name or ID and timeout value. The timeout
value is the time before a forcible stop is used to stop the container.  If the container cannot be found by
name or ID, a [ContainerNotFound](#ContainerNotFound)  error will be returned; otherwise, the ID of the
container will be returned.
### <a name="RestartPod"></a>func RestartPod
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method RestartPod(name: [string](https://godoc.org/builtin#string)) [string](https://godoc.org/builtin#string)</div>
RestartPod will restart containers in a pod given a pod name or ID. Containers in
the pod that are running will be stopped, then all stopped containers will be run.
If the pod cannot be found by name or ID, a [PodNotFound](#PodNotFound) error will be returned.
Containers in a pod are restarted independently. If there is an error restarting one container, the ID of those containers
will be returned in a list, along with the ID of the pod in a [PodContainerError](#PodContainerError).
If the pod was restarted with no errors, the pod ID is returned.
#### Example
~~~
$ varlink call -m unix:/run/podman/io.podman/io.podman.RestartPod '{"name": "135d71b9495f"}'
{
  "pod": "135d71b9495f7c3967f536edad57750bfdb569336cd107d8aabab45565ffcfb6"
}
~~~
### <a name="SearchImage"></a>func SearchImage
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method SearchImage(name: [string](https://godoc.org/builtin#string), limit: [int](https://godoc.org/builtin#int)) [ImageSearch](#ImageSearch)</div>
SearchImage takes the string of an image name and a limit of searches from each registries to be returned.  SearchImage
will then use a glob-like match to find the image you are searching for.  The images are returned in an array of
ImageSearch structures which contain information about the image as well as its fully-qualified name.
### <a name="StartContainer"></a>func StartContainer
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method StartContainer(name: [string](https://godoc.org/builtin#string)) [string](https://godoc.org/builtin#string)</div>
StartContainer starts a created or stopped container. It takes the name or ID of container.  It returns
the container ID once started.  If the container cannot be found, a [ContainerNotFound](#ContainerNotFound)
error will be returned.  See also [CreateContainer](#CreateContainer).
### <a name="StartPod"></a>func StartPod
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method StartPod(name: [string](https://godoc.org/builtin#string)) [string](https://godoc.org/builtin#string)</div>
StartPod starts containers in a pod.  It takes the name or ID of pod.  If the pod cannot be found, a [PodNotFound](#PodNotFound)
error will be returned.  Containers in a pod are started independently. If there is an error starting one container, the ID of those containers
will be returned in a list, along with the ID of the pod in a [PodContainerError](#PodContainerError).
If the pod was started with no errors, the pod ID is returned.
See also [CreatePod](#CreatePod).
#### Example
~~~
$ varlink call -m unix:/run/podman/io.podman/io.podman.StartPod '{"name": "135d71b9495f"}'
{
  "pod": "135d71b9495f7c3967f536edad57750bfdb569336cd107d8aabab45565ffcfb6",
}
~~~
### <a name="StopContainer"></a>func StopContainer
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method StopContainer(name: [string](https://godoc.org/builtin#string), timeout: [int](https://godoc.org/builtin#int)) [string](https://godoc.org/builtin#string)</div>
StopContainer stops a container given a timeout.  It takes the name or ID of a container as well as a
timeout value.  The timeout value the time before a forcible stop to the container is applied.  It
returns the container ID once stopped. If the container cannot be found, a [ContainerNotFound](#ContainerNotFound)
error will be returned instead. See also [KillContainer](KillContainer).
#### Error
~~~
$ varlink call -m unix:/run/podman/io.podman/io.podman.StopContainer '{"name": "135d71b9495f", "timeout": 5}'
{
  "container": "135d71b9495f7c3967f536edad57750bfdb569336cd107d8aabab45565ffcfb6"
}
~~~
### <a name="StopPod"></a>func StopPod
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method StopPod(name: [string](https://godoc.org/builtin#string), timeout: [int](https://godoc.org/builtin#int)) [string](https://godoc.org/builtin#string)</div>
StopPod stops containers in a pod.  It takes the name or ID of a pod and a timeout.
If the pod cannot be found, a [PodNotFound](#PodNotFound) error will be returned instead.
Containers in a pod are stopped independently. If there is an error stopping one container, the ID of those containers
will be returned in a list, along with the ID of the pod in a [PodContainerError](#PodContainerError).
If the pod was stopped with no errors, the pod ID is returned.
See also [KillPod](KillPod).
#### Example
~~~
$ varlink call -m unix:/run/podman/io.podman/io.podman.StopPod '{"name": "135d71b9495f"}'
{
  "pod": "135d71b9495f7c3967f536edad57750bfdb569336cd107d8aabab45565ffcfb6"
}
~~~
### <a name="TagImage"></a>func TagImage
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method TagImage(name: [string](https://godoc.org/builtin#string), tagged: [string](https://godoc.org/builtin#string)) [string](https://godoc.org/builtin#string)</div>
TagImage takes the name or ID of an image in local storage as well as the desired tag name.  If the image cannot
be found, an [ImageNotFound](#ImageNotFound) error will be returned; otherwise, the ID of the image is returned on success.
### <a name="TopPod"></a>func TopPod
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method TopPod() [NotImplemented](#NotImplemented)</div>
This method has not been implemented yet.
### <a name="UnmountContainer"></a>func UnmountContainer
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method UnmountContainer(name: [string](https://godoc.org/builtin#string), force: [bool](https://godoc.org/builtin#bool)) </div>
UnmountContainer umounts a container by its name or full/partial container ID.
#### Example
~~~
$ varlink call -m unix:/run/podman/io.podman/io.podman.UnmountContainer '{"name": "jolly_shannon", "force": false}'
{}
~~~
### <a name="UnpauseContainer"></a>func UnpauseContainer
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method UnpauseContainer(name: [string](https://godoc.org/builtin#string)) [string](https://godoc.org/builtin#string)</div>
UnpauseContainer takes the name or ID of container and unpauses a paused container.  If the container cannot be
found, a [ContainerNotFound](#ContainerNotFound) error will be returned; otherwise the ID of the container is returned.
See also [PauseContainer](#PauseContainer).
### <a name="UnpausePod"></a>func UnpausePod
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method UnpausePod(name: [string](https://godoc.org/builtin#string)) [string](https://godoc.org/builtin#string)</div>
UnpausePod takes the name or ID of a pod and unpauses the paused containers associated with it.  If the pod cannot be
found, a [PodNotFound](#PodNotFound) error will be returned.
Containers in a pod are unpaused independently. If there is an error unpausing one container, the ID of those containers
will be returned in a list, along with the ID of the pod in a [PodContainerError](#PodContainerError).
If the pod was unpaused with no errors, the pod ID is returned.
See also [PausePod](#PausePod).
#### Example
~~~
$ varlink call -m unix:/run/podman/io.podman/io.podman.UnpausePod '{"name": "foobar"}'
{
  "pod": "1840835294cf076a822e4e12ba4152411f131bd869e7f6a4e8b16df9b0ea5c7f"
}
~~~
### <a name="UpdateContainer"></a>func UpdateContainer
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method UpdateContainer() [NotImplemented](#NotImplemented)</div>
This method has not be implemented yet.
### <a name="WaitContainer"></a>func WaitContainer
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method WaitContainer(name: [string](https://godoc.org/builtin#string)) [int](https://godoc.org/builtin#int)</div>
WaitContainer takes the name or ID of a container and waits until the container stops.  Upon stopping, the return
code of the container is returned. If the container container cannot be found by ID or name,
a [ContainerNotFound](#ContainerNotFound) error is returned.
### <a name="WaitPod"></a>func WaitPod
<div style="background-color: #E8E8E8; padding: 15px; margin: 10px; border-radius: 10px;">

method WaitPod() [NotImplemented](#NotImplemented)</div>
This method has not be implemented yet.
## Types
### <a name="BuildInfo"></a>type BuildInfo

BuildInfo is used to describe user input for building images

dockerfile [[]string](#[]string)

tags [[]string](#[]string)

add_hosts [[]string](#[]string)

cgroup_parent [string](https://godoc.org/builtin#string)

cpu_period [int](https://godoc.org/builtin#int)

cpu_quota [int](https://godoc.org/builtin#int)

cpu_shares [int](https://godoc.org/builtin#int)

cpuset_cpus [string](https://godoc.org/builtin#string)

cpuset_mems [string](https://godoc.org/builtin#string)

memory [string](https://godoc.org/builtin#string)

memory_swap [string](https://godoc.org/builtin#string)

security_opts [[]string](#[]string)

shm_size [string](https://godoc.org/builtin#string)

ulimit [[]string](#[]string)

volume [[]string](#[]string)

squash [bool](https://godoc.org/builtin#bool)

pull [bool](https://godoc.org/builtin#bool)

pull_always [bool](https://godoc.org/builtin#bool)

force_rm [bool](https://godoc.org/builtin#bool)

rm [bool](https://godoc.org/builtin#bool)

label [[]string](#[]string)

annotations [[]string](#[]string)

build_args [map[string]](#map[string])

image_format [string](https://godoc.org/builtin#string)
### <a name="BuildResponse"></a>type BuildResponse

BuildResponse is used to describe the responses for building images

logs [[]string](#[]string)

id [string](https://godoc.org/builtin#string)
### <a name="ContainerChanges"></a>type ContainerChanges

ContainerChanges describes the return struct for ListContainerChanges

changed [[]string](#[]string)

added [[]string](#[]string)

deleted [[]string](#[]string)
### <a name="ContainerMount"></a>type ContainerMount

ContainerMount describes the struct for mounts in a container

destination [string](https://godoc.org/builtin#string)

type [string](https://godoc.org/builtin#string)

source [string](https://godoc.org/builtin#string)

options [[]string](#[]string)
### <a name="ContainerNameSpace"></a>type ContainerNameSpace

ContainerNamespace describes the namespace structure for an existing container

user [string](https://godoc.org/builtin#string)

uts [string](https://godoc.org/builtin#string)

pidns [string](https://godoc.org/builtin#string)

pid [string](https://godoc.org/builtin#string)

cgroup [string](https://godoc.org/builtin#string)

net [string](https://godoc.org/builtin#string)

mnt [string](https://godoc.org/builtin#string)

ipc [string](https://godoc.org/builtin#string)
### <a name="ContainerPortMappings"></a>type ContainerPortMappings

ContainerPortMappings describes the struct for portmappings in an existing container

host_port [string](https://godoc.org/builtin#string)

host_ip [string](https://godoc.org/builtin#string)

protocol [string](https://godoc.org/builtin#string)

container_port [string](https://godoc.org/builtin#string)
### <a name="ContainerStats"></a>type ContainerStats

ContainerStats is the return struct for the stats of a container

id [string](https://godoc.org/builtin#string)

name [string](https://godoc.org/builtin#string)

cpu [float](https://golang.org/src/builtin/builtin.go#L58)

cpu_nano [int](https://godoc.org/builtin#int)

system_nano [int](https://godoc.org/builtin#int)

mem_usage [int](https://godoc.org/builtin#int)

mem_limit [int](https://godoc.org/builtin#int)

mem_perc [float](https://golang.org/src/builtin/builtin.go#L58)

net_input [int](https://godoc.org/builtin#int)

net_output [int](https://godoc.org/builtin#int)

block_output [int](https://godoc.org/builtin#int)

block_input [int](https://godoc.org/builtin#int)

pids [int](https://godoc.org/builtin#int)
### <a name="Create"></a>type Create

Create is an input structure for creating containers. It closely resembles the
CreateConfig structure in libpod/pkg/spec.

args [[]string](#[]string)

cap_add [[]string](#[]string)

cap_drop [[]string](#[]string)

conmon_pidfile [string](https://godoc.org/builtin#string)

cgroup_parent [string](https://godoc.org/builtin#string)

command [[]string](#[]string)

detach [bool](https://godoc.org/builtin#bool)

devices [[]string](#[]string)

dns_opt [[]string](#[]string)

dns_search [[]string](#[]string)

dns_servers [[]string](#[]string)

entrypoint [[]string](#[]string)

env [map[string]](#map[string])

exposed_ports [[]string](#[]string)

gidmap [[]string](#[]string)

group_add [[]string](#[]string)

host_add [[]string](#[]string)

hostname [string](https://godoc.org/builtin#string)

image [string](https://godoc.org/builtin#string)

image_id [string](https://godoc.org/builtin#string)

builtin_imgvolumes [[]string](#[]string)

id_mappings [IDMappingOptions](#IDMappingOptions)

image_volume_type [string](https://godoc.org/builtin#string)

interactive [bool](https://godoc.org/builtin#bool)

ipc_mode [string](https://godoc.org/builtin#string)

labels [map[string]](#map[string])

log_driver [string](https://godoc.org/builtin#string)

log_driver_opt [[]string](#[]string)

name [string](https://godoc.org/builtin#string)

net_mode [string](https://godoc.org/builtin#string)

network [string](https://godoc.org/builtin#string)

pid_mode [string](https://godoc.org/builtin#string)

pod [string](https://godoc.org/builtin#string)

privileged [bool](https://godoc.org/builtin#bool)

publish [[]string](#[]string)

publish_all [bool](https://godoc.org/builtin#bool)

quiet [bool](https://godoc.org/builtin#bool)

readonly_rootfs [bool](https://godoc.org/builtin#bool)

resources [CreateResourceConfig](#CreateResourceConfig)

rm [bool](https://godoc.org/builtin#bool)

shm_dir [string](https://godoc.org/builtin#string)

stop_signal [int](https://godoc.org/builtin#int)

stop_timeout [int](https://godoc.org/builtin#int)

subuidmap [string](https://godoc.org/builtin#string)

subgidmap [string](https://godoc.org/builtin#string)

subuidname [string](https://godoc.org/builtin#string)

subgidname [string](https://godoc.org/builtin#string)

sys_ctl [map[string]](#map[string])

tmpfs [[]string](#[]string)

tty [bool](https://godoc.org/builtin#bool)

uidmap [[]string](#[]string)

userns_mode [string](https://godoc.org/builtin#string)

user [string](https://godoc.org/builtin#string)

uts_mode [string](https://godoc.org/builtin#string)

volumes [[]string](#[]string)

work_dir [string](https://godoc.org/builtin#string)

mount_label [string](https://godoc.org/builtin#string)

process_label [string](https://godoc.org/builtin#string)

no_new_privs [bool](https://godoc.org/builtin#bool)

apparmor_profile [string](https://godoc.org/builtin#string)

seccomp_profile_path [string](https://godoc.org/builtin#string)

security_opts [[]string](#[]string)
### <a name="CreateResourceConfig"></a>type CreateResourceConfig

CreateResourceConfig is an input structure used to describe host attributes during
container creation.  It is only valid inside a [Create](#Create) type.

blkio_weight [int](https://godoc.org/builtin#int)

blkio_weight_device [[]string](#[]string)

cpu_period [int](https://godoc.org/builtin#int)

cpu_quota [int](https://godoc.org/builtin#int)

cpu_rt_period [int](https://godoc.org/builtin#int)

cpu_rt_runtime [int](https://godoc.org/builtin#int)

cpu_shares [int](https://godoc.org/builtin#int)

cpus [float](https://golang.org/src/builtin/builtin.go#L58)

cpuset_cpus [string](https://godoc.org/builtin#string)

cpuset_mems [string](https://godoc.org/builtin#string)

device_read_bps [[]string](#[]string)

device_read_iops [[]string](#[]string)

device_write_bps [[]string](#[]string)

device_write_iops [[]string](#[]string)

disable_oomkiller [bool](https://godoc.org/builtin#bool)

kernel_memory [int](https://godoc.org/builtin#int)

memory [int](https://godoc.org/builtin#int)

memory_reservation [int](https://godoc.org/builtin#int)

memory_swap [int](https://godoc.org/builtin#int)

memory_swappiness [int](https://godoc.org/builtin#int)

oom_score_adj [int](https://godoc.org/builtin#int)

pids_limit [int](https://godoc.org/builtin#int)

shm_size [int](https://godoc.org/builtin#int)

ulimit [[]string](#[]string)
### <a name="IDMap"></a>type IDMap

IDMap is used to describe user name spaces during container creation

container_id [int](https://godoc.org/builtin#int)

host_id [int](https://godoc.org/builtin#int)

size [int](https://godoc.org/builtin#int)
### <a name="IDMappingOptions"></a>type IDMappingOptions

IDMappingOptions is an input structure used to described ids during container creation.

host_uid_mapping [bool](https://godoc.org/builtin#bool)

host_gid_mapping [bool](https://godoc.org/builtin#bool)

uid_map [IDMap](#IDMap)

gid_map [IDMap](#IDMap)
### <a name="ImageHistory"></a>type ImageHistory

ImageHistory describes the returned structure from ImageHistory.

id [string](https://godoc.org/builtin#string)

created [string](https://godoc.org/builtin#string)

createdBy [string](https://godoc.org/builtin#string)

tags [[]string](#[]string)

size [int](https://godoc.org/builtin#int)

comment [string](https://godoc.org/builtin#string)
### <a name="ImageInList"></a>type ImageInList

ImageInList describes the structure that is returned in
ListImages.

id [string](https://godoc.org/builtin#string)

parentId [string](https://godoc.org/builtin#string)

repoTags [[]string](#[]string)

repoDigests [[]string](#[]string)

created [string](https://godoc.org/builtin#string)

size [int](https://godoc.org/builtin#int)

virtualSize [int](https://godoc.org/builtin#int)

containers [int](https://godoc.org/builtin#int)

labels [map[string]](#map[string])
### <a name="ImageSearch"></a>type ImageSearch

ImageSearch is the returned structure for SearchImage.  It is returned
in array form.

description [string](https://godoc.org/builtin#string)

is_official [bool](https://godoc.org/builtin#bool)

is_automated [bool](https://godoc.org/builtin#bool)

name [string](https://godoc.org/builtin#string)

star_count [int](https://godoc.org/builtin#int)
### <a name="InfoDistribution"></a>type InfoDistribution

InfoDistribution describes the the host's distribution

distribution [string](https://godoc.org/builtin#string)

version [string](https://godoc.org/builtin#string)
### <a name="InfoGraphStatus"></a>type InfoGraphStatus

InfoGraphStatus describes the detailed status of the storage driver

backing_filesystem [string](https://godoc.org/builtin#string)

native_overlay_diff [string](https://godoc.org/builtin#string)

supports_d_type [string](https://godoc.org/builtin#string)
### <a name="InfoHost"></a>type InfoHost

InfoHost describes the host stats portion of PodmanInfo

buildah_version [string](https://godoc.org/builtin#string)

distribution [InfoDistribution](#InfoDistribution)

mem_free [int](https://godoc.org/builtin#int)

mem_total [int](https://godoc.org/builtin#int)

swap_free [int](https://godoc.org/builtin#int)

swap_total [int](https://godoc.org/builtin#int)

arch [string](https://godoc.org/builtin#string)

cpus [int](https://godoc.org/builtin#int)

hostname [string](https://godoc.org/builtin#string)

kernel [string](https://godoc.org/builtin#string)

os [string](https://godoc.org/builtin#string)

uptime [string](https://godoc.org/builtin#string)
### <a name="InfoPodmanBinary"></a>type InfoPodmanBinary

InfoPodman provides details on the podman binary

compiler [string](https://godoc.org/builtin#string)

go_version [string](https://godoc.org/builtin#string)

podman_version [string](https://godoc.org/builtin#string)

git_commit [string](https://godoc.org/builtin#string)
### <a name="InfoStore"></a>type InfoStore

InfoStore describes the host's storage informatoin

containers [int](https://godoc.org/builtin#int)

images [int](https://godoc.org/builtin#int)

graph_driver_name [string](https://godoc.org/builtin#string)

graph_driver_options [string](https://godoc.org/builtin#string)

graph_root [string](https://godoc.org/builtin#string)

graph_status [InfoGraphStatus](#InfoGraphStatus)

run_root [string](https://godoc.org/builtin#string)
### <a name="ListContainerData"></a>type ListContainerData

ListContainer is the returned struct for an individual container

id [string](https://godoc.org/builtin#string)

image [string](https://godoc.org/builtin#string)

imageid [string](https://godoc.org/builtin#string)

command [[]string](#[]string)

createdat [string](https://godoc.org/builtin#string)

runningfor [string](https://godoc.org/builtin#string)

status [string](https://godoc.org/builtin#string)

ports [ContainerPortMappings](#ContainerPortMappings)

rootfssize [int](https://godoc.org/builtin#int)

rwsize [int](https://godoc.org/builtin#int)

names [string](https://godoc.org/builtin#string)

labels [map[string]](#map[string])

mounts [ContainerMount](#ContainerMount)

containerrunning [bool](https://godoc.org/builtin#bool)

namespaces [ContainerNameSpace](#ContainerNameSpace)
### <a name="ListPodContainerInfo"></a>type ListPodContainerInfo

ListPodContainerInfo is a returned struct for describing containers
in a pod.

name [string](https://godoc.org/builtin#string)

id [string](https://godoc.org/builtin#string)

status [string](https://godoc.org/builtin#string)
### <a name="ListPodData"></a>type ListPodData

ListPodData is the returned struct for an individual pod

id [string](https://godoc.org/builtin#string)

name [string](https://godoc.org/builtin#string)

createdat [string](https://godoc.org/builtin#string)

cgroup [string](https://godoc.org/builtin#string)

status [string](https://godoc.org/builtin#string)

labels [map[string]](#map[string])

numberofcontainers [string](https://godoc.org/builtin#string)

containersinfo [ListPodContainerInfo](#ListPodContainerInfo)
### <a name="NotImplemented"></a>type NotImplemented



comment [string](https://godoc.org/builtin#string)
### <a name="PodContainerErrorData"></a>type PodContainerErrorData



containerid [string](https://godoc.org/builtin#string)

reason [string](https://godoc.org/builtin#string)
### <a name="PodCreate"></a>type PodCreate

PodCreate is an input structure for creating pods.
It emulates options to podman pod create. The infraCommand and
infraImage options are currently NotSupported.

name [string](https://godoc.org/builtin#string)

cgroupParent [string](https://godoc.org/builtin#string)

labels [map[string]](#map[string])

share [[]string](#[]string)

infra [bool](https://godoc.org/builtin#bool)

infraCommand [string](https://godoc.org/builtin#string)

infraImage [string](https://godoc.org/builtin#string)

publish [[]string](#[]string)
### <a name="PodmanInfo"></a>type PodmanInfo

PodmanInfo describes the Podman host and build

host [InfoHost](#InfoHost)

registries [[]string](#[]string)

insecure_registries [[]string](#[]string)

store [InfoStore](#InfoStore)

podman [InfoPodmanBinary](#InfoPodmanBinary)
### <a name="Runlabel"></a>type Runlabel

Runlabel describes the required input for container runlabel

image [string](https://godoc.org/builtin#string)

authfile [string](https://godoc.org/builtin#string)

certDir [string](https://godoc.org/builtin#string)

creds [string](https://godoc.org/builtin#string)

display [bool](https://godoc.org/builtin#bool)

name [string](https://godoc.org/builtin#string)

pull [bool](https://godoc.org/builtin#bool)

signaturePolicyPath [string](https://godoc.org/builtin#string)

tlsVerify [bool](https://godoc.org/builtin#bool)

label [string](https://godoc.org/builtin#string)

extraArgs [[]string](#[]string)

opts [map[string]](#map[string])
### <a name="Sockets"></a>type Sockets

Sockets describes sockets location for a container

container_id [string](https://godoc.org/builtin#string)

io_socket [string](https://godoc.org/builtin#string)

control_socket [string](https://godoc.org/builtin#string)
### <a name="StringResponse"></a>type StringResponse



message [string](https://godoc.org/builtin#string)
### <a name="Version"></a>type Version

Version is the structure returned by GetVersion

version [string](https://godoc.org/builtin#string)

go_version [string](https://godoc.org/builtin#string)

git_commit [string](https://godoc.org/builtin#string)

built [int](https://godoc.org/builtin#int)

os_arch [string](https://godoc.org/builtin#string)
## Errors
### <a name="ContainerNotFound"></a>type ContainerNotFound

ContainerNotFound means the container could not be found by the provided name or ID in local storage.
### <a name="ErrorOccurred"></a>type ErrorOccurred

ErrorOccurred is a generic error for an error that occurs during the execution.  The actual error message
is includes as part of the error's text.
### <a name="ImageNotFound"></a>type ImageNotFound

ImageNotFound means the image could not be found by the provided name or ID in local storage.
### <a name="NoContainerRunning"></a>type NoContainerRunning

NoContainerRunning means none of the containers requested are running in a command that requires a running container.
### <a name="NoContainersInPod"></a>type NoContainersInPod

NoContainersInPod means a pod has no containers on which to perform the operation. It contains
the pod ID.
### <a name="PodContainerError"></a>type PodContainerError

PodContainerError means a container associated with a pod failed to preform an operation. It contains
a container ID of the container that failed.
### <a name="PodNotFound"></a>type PodNotFound

PodNotFound means the pod could not be found by the provided name or ID in local storage.
### <a name="RuntimeError"></a>type RuntimeError

RuntimeErrors generally means a runtime could not be found or gotten.
