package reservation

import (
	"context"
	"errors"

	rsv "github.com/eifzed/antre-app/internal/entity/reservation"
	"github.com/eifzed/antre-app/internal/handler/http/middleware/auth"
	"github.com/eifzed/antre-app/lib/common/commonerr"
	"github.com/eifzed/antre-app/lib/common/databaseerr"
	"github.com/eifzed/antre-app/lib/utility"
	pkgErr "github.com/pkg/errors"
)

func (uc *ReservationUC) RegisterShop(ctx context.Context, shopRegistData rsv.ShopRegistration) error {
	userDetail, exist := auth.GetUserDetailFromContext(ctx)
	if !exist {
		return commonerr.ErrorUnauthorized("user is not authorized")
	}

	if !utility.RoleExistInSlice(uc.Config.Roles.Owner, userDetail.Roles) {
		return commonerr.ErrorUnauthorized("User does not have owner role")
	}

	_, err := uc.ReservationDB.GetDtlShopByOwnerID(ctx, userDetail.UserID)
	if err != nil && !errors.Is(err, databaseerr.ErrorDataNotFound) {
		return pkgErr.Wrap(err, wrapPrefixRegisterShop+"GetDtlShopByOwnerID")
	}

	if errors.Is(err, databaseerr.ErrorDataNotFound) {
		ctx, err = uc.Transaction.Begin(ctx)
		if err != nil {
			return pkgErr.Wrap(err, wrapPrefixRegisterShop+"Transaction.Begin")
		}
		defer uc.Transaction.Finish(ctx, &err)
		//register shop
		dtlShop := rsv.DtlShop{
			ShopID:         shopRegistData.ShopID,
			OwnerID:        userDetail.UserID,
			ShopName:       shopRegistData.ShopName,
			ShopType:       shopRegistData.ShopType,
			PostalCode:     shopRegistData.PostalCode,
			Address:        shopRegistData.Address,
			OpenHour:       shopRegistData.OpenHour,
			CloseHour:      shopRegistData.CloseHour,
			ShopPictureURL: shopRegistData.ShopPictureURL,
		}
		err = uc.ReservationDB.InsertDtlShopByOwnerID(ctx, &dtlShop)
		if err != nil {
			return pkgErr.Wrap(err, wrapPrefixRegisterShop+"InsertDtlShopByOwnerID")
		}
		shopRegistData.ShopID = dtlShop.ShopID

		err = uc.insertShopGoodService(ctx, shopRegistData)
		if err != nil {
			return pkgErr.Wrap(err, wrapPrefixRegisterShop+"insertShopGoodService")
		}
		err = uc.insertShopCategory(ctx, shopRegistData)
		if err != nil {
			return pkgErr.Wrap(err, wrapPrefixRegisterShop+"insertShopCategory")
		}
	}
	return nil
}

func (uc *ReservationUC) insertShopCategory(ctx context.Context, shopRegistData rsv.ShopRegistration) error {
	categories := []rsv.MapShopCategory{}
	if shopRegistData.CategoryLv0 > 0 {
		categories = append(categories, rsv.MapShopCategory{
			ShopID:     shopRegistData.ShopID,
			CategoryID: shopRegistData.CategoryLv0,
		})
	}
	// TODO: add level 1 and 2 category
	err := uc.ReservationDB.InsertMapShopCategory(ctx, categories...)
	if err != nil {
		return pkgErr.Wrap(err, "insertShopCategory.InsertMapShopCategory")
	}
	return nil
}

func (uc *ReservationUC) insertShopGoodService(ctx context.Context, shopRegistData rsv.ShopRegistration) error {
	shopID := shopRegistData.ShopID
	for i := range shopRegistData.GoodServiceOptions {
		shopRegistData.GoodServiceOptions[i].ShopID = shopID
	}
	err := uc.ReservationDB.InsertMapShopGoodService(ctx, shopRegistData.GoodServiceOptions...)
	if err != nil {
		return pkgErr.Wrap(err, "insertShopGoodService.InsertMapShopGoodService")
	}
	return nil
}
