package main

import (
	"math/rand"
	"time"

	"github.com/aaasen/colony/controller"
	"github.com/aaasen/colony/model"
	"github.com/aaasen/colony/view"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	rendererChan := make(chan view.Renderer)
	controlChan := make(chan bool)
	modelKeyChan := make(chan int)

	model := model.NewModel(rendererChan, modelKeyChan, controlChan)
	go model.Listen()

	view := view.NewView(rendererChan, controlChan)
	view.Init()
	defer view.Destroy()
	go view.Listen()

	controller := controller.NewKeyListener()
	controller.Subscribe(modelKeyChan)
	go controller.Listen()

	time.Sleep(time.Second * 100)

}
