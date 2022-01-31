package reservation

import "time"

type HstReservation struct {
	HstReservationID int64      `json:"hst_reservation_id" xorm:"hst_reservation_id pk autoincr"`
	ReservationID    int64      `json:"reservation_id" xorm:"reservation_id"`
	CustomerID       int64      `json:"customer_id" xorm:"customer_id"`
	ShopID           int64      `json:"shop_id" xorm:"shop_id"`
	ReservationTime  *time.Time `json:"reservation_time" xorm:"reservation_time"`
	ReservationType  string     `json:"reservation_type" xorm:"reservation_type"`
	StatusID         int64      `json:"status_id" xorm:"status_id"`
	CustomerNote     string     `json:"customer_note" xorm:"customer_note"`
	ShopNote         string     `json:"shop_note" xorm:"shop_note"`
	UpdaterID        int64      `json:"updater_id" xorm:"updater_id"`
	Reason           string     `json:"reason" xorm:"reason"`
	UpdateType       string     `json:"update_type" xorm:"update_type"`
	CreateTime       time.Time  `json:"create_time" xorm:"create_time created"`
	UpdateTime       time.Time  `json:"update_time" xorm:"update_time updated"`
	DeleteTime       *time.Time `json:"delete_time" xorm:"delete_time deleted"`
}
