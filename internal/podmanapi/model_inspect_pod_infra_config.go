/*
 * Provides a container compatible interface.
 *
 * This documentation describes the Podman v2.0 RESTful API. It replaces the Podman v1.0 API and was initially delivered along with Podman v2.0.  It consists of a Docker-compatible API and a Libpod API providing support for Podman’s unique features such as pods.  To start the service and keep it running for 5,000 seconds (-t 0 runs forever):  podman system service -t 5000 &  You can then use cURL on the socket using requests documented below.  NOTE: if you install the package podman-docker, it will create a symbolic link for /var/run/docker.sock to /run/podman/podman.sock  See podman-service(1) for more information.  Quick Examples:  'podman info'  curl --unix-socket /run/podman/podman.sock http://d/v1.0.0/libpod/info  'podman pull quay.io/containers/podman'  curl -XPOST --unix-socket /run/podman/podman.sock -v 'http://d/v1.0.0/images/create?fromImage=quay.io%2Fcontainers%2Fpodman'  'podman list images'  curl --unix-socket /run/podman/podman.sock -v 'http://d/v1.0.0/libpod/images/json' | jq
 *
 * API version: 0.0.1
 * Contact: podman@lists.podman.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// InspectPodInfraConfig contains the configuration of the pod's infra container.
type InspectPodInfraConfig struct {
	// DNSOption is a set of DNS options that will be used by the infra container's resolv.conf and shared with the remainder of the pod.
	DNSOption []string `json:"DNSOption,omitempty"`
	// DNSSearch is a set of DNS search domains that will be used by the infra container's resolv.conf and shared with the remainder of the pod.
	DNSSearch []string `json:"DNSSearch,omitempty"`
	// DNSServer is a set of DNS Servers that will be used by the infra container's resolv.conf and shared with the remainder of the pod.
	DNSServer []string `json:"DNSServer,omitempty"`
	// HostAdd adds a number of hosts to the infra container's resolv.conf which will be shared with the rest of the pod.
	HostAdd []string `json:"HostAdd,omitempty"`
	// HostNetwork is whether the infra container (and thus the whole pod) will use the host's network and not create a network namespace.
	HostNetwork bool `json:"HostNetwork,omitempty"`
	// NetworkOptions are additional options for each network
	NetworkOptions map[string][]string `json:"NetworkOptions,omitempty"`
	// Networks is a list of CNI networks the pod will join.
	Networks []string `json:"Networks,omitempty"`
	// NoManageHosts indicates that the pod will not manage /etc/hosts and instead each container will handle their own.
	NoManageHosts bool `json:"NoManageHosts,omitempty"`
	// NoManageResolvConf indicates that the pod will not manage resolv.conf and instead each container will handle their own.
	NoManageResolvConf bool `json:"NoManageResolvConf,omitempty"`
	// PortBindings are ports that will be forwarded to the infra container and then shared with the pod.
	PortBindings map[string][]InspectHostPort `json:"PortBindings,omitempty"`
	StaticIP *[]int32 `json:"StaticIP,omitempty"`
	StaticMAC *[]int32 `json:"StaticMAC,omitempty"`
}