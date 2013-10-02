package main

import (
	"github.com/aaasen/colony/model"
	"github.com/aaasen/colony/view"

	"time"
)

func main() {

	rendererChan := make(chan view.Renderer)
	controlChan := make(chan bool)

	model := model.NewModel(rendererChan, controlChan)
	go model.Listen()

	view := view.NewView(rendererChan)
	view.Init()
	defer view.Destroy()
	go view.Listen()

	time.Sleep(time.Second * 10)

}
