package antre

import (
	"github.com/eifzed/antre-app/internal/config"
	"github.com/eifzed/antre-app/internal/entity/repo/antre"
)

type AntreUC struct {
	AntreDB antre.AntreDBInterface
	Config  *config.Config
}

func NewAntreUC(antre *AntreUC) *AntreUC {
	return antre
}
