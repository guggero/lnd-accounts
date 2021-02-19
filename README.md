# This project has been archived!

A new approach with a different architecture can be found here:
 - [RPC message interceptor for lnd (replaces the component in aperture)](https://github.com/guggero/lnd/tree/macaroon-interceptor)
 - [Account system in faraday (replaces this repository)](https://github.com/guggero/faraday/tree/lnd-accounts)


# Accounts plugin for lnd

This is the continuation of the work [previously started as a PR in lnd](https://github.com/lightningnetwork/lnd/pull/2390).

The goal is to add the concept of "accounts" to `lnd`. An account is simply a
virtual balance of some satoshis in an `lnd` node. With macaroons as a
capability based authentication method, this means that users can be given
restricted macaroons that only allows them to access one account in `lnd`.

This would allow the following features:
 - Make an `lnd` node multi-tenant capable. Create multiple accounts, let other
   users (family members, friends) connect to your node by giving them
   restricted macaroons.
 - When using an account-restricted macaroon, the user would only see the
   balance of their account and only invoices/payments that were paid/received
   with that account.
 - Invoices created by an account will be credited to that account once they
   are paid.
 - Concepts not important to "managed" users like channels and peers can be
   hidden (=empty response) from them.

So the goal is pretty simple: **Give your friends and family members access to
the `lnd` node you manage by simply letting them scan a QR code**.

## Plugin? Lnd doesn't have plugins!

Yes, that's correct. Unlike `c-lightning` that fully supports plugins, `lnd`
itself doesn't allow developers to directly run code inside the core `lnd`
process. Instead, `lnd` gives developers a wide range of gRPC APIs that can be
used to interact with a node.

With [aperture](https://github.com/lightninglabs/aperture) Lightning Labs
released a gRPC and HTTP/2 reverse proxy that can (among other things) be put
in front of an `lnd` node and encapsulate/intercept all gRPC calls made to it.

So if we can't put plugins into `lnd`, what's the next best thing to put plugins
in? Right, a reverse proxy that can be placed in front of `lnd` and has the
ability to offer the following things to plugin developers:
 - Intercept _any_ gRPC call made to `lnd` and have the ability to change both
   the request and response message.
 - Give the plugin a direct connection to `lnd` so it can create their own
   subscriptions or interact with `lnd` independently of the user.
 - Give the plugin the ability to add their own gRPC methods to the interface,
   extending the functionality of the `lnd` node.
 - Experiment with less critical code outside of the main `lnd` process. If the
   plugin crashes, no harm comes to the node.

## How can I write my own Aperture plugin?

Right now, this is barely a proof of concept. A lot of work still has to go into
this and nothing is finished yet.

But if you want to start writing plugins, you can do so with my fork of Aperture:
https://github.com/guggero/aperture/tree/lnd-plugin

A bare minimum plugin would look like this:

```go
package myplugin

import (
	"fmt"
	"path"

	"github.com/lightninglabs/aperture/plugin/shared"
	"github.com/lightninglabs/aperture/plugin/author"
	"github.com/lightninglabs/aperture/plugin/pluginrpc"
	"github.com/lightninglabs/loop/lndclient"
	"github.com/lightningnetwork/lnd/lnrpc"
	"google.golang.org/grpc"
)

const (
	PluginName = "example-plugin"
)

func main() {
	p := &LndExamplePlugin{}
	author.RunPluginServer(PluginName, p, &shared.RpcPackage{
		Name:     "examplerpc",
		Services: []interface{}{
                // Add definitions for your plugin's RPC package here.
                (*examplerpc.ExampleServer)(nil),
        },
	})
}

type LndExamplePlugin struct {
	*ExampleRpcServer

	lnd lnrpc.LightningClient
}

func (e *LndExamplePlugin) Init(request *pluginrpc.StartRequest) error {
	var err error
	e.ExampleRpcServer = &ExampleRpcServer{}
	e.lnd, err = lndclient.NewBasicClient(
		request.LndHost, request.LndCert, path.Dir(request.LndMacaroon),
		"", lndclient.MacFilename(path.Base(request.LndMacaroon)),
	)
	fmt.Printf("Successfully connected to lnd at %s", request.LndHost)
	return err
}

func (e *LndExamplePlugin) RegisterServices(s *grpc.Server) {
	examplerpc.RegisterExampleServer(s, e)
}

func (e *LndExamplePlugin) Stop() error {
	return nil
}

// HookLightningNewAddressReq is an example of how an incoming gRPC request
// message for a call can be intercepted and the message manipulated.
// The naming pattern for a request interception hook is:
//    Hook<RPCServiceName><RPCMethodName>Req
func (e *LndExamplePlugin) HookLightningNewAddressReq(req *lnrpc.NewAddressRequest) error {
	fmt.Printf("Intercepted request with type: %v\n", req)
	req.Type = lnrpc.AddressType_NESTED_PUBKEY_HASH
	return nil
}

// HookLightningNewAddressRes is an example of how an outgoing gRCP response
// message for a call can be intercepted and the message manipulated.
// The naming pattern for a response interception hook is:
//    Hook<RPCServiceName><RPCMethodName>Res
func (e *LndExamplePlugin) HookLightningNewAddressRes(res *lnrpc.NewAddressResponse) error {
	fmt.Printf("Intercepted response with new address %s\n", res.Address)
	res.Address = "fooooo"
	return nil
}
```

You could then compile this plugin to a binary called `plugin-example-plugin`
and place it into the directory where Aperture will pick it up, for example
`~/.aperture/plugins` with the following configuration:

```yaml
listenaddr: "localhost:11110"
staticroot: "."
debuglevel: "trace"

plugins:
  dir: /home/user/.aperture/plugins
  lndhost: 127.0.0.1:10011
  lndcert: /home/user/regtest/alice/tls.cert
  lndmac: /home/user/regtest/alice/admin.macaroon
  enabled:
    - lnd-accounts

services:
  - name: alice-rpc
    hostregexp: '^localhost:?\d*$'
    pathregexp: '^/(lnrpc|signrpc|walletrpc|examplerpc).*/.*$'
    address: "127.0.0.1:10011"
    protocol: https
    tlscertpath: "/home/user/regtest/alice/tls.cert"
    auth: off
    rpcservicetype: lnrpc
```
