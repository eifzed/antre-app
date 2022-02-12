package order

// Table constant
const (
	tblHstOrder           = "ant_hst_order"
	tblDtlShop            = "ant_dtl_shop"
	tblMapShopGoodService = "ant_map_shop_good_service"
	tblMapShopCategory    = "ant_map_shop_category"
)

// error wraps
const (
	wrapPrefix                    = "repo.antre.order."
	wrapPrefixGetDtlShopByOwnerID = wrapPrefix + "GetDtlShopByOwnerID."
)
