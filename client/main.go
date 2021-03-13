package client

import (
	"context"
	"fmt"
	proto "github.com/Tackeyyyyyyyy/grpc-gateway-sample/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

func Echo(conn *grpc.ClientConn, name string) {
	client := proto.NewSayHelloClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, err := client.Echo(ctx, &proto.HelloRequest{UserName: name})
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("message: %s\n", resp.Message)
}

func Call(port string) {
	addr := fmt.Sprintf("localhost:" + port)

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	Echo(conn, "Tackeyyyyyyyy !!")
	defer conn.Close()
}
