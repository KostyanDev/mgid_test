package main

import (
	"google.golang.org/grpc"
	"net"
)

func main() {

	conn, err := grpc.Dial(net.JoinHostPort(cfg.DevicesInternalGRPCConfig.Host, cfg.DevicesInternalGRPCConfig.Port), grpc.WithInsecure())
	if err != nil {
		logger.Fatal(err)
	}

}
