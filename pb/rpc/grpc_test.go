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

const port = ":8080"

func (s *server) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	return &HelloReply{Message: "0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"}, nil
}
func (s *server) ListFeatures(in *HelloRequest, stream Greeter_ListFeaturesServer) error {
	/*for  {
		time.Sleep(time.Second)
		err:=stream.Send(&HelloReply{
			Message: in.Name,
		})
		fmt.Println(err)
	}*/
	time.Sleep(time.Hour)

	return nil
}

func TestServer(t *testing.T) {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	RegisterGreeterServer(s, &server{})
	log.Println("rpc服务已经开启")
	s.Serve(lis)
}

func TestClientNginxProxy(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:80", grpc.WithInsecure())
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
	conn, err := grpc.DialContext(context.TODO(), "ips:///127.0.0.1:50000,127.0.0.1:50001,127.0.0.1:50002,", grpc.WithInsecure(),
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

func TestClient(t *testing.T) {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := NewGreeterClient(conn)
	resp, err := c.ListFeatures(context.TODO(), &HelloRequest{
		Name: "hello",
	})

	for {
		msg, err := resp.Recv()
		fmt.Println(msg, err)
	}
}

var data = "0123456789012345678901234567890123456789"

func BenchmarkClient(b *testing.B) {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	c := NewGreeterClient(conn)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	for i := 0; i < b.N; i++ {
		c.SayHello(context.TODO(), &HelloRequest{Name: "0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"})
	}
}

func TestData(t *testing.T) {
	for i := 0; i < 1024; i++ {
		fmt.Printf("0")
	}
	fmt.Println()
}

// docker run -it -d -p 8889:8888 -p 8091:8090 -v /usr/local/app/tars/tarsnode2/:/usr/local/app/tars/tarsnode/ -v /usr/local/app/tars/app_log/tars/tarsnode2/:/usr/local/app/tars/app_log/tars/tarsnode/ -v /usr/local/app/tars/app_log/goim/comet2:/usr/local/app/tars/app_log/goim/comet centos:7.9.2009 /bin/bash -c 'sh /usr/local/app/tars/tarsnode/util/execute.sh'
// docker run -it -p 8889:8888 -p 8091:8090 -v /usr/local/app/tars/tarsnode2/:/usr/local/app/tars/tarsnode/ -v /usr/local/app/tars/app_log/tars/tarsnode2/:/usr/local/app/tars/app_log/tars/tarsnode/ -v /usr/local/app/tars/app_log/goim/comet2:/usr/local/app/tars/app_log/goim/comet centos:7.9.2009 /bin/bash /usr/local/app/tars/tarsnode/util/execute.sh


exec user process caused: exec format error