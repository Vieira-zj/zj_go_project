package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	mysvc "tools.app/services/k8sio"
)

func main() {
	log.Println("K8S Client app start.")

	const desc = "(optional) absolute path to the kubeconfig file"
	defaultPath := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	kubeconfig := flag.String("kubeconfig", defaultPath, desc)
	flag.Parse()

	printClusterInfo(*kubeconfig)
	execRemoteCommand(*kubeconfig)
}

func printClusterInfo(kubeconfig string) {
	client, err := mysvc.NewK8SClient(kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	if err := client.PrintNumberOfAllPods(); err != nil {
		panic(err.Error())
	}

	const (
		namespace = "mini-test-ns"
		podname   = "hello-minikube-59ddd8676b-vkl26"
	)
	if err := client.PrintPodInfo(namespace, podname); err != nil {
		panic(err.Error())
	}
}

func execRemoteCommand(kubeconfig string) {
	rc, err := mysvc.NewRemoteCMD(kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	var (
		namespace = "mini-test-ns"
		podname   = "hello-minikube-59ddd8676b-vkl26"
	)
	if err := rc.RemoteExecAndPrint(namespace, podname, "ls -l"); err != nil {
		panic(err.Error())
	}
}
