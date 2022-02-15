package order

// Table constant
const (
	tblHstOrder            = "ant_hst_order"
	tblDtlShop             = "ant_dtl_shop"
	tblMapShopGoodService  = "ant_map_shop_good_service"
	tblMapShopCategory     = "ant_map_shop_category"
	tblMapOrderGoodService = "ant_map_order_good_service"
	tblTrxOrder            = "ant_trx_order"
	tblMstStatus           = "ant_mst_status"
)

// error wraps
const (
	wrapPrefix                              = "repo.antre.order."
	wrapPrefixGetDtlShopByOwnerID           = wrapPrefix + "GetDtlShopByOwnerID."
	wrapPrefixInsertMapOrderGoodService     = wrapPrefix + "InsertMapOrderGoodService."
	wrapPrefixGetMapShopGoodServiceByShopID = wrapPrefix + "GetMapShopGoodServiceByShopID."
	wrapPrefixInsertMapShopGoodService      = wrapPrefix + "InsertMapShopGoodService."
	wrapPrefixGetDtlShopByShopID            = wrapPrefix + "GetDtlShopByShopID."
	wrapPrefixGetOrdersByCustomerID         = wrapPrefix + "GetOrdersByCustomerID."
)
