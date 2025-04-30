package types

import "github.com/civet148/sqlca/v2"

// InstType define instance type
type InstType string

const (
	InstType_SPOT    InstType = "SPOT"    //现货
	InstType_SWAP    InstType = "SWAP"    //永续合约
	InstType_FUTURES InstType = "FUTURES" //交割合约
	InstType_OPTION  InstType = "OPTION"  //期权
)

// 合约信息
type InstrumentDetail struct {
	InstType     string        `json:"instType"`     // 产品类型：永续合约
	InstId       string        `json:"instId"`       // 合约ID（格式：基础货币-计价货币-类型）
	Uly          string        `json:"uly"`          // 标的资产（如PEPE-USDT指数）
	Category     string        `json:"category"`     // 手续费档位（1表示普通档）
	BaseCcy      string        `json:"baseCcy"`      // 基础货币（如PEPE）
	QuoteCcy     string        `json:"quoteCcy"`     // 计价货币（如USDT）
	SettleCcy    string        `json:"settleCcy"`    // 结算货币（通常与计价货币相同）
	CtVal        sqlca.Decimal `json:"ctVal"`        // 单张合约面值（例如：PEPE-USDT价格是0.000006 CtValue=10000000 那么一张的价值=60USDT)
	CtMult       sqlca.Decimal `json:"ctMult"`       // 合约乘数（通常为1）
	CtValCcy     string        `json:"ctValCcy"`     // 面值计价货币（USD与USDT 1:1锚定）
	OptType      string        `json:"optType"`      // 期权类型（永续合约为空）
	Stk          string        `json:"stk"`          // 行权价（期权专用，永续合约为空）
	ListTime     string        `json:"listTime"`     // 合约上线时间（Unix时间戳，毫秒）
	ExpTime      string        `json:"expTime"`      // 到期时间（永续合约为空）
	Lever        sqlca.Decimal `json:"lever"`        // 当前最大杠杆倍数（如75倍）
	TickSz       string        `json:"tickSz"`       // 价格最小变动单位（0.0000001 USDT）
	LotSz        sqlca.Decimal `json:"lotSz"`        // 数量最小单位（0.1张）
	MinSz        sqlca.Decimal `json:"minSz"`        // 最小下单数量（0.1张）
	CtType       string        `json:"ctType"`       // 合约类型（linear=正向，inverse=反向）
	Alias        string        `json:"alias"`        // 合约别名（永续合约通常为"swap"）
	State        string        `json:"state"`        // 合约状态（live=交易中）
	MaxLmtSz     sqlca.Decimal `json:"maxLmtSz"`     // 单笔限价单最大张数
	MaxMktSz     sqlca.Decimal `json:"maxMktSz"`     // 单笔市价单最大张数
	MaxTwapSz    sqlca.Decimal `json:"maxTwapSz"`    // 单笔TWAP订单最大张数
	MaxIcebergSz sqlca.Decimal `json:"maxIcebergSz"` // 单笔冰山订单最大张数
	MaxTriggerSz sqlca.Decimal `json:"maxTriggerSz"` // 单笔条件单最大张数
	MaxStopSz    sqlca.Decimal `json:"maxStopSz"`    // 单笔止损单最大张数
}
