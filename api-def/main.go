package main

import (
	"flag"
	"fmt"
	proto "proto/hello"
	"time"

	// "github.com/coreos/etcd/mvcc/mvccpb"
	// mvccpb "github.com/etcd-io/etcd/mvcc/mvccpb"
	grpclib "pkg/etcdv3"

	clientv3 "go.etcd.io/etcd/client/v3"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

const schema = "ns"

var (
	ServiceName = flag.String("ServiceName", "greet_service", "service name")        //服务名称
	EtcdAddr    = flag.String("EtcdAddr", "127.0.0.1:2379", "register etcd address") //etcd的地址
)

var cli *clientv3.Client

func main() {
	flag.Parse()

	//注册etcd解析器
	r := grpclib.NewResolver(*EtcdAddr, *ServiceName)
	resolver.Register(r)

	//客户端连接服务器(负载均衡：轮询) 会同步调用r.Build()
	conn, err := grpc.Dial(r.Scheme()+"://author/"+*ServiceName, grpc.WithBalancerName("round_robin"), grpc.WithInsecure())
	if err != nil {
		fmt.Println("连接服务器失败：", err)
	}
	defer conn.Close()

	//获得grpc句柄
	c := proto.NewGreetClient(conn)
	ticker := time.NewTicker(1 * time.Second)
	for range ticker.C {
		fmt.Println("Morning 调用...")
		resp1, err := c.Morning(
			context.Background(),
			&proto.GreetRequest{Name: "JetWu"},
		)
		if err != nil {
			fmt.Println("Morning调用失败：", err)
			return
		}
		fmt.Printf("Morning 响应：%s，来自：%s\n", resp1.Message, resp1.From)

		fmt.Println("Night 调用...")
		resp2, err := c.Night(
			context.Background(),
			&proto.GreetRequest{Name: "JetWu"},
		)
		if err != nil {
			fmt.Println("Night调用失败：", err)
			return
		}
		fmt.Printf("Night 响应：%s，来自：%s\n", resp2.Message, resp2.From)
	}
}
