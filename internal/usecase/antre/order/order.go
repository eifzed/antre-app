package order

import (
	"context"

	"github.com/eifzed/antre-app/internal/config"
	"github.com/eifzed/antre-app/internal/entity/order"
	orderRepo "github.com/eifzed/antre-app/internal/entity/repo/antre/order"
	"github.com/eifzed/antre-app/internal/handler/http/middleware/auth"
	"github.com/eifzed/antre-app/lib/common/commonerr"
	db "github.com/eifzed/antre-app/lib/database/xorm"
	"github.com/eifzed/antre-app/lib/utility"
	pkgErr "github.com/pkg/errors"
)

const (
	wrapPrefix              = "usecase.antre.order"
	wrapPrefixRegisterShop  = wrapPrefix + "RegisterShop."
	wrapPrefixRegisterOrder = wrapPrefix + "RegisterOrder."
)

type OrderUC struct {
	OrderDB     orderRepo.Order
	Config      *config.Config
	Transaction *db.DBTransaction
}

func NewOrderUC(resrvation *OrderUC) *OrderUC {
	return resrvation
}

func (uc *OrderUC) GetOrderByID(ctx context.Context, orderID int64) (*order.TrxOrder, error) {
	// TODO: get user detail from context
	return uc.OrderDB.GetOrderByID(ctx, orderID)
}

func (uc *OrderUC) RegisterOrder(ctx context.Context, orderData order.OrderRegistration) error {
	userDetail, isExist := auth.GetUserDetailFromContext(ctx)
	if !isExist {
		return commonerr.ErrorForbidden("user does not exist")
	}
	if !utility.RoleExistInSlice(uc.Config.Roles.Customer, userDetail.Roles) {
		return commonerr.ErrorUnauthorized("User does not have customer role")
	}

	ctx, err := uc.Transaction.Begin(ctx)
	defer uc.Transaction.Finish(ctx, &err)
	trxOrder := order.TrxOrder{
		CustomerID:   userDetail.UserID,
		ShopID:       orderData.ShopID,
		StatusID:     StatusRegistered,
		CustomerNote: orderData.CustomerNote,
	}

	err = uc.OrderDB.InsertTrxOrder(ctx, &trxOrder)
	if err != nil {
		return pkgErr.Wrap(err, wrapPrefixRegisterOrder+"InsertTrxOrder")
	}
	err = uc.insertMapOrderGoodService(ctx, orderData)
	if err != nil {
		return pkgErr.Wrap(err, wrapPrefixRegisterOrder+"insertMapOrderGoodService")
	}

	hstOrder := &order.HstOrder{
		OrderID:      trxOrder.OrderID,
		CustomerID:   userDetail.UserID,
		ShopID:       orderData.ShopID,
		StatusID:     StatusRegistered,
		CustomerNote: orderData.CustomerNote,
		UpdaterID:    userDetail.UserID,
		Reason:       ReasonNewOrder,
	}
	err = uc.OrderDB.InsertHstOrder(ctx, hstOrder)
	if err != nil {
		return pkgErr.Wrap(err, wrapPrefixRegisterOrder+"InsertTrxOrder")
	}
	return nil
}

func (uc *OrderUC) insertMapOrderGoodService(ctx context.Context, order order.OrderRegistration) error {
	return nil
}
