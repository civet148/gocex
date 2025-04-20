package logic

import (
	"context"
	"github.com/civet148/gocex/internal/api"
	"github.com/civet148/gocex/internal/locker"
	"github.com/civet148/log"
	"github.com/civet148/sqlca/v2"
	"time"
)

type Ticker interface {
	GetCurrentPrice() sqlca.Decimal
	GetLowestPrice() sqlca.Decimal
	GetHighestPrice() sqlca.Decimal
}

type TickerLogic struct {
	cex          api.CexApi
	currentPrice sqlca.Decimal
	lowestPrice  sqlca.Decimal
	highestPrice sqlca.Decimal
	canceler     context.CancelFunc
}

func NewTickerLogic(cex api.CexApi, symbol string, interval time.Duration) *TickerLogic {
	ticker := &TickerLogic{
		cex: cex,
	}
	return ticker.start(symbol, interval)
}

func (l *TickerLogic) start(symbol string, interval time.Duration) *TickerLogic {
	ctx, canceler := context.WithCancel(context.Background())
	l.canceler = canceler
	go l.startTicker(ctx, symbol, interval)
	return l
}

func (l *TickerLogic) startTicker(ctx context.Context, symbol string, interval time.Duration) {
	tc := time.NewTicker(interval)
	err := l.refreshMarketPrice(symbol)
	if err != nil {
		log.Printf("get %s market price error: %s", symbol, err)
		return
	}

	for {
		select {
		case <-tc.C:
			_ = l.refreshMarketPrice(symbol)
		case <-ctx.Done():
			return
		}
		time.Sleep(time.Second)
	}
}

func (l *TickerLogic) refreshMarketPrice(symbol string) error {
	ts, err := l.cex.GetTickerPrice(symbol)
	if err != nil {
		return log.Errorf(err)
	}
	if len(ts) == 0 {
		return log.Errorf("symbol %s price ticker not found", symbol)
	}
	price := ts[0]

	unlock := locker.Lock()
	defer unlock()

	marketPrice := price.AskPx
	l.currentPrice = marketPrice

	if l.lowestPrice.IsZero() {
		l.lowestPrice = marketPrice
	} else if l.lowestPrice.GreaterThan(l.currentPrice) {
		l.lowestPrice = marketPrice
	}
	if l.highestPrice.IsZero() {
		l.highestPrice = marketPrice
	} else if l.highestPrice.LessThan(marketPrice) {
		l.highestPrice = marketPrice
	}
	log.Printf("[%s] market price [%v]", symbol, l.currentPrice)
	return nil
}

func (l *TickerLogic) Stop() {
	l.canceler()
}

func (l *TickerLogic) GetCurrentPrice() sqlca.Decimal {
	return l.currentPrice
}

func (l *TickerLogic) GetLowestPrice() sqlca.Decimal {
	return l.lowestPrice
}

func (l *TickerLogic) GetHighestPrice() sqlca.Decimal {
	return l.highestPrice
}
