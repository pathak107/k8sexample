package restServices

import (
	"fmt"
	"net/http"

	"github.com/pathak107/k8Example/k8service"
	"k8s.io/client-go/kubernetes"

	"encoding/json"
)

type response struct {
	Status string
	Err    string
	Data   []string
}

var clientset kubernetes.Interface
var k8s k8service.K8serviceInterface

func InitRoutes(cs kubernetes.Interface, k8service k8service.K8serviceInterface) {
	clientset = cs
	k8s = k8service
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/pods", GetPodsHandler)
	http.HandleFunc("/deployments", GetDeploymentsHandler)
	http.HandleFunc("/services", GetServicesHandler)
	http.HandleFunc("/secrets", GetSecretsHandler)
}

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Welcome to K8 Example")
}

func GetPodsHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("content-type", "application/json")
	podNames, err := k8s.GetPods(clientset)
	if err != nil {
		fmt.Print(err.Error())
		res := &response{"failure", err.Error(), podNames}
		out, _ := json.Marshal(res)
		http.Error(rw, string(out), http.StatusInternalServerError)
		return
	}

	res := &response{"success", "", podNames}
	out, _ := json.Marshal(res)
	fmt.Fprint(rw, string(out))
}

func GetDeploymentsHandler(rw http.ResponseWriter, r *http.Request) {
	deploymentNames, err := k8s.GetDeployments(clientset)
	if err != nil {
		fmt.Print(err.Error())
		res := &response{"failure", err.Error(), deploymentNames}
		out, _ := json.Marshal(res)
		http.Error(rw, string(out), http.StatusInternalServerError)
		return
	}
	rw.Header().Set("content-type", "application/json")
	res := response{"success", "", deploymentNames}
	out, _ := json.Marshal(res)
	fmt.Fprint(rw, string(out))
}

func GetServicesHandler(rw http.ResponseWriter, r *http.Request) {
	serviceNames, err := k8s.GetServices(clientset)
	if err != nil {
		fmt.Print(err.Error())
		res := &response{"failure", err.Error(), serviceNames}
		out, _ := json.Marshal(res)
		http.Error(rw, string(out), http.StatusInternalServerError)
		return
	}
	rw.Header().Set("content-type", "application/json")
	res := response{"success", "", serviceNames}
	out, _ := json.Marshal(res)
	fmt.Fprint(rw, string(out))
}

func GetSecretsHandler(rw http.ResponseWriter, r *http.Request) {
	secretNames, err := k8s.GetSecrets(clientset)
	if err != nil {
		fmt.Print(err.Error())
		res := &response{"failure", err.Error(), secretNames}
		out, _ := json.Marshal(res)
		http.Error(rw, string(out), http.StatusInternalServerError)
		return
	}
	rw.Header().Set("content-type", "application/json")
	res := response{"success", "", secretNames}
	out, _ := json.Marshal(res)
	fmt.Fprint(rw, string(out))
}
