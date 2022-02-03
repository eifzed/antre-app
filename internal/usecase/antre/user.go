package antre

import (
	"context"
	"errors"

	"github.com/eifzed/antre-app/internal/entity/auth"
	"github.com/eifzed/antre-app/lib/common/databaseerr"
)

func (uc *AntreUC) RegisterNewUser(ctx context.Context, params auth.User) error {
	_, err := uc.AntreDB.GetUserByEmail(ctx, params.Email)
	if err != nil && !errors.Is(err, databaseerr.ErrorDataNotFound) {
		return err
	}
	if errors.Is(err, databaseerr.ErrorDataNotFound) {
		err = uc.AntreDB.InsertUser(ctx, &params)
		if err != nil {
			return nil
		}
	}

	return nil
}
