package controller

import (
	"log"
)

func KeyHandler(key, state int) {
	log.Printf("key: %v, state: %v", key, state)
	if key == KeyA {
		log.Println("a")
	}
}
