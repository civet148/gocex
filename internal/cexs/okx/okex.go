package okx

import (
	"github.com/civet148/gocex/internal/api"
	"github.com/civet148/gocex/internal/config"
	"github.com/civet148/gocex/internal/types"
	"github.com/tbtc-bot/go-okex"
)

type CexOkex struct {
	api.CexUnimplement
	client *okex.Client
}

func init() {
	api.RegisterCex(types.CexNameOkex, NewCex)
}

func NewCex(c *config.Config) api.CexApi {
	client := okex.NewClient(c.ApiKey, c.ApiSecret, c.ApiPassphrase)
	client.Debug = c.Debug
	return &CexOkex{
		client: client,
	}
}

func (m *CexOkex) Name() string {
	return string(types.CexNameOkex)
}
