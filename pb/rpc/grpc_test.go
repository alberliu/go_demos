package rpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"testing"
	"time"
)

// protoc --go_out=plugins=grpc:. *.proto

// 服务器端的单向调用的拦截器
func UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("before handling. Info: %+v", info)
	log.Printf("before handling. req: %+v", req)
	resp, err := handler(ctx, req)
	log.Printf("after handling. resp: %+v", resp)
	return resp, err
}

// 服务器端stream调用的拦截器
func StreamServerInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Printf("before handling. Info: %+v", info)
	err := handler(srv, ss)
	log.Printf("after handling. err: %v", err)
	return err
}

type server struct{}

const port = ":50001"

func (s *server) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	fmt.Println("ctx", md, ok)
	return &HelloReply{Message: port}, nil
}

func TestServer(t *testing.T) {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.StreamInterceptor(StreamServerInterceptor),
		grpc.UnaryInterceptor(UnaryServerInterceptor))
	RegisterGreeterServer(s, &server{})
	log.Println("rpc服务已经开启")
	s.Serve(lis)
}

func TestClientNginxProxy(t *testing.T) {
	conn, err := grpc.Dial("localhost:80", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := NewGreeterClient(conn)
	fmt.Println(c.SayHello(metadata.NewOutgoingContext(context.TODO(), metadata.Pairs("key", "val")), &HelloRequest{Name: "hello"}))
	fmt.Println(c.SayHello(metadata.NewOutgoingContext(context.TODO(), metadata.Pairs("key", "val")), &HelloRequest{Name: "hello"}))
	fmt.Println(c.SayHello(metadata.NewOutgoingContext(context.TODO(), metadata.Pairs("key", "val")), &HelloRequest{Name: "hello"}))
	fmt.Println(c.SayHello(metadata.NewOutgoingContext(context.TODO(), metadata.Pairs("key", "val")), &HelloRequest{Name: "hello"}))
	fmt.Println(c.SayHello(metadata.NewOutgoingContext(context.TODO(), metadata.Pairs("key", "val")), &HelloRequest{Name: "hello"}))
	fmt.Println(c.SayHello(metadata.NewOutgoingContext(context.TODO(), metadata.Pairs("key", "val")), &HelloRequest{Name: "hello"}))
	fmt.Println(c.SayHello(metadata.NewOutgoingContext(context.TODO(), metadata.Pairs("key", "val")), &HelloRequest{Name: "hello"}))
	fmt.Println(c.SayHello(metadata.NewOutgoingContext(context.TODO(), metadata.Pairs("key", "val")), &HelloRequest{Name: "hello"}))
	fmt.Println(c.SayHello(metadata.NewOutgoingContext(context.TODO(), metadata.Pairs("key", "val")), &HelloRequest{Name: "hello"}))
	fmt.Println(c.SayHello(metadata.NewOutgoingContext(context.TODO(), metadata.Pairs("key", "val")), &HelloRequest{Name: "hello"}))
	fmt.Println(c.SayHello(metadata.NewOutgoingContext(context.TODO(), metadata.Pairs("key", "val")), &HelloRequest{Name: "hello"}))
	fmt.Println(c.SayHello(metadata.NewOutgoingContext(context.TODO(), metadata.Pairs("key", "val")), &HelloRequest{Name: "hello"}))
}

func TestClientStatus(t *testing.T) {
	conn, err := grpc.Dial("localhost:80", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := NewGreeterClient(conn)
	resp, err := c.SayHello(metadata.NewOutgoingContext(context.TODO(), metadata.Pairs("key", "val")), &HelloRequest{Name: "hello"})

	s, ok := status.FromError(err)
	fmt.Println(ok)
	fmt.Println("code", s.Code())
	fmt.Println("message", s.Message())
	fmt.Println(s.Details())

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
}

func TestClientWithContext(t *testing.T) {
	conn, err := grpc.DialContext(context.TODO(), "custom///123456", grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := NewGreeterClient(conn)
	fmt.Println(c.SayHello(metadata.NewOutgoingContext(context.TODO(), metadata.Pairs("key", "val")), &HelloRequest{Name: "hello"}))

}

func TestClientLB(t *testing.T) {
	conn, err := grpc.DialContext(context.TODO(), "custom:///123456", grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := NewGreeterClient(conn)
	for {
		time.Sleep(1 * time.Second)
		fmt.Println(c.SayHello(metadata.NewOutgoingContext(context.TODO(), metadata.Pairs("key", "val")), &HelloRequest{Name: "hello"}))
	}
}
