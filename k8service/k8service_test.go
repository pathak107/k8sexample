package k8service

import (
	"context"
	"testing"

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
