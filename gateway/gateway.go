package gateway

import "fmt"

type Gateway struct {
	// already implemented
}

func (g *Gateway) Send(userId, message string) {
	printMessage := fmt.Sprintf("sending message to user: %s, via Gateway. Message: %s", userId, message)
	fmt.Println(printMessage)
}
