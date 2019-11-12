package k8sio

import (
	"io"
	"log"
	"os"
	"strings"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/remotecommand"
)

// RemoteExec runs remote commands on given pod in namespace.
type RemoteExec struct {
	k8sClient *K8SClient
}

// NewRemoteExec returns an instance of RemoteExec.
func NewRemoteExec(kubeConfig string) (rc *RemoteExec, err error) {
	client, err := NewK8SClient(kubeConfig)
	if err != nil {
		return nil, err
	}

	return &RemoteExec{
		k8sClient: client,
	}, nil
}

// RemoteExecOptions contains exec options for running remote command.
type RemoteExecOptions struct {
	Namespace string
	PodName   string
	Command   []string

	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

// RunCommandWithStdout runs a command on a given pod in a namespace with stdout and stderr.
func (re *RemoteExec) RunCommandWithStdout(namespace, podName, command string) error {
	return re.RemoteExec(RemoteExecOptions{
		Namespace: namespace,
		PodName:   podName,
		Command:   strings.Split(command, " "),
		Stdin:     strings.NewReader(""),
		Stdout:    os.Stdout,
		Stderr:    os.Stderr,
	})
}

// RemoteExec runs a command on a given pod in a namespace.
func (re *RemoteExec) RemoteExec(options RemoteExecOptions) error {
	pod, err := re.k8sClient.GetPod(options.Namespace, options.PodName)
	if err != nil {
		return err
	}

	containerName := pod.Spec.Containers[0].Name
	podExecOptions := &v1.PodExecOptions{
		Container: containerName,
		Command:   options.Command,
		Stdin:     true,
		Stdout:    true,
		Stderr:    true,
	}
	execRequest := re.k8sClient.KubeClient.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(options.PodName).
		Namespace(options.Namespace).
		SubResource("exec")
	execRequest.VersionedParams(podExecOptions, scheme.ParameterCodec)

	if IsDebug {
		log.Printf("URL: %s\n", execRequest.URL())
		log.Printf("Executing command \"%s\" on container [%s] in pod [%s].\n",
			options.Command, containerName, options.PodName)
	}

	exec, err := remotecommand.NewSPDYExecutor(re.k8sClient.KubeConfig, "POST", execRequest.URL())
	if err != nil {
		return err
	}
	return exec.Stream(remotecommand.StreamOptions{
		Stdin:  options.Stdin,
		Stdout: options.Stdout,
		Stderr: options.Stderr,
	})
}
