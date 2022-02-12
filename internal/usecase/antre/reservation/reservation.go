package reservation

import (
	"context"

	"github.com/eifzed/antre-app/internal/config"
	rsvRepo "github.com/eifzed/antre-app/internal/entity/repo/antre/reservation"
	rsv "github.com/eifzed/antre-app/internal/entity/reservation"
	"github.com/eifzed/antre-app/internal/handler/http/middleware/auth"
	"github.com/eifzed/antre-app/lib/common/commonerr"
	db "github.com/eifzed/antre-app/lib/database/xorm"
	"github.com/eifzed/antre-app/lib/utility"
	pkgErr "github.com/pkg/errors"
)

const (
	wrapPrefix                    = "usecase.antre.reservation"
	wrapPrefixRegisterShop        = wrapPrefix + "RegisterShop."
	wrapPrefixRegisterReservation = wrapPrefix + "RegisterReservation."
)

type ReservationUC struct {
	ReservationDB rsvRepo.Reservation
	Config        *config.Config
	Transaction   *db.DBTransaction
}

func NewReservationUC(resrvation *ReservationUC) *ReservationUC {
	return resrvation
}

func (uc *ReservationUC) GetReservationByID(ctx context.Context, rsvID int64) (*rsv.TrxReservation, error) {
	// TODO: get user detail from context
	return uc.ReservationDB.GetReservationByID(ctx, rsvID)
}

func (uc *ReservationUC) RegisterReservation(ctx context.Context, reservation rsv.RerservationRegistration) error {
	userDetail, isExist := auth.GetUserDetailFromContext(ctx)
	if !isExist {
		return commonerr.ErrorForbidden("user does not exist")
	}
	if !utility.RoleExistInSlice(uc.Config.Roles.Customer, userDetail.Roles) {
		return commonerr.ErrorUnauthorized("User does not have customer role")
	}

	ctx, err := uc.Transaction.Begin(ctx)
	defer uc.Transaction.Finish(ctx, &err)

	trxReservation := rsv.TrxReservation{
		CustomerID:   userDetail.UserID,
		ShopID:       reservation.ShopID,
		StatusID:     StatusRegistered,
		CustomerNote: reservation.CustomerNote,
	}

	err = uc.ReservationDB.InsertTrxReservation(ctx, &trxReservation)
	if err != nil {
		return pkgErr.Wrap(err, wrapPrefixRegisterReservation+"InsertTrxReservation")
	}
	err = uc.insertMapReservationGoodService(ctx, reservation)
	if err != nil {
		return pkgErr.Wrap(err, wrapPrefixRegisterReservation+"insertMapReservationGoodService")
	}

	hstReservation := &rsv.HstReservation{
		ReservationID: trxReservation.ReservationID,
		CustomerID:    userDetail.UserID,
		ShopID:        reservation.ShopID,
		StatusID:      StatusRegistered,
		CustomerNote:  reservation.CustomerNote,
		UpdaterID:     userDetail.UserID,
		Reason:        ReasonNewReservation,
	}
	err = uc.ReservationDB.InsertHstReservation(ctx, hstReservation)
	if err != nil {
		return pkgErr.Wrap(err, wrapPrefixRegisterReservation+"InsertTrxReservation")
	}
	return nil
}

func (uc *ReservationUC) insertMapReservationGoodService(ctx context.Context, reservation rsv.RerservationRegistration) error {
	return nil
}
