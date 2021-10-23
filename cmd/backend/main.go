package main

import (
	"fmt"
	"log"
	"net"

	"github.com/TrashRaiders/garbage-truck/pkg/protos"
	"github.com/TrashRaiders/garbage-truck/pkg/shop"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	logger := logrus.StandardLogger()

	port := 8080
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	protos.RegisterShopServiceServer(grpcServer, shop.NewService())

	logger.Infof("backend listening on port: %v", port)

	grpcServer.Serve(lis)
}
