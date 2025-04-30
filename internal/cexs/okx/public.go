package okx

import (
	"context"
	"github.com/civet148/gocex/internal/types"
	"github.com/civet148/log"
	"github.com/jinzhu/copier"
	"github.com/tbtc-bot/go-okex"
)

func (m *CexOkex) GetInstrument(ctx context.Context, instId string, instType types.InstType) (details []*types.InstrumentDetail, err error) {
	svc := m.client.NewGetInstrumentsService()
	if instId == "" || instType == "" {
		return nil, log.Errorf("inst id and type is requires")
	}
	if instType == types.InstType_SWAP {
		instId = types.ToSwapInstId(m.Name(), instId)
	}
	res, err := svc.InstrumentId(instId).InstrumentType(okex.InstType(instType)).Do(ctx)
	if err != nil {
		return nil, log.Errorf("get ticker error: %s", err)
	}
	_ = copier.Copy(&details, &res.Data)
	return details, nil
}
