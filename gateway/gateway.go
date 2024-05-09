package gateway

import "fmt"

type Gateway struct {
	// already implemented
}

func (g *Gateway) Send(userId, _ string) {
	fmt.Println("sending message to user", userId)
}
