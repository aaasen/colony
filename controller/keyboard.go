package controller

import (
	"log"

	"github.com/jteeuwen/glfw"
)

const (
	_ = iota
	SayHey
)

var defaultKeyBindings = map[KeyEvent]int{
	KeyEvent{KeyA, glfw.KeyPress}: SayHey,
}

type KeyEvent struct {
	Key, State int
}

var keyEventChannel = make(chan KeyEvent)

func KeyHandler(key, state int) {
	log.Printf("key: %v, state: %v", key, state)
	keyEventChannel <- KeyEvent{key, state}
}
