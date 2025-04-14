package config

import "time"

type Contract struct {
	Symbol         string        `yaml:"Symbol"`         //下单交易对
	PullBackRate   float64       `yaml:"PullBackRate"`   //价格回调比例
	OrderInterval  time.Duration `yaml:"OrderInterval"`  //下单时间间隔
	TickerInterval time.Duration `yaml:"TickerInterval"` //市场价监控时间间隔
}

type Config struct {
	ApiKey        string   `yaml:"ApiKey" env:"API_KEY"`               //交易所API Key
	ApiSecret     string   `yaml:"ApiSecret" env:"API_SECRET"`         //交易所API密钥
	ApiPassphrase string   `yaml:"ApiPassphrase" env:"API_PASSPHRASE"` //交易所API密码
	CexName       string   `yaml:"CexName"`                            //交易所名称(okex/binance...)
	Contract      Contract `yaml:"Contract"`                           //合约配置
}
