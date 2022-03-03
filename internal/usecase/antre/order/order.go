package order

import (
	"context"
	"errors"

	"github.com/eifzed/antre-app/internal/config"
	"github.com/eifzed/antre-app/internal/entity/order"
	orderRepo "github.com/eifzed/antre-app/internal/entity/repo/antre/order"
	"github.com/eifzed/antre-app/internal/handler/http/middleware/auth"
	"github.com/eifzed/antre-app/lib/common/commonerr"
	"github.com/eifzed/antre-app/lib/common/databaseerr"
	db "github.com/eifzed/antre-app/lib/database/xorm"
	"github.com/eifzed/antre-app/lib/utility"
	pkgErr "github.com/pkg/errors"
)

const (
	wrapPrefix                        = "usecase.antre.order"
	wrapPrefixRegisterShop            = wrapPrefix + "RegisterShop."
	wrapPrefixRegisterOrder           = wrapPrefix + "RegisterOrder."
	wrapPrefixGetUserOrders           = wrapPrefix + "GetUserOrders."
	wrapPrefixGetProductsListByShopID = wrapPrefix + "GetProductsListByShopID."
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
	for i := range orderData.Orders {
		orderData.Orders[i].OrderID = trxOrder.OrderID
	}
	err = uc.insertMapOrderProduct(ctx, orderData)
	if err != nil {
		return pkgErr.Wrap(err, wrapPrefixRegisterOrder+"insertMapOrderProduct")
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

func (uc *OrderUC) insertMapOrderProduct(ctx context.Context, orderData order.OrderRegistration) error {
	if orderData.Orders == nil || len(orderData.Orders) == 0 {
		return commonerr.ErrorBadRequest("order options", "empty order options")
	}
	goodServiceIDs := []int64{}
	for _, o := range orderData.Orders {
		goodServiceIDs = append(goodServiceIDs, o.ProductID)
	}
	// validate that orders exist in shop options
	goodServiceList, err := uc.OrderDB.GetMapShopProductByShopID(ctx, orderData.ShopID, goodServiceIDs...)
	if err != nil {
		return pkgErr.Wrap(err, wrapPrefixRegisterOrder+"GetMapShopProductByShopID")
	}
	if len(goodServiceList) != len(goodServiceIDs) {
		return commonerr.ErrorBadRequest("good/service id", "invalid good/service IDs")
	}

	// get price of each good/service
	for i, r := range orderData.Orders {
		for _, o := range goodServiceList {
			if r.ProductID == o.ID {
				orderData.Orders[i].PricePerItemIDR = o.PriceIDR
				break
			}
		}
	}

	err = uc.OrderDB.InsertMapOrderProduct(ctx, orderData.Orders...)
	if err != nil {
		return pkgErr.Wrap(err, wrapPrefixRegisterOrder+"InsertMapOrderProduct")
	}
	return nil
}

func (uc *OrderUC) GetCustomerOrders(ctx context.Context) (order.DtlOrderList, error) {
	result := order.DtlOrderList{}
	userDetail, isExist := auth.GetUserDetailFromContext(ctx)
	if !isExist {
		return result, commonerr.ErrorForbidden("user does not exist")
	}
	if !utility.RoleExistInSlice(uc.Config.Roles.Customer, userDetail.Roles) {
		return result, commonerr.ErrorForbidden("user does not have customer role")
	}

	dtlOrder, err := uc.OrderDB.GetDtlOrdersByCustomerID(ctx, userDetail.UserID)
	if err != nil && !errors.Is(err, databaseerr.ErrorDataNotFound) {
		return result, pkgErr.Wrap(err, wrapPrefixGetUserOrders+"GetTrxOrderByCustomerID")
	}
	if errors.Is(err, databaseerr.ErrorDataNotFound) {
		return result, nil
	}
	for i, dtl := range dtlOrder {
		mapOrderProduct, err := uc.OrderDB.GetMapOrderProductByOrderID(ctx, dtl.OrderID)
		if err != nil {
			return result, pkgErr.Wrap(err, wrapPrefixGetUserOrders+"GetMapOrderProductByOrderID")
		}
		dtlOrder[i].Orders = mapOrderProduct
	}
	result.TotalOrder = int64(len(dtlOrder))
	result.OrderDetailList = dtlOrder
	return result, nil
}
