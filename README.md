# Overview
This project creates a grpc-go client which does round robin load balancing to send requests to multiple servers. It requires at least one grpc hello world server to be running.

# Instructions
1. Set up as many grpc greeter servers as you would like from this example:  https://github.com/grpc/grpc-go/tree/master/examples/helloworld/greeter_server. Give each server a unique port.
2. Start up the servers.
3. Change the `testAddresses` variable to match up with all of your hosts and ports.
4. Run the client code. It will send rpcs to available servers using a simple round-robin strategy.