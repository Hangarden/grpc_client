package main

import (
	sampleV1 "awesomeProject/sample/client/sample/v1"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

const (
	sampleServerHost = "localhost"
	sampleServerPort = 5000
)

func main() {
	ctx := context.Background()
	ctxTimeout, _ := context.WithTimeout(ctx, time.Second*5)

	fmt.Println("Connect gRPC Sample Server : ", sampleServerHost, sampleServerPort)
	sampleServerAddress := fmt.Sprintf("%v:%v", sampleServerHost, sampleServerPort)
	var err error

	sampleClientConn, err := grpc.DialContext(ctxTimeout, sampleServerAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Failed to dial gRPC Sample Server:", err)
	}
	sampleClient := sampleV1.NewSampleServiceClient(sampleClientConn)

	request := &sampleV1.GetInfoInfoRequest{
		SendMessage: "hello",
	}

	resMessage, er := sampleClient.GetInfo(ctxTimeout, request)
	if err != nil {
		fmt.Println("GetInfo Error:", er)
	} else {
		fmt.Println("GetInfo Data:", resMessage)
	}
}
