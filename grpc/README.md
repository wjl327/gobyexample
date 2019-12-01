编译命令：

    在gprc目录下执行：
    protoc -I protos/ protos/hello.proto --go_out=plugins=grpc:.
    
    protoc -I. -I/usr/local/include  \
      -I$GOPATH/src \
      -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
      --grpc-gateway_out=logtostderr=true:. \
      protos/hi.proto
      
    protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. protos/hi.proto
    
    protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.12.1/third_party/googleapis --go_out=plugins=grpc:. protos/hi.proto