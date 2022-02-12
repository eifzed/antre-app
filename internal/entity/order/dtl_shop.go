package order

import "time"

type ShopRegistration struct {
	ShopID int64 `json:"-" xorm:"shop_id pk autoincr"`
	// CompanyID          int64               `json:"company_id" xorm:"company_id"`
	// PICID              int64               `json:"pic_id" xorm:"pic_id"`
	// OwnerID            int64               `json:"owner_id" xorm:"owner_id"`
	ShopName           string              `json:"shop_name" xorm:"shop_name"`
	ShopType           string              `json:"shop_type" xorm:"shop_type"`
	ShopDescription    string              `json:"shop_description" xorm:"shop_description"`
	Address            string              `json:"address" xorm:"address"`
	ShopPictureURL     string              `json:"shop_picture_url" xorm:"shop_picture_url"`
	PostalCode         int64               `json:"postal_code" xorm:"postal_code"`
	GoodServiceOptions []GoodServiceOption `json:"good_service_options"`
	OpenHour           int                 `json:"open_hour" xorm:"open_hour"`
	CloseHour          int                 `json:"close_hour" xorm:"close_hour"`
	CategoryLv0        int64               `json:"category_lv0" xorm:"category"`
	CreateTime         time.Time           `xorm:"create_time created"`
	UpdateTime         time.Time           `xorm:"update_time updated"`
	DeleteTime         *time.Time          `xorm:"delete_time deleted"`
}

type GoodServiceType string

const (
	GoodType    GoodServiceType = "good"
	ServiceType GoodServiceType = "service"
)

type GoodServiceOption struct {
	ID                 int64           `json:"id" xorm:"id pk autoincr"`
	ShopID             int64           `json:"-" xorm:"shop_id"`
	Type               GoodServiceType `json:"type" xorm:"type"`
	Name               string          `json:"name" xorm:"name"`
	Description        string          `json:"description" xorm:"description"`
	PriceIDR           int64           `json:"price_idr" xorm:"price_idr"`
	ProcessTimeMinutes int64           `json:"process_time_minutes" xorm:"process_time_minutes"`
	PictureURL         string          `json:"picture_url" xorm:"picture_url"`
	CreateTime         time.Time       `xorm:"create_time created"`
	UpdateTime         time.Time       `xorm:"update_time updated"`
	DeleteTime         *time.Time      `xorm:"delete_time deleted"`
}

type DtlShop struct {
	ShopID         int64      `xorm:"shop_id pk autoincr"`
	OwnerID        int64      `xorm:"owner_id"`
	ShopName       string     `xorm:"shop_name"`
	ShopType       string     `xorm:"shop_type"`
	Address        string     `xorm:"address"`
	PostalCode     int64      `xorm:"postal_code"`
	OpenHour       int        `json:"open_hour" xorm:"open_hour"`
	CloseHour      int        `json:"close_hour" xorm:"close_hour"`
	ShopPictureURL string     `json:"shßop_picture_url" xorm:"shop_picture_url"`
	CreateTime     time.Time  `xorm:"create_time created"`
	UpdateTime     time.Time  `xorm:"update_time updated"`
	DeleteTime     *time.Time `xorm:"delete_time deleted"`
}

type MapShopCategory struct {
	ShopID     int64 `xorm:"shop_id"`
	CategoryID int64 `xorm:"category_id"`
}
