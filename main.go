package main

import (
	client "github.com/Tackeyyyyyyyy/grpc-gateway-sample/client"
	gateway "github.com/Tackeyyyyyyyy/grpc-gateway-sample/gateway"
	server "github.com/Tackeyyyyyyyy/grpc-gateway-sample/server"
	"time"
)

func main() {
	serverPort := "19203"
	gwPort := "12345"
	go func() {
		server.Start(serverPort)
	}()
	time.Sleep(time.Second * 2)
	client.Call(serverPort)
	gateway.Start(serverPort, gwPort)
	return
}