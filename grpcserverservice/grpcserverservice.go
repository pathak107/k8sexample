package grpcserverservice

import (
	"context"

	pb "github.com/pathak107/k8Example/k8grpc"
	"github.com/pathak107/k8Example/k8service"
	"k8s.io/client-go/kubernetes"

	"fmt"
)

type K8GrpcServer struct {
	pb.UnimplementedK8ServiceServer
}

var k8s k8service.K8serviceInterface
var clientset kubernetes.Interface

func Init(k8svc k8service.K8serviceInterface, cs kubernetes.Interface) {
	clientset = cs
	k8s = k8svc
}

func (s *K8GrpcServer) GetPods(ctx context.Context, in *pb.Request) (*pb.PodList, error) {
	pods, err := k8s.GetPods(clientset)
	if err != nil {
		fmt.Printf("Error occured in retrieving pods : %v", err)
		return nil, err
	}
	return &pb.PodList{Pods: pods}, nil
}

func (s *K8GrpcServer) GetDeployments(ctx context.Context, in *pb.Request) (*pb.DeploymentList, error) {
	deps, err := k8s.GetDeployments(clientset)
	if err != nil {
		fmt.Printf("Error occured in retrieving deployments : %v", err)
		return nil, err
	}
	return &pb.DeploymentList{Deployments: deps}, nil
}
func (s *K8GrpcServer) GetServices(ctx context.Context, in *pb.Request) (*pb.ServiceList, error) {
	services, err := k8s.GetServices(clientset)
	if err != nil {
		fmt.Printf("Error occured in retrieving services : %v", err)
		return nil, err
	}
	return &pb.ServiceList{Services: services}, nil
}
func (s *K8GrpcServer) GetSecrets(ctx context.Context, in *pb.Request) (*pb.SecretList, error) {
	secrets, err := k8s.GetSecrets(clientset)
	if err != nil {
		fmt.Printf("Error occured in retrieving secrets : %v", err)
		return nil, err
	}
	return &pb.SecretList{Secrets: secrets}, nil
}
