package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bogatyr285/test_task/aws"
	grpcServer "github.com/bogatyr285/test_task/grpc"
)

const (
	grpcServerAddr = ":5000"
)

func main() {
	awsUploader := &aws.AWSUploader{}
	grpcFilesTmpDir := "./tmp"
	grpcController := grpcServer.NewGRPCFilesAPIController(awsUploader, grpcFilesTmpDir)
	grpcServer, err := grpcServer.NewGRPCServer(grpcServerAddr, grpcController)
	if err != nil {
		panic(err)
	}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	grpcServer.Start()

	<-sigchan

	grpcServer.Stop()
}
