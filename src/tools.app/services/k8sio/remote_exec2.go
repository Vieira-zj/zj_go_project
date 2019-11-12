package k8sio

import (
	"bytes"
	"io"
	"log"
	"net/url"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
)

// RemoteExec2 runs command on given pod in namespace.
type RemoteExec2 struct {
	k8sClient *K8SClient
}

// NewRemoteExec2 returns an instance of RemoteExec2.
func NewRemoteExec2(kubeConfig string) (*RemoteExec2, error) {
	client, err := NewK8SClient(kubeConfig)
	if err != nil {
		return nil, err
	}

	return &RemoteExec2{
		k8sClient: client,
	}, nil
}

// ExecOptions2 includes exec options, and passed to ExecWithOptions.
type ExecOptions2 struct {
	Command       []string
	Namespace     string
	PodName       string
	ContainerName string

	Stdin              io.Reader
	CaptureStdout      bool
	CaptureStderr      bool
	PreserveWhitespace bool
}

// RunCommandInPod runs a command in given pod, and returns stdout and stderr string.
func (re2 *RemoteExec2) RunCommandInPod(namespace, podName string, cmd ...string) (string, string, error) {
	pod, err := re2.k8sClient.GetPod(namespace, podName)
	if err != nil {
		return "", "", err
	}
	return re2.RunCommandInContainer(namespace, podName, pod.Spec.Containers[0].Name, cmd...)
}

// RunCommandInContainer runs a command in given container, and returns stdout and stderr string.
func (re2 *RemoteExec2) RunCommandInContainer(
	namespace, podName, containerName string, cmd ...string) (string, string, error) {
	return re2.ExecWithOptions(ExecOptions2{
		Command:       cmd,
		Namespace:     namespace,
		PodName:       podName,
		ContainerName: containerName,

		Stdin:         nil,
		CaptureStdout: true,
		CaptureStderr: true,
	})
}

// ExecWithOptions executes a command in given container, and returns stdout and stderr string.
func (re2 *RemoteExec2) ExecWithOptions(options ExecOptions2) (string, string, error) {
	const tty = false
	req := re2.k8sClient.KubeClient.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(options.PodName).
		Namespace(options.Namespace).
		SubResource("exec").
		Param("container", options.ContainerName)
	req.VersionedParams(&v1.PodExecOptions{
		Container: options.ContainerName,
		Command:   options.Command,
		Stdin:     options.Stdin != nil,
		Stdout:    options.CaptureStdout,
		Stderr:    options.CaptureStderr,
		TTY:       tty,
	}, scheme.ParameterCodec)

	if IsDebug {
		log.Printf("URL: %s\n", req.URL())
		log.Printf("Executing command \"%s\" on container [%s] in pod [%s].\n",
			options.Command, options.ContainerName, options.PodName)
	}

	var stdout, stderr bytes.Buffer
	err := execute("POST", req.URL(), re2.k8sClient.KubeConfig, options.Stdin, &stdout, &stderr, tty)
	return stdout.String(), stderr.String(), err
}

func execute(method string, url *url.URL, config *restclient.Config, stdin io.Reader, stdout, stderr io.Writer, tty bool) error {
	exec, err := remotecommand.NewSPDYExecutor(config, method, url)
	if err != nil {
		return err
	}

	return exec.Stream(remotecommand.StreamOptions{
		Stdin:  stdin,
		Stdout: stdout,
		Stderr: stderr,
		Tty:    tty,
	})
}
