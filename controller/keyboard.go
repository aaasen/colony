package controller

import (
	"log"

	"github.com/jteeuwen/glfw"
)

func KeyHandler(key, state int) {
	log.Printf("key: %v, state: %v", key, state)
	if key == KeyA && state == glfw.KeyPress {
		log.Println("A pressed")
	}
}
