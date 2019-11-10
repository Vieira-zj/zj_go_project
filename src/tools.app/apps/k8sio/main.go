package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	mysvc "tools.app/services/k8sio"
)

func main() {
	// Refer: https://github.com/kubernetes/kubernetes/tree/v1.6.1/test/e2e/framework
	log.Println("K8S Client app start.")

	const desc = "(optional) absolute path to the kubeconfig file"
	defaultPath := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	kubeconfig := flag.String("kubeconfig", defaultPath, desc)
	flag.Parse()

	execCliCommand()
	if false {
		printClusterInfo(*kubeconfig)
		execRemoteCommand(*kubeconfig)
	}
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

func execCliCommand() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("go-cli$ ")
		cmdStr, err := reader.ReadString('\n')
		if err != nil {
			panic(err.Error())
		}
		if err = runCommand(cmdStr); err != nil {
			panic(err.Error())
		}
	}
}

func runCommand(cmdStr string) error {
	cmdStr = strings.TrimSuffix(cmdStr, "\n")
	cmdSlice := strings.Fields(cmdStr)
	switch cmdSlice[0] {
	case "exit":
		log.Println("go-cli shell exit.")
		os.Exit(0)
	}
	cmd := exec.Command(cmdSlice[0], cmdSlice[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
