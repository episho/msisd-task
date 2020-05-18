package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	msisdnsvc "msisdn/pkg/pb"
	"msisdn/pkg/service"
	"net"
)

func main() {

	netListener := getNetListener(1234)
	gRPCServer := grpc.NewServer()

	msisdnServer := service.NewMsisdnServer()
	msisdnsvc.RegisterMsisdnServiceServer(gRPCServer, msisdnServer)

	// start the server
	if err := gRPCServer.Serve(netListener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}

func getNetListener(port uint) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Sprintf("failed to listen: %v", err))
	}

	return lis
}



