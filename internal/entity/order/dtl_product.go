package order

type ProductsList struct {
	Total    int64     `json:"total"`
	Products []Product `json:"products"`
}

type Product struct {
	ID         int64   `json:"id" xorm:"id pk autoincr"`
	Name       string  `json:"name" xorm:"name"`
	PriceIDR   string  `json:"price" xorm:"price_idr"`
	PictureURL string  `json:"picture_url" xorm:"picture_url"`
	Rating     float32 `json:"rating" xorm:"rating"`
}
