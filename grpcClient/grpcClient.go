package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/pathak107/k8Example/k8grpc"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial("192.168.64.2:30001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	client := pb.NewK8ServiceClient(conn)

	pods, err := client.GetPods(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", pods.Pods[:])

	dep, err := client.GetDeployments(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", dep.Deployments[:])

	secrets, err := client.GetSecrets(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", secrets.Secrets[:])

	services, err := client.GetServices(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", services.Services[:])
}
