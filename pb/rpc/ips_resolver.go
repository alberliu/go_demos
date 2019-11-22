package rpc

import (
	"fmt"
	_ "google.golang.org/grpc/balancer/grpclb"
	"google.golang.org/grpc/resolver"
	"strings"
)

func init() {
	fmt.Println("init")
	RegisterResolver()
}

func RegisterResolver() {
	resolver.Register(NewCustomBuilder())
}

type IPsBuilder struct {
}

func NewCustomBuilder() resolver.Builder {
	return &IPsBuilder{}
}

func (b *IPsBuilder) Build(target resolver.Target, clientConn resolver.ClientConn, opts resolver.BuildOption) (resolver.Resolver, error) {
	ips := strings.Split(target.Endpoint, ",")

	state := resolver.State{
		Addresses: getAddresses(ips),
	}
	clientConn.UpdateState(state)
	return &IPsResolver{
		ips:        ips,
		clientConn: clientConn,
	}, nil
}

func (b *IPsBuilder) Scheme() string {
	return "ips"
}

func getAddresses(ips []string) []resolver.Address {
	addresses := make([]resolver.Address, len(ips))
	for i := range ips {
		addresses[i].Addr = ips[i]
		addresses[i].Type = resolver.Backend
	}
	return addresses
}

type IPsResolver struct {
	ips        []string
	clientConn resolver.ClientConn
}

func (r *IPsResolver) ResolveNow(opt resolver.ResolveNowOption) {
	state := resolver.State{
		Addresses: getAddresses(r.ips),
	}
	r.clientConn.UpdateState(state)
}

func (r *IPsResolver) Close() {
	fmt.Println("Close")
}
