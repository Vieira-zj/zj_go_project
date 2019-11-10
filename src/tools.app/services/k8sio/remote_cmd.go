package k8sio

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/remotecommand"
)

// RemoteCMD runs remote commands in pod shell env.
type RemoteCMD struct {
	k8sClient *K8SClient
}

// NewRemoteCMD returns an instance of remoteCMD.
func NewRemoteCMD(kubeconfig string) (rc *RemoteCMD, err error) {
	client, err := NewK8SClient(kubeconfig)
	if err != nil {
		return nil, err
	}

	return &RemoteCMD{
		k8sClient: client,
	}, nil
}

// RemoteExecOption contains the options for exec remote command.
type RemoteExecOption struct {
	Namespace string
	PodName   string
	Command   string
	Stdin     io.Reader
	Stdout    io.Writer
	Stderr    io.Writer
}

// RemoteExecAndPrint executes a shell command on a given pod in a namespace, and prints results.
func (rc *RemoteCMD) RemoteExecAndPrint(namespace, podName, command string) error {
	return rc.RemoteExec(RemoteExecOption{
		Namespace: namespace,
		PodName:   podName,
		Command:   command,
		Stdin:     strings.NewReader(""),
		Stdout:    os.Stdout,
		Stderr:    os.Stderr,
	})
}

// RemoteExec executes a shell command on a given pod in a namespace.
func (rc *RemoteCMD) RemoteExec(options RemoteExecOption) error {
	pod, err := rc.k8sClient.GetPod(options.Namespace, options.PodName)
	if err != nil {
		return err
	}

	containerName := pod.Spec.Containers[0].Name
	podExecOptions := &v1.PodExecOptions{
		Container: containerName,
		Command:   strings.Split(options.Command, " "),
		Stdout:    true,
		Stderr:    true,
		Stdin:     true,
	}
	execRequest := rc.k8sClient.KubeClient.CoreV1().RESTClient().Post().
		Resource("pods").Name(options.PodName).Namespace(options.Namespace).SubResource("exec")
	execRequest.VersionedParams(podExecOptions, scheme.ParameterCodec)

	log.Printf("URL: %s\n", execRequest.URL())
	fmt.Printf("Executing command \"%s\" on container [%s] in pod [%s].\n",
		options.Command, containerName, options.PodName)

	exec, err := remotecommand.NewSPDYExecutor(rc.k8sClient.KubeConfig, "POST", execRequest.URL())
	if err != nil {
		return err
	}
	if err = exec.Stream(remotecommand.StreamOptions{
		Stdin:  options.Stdin,
		Stdout: options.Stdout,
		Stderr: options.Stderr,
	}); err != nil {
		return err
	}

	return nil
}
