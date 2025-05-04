package config

import "time"

type Config struct {
	Debug           bool          `yaml:"Debug"`                              //开启调试模式
	ApiKey          string        `yaml:"ApiKey" env:"API_KEY"`               //交易所API Key
	ApiSecret       string        `yaml:"ApiSecret" env:"API_SECRET"`         //交易所API密钥
	ApiPassphrase   string        `yaml:"ApiPassphrase" env:"API_PASSPHRASE"` //交易所API密码
	Simulate        bool          `yaml:"Simulate"`                           //模拟交易
	CexName         string        `yaml:"CexName"`                            //交易所名称(okex/binance...)
	Symbol          string        `yaml:"Symbol"`                             //下单交易对
	Leverage        int32         `yaml:"Leverage"`                           //杠杆倍数
	Continuous      int32         `yaml:"continuous"`                         //价格持续上涨次数
	CheckDur        time.Duration `yaml:"CheckDur"`                           //市价检查间隔
	FastRise        float64       `yaml:"FastRise"`                           //暴涨或暴跌
	RiseThreshold   float64       `yaml:"RiseThreshold"`                      //涨幅触发阈值
	StopLossPct     float64       `yaml:"StopLossPct"`                        //止损百分比
	TakeProfitPct   float64       `yaml:"TakeProfitPct"`                      //止盈百分比
	PullBackRate    float64       `yaml:"PullBackRate"`                       //价格回调比例
	TradeAmountRate float64       `yaml:"TradeAmountRate"`                    //交易金额比例（余额百分比）
}
