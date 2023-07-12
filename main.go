package main

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"
	"things-auth-service/authgrpc"

	"github.com/ngaut/log"
	"google.golang.org/grpc"
)

/*
*
* 设备列表
*
 */
var devices map[string]bool = map[string]bool{
	"AABBCC00000001": true,
	"AABBCC00000002": true,
	"AABBCC00000003": true,
	"AABBCC00000004": true,
	"AABBCC00000005": true,
	"AABBCC00000006": true,
	"AABBCC00000007": true,
	"AABBCC00000008": true,
}

type _rpcServer struct {
	authgrpc.UnimplementedAuthenticationServer
}

func (*_rpcServer) CheckAuth(ctx context.Context,
	request *authgrpc.AuthRequest) (response *authgrpc.AuthResponse, err error) {
	response = new(authgrpc.AuthResponse)
	log.Debug("CheckAuth => ", request.String())
	if devices[request.ClientId] {
		response.Result = true
		response.Msg = "AUTH SUCCESS"
		response.IsSuperuser = false
	} else {
		response.Result = false
		response.Msg = "AUTH FAILURE"
		response.IsSuperuser = false
	}
	return response, nil
}
func (*_rpcServer) CheckACL(ctx context.Context,
	request *authgrpc.ACLRequest) (response *authgrpc.ACLResponse, err error) {
	log.Debug("CheckACL => ", request.String())
	response = new(authgrpc.ACLResponse)
	response.Result = true
	response.Msg = "OK"
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
	log.Info("Mqtt Auth Proxy Server started @", listener.Addr())
	rpcServer.Serve(listener)
}

//go:generate ./gen_proto.sh

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGABRT, syscall.SIGTERM)
	go startServer()
	s := <-c
	log.Warn("Received stop signal:", s)
	os.Exit(0)
}
