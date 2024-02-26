package proxy

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type ProxyService struct {
	cl *rpc.Client
}

func (p *ProxyService) Apply(command string) error {
	// Future optimization:
	// leader coordination,
	// batching requests and so on...
	return p.cl.Call("Replica.Apply", command, nil)
}

func Proxy() {

	proxy := new(ProxyService)
	rpc.Register(proxy)
	rpc.HandleHTTP()

	cl, err := rpc.DialHTTP("tcp", ":todoport")

	if err != nil {
		log.Fatal(err)
	}

	proxy.cl = cl

	l, err := net.Listen("tcp", ":todoport")

	if err != nil {
		log.Fatal(err)
	}

	http.Serve(l, nil)
}
