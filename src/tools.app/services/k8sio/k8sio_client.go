package k8sio

import (
	"fmt"
	"log"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// K8SClient inculdes utils for k8s client.
type K8SClient struct {
	clientset *kubernetes.Clientset
}

// NewK8SClient returns an instance of K8SClient.
func NewK8SClient(kubeconfig string) (*K8SClient, error) {
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}
	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	client := &K8SClient{
		clientset: clientset,
	}
	return client, nil
}

// PrintNumberOfAllPods prints the number of pods in the cluster.
func (kc *K8SClient) PrintNumberOfAllPods() error {
	pods, err := kc.clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		return err
	}
	log.Printf("There are %d pods in the cluster\n", len(pods.Items))
	return nil
}

// PrintPodInfo prints pod by namespace and name.
func (kc *K8SClient) PrintPodInfo(namespace, podname string) error {
	pod, err := kc.clientset.CoreV1().Pods(namespace).Get(podname, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		return fmt.Errorf("Pod %s in namespace %s not found", podname, namespace)
	}
	if statusError, isStatus := err.(*errors.StatusError); isStatus {
		return fmt.Errorf("Error getting pod %s in namespace %s: %v",
			podname, namespace, statusError.ErrStatus.Message)
	}
	if err != nil {
		return err
	}
	log.Printf("Found pod %s in namespace %s\n", podname, namespace)

	log.Println("Number of containers in pod:", len(pod.Status.ContainerStatuses))
	container := pod.Status.ContainerStatuses[0]
	log.Printf("Container Info: name=%s, isReady=%v, Image=%s",
		container.Name, container.Ready, container.Image)
	return nil
}
