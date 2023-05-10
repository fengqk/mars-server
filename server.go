package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/fengqk/mars-base/base"
	"github.com/fengqk/mars-server/db"
	"github.com/fengqk/mars-server/game"
	"github.com/fengqk/mars-server/gate"
	"github.com/fengqk/mars-server/message"
)

func main() {
	args := os.Args
	base.LOG.Init(args[1])
	base.SEVERNAME = args[1]

	switch base.SEVERNAME {
	case "gate":
		gate.SERVER.Init()
	case "game":
		game.SERVER.Init()
	case "db":
		db.SERVER.Init()
	}

	Init()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	s := <-c

	fmt.Printf("server %s exit,signal: %v", args[1], s)

	Exit()
}

func Init() {
	message.Init()
}

func Exit() {
}
