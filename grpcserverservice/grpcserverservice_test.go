package grpcserverservice

import (
	"context"
	"log"
	"net"
	"testing"

	pb "github.com/pathak107/k8Example/k8grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
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

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)

	grpcServer := grpc.NewServer()
	s := K8GrpcServer{}
	pb.RegisterK8ServiceServer(grpcServer, &s)
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

var cs kubernetes.Interface

func TestGetPodsGrpc(t *testing.T) {
	k8s := &k8serviceMock{}
	cs = testclient.NewSimpleClientset()
	Init(k8s, cs)

	GetPodsMock = func(clientset kubernetes.Interface) ([]string, error) {
		// return []string{"pods1", "pods2", "pods3"}, nil
		return []string{"pod1", "pod2", "pod3"}, nil
	}

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewK8ServiceClient(conn)
	resp, err := client.GetPods(ctx, &pb.Request{})
	if err != nil {
		t.Fatalf("SayHello failed: %v", err)
	}
	log.Printf("Response: %+v", resp)
	// Test for output here.
}

func TestGetDeploymentsGrpc(t *testing.T) {
	k8s := &k8serviceMock{}
	cs = testclient.NewSimpleClientset()
	Init(k8s, cs)

	GetDeploymentsMock = func(clientset kubernetes.Interface) ([]string, error) {
		// return []string{"pods1", "pods2", "pods3"}, nil
		return []string{"dep1", "dep2", "dep3"}, nil
	}

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewK8ServiceClient(conn)
	resp, err := client.GetDeployments(ctx, &pb.Request{})
	if err != nil {
		t.Fatalf("SayHello failed: %v", err)
	}
	log.Printf("Response: %+v", resp)
	// Test for output here.
}

func TestGetServicesGrpc(t *testing.T) {
	k8s := &k8serviceMock{}
	cs = testclient.NewSimpleClientset()
	Init(k8s, cs)

	GetServicesMock = func(clientset kubernetes.Interface) ([]string, error) {
		// return []string{"pods1", "pods2", "pods3"}, nil
		return []string{"svc1", "svc2", "svc3"}, nil
	}

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewK8ServiceClient(conn)
	resp, err := client.GetServices(ctx, &pb.Request{})
	if err != nil {
		t.Fatalf("SayHello failed: %v", err)
	}
	log.Printf("Response: %+v", resp)
	// Test for output here.
}

func TestGetSecretsGrpc(t *testing.T) {
	k8s := &k8serviceMock{}
	cs = testclient.NewSimpleClientset()
	Init(k8s, cs)

	GetSecretsMock = func(clientset kubernetes.Interface) ([]string, error) {
		// return []string{"pods1", "pods2", "pods3"}, nil
		return []string{"sec1", "sec2", "sec3"}, nil
	}

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewK8ServiceClient(conn)
	resp, err := client.GetSecrets(ctx, &pb.Request{})
	if err != nil {
		t.Fatalf("SayHello failed: %v", err)
	}
	log.Printf("Response: %+v", resp)
	// Test for output here.
}
