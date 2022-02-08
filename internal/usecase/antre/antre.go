package antre

import (
	"github.com/eifzed/antre-app/internal/config"
	"github.com/eifzed/antre-app/internal/entity/repo/antre"
	db "github.com/eifzed/antre-app/lib/database/xorm"
)

const (
	wrapPrefix                    = "usercase.antre."
	wrapPrefixRegisterNewAccount  = wrapPrefix + "RegisterNewAccount."
	wrapPrefixAssignNewRoleToUser = wrapPrefix + "AssignNewRoleToUser."
	wrapPrefixLogin               = wrapPrefix + "Login."
)

type AntreUC struct {
	AntreDB     antre.AntreDBInterface
	Config      *config.Config
	Transaction *db.DBTransaction
}

func NewAntreUC(antre *AntreUC) *AntreUC {
	return antre
}
