package order

import "time"

type DtlOrderList struct {
	TotalOrder      int64      `json:"total_orders"`
	OrderDetailList []DtlOrder `json:"order_detail_list"`
}

type DtlOrder struct {
	OrderID      int64                 `json:"order_id" xorm:"order_id"`
	ShopID       int64                 `json:"shop_id" xorm:"shop_id"`
	ShopName     string                `json:"shop_name" xorm:"shop_name"`
	Orders       []MapOrderGoodService `json:"orders" xorm:"-"`
	Status       string                `json:"status_name" xorm:"status_name"`
	CustomerNote string                `json:"customer_note" xorm:"customer_note"`
	ShopNote     string                `json:"shop_note" xorm:"shop_note"`
	CreateTime   time.Time             `json:"create_time" xorm:"create_time"`
}
