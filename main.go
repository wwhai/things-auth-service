package main

import (
	"context"
	"net"
	"os"
	"os/signal"
	"strings"
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
	"admin":                true,
	"wankeyun001":          true,
	"wankeyun002":          true,
	"ESP8266D1001":         true,
	"Arduino-Eth001":       true,
	"Arduino-Eth002":       true,
	"ESP826612E-MOD001":    true,
	"ESP826612E-MOD002":    true,
	"ESP826612E-MOD003":    true,
	"ESP826612E-MOD004":    true,
	"ESP826612E-MOD005":    true,
	"ESP826612E-LORA-001":  true,
	"ESP826612E-LORA-002":  true,
	"ESP826612E-LORA-003":  true,
	"ESP826612E-LORA-004":  true,
	"ESP826612E-LORA-005":  true,
	"ESP826612E-STC51-001": true,
	"PLC-001":              true,
	"PLC-002":              true,
	"RULEX-001":            true,
	"RULEX-002":            true,
	"RULEX-003":            true,
}

type _rpcServer struct {
	authgrpc.UnimplementedAuthenticationServer
}

func (*_rpcServer) CheckAuth(ctx context.Context, request *authgrpc.AuthRequest) (response *authgrpc.AuthResponse, err error) {
	response = new(authgrpc.AuthResponse)
	log.Debug("CheckAuth => ", request)
	if strings.HasPrefix(request.ClientId, "rulex-dev-") || devices[request.ClientId] {
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
func (*_rpcServer) CheckACL(ctx context.Context, request *authgrpc.ACLRequest) (response *authgrpc.ACLResponse, err error) {
	log.Debug("CheckACL => ", request)
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
	log.Info("Server started @", listener.Addr())
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
