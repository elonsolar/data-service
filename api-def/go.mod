module github.com/elonsolar/data-service/api-def

go 1.16

require (
	// github.com/etcd-io/etcd v3.3.24+incompatible
	go.etcd.io/etcd/client/v3 v3.5.0
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4
	google.golang.org/grpc v1.40.0
	proto v0.0.1
	pkg v0.0.1
)

replace pkg v0.0.1 => ../pkg

replace proto v0.0.1 => ../proto

// replace go.etcd.io/etcd/mvcc/mvccpb =>       github.com/etcd-io/etcd/mvcc/mvccpb v3.3.24+incompatible
