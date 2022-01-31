package reservation

import "time"

type User struct {
	UserID      string     `xorm:"user_id"`
	Name        string     `json:"name" xorm:"name"`
	PhoneNumber string     `json:"phone_number" xorm:"phone_number"`
	Email       string     `json:"email" xorm:"email"`
	Address     string     `json:"address" xorm:"address"`
	UserType    string     `json:"user_type" xorm:"user_type"`
	PostalCode  string     `json:"postal_code" xorm:"postal_code"`
	CreateTime  *time.Time `xorm:"create_time created"`
	UpdateTime  *time.Time `xorm:"update_time updated"`
	DeleteTime  time.Time  `xorm:"delete_time deleted"`
}
