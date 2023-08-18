
# go-rpc-grpc

1. protoc --go_out=.. prod.proto
2. protoc --go-grpc_out=.. prod.proto
3. protoc --grpc-gateway_out=logtostderr=true:.. prod.proto
