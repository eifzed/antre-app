package reservation

import (
	"time"
)

type TrxReservation struct {
	ReservationID int64      `json:"reservation_id" xorm:"reservation_id pk autoincr"`
	OwnerID       int64      `json:"owner_id" xorm:"owner_id"`
	CustomerID    int64      `json:"customer_id" xorm:"customer_id"`
	StatusID      int64      `json:"status_id" xorm:"status_id"`
	CreateTime    *time.Time `json:"-" xorm:"create_time created"`
	DeleteTime    *time.Time `json:"-" xorm:"delete_time deleted"`
	UpdateTime    *time.Time `json:"-" xorm:"update_time updated"`
}
