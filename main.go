package main

import (
	"log"
	"net"
	"net/http"

	// "net/http"

	"flag"
	"fmt"

	"github.com/pathak107/k8Example/grpcserverservice"
	pb "github.com/pathak107/k8Example/k8grpc"
	"github.com/pathak107/k8Example/k8service"

	"github.com/pathak107/k8Example/restServices"
	"google.golang.org/grpc"
)

func main() {
	clientset, err := k8service.ConnectToK8Cluster()
	if err != nil {
		log.Fatalf("Failed to create clientset with error %v", err)
	}
	k8s := k8service.NewK8Service()

	var server string
	flag.StringVar(&server, "server", "rest", "State the server type server=rest or server=grpc")
	flag.Parse()

	if server == "rest" {
		restServices.InitRoutes(clientset, k8s)
		log.Fatal(http.ListenAndServe(":3000", nil))

	} else {
		grpcserverservice.Init(k8s, clientset)
		lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", 9000))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		grpcServer := grpc.NewServer()

		s := grpcserverservice.K8GrpcServer{}

		pb.RegisterK8ServiceServer(grpcServer, &s)

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %s", err)
		}
	}

}
