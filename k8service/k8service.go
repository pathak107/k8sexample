package k8service

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	clientcmd "k8s.io/client-go/tools/clientcmd"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type K8serviceInterface interface {
	GetPods(clientset kubernetes.Interface) ([]string, error)
	GetDeployments(clientset kubernetes.Interface) ([]string, error)
	GetSecrets(clientset kubernetes.Interface) ([]string, error)
	GetServices(clientset kubernetes.Interface) ([]string, error)
}

type K8serviceStruct struct{}

func NewK8Service() K8serviceInterface {
	K8service := &K8serviceStruct{}
	return K8service
}

func ConnectToK8Cluster() (*kubernetes.Clientset, error) {
	//var config *rest.Config
	config, err := rest.InClusterConfig()
	if err != nil {
		home, exists := os.LookupEnv("HOME")
		if !exists {
			home = "/root"
		}
		configPath := filepath.Join(home, ".kube", "config")
		config, err = clientcmd.BuildConfigFromFlags("", configPath)
		if err != nil {
			fmt.Printf("The kubeconfig cannot be loaded: %v\n", err)
			return nil, err
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("The kubeconfig cannot be loaded: %v\n", err)
		return nil, err
	}

	return clientset, nil
}

func (k *K8serviceStruct) GetPods(clientset kubernetes.Interface) ([]string, error) {
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	var podNames []string
	for _, pod := range pods.Items {
		podNames = append(podNames, pod.Name)
	}
	return podNames, nil
}
func (k *K8serviceStruct) GetDeployments(clientset kubernetes.Interface) ([]string, error) {
	deployments, err := clientset.AppsV1().Deployments("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Error in retrieving deployments: %v", err.Error())
		return nil, err
	}
	var deploymentNames []string
	for _, deployment := range deployments.Items {
		deploymentNames = append(deploymentNames, deployment.Name)
	}
	return deploymentNames, nil
}
func (k *K8serviceStruct) GetSecrets(clientset kubernetes.Interface) ([]string, error) {
	secrets, err := clientset.CoreV1().Secrets("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	var secretNames []string
	for _, secret := range secrets.Items {
		secretNames = append(secretNames, secret.Name)
	}
	return secretNames, nil
}
func (k *K8serviceStruct) GetServices(clientset kubernetes.Interface) ([]string, error) {
	services, err := clientset.CoreV1().Services("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	var serviceNames []string
	for _, service := range services.Items {
		serviceNames = append(serviceNames, service.Name)
	}
	return serviceNames, nil
}
