package rateLimited

import "NotificationService-rl/dataBase/postgres"

type Repository struct {
	dbRepo postgres.DbRepositoryBase
}

type RepositoryInterface interface {
	SaveAttempt(data NotificationAttempt) error
	GetAttempt(emailRecipient, typeId string) (int, string)
	GetRules(notificationType string) (int, string)
}

type TypeNotification struct {
	TypeId   string `gorm:"column:type_id" json:"type_id"`
	TypeName string `gorm:"column:type_name" json:"type_name"`
	RlRuleId int    `gorm:"column:rl_rule_id" json:"rl_rule_id"`
}

type RateLimitRule struct {
	RlRuleId     int    `gorm:"column:rl_rule_id" json:"rl_rule_id"`
	MaxShipments int    `gorm:"column:max_shipments" json:"max_shipments"`
	TimeShipment string `gorm:"column:time_shipment" json:"time_shipment"`
}

type NotificationAttempt struct {
	EmailRecipient   string `gorm:"primary_key;column:email_recipient" json:"email_recipient"`
	TypeId           string `gorm:"primary_key;column:type_id" json:"type_id"`
	ShipmentCount    int    `gorm:"column:shipment_count" json:"shipment_count"`
	LastNotification string `gorm:"column:last_notification" json:"last_notification"`
}

func (TypeNotification) TableName() string {
	return "type_notification"
}

func (RateLimitRule) TableName() string {
	return "rate_limit_rule"
}

func (NotificationAttempt) TableName() string {
	return "notification_attempt"
}

func NewRepository(dbRepo postgres.DbRepositoryBase) RepositoryInterface {
	return &Repository{dbRepo: dbRepo}
}

func (r Repository) SaveAttempt(data NotificationAttempt) error {
	return r.dbRepo.Save(&data)
}

func (r Repository) GetAttempt(emailRecipient, typeId string) (int, string) {
	condition := "email_recipient = ? AND type_id = ?"
	parameters := []interface{}{emailRecipient, typeId}
	notificationAttempt := new([]NotificationAttempt)
	r.dbRepo.GetByConditions(notificationAttempt, condition, parameters)
	if len(*notificationAttempt) == 0 {
		return 0, ""
	}
	shipmentCount := (*notificationAttempt)[0].ShipmentCount
	lasNotification := (*notificationAttempt)[0].LastNotification
	return shipmentCount, lasNotification
}

func (r Repository) GetRules(notificationType string) (int, string) {
	condition := "type_id = ?"
	parameters := []interface{}{notificationType}
	RLRule := new([]RateLimitRule)
	r.dbRepo.GetByConditions(RLRule, condition, parameters)
	if len(*RLRule) == 0 {
		return 0, ""
	}
	return (*RLRule)[0].MaxShipments, (*RLRule)[0].TimeShipment
}
