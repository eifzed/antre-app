package order

import (
	"time"
)

type TrxOrder struct {
	OrderID      int64      `json:"order_id" xorm:"order_id pk autoincr"`
	CustomerID   int64      `json:"customer_id" xorm:"customer_id"`
	ShopID       int64      `json:"shop_id" xorm:"shop_id"`
	StatusID     int64      `json:"status_id" xorm:"status_id"`
	CustomerNote string     `json:"customer_note" xorm:"customer_note"`
	ShopNote     string     `json:"shop_note" xorm:"shop_note"`
	CreateTime   time.Time  `json:"-" xorm:"create_time created"`
	UpdateTime   time.Time  `json:"-" xorm:"update_time updated"`
	DeleteTime   *time.Time `json:"-" xorm:"delete_time deleted"`
}

type OrderRegistration struct {
	ShopID       int64              `json:"shop_id"`
	Orders       []*MapOrderProduct `json:"orders"`
	CustomerNote string             `json:"customer_note"`
}

type MapOrderProduct struct {
	ID              int64      `json:"-" xorm:"id pk autoincr"`
	OrderID         int64      `json:"-" xorm:"order_id"`
	ProductID       int64      `json:"product_id" xorm:"product_id"`
	Quantity        int        `json:"quantity" xorm:"quantity"`
	Note            string     `json:"note" xorm:"note"`
	PricePerItemIDR int64      `json:"-" xorm:"price_per_item_idr"`
	CreateTime      time.Time  `json:"-" xorm:"create_time created"`
	UpdateTime      time.Time  `json:"-" xorm:"update_time updated"`
	DeleteTime      *time.Time `json:"-" xorm:"delete_time deleted"`
}
