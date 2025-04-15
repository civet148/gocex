package config

import "time"

type Market struct {
	PriceFlat float64 `yaml:"PriceFlat"` //正负范围内视为持平
	SlowRise  float64 `yaml:"SlowRise"`  //慢涨
	FastRise  float64 `yaml:"FastRise"`  //暴涨
}

type Contract struct {
	Symbol           string        `yaml:"Symbol"`           //下单交易对
	PullBackRate     float64       `yaml:"PullBackRate"`     //价格回调比例
	TradeAmountRate  float64       `yaml:"TradeAmountRate"`  //交易金额比例（余额百分比）
	LeverMultiple    int32         `yaml:"LeverMultiple"`    //杠杆倍数
	TickerInterval   time.Duration `yaml:"TickerInterval"`   //市场价监控时间间隔
	OrderInterval    time.Duration `yaml:"OrderInterval"`    //下单时间间隔
	ClosePosInterval time.Duration `yaml:"ClosePosInterval"` //订单下单后多长时间内不允许平仓
}

type Config struct {
	ApiKey        string   `yaml:"ApiKey" env:"API_KEY"`               //交易所API Key
	ApiSecret     string   `yaml:"ApiSecret" env:"API_SECRET"`         //交易所API密钥
	ApiPassphrase string   `yaml:"ApiPassphrase" env:"API_PASSPHRASE"` //交易所API密码
	CexName       string   `yaml:"CexName"`                            //交易所名称(okex/binance...)
	Market        Market   `yaml:"Market"`                             //市场价格规则配置
	Contract      Contract `yaml:"Contract"`                           //合约配置
}
