package handlers

import (
	"fmt"
)

func HandleMessage(messageEvent interface{}) {
	fmt.Println("Received message", messageEvent)
}
