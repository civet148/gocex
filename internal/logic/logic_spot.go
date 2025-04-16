package logic

import (
	"github.com/civet148/gocex/internal/api"
	"github.com/civet148/gocex/internal/config"
)

type SpotLogic struct {
	cex api.CexApi
	cfg *config.Config
}

func NewSpotLogic(c *config.Config, cex api.CexApi) *SpotLogic {
	return &SpotLogic{
		cfg: c,
		cex: cex,
	}
}

func (l *SpotLogic) Exec() {

}
