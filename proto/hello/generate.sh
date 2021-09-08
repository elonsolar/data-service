
# go install google.golang.org/protobuf/cmd/protoc-gen-go
#go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
#option go_package = ".;message";
protoc -I:/Users/chenxiangqian/mn/data-service/proto/hello --go-grpc_out=/Users/chenxiangqian/mn/data-service/proto/hello --go_out=/Users/chenxiangqian/mn/data-service/proto/hello  /Users/chenxiangqian/mn/data-service/proto/hello/hello.proto