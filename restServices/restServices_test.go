package restServices

import (
	"errors"
	"fmt"
	"testing"

	"net/http"
	"net/http/httptest"

	"k8s.io/client-go/kubernetes"
	testclient "k8s.io/client-go/kubernetes/fake"
)

var GetPodsMock func(clientset kubernetes.Interface) ([]string, error)
var GetDeploymentsMock func(clientset kubernetes.Interface) ([]string, error)
var GetSecretsMock func(clientset kubernetes.Interface) ([]string, error)
var GetServicesMock func(clientset kubernetes.Interface) ([]string, error)

type k8serviceMock struct{}

func (k *k8serviceMock) GetPods(clientset kubernetes.Interface) ([]string, error) {
	return GetPodsMock(cs)
}
func (k *k8serviceMock) GetDeployments(clientset kubernetes.Interface) ([]string, error) {
	return GetDeploymentsMock(cs)
}
func (k *k8serviceMock) GetSecrets(clientset kubernetes.Interface) ([]string, error) {
	return GetSecretsMock(cs)
}
func (k *k8serviceMock) GetServices(clientset kubernetes.Interface) ([]string, error) {
	return GetServicesMock(cs)
}

var cs kubernetes.Interface

func TestGetPods(t *testing.T) {
	k8s := &k8serviceMock{}
	cs = testclient.NewSimpleClientset()
	InitRoutes(clientset, k8s)

	t.Run("Return a list of pods", func(t *testing.T) {
		GetPodsMock = func(clientset kubernetes.Interface) ([]string, error) {
			// return []string{"pods1", "pods2", "pods3"}, nil
			return []string{"pod1", "pod2", "pod3"}, nil
		}

		req, err := http.NewRequest("GET", "/pods", nil)
		if err != nil {
			t.Fatal(err)
		}
		res := httptest.NewRecorder()
		handler := http.HandlerFunc(GetPodsHandler)
		handler.ServeHTTP(res, req)

		// Check the status code is what we expect.
		if status := res.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		// Check the response body is what we expect.
		expected := `{"Status":"success","Err":"","Data":["pod1","pod2","pod3"]}`
		if res.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				res.Body.String(), expected)
		}
	})

	t.Run("Return error from getPods function", func(t *testing.T) {
		GetPodsMock = func(clientset kubernetes.Interface) ([]string, error) {
			// return []string{"pods1", "pods2", "pods3"}, nil
			return nil, errors.New("Some server error in getting pods")
		}

		req, err := http.NewRequest("GET", "/pods", nil)
		if err != nil {
			t.Fatal(err)
		}
		res := httptest.NewRecorder()
		handler := http.HandlerFunc(GetPodsHandler)
		handler.ServeHTTP(res, req)

		// Check the status code is what we expect.
		if status := res.Code; status != http.StatusInternalServerError {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		// Check the response body is what we expect.
		expected := `{"Status":"failure","Err":"Some server error in getting pods","Data":null}`
		if res.Body.String() != expected {
			fmt.Print(res.Body.String() == expected)
			t.Errorf("handler returned unexpected body: got %v want %v",
				res.Body.String(), expected)
		}
	})

}

func TestGetDeployments(t *testing.T) {
	k8s := &k8serviceMock{}
	cs = testclient.NewSimpleClientset()
	InitRoutes(clientset, k8s)

	t.Run("Return a list of deployments", func(t *testing.T) {
		GetDeploymentsMock = func(clientset kubernetes.Interface) ([]string, error) {
			// return []string{"pods1", "pods2", "pods3"}, nil
			return []string{"dep1", "dep2", "dep3"}, nil
		}

		req, err := http.NewRequest("GET", "/deployments", nil)
		if err != nil {
			t.Fatal(err)
		}
		res := httptest.NewRecorder()
		handler := http.HandlerFunc(GetDeploymentsHandler)
		handler.ServeHTTP(res, req)

		// Check the status code is what we expect.
		if status := res.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		// Check the response body is what we expect.
		expected := `{"Status":"success","Err":"","Data":["dep1","dep2","dep3"]}`
		if res.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				res.Body.String(), expected)
		}
	})

	t.Run("Return error from getDeployments function", func(t *testing.T) {
		GetDeploymentsMock = func(clientset kubernetes.Interface) ([]string, error) {
			return nil, errors.New("Some server error")
		}

		req, err := http.NewRequest("GET", "/deployments", nil)
		if err != nil {
			t.Fatal(err)
		}
		res := httptest.NewRecorder()
		handler := http.HandlerFunc(GetDeploymentsHandler)
		handler.ServeHTTP(res, req)

		// Check the status code is what we expect.
		if status := res.Code; status != http.StatusInternalServerError {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		// Check the response body is what we expect.
		expected := `{"Status":"failure","Err":"Some server error","Data":null}
		`
		fmt.Println(res.Body.String())
		if res.Body.String() != expected {
			fmt.Print(res.Body.String() == expected)
			t.Errorf("handler returned unexpected body: got %v want %v",
				res.Body.String(), expected)
		}
	})

}
