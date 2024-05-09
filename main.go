package main

import (
	"rateLimited"
)

func main() {
	service := rateLimited.NewNotificationServiceImpl(Gateway{})
	service.Send("news", "user", "news 1")
	service.Send("news", "user", "news 2")
	service.Send("news", "user", "news 3")
	service.Send("news", "another user", "news 1")
	service.Send("update", "user", "update 1")
}