package server

import (
	proto "github.com/Tackeyyyyyyyy/grpc-gateway-sample/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type helloService struct{}

func (hs *helloService) Echo(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	return &proto.HelloResponse{
		Message: "hello, " + req.UserName,
	}, nil
}

func Start(port string) {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("server listen: " + listen.Addr().String())
	server := grpc.NewServer()
	proto.RegisterSayHelloServer(server, &helloService{})

	if err := server.Serve(listen); err != nil {
		log.Fatalln(err)
	}
	return
}
