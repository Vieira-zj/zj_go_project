package k8sio

import (
	"fmt"
	"log"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// K8SClient inculdes utils for k8s client.
type K8SClient struct {
	KubeConfig *restclient.Config
	KubeClient *kubernetes.Clientset
}

// NewK8SClient returns an instance of K8SClient.
func NewK8SClient(kubeconfig string) (*K8SClient, error) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	client := &K8SClient{
		KubeConfig: config,
		KubeClient: clientset,
	}
	return client, nil
}

// PrintNumberOfAllPods prints the number of all pods in cluster.
func (kc *K8SClient) PrintNumberOfAllPods() error {
	pods, err := kc.KubeClient.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		return err
	}
	log.Printf("There are %d pods in the cluster.\n", len(pods.Items))
	return nil
}

// PrintPodInfo prints pod info by namespace and name.
func (kc *K8SClient) PrintPodInfo(namespace, podName string) error {
	pod, err := kc.GetPod(namespace, podName)
	if err != nil {
		return err
	}

	log.Printf("Number of containers in pod [%s]: %d\n", podName, len(pod.Status.ContainerStatuses))
	container := pod.Status.ContainerStatuses[0]
	log.Printf("Container Info: name=%s, isReady=%v, Image=%s",
		container.Name, container.Ready, container.Image)
	return nil
}

// GetPod returns a pod by namespace and name.
func (kc *K8SClient) GetPod(namespace, podName string) (pod *v1.Pod, err error) {
	pod, err = kc.KubeClient.CoreV1().Pods(namespace).Get(podName, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		return nil, fmt.Errorf("Pod [%s] in namespace [%s] not found", podName, namespace)
	}
	if statusError, isStatus := err.(*errors.StatusError); isStatus {
		return nil, fmt.Errorf("Error getting pod [%s] in namespace [%s]: %v",
			podName, namespace, statusError.ErrStatus.Message)
	}
	if err != nil {
		return nil, err
	}

	log.Printf("Found pod [%s] in namespace [%s].\n", podName, namespace)
	return pod, nil
}
