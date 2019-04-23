# 交易合约

获取交易合约的基本属性
* 有些字段是某些交易品种的专有属性，如行权价只有期权合约才有,根据数据源的情况goshare会尽量填充这些字段

## 获取合约列表

```
import (
  "github.com/mineralres/goshare"
)

var s goshare.SSEOfficialSource
// 获取上证所50ETF期权列表
list, err := s.GetSSEStockOptionList()
// 上证股票列表
list, err := s.GetSSEStockOptionList()

```

## 合约字段详情

```
// 合约行情信息
message InstrumentInfo {
  /// 合约名称.
  string symbol_name = 1;
  /// 是否单边大保
  int32 position_rank = 2;
  /// 更新时间.
  int64 update_time = 3;
  // 交易日
  int32 update_trading_day = 4;
  /// 涨停价.
  double upper_limit_price = 5;
  /// 跌停价.
  double lower_limit_price = 6;
  // 昨收
  double pre_close_price = 7;
  // 昨结
  double pre_settlement_price = 8;
  // 昨成交
  int32 pre_volume = 9;
  // 昨持仓
  int32 pre_position = 10;
  // 今天结算价
  double settlement_price = 11;
  // 创建日期
  int32 create_date = 12;
  // 上市日期
  int32 open_date = 13;
  // 到期日期
  int32 expire_date = 14;
  // 开始交割日期
  int32 start_deliver_date = 15;
  // 结束交割日期
  int32 end_deliver_date = 16;
  // 是否T+0
  int32 is_close_today_allowed = 17;
  // 市价最大下单量
  int32 max_market_order_volume = 18;
  // 市价最小下单量
  int32 min_market_order_volume = 19;
  // 限价最大下单量
  int32 max_limit_order_volume = 20;
  // 限价最小下单量
  int32 min_limit_order_volume = 21;
  // 基础合约
  Symbol base_symbol = 22;
  // 行权价
  double strike_price = 23;
  ///期权类型
  OptionCallPutType call_put_type = 24;
  ///合约基础商品乘数
  double underlying_multiple = 25;
  ///组合类型
  int32 combination_type = 26;
  ///最小买下单单位
  int32 min_buy_volume = 27;
  ///最小卖下单单位
  int32 min_sell_volume = 28;
  ///合约标识码
  Symbol instrument_code = 29;
  // 是否在交易
  bool is_trading = 30;
  // 行权方式类型
  OptionDeliveryDateType delivery_date_type = 31;
}
```