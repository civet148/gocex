package logic

import (
	"context"
	"github.com/civet148/gocex/internal/api"
	"github.com/civet148/gocex/internal/types"
	"github.com/civet148/log"
	"github.com/civet148/sqlca/v2"
)

type Ticker interface {
	GetCurrentPrice(instId string) (price sqlca.Decimal, err error)
	GetAvailableUSDT() (available sqlca.Decimal, err error)
}

type TickerLogic struct {
	cex api.CexApi
}

func NewTickerLogic(cex api.CexApi) *TickerLogic {
	return &TickerLogic{
		cex: cex,
	}
}

func (l *TickerLogic) GetCurrentPrice(instId string) (price sqlca.Decimal, err error) {
	ts, err := l.cex.GetTickerPrice(context.Background(), instId)
	if err != nil {
		return price, log.Errorf(err)
	}
	if len(ts) == 0 {
		return price, log.Errorf("instId %s price ticker not found", instId)
	}
	price = ts[0].AskPx
	return price, nil
}

func (l *TickerLogic) GetAvailableUSDT() (sqlca.Decimal, error) {
	bal, err := l.cex.GetBalance(context.Background(), types.USDT)
	if err != nil {
		return sqlca.Decimal{}, log.Errorf(err)
	}
	if len(bal.Details) == 0 {
		return sqlca.NewDecimal(0), nil
	}
	for _, d := range bal.Details {
		if d.Ccy == types.USDT {
			return d.AvailBal, nil
		}
	}
	return sqlca.NewDecimal(0), nil
}
