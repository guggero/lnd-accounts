module github.com/guggero/lnd-accounts

go 1.13

require (
	github.com/btcsuite/btclog v0.0.0-20170628155309-84c8d2346e9f
	github.com/btcsuite/btcutil v1.0.2
	github.com/coreos/bbolt v1.3.3
	github.com/golang/protobuf v1.3.4
	github.com/lightninglabs/aperture v0.0.0-20200521002041-64ea3fbcaca1
	github.com/lightninglabs/loop v0.6.2-beta.0.20200521180902-bd2d39bd104b
	github.com/lightningnetwork/lnd v0.10.1-beta.rc1
	github.com/prometheus/common v0.4.0
	github.com/urfave/cli v1.20.0
	google.golang.org/genproto v0.0.0-20190927181202-20e1ac93f88c
	google.golang.org/grpc v1.27.1
	gopkg.in/macaroon-bakery.v2 v2.0.1
	gopkg.in/macaroon.v2 v2.1.0
)

replace github.com/lightninglabs/aperture => ../../lightninglabs/aperture

replace github.com/lightningnetwork/lnd => github.com/lightningnetwork/lnd v0.10.0-beta.rc6.0.20200522002150-c2e7ca9b5f3e
