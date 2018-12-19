package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/resolver/manual"
)

const (
	addressOne  = "localhost:50051"
	addressTwo  = "localhost:50052"
	defaultName = "world"
)

func main() {
	// The secret sauce, not needed for testing
	// resolver.SetDefaultScheme("dns")

	r, _ := manual.GenerateAndRegisterManualResolver()

	// Set up a connection to the server.
	conn, err := grpc.Dial(r.Scheme()+":///test.server", grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name))

	var resolvedAddrs []resolver.Address
	resolvedAddrs = append(resolvedAddrs, resolver.Address{Addr: "localhost:50051"})
	resolvedAddrs = append(resolvedAddrs, resolver.Address{Addr: "localhost:50052"})

	r.NewAddress(resolvedAddrs)

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName

	// Contact the servers in round-robin manner.
	for i := 0; i < 10; i++ {
		// Add pause for testing
		// duration := time.Duration(5) * time.Second // Pause for 5 seconds
		// time.Sleep(duration)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		// ctx := context.Background()
		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.Message)
	}
}
