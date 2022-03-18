package main

import (
	"context"
	"net"
	"things-auth-service/authgrpc"

	"github.com/ngaut/log"
	"google.golang.org/grpc"
)

type _rpcServer struct {
	authgrpc.UnimplementedAuthenticationServer
}

func (*_rpcServer) CheckAuth(ctx context.Context, request *authgrpc.AuthRequest) (response *authgrpc.AuthResponse, err error) {
	response = new(authgrpc.AuthResponse)
	response.Result = true
	response.Msg = "OK"
	response.IsSuperuser = false
	log.Debug("CheckAuth========> ", request)
	return response, nil
}
func (*_rpcServer) CheckACL(ctx context.Context, request *authgrpc.ACLRequest) (response *authgrpc.ACLResponse, err error) {
	response = new(authgrpc.ACLResponse)
	response.Result = false
	response.Msg = "OK"
	log.Debug("CheckACL========> ", request)
	return response, nil
}
func startServer() {
	listener, err := net.Listen("tcp", ":1994")
	if err != nil {
		log.Fatal(err)
		return
	}
	rpcServer := grpc.NewServer()
	authgrpc.RegisterAuthenticationServer(rpcServer, new(_rpcServer))
	go func(c context.Context) {
		log.Info("rpcCodecServer started on", listener.Addr())
		rpcServer.Serve(listener)
	}(context.TODO())

}

//go:generate ./gen_proto.sh

func main() {
	startServer()

	for {
	}
}
