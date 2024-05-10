package rateLimited

import (
	"NotificationService-rl/gateway"
	"fmt"
	"time"
)

type NotificationService struct {
	repo    RepositoryInterface
	gateway gateway.Gateway
}

type NotificationServiceInterface interface {
	Send(typeStr, userEmail, message string)
}

func NewNotificationServiceImpl(gateway gateway.Gateway, repo RepositoryInterface) NotificationServiceInterface {
	return NotificationService{
		repo:    repo,
		gateway: gateway,
	}
}

func (n NotificationService) Send(typeStr, userEmail, message string) {
	actualTime := time.Now()
	attempt, lastNotification := n.repo.GetAttempt(userEmail, typeStr)
	maxShipments, timeRule := n.repo.GetRules(typeStr)
	if maxShipments == 0 {
		fmt.Println("no rules found for type: ", typeStr)
		return
	}
	lastNotificationTime, errTime := time.Parse(time.RFC3339, lastNotification)
	if errTime != nil {
		lastNotificationTime = actualTime
	}
	if attempt != 0 || !n.ApplyValidations(attempt, maxShipments, timeRule, actualTime, lastNotificationTime) {
		errMessage := fmt.Sprintf("Rule -> type: %s, max shipments: %v, time shipment: %s", typeStr, maxShipments, timeRule)
		fmt.Println(errMessage)
		fmt.Println("error validating last notification. Attempt: ", attempt, " Last Notification: ", lastNotification)
		n.repo.SaveAttempt(NotificationAttempt{
			EmailRecipient:   userEmail,
			TypeId:           typeStr,
			ShipmentCount:    0,
			LastNotification: time.Time{}.String(),
		})
		return
	}
	n.gateway.Send(userEmail, message)
	fmt.Println("sending message to user", userEmail)
	n.repo.SaveAttempt(NotificationAttempt{
		EmailRecipient:   userEmail,
		TypeId:           typeStr,
		ShipmentCount:    attempt + 1,
		LastNotification: actualTime.String(),
	})
	return
}

func (n NotificationService) ApplyValidations(attempt, maxShipments int, timeRule string, actualTime, lastNotificationTime time.Time) bool {
	timeDiff := actualTime.Sub(lastNotificationTime).Minutes()
	maxShipmentsState := attempt <= maxShipments
	switch timeRule {
	case "m":
		return timeDiff <= 1 || maxShipmentsState //minutes
	case "h":
		return timeDiff <= 60 || maxShipmentsState //hours
	case "d":
		return timeDiff <= 24*60 || maxShipmentsState //days
	case "W":
		return timeDiff <= 5*24*60 || maxShipmentsState //weeks
	case "M":
		return timeDiff <= 30*24*60 || maxShipmentsState //months
	case "Y":
		return timeDiff <= 365*24*60 || maxShipmentsState //years
	default:
		return false
	}
}
