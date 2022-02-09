package main

import (
	"fmt"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/pathak107/k8Example/k8grpc"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	// conn, err := grpc.Dial("192.168.64.2:30001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	client := pb.NewK8ServiceClient(conn)

	pods, err := client.GetPods(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("Error when calling GetPods grpc: %s", err)
	}
	fmt.Printf("\nCalling grpc GetPods:\nResponse:\n %s\n", pods.Pods[:])

	dep, err := client.GetDeployments(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("Error when calling GetDeployments: %s", err)
	}
	fmt.Printf("\nCalling grpc Deployments:\nResponse:\n %s\n", dep.Deployments[:])

	secrets, err := client.GetSecrets(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("Error when calling GetSecrets: %s", err)
	}
	fmt.Printf("\nCalling grpc GetSecrets:\nResponse:\n %s\n", secrets.Secrets[:])

	services, err := client.GetServices(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("Error when calling getservices: %s", err)
	}
	fmt.Printf("\nCalling grpc GetServices:\nResponse:\n %s\n", services.Services[:])
}
