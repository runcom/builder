package libpod

import (
	"fmt"
	"strings"

	"github.com/containers/libpod/pkg/lookup"
	"github.com/containers/libpod/pkg/util"
	"github.com/cri-o/ocicni/pkg/ocicni"
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// InspectForKube takes a slice of libpod containers and generates
// one v1.Pod description that includes just a single container.
func (c *Container) InspectForKube() (*v1.Pod, error) {
	// Generate the v1.Pod yaml description
	return simplePodWithV1Container(c)
}

// simplePodWithV1Container is a function used by inspect when kube yaml needs to be generated
// for a single container.  we "insert" that container description in a pod.
func simplePodWithV1Container(ctr *Container) (*v1.Pod, error) {
	var containers []v1.Container
	result, err := containerToV1Container(ctr)
	if err != nil {
		return nil, err
	}
	containers = append(containers, result)

	tm := v12.TypeMeta{
		Kind:       "Pod",
		APIVersion: "v1",
	}

	// Add a label called "app" with the containers name as a value
	labels := make(map[string]string)
	labels["app"] = removeUnderscores(ctr.Name())
	om := v12.ObjectMeta{
		// The name of the pod is container_name-libpod
		Name:   fmt.Sprintf("%s-libpod", removeUnderscores(ctr.Name())),
		Labels: labels,
		// CreationTimestamp seems to be required, so adding it; in doing so, the timestamp
		// will reflect time this is run (not container create time) because the conversion
		// of the container create time to v1 Time is probably not warranted nor worthwhile.
		CreationTimestamp: v12.Now(),
	}
	ps := v1.PodSpec{
		Containers: containers,
	}
	p := v1.Pod{
		TypeMeta:   tm,
		ObjectMeta: om,
		Spec:       ps,
	}
	return &p, nil
}

// containerToV1Container converts information we know about a libpod container
// to a V1.Container specification.
func containerToV1Container(c *Container) (v1.Container, error) {
	kubeContainer := v1.Container{}
	kubeSec, err := generateKubeSecurityContext(c)
	if err != nil {
		return kubeContainer, err
	}

	if len(c.config.Spec.Linux.Devices) > 0 {
		// TODO Enable when we can support devices and their names
		devices, err := generateKubeVolumeDeviceFromLinuxDevice(c.Spec().Linux.Devices)
		if err != nil {
			return kubeContainer, err
		}
		kubeContainer.VolumeDevices = devices
		return kubeContainer, errors.Wrapf(ErrNotImplemented, "linux devices")
	}

	if len(c.config.UserVolumes) > 0 {
		// TODO When we until we can resolve what the volume name should be, this is disabled
		// Volume names need to be coordinated "globally" in the kube files.
		volumes, err := libpodMountsToKubeVolumeMounts(c)
		if err != nil {
			return kubeContainer, err
		}
		kubeContainer.VolumeMounts = volumes
		return kubeContainer, errors.Wrapf(ErrNotImplemented, "volume names")
	}

	envVariables, err := libpodEnvVarsToKubeEnvVars(c.config.Spec.Process.Env)
	if err != nil {
		return kubeContainer, nil
	}

	ports, err := ocicniPortMappingToContainerPort(c.PortMappings())
	if err != nil {
		return kubeContainer, nil
	}

	containerCommands := c.Command()
	kubeContainer.Name = removeUnderscores(c.Name())

	_, image := c.Image()
	kubeContainer.Image = image
	kubeContainer.Stdin = c.Stdin()
	kubeContainer.Command = containerCommands
	// TODO need to figure out how we handle command vs entry point.  Kube appears to prefer entrypoint.
	// right now we just take the container's command
	//container.Args = args
	kubeContainer.WorkingDir = c.WorkingDir()
	kubeContainer.Ports = ports
	// This should not be applicable
	//container.EnvFromSource =
	kubeContainer.Env = envVariables
	// TODO enable resources when we can support naming conventions
	//container.Resources
	kubeContainer.SecurityContext = kubeSec
	kubeContainer.StdinOnce = false
	kubeContainer.TTY = c.config.Spec.Process.Terminal

	return kubeContainer, nil
}

// ocicniPortMappingToContainerPort takes an ocicni portmapping and converts
// it to a v1.ContainerPort format for kube output
func ocicniPortMappingToContainerPort(portMappings []ocicni.PortMapping) ([]v1.ContainerPort, error) {
	var containerPorts []v1.ContainerPort
	for _, p := range portMappings {
		var protocol v1.Protocol
		switch strings.ToUpper(p.Protocol) {
		case "TCP":
			protocol = v1.ProtocolTCP
		case "UDP":
			protocol = v1.ProtocolUDP
		default:
			return containerPorts, errors.Errorf("unknown network protocol %s", p.Protocol)
		}
		cp := v1.ContainerPort{
			// Name will not be supported
			HostPort:      p.HostPort,
			HostIP:        p.HostIP,
			ContainerPort: p.ContainerPort,
			Protocol:      protocol,
		}
		containerPorts = append(containerPorts, cp)
	}
	return containerPorts, nil
}

// libpodEnvVarsToKubeEnvVars converts a key=value string slice to []v1.EnvVar
func libpodEnvVarsToKubeEnvVars(envs []string) ([]v1.EnvVar, error) {
	var envVars []v1.EnvVar
	for _, e := range envs {
		splitE := strings.SplitN(e, "=", 2)
		if len(splitE) != 2 {
			return envVars, errors.Errorf("environment variable %s is malformed; should be key=value", e)
		}
		ev := v1.EnvVar{
			Name:  splitE[0],
			Value: splitE[1],
		}
		envVars = append(envVars, ev)
	}
	return envVars, nil
}

// Is this worth it?
func libpodMaxAndMinToResourceList(c *Container) (v1.ResourceList, v1.ResourceList) { //nolint
	// It does not appear we can properly calculate CPU resources from the information
	// we know in libpod.  Libpod knows CPUs by time, shares, etc.

	// We also only know about a memory limit; no memory minimum
	maxResources := make(map[v1.ResourceName]resource.Quantity)
	minResources := make(map[v1.ResourceName]resource.Quantity)
	config := c.Config()
	maxMem := config.Spec.Linux.Resources.Memory.Limit

	_ = maxMem

	return maxResources, minResources
}

func generateKubeVolumeMount(hostSourcePath string, mounts []specs.Mount) (v1.VolumeMount, error) {
	vm := v1.VolumeMount{}
	for _, m := range mounts {
		if m.Source == hostSourcePath {
			// TODO Name is not provided and is required by Kube; therefore, this is disabled earlier
			//vm.Name =
			vm.MountPath = m.Source
			vm.SubPath = m.Destination
			if util.StringInSlice("ro", m.Options) {
				vm.ReadOnly = true
			}
			return vm, nil
		}
	}
	return vm, errors.New("unable to find mount source")
}

// libpodMountsToKubeVolumeMounts converts the containers mounts to a struct kube understands
func libpodMountsToKubeVolumeMounts(c *Container) ([]v1.VolumeMount, error) {
	// At this point, I dont think we can distinguish between the default
	// volume mounts and user added ones.  For now, we pass them all.
	var vms []v1.VolumeMount
	for _, hostSourcePath := range c.config.UserVolumes {
		vm, err := generateKubeVolumeMount(hostSourcePath, c.config.Spec.Mounts)
		if err != nil {
			return vms, err
		}
		vms = append(vms, vm)
	}
	return vms, nil
}

// generateKubeSecurityContext generates a securityContext based on the existing container
func generateKubeSecurityContext(c *Container) (*v1.SecurityContext, error) {
	priv := c.Privileged()
	ro := c.IsReadOnly()
	allowPrivEscalation := !c.Spec().Process.NoNewPrivileges

	// TODO enable use of capabilities when we can figure out how to extract cap-add|remove
	//caps := v1.Capabilities{
	//	//Add: c.config.Spec.Process.Capabilities
	//}
	sc := v1.SecurityContext{
		// TODO enable use of capabilities when we can figure out how to extract cap-add|remove
		//Capabilities: &caps,
		Privileged: &priv,
		// TODO How do we know if selinux were passed into podman
		//SELinuxOptions:
		// RunAsNonRoot is an optional parameter; our first implementations should be root only; however
		// I'm leaving this as a bread-crumb for later
		//RunAsNonRoot:             &nonRoot,
		ReadOnlyRootFilesystem:   &ro,
		AllowPrivilegeEscalation: &allowPrivEscalation,
	}

	if c.User() != "" {
		// It is *possible* that
		logrus.Debugf("Looking in container for user: %s", c.User())
		u, err := lookup.GetUser(c.state.Mountpoint, c.User())
		if err != nil {
			return nil, err
		}
		user := int64(u.Uid)
		sc.RunAsUser = &user
	}
	return &sc, nil
}

// generateKubeVolumeDeviceFromLinuxDevice takes a list of devices and makes a VolumeDevice struct for kube
func generateKubeVolumeDeviceFromLinuxDevice(devices []specs.LinuxDevice) ([]v1.VolumeDevice, error) {
	var volumeDevices []v1.VolumeDevice
	for _, d := range devices {
		vd := v1.VolumeDevice{
			// TBD How are we going to sync up these names
			//Name:
			DevicePath: d.Path,
		}
		volumeDevices = append(volumeDevices, vd)
	}
	return volumeDevices, nil
}

func removeUnderscores(s string) string {
	return strings.Replace(s, "_", "", -1)
}
