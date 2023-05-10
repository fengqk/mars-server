package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/fengqk/mars-base/base"
	"github.com/fengqk/mars-server/gate"
)

func Init() {
	common.Init()
}

func Exit() {

}

func main() {
	args := os.Args
	base.LOG.Init(args[1])
	base.SEVERNAME = args[1]

	switch base.SEVERNAME {
	case "gate":
		gate.SERVER.Init()
	}

	Init()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	s := <-c

	fmt.Printf("server %s exit,signal: %v", args[1], s)

	Exit()
}
