// @generated Code generated by thrift-gen. Do not modify.

// Package ping is generated code used to make or handle TChannel calls using Thrift.
package ping

import (
	"fmt"

	athrift "github.com/apache/thrift/lib/go/thrift"
	"github.com/temporalio/tchannel-go/thrift"
)

// Interfaces for the service and client for the services defined in the IDL.

// TChanPingPongService is the interface that defines the server handler and client interface.
type TChanPingPongService interface {
	Ping(ctx thrift.Context, request *Ping) (*Pong, error)
}

// Implementation of a client and service handler.

type tchanPingPongServiceClient struct {
	thriftService string
	client        thrift.TChanClient
}

func NewTChanPingPongServiceInheritedClient(thriftService string, client thrift.TChanClient) *tchanPingPongServiceClient {
	return &tchanPingPongServiceClient{
		thriftService,
		client,
	}
}

// NewTChanPingPongServiceClient creates a client that can be used to make remote calls.
func NewTChanPingPongServiceClient(client thrift.TChanClient) TChanPingPongService {
	return NewTChanPingPongServiceInheritedClient("PingPongService", client)
}

func (c *tchanPingPongServiceClient) Ping(ctx thrift.Context, request *Ping) (*Pong, error) {
	var resp PingPongServicePingResult
	args := PingPongServicePingArgs{
		Request: request,
	}
	success, err := c.client.Call(ctx, c.thriftService, "Ping", &args, &resp)
	if err == nil && !success {
		switch {
		default:
			err = fmt.Errorf("received no result or unknown exception for Ping")
		}
	}

	return resp.GetSuccess(), err
}

type tchanPingPongServiceServer struct {
	handler TChanPingPongService
}

// NewTChanPingPongServiceServer wraps a handler for TChanPingPongService so it can be
// registered with a thrift.Server.
func NewTChanPingPongServiceServer(handler TChanPingPongService) thrift.TChanServer {
	return &tchanPingPongServiceServer{
		handler,
	}
}

func (s *tchanPingPongServiceServer) Service() string {
	return "PingPongService"
}

func (s *tchanPingPongServiceServer) Methods() []string {
	return []string{
		"Ping",
	}
}

func (s *tchanPingPongServiceServer) Handle(ctx thrift.Context, methodName string, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	switch methodName {
	case "Ping":
		return s.handlePing(ctx, protocol)

	default:
		return false, nil, fmt.Errorf("method %v not found in service %v", methodName, s.Service())
	}
}

func (s *tchanPingPongServiceServer) handlePing(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req PingPongServicePingArgs
	var res PingPongServicePingResult

	if err := req.Read(ctx, protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.Ping(ctx, req.Request)

	if err != nil {
		return false, nil, err
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}
