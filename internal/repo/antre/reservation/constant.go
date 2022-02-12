package reservation

// Table constant
const (
	tblHstReservation     = "ant_hst_reservation"
	tblDtlShop            = "ant_dtl_shop"
	tblMapShopGoodService = "ant_map_shop_good_service"
	tblMapShopCategory    = "ant_map_shop_category"
)

// error wraps
const (
	wrapPrefix                    = "repo.antre.reservation."
	wrapPrefixGetDtlShopByOwnerID = wrapPrefix + "GetDtlShopByOwnerID."
)
