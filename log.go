package main

import (
	"github.com/btcsuite/btclog"
	"github.com/lightninglabs/aperture/plugin"
	"github.com/lightninglabs/aperture/plugin/author"
	"github.com/lightninglabs/aperture/plugin/shared"
	"github.com/lightningnetwork/lnd/build"
)

const Subsystem = "ACCT"

var (
	logWriter = build.NewRotatingLogWriter()

	log = build.NewSubLogger(Subsystem, logWriter.GenSubLogger)
)

// The default amount of logging is none.
func init() {
	setSubLogger("MAIN", log, nil)
	addSubLogger(plugin.Subsystem, plugin.UseLogger)
	addSubLogger(author.Subsystem, author.UseLogger)
	addSubLogger(shared.Subsystem, shared.UseLogger)
}

// addSubLogger is a helper method to conveniently create and register the
// logger of a sub system.
func addSubLogger(subsystem string, useLogger func(btclog.Logger)) {
	logger := build.NewSubLogger(subsystem, logWriter.GenSubLogger)
	setSubLogger(subsystem, logger, useLogger)
}

// setSubLogger is a helper method to conveniently register the logger of a sub
// system.
func setSubLogger(subsystem string, logger btclog.Logger,
	useLogger func(btclog.Logger)) {

	logWriter.RegisterSubLogger(subsystem, logger)
	if useLogger != nil {
		useLogger(logger)
	}
}
