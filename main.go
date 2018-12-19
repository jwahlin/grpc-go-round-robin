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

// Change these addresses to match up with all of your greeter server ports
var testAddresses = []string{
	"localhost:50051",
	"localhost:50052",
}

func main() {
	// Not needed for testing, but may be needed later
	// resolver.SetDefaultScheme("dns")

	r, _ := manual.GenerateAndRegisterManualResolver()

	// Set up a connection to the server.
	conn, err := grpc.Dial(r.Scheme()+":///test.server", grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name))

	var resolvedAddrs []resolver.Address
	for i := range testAddresses {
		resolvedAddrs = append(resolvedAddrs, resolver.Address{Addr: testAddresses[i]})
	}

	r.NewAddress(resolvedAddrs)

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the servers in round-robin manner.
	for i := 0; i < 10; i++ {
		// Add pause for testing
		// duration := time.Duration(5) * time.Second // Pause for 5 seconds
		// time.Sleep(duration)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		// ctx := context.Background()
		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "world"})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.Message)
	}
}
