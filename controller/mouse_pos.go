package controller

import (
	"log"
)

func MousePosHandler(x, y int) {
	log.Printf("(%v, %v)", x, y)
}
