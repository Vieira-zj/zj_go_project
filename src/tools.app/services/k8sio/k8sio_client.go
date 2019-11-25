package k8sio

import (
	"fmt"
	"log"
	"strings"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var client *K8SClient

// IsDebug a flag for printing debug message.
var IsDebug = false

// K8SClient inculdes k8s client utils.
type K8SClient struct {
	KubeConfig *restclient.Config
	KubeClient *kubernetes.Clientset
}

// NewK8SClient returns an instance of K8SClient.
func NewK8SClient(kubeConfig string) (*K8SClient, error) {
	if client != nil {
		return client, nil
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	client = &K8SClient{
		KubeConfig: config,
		KubeClient: clientset,
	}
	return client, nil
}

// ------------------------------
// Print k8s cluster info for debug.
// ------------------------------

// PrintNumberOfAllPods prints the number of all pods in cluster.
func (kc *K8SClient) PrintNumberOfAllPods() error {
	pods, err := kc.KubeClient.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		return err
	}

	numbers := len(pods.Items)
	log.Printf("All pods count: %d\n", numbers)

	names := make([]string, numbers, numbers)
	for idx, pod := range pods.Items {
		names[idx] = pod.Name
	}
	log.Println("All pods:", strings.Join(names, ","))
	return nil
}

// PrintPodInfo prints pod info by given namespace and name.
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

// GetPod returns a pod by given namespace and name.
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

	if IsDebug {
		log.Printf("Found pod [%s] in namespace [%s].\n", podName, namespace)
	}
	return pod, nil
}

// CheckPod checks the given pod is validate.
func (kc *K8SClient) CheckPod(namespace, podName, containerName string) (bool, error) {
	pod, err := kc.KubeClient.CoreV1().Pods(namespace).Get(podName, metav1.GetOptions{})
	if err != nil {
		return false, err
	}

	if pod.Status.Phase == v1.PodSucceeded || pod.Status.Phase == v1.PodFailed {
		return false, fmt.Errorf("Cannot exec in a container of [%s] pod", pod.Status.Phase)
	}

	for _, c := range pod.Spec.Containers {
		if c.Name == containerName {
			return true, nil
		}
	}
	return false, fmt.Errorf("No container [%s] found in pod [%s]", containerName, podName)
}

// ------------------------------
// Get name string for ns, pods and containers.
// ------------------------------

// GetAllNamespacesName returns list of names of all namespaces.
func (kc *K8SClient) GetAllNamespacesName() ([]string, error) {
	ns, err := kc.KubeClient.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		return []string{}, nil
	}

	retNsNames := make([]string, len(ns.Items), len(ns.Items))
	for idx, item := range ns.Items {
		retNsNames[idx] = item.Name
	}
	return retNsNames, nil
}

// GetPodNamesByNamespace returns list of names of pods in given namespace.
func (kc *K8SClient) GetPodNamesByNamespace(namespace string) ([]string, error) {
	pods, err := kc.KubeClient.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		return []string{}, nil
	}

	retPodNames := make([]string, len(pods.Items), len(pods.Items))
	for idx, pod := range pods.Items {
		retPodNames[idx] = pod.Name
	}
	return retPodNames, nil
}

// GetContainerNamesByNsAndPod returns list of names of containers in given namespace and pod.
func (kc *K8SClient) GetContainerNamesByNsAndPod(namespace, podName string) ([]string, error) {
	pod, err := kc.KubeClient.CoreV1().Pods(namespace).Get(podName, metav1.GetOptions{})
	if err != nil {
		return []string{}, err
	}

	containers := pod.Spec.Containers
	retContainerNames := make([]string, len(containers), len(containers))
	for idx, container := range containers {
		retContainerNames[idx] = container.Name
	}
	return retContainerNames, nil
}
