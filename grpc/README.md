编译命令：

    在gprc目录下执行：
    protoc -I protos/ protos/hello.proto --go_out=plugins=grpc:.
    
    protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. protos/hi.proto
    
    protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. protos/hi.proto
    
    protoc --proto_path=. --descriptor_set_out=./client/hello.protoset --include_imports protos/hello.proto