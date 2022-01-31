package reservation

import (
	"context"
	"errors"

	rsv "github.com/eifzed/antre-app/internal/entity/reservation"
	"github.com/eifzed/antre-app/lib/common/databaseerr"
)

func (uc *ReservationUC) RegisterNewUser(ctx context.Context, params rsv.User) error {
	_, err := uc.ReservationDB.GetUserByEmail(ctx, params.Email)
	if err != nil && !errors.Is(err, databaseerr.ErrorDataNotFound) {
		return err
	}
	if errors.Is(err, databaseerr.ErrorDataNotFound) {
		err = uc.ReservationDB.InsertUser(ctx, &params)
		if err != nil {
			return nil
		}
	}

	return nil
}
