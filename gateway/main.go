package gateway

import (
	"flag"
	"fmt"
	gw "github.com/Tackeyyyyyyyy/grpc-gateway-sample/proto"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func run(serverPort string, gwPort string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	endpoint := fmt.Sprintf(":" + serverPort)
	err := gw.RegisterSayHelloHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		return err
	}
	log.Printf("gateway port:" + gwPort)
	log.Printf("server listen: " + serverPort)
	return http.ListenAndServe(":"+gwPort, mux)
}

func Start(serverPort string, gwPort string) {
	flag.Parse()
	defer glog.Flush()
	if err := run(serverPort, gwPort); err != nil {
		glog.Fatal(err)
	}
}