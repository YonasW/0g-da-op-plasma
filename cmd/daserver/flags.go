package main

import (
	"errors"
	"fmt"

	"github.com/urfave/cli/v2"

	opservice "github.com/ethereum-optimism/optimism/op-service"
	oplog "github.com/ethereum-optimism/optimism/op-service/log"
)

const (
	ListenAddrFlagName  = "addr"
	PortFlagName        = "port"
	GenericCommFlagName = "generic-commitment"
	ZgServerFlagName    = "zg.server"
)

const EnvVarPrefix = "OP_PLASMA_DA_SERVER"

func prefixEnvVars(name string) []string {
	return opservice.PrefixEnvVar(EnvVarPrefix, name)
}

var (
	ListenAddrFlag = &cli.StringFlag{
		Name:    ListenAddrFlagName,
		Usage:   "server listening address",
		Value:   "127.0.0.1",
		EnvVars: prefixEnvVars("ADDR"),
	}
	PortFlag = &cli.IntFlag{
		Name:    PortFlagName,
		Usage:   "server listening port",
		Value:   3100,
		EnvVars: prefixEnvVars("PORT"),
	}
	GenericCommFlag = &cli.BoolFlag{
		Name:    GenericCommFlagName,
		Usage:   "enable generic commitments for testing. Not for production use.",
		EnvVars: prefixEnvVars("GENERIC_COMMITMENT"),
		Value:   false,
	}
	ZgServerFlag = &cli.StringFlag{
		Name:    ZgServerFlagName,
		Usage:   "zg server endpoint",
		Value:   "localhost:51001",
		EnvVars: prefixEnvVars("ZG_SERVER"),
	}
)

var requiredFlags = []cli.Flag{
	ListenAddrFlag,
	PortFlag,
}

var optionalFlags = []cli.Flag{
	GenericCommFlag,
	ZgServerFlag,
}

func init() {
	optionalFlags = append(optionalFlags, oplog.CLIFlags(EnvVarPrefix)...)
	Flags = append(requiredFlags, optionalFlags...)
}

// Flags contains the list of configuration options available to the binary.
var Flags []cli.Flag

type CLIConfig struct {
	UseGenericComm bool
	ZgServer       string
}

func ReadCLIConfig(ctx *cli.Context) CLIConfig {
	return CLIConfig{
		UseGenericComm: ctx.Bool(GenericCommFlagName),
		ZgServer:       ctx.String(ZgServerFlagName),
	}
}

func (c CLIConfig) Check() error {
	if !c.ZgEnabled() {
		return errors.New("at least one storage backend must be enabled")
	}

	return nil
}

func (c CLIConfig) ZgEnabled() bool {
	return !(c.ZgServer == "")
}

func (c CLIConfig) ZgConfig() ZgConfig {
	return ZgConfig{
		server: c.ZgServer,
	}
}

func CheckRequired(ctx *cli.Context) error {
	for _, f := range requiredFlags {
		if !ctx.IsSet(f.Names()[0]) {
			return fmt.Errorf("flag %s is required", f.Names()[0])
		}
	}
	return nil
}
