package main

import (
	"log"

	"github.com/jteeuwen/glfw"
)

func initGLFW() {
	if err := glfw.Init(); err != nil {
		log.Fatalf("glfw: %s\n", err)
		return
	}

	glfw.OpenWindowHint(glfw.WindowNoResize, 1)

	if err := glfw.OpenWindow(Width, Height, 0, 0, 0, 0, 16, 0, glfw.Windowed); err != nil {
		log.Fatalf("glfw: %s\n", err)
		return
	}

	glfw.SetSwapInterval(1)
	glfw.SetWindowTitle(Title)
}

func terminateGLFW() {
	glfw.Terminate()
	glfw.CloseWindow()
}
