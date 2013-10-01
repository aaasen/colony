package controller

import (
	"log"

	"github.com/jteeuwen/glfw"
)

func MouseButtonHandler(button, state int) {
	if button == glfw.Mouse1 && state == glfw.KeyPress {
		log.Println("LMB pressed")
	}

	log.Printf("button: %v, state: %v", button, state)
}
