package reservation

import (
	"time"
)

type TrxReservation struct {
	ReservationID int64      `json:"reservation_id" xorm:"reservation_id pk autoincr"`
	CustomerID    int64      `json:"customer_id" xorm:"customer_id"`
	ShopID        int64      `json:"shop_id" xorm:"shop_id"`
	StatusID      int64      `json:"status_id" xorm:"status_id"`
	CustomerNote  string     `json:"customer_note" xorm:"customer_note"`
	ShopNote      string     `json:"shop_note" xorm:"shop_note"`
	CreateTime    time.Time  `json:"-" xorm:"create_time created"`
	UpdateTime    time.Time  `json:"-" xorm:"update_time updated"`
	DeleteTime    *time.Time `json:"-" xorm:"delete_time deleted"`
}

type RerservationRegistration struct {
	CustomerID   int64                       `json:"customer_id"`
	ShopID       int64                       `json:"shop_id"`
	Reservations []MapReservationGoodService `json:"reservations"`
	CustomerNote string                      `json:"customer_note" xorm:"customer_note"`
}

type MapReservationGoodService struct {
	ID              int64      `xorm:"id pk autoincr"`
	ReservationID   int64      `json:"-" xorm:"reservation_id"`
	GoodServiceID   int64      `json:"good_service_id" xorm:"good_service_id"`
	Quantity        int        `json:"quantity" xorm:"quantity"`
	Note            string     `json:"note" xorm:"note"`
	PricePerItemIDR int64      `json:"-" xorm:"price_per_item_idr"`
	CreateTime      *time.Time `json:"-" xorm:"create_time created"`
	DeleteTime      *time.Time `json:"-" xorm:"delete_time deleted"`
	UpdateTime      *time.Time `json:"-" xorm:"update_time updated"`
}
