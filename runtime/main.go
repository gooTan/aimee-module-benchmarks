package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/JBailes/aimee/server-go/bus"
	handler "github.com/JBailes/aimee/server-go/modules/benchmarks"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s DAEMON_MODULE_BUS_SOCKET\n", os.Args[0])
		os.Exit(2)
	}
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	config := bus.ModuleProcessConfig{
		SocketPath: os.Args[1], ModuleName: "benchmarks",
		PrincipalClass: 1, PrincipalRef: 25,
		Stages: []bus.ModuleStage{
		{EventKind: 10497, StageID: 1},
		{EventKind: 10498, StageID: 2},
		},
		Handler: handler.Handle,
	}
	if err := bus.RunModuleProcess(ctx, config); err != nil {
		fmt.Fprintf(os.Stderr, "aimee-module-benchmarks: %v\n", err)
		os.Exit(1)
	}
}
