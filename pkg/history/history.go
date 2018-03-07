package history

import "github.com/mineralres/goshare/aproto"

// GetKData 请求历史K线数据
/*
symbol：股票代码，即6位数字代码，或者指数代码（sh=上证指数 sz=深圳成指 hs300=沪深300指数 sz50=上证50 zxb=中小板 cyb=创业板）
startDate：开始日期，格式20180307
endDate：结束日期，格式20180307
period：周期
retryCount：当网络异常后重试次数，默认为3
*/
func GetKData(symbol *aproto.Symbol, period aproto.PeriodType, startDate, endDate, retryCount int) (*aproto.KlineSeries, error) {
	var ret aproto.KlineSeries
	return &ret, nil
}
