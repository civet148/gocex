package logic

import (
	"github.com/civet148/gocex/internal/config"
	"github.com/civet148/log"
)

type CexLogic struct {
}

func NewCexLogic(c *config.Config) *CexLogic {
	return &CexLogic{}
}

func (m *CexLogic) Run() error {
	log.Infof("running...")
	return nil
}
