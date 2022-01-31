package reservation

import (
	"time"
)

type TrxReservation struct {
	ReservationID   int64      `json:"reservation_id" xorm:"reservation_id pk autoincr"`
	CustomerID      int64      `json:"customer_id" xorm:"customer_id"`
	OwnerID         int64      `json:"owner_id" xorm:"owner_id"`
	ShopID          int64      `json:"shop_id" xorm:"shop_id"`
	StatusID        int64      `json:"status_id" xorm:"status_id"`
	ReservationTime *time.Time `json:"reservation_time" xorm:"reservation_time"`
	ReservationType string     `json:"reservation_type" xorm:"reservation_type"`
	CustomerNote    string     `json:"customer_note" xorm:"customer_note"`
	ShopNote        string     `json:"shop_note" xorm:"shop_note"`
	CreateTime      time.Time  `json:"-" xorm:"create_time created"`
	UpdateTime      time.Time  `json:"-" xorm:"update_time updated"`
	DeleteTime      *time.Time `json:"-" xorm:"delete_time deleted"`
}

type RsvRegistration struct {
	BusinessID      int64      `json:"business_id" xorm:"reservation_id pk autoincr"`
	ReservationTime *time.Time `json:"reservation_time" xorm:"owner_id"`
	ReservationType string     `json:"reservatioin_type" xorm:"reservatioin_type"`
	ReservationNote string     `json:"reservation_note" xorm:"status_id"`
	CreateTime      *time.Time `json:"-" xorm:"create_time created"`
	DeleteTime      *time.Time `json:"-" xorm:"delete_time deleted"`
	UpdateTime      *time.Time `json:"-" xorm:"update_time updated"`
}
