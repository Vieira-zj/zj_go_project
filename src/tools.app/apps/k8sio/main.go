package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	mysvc "tools.app/services/k8sio"
)

func main() {
	log.Println("k8s client app start.")

	defaultPath := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	desc := "(optional) absolute path to the kubeconfig file"
	kubeconfig := flag.String("kubeconfig", defaultPath, desc)
	flag.Parse()

	client, err := mysvc.NewK8SClient(*kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	client.PrintNumberOfAllPods()
	client.PrintPodInfo("mini-test-ns", "hello-minikube-59ddd8676b-vkl26")
}
