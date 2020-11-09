# Go API client for swagger

This documentation describes the Podman v2.0 RESTful API. It replaces the Podman v1.0 API and was initially delivered along with Podman v2.0.  It consists of a Docker-compatible API and a Libpod API providing support for Podman’s unique features such as pods.  To start the service and keep it running for 5,000 seconds (-t 0 runs forever):  podman system service -t 5000 &  You can then use cURL on the socket using requests documented below.  NOTE: if you install the package podman-docker, it will create a symbolic link for /var/run/docker.sock to /run/podman/podman.sock  See podman-service(1) for more information.  Quick Examples:  'podman info'  curl --unix-socket /run/podman/podman.sock http://d/v1.0.0/libpod/info  'podman pull quay.io/containers/podman'  curl -XPOST --unix-socket /run/podman/podman.sock -v 'http://d/v1.0.0/images/create?fromImage=quay.io%2Fcontainers%2Fpodman'  'podman list images'  curl --unix-socket /run/podman/podman.sock -v 'http://d/v1.0.0/libpod/images/json' | jq

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 0.0.1
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen
For more information, please visit [https://podman.io/community/](https://podman.io/community/)

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *http://podman.io/*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ContainersApi* | [**ChangesContainer**](docs/ContainersApi.md#changescontainer) | **Get** /libpod/containers/{name}/changes | Report on changes to container&#x27;s filesystem; adds, deletes or modifications.
*ContainersApi* | [**LibpodAttachContainer**](docs/ContainersApi.md#libpodattachcontainer) | **Post** /libpod/containers/{name}/attach | Attach to a container
*ContainersApi* | [**LibpodCheckpointContainer**](docs/ContainersApi.md#libpodcheckpointcontainer) | **Post** /libpod/containers/{name}/checkpoint | Checkpoint a container
*ContainersApi* | [**LibpodCommitContainer**](docs/ContainersApi.md#libpodcommitcontainer) | **Post** /libpod/commit | Commit
*ContainersApi* | [**LibpodContainerExists**](docs/ContainersApi.md#libpodcontainerexists) | **Get** /libpod/containers/{name}/exists | Check if container exists
*ContainersApi* | [**LibpodCreateContainer**](docs/ContainersApi.md#libpodcreatecontainer) | **Post** /libpod/containers/create | Create a container
*ContainersApi* | [**LibpodExportContainer**](docs/ContainersApi.md#libpodexportcontainer) | **Get** /libpod/containers/{name}/export | Export a container
*ContainersApi* | [**LibpodGenerateKube**](docs/ContainersApi.md#libpodgeneratekube) | **Get** /libpod/generate/{name:.*}/kube | Generate a Kubernetes YAML file.
*ContainersApi* | [**LibpodGenerateSystemd**](docs/ContainersApi.md#libpodgeneratesystemd) | **Get** /libpod/generate/{name:.*}/systemd | Generate Systemd Units
*ContainersApi* | [**LibpodGetContainer**](docs/ContainersApi.md#libpodgetcontainer) | **Get** /libpod/containers/{name}/json | Inspect container
*ContainersApi* | [**LibpodInitContainer**](docs/ContainersApi.md#libpodinitcontainer) | **Post** /libpod/containers/{name}/init | Initialize a container
*ContainersApi* | [**LibpodKillContainer**](docs/ContainersApi.md#libpodkillcontainer) | **Post** /libpod/containers/{name}/kill | Kill container
*ContainersApi* | [**LibpodListContainers**](docs/ContainersApi.md#libpodlistcontainers) | **Get** /libpod/containers/json | List containers
*ContainersApi* | [**LibpodLogsFromContainer**](docs/ContainersApi.md#libpodlogsfromcontainer) | **Get** /libpod/containers/{name}/logs | Get container logs
*ContainersApi* | [**LibpodMountContainer**](docs/ContainersApi.md#libpodmountcontainer) | **Post** /libpod/containers/{name}/mount | Mount a container
*ContainersApi* | [**LibpodPauseContainer**](docs/ContainersApi.md#libpodpausecontainer) | **Post** /libpod/containers/{name}/pause | Pause a container
*ContainersApi* | [**LibpodPlayKube**](docs/ContainersApi.md#libpodplaykube) | **Post** /libpod/play/kube | Play a Kubernetes YAML file.
*ContainersApi* | [**LibpodPruneContainers**](docs/ContainersApi.md#libpodprunecontainers) | **Post** /libpod/containers/prune | Delete stopped containers
*ContainersApi* | [**LibpodPutArchive**](docs/ContainersApi.md#libpodputarchive) | **Post** /libpod/containers/{name}/copy | Copy files into a container
*ContainersApi* | [**LibpodRemoveContainer**](docs/ContainersApi.md#libpodremovecontainer) | **Delete** /libpod/containers/{name} | Delete container
*ContainersApi* | [**LibpodResizeContainer**](docs/ContainersApi.md#libpodresizecontainer) | **Post** /libpod/containers/{name}/resize | Resize a container&#x27;s TTY
*ContainersApi* | [**LibpodRestartContainer**](docs/ContainersApi.md#libpodrestartcontainer) | **Post** /libpod/containers/{name}/restart | Restart a container
*ContainersApi* | [**LibpodRestoreContainer**](docs/ContainersApi.md#libpodrestorecontainer) | **Post** /libpod/containers/{name}/restore | Restore a container
*ContainersApi* | [**LibpodRunHealthCheck**](docs/ContainersApi.md#libpodrunhealthcheck) | **Get** /libpod/containers/{name:.*}/healthcheck | Run a container&#x27;s healthcheck
*ContainersApi* | [**LibpodShowMountedContainers**](docs/ContainersApi.md#libpodshowmountedcontainers) | **Get** /libpod/containers/showmounted | Show mounted containers
*ContainersApi* | [**LibpodStartContainer**](docs/ContainersApi.md#libpodstartcontainer) | **Post** /libpod/containers/{name}/start | Start a container
*ContainersApi* | [**LibpodStatsContainer**](docs/ContainersApi.md#libpodstatscontainer) | **Get** /libpod/containers/{name}/stats | Get stats for a container
*ContainersApi* | [**LibpodStatsContainers**](docs/ContainersApi.md#libpodstatscontainers) | **Get** /libpod/containers/stats | Get stats for one or more containers
*ContainersApi* | [**LibpodStopContainer**](docs/ContainersApi.md#libpodstopcontainer) | **Post** /libpod/containers/{name}/stop | Stop a container
*ContainersApi* | [**LibpodTopContainer**](docs/ContainersApi.md#libpodtopcontainer) | **Get** /libpod/containers/{name}/top | List processes
*ContainersApi* | [**LibpodUnmountContainer**](docs/ContainersApi.md#libpodunmountcontainer) | **Post** /libpod/containers/{name}/unmount | Unmount a container
*ContainersApi* | [**LibpodUnpauseContainer**](docs/ContainersApi.md#libpodunpausecontainer) | **Post** /libpod/containers/{name}/unpause | Unpause Container
*ContainersApi* | [**LibpodWaitContainer**](docs/ContainersApi.md#libpodwaitcontainer) | **Post** /libpod/containers/{name}/wait | Wait on a container
*ContainersCompatApi* | [**AttachContainer**](docs/ContainersCompatApi.md#attachcontainer) | **Post** /containers/{name}/attach | Attach to a container
*ContainersCompatApi* | [**ChangesContainer**](docs/ContainersCompatApi.md#changescontainer) | **Get** /libpod/containers/{name}/changes | Report on changes to container&#x27;s filesystem; adds, deletes or modifications.
*ContainersCompatApi* | [**CommitContainer**](docs/ContainersCompatApi.md#commitcontainer) | **Post** /commit | New Image
*ContainersCompatApi* | [**CreateContainer**](docs/ContainersCompatApi.md#createcontainer) | **Post** /containers/create | Create a container
*ContainersCompatApi* | [**ExportContainer**](docs/ContainersCompatApi.md#exportcontainer) | **Get** /containers/{name}/export | Export a container
*ContainersCompatApi* | [**GetArchive**](docs/ContainersCompatApi.md#getarchive) | **Get** /containers/{name}/archive | Get files from a container
*ContainersCompatApi* | [**GetContainer**](docs/ContainersCompatApi.md#getcontainer) | **Get** /containers/{name}/json | Inspect container
*ContainersCompatApi* | [**KillContainer**](docs/ContainersCompatApi.md#killcontainer) | **Post** /containers/{name}/kill | Kill container
*ContainersCompatApi* | [**LibpodGetArchive**](docs/ContainersCompatApi.md#libpodgetarchive) | **Get** /libpod/containers/{name}/copy | Copy files from a container
*ContainersCompatApi* | [**ListContainers**](docs/ContainersCompatApi.md#listcontainers) | **Get** /containers/json | List containers
*ContainersCompatApi* | [**LogsFromContainer**](docs/ContainersCompatApi.md#logsfromcontainer) | **Get** /containers/{name}/logs | Get container logs
*ContainersCompatApi* | [**PauseContainer**](docs/ContainersCompatApi.md#pausecontainer) | **Post** /containers/{name}/pause | Pause container
*ContainersCompatApi* | [**PruneContainers**](docs/ContainersCompatApi.md#prunecontainers) | **Post** /containers/prune | Delete stopped containers
*ContainersCompatApi* | [**PutArchive**](docs/ContainersCompatApi.md#putarchive) | **Put** /containers/{name}/archive | Put files into a container
*ContainersCompatApi* | [**RemoveContainer**](docs/ContainersCompatApi.md#removecontainer) | **Delete** /containers/{name} | Remove a container
*ContainersCompatApi* | [**ResizeContainer**](docs/ContainersCompatApi.md#resizecontainer) | **Post** /containers/{name}/resize | Resize a container&#x27;s TTY
*ContainersCompatApi* | [**RestartContainer**](docs/ContainersCompatApi.md#restartcontainer) | **Post** /containers/{name}/restart | Restart container
*ContainersCompatApi* | [**StartContainer**](docs/ContainersCompatApi.md#startcontainer) | **Post** /containers/{name}/start | Start a container
*ContainersCompatApi* | [**StatsContainer**](docs/ContainersCompatApi.md#statscontainer) | **Get** /containers/{name}/stats | Get stats for a container
*ContainersCompatApi* | [**StopContainer**](docs/ContainersCompatApi.md#stopcontainer) | **Post** /containers/{name}/stop | Stop a container
*ContainersCompatApi* | [**TopContainer**](docs/ContainersCompatApi.md#topcontainer) | **Get** /containers/{name}/top | List processes running inside a container
*ContainersCompatApi* | [**UnpauseContainer**](docs/ContainersCompatApi.md#unpausecontainer) | **Post** /containers/{name}/unpause | Unpause container
*ContainersCompatApi* | [**WaitContainer**](docs/ContainersCompatApi.md#waitcontainer) | **Post** /containers/{name}/wait | Wait on a container
*ExecApi* | [**LibpodCreateExec**](docs/ExecApi.md#libpodcreateexec) | **Post** /libpod/containers/{name}/exec | Create an exec instance
*ExecApi* | [**LibpodInspectExec**](docs/ExecApi.md#libpodinspectexec) | **Get** /libpod/exec/{id}/json | Inspect an exec instance
*ExecApi* | [**LibpodResizeExec**](docs/ExecApi.md#libpodresizeexec) | **Post** /libpod/exec/{id}/resize | Resize an exec instance
*ExecApi* | [**LibpodStartExec**](docs/ExecApi.md#libpodstartexec) | **Post** /libpod/exec/{id}/start | Start an exec instance
*ExecCompatApi* | [**CreateExec**](docs/ExecCompatApi.md#createexec) | **Post** /containers/{name}/exec | Create an exec instance
*ExecCompatApi* | [**InspectExec**](docs/ExecCompatApi.md#inspectexec) | **Get** /exec/{id}/json | Inspect an exec instance
*ExecCompatApi* | [**ResizeExec**](docs/ExecCompatApi.md#resizeexec) | **Post** /exec/{id}/resize | Resize an exec instance
*ExecCompatApi* | [**StartExec**](docs/ExecCompatApi.md#startexec) | **Post** /exec/{id}/start | Start an exec instance
*ImagesApi* | [**LibpodBuildImage**](docs/ImagesApi.md#libpodbuildimage) | **Post** /libpod/build | Create image
*ImagesApi* | [**LibpodChangesImages**](docs/ImagesApi.md#libpodchangesimages) | **Get** /libpod/images/{name}/changes | Report on changes to images&#x27;s filesystem; adds, deletes or modifications.
*ImagesApi* | [**LibpodExportImage**](docs/ImagesApi.md#libpodexportimage) | **Get** /libpod/images/{name:.*}/get | Export an image
*ImagesApi* | [**LibpodExportImages**](docs/ImagesApi.md#libpodexportimages) | **Get** /libpod/images/export | Export multiple images
*ImagesApi* | [**LibpodImageExists**](docs/ImagesApi.md#libpodimageexists) | **Get** /libpod/images/{name:.*}/exists | Image exists
*ImagesApi* | [**LibpodImageHistory**](docs/ImagesApi.md#libpodimagehistory) | **Get** /libpod/images/{name:.*}/history | History of an image
*ImagesApi* | [**LibpodImageTree**](docs/ImagesApi.md#libpodimagetree) | **Get** /libpod/images/{name:.*}/tree | Image tree
*ImagesApi* | [**LibpodImagesImport**](docs/ImagesApi.md#libpodimagesimport) | **Post** /libpod/images/import | Import image
*ImagesApi* | [**LibpodImagesLoad**](docs/ImagesApi.md#libpodimagesload) | **Post** /libpod/images/load | Load image
*ImagesApi* | [**LibpodImagesPull**](docs/ImagesApi.md#libpodimagespull) | **Post** /libpod/images/pull | Pull images
*ImagesApi* | [**LibpodImagesRemove**](docs/ImagesApi.md#libpodimagesremove) | **Delete** /libpod/images/remove | Remove one or more images from the storage.
*ImagesApi* | [**LibpodInspectImage**](docs/ImagesApi.md#libpodinspectimage) | **Get** /libpod/images/{name:.*}/json | Inspect an image
*ImagesApi* | [**LibpodListImages**](docs/ImagesApi.md#libpodlistimages) | **Get** /libpod/images/json | List Images
*ImagesApi* | [**LibpodPruneImages**](docs/ImagesApi.md#libpodpruneimages) | **Post** /libpod/images/prune | Prune unused images
*ImagesApi* | [**LibpodPushImage**](docs/ImagesApi.md#libpodpushimage) | **Post** /libpod/images/{name:.*}/push | Push Image
*ImagesApi* | [**LibpodRemoveImage**](docs/ImagesApi.md#libpodremoveimage) | **Delete** /libpod/images/{name:.*} | Remove an image from the local storage.
*ImagesApi* | [**LibpodSearchImages**](docs/ImagesApi.md#libpodsearchimages) | **Get** /libpod/images/search | Search images
*ImagesApi* | [**LibpodTagImage**](docs/ImagesApi.md#libpodtagimage) | **Post** /libpod/images/{name:.*}/tag | Tag an image
*ImagesApi* | [**LibpodUntagImage**](docs/ImagesApi.md#libpoduntagimage) | **Post** /libpod/images/{name:.*}/untag | Untag an image
*ImagesCompatApi* | [**BuildImage**](docs/ImagesCompatApi.md#buildimage) | **Post** /build | Create image
*ImagesCompatApi* | [**CreateImage**](docs/ImagesCompatApi.md#createimage) | **Post** /images/create | Create an image
*ImagesCompatApi* | [**ExportImage**](docs/ImagesCompatApi.md#exportimage) | **Get** /images/{name:.*}/get | Export an image
*ImagesCompatApi* | [**ImageHistory**](docs/ImagesCompatApi.md#imagehistory) | **Get** /images/{name:.*}/history | History of an image
*ImagesCompatApi* | [**ImportImage**](docs/ImagesCompatApi.md#importimage) | **Post** /images/load | Import image
*ImagesCompatApi* | [**InspectImage**](docs/ImagesCompatApi.md#inspectimage) | **Get** /images/{name:.*}/json | Inspect an image
*ImagesCompatApi* | [**ListImages**](docs/ImagesCompatApi.md#listimages) | **Get** /images/json | List Images
*ImagesCompatApi* | [**PruneImages**](docs/ImagesCompatApi.md#pruneimages) | **Post** /images/prune | Prune unused images
*ImagesCompatApi* | [**PushImage**](docs/ImagesCompatApi.md#pushimage) | **Post** /images/{name:.*}/push | Push Image
*ImagesCompatApi* | [**RemoveImage**](docs/ImagesCompatApi.md#removeimage) | **Delete** /images/{name:.*} | Remove Image
*ImagesCompatApi* | [**SearchImages**](docs/ImagesCompatApi.md#searchimages) | **Get** /images/search | Search images
*ImagesCompatApi* | [**TagImage**](docs/ImagesCompatApi.md#tagimage) | **Post** /images/{name:.*}/tag | Tag an image
*ManifestsApi* | [**AddManifest**](docs/ManifestsApi.md#addmanifest) | **Post** /libpod/manifests/{name:.*}/add | 
*ManifestsApi* | [**Create**](docs/ManifestsApi.md#create) | **Post** /libpod/manifests/create | Create
*ManifestsApi* | [**Inspect**](docs/ManifestsApi.md#inspect) | **Get** /libpod/manifests/{name:.*}/json | Inspect
*ManifestsApi* | [**PushManifest**](docs/ManifestsApi.md#pushmanifest) | **Post** /libpod/manifests/{name}/push | Push
*ManifestsApi* | [**RemoveManifest**](docs/ManifestsApi.md#removemanifest) | **Delete** /libpod/manifests/{name:.*} | Remove
*NetworksApi* | [**LibpodCreateNetwork**](docs/NetworksApi.md#libpodcreatenetwork) | **Post** /libpod/networks/create | Create network
*NetworksApi* | [**LibpodInspectNetwork**](docs/NetworksApi.md#libpodinspectnetwork) | **Get** /libpod/networks/{name}/json | Inspect a network
*NetworksApi* | [**LibpodListNetwork**](docs/NetworksApi.md#libpodlistnetwork) | **Get** /libpod/networks/json | List networks
*NetworksApi* | [**LibpodRemoveNetwork**](docs/NetworksApi.md#libpodremovenetwork) | **Delete** /libpod/networks/{name} | Remove a network
*NetworksCompatApi* | [**CompatCreateNetwork**](docs/NetworksCompatApi.md#compatcreatenetwork) | **Post** /networks/create | Create network
*NetworksCompatApi* | [**CompatInspectNetwork**](docs/NetworksCompatApi.md#compatinspectnetwork) | **Get** /networks/{name} | Inspect a network
*NetworksCompatApi* | [**CompatListNetwork**](docs/NetworksCompatApi.md#compatlistnetwork) | **Get** /networks | List networks
*NetworksCompatApi* | [**CompatRemoveNetwork**](docs/NetworksCompatApi.md#compatremovenetwork) | **Delete** /networks/{name} | Remove a network
*PodsApi* | [**CreatePod**](docs/PodsApi.md#createpod) | **Post** /libpod/pods/create | Create a pod
*PodsApi* | [**InspectPod**](docs/PodsApi.md#inspectpod) | **Get** /libpod/pods/{name}/json | Inspect pod
*PodsApi* | [**KillPod**](docs/PodsApi.md#killpod) | **Post** /libpod/pods/{name}/kill | Kill a pod
*PodsApi* | [**LibpodGenerateKube**](docs/PodsApi.md#libpodgeneratekube) | **Get** /libpod/generate/{name:.*}/kube | Generate a Kubernetes YAML file.
*PodsApi* | [**LibpodGenerateSystemd**](docs/PodsApi.md#libpodgeneratesystemd) | **Get** /libpod/generate/{name:.*}/systemd | Generate Systemd Units
*PodsApi* | [**LibpodPlayKube**](docs/PodsApi.md#libpodplaykube) | **Post** /libpod/play/kube | Play a Kubernetes YAML file.
*PodsApi* | [**ListPods**](docs/PodsApi.md#listpods) | **Get** /libpod/pods/json | List pods
*PodsApi* | [**PausePod**](docs/PodsApi.md#pausepod) | **Post** /libpod/pods/{name}/pause | Pause a pod
*PodsApi* | [**PodExists**](docs/PodsApi.md#podexists) | **Get** /libpod/pods/{name}/exists | Pod exists
*PodsApi* | [**PrunePods**](docs/PodsApi.md#prunepods) | **Post** /libpod/pods/prune | Prune unused pods
*PodsApi* | [**RemovePod**](docs/PodsApi.md#removepod) | **Delete** /libpod/pods/{name} | Remove pod
*PodsApi* | [**RestartPod**](docs/PodsApi.md#restartpod) | **Post** /libpod/pods/{name}/restart | Restart a pod
*PodsApi* | [**StartPod**](docs/PodsApi.md#startpod) | **Post** /libpod/pods/{name}/start | Start a pod
*PodsApi* | [**StatsPod**](docs/PodsApi.md#statspod) | **Get** /libpod/pods/stats | Get stats for one or more pods
*PodsApi* | [**StopPod**](docs/PodsApi.md#stoppod) | **Post** /libpod/pods/{name}/stop | Stop a pod
*PodsApi* | [**TopPod**](docs/PodsApi.md#toppod) | **Get** /libpod/pods/{name}/top | List processes
*PodsApi* | [**UnpausePod**](docs/PodsApi.md#unpausepod) | **Post** /libpod/pods/{name}/unpause | Unpause a pod
*SystemApi* | [**Df**](docs/SystemApi.md#df) | **Get** /libpod/system/df | Show disk usage
*SystemApi* | [**LibpodGetEvents**](docs/SystemApi.md#libpodgetevents) | **Get** /libpod/events | Get events
*SystemApi* | [**LibpodGetInfo**](docs/SystemApi.md#libpodgetinfo) | **Get** /libpod/info | Get info
*SystemApi* | [**LibpodPingGet**](docs/SystemApi.md#libpodpingget) | **Get** /libpod/_ping | Ping service
*SystemApi* | [**PruneSystem**](docs/SystemApi.md#prunesystem) | **Post** /libpod/system/prune | Prune unused data
*SystemApi* | [**SystemVersion**](docs/SystemApi.md#systemversion) | **Get** /libpod/version | Component Version information
*SystemCompatApi* | [**CompatSystemVersion**](docs/SystemCompatApi.md#compatsystemversion) | **Get** /version | Component Version information
*SystemCompatApi* | [**GetEvents**](docs/SystemCompatApi.md#getevents) | **Get** /events | Get events
*SystemCompatApi* | [**GetInfo**](docs/SystemCompatApi.md#getinfo) | **Get** /info | Get info
*SystemCompatApi* | [**LibpodPingGet**](docs/SystemCompatApi.md#libpodpingget) | **Get** /libpod/_ping | Ping service
*VolumesApi* | [**LibpodCreateVolume**](docs/VolumesApi.md#libpodcreatevolume) | **Post** /libpod/volumes/create | Create a volume
*VolumesApi* | [**LibpodInspectVolume**](docs/VolumesApi.md#libpodinspectvolume) | **Get** /libpod/volumes/{name}/json | Inspect volume
*VolumesApi* | [**LibpodListVolumes**](docs/VolumesApi.md#libpodlistvolumes) | **Get** /libpod/volumes/json | List volumes
*VolumesApi* | [**LibpodPruneVolumes**](docs/VolumesApi.md#libpodprunevolumes) | **Post** /libpod/volumes/prune | Prune volumes
*VolumesApi* | [**LibpodRemoveVolume**](docs/VolumesApi.md#libpodremovevolume) | **Delete** /libpod/volumes/{name} | Remove volume
*VolumesCompatApi* | [**CreateVolume**](docs/VolumesCompatApi.md#createvolume) | **Post** /volumes/create | Create a volume
*VolumesCompatApi* | [**InspectVolume**](docs/VolumesCompatApi.md#inspectvolume) | **Get** /volumes/{name} | Inspect volume
*VolumesCompatApi* | [**ListVolumes**](docs/VolumesCompatApi.md#listvolumes) | **Get** /volumes | List volumes
*VolumesCompatApi* | [**PruneVolumes**](docs/VolumesCompatApi.md#prunevolumes) | **Post** /volumes/prune | Prune volumes
*VolumesCompatApi* | [**RemoveVolume**](docs/VolumesCompatApi.md#removevolume) | **Delete** /volumes/{name} | Remove volume

## Documentation For Models

 - [Address](docs/Address.md)
 - [AuthenticateOkBody](docs/AuthenticateOkBody.md)
 - [AutoUserNsOptions](docs/AutoUserNsOptions.md)
 - [Body](docs/Body.md)
 - [Body1](docs/Body1.md)
 - [Body10](docs/Body10.md)
 - [Body11](docs/Body11.md)
 - [Body2](docs/Body2.md)
 - [Body3](docs/Body3.md)
 - [Body4](docs/Body4.md)
 - [Body5](docs/Body5.md)
 - [Body6](docs/Body6.md)
 - [Body7](docs/Body7.md)
 - [Body8](docs/Body8.md)
 - [Body9](docs/Body9.md)
 - [CgroupConfig](docs/CgroupConfig.md)
 - [Change](docs/Change.md)
 - [Config](docs/Config.md)
 - [ConfigReference](docs/ConfigReference.md)
 - [ConmonInfo](docs/ConmonInfo.md)
 - [ContainerBasicConfig](docs/ContainerBasicConfig.md)
 - [ContainerCgroupConfig](docs/ContainerCgroupConfig.md)
 - [ContainerChangeResponseItem](docs/ContainerChangeResponseItem.md)
 - [ContainerCreateCreatedBody](docs/ContainerCreateCreatedBody.md)
 - [ContainerHealthCheckConfig](docs/ContainerHealthCheckConfig.md)
 - [ContainerNamedVolume](docs/ContainerNamedVolume.md)
 - [ContainerNetworkConfig](docs/ContainerNetworkConfig.md)
 - [ContainerNode](docs/ContainerNode.md)
 - [ContainerResourceConfig](docs/ContainerResourceConfig.md)
 - [ContainerSecurityConfig](docs/ContainerSecurityConfig.md)
 - [ContainerSize](docs/ContainerSize.md)
 - [ContainerState](docs/ContainerState.md)
 - [ContainerStorageConfig](docs/ContainerStorageConfig.md)
 - [ContainerStore](docs/ContainerStore.md)
 - [ContainerTopOkBody](docs/ContainerTopOkBody.md)
 - [ContainerUpdateOkBody](docs/ContainerUpdateOkBody.md)
 - [ContainerWaitOkBody](docs/ContainerWaitOkBody.md)
 - [ContainerWaitOkBodyError](docs/ContainerWaitOkBodyError.md)
 - [ContainersPruneReport](docs/ContainersPruneReport.md)
 - [CreateConfig](docs/CreateConfig.md)
 - [CreateResourceConfig](docs/CreateResourceConfig.md)
 - [Data](docs/Data.md)
 - [DeviceMapping](docs/DeviceMapping.md)
 - [DeviceRequest](docs/DeviceRequest.md)
 - [DistributionInfo](docs/DistributionInfo.md)
 - [DockerVolumeCreate](docs/DockerVolumeCreate.md)
 - [EndpointIpamConfig](docs/EndpointIpamConfig.md)
 - [EndpointResource](docs/EndpointResource.md)
 - [EndpointSettings](docs/EndpointSettings.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [GraphDriverData](docs/GraphDriverData.md)
 - [Health](docs/Health.md)
 - [HealthCheckLog](docs/HealthCheckLog.md)
 - [HealthCheckResults](docs/HealthCheckResults.md)
 - [HealthConfig](docs/HealthConfig.md)
 - [HealthcheckResult](docs/HealthcheckResult.md)
 - [History](docs/History.md)
 - [HistoryResponseItem](docs/HistoryResponseItem.md)
 - [HostConfig](docs/HostConfig.md)
 - [HostInfo](docs/HostInfo.md)
 - [IdMap](docs/IdMap.md)
 - [IdMappingOptions](docs/IdMappingOptions.md)
 - [IdMappings](docs/IdMappings.md)
 - [IdResponse](docs/IdResponse.md)
 - [ImageConfig](docs/ImageConfig.md)
 - [ImageDeleteResponse](docs/ImageDeleteResponse.md)
 - [ImageDeleteResponseItem](docs/ImageDeleteResponseItem.md)
 - [ImageImportReport](docs/ImageImportReport.md)
 - [ImageLayer](docs/ImageLayer.md)
 - [ImageLoadReport](docs/ImageLoadReport.md)
 - [ImageMetadata](docs/ImageMetadata.md)
 - [ImageStore](docs/ImageStore.md)
 - [ImageSummary](docs/ImageSummary.md)
 - [Info](docs/Info.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse2001](docs/InlineResponse2001.md)
 - [InlineResponse20010](docs/InlineResponse20010.md)
 - [InlineResponse20011](docs/InlineResponse20011.md)
 - [InlineResponse20012](docs/InlineResponse20012.md)
 - [InlineResponse20013](docs/InlineResponse20013.md)
 - [InlineResponse20014](docs/InlineResponse20014.md)
 - [InlineResponse20015](docs/InlineResponse20015.md)
 - [InlineResponse20016](docs/InlineResponse20016.md)
 - [InlineResponse20017](docs/InlineResponse20017.md)
 - [InlineResponse20018](docs/InlineResponse20018.md)
 - [InlineResponse20019](docs/InlineResponse20019.md)
 - [InlineResponse2002](docs/InlineResponse2002.md)
 - [InlineResponse2003](docs/InlineResponse2003.md)
 - [InlineResponse2003Error](docs/InlineResponse2003Error.md)
 - [InlineResponse2004](docs/InlineResponse2004.md)
 - [InlineResponse2005](docs/InlineResponse2005.md)
 - [InlineResponse2006](docs/InlineResponse2006.md)
 - [InlineResponse2007](docs/InlineResponse2007.md)
 - [InlineResponse2008](docs/InlineResponse2008.md)
 - [InlineResponse2009](docs/InlineResponse2009.md)
 - [InlineResponse201](docs/InlineResponse201.md)
 - [InlineResponse400](docs/InlineResponse400.md)
 - [InspectAdditionalNetwork](docs/InspectAdditionalNetwork.md)
 - [InspectBlkioThrottleDevice](docs/InspectBlkioThrottleDevice.md)
 - [InspectBlkioWeightDevice](docs/InspectBlkioWeightDevice.md)
 - [InspectContainerConfig](docs/InspectContainerConfig.md)
 - [InspectContainerHostConfig](docs/InspectContainerHostConfig.md)
 - [InspectContainerState](docs/InspectContainerState.md)
 - [InspectDevice](docs/InspectDevice.md)
 - [InspectHostPort](docs/InspectHostPort.md)
 - [InspectLogConfig](docs/InspectLogConfig.md)
 - [InspectMount](docs/InspectMount.md)
 - [InspectNetworkSettings](docs/InspectNetworkSettings.md)
 - [InspectPodContainerInfo](docs/InspectPodContainerInfo.md)
 - [InspectPodInfraConfig](docs/InspectPodInfraConfig.md)
 - [InspectRestartPolicy](docs/InspectRestartPolicy.md)
 - [InspectUlimit](docs/InspectUlimit.md)
 - [IpNet](docs/IpNet.md)
 - [Ipam](docs/Ipam.md)
 - [IpamConfig](docs/IpamConfig.md)
 - [IpcConfig](docs/IpcConfig.md)
 - [LibpodContainersPruneReport](docs/LibpodContainersPruneReport.md)
 - [LibpodImagesPullReport](docs/LibpodImagesPullReport.md)
 - [LibpodImagesRemoveReport](docs/LibpodImagesRemoveReport.md)
 - [LinuxBlockIo](docs/LinuxBlockIo.md)
 - [LinuxBlockIoDevice](docs/LinuxBlockIoDevice.md)
 - [LinuxCpu](docs/LinuxCpu.md)
 - [LinuxDevice](docs/LinuxDevice.md)
 - [LinuxDeviceCgroup](docs/LinuxDeviceCgroup.md)
 - [LinuxHugepageLimit](docs/LinuxHugepageLimit.md)
 - [LinuxInterfacePriority](docs/LinuxInterfacePriority.md)
 - [LinuxMemory](docs/LinuxMemory.md)
 - [LinuxNetwork](docs/LinuxNetwork.md)
 - [LinuxPids](docs/LinuxPids.md)
 - [LinuxRdma](docs/LinuxRdma.md)
 - [LinuxResources](docs/LinuxResources.md)
 - [LinuxThrottleDevice](docs/LinuxThrottleDevice.md)
 - [LinuxWeightDevice](docs/LinuxWeightDevice.md)
 - [List](docs/List.md)
 - [ListContainer](docs/ListContainer.md)
 - [ListContainerNamespaces](docs/ListContainerNamespaces.md)
 - [ListPodContainer](docs/ListPodContainer.md)
 - [ListPodsReport](docs/ListPodsReport.md)
 - [ListRegistriesReport](docs/ListRegistriesReport.md)
 - [LogConfig](docs/LogConfig.md)
 - [ManifestAddOpts](docs/ManifestAddOpts.md)
 - [Mount](docs/Mount.md)
 - [MountPoint](docs/MountPoint.md)
 - [NamedVolume](docs/NamedVolume.md)
 - [Namespace](docs/Namespace.md)
 - [NetworkConfig](docs/NetworkConfig.md)
 - [NetworkCreate](docs/NetworkCreate.md)
 - [NetworkCreateOptions](docs/NetworkCreateOptions.md)
 - [NetworkCreateReport](docs/NetworkCreateReport.md)
 - [NetworkCreateRequest](docs/NetworkCreateRequest.md)
 - [NetworkListReport](docs/NetworkListReport.md)
 - [NetworkResource](docs/NetworkResource.md)
 - [NetworkRmReport](docs/NetworkRmReport.md)
 - [NetworkSettings](docs/NetworkSettings.md)
 - [OciRuntimeInfo](docs/OciRuntimeInfo.md)
 - [OverlayVolume](docs/OverlayVolume.md)
 - [PeerInfo](docs/PeerInfo.md)
 - [PidConfig](docs/PidConfig.md)
 - [PlayKubePod](docs/PlayKubePod.md)
 - [PlayKubeReport](docs/PlayKubeReport.md)
 - [Plugin](docs/Plugin.md)
 - [PluginConfig](docs/PluginConfig.md)
 - [PluginConfigArgs](docs/PluginConfigArgs.md)
 - [PluginConfigInterface](docs/PluginConfigInterface.md)
 - [PluginConfigLinux](docs/PluginConfigLinux.md)
 - [PluginConfigNetwork](docs/PluginConfigNetwork.md)
 - [PluginConfigRootfs](docs/PluginConfigRootfs.md)
 - [PluginConfigUser](docs/PluginConfigUser.md)
 - [PluginDevice](docs/PluginDevice.md)
 - [PluginEnv](docs/PluginEnv.md)
 - [PluginInterfaceType](docs/PluginInterfaceType.md)
 - [PluginMount](docs/PluginMount.md)
 - [PluginSettings](docs/PluginSettings.md)
 - [PodBasicConfig](docs/PodBasicConfig.md)
 - [PodCgroupConfig](docs/PodCgroupConfig.md)
 - [PodCreateConfig](docs/PodCreateConfig.md)
 - [PodKillReport](docs/PodKillReport.md)
 - [PodNetworkConfig](docs/PodNetworkConfig.md)
 - [PodPauseReport](docs/PodPauseReport.md)
 - [PodPruneReport](docs/PodPruneReport.md)
 - [PodRestartReport](docs/PodRestartReport.md)
 - [PodRmReport](docs/PodRmReport.md)
 - [PodSpecGenerator](docs/PodSpecGenerator.md)
 - [PodStartReport](docs/PodStartReport.md)
 - [PodStatsReport](docs/PodStatsReport.md)
 - [PodStopReport](docs/PodStopReport.md)
 - [PodUnpauseReport](docs/PodUnpauseReport.md)
 - [Port](docs/Port.md)
 - [PortBinding](docs/PortBinding.md)
 - [PortMapping](docs/PortMapping.md)
 - [PosixRlimit](docs/PosixRlimit.md)
 - [RemoteSocket](docs/RemoteSocket.md)
 - [Report](docs/Report.md)
 - [RestartPolicy](docs/RestartPolicy.md)
 - [RootFs](docs/RootFs.md)
 - [Schema2HealthConfig](docs/Schema2HealthConfig.md)
 - [SecurityConfig](docs/SecurityConfig.md)
 - [ServiceInfo](docs/ServiceInfo.md)
 - [ServiceUpdateResponse](docs/ServiceUpdateResponse.md)
 - [SlirpInfo](docs/SlirpInfo.md)
 - [SpecGenerator](docs/SpecGenerator.md)
 - [StoreInfo](docs/StoreInfo.md)
 - [SystemDfContainerReport](docs/SystemDfContainerReport.md)
 - [SystemDfImageReport](docs/SystemDfImageReport.md)
 - [SystemDfVolumeReport](docs/SystemDfVolumeReport.md)
 - [Task](docs/Task.md)
 - [ThrottleDevice](docs/ThrottleDevice.md)
 - [Ulimit](docs/Ulimit.md)
 - [UserConfig](docs/UserConfig.md)
 - [UtsConfig](docs/UtsConfig.md)
 - [Version](docs/Version.md)
 - [Volume](docs/Volume.md)
 - [VolumeCreate](docs/VolumeCreate.md)
 - [VolumeCreateBody](docs/VolumeCreateBody.md)
 - [VolumeInfo](docs/VolumeInfo.md)
 - [VolumeListBody](docs/VolumeListBody.md)
 - [VolumeListOkBody](docs/VolumeListOkBody.md)
 - [VolumePruneReport](docs/VolumePruneReport.md)
 - [VolumeUsageData](docs/VolumeUsageData.md)
 - [WeightDevice](docs/WeightDevice.md)

## Documentation For Authorization
 Endpoints do not require authorization.


## Author

podman@lists.podman.io