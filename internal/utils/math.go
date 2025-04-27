package utils

import (
	"fmt"
	"github.com/civet148/sqlca/v2"
	"github.com/shopspring/decimal"
	"strconv"
)

func FormatDecimal(d sqlca.Decimal, precision int) string {
	// 获取整数部分
	intPart := d.IntPart()

	// 获取小数部分并格式化为指定精度
	fracPart := d.Sub(decimal.NewFromInt(intPart))
	fracStr := fracPart.StringFixed(int32(precision))[2:] // 去掉前面的"0."

	// 将整数和小数组合
	return fmt.Sprintf("%d.%s", intPart, fracStr)
}

func AnyToString(v any) string {
	return fmt.Sprintf("%v", v)
}

func StringToUint64(v string) uint64 {
	u, _ := strconv.ParseUint(v, 10, 64)
	return u
}

func StringToInt32(v string) int32 {
	u, _ := strconv.ParseUint(v, 10, 32)
	return int32(u)
}
