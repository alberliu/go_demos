package http

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
	"testing"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Add(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func TestServer(t *testing.T) {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()

	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestClient(t *testing.T) {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Println("dialing:", err)
		return
	}

	args := Args{17, 8}
	var reply int
	err = client.Call("Arith.Add", args, &reply)
	if err != nil {
		log.Println("arith error:", err)
		return
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	err = client.Call("Arith.Add", args, &reply)
	if err != nil {
		log.Println("arith error:", err)
		return
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)
}
