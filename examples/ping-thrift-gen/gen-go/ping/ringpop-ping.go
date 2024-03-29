// @generated Code generated by thrift-gen. Do not modify.

package ping

import (
	"errors"
	"fmt"

	"github.com/temporalio/ringpop-go"
	"github.com/temporalio/ringpop-go/forward"
	"github.com/temporalio/ringpop-go/router"
	"github.com/temporalio/tchannel-go"
	"github.com/temporalio/tchannel-go/thrift"
)

type RingpopPingPongServiceAdapter struct {
	impl    TChanPingPongService
	ringpop ringpop.Interface
	ch      *tchannel.Channel
	config  PingPongServiceConfiguration
	router  router.Router
}

// PingPongServiceConfiguration contains the forwarding configuration for the PingPongService service. It has a field for every endpoint defined in the service. In this field the endpoint specific forward configuration can be stored. Populating these fields is optional, default behaviour is to call the service implementation locally to the process where the call came in.
type PingPongServiceConfiguration struct {
	// Ping holds the forwarding configuration for the Ping endpoint defined in the service
	Ping *PingPongServicePingConfiguration
}

func (c *PingPongServiceConfiguration) validate() error {
	if c.Ping != nil {
		if c.Ping.Key == nil {
			return errors.New("configuration for endpoint Ping is missing a Key function")
		}
	}
	return nil
}

// NewRingpopPingPongServiceAdapter creates an implementation of the TChanPingPongService interface. This specific implementation will use to configuration provided during construction to deterministically route calls to nodes from a ringpop cluster. The channel should be the channel on which the service exposes its endpoints. Forwarded calls, calls to unconfigured endpoints and calls that already were executed on the right machine will be passed on the the implementation passed in during construction.
//
// Example usage:
//  import "github.com/temporalio/tchannel-go/thrift"
//
//  var server thrift.Server
//  server = ...
//
//  var handler TChanPingPongService
//  handler = &YourImplementation{}
//
//  adapter, _ := NewRingpopPingPongServiceAdapter(handler, ringpop, channel,
//    PingPongServiceConfiguration{
//      Ping: &PingPongServicePingConfiguration: {
//        Key: func(ctx thrift.Context, request *Ping) (shardKey string, err error) {
//          return "calculated-shard-key", nil
//        },
//      },
//    },
//  )
//  server.Register(NewTChanPingPongServiceServer(adapter))
func NewRingpopPingPongServiceAdapter(
	impl TChanPingPongService,
	rp ringpop.Interface,
	ch *tchannel.Channel,
	config PingPongServiceConfiguration,
) (TChanPingPongService, error) {
	err := config.validate()
	if err != nil {
		return nil, err
	}

	adapter := &RingpopPingPongServiceAdapter{
		impl:    impl,
		ringpop: rp,
		ch:      ch,
		config:  config,
	}
	// create ringpop router for routing based on ring membership
	adapter.router = router.New(rp, adapter, ch)

	return adapter, nil
}

// GetLocalClient satisfies the ClientFactory interface of ringpop-go/router
func (a *RingpopPingPongServiceAdapter) GetLocalClient() interface{} {
	return a.impl
}

// MakeRemoteClient satisfies the ClientFactory interface of ringpop-go/router
func (a *RingpopPingPongServiceAdapter) MakeRemoteClient(client thrift.TChanClient) interface{} {
	return NewTChanPingPongServiceClient(client)
}

// PingPongServicePingConfiguration contains the configuration on how to route calls to the thrift endpoint PingPongService::Ping.
type PingPongServicePingConfiguration struct {
	// Key is a closure that generates a routable key based on the parameters of the incomming request.
	Key func(ctx thrift.Context, request *Ping) (string, error)
}

// Ping satisfies the TChanPingPongService interface. This function uses the configuration for Ping to determine the host to execute the call on. When it decides the call needs to be executed in the current process it will forward the invocation to its local implementation.
func (a *RingpopPingPongServiceAdapter) Ping(ctx thrift.Context, request *Ping) (r *Pong, err error) {
	// check if the function should be called locally
	if a.config.Ping == nil || forward.DeleteForwardedHeader(ctx) {
		return a.impl.Ping(ctx, request)
	}

	// find the key to shard on
	ringpopKey, err := a.config.Ping.Key(ctx, request)
	if err != nil {
		return r, fmt.Errorf("could not get key: %q", err)
	}

	clientInterface, isRemote, err := a.router.GetClient(ringpopKey)
	if err != nil {
		return r, err
	}

	client := clientInterface.(TChanPingPongService)
	if isRemote {
		ctx = forward.SetForwardedHeader(ctx, []string{ringpopKey})
	}
	return client.Ping(ctx, request)
}
