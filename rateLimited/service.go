package rateLimited

import (
	"fmt"
)

type NotificationService interface {
	Send(string, string, string)
}

type NotificationServiceImpl struct {
	gateway Gateway
}

func NewNotificationServiceImpl(gateway Gateway) *NotificationServiceImpl {
	return &NotificationServiceImpl{
		gateway: gateway,
	}
}

func (n *NotificationServiceImpl) Send(typeStr, userId, message string) {
	fmt.Println("sending message to user", userId)
}

func ApplyValidations() {

}
