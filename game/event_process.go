package game

import "github.com/fengqk/mars-base/actor"

type (
	IEventProcess interface {
		actor.IActor
	}

	EventProcess struct {
		actor.Actor
	}
)

func (e *EventProcess) Init() {
	e.Actor.Init()
	actor.MGR.RegisterActor(e)
	e.Actor.Start()
}
