
# go-rpc-grpc

protoc --go_out=.. prod.proto
protoc --go-grpc_out=.. prod.proto
protoc --grpc-gateway_out=logtostderr=true:.. prod.proto
