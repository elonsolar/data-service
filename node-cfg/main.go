package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	proto "proto/hello"
	"strconv"
	"syscall"
	"time"

	grpclib "pkg/etcdv3"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var host = "127.0.0.1" //服务器主机
var (
	Port        = flag.Int("Port", 3001, "listening port")                           //服务器监听端口
	ServiceName = flag.String("ServiceName", "greet_service", "service name")        //服务名称
	EtcdAddr    = flag.String("EtcdAddr", "127.0.0.1:2379", "register etcd address") //etcd的地址
)

//rpc服务接口
type greetServer struct {
	proto.UnimplementedGreetServer
}

func (gs *greetServer) Morning(ctx context.Context, req *proto.GreetRequest) (*proto.GreetResponse, error) {
	fmt.Printf("Morning 调用: %s\n", req.Name)
	return &proto.GreetResponse{
		Message: "Good morning, " + req.Name,
		From:    fmt.Sprintf("127.0.0.1:%d", *Port),
	}, nil
}

func (gs *greetServer) Night(ctx context.Context, req *proto.GreetRequest) (*proto.GreetResponse, error) {
	fmt.Printf("Night 调用: %s\n", req.Name)
	return &proto.GreetResponse{
		Message: "Good night, " + req.Name,
		From:    fmt.Sprintf("127.0.0.1:%d", *Port),
	}, nil
}

//将服务地址注册

func main() {
	flag.Parse()

	//监听网络
	listener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", *Port))
	if err != nil {
		fmt.Println("监听网络失败：", err)
		return
	}
	defer listener.Close()

	//创建grpc句柄
	srv := grpc.NewServer()
	defer srv.GracefulStop()

	//将greetServer结构体注册到grpc服务中
	proto.RegisterGreetServer(srv, &greetServer{})

	//将服务地址注册到etcd中
	serverAddr := fmt.Sprintf("%s:%d", host, *Port)
	fmt.Printf("greeting server address: %s\n", serverAddr)
	grpclib.Register(*EtcdAddr, *ServiceName, host, strconv.Itoa(*Port), time.Second*10, 5)

	//关闭信号处理
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		s := <-ch
		grpclib.UnRegister()
		if i, ok := s.(syscall.Signal); ok {
			os.Exit(int(i))
		} else {
			os.Exit(0)
		}
	}()

	//监听服务
	err = srv.Serve(listener)
	if err != nil {
		fmt.Println("监听异常：", err)
		return
	}
}
