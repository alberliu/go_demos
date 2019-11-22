package rpc

import (
	"fmt"
	_ "google.golang.org/grpc/balancer/grpclb"
	"google.golang.org/grpc/resolver"
)

func init() {
	fmt.Println("init")
	RegisterResolver()
}

func RegisterResolver() {
	resolver.Register(NewCustomBuilder())
}

type CustomBuilder struct {
}

func NewCustomBuilder() resolver.Builder {
	return &CustomBuilder{}
}

func (b *CustomBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOption) (resolver.Resolver, error) {
	cr := &CustomResolver{
		cc: cc,
	}
	fmt.Printf("%#v\n", target)

	state := resolver.State{
		Addresses: []resolver.Address{
			{
				Addr: "127.0.0.1:50000",
				Type: resolver.Backend,
			},
			{
				Addr: "127.0.0.1:50001",
				Type: resolver.Backend,
			},
		},
	}
	cc.UpdateState(state)
	return cr, nil
}

func (b *CustomBuilder) Scheme() string {
	return "custom"
}

type CustomResolver struct {
	cc resolver.ClientConn
}

func (r *CustomResolver) ResolveNow(opt resolver.ResolveNowOption) {
	fmt.Println("ResolveNow")

}

func (r *CustomResolver) Close() {
	fmt.Println("Close")
}
