package k8service

import (
	"context"
	"testing"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	testclient "k8s.io/client-go/kubernetes/fake"
)

func TestGetPods(t *testing.T) {
	K8service := NewK8Service()
	clientset := testclient.NewSimpleClientset()

	pod := &v1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-pod",
			Namespace: "default",
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:            "nginx",
					Image:           "nginx",
					ImagePullPolicy: "Always",
				},
			},
		},
	}
	clientset.CoreV1().Pods(pod.Namespace).Create(context.TODO(), pod, metav1.CreateOptions{})
	pods, err := K8service.GetPods(clientset)
	if err != nil {
		t.Errorf("Test failed with error: %v", err)
	} else {
		if pods[0] != pod.Name {
			t.Errorf("Test failed expected pod name: %v , got : %v", pod.Name, pods[0])
		}
	}

}

func int32Ptr(i int32) *int32 { return &i }
func TestGetDeployments(t *testing.T) {
	K8service := NewK8Service()
	clientset := testclient.NewSimpleClientset()

	dep := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-deployment",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "demo",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "demo",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "web",
							Image: "nginx:1.12",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}
	clientset.AppsV1().Deployments(dep.Namespace).Create(context.TODO(), dep, metav1.CreateOptions{})
	deps, err := K8service.GetDeployments(clientset)
	if err != nil {
		t.Errorf("Test failed with error: %v", err)
	} else {
		if deps[0] != dep.Name {
			t.Errorf("Test failed expected pod name: %v , got : %v", dep.Name, deps[0])
		}
	}

}

func TestGetServices(t *testing.T) {
	K8service := NewK8Service()
	clientset := testclient.NewSimpleClientset()

	service := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "myservice",
			Namespace: "default",
			Labels: map[string]string{
				"app": "myapp",
			},
		},
		Spec: v1.ServiceSpec{
			Ports:     nil,
			Selector:  nil,
			ClusterIP: "",
		}}
	clientset.CoreV1().Services(service.Namespace).Create(context.TODO(), service, metav1.CreateOptions{})
	services, err := K8service.GetServices(clientset)
	if err != nil {
		t.Errorf("Test failed with error: %v", err)
	} else {
		if services[0] != service.Name {
			t.Errorf("Test failed expected service name: %v , got : %v", service.Name, services[0])
		}
	}

}

func TestGetSecrets(t *testing.T) {
	K8service := NewK8Service()
	clientset := testclient.NewSimpleClientset()

	secret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "pull-image-secret",
			Namespace: "aaaaaa",
		},
		Type: "kubernetes.io/dockerconfigjson",
		Data: map[string][]byte{".dockerconfigjson": nil},
	}
	clientset.CoreV1().Secrets(secret.Namespace).Create(context.TODO(), secret, metav1.CreateOptions{})
	secrets, err := K8service.GetSecrets(clientset)
	if err != nil {
		t.Errorf("Test failed with error: %v", err)
	} else {
		if secrets[0] != secret.Name {
			t.Errorf("Test failed expected secret name: %v , got : %v", secret.Name, secrets[0])
		}
	}

}
