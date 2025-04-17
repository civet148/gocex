package okx

import (
	"github.com/civet148/gocex/internal/api"
	"github.com/civet148/gocex/internal/config"
	"github.com/civet148/gocex/internal/types"
	"github.com/tbtc-bot/go-okex"
)

type CexOkex struct {
	api.CexUnimplInterface
	client *okex.Client
}

func init() {
	api.RegisterCex(types.CexNameOkex, NewCex)
}

func NewCex(c *config.Config) api.CexApi {
	return &CexOkex{
		client: okex.NewClient(c.ApiKey, c.ApiSecret, c.ApiPassphrase),
	}
}

func (c *CexOkex) Name() string {
	return string(types.CexNameOkex)
}
