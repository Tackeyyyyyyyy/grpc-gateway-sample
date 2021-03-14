# grpc-gateway-sample

gRPC-GatewayをgRPCサーバをRESTから呼び出すサンプル。


## USAGE

````````
brew install buf
protoc -I.   -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis   --go_out=plugins=grpc:. proto/sample.proto 
protoc -I.   -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis   --grpc-gateway_out=.  proto/sample.proto 
go run main.go
````````

### Call
```````
curl --location --request POST 'http://localhost:12345/api/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "test"
}'

{"encryptedId":"j4sf170oid","name":"test"}
```````
