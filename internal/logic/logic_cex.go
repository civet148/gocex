package logic

import (
	"github.com/civet148/gocex/internal/api"
	_ "github.com/civet148/gocex/internal/cexs/okx"
	"github.com/civet148/gocex/internal/config"
	"github.com/civet148/gocex/internal/types"
)

type CexLogic struct {
	cfg *config.Config
	cex api.CexApi
}

func NewCexLogic(cfg *config.Config) *CexLogic {
	return &CexLogic{
		cfg: cfg,
		cex: api.NewCex(types.CexName(cfg.CexName), cfg),
	}
}

func (m *CexLogic) Run() error {
	contract := NewContractLogic(m.cfg, m.cex)
	return contract.Exec()
}
