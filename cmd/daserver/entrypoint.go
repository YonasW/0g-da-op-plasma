package main

import (
	"fmt"

	zg "github.com/0glabs/op-plasma-0g"
	oplog "github.com/ethereum-optimism/optimism/op-service/log"
	"github.com/ethereum-optimism/optimism/op-service/opio"
	"github.com/urfave/cli/v2"
)

func StartDAServer(cliCtx *cli.Context) error {
	if err := CheckRequired(cliCtx); err != nil {
		return err
	}

	cfg := ReadCLIConfig(cliCtx)
	if err := cfg.Check(); err != nil {
		return err
	}

	logCfg := oplog.ReadCLIConfig(cliCtx)

	l := oplog.NewLogger(oplog.AppOut(cliCtx), logCfg)
	oplog.SetGlobalLogHandler(l.Handler())

	l.Info("Initializing Plasma DA server...")

	var store zg.KVStore

	if cfg.ZgEnabled() {
		l.Info("Using zg storage", "url", cfg.ZgServer)
		s, err := NewZgStore(cliCtx.Context, cfg.ZgConfig(), l)
		if err != nil {
			return fmt.Errorf("failed to create zg store: %w", err)
		}
		store = s
	} else {
		return fmt.Errorf("store must be specified")
	}

	server := zg.NewDAServer(cliCtx.String(ListenAddrFlagName), cliCtx.Int(PortFlagName), store, l, cfg.UseGenericComm)

	if err := server.Start(); err != nil {
		return fmt.Errorf("failed to start the DA server")
	} else {
		l.Info("Started DA Server")
	}

	defer func() {
		if err := server.Stop(); err != nil {
			l.Error("failed to stop DA server", "err", err)
		}
	}()

	opio.BlockOnInterrupts()

	return nil
}
