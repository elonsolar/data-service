module github.com/elonsolar/data-service/node-cfg

go 1.16

require (
	go.etcd.io/etcd/client/v3 v3.5.0
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4
	google.golang.org/grpc v1.40.0
	pkg v0.0.1
	proto v0.0.1
)

replace proto v0.0.1 => ../proto

replace pkg v0.0.1 => ../pkg
