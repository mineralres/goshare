package ftdc

import (
	"encoding/binary"
	"io"
	"strings"
)

func writeLen(w io.Writer, source string, fixedLen int) {
	if len(source) < fixedLen {
		io.WriteString(w, source)
		b := make([]byte, fixedLen-len(source))
		w.Write(b)
		return
	}
	io.WriteString(w, source[:fixedLen])
}

func readLen(r io.Reader, fixedLen int) string {
	buf := make([]byte, fixedLen)
	n, err := r.Read(buf)
	if err != nil || n != fixedLen {
		return ""
	}
	return strings.Trim(string(buf), "\u0000")
}

type  CThostFtdcDisseminationField struct {
    SequenceSeries int16;
    SequenceNo int32;
}

func (m* CThostFtdcDisseminationField) Marshal(w io.Writer) {
    binary.Write(w, binary.BigEndian, &m.SequenceSeries)
    binary.Write(w, binary.BigEndian, &m.SequenceNo)
}

func (m* CThostFtdcDisseminationField) Unmarshal(r io.Reader) {
    binary.Read(r, binary.BigEndian, &m.SequenceSeries)
    binary.Read(r, binary.BigEndian, &m.SequenceNo)
}

type  CThostFtdcReqUserLoginField struct {
    TradingDay string;
    BrokerID string;
    UserID string;
    Password string;
    UserProductInfo string;
    InterfaceProductInfo string;
    ProtocolInfo string;
    MacAddress string;
    OneTimePassword string;
    ClientIPAddress string;
    LoginRemark string;
    ClientIPPort int32;
}

func (m* CThostFtdcReqUserLoginField) Marshal(w io.Writer) {
    writeLen(w, m.TradingDay, 9)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.Password, 41)
    writeLen(w, m.UserProductInfo, 11)
    writeLen(w, m.InterfaceProductInfo, 11)
    writeLen(w, m.ProtocolInfo, 11)
    writeLen(w, m.MacAddress, 21)
    writeLen(w, m.OneTimePassword, 41)
    writeLen(w, m.ClientIPAddress, 16)
    writeLen(w, m.LoginRemark, 36)
    binary.Write(w, binary.BigEndian, &m.ClientIPPort)
}

func (m* CThostFtdcReqUserLoginField) Unmarshal(r io.Reader) {
    m.TradingDay = readLen(r, 9)
    m.BrokerID = readLen(r, 11)
    m.UserID = readLen(r, 16)
    m.Password = readLen(r, 41)
    m.UserProductInfo = readLen(r, 11)
    m.InterfaceProductInfo = readLen(r, 11)
    m.ProtocolInfo = readLen(r, 11)
    m.MacAddress = readLen(r, 21)
    m.OneTimePassword = readLen(r, 41)
    m.ClientIPAddress = readLen(r, 16)
    m.LoginRemark = readLen(r, 36)
    binary.Read(r, binary.BigEndian, &m.ClientIPPort)
}

type  CThostFtdcRspUserLoginField struct {
    TradingDay string;
    LoginTime string;
    BrokerID string;
    UserID string;
    SystemName string;
    FrontID int32;
    SessionID int32;
    MaxOrderRef string;
    SHFETime string;
    DCETime string;
    CZCETime string;
    FFEXTime string;
    INETime string;
}

func (m* CThostFtdcRspUserLoginField) Marshal(w io.Writer) {
    writeLen(w, m.TradingDay, 9)
    writeLen(w, m.LoginTime, 9)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.SystemName, 41)
    binary.Write(w, binary.BigEndian, &m.FrontID)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.MaxOrderRef, 13)
    writeLen(w, m.SHFETime, 9)
    writeLen(w, m.DCETime, 9)
    writeLen(w, m.CZCETime, 9)
    writeLen(w, m.FFEXTime, 9)
    writeLen(w, m.INETime, 9)
}

func (m* CThostFtdcRspUserLoginField) Unmarshal(r io.Reader) {
    m.TradingDay = readLen(r, 9)
    m.LoginTime = readLen(r, 9)
    m.BrokerID = readLen(r, 11)
    m.UserID = readLen(r, 16)
    m.SystemName = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.FrontID)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.MaxOrderRef = readLen(r, 13)
    m.SHFETime = readLen(r, 9)
    m.DCETime = readLen(r, 9)
    m.CZCETime = readLen(r, 9)
    m.FFEXTime = readLen(r, 9)
    m.INETime = readLen(r, 9)
}

type  CThostFtdcUserLogoutField struct {
    BrokerID string;
    UserID string;
}

func (m* CThostFtdcUserLogoutField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.UserID, 16)
}

func (m* CThostFtdcUserLogoutField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.UserID = readLen(r, 16)
}

type  CThostFtdcForceUserLogoutField struct {
    BrokerID string;
    UserID string;
}

func (m* CThostFtdcForceUserLogoutField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.UserID, 16)
}

func (m* CThostFtdcForceUserLogoutField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.UserID = readLen(r, 16)
}

type  CThostFtdcReqAuthenticateField struct {
    BrokerID string;
    UserID string;
    UserProductInfo string;
    AuthCode string;
    AppID string;
}

func (m* CThostFtdcReqAuthenticateField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.UserProductInfo, 11)
    writeLen(w, m.AuthCode, 17)
    writeLen(w, m.AppID, 33)
}

func (m* CThostFtdcReqAuthenticateField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.UserID = readLen(r, 16)
    m.UserProductInfo = readLen(r, 11)
    m.AuthCode = readLen(r, 17)
    m.AppID = readLen(r, 33)
}

type  CThostFtdcRspAuthenticateField struct {
    BrokerID string;
    UserID string;
    UserProductInfo string;
    AppID string;
    AppType uint8;
}

func (m* CThostFtdcRspAuthenticateField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.UserProductInfo, 11)
    writeLen(w, m.AppID, 33)
    binary.Write(w, binary.BigEndian, &m.AppType)
}

func (m* CThostFtdcRspAuthenticateField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.UserID = readLen(r, 16)
    m.UserProductInfo = readLen(r, 11)
    m.AppID = readLen(r, 33)
    binary.Read(r, binary.BigEndian, &m.AppType)
}

type  CThostFtdcAuthenticationInfoField struct {
    BrokerID string;
    UserID string;
    UserProductInfo string;
    AuthInfo string;
    IsResult int32;
    AppID string;
    AppType uint8;
}

func (m* CThostFtdcAuthenticationInfoField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.UserProductInfo, 11)
    writeLen(w, m.AuthInfo, 129)
    binary.Write(w, binary.BigEndian, &m.IsResult)
    writeLen(w, m.AppID, 33)
    binary.Write(w, binary.BigEndian, &m.AppType)
}

func (m* CThostFtdcAuthenticationInfoField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.UserID = readLen(r, 16)
    m.UserProductInfo = readLen(r, 11)
    m.AuthInfo = readLen(r, 129)
    binary.Read(r, binary.BigEndian, &m.IsResult)
    m.AppID = readLen(r, 33)
    binary.Read(r, binary.BigEndian, &m.AppType)
}

type  CThostFtdcRspUserLoginSafelyField struct {
    TradingDay string;
    LoginTime string;
    BrokerID string;
    UserID string;
    SystemName string;
    FrontID int32;
    SessionID int32;
    MaxOrderRef string;
    SHFETime string;
    DCETime string;
    CZCETime string;
    FFEXTime string;
    INETime string;
    RandomString string;
}

func (m* CThostFtdcRspUserLoginSafelyField) Marshal(w io.Writer) {
    writeLen(w, m.TradingDay, 9)
    writeLen(w, m.LoginTime, 9)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.SystemName, 41)
    binary.Write(w, binary.BigEndian, &m.FrontID)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.MaxOrderRef, 13)
    writeLen(w, m.SHFETime, 9)
    writeLen(w, m.DCETime, 9)
    writeLen(w, m.CZCETime, 9)
    writeLen(w, m.FFEXTime, 9)
    writeLen(w, m.INETime, 9)
    writeLen(w, m.RandomString, 17)
}

func (m* CThostFtdcRspUserLoginSafelyField) Unmarshal(r io.Reader) {
    m.TradingDay = readLen(r, 9)
    m.LoginTime = readLen(r, 9)
    m.BrokerID = readLen(r, 11)
    m.UserID = readLen(r, 16)
    m.SystemName = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.FrontID)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.MaxOrderRef = readLen(r, 13)
    m.SHFETime = readLen(r, 9)
    m.DCETime = readLen(r, 9)
    m.CZCETime = readLen(r, 9)
    m.FFEXTime = readLen(r, 9)
    m.INETime = readLen(r, 9)
    m.RandomString = readLen(r, 17)
}

type  CThostFtdcTransferHeaderField struct {
    Version string;
    TradeCode string;
    TradeDate string;
    TradeTime string;
    TradeSerial string;
    FutureID string;
    BankID string;
    BankBrchID string;
    OperNo string;
    DeviceID string;
    RecordNum string;
    SessionID int32;
    RequestID int32;
}

func (m* CThostFtdcTransferHeaderField) Marshal(w io.Writer) {
    writeLen(w, m.Version, 4)
    writeLen(w, m.TradeCode, 7)
    writeLen(w, m.TradeDate, 9)
    writeLen(w, m.TradeTime, 9)
    writeLen(w, m.TradeSerial, 9)
    writeLen(w, m.FutureID, 11)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBrchID, 5)
    writeLen(w, m.OperNo, 17)
    writeLen(w, m.DeviceID, 3)
    writeLen(w, m.RecordNum, 7)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    binary.Write(w, binary.BigEndian, &m.RequestID)
}

func (m* CThostFtdcTransferHeaderField) Unmarshal(r io.Reader) {
    m.Version = readLen(r, 4)
    m.TradeCode = readLen(r, 7)
    m.TradeDate = readLen(r, 9)
    m.TradeTime = readLen(r, 9)
    m.TradeSerial = readLen(r, 9)
    m.FutureID = readLen(r, 11)
    m.BankID = readLen(r, 4)
    m.BankBrchID = readLen(r, 5)
    m.OperNo = readLen(r, 17)
    m.DeviceID = readLen(r, 3)
    m.RecordNum = readLen(r, 7)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    binary.Read(r, binary.BigEndian, &m.RequestID)
}

type  CThostFtdcTransferBankToFutureReqField struct {
    FutureAccount string;
    FuturePwdFlag uint8;
    FutureAccPwd string;
    TradeAmt float64;
    CustFee float64;
    CurrencyCode string;
}

func (m* CThostFtdcTransferBankToFutureReqField) Marshal(w io.Writer) {
    writeLen(w, m.FutureAccount, 13)
    binary.Write(w, binary.BigEndian, &m.FuturePwdFlag)
    writeLen(w, m.FutureAccPwd, 17)
    binary.Write(w, binary.BigEndian, &m.TradeAmt)
    binary.Write(w, binary.BigEndian, &m.CustFee)
    writeLen(w, m.CurrencyCode, 4)
}

func (m* CThostFtdcTransferBankToFutureReqField) Unmarshal(r io.Reader) {
    m.FutureAccount = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.FuturePwdFlag)
    m.FutureAccPwd = readLen(r, 17)
    binary.Read(r, binary.BigEndian, &m.TradeAmt)
    binary.Read(r, binary.BigEndian, &m.CustFee)
    m.CurrencyCode = readLen(r, 4)
}

type  CThostFtdcTransferBankToFutureRspField struct {
    RetCode string;
    RetInfo string;
    FutureAccount string;
    TradeAmt float64;
    CustFee float64;
    CurrencyCode string;
}

func (m* CThostFtdcTransferBankToFutureRspField) Marshal(w io.Writer) {
    writeLen(w, m.RetCode, 5)
    writeLen(w, m.RetInfo, 129)
    writeLen(w, m.FutureAccount, 13)
    binary.Write(w, binary.BigEndian, &m.TradeAmt)
    binary.Write(w, binary.BigEndian, &m.CustFee)
    writeLen(w, m.CurrencyCode, 4)
}

func (m* CThostFtdcTransferBankToFutureRspField) Unmarshal(r io.Reader) {
    m.RetCode = readLen(r, 5)
    m.RetInfo = readLen(r, 129)
    m.FutureAccount = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.TradeAmt)
    binary.Read(r, binary.BigEndian, &m.CustFee)
    m.CurrencyCode = readLen(r, 4)
}

type  CThostFtdcTransferFutureToBankReqField struct {
    FutureAccount string;
    FuturePwdFlag uint8;
    FutureAccPwd string;
    TradeAmt float64;
    CustFee float64;
    CurrencyCode string;
}

func (m* CThostFtdcTransferFutureToBankReqField) Marshal(w io.Writer) {
    writeLen(w, m.FutureAccount, 13)
    binary.Write(w, binary.BigEndian, &m.FuturePwdFlag)
    writeLen(w, m.FutureAccPwd, 17)
    binary.Write(w, binary.BigEndian, &m.TradeAmt)
    binary.Write(w, binary.BigEndian, &m.CustFee)
    writeLen(w, m.CurrencyCode, 4)
}

func (m* CThostFtdcTransferFutureToBankReqField) Unmarshal(r io.Reader) {
    m.FutureAccount = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.FuturePwdFlag)
    m.FutureAccPwd = readLen(r, 17)
    binary.Read(r, binary.BigEndian, &m.TradeAmt)
    binary.Read(r, binary.BigEndian, &m.CustFee)
    m.CurrencyCode = readLen(r, 4)
}

type  CThostFtdcTransferFutureToBankRspField struct {
    RetCode string;
    RetInfo string;
    FutureAccount string;
    TradeAmt float64;
    CustFee float64;
    CurrencyCode string;
}

func (m* CThostFtdcTransferFutureToBankRspField) Marshal(w io.Writer) {
    writeLen(w, m.RetCode, 5)
    writeLen(w, m.RetInfo, 129)
    writeLen(w, m.FutureAccount, 13)
    binary.Write(w, binary.BigEndian, &m.TradeAmt)
    binary.Write(w, binary.BigEndian, &m.CustFee)
    writeLen(w, m.CurrencyCode, 4)
}

func (m* CThostFtdcTransferFutureToBankRspField) Unmarshal(r io.Reader) {
    m.RetCode = readLen(r, 5)
    m.RetInfo = readLen(r, 129)
    m.FutureAccount = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.TradeAmt)
    binary.Read(r, binary.BigEndian, &m.CustFee)
    m.CurrencyCode = readLen(r, 4)
}

type  CThostFtdcTransferQryBankReqField struct {
    FutureAccount string;
    FuturePwdFlag uint8;
    FutureAccPwd string;
    CurrencyCode string;
}

func (m* CThostFtdcTransferQryBankReqField) Marshal(w io.Writer) {
    writeLen(w, m.FutureAccount, 13)
    binary.Write(w, binary.BigEndian, &m.FuturePwdFlag)
    writeLen(w, m.FutureAccPwd, 17)
    writeLen(w, m.CurrencyCode, 4)
}

func (m* CThostFtdcTransferQryBankReqField) Unmarshal(r io.Reader) {
    m.FutureAccount = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.FuturePwdFlag)
    m.FutureAccPwd = readLen(r, 17)
    m.CurrencyCode = readLen(r, 4)
}

type  CThostFtdcTransferQryBankRspField struct {
    RetCode string;
    RetInfo string;
    FutureAccount string;
    TradeAmt float64;
    UseAmt float64;
    FetchAmt float64;
    CurrencyCode string;
}

func (m* CThostFtdcTransferQryBankRspField) Marshal(w io.Writer) {
    writeLen(w, m.RetCode, 5)
    writeLen(w, m.RetInfo, 129)
    writeLen(w, m.FutureAccount, 13)
    binary.Write(w, binary.BigEndian, &m.TradeAmt)
    binary.Write(w, binary.BigEndian, &m.UseAmt)
    binary.Write(w, binary.BigEndian, &m.FetchAmt)
    writeLen(w, m.CurrencyCode, 4)
}

func (m* CThostFtdcTransferQryBankRspField) Unmarshal(r io.Reader) {
    m.RetCode = readLen(r, 5)
    m.RetInfo = readLen(r, 129)
    m.FutureAccount = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.TradeAmt)
    binary.Read(r, binary.BigEndian, &m.UseAmt)
    binary.Read(r, binary.BigEndian, &m.FetchAmt)
    m.CurrencyCode = readLen(r, 4)
}

type  CThostFtdcTransferQryDetailReqField struct {
    FutureAccount string;
}

func (m* CThostFtdcTransferQryDetailReqField) Marshal(w io.Writer) {
    writeLen(w, m.FutureAccount, 13)
}

func (m* CThostFtdcTransferQryDetailReqField) Unmarshal(r io.Reader) {
    m.FutureAccount = readLen(r, 13)
}

type  CThostFtdcTransferQryDetailRspField struct {
    TradeDate string;
    TradeTime string;
    TradeCode string;
    FutureSerial int32;
    FutureID string;
    FutureAccount string;
    BankSerial int32;
    BankID string;
    BankBrchID string;
    BankAccount string;
    CertCode string;
    CurrencyCode string;
    TxAmount float64;
    Flag uint8;
}

func (m* CThostFtdcTransferQryDetailRspField) Marshal(w io.Writer) {
    writeLen(w, m.TradeDate, 9)
    writeLen(w, m.TradeTime, 9)
    writeLen(w, m.TradeCode, 7)
    binary.Write(w, binary.BigEndian, &m.FutureSerial)
    writeLen(w, m.FutureID, 11)
    writeLen(w, m.FutureAccount, 22)
    binary.Write(w, binary.BigEndian, &m.BankSerial)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBrchID, 5)
    writeLen(w, m.BankAccount, 41)
    writeLen(w, m.CertCode, 21)
    writeLen(w, m.CurrencyCode, 4)
    binary.Write(w, binary.BigEndian, &m.TxAmount)
    binary.Write(w, binary.BigEndian, &m.Flag)
}

func (m* CThostFtdcTransferQryDetailRspField) Unmarshal(r io.Reader) {
    m.TradeDate = readLen(r, 9)
    m.TradeTime = readLen(r, 9)
    m.TradeCode = readLen(r, 7)
    binary.Read(r, binary.BigEndian, &m.FutureSerial)
    m.FutureID = readLen(r, 11)
    m.FutureAccount = readLen(r, 22)
    binary.Read(r, binary.BigEndian, &m.BankSerial)
    m.BankID = readLen(r, 4)
    m.BankBrchID = readLen(r, 5)
    m.BankAccount = readLen(r, 41)
    m.CertCode = readLen(r, 21)
    m.CurrencyCode = readLen(r, 4)
    binary.Read(r, binary.BigEndian, &m.TxAmount)
    binary.Read(r, binary.BigEndian, &m.Flag)
}

type  CThostFtdcRspInfoField struct {
    ErrorID int32;
    ErrorMsg string;
}

func (m* CThostFtdcRspInfoField) Marshal(w io.Writer) {
    binary.Write(w, binary.BigEndian, &m.ErrorID)
    writeLen(w, m.ErrorMsg, 81)
}

func (m* CThostFtdcRspInfoField) Unmarshal(r io.Reader) {
    binary.Read(r, binary.BigEndian, &m.ErrorID)
    m.ErrorMsg = readLen(r, 81)
}

type  CThostFtdcExchangeField struct {
    ExchangeID string;
    ExchangeName string;
    ExchangeProperty uint8;
}

func (m* CThostFtdcExchangeField) Marshal(w io.Writer) {
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ExchangeName, 61)
    binary.Write(w, binary.BigEndian, &m.ExchangeProperty)
}

func (m* CThostFtdcExchangeField) Unmarshal(r io.Reader) {
    m.ExchangeID = readLen(r, 9)
    m.ExchangeName = readLen(r, 61)
    binary.Read(r, binary.BigEndian, &m.ExchangeProperty)
}

type  CThostFtdcProductField struct {
    ProductID string;
    ProductName string;
    ExchangeID string;
    ProductClass uint8;
    VolumeMultiple int32;
    PriceTick float64;
    MaxMarketOrderVolume int32;
    MinMarketOrderVolume int32;
    MaxLimitOrderVolume int32;
    MinLimitOrderVolume int32;
    PositionType uint8;
    PositionDateType uint8;
    CloseDealType uint8;
    TradeCurrencyID string;
    MortgageFundUseRange uint8;
    ExchangeProductID string;
    UnderlyingMultiple float64;
}

func (m* CThostFtdcProductField) Marshal(w io.Writer) {
    writeLen(w, m.ProductID, 31)
    writeLen(w, m.ProductName, 21)
    writeLen(w, m.ExchangeID, 9)
    binary.Write(w, binary.BigEndian, &m.ProductClass)
    binary.Write(w, binary.BigEndian, &m.VolumeMultiple)
    binary.Write(w, binary.BigEndian, &m.PriceTick)
    binary.Write(w, binary.BigEndian, &m.MaxMarketOrderVolume)
    binary.Write(w, binary.BigEndian, &m.MinMarketOrderVolume)
    binary.Write(w, binary.BigEndian, &m.MaxLimitOrderVolume)
    binary.Write(w, binary.BigEndian, &m.MinLimitOrderVolume)
    binary.Write(w, binary.BigEndian, &m.PositionType)
    binary.Write(w, binary.BigEndian, &m.PositionDateType)
    binary.Write(w, binary.BigEndian, &m.CloseDealType)
    writeLen(w, m.TradeCurrencyID, 4)
    binary.Write(w, binary.BigEndian, &m.MortgageFundUseRange)
    writeLen(w, m.ExchangeProductID, 31)
    binary.Write(w, binary.BigEndian, &m.UnderlyingMultiple)
}

func (m* CThostFtdcProductField) Unmarshal(r io.Reader) {
    m.ProductID = readLen(r, 31)
    m.ProductName = readLen(r, 21)
    m.ExchangeID = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.ProductClass)
    binary.Read(r, binary.BigEndian, &m.VolumeMultiple)
    binary.Read(r, binary.BigEndian, &m.PriceTick)
    binary.Read(r, binary.BigEndian, &m.MaxMarketOrderVolume)
    binary.Read(r, binary.BigEndian, &m.MinMarketOrderVolume)
    binary.Read(r, binary.BigEndian, &m.MaxLimitOrderVolume)
    binary.Read(r, binary.BigEndian, &m.MinLimitOrderVolume)
    binary.Read(r, binary.BigEndian, &m.PositionType)
    binary.Read(r, binary.BigEndian, &m.PositionDateType)
    binary.Read(r, binary.BigEndian, &m.CloseDealType)
    m.TradeCurrencyID = readLen(r, 4)
    binary.Read(r, binary.BigEndian, &m.MortgageFundUseRange)
    m.ExchangeProductID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.UnderlyingMultiple)
}

type  CThostFtdcInstrumentField struct {
    InstrumentID string;
    ExchangeID string;
    InstrumentName string;
    ExchangeInstID string;
    ProductID string;
    ProductClass uint8;
    DeliveryYear int32;
    DeliveryMonth int32;
    MaxMarketOrderVolume int32;
    MinMarketOrderVolume int32;
    MaxLimitOrderVolume int32;
    MinLimitOrderVolume int32;
    VolumeMultiple int32;
    PriceTick float64;
    CreateDate string;
    OpenDate string;
    ExpireDate string;
    StartDelivDate string;
    EndDelivDate string;
    InstLifePhase uint8;
    IsTrading int32;
    PositionType uint8;
    PositionDateType uint8;
    LongMarginRatio float64;
    ShortMarginRatio float64;
    MaxMarginSideAlgorithm uint8;
    UnderlyingInstrID string;
    StrikePrice float64;
    OptionsType uint8;
    UnderlyingMultiple float64;
    CombinationType uint8;
    MinBuyVolume int32;
    MinSellVolume int32;
    InstrumentCode string;
}

func (m* CThostFtdcInstrumentField) Marshal(w io.Writer) {
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.InstrumentName, 21)
    writeLen(w, m.ExchangeInstID, 31)
    writeLen(w, m.ProductID, 31)
    binary.Write(w, binary.BigEndian, &m.ProductClass)
    binary.Write(w, binary.BigEndian, &m.DeliveryYear)
    binary.Write(w, binary.BigEndian, &m.DeliveryMonth)
    binary.Write(w, binary.BigEndian, &m.MaxMarketOrderVolume)
    binary.Write(w, binary.BigEndian, &m.MinMarketOrderVolume)
    binary.Write(w, binary.BigEndian, &m.MaxLimitOrderVolume)
    binary.Write(w, binary.BigEndian, &m.MinLimitOrderVolume)
    binary.Write(w, binary.BigEndian, &m.VolumeMultiple)
    binary.Write(w, binary.BigEndian, &m.PriceTick)
    writeLen(w, m.CreateDate, 9)
    writeLen(w, m.OpenDate, 9)
    writeLen(w, m.ExpireDate, 9)
    writeLen(w, m.StartDelivDate, 9)
    writeLen(w, m.EndDelivDate, 9)
    binary.Write(w, binary.BigEndian, &m.InstLifePhase)
    binary.Write(w, binary.BigEndian, &m.IsTrading)
    binary.Write(w, binary.BigEndian, &m.PositionType)
    binary.Write(w, binary.BigEndian, &m.PositionDateType)
    binary.Write(w, binary.BigEndian, &m.LongMarginRatio)
    binary.Write(w, binary.BigEndian, &m.ShortMarginRatio)
    binary.Write(w, binary.BigEndian, &m.MaxMarginSideAlgorithm)
    writeLen(w, m.UnderlyingInstrID, 31)
    binary.Write(w, binary.BigEndian, &m.StrikePrice)
    binary.Write(w, binary.BigEndian, &m.OptionsType)
    binary.Write(w, binary.BigEndian, &m.UnderlyingMultiple)
    binary.Write(w, binary.BigEndian, &m.CombinationType)
    binary.Write(w, binary.BigEndian, &m.MinBuyVolume)
    binary.Write(w, binary.BigEndian, &m.MinSellVolume)
    writeLen(w, m.InstrumentCode, 31)
}

func (m* CThostFtdcInstrumentField) Unmarshal(r io.Reader) {
    m.InstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
    m.InstrumentName = readLen(r, 21)
    m.ExchangeInstID = readLen(r, 31)
    m.ProductID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.ProductClass)
    binary.Read(r, binary.BigEndian, &m.DeliveryYear)
    binary.Read(r, binary.BigEndian, &m.DeliveryMonth)
    binary.Read(r, binary.BigEndian, &m.MaxMarketOrderVolume)
    binary.Read(r, binary.BigEndian, &m.MinMarketOrderVolume)
    binary.Read(r, binary.BigEndian, &m.MaxLimitOrderVolume)
    binary.Read(r, binary.BigEndian, &m.MinLimitOrderVolume)
    binary.Read(r, binary.BigEndian, &m.VolumeMultiple)
    binary.Read(r, binary.BigEndian, &m.PriceTick)
    m.CreateDate = readLen(r, 9)
    m.OpenDate = readLen(r, 9)
    m.ExpireDate = readLen(r, 9)
    m.StartDelivDate = readLen(r, 9)
    m.EndDelivDate = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.InstLifePhase)
    binary.Read(r, binary.BigEndian, &m.IsTrading)
    binary.Read(r, binary.BigEndian, &m.PositionType)
    binary.Read(r, binary.BigEndian, &m.PositionDateType)
    binary.Read(r, binary.BigEndian, &m.LongMarginRatio)
    binary.Read(r, binary.BigEndian, &m.ShortMarginRatio)
    binary.Read(r, binary.BigEndian, &m.MaxMarginSideAlgorithm)
    m.UnderlyingInstrID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.StrikePrice)
    binary.Read(r, binary.BigEndian, &m.OptionsType)
    binary.Read(r, binary.BigEndian, &m.UnderlyingMultiple)
    binary.Read(r, binary.BigEndian, &m.CombinationType)
    binary.Read(r, binary.BigEndian, &m.MinBuyVolume)
    binary.Read(r, binary.BigEndian, &m.MinSellVolume)
    m.InstrumentCode = readLen(r, 31)
}

type  CThostFtdcBrokerField struct {
    BrokerID string;
    BrokerAbbr string;
    BrokerName string;
    IsActive int32;
}

func (m* CThostFtdcBrokerField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.BrokerAbbr, 9)
    writeLen(w, m.BrokerName, 81)
    binary.Write(w, binary.BigEndian, &m.IsActive)
}

func (m* CThostFtdcBrokerField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.BrokerAbbr = readLen(r, 9)
    m.BrokerName = readLen(r, 81)
    binary.Read(r, binary.BigEndian, &m.IsActive)
}

type  CThostFtdcTraderField struct {
    ExchangeID string;
    TraderID string;
    ParticipantID string;
    Password string;
    InstallCount int32;
    BrokerID string;
}

func (m* CThostFtdcTraderField) Marshal(w io.Writer) {
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.TraderID, 21)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.Password, 41)
    binary.Write(w, binary.BigEndian, &m.InstallCount)
    writeLen(w, m.BrokerID, 11)
}

func (m* CThostFtdcTraderField) Unmarshal(r io.Reader) {
    m.ExchangeID = readLen(r, 9)
    m.TraderID = readLen(r, 21)
    m.ParticipantID = readLen(r, 11)
    m.Password = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.InstallCount)
    m.BrokerID = readLen(r, 11)
}

type  CThostFtdcInvestorField struct {
    InvestorID string;
    BrokerID string;
    InvestorGroupID string;
    InvestorName string;
    IdentifiedCardType uint8;
    IdentifiedCardNo string;
    IsActive int32;
    Telephone string;
    Address string;
    OpenDate string;
    Mobile string;
    CommModelID string;
    MarginModelID string;
}

func (m* CThostFtdcInvestorField) Marshal(w io.Writer) {
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorGroupID, 13)
    writeLen(w, m.InvestorName, 81)
    binary.Write(w, binary.BigEndian, &m.IdentifiedCardType)
    writeLen(w, m.IdentifiedCardNo, 51)
    binary.Write(w, binary.BigEndian, &m.IsActive)
    writeLen(w, m.Telephone, 41)
    writeLen(w, m.Address, 101)
    writeLen(w, m.OpenDate, 9)
    writeLen(w, m.Mobile, 41)
    writeLen(w, m.CommModelID, 13)
    writeLen(w, m.MarginModelID, 13)
}

func (m* CThostFtdcInvestorField) Unmarshal(r io.Reader) {
    m.InvestorID = readLen(r, 13)
    m.BrokerID = readLen(r, 11)
    m.InvestorGroupID = readLen(r, 13)
    m.InvestorName = readLen(r, 81)
    binary.Read(r, binary.BigEndian, &m.IdentifiedCardType)
    m.IdentifiedCardNo = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.IsActive)
    m.Telephone = readLen(r, 41)
    m.Address = readLen(r, 101)
    m.OpenDate = readLen(r, 9)
    m.Mobile = readLen(r, 41)
    m.CommModelID = readLen(r, 13)
    m.MarginModelID = readLen(r, 13)
}

type  CThostFtdcTradingCodeField struct {
    InvestorID string;
    BrokerID string;
    ExchangeID string;
    ClientID string;
    IsActive int32;
    ClientIDType string;
    BranchID string;
    BizType uint8;
    InvestUnitID string;
}

func (m* CThostFtdcTradingCodeField) Marshal(w io.Writer) {
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ClientID, 11)
    binary.Write(w, binary.BigEndian, &m.IsActive)
    writeLen(w, m.ClientIDType, 11)
    writeLen(w, m.BranchID, 9)
    binary.Write(w, binary.BigEndian, &m.BizType)
    writeLen(w, m.InvestUnitID, 17)
}

func (m* CThostFtdcTradingCodeField) Unmarshal(r io.Reader) {
    m.InvestorID = readLen(r, 13)
    m.BrokerID = readLen(r, 11)
    m.ExchangeID = readLen(r, 9)
    m.ClientID = readLen(r, 11)
    binary.Read(r, binary.BigEndian, &m.IsActive)
    m.ClientIDType = readLen(r, 11)
    m.BranchID = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.BizType)
    m.InvestUnitID = readLen(r, 17)
}

type  CThostFtdcPartBrokerField struct {
    BrokerID string;
    ExchangeID string;
    ParticipantID string;
    IsActive int32;
}

func (m* CThostFtdcPartBrokerField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ParticipantID, 11)
    binary.Write(w, binary.BigEndian, &m.IsActive)
}

func (m* CThostFtdcPartBrokerField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.ExchangeID = readLen(r, 9)
    m.ParticipantID = readLen(r, 11)
    binary.Read(r, binary.BigEndian, &m.IsActive)
}

type  CThostFtdcSuperUserField struct {
    UserID string;
    UserName string;
    Password string;
    IsActive int32;
}

func (m* CThostFtdcSuperUserField) Marshal(w io.Writer) {
    writeLen(w, m.UserID, 16)
    writeLen(w, m.UserName, 81)
    writeLen(w, m.Password, 41)
    binary.Write(w, binary.BigEndian, &m.IsActive)
}

func (m* CThostFtdcSuperUserField) Unmarshal(r io.Reader) {
    m.UserID = readLen(r, 16)
    m.UserName = readLen(r, 81)
    m.Password = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.IsActive)
}

type  CThostFtdcSuperUserFunctionField struct {
    UserID string;
    FunctionCode uint8;
}

func (m* CThostFtdcSuperUserFunctionField) Marshal(w io.Writer) {
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.FunctionCode)
}

func (m* CThostFtdcSuperUserFunctionField) Unmarshal(r io.Reader) {
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.FunctionCode)
}

type  CThostFtdcInvestorGroupField struct {
    BrokerID string;
    InvestorGroupID string;
    InvestorGroupName string;
}

func (m* CThostFtdcInvestorGroupField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorGroupID, 13)
    writeLen(w, m.InvestorGroupName, 41)
}

func (m* CThostFtdcInvestorGroupField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorGroupID = readLen(r, 13)
    m.InvestorGroupName = readLen(r, 41)
}

type  CThostFtdcTradingAccountField struct {
    BrokerID string;
    AccountID string;
    PreMortgage float64;
    PreCredit float64;
    PreDeposit float64;
    PreBalance float64;
    PreMargin float64;
    InterestBase float64;
    Interest float64;
    Deposit float64;
    Withdraw float64;
    FrozenMargin float64;
    FrozenCash float64;
    FrozenCommission float64;
    CurrMargin float64;
    CashIn float64;
    Commission float64;
    CloseProfit float64;
    PositionProfit float64;
    Balance float64;
    Available float64;
    WithdrawQuota float64;
    Reserve float64;
    TradingDay string;
    SettlementID int32;
    Credit float64;
    Mortgage float64;
    ExchangeMargin float64;
    DeliveryMargin float64;
    ExchangeDeliveryMargin float64;
    ReserveBalance float64;
    CurrencyID string;
    PreFundMortgageIn float64;
    PreFundMortgageOut float64;
    FundMortgageIn float64;
    FundMortgageOut float64;
    FundMortgageAvailable float64;
    MortgageableFund float64;
    SpecProductMargin float64;
    SpecProductFrozenMargin float64;
    SpecProductCommission float64;
    SpecProductFrozenCommission float64;
    SpecProductPositionProfit float64;
    SpecProductCloseProfit float64;
    SpecProductPositionProfitByAlg float64;
    SpecProductExchangeMargin float64;
    BizType uint8;
    FrozenSwap float64;
    RemainSwap float64;
}

func (m* CThostFtdcTradingAccountField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.AccountID, 13)
    binary.Write(w, binary.BigEndian, &m.PreMortgage)
    binary.Write(w, binary.BigEndian, &m.PreCredit)
    binary.Write(w, binary.BigEndian, &m.PreDeposit)
    binary.Write(w, binary.BigEndian, &m.PreBalance)
    binary.Write(w, binary.BigEndian, &m.PreMargin)
    binary.Write(w, binary.BigEndian, &m.InterestBase)
    binary.Write(w, binary.BigEndian, &m.Interest)
    binary.Write(w, binary.BigEndian, &m.Deposit)
    binary.Write(w, binary.BigEndian, &m.Withdraw)
    binary.Write(w, binary.BigEndian, &m.FrozenMargin)
    binary.Write(w, binary.BigEndian, &m.FrozenCash)
    binary.Write(w, binary.BigEndian, &m.FrozenCommission)
    binary.Write(w, binary.BigEndian, &m.CurrMargin)
    binary.Write(w, binary.BigEndian, &m.CashIn)
    binary.Write(w, binary.BigEndian, &m.Commission)
    binary.Write(w, binary.BigEndian, &m.CloseProfit)
    binary.Write(w, binary.BigEndian, &m.PositionProfit)
    binary.Write(w, binary.BigEndian, &m.Balance)
    binary.Write(w, binary.BigEndian, &m.Available)
    binary.Write(w, binary.BigEndian, &m.WithdrawQuota)
    binary.Write(w, binary.BigEndian, &m.Reserve)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.SettlementID)
    binary.Write(w, binary.BigEndian, &m.Credit)
    binary.Write(w, binary.BigEndian, &m.Mortgage)
    binary.Write(w, binary.BigEndian, &m.ExchangeMargin)
    binary.Write(w, binary.BigEndian, &m.DeliveryMargin)
    binary.Write(w, binary.BigEndian, &m.ExchangeDeliveryMargin)
    binary.Write(w, binary.BigEndian, &m.ReserveBalance)
    writeLen(w, m.CurrencyID, 4)
    binary.Write(w, binary.BigEndian, &m.PreFundMortgageIn)
    binary.Write(w, binary.BigEndian, &m.PreFundMortgageOut)
    binary.Write(w, binary.BigEndian, &m.FundMortgageIn)
    binary.Write(w, binary.BigEndian, &m.FundMortgageOut)
    binary.Write(w, binary.BigEndian, &m.FundMortgageAvailable)
    binary.Write(w, binary.BigEndian, &m.MortgageableFund)
    binary.Write(w, binary.BigEndian, &m.SpecProductMargin)
    binary.Write(w, binary.BigEndian, &m.SpecProductFrozenMargin)
    binary.Write(w, binary.BigEndian, &m.SpecProductCommission)
    binary.Write(w, binary.BigEndian, &m.SpecProductFrozenCommission)
    binary.Write(w, binary.BigEndian, &m.SpecProductPositionProfit)
    binary.Write(w, binary.BigEndian, &m.SpecProductCloseProfit)
    binary.Write(w, binary.BigEndian, &m.SpecProductPositionProfitByAlg)
    binary.Write(w, binary.BigEndian, &m.SpecProductExchangeMargin)
    binary.Write(w, binary.BigEndian, &m.BizType)
    binary.Write(w, binary.BigEndian, &m.FrozenSwap)
    binary.Write(w, binary.BigEndian, &m.RemainSwap)
}

func (m* CThostFtdcTradingAccountField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.AccountID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.PreMortgage)
    binary.Read(r, binary.BigEndian, &m.PreCredit)
    binary.Read(r, binary.BigEndian, &m.PreDeposit)
    binary.Read(r, binary.BigEndian, &m.PreBalance)
    binary.Read(r, binary.BigEndian, &m.PreMargin)
    binary.Read(r, binary.BigEndian, &m.InterestBase)
    binary.Read(r, binary.BigEndian, &m.Interest)
    binary.Read(r, binary.BigEndian, &m.Deposit)
    binary.Read(r, binary.BigEndian, &m.Withdraw)
    binary.Read(r, binary.BigEndian, &m.FrozenMargin)
    binary.Read(r, binary.BigEndian, &m.FrozenCash)
    binary.Read(r, binary.BigEndian, &m.FrozenCommission)
    binary.Read(r, binary.BigEndian, &m.CurrMargin)
    binary.Read(r, binary.BigEndian, &m.CashIn)
    binary.Read(r, binary.BigEndian, &m.Commission)
    binary.Read(r, binary.BigEndian, &m.CloseProfit)
    binary.Read(r, binary.BigEndian, &m.PositionProfit)
    binary.Read(r, binary.BigEndian, &m.Balance)
    binary.Read(r, binary.BigEndian, &m.Available)
    binary.Read(r, binary.BigEndian, &m.WithdrawQuota)
    binary.Read(r, binary.BigEndian, &m.Reserve)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.SettlementID)
    binary.Read(r, binary.BigEndian, &m.Credit)
    binary.Read(r, binary.BigEndian, &m.Mortgage)
    binary.Read(r, binary.BigEndian, &m.ExchangeMargin)
    binary.Read(r, binary.BigEndian, &m.DeliveryMargin)
    binary.Read(r, binary.BigEndian, &m.ExchangeDeliveryMargin)
    binary.Read(r, binary.BigEndian, &m.ReserveBalance)
    m.CurrencyID = readLen(r, 4)
    binary.Read(r, binary.BigEndian, &m.PreFundMortgageIn)
    binary.Read(r, binary.BigEndian, &m.PreFundMortgageOut)
    binary.Read(r, binary.BigEndian, &m.FundMortgageIn)
    binary.Read(r, binary.BigEndian, &m.FundMortgageOut)
    binary.Read(r, binary.BigEndian, &m.FundMortgageAvailable)
    binary.Read(r, binary.BigEndian, &m.MortgageableFund)
    binary.Read(r, binary.BigEndian, &m.SpecProductMargin)
    binary.Read(r, binary.BigEndian, &m.SpecProductFrozenMargin)
    binary.Read(r, binary.BigEndian, &m.SpecProductCommission)
    binary.Read(r, binary.BigEndian, &m.SpecProductFrozenCommission)
    binary.Read(r, binary.BigEndian, &m.SpecProductPositionProfit)
    binary.Read(r, binary.BigEndian, &m.SpecProductCloseProfit)
    binary.Read(r, binary.BigEndian, &m.SpecProductPositionProfitByAlg)
    binary.Read(r, binary.BigEndian, &m.SpecProductExchangeMargin)
    binary.Read(r, binary.BigEndian, &m.BizType)
    binary.Read(r, binary.BigEndian, &m.FrozenSwap)
    binary.Read(r, binary.BigEndian, &m.RemainSwap)
}

type  CThostFtdcInvestorPositionField struct {
    InstrumentID string;
    BrokerID string;
    InvestorID string;
    PosiDirection uint8;
    HedgeFlag uint8;
    PositionDate uint8;
    YdPosition int32;
    Position int32;
    LongFrozen int32;
    ShortFrozen int32;
    LongFrozenAmount float64;
    ShortFrozenAmount float64;
    OpenVolume int32;
    CloseVolume int32;
    OpenAmount float64;
    CloseAmount float64;
    PositionCost float64;
    PreMargin float64;
    UseMargin float64;
    FrozenMargin float64;
    FrozenCash float64;
    FrozenCommission float64;
    CashIn float64;
    Commission float64;
    CloseProfit float64;
    PositionProfit float64;
    PreSettlementPrice float64;
    SettlementPrice float64;
    TradingDay string;
    SettlementID int32;
    OpenCost float64;
    ExchangeMargin float64;
    CombPosition int32;
    CombLongFrozen int32;
    CombShortFrozen int32;
    CloseProfitByDate float64;
    CloseProfitByTrade float64;
    TodayPosition int32;
    MarginRateByMoney float64;
    MarginRateByVolume float64;
    StrikeFrozen int32;
    StrikeFrozenAmount float64;
    AbandonFrozen int32;
    ExchangeID string;
    YdStrikeFrozen int32;
    InvestUnitID string;
}

func (m* CThostFtdcInvestorPositionField) Marshal(w io.Writer) {
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.PosiDirection)
    binary.Write(w, binary.BigEndian, &m.HedgeFlag)
    binary.Write(w, binary.BigEndian, &m.PositionDate)
    binary.Write(w, binary.BigEndian, &m.YdPosition)
    binary.Write(w, binary.BigEndian, &m.Position)
    binary.Write(w, binary.BigEndian, &m.LongFrozen)
    binary.Write(w, binary.BigEndian, &m.ShortFrozen)
    binary.Write(w, binary.BigEndian, &m.LongFrozenAmount)
    binary.Write(w, binary.BigEndian, &m.ShortFrozenAmount)
    binary.Write(w, binary.BigEndian, &m.OpenVolume)
    binary.Write(w, binary.BigEndian, &m.CloseVolume)
    binary.Write(w, binary.BigEndian, &m.OpenAmount)
    binary.Write(w, binary.BigEndian, &m.CloseAmount)
    binary.Write(w, binary.BigEndian, &m.PositionCost)
    binary.Write(w, binary.BigEndian, &m.PreMargin)
    binary.Write(w, binary.BigEndian, &m.UseMargin)
    binary.Write(w, binary.BigEndian, &m.FrozenMargin)
    binary.Write(w, binary.BigEndian, &m.FrozenCash)
    binary.Write(w, binary.BigEndian, &m.FrozenCommission)
    binary.Write(w, binary.BigEndian, &m.CashIn)
    binary.Write(w, binary.BigEndian, &m.Commission)
    binary.Write(w, binary.BigEndian, &m.CloseProfit)
    binary.Write(w, binary.BigEndian, &m.PositionProfit)
    binary.Write(w, binary.BigEndian, &m.PreSettlementPrice)
    binary.Write(w, binary.BigEndian, &m.SettlementPrice)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.SettlementID)
    binary.Write(w, binary.BigEndian, &m.OpenCost)
    binary.Write(w, binary.BigEndian, &m.ExchangeMargin)
    binary.Write(w, binary.BigEndian, &m.CombPosition)
    binary.Write(w, binary.BigEndian, &m.CombLongFrozen)
    binary.Write(w, binary.BigEndian, &m.CombShortFrozen)
    binary.Write(w, binary.BigEndian, &m.CloseProfitByDate)
    binary.Write(w, binary.BigEndian, &m.CloseProfitByTrade)
    binary.Write(w, binary.BigEndian, &m.TodayPosition)
    binary.Write(w, binary.BigEndian, &m.MarginRateByMoney)
    binary.Write(w, binary.BigEndian, &m.MarginRateByVolume)
    binary.Write(w, binary.BigEndian, &m.StrikeFrozen)
    binary.Write(w, binary.BigEndian, &m.StrikeFrozenAmount)
    binary.Write(w, binary.BigEndian, &m.AbandonFrozen)
    writeLen(w, m.ExchangeID, 9)
    binary.Write(w, binary.BigEndian, &m.YdStrikeFrozen)
    writeLen(w, m.InvestUnitID, 17)
}

func (m* CThostFtdcInvestorPositionField) Unmarshal(r io.Reader) {
    m.InstrumentID = readLen(r, 31)
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.PosiDirection)
    binary.Read(r, binary.BigEndian, &m.HedgeFlag)
    binary.Read(r, binary.BigEndian, &m.PositionDate)
    binary.Read(r, binary.BigEndian, &m.YdPosition)
    binary.Read(r, binary.BigEndian, &m.Position)
    binary.Read(r, binary.BigEndian, &m.LongFrozen)
    binary.Read(r, binary.BigEndian, &m.ShortFrozen)
    binary.Read(r, binary.BigEndian, &m.LongFrozenAmount)
    binary.Read(r, binary.BigEndian, &m.ShortFrozenAmount)
    binary.Read(r, binary.BigEndian, &m.OpenVolume)
    binary.Read(r, binary.BigEndian, &m.CloseVolume)
    binary.Read(r, binary.BigEndian, &m.OpenAmount)
    binary.Read(r, binary.BigEndian, &m.CloseAmount)
    binary.Read(r, binary.BigEndian, &m.PositionCost)
    binary.Read(r, binary.BigEndian, &m.PreMargin)
    binary.Read(r, binary.BigEndian, &m.UseMargin)
    binary.Read(r, binary.BigEndian, &m.FrozenMargin)
    binary.Read(r, binary.BigEndian, &m.FrozenCash)
    binary.Read(r, binary.BigEndian, &m.FrozenCommission)
    binary.Read(r, binary.BigEndian, &m.CashIn)
    binary.Read(r, binary.BigEndian, &m.Commission)
    binary.Read(r, binary.BigEndian, &m.CloseProfit)
    binary.Read(r, binary.BigEndian, &m.PositionProfit)
    binary.Read(r, binary.BigEndian, &m.PreSettlementPrice)
    binary.Read(r, binary.BigEndian, &m.SettlementPrice)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.SettlementID)
    binary.Read(r, binary.BigEndian, &m.OpenCost)
    binary.Read(r, binary.BigEndian, &m.ExchangeMargin)
    binary.Read(r, binary.BigEndian, &m.CombPosition)
    binary.Read(r, binary.BigEndian, &m.CombLongFrozen)
    binary.Read(r, binary.BigEndian, &m.CombShortFrozen)
    binary.Read(r, binary.BigEndian, &m.CloseProfitByDate)
    binary.Read(r, binary.BigEndian, &m.CloseProfitByTrade)
    binary.Read(r, binary.BigEndian, &m.TodayPosition)
    binary.Read(r, binary.BigEndian, &m.MarginRateByMoney)
    binary.Read(r, binary.BigEndian, &m.MarginRateByVolume)
    binary.Read(r, binary.BigEndian, &m.StrikeFrozen)
    binary.Read(r, binary.BigEndian, &m.StrikeFrozenAmount)
    binary.Read(r, binary.BigEndian, &m.AbandonFrozen)
    m.ExchangeID = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.YdStrikeFrozen)
    m.InvestUnitID = readLen(r, 17)
}

type  CThostFtdcInstrumentMarginRateField struct {
    InstrumentID string;
    InvestorRange uint8;
    BrokerID string;
    InvestorID string;
    HedgeFlag uint8;
    LongMarginRatioByMoney float64;
    LongMarginRatioByVolume float64;
    ShortMarginRatioByMoney float64;
    ShortMarginRatioByVolume float64;
    IsRelative int32;
    ExchangeID string;
    InvestUnitID string;
}

func (m* CThostFtdcInstrumentMarginRateField) Marshal(w io.Writer) {
    writeLen(w, m.InstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.InvestorRange)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.HedgeFlag)
    binary.Write(w, binary.BigEndian, &m.LongMarginRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.LongMarginRatioByVolume)
    binary.Write(w, binary.BigEndian, &m.ShortMarginRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.ShortMarginRatioByVolume)
    binary.Write(w, binary.BigEndian, &m.IsRelative)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.InvestUnitID, 17)
}

func (m* CThostFtdcInstrumentMarginRateField) Unmarshal(r io.Reader) {
    m.InstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.InvestorRange)
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.HedgeFlag)
    binary.Read(r, binary.BigEndian, &m.LongMarginRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.LongMarginRatioByVolume)
    binary.Read(r, binary.BigEndian, &m.ShortMarginRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.ShortMarginRatioByVolume)
    binary.Read(r, binary.BigEndian, &m.IsRelative)
    m.ExchangeID = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
}

type  CThostFtdcInstrumentCommissionRateField struct {
    InstrumentID string;
    InvestorRange uint8;
    BrokerID string;
    InvestorID string;
    OpenRatioByMoney float64;
    OpenRatioByVolume float64;
    CloseRatioByMoney float64;
    CloseRatioByVolume float64;
    CloseTodayRatioByMoney float64;
    CloseTodayRatioByVolume float64;
    ExchangeID string;
    BizType uint8;
    InvestUnitID string;
}

func (m* CThostFtdcInstrumentCommissionRateField) Marshal(w io.Writer) {
    writeLen(w, m.InstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.InvestorRange)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.OpenRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.OpenRatioByVolume)
    binary.Write(w, binary.BigEndian, &m.CloseRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.CloseRatioByVolume)
    binary.Write(w, binary.BigEndian, &m.CloseTodayRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.CloseTodayRatioByVolume)
    writeLen(w, m.ExchangeID, 9)
    binary.Write(w, binary.BigEndian, &m.BizType)
    writeLen(w, m.InvestUnitID, 17)
}

func (m* CThostFtdcInstrumentCommissionRateField) Unmarshal(r io.Reader) {
    m.InstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.InvestorRange)
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.OpenRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.OpenRatioByVolume)
    binary.Read(r, binary.BigEndian, &m.CloseRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.CloseRatioByVolume)
    binary.Read(r, binary.BigEndian, &m.CloseTodayRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.CloseTodayRatioByVolume)
    m.ExchangeID = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.BizType)
    m.InvestUnitID = readLen(r, 17)
}

type  CThostFtdcDepthMarketDataField struct {
    TradingDay string;
    InstrumentID string;
    ExchangeID string;
    ExchangeInstID string;
    LastPrice float64;
    PreSettlementPrice float64;
    PreClosePrice float64;
    PreOpenInterest float64;
    OpenPrice float64;
    HighestPrice float64;
    LowestPrice float64;
    Volume int32;
    Turnover float64;
    OpenInterest float64;
    ClosePrice float64;
    SettlementPrice float64;
    UpperLimitPrice float64;
    LowerLimitPrice float64;
    PreDelta float64;
    CurrDelta float64;
    UpdateTime string;
    UpdateMillisec int32;
    BidPrice1 float64;
    BidVolume1 int32;
    AskPrice1 float64;
    AskVolume1 int32;
    BidPrice2 float64;
    BidVolume2 int32;
    AskPrice2 float64;
    AskVolume2 int32;
    BidPrice3 float64;
    BidVolume3 int32;
    AskPrice3 float64;
    AskVolume3 int32;
    BidPrice4 float64;
    BidVolume4 int32;
    AskPrice4 float64;
    AskVolume4 int32;
    BidPrice5 float64;
    BidVolume5 int32;
    AskPrice5 float64;
    AskVolume5 int32;
    AveragePrice float64;
    ActionDay string;
}

func (m* CThostFtdcDepthMarketDataField) Marshal(w io.Writer) {
    writeLen(w, m.TradingDay, 9)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ExchangeInstID, 31)
    binary.Write(w, binary.BigEndian, &m.LastPrice)
    binary.Write(w, binary.BigEndian, &m.PreSettlementPrice)
    binary.Write(w, binary.BigEndian, &m.PreClosePrice)
    binary.Write(w, binary.BigEndian, &m.PreOpenInterest)
    binary.Write(w, binary.BigEndian, &m.OpenPrice)
    binary.Write(w, binary.BigEndian, &m.HighestPrice)
    binary.Write(w, binary.BigEndian, &m.LowestPrice)
    binary.Write(w, binary.BigEndian, &m.Volume)
    binary.Write(w, binary.BigEndian, &m.Turnover)
    binary.Write(w, binary.BigEndian, &m.OpenInterest)
    binary.Write(w, binary.BigEndian, &m.ClosePrice)
    binary.Write(w, binary.BigEndian, &m.SettlementPrice)
    binary.Write(w, binary.BigEndian, &m.UpperLimitPrice)
    binary.Write(w, binary.BigEndian, &m.LowerLimitPrice)
    binary.Write(w, binary.BigEndian, &m.PreDelta)
    binary.Write(w, binary.BigEndian, &m.CurrDelta)
    writeLen(w, m.UpdateTime, 9)
    binary.Write(w, binary.BigEndian, &m.UpdateMillisec)
    binary.Write(w, binary.BigEndian, &m.BidPrice1)
    binary.Write(w, binary.BigEndian, &m.BidVolume1)
    binary.Write(w, binary.BigEndian, &m.AskPrice1)
    binary.Write(w, binary.BigEndian, &m.AskVolume1)
    binary.Write(w, binary.BigEndian, &m.BidPrice2)
    binary.Write(w, binary.BigEndian, &m.BidVolume2)
    binary.Write(w, binary.BigEndian, &m.AskPrice2)
    binary.Write(w, binary.BigEndian, &m.AskVolume2)
    binary.Write(w, binary.BigEndian, &m.BidPrice3)
    binary.Write(w, binary.BigEndian, &m.BidVolume3)
    binary.Write(w, binary.BigEndian, &m.AskPrice3)
    binary.Write(w, binary.BigEndian, &m.AskVolume3)
    binary.Write(w, binary.BigEndian, &m.BidPrice4)
    binary.Write(w, binary.BigEndian, &m.BidVolume4)
    binary.Write(w, binary.BigEndian, &m.AskPrice4)
    binary.Write(w, binary.BigEndian, &m.AskVolume4)
    binary.Write(w, binary.BigEndian, &m.BidPrice5)
    binary.Write(w, binary.BigEndian, &m.BidVolume5)
    binary.Write(w, binary.BigEndian, &m.AskPrice5)
    binary.Write(w, binary.BigEndian, &m.AskVolume5)
    binary.Write(w, binary.BigEndian, &m.AveragePrice)
    writeLen(w, m.ActionDay, 9)
}

func (m* CThostFtdcDepthMarketDataField) Unmarshal(r io.Reader) {
    m.TradingDay = readLen(r, 9)
    m.InstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
    m.ExchangeInstID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.LastPrice)
    binary.Read(r, binary.BigEndian, &m.PreSettlementPrice)
    binary.Read(r, binary.BigEndian, &m.PreClosePrice)
    binary.Read(r, binary.BigEndian, &m.PreOpenInterest)
    binary.Read(r, binary.BigEndian, &m.OpenPrice)
    binary.Read(r, binary.BigEndian, &m.HighestPrice)
    binary.Read(r, binary.BigEndian, &m.LowestPrice)
    binary.Read(r, binary.BigEndian, &m.Volume)
    binary.Read(r, binary.BigEndian, &m.Turnover)
    binary.Read(r, binary.BigEndian, &m.OpenInterest)
    binary.Read(r, binary.BigEndian, &m.ClosePrice)
    binary.Read(r, binary.BigEndian, &m.SettlementPrice)
    binary.Read(r, binary.BigEndian, &m.UpperLimitPrice)
    binary.Read(r, binary.BigEndian, &m.LowerLimitPrice)
    binary.Read(r, binary.BigEndian, &m.PreDelta)
    binary.Read(r, binary.BigEndian, &m.CurrDelta)
    m.UpdateTime = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.UpdateMillisec)
    binary.Read(r, binary.BigEndian, &m.BidPrice1)
    binary.Read(r, binary.BigEndian, &m.BidVolume1)
    binary.Read(r, binary.BigEndian, &m.AskPrice1)
    binary.Read(r, binary.BigEndian, &m.AskVolume1)
    binary.Read(r, binary.BigEndian, &m.BidPrice2)
    binary.Read(r, binary.BigEndian, &m.BidVolume2)
    binary.Read(r, binary.BigEndian, &m.AskPrice2)
    binary.Read(r, binary.BigEndian, &m.AskVolume2)
    binary.Read(r, binary.BigEndian, &m.BidPrice3)
    binary.Read(r, binary.BigEndian, &m.BidVolume3)
    binary.Read(r, binary.BigEndian, &m.AskPrice3)
    binary.Read(r, binary.BigEndian, &m.AskVolume3)
    binary.Read(r, binary.BigEndian, &m.BidPrice4)
    binary.Read(r, binary.BigEndian, &m.BidVolume4)
    binary.Read(r, binary.BigEndian, &m.AskPrice4)
    binary.Read(r, binary.BigEndian, &m.AskVolume4)
    binary.Read(r, binary.BigEndian, &m.BidPrice5)
    binary.Read(r, binary.BigEndian, &m.BidVolume5)
    binary.Read(r, binary.BigEndian, &m.AskPrice5)
    binary.Read(r, binary.BigEndian, &m.AskVolume5)
    binary.Read(r, binary.BigEndian, &m.AveragePrice)
    m.ActionDay = readLen(r, 9)
}

type  CThostFtdcInstrumentTradingRightField struct {
    InstrumentID string;
    InvestorRange uint8;
    BrokerID string;
    InvestorID string;
    TradingRight uint8;
    ExchangeID string;
    BizType uint8;
}

func (m* CThostFtdcInstrumentTradingRightField) Marshal(w io.Writer) {
    writeLen(w, m.InstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.InvestorRange)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.TradingRight)
    writeLen(w, m.ExchangeID, 9)
    binary.Write(w, binary.BigEndian, &m.BizType)
}

func (m* CThostFtdcInstrumentTradingRightField) Unmarshal(r io.Reader) {
    m.InstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.InvestorRange)
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.TradingRight)
    m.ExchangeID = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.BizType)
}

type  CThostFtdcBrokerUserField struct {
    BrokerID string;
    UserID string;
    UserName string;
    UserType uint8;
    IsActive int32;
    IsUsingOTP int32;
    IsAuthForce int32;
}

func (m* CThostFtdcBrokerUserField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.UserName, 81)
    binary.Write(w, binary.BigEndian, &m.UserType)
    binary.Write(w, binary.BigEndian, &m.IsActive)
    binary.Write(w, binary.BigEndian, &m.IsUsingOTP)
    binary.Write(w, binary.BigEndian, &m.IsAuthForce)
}

func (m* CThostFtdcBrokerUserField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.UserID = readLen(r, 16)
    m.UserName = readLen(r, 81)
    binary.Read(r, binary.BigEndian, &m.UserType)
    binary.Read(r, binary.BigEndian, &m.IsActive)
    binary.Read(r, binary.BigEndian, &m.IsUsingOTP)
    binary.Read(r, binary.BigEndian, &m.IsAuthForce)
}

type  CThostFtdcBrokerUserPasswordField struct {
    BrokerID string;
    UserID string;
    Password string;
    LastUpdateTime string;
    LastLoginTime string;
    ExpireDate string;
    WeakExpireDate string;
}

func (m* CThostFtdcBrokerUserPasswordField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.Password, 41)
    writeLen(w, m.LastUpdateTime, 17)
    writeLen(w, m.LastLoginTime, 17)
    writeLen(w, m.ExpireDate, 9)
    writeLen(w, m.WeakExpireDate, 9)
}

func (m* CThostFtdcBrokerUserPasswordField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.UserID = readLen(r, 16)
    m.Password = readLen(r, 41)
    m.LastUpdateTime = readLen(r, 17)
    m.LastLoginTime = readLen(r, 17)
    m.ExpireDate = readLen(r, 9)
    m.WeakExpireDate = readLen(r, 9)
}

type  CThostFtdcBrokerUserFunctionField struct {
    BrokerID string;
    UserID string;
    BrokerFunctionCode uint8;
}

func (m* CThostFtdcBrokerUserFunctionField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.BrokerFunctionCode)
}

func (m* CThostFtdcBrokerUserFunctionField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.BrokerFunctionCode)
}

type  CThostFtdcTraderOfferField struct {
    ExchangeID string;
    TraderID string;
    ParticipantID string;
    Password string;
    InstallID int32;
    OrderLocalID string;
    TraderConnectStatus uint8;
    ConnectRequestDate string;
    ConnectRequestTime string;
    LastReportDate string;
    LastReportTime string;
    ConnectDate string;
    ConnectTime string;
    StartDate string;
    StartTime string;
    TradingDay string;
    BrokerID string;
    MaxTradeID string;
    MaxOrderMessageReference string;
    BizType uint8;
}

func (m* CThostFtdcTraderOfferField) Marshal(w io.Writer) {
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.TraderID, 21)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.Password, 41)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    writeLen(w, m.OrderLocalID, 13)
    binary.Write(w, binary.BigEndian, &m.TraderConnectStatus)
    writeLen(w, m.ConnectRequestDate, 9)
    writeLen(w, m.ConnectRequestTime, 9)
    writeLen(w, m.LastReportDate, 9)
    writeLen(w, m.LastReportTime, 9)
    writeLen(w, m.ConnectDate, 9)
    writeLen(w, m.ConnectTime, 9)
    writeLen(w, m.StartDate, 9)
    writeLen(w, m.StartTime, 9)
    writeLen(w, m.TradingDay, 9)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.MaxTradeID, 21)
    writeLen(w, m.MaxOrderMessageReference, 7)
    binary.Write(w, binary.BigEndian, &m.BizType)
}

func (m* CThostFtdcTraderOfferField) Unmarshal(r io.Reader) {
    m.ExchangeID = readLen(r, 9)
    m.TraderID = readLen(r, 21)
    m.ParticipantID = readLen(r, 11)
    m.Password = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    m.OrderLocalID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.TraderConnectStatus)
    m.ConnectRequestDate = readLen(r, 9)
    m.ConnectRequestTime = readLen(r, 9)
    m.LastReportDate = readLen(r, 9)
    m.LastReportTime = readLen(r, 9)
    m.ConnectDate = readLen(r, 9)
    m.ConnectTime = readLen(r, 9)
    m.StartDate = readLen(r, 9)
    m.StartTime = readLen(r, 9)
    m.TradingDay = readLen(r, 9)
    m.BrokerID = readLen(r, 11)
    m.MaxTradeID = readLen(r, 21)
    m.MaxOrderMessageReference = readLen(r, 7)
    binary.Read(r, binary.BigEndian, &m.BizType)
}

type  CThostFtdcSettlementInfoField struct {
    TradingDay string;
    SettlementID int32;
    BrokerID string;
    InvestorID string;
    SequenceNo int32;
    Content string;
    AccountID string;
    CurrencyID string;
}

func (m* CThostFtdcSettlementInfoField) Marshal(w io.Writer) {
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.SettlementID)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.SequenceNo)
    writeLen(w, m.Content, 501)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.CurrencyID, 4)
}

func (m* CThostFtdcSettlementInfoField) Unmarshal(r io.Reader) {
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.SettlementID)
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.SequenceNo)
    m.Content = readLen(r, 501)
    m.AccountID = readLen(r, 13)
    m.CurrencyID = readLen(r, 4)
}

type  CThostFtdcInstrumentMarginRateAdjustField struct {
    InstrumentID string;
    InvestorRange uint8;
    BrokerID string;
    InvestorID string;
    HedgeFlag uint8;
    LongMarginRatioByMoney float64;
    LongMarginRatioByVolume float64;
    ShortMarginRatioByMoney float64;
    ShortMarginRatioByVolume float64;
    IsRelative int32;
}

func (m* CThostFtdcInstrumentMarginRateAdjustField) Marshal(w io.Writer) {
    writeLen(w, m.InstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.InvestorRange)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.HedgeFlag)
    binary.Write(w, binary.BigEndian, &m.LongMarginRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.LongMarginRatioByVolume)
    binary.Write(w, binary.BigEndian, &m.ShortMarginRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.ShortMarginRatioByVolume)
    binary.Write(w, binary.BigEndian, &m.IsRelative)
}

func (m* CThostFtdcInstrumentMarginRateAdjustField) Unmarshal(r io.Reader) {
    m.InstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.InvestorRange)
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.HedgeFlag)
    binary.Read(r, binary.BigEndian, &m.LongMarginRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.LongMarginRatioByVolume)
    binary.Read(r, binary.BigEndian, &m.ShortMarginRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.ShortMarginRatioByVolume)
    binary.Read(r, binary.BigEndian, &m.IsRelative)
}

type  CThostFtdcExchangeMarginRateField struct {
    BrokerID string;
    InstrumentID string;
    HedgeFlag uint8;
    LongMarginRatioByMoney float64;
    LongMarginRatioByVolume float64;
    ShortMarginRatioByMoney float64;
    ShortMarginRatioByVolume float64;
    ExchangeID string;
}

func (m* CThostFtdcExchangeMarginRateField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.HedgeFlag)
    binary.Write(w, binary.BigEndian, &m.LongMarginRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.LongMarginRatioByVolume)
    binary.Write(w, binary.BigEndian, &m.ShortMarginRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.ShortMarginRatioByVolume)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcExchangeMarginRateField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.HedgeFlag)
    binary.Read(r, binary.BigEndian, &m.LongMarginRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.LongMarginRatioByVolume)
    binary.Read(r, binary.BigEndian, &m.ShortMarginRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.ShortMarginRatioByVolume)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcExchangeMarginRateAdjustField struct {
    BrokerID string;
    InstrumentID string;
    HedgeFlag uint8;
    LongMarginRatioByMoney float64;
    LongMarginRatioByVolume float64;
    ShortMarginRatioByMoney float64;
    ShortMarginRatioByVolume float64;
    ExchLongMarginRatioByMoney float64;
    ExchLongMarginRatioByVolume float64;
    ExchShortMarginRatioByMoney float64;
    ExchShortMarginRatioByVolume float64;
    NoLongMarginRatioByMoney float64;
    NoLongMarginRatioByVolume float64;
    NoShortMarginRatioByMoney float64;
    NoShortMarginRatioByVolume float64;
}

func (m* CThostFtdcExchangeMarginRateAdjustField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.HedgeFlag)
    binary.Write(w, binary.BigEndian, &m.LongMarginRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.LongMarginRatioByVolume)
    binary.Write(w, binary.BigEndian, &m.ShortMarginRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.ShortMarginRatioByVolume)
    binary.Write(w, binary.BigEndian, &m.ExchLongMarginRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.ExchLongMarginRatioByVolume)
    binary.Write(w, binary.BigEndian, &m.ExchShortMarginRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.ExchShortMarginRatioByVolume)
    binary.Write(w, binary.BigEndian, &m.NoLongMarginRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.NoLongMarginRatioByVolume)
    binary.Write(w, binary.BigEndian, &m.NoShortMarginRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.NoShortMarginRatioByVolume)
}

func (m* CThostFtdcExchangeMarginRateAdjustField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.HedgeFlag)
    binary.Read(r, binary.BigEndian, &m.LongMarginRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.LongMarginRatioByVolume)
    binary.Read(r, binary.BigEndian, &m.ShortMarginRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.ShortMarginRatioByVolume)
    binary.Read(r, binary.BigEndian, &m.ExchLongMarginRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.ExchLongMarginRatioByVolume)
    binary.Read(r, binary.BigEndian, &m.ExchShortMarginRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.ExchShortMarginRatioByVolume)
    binary.Read(r, binary.BigEndian, &m.NoLongMarginRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.NoLongMarginRatioByVolume)
    binary.Read(r, binary.BigEndian, &m.NoShortMarginRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.NoShortMarginRatioByVolume)
}

type  CThostFtdcExchangeRateField struct {
    BrokerID string;
    FromCurrencyID string;
    FromCurrencyUnit float64;
    ToCurrencyID string;
    ExchangeRate float64;
}

func (m* CThostFtdcExchangeRateField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.FromCurrencyID, 4)
    binary.Write(w, binary.BigEndian, &m.FromCurrencyUnit)
    writeLen(w, m.ToCurrencyID, 4)
    binary.Write(w, binary.BigEndian, &m.ExchangeRate)
}

func (m* CThostFtdcExchangeRateField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.FromCurrencyID = readLen(r, 4)
    binary.Read(r, binary.BigEndian, &m.FromCurrencyUnit)
    m.ToCurrencyID = readLen(r, 4)
    binary.Read(r, binary.BigEndian, &m.ExchangeRate)
}

type  CThostFtdcSettlementRefField struct {
    TradingDay string;
    SettlementID int32;
}

func (m* CThostFtdcSettlementRefField) Marshal(w io.Writer) {
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.SettlementID)
}

func (m* CThostFtdcSettlementRefField) Unmarshal(r io.Reader) {
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.SettlementID)
}

type  CThostFtdcCurrentTimeField struct {
    CurrDate string;
    CurrTime string;
    CurrMillisec int32;
    ActionDay string;
}

func (m* CThostFtdcCurrentTimeField) Marshal(w io.Writer) {
    writeLen(w, m.CurrDate, 9)
    writeLen(w, m.CurrTime, 9)
    binary.Write(w, binary.BigEndian, &m.CurrMillisec)
    writeLen(w, m.ActionDay, 9)
}

func (m* CThostFtdcCurrentTimeField) Unmarshal(r io.Reader) {
    m.CurrDate = readLen(r, 9)
    m.CurrTime = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.CurrMillisec)
    m.ActionDay = readLen(r, 9)
}

type  CThostFtdcCommPhaseField struct {
    TradingDay string;
    CommPhaseNo int16;
    SystemID string;
}

func (m* CThostFtdcCommPhaseField) Marshal(w io.Writer) {
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.CommPhaseNo)
    writeLen(w, m.SystemID, 21)
}

func (m* CThostFtdcCommPhaseField) Unmarshal(r io.Reader) {
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.CommPhaseNo)
    m.SystemID = readLen(r, 21)
}

type  CThostFtdcLoginInfoField struct {
    FrontID int32;
    SessionID int32;
    BrokerID string;
    UserID string;
    LoginDate string;
    LoginTime string;
    IPAddress string;
    UserProductInfo string;
    InterfaceProductInfo string;
    ProtocolInfo string;
    SystemName string;
    Password string;
    MaxOrderRef string;
    SHFETime string;
    DCETime string;
    CZCETime string;
    FFEXTime string;
    MacAddress string;
    OneTimePassword string;
    INETime string;
    IsQryControl int32;
    LoginRemark string;
    SmsCode string;
    EncryptType int32;
}

func (m* CThostFtdcLoginInfoField) Marshal(w io.Writer) {
    binary.Write(w, binary.BigEndian, &m.FrontID)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.LoginDate, 9)
    writeLen(w, m.LoginTime, 9)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.UserProductInfo, 11)
    writeLen(w, m.InterfaceProductInfo, 11)
    writeLen(w, m.ProtocolInfo, 11)
    writeLen(w, m.SystemName, 41)
    writeLen(w, m.Password, 41)
    writeLen(w, m.MaxOrderRef, 13)
    writeLen(w, m.SHFETime, 9)
    writeLen(w, m.DCETime, 9)
    writeLen(w, m.CZCETime, 9)
    writeLen(w, m.FFEXTime, 9)
    writeLen(w, m.MacAddress, 21)
    writeLen(w, m.OneTimePassword, 41)
    writeLen(w, m.INETime, 9)
    binary.Write(w, binary.BigEndian, &m.IsQryControl)
    writeLen(w, m.LoginRemark, 36)
    writeLen(w, m.SmsCode, 13)
    binary.Write(w, binary.BigEndian, &m.EncryptType)
}

func (m* CThostFtdcLoginInfoField) Unmarshal(r io.Reader) {
    binary.Read(r, binary.BigEndian, &m.FrontID)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.BrokerID = readLen(r, 11)
    m.UserID = readLen(r, 16)
    m.LoginDate = readLen(r, 9)
    m.LoginTime = readLen(r, 9)
    m.IPAddress = readLen(r, 16)
    m.UserProductInfo = readLen(r, 11)
    m.InterfaceProductInfo = readLen(r, 11)
    m.ProtocolInfo = readLen(r, 11)
    m.SystemName = readLen(r, 41)
    m.Password = readLen(r, 41)
    m.MaxOrderRef = readLen(r, 13)
    m.SHFETime = readLen(r, 9)
    m.DCETime = readLen(r, 9)
    m.CZCETime = readLen(r, 9)
    m.FFEXTime = readLen(r, 9)
    m.MacAddress = readLen(r, 21)
    m.OneTimePassword = readLen(r, 41)
    m.INETime = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.IsQryControl)
    m.LoginRemark = readLen(r, 36)
    m.SmsCode = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.EncryptType)
}

type  CThostFtdcLogoutAllField struct {
    FrontID int32;
    SessionID int32;
    SystemName string;
}

func (m* CThostFtdcLogoutAllField) Marshal(w io.Writer) {
    binary.Write(w, binary.BigEndian, &m.FrontID)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.SystemName, 41)
}

func (m* CThostFtdcLogoutAllField) Unmarshal(r io.Reader) {
    binary.Read(r, binary.BigEndian, &m.FrontID)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.SystemName = readLen(r, 41)
}

type  CThostFtdcFrontStatusField struct {
    FrontID int32;
    LastReportDate string;
    LastReportTime string;
    IsActive int32;
}

func (m* CThostFtdcFrontStatusField) Marshal(w io.Writer) {
    binary.Write(w, binary.BigEndian, &m.FrontID)
    writeLen(w, m.LastReportDate, 9)
    writeLen(w, m.LastReportTime, 9)
    binary.Write(w, binary.BigEndian, &m.IsActive)
}

func (m* CThostFtdcFrontStatusField) Unmarshal(r io.Reader) {
    binary.Read(r, binary.BigEndian, &m.FrontID)
    m.LastReportDate = readLen(r, 9)
    m.LastReportTime = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.IsActive)
}

type  CThostFtdcUserPasswordUpdateField struct {
    BrokerID string;
    UserID string;
    OldPassword string;
    NewPassword string;
    EncryptType int32;
}

func (m* CThostFtdcUserPasswordUpdateField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.OldPassword, 41)
    writeLen(w, m.NewPassword, 41)
    binary.Write(w, binary.BigEndian, &m.EncryptType)
}

func (m* CThostFtdcUserPasswordUpdateField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.UserID = readLen(r, 16)
    m.OldPassword = readLen(r, 41)
    m.NewPassword = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.EncryptType)
}

type  CThostFtdcInputOrderField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    OrderRef string;
    UserID string;
    OrderPriceType uint8;
    Direction uint8;
    CombOffsetFlag string;
    CombHedgeFlag string;
    LimitPrice float64;
    VolumeTotalOriginal int32;
    TimeCondition uint8;
    GTDDate string;
    VolumeCondition uint8;
    MinVolume int32;
    ContingentCondition uint8;
    StopPrice float64;
    ForceCloseReason uint8;
    IsAutoSuspend int32;
    BusinessUnit string;
    RequestID int32;
    UserForceClose int32;
    IsSwapOrder int32;
    ExchangeID string;
    InvestUnitID string;
    AccountID string;
    CurrencyID string;
    ClientID string;
    IPAddress string;
    MacAddress string;
}

func (m* CThostFtdcInputOrderField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.OrderRef, 13)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.OrderPriceType)
    binary.Write(w, binary.BigEndian, &m.Direction)
    writeLen(w, m.CombOffsetFlag, 5)
    writeLen(w, m.CombHedgeFlag, 5)
    binary.Write(w, binary.BigEndian, &m.LimitPrice)
    binary.Write(w, binary.BigEndian, &m.VolumeTotalOriginal)
    binary.Write(w, binary.BigEndian, &m.TimeCondition)
    writeLen(w, m.GTDDate, 9)
    binary.Write(w, binary.BigEndian, &m.VolumeCondition)
    binary.Write(w, binary.BigEndian, &m.MinVolume)
    binary.Write(w, binary.BigEndian, &m.ContingentCondition)
    binary.Write(w, binary.BigEndian, &m.StopPrice)
    binary.Write(w, binary.BigEndian, &m.ForceCloseReason)
    binary.Write(w, binary.BigEndian, &m.IsAutoSuspend)
    writeLen(w, m.BusinessUnit, 21)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    binary.Write(w, binary.BigEndian, &m.UserForceClose)
    binary.Write(w, binary.BigEndian, &m.IsSwapOrder)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.InvestUnitID, 17)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.CurrencyID, 4)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
}

func (m* CThostFtdcInputOrderField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.OrderRef = readLen(r, 13)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.OrderPriceType)
    binary.Read(r, binary.BigEndian, &m.Direction)
    m.CombOffsetFlag = readLen(r, 5)
    m.CombHedgeFlag = readLen(r, 5)
    binary.Read(r, binary.BigEndian, &m.LimitPrice)
    binary.Read(r, binary.BigEndian, &m.VolumeTotalOriginal)
    binary.Read(r, binary.BigEndian, &m.TimeCondition)
    m.GTDDate = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.VolumeCondition)
    binary.Read(r, binary.BigEndian, &m.MinVolume)
    binary.Read(r, binary.BigEndian, &m.ContingentCondition)
    binary.Read(r, binary.BigEndian, &m.StopPrice)
    binary.Read(r, binary.BigEndian, &m.ForceCloseReason)
    binary.Read(r, binary.BigEndian, &m.IsAutoSuspend)
    m.BusinessUnit = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    binary.Read(r, binary.BigEndian, &m.UserForceClose)
    binary.Read(r, binary.BigEndian, &m.IsSwapOrder)
    m.ExchangeID = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
    m.AccountID = readLen(r, 13)
    m.CurrencyID = readLen(r, 4)
    m.ClientID = readLen(r, 11)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
}

type  CThostFtdcOrderField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    OrderRef string;
    UserID string;
    OrderPriceType uint8;
    Direction uint8;
    CombOffsetFlag string;
    CombHedgeFlag string;
    LimitPrice float64;
    VolumeTotalOriginal int32;
    TimeCondition uint8;
    GTDDate string;
    VolumeCondition uint8;
    MinVolume int32;
    ContingentCondition uint8;
    StopPrice float64;
    ForceCloseReason uint8;
    IsAutoSuspend int32;
    BusinessUnit string;
    RequestID int32;
    OrderLocalID string;
    ExchangeID string;
    ParticipantID string;
    ClientID string;
    ExchangeInstID string;
    TraderID string;
    InstallID int32;
    OrderSubmitStatus uint8;
    NotifySequence int32;
    TradingDay string;
    SettlementID int32;
    OrderSysID string;
    OrderSource uint8;
    OrderStatus uint8;
    OrderType uint8;
    VolumeTraded int32;
    VolumeTotal int32;
    InsertDate string;
    InsertTime string;
    ActiveTime string;
    SuspendTime string;
    UpdateTime string;
    CancelTime string;
    ActiveTraderID string;
    ClearingPartID string;
    SequenceNo int32;
    FrontID int32;
    SessionID int32;
    UserProductInfo string;
    StatusMsg string;
    UserForceClose int32;
    ActiveUserID string;
    BrokerOrderSeq int32;
    RelativeOrderSysID string;
    ZCETotalTradedVolume int32;
    IsSwapOrder int32;
    BranchID string;
    InvestUnitID string;
    AccountID string;
    CurrencyID string;
    IPAddress string;
    MacAddress string;
}

func (m* CThostFtdcOrderField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.OrderRef, 13)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.OrderPriceType)
    binary.Write(w, binary.BigEndian, &m.Direction)
    writeLen(w, m.CombOffsetFlag, 5)
    writeLen(w, m.CombHedgeFlag, 5)
    binary.Write(w, binary.BigEndian, &m.LimitPrice)
    binary.Write(w, binary.BigEndian, &m.VolumeTotalOriginal)
    binary.Write(w, binary.BigEndian, &m.TimeCondition)
    writeLen(w, m.GTDDate, 9)
    binary.Write(w, binary.BigEndian, &m.VolumeCondition)
    binary.Write(w, binary.BigEndian, &m.MinVolume)
    binary.Write(w, binary.BigEndian, &m.ContingentCondition)
    binary.Write(w, binary.BigEndian, &m.StopPrice)
    binary.Write(w, binary.BigEndian, &m.ForceCloseReason)
    binary.Write(w, binary.BigEndian, &m.IsAutoSuspend)
    writeLen(w, m.BusinessUnit, 21)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    writeLen(w, m.OrderLocalID, 13)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.ExchangeInstID, 31)
    writeLen(w, m.TraderID, 21)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    binary.Write(w, binary.BigEndian, &m.OrderSubmitStatus)
    binary.Write(w, binary.BigEndian, &m.NotifySequence)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.SettlementID)
    writeLen(w, m.OrderSysID, 21)
    binary.Write(w, binary.BigEndian, &m.OrderSource)
    binary.Write(w, binary.BigEndian, &m.OrderStatus)
    binary.Write(w, binary.BigEndian, &m.OrderType)
    binary.Write(w, binary.BigEndian, &m.VolumeTraded)
    binary.Write(w, binary.BigEndian, &m.VolumeTotal)
    writeLen(w, m.InsertDate, 9)
    writeLen(w, m.InsertTime, 9)
    writeLen(w, m.ActiveTime, 9)
    writeLen(w, m.SuspendTime, 9)
    writeLen(w, m.UpdateTime, 9)
    writeLen(w, m.CancelTime, 9)
    writeLen(w, m.ActiveTraderID, 21)
    writeLen(w, m.ClearingPartID, 11)
    binary.Write(w, binary.BigEndian, &m.SequenceNo)
    binary.Write(w, binary.BigEndian, &m.FrontID)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.UserProductInfo, 11)
    writeLen(w, m.StatusMsg, 81)
    binary.Write(w, binary.BigEndian, &m.UserForceClose)
    writeLen(w, m.ActiveUserID, 16)
    binary.Write(w, binary.BigEndian, &m.BrokerOrderSeq)
    writeLen(w, m.RelativeOrderSysID, 21)
    binary.Write(w, binary.BigEndian, &m.ZCETotalTradedVolume)
    binary.Write(w, binary.BigEndian, &m.IsSwapOrder)
    writeLen(w, m.BranchID, 9)
    writeLen(w, m.InvestUnitID, 17)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.CurrencyID, 4)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
}

func (m* CThostFtdcOrderField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.OrderRef = readLen(r, 13)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.OrderPriceType)
    binary.Read(r, binary.BigEndian, &m.Direction)
    m.CombOffsetFlag = readLen(r, 5)
    m.CombHedgeFlag = readLen(r, 5)
    binary.Read(r, binary.BigEndian, &m.LimitPrice)
    binary.Read(r, binary.BigEndian, &m.VolumeTotalOriginal)
    binary.Read(r, binary.BigEndian, &m.TimeCondition)
    m.GTDDate = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.VolumeCondition)
    binary.Read(r, binary.BigEndian, &m.MinVolume)
    binary.Read(r, binary.BigEndian, &m.ContingentCondition)
    binary.Read(r, binary.BigEndian, &m.StopPrice)
    binary.Read(r, binary.BigEndian, &m.ForceCloseReason)
    binary.Read(r, binary.BigEndian, &m.IsAutoSuspend)
    m.BusinessUnit = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    m.OrderLocalID = readLen(r, 13)
    m.ExchangeID = readLen(r, 9)
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.ExchangeInstID = readLen(r, 31)
    m.TraderID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    binary.Read(r, binary.BigEndian, &m.OrderSubmitStatus)
    binary.Read(r, binary.BigEndian, &m.NotifySequence)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.SettlementID)
    m.OrderSysID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.OrderSource)
    binary.Read(r, binary.BigEndian, &m.OrderStatus)
    binary.Read(r, binary.BigEndian, &m.OrderType)
    binary.Read(r, binary.BigEndian, &m.VolumeTraded)
    binary.Read(r, binary.BigEndian, &m.VolumeTotal)
    m.InsertDate = readLen(r, 9)
    m.InsertTime = readLen(r, 9)
    m.ActiveTime = readLen(r, 9)
    m.SuspendTime = readLen(r, 9)
    m.UpdateTime = readLen(r, 9)
    m.CancelTime = readLen(r, 9)
    m.ActiveTraderID = readLen(r, 21)
    m.ClearingPartID = readLen(r, 11)
    binary.Read(r, binary.BigEndian, &m.SequenceNo)
    binary.Read(r, binary.BigEndian, &m.FrontID)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.UserProductInfo = readLen(r, 11)
    m.StatusMsg = readLen(r, 81)
    binary.Read(r, binary.BigEndian, &m.UserForceClose)
    m.ActiveUserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.BrokerOrderSeq)
    m.RelativeOrderSysID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.ZCETotalTradedVolume)
    binary.Read(r, binary.BigEndian, &m.IsSwapOrder)
    m.BranchID = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
    m.AccountID = readLen(r, 13)
    m.CurrencyID = readLen(r, 4)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
}

type  CThostFtdcExchangeOrderField struct {
    OrderPriceType uint8;
    Direction uint8;
    CombOffsetFlag string;
    CombHedgeFlag string;
    LimitPrice float64;
    VolumeTotalOriginal int32;
    TimeCondition uint8;
    GTDDate string;
    VolumeCondition uint8;
    MinVolume int32;
    ContingentCondition uint8;
    StopPrice float64;
    ForceCloseReason uint8;
    IsAutoSuspend int32;
    BusinessUnit string;
    RequestID int32;
    OrderLocalID string;
    ExchangeID string;
    ParticipantID string;
    ClientID string;
    ExchangeInstID string;
    TraderID string;
    InstallID int32;
    OrderSubmitStatus uint8;
    NotifySequence int32;
    TradingDay string;
    SettlementID int32;
    OrderSysID string;
    OrderSource uint8;
    OrderStatus uint8;
    OrderType uint8;
    VolumeTraded int32;
    VolumeTotal int32;
    InsertDate string;
    InsertTime string;
    ActiveTime string;
    SuspendTime string;
    UpdateTime string;
    CancelTime string;
    ActiveTraderID string;
    ClearingPartID string;
    SequenceNo int32;
    BranchID string;
    IPAddress string;
    MacAddress string;
}

func (m* CThostFtdcExchangeOrderField) Marshal(w io.Writer) {
    binary.Write(w, binary.BigEndian, &m.OrderPriceType)
    binary.Write(w, binary.BigEndian, &m.Direction)
    writeLen(w, m.CombOffsetFlag, 5)
    writeLen(w, m.CombHedgeFlag, 5)
    binary.Write(w, binary.BigEndian, &m.LimitPrice)
    binary.Write(w, binary.BigEndian, &m.VolumeTotalOriginal)
    binary.Write(w, binary.BigEndian, &m.TimeCondition)
    writeLen(w, m.GTDDate, 9)
    binary.Write(w, binary.BigEndian, &m.VolumeCondition)
    binary.Write(w, binary.BigEndian, &m.MinVolume)
    binary.Write(w, binary.BigEndian, &m.ContingentCondition)
    binary.Write(w, binary.BigEndian, &m.StopPrice)
    binary.Write(w, binary.BigEndian, &m.ForceCloseReason)
    binary.Write(w, binary.BigEndian, &m.IsAutoSuspend)
    writeLen(w, m.BusinessUnit, 21)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    writeLen(w, m.OrderLocalID, 13)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.ExchangeInstID, 31)
    writeLen(w, m.TraderID, 21)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    binary.Write(w, binary.BigEndian, &m.OrderSubmitStatus)
    binary.Write(w, binary.BigEndian, &m.NotifySequence)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.SettlementID)
    writeLen(w, m.OrderSysID, 21)
    binary.Write(w, binary.BigEndian, &m.OrderSource)
    binary.Write(w, binary.BigEndian, &m.OrderStatus)
    binary.Write(w, binary.BigEndian, &m.OrderType)
    binary.Write(w, binary.BigEndian, &m.VolumeTraded)
    binary.Write(w, binary.BigEndian, &m.VolumeTotal)
    writeLen(w, m.InsertDate, 9)
    writeLen(w, m.InsertTime, 9)
    writeLen(w, m.ActiveTime, 9)
    writeLen(w, m.SuspendTime, 9)
    writeLen(w, m.UpdateTime, 9)
    writeLen(w, m.CancelTime, 9)
    writeLen(w, m.ActiveTraderID, 21)
    writeLen(w, m.ClearingPartID, 11)
    binary.Write(w, binary.BigEndian, &m.SequenceNo)
    writeLen(w, m.BranchID, 9)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
}

func (m* CThostFtdcExchangeOrderField) Unmarshal(r io.Reader) {
    binary.Read(r, binary.BigEndian, &m.OrderPriceType)
    binary.Read(r, binary.BigEndian, &m.Direction)
    m.CombOffsetFlag = readLen(r, 5)
    m.CombHedgeFlag = readLen(r, 5)
    binary.Read(r, binary.BigEndian, &m.LimitPrice)
    binary.Read(r, binary.BigEndian, &m.VolumeTotalOriginal)
    binary.Read(r, binary.BigEndian, &m.TimeCondition)
    m.GTDDate = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.VolumeCondition)
    binary.Read(r, binary.BigEndian, &m.MinVolume)
    binary.Read(r, binary.BigEndian, &m.ContingentCondition)
    binary.Read(r, binary.BigEndian, &m.StopPrice)
    binary.Read(r, binary.BigEndian, &m.ForceCloseReason)
    binary.Read(r, binary.BigEndian, &m.IsAutoSuspend)
    m.BusinessUnit = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    m.OrderLocalID = readLen(r, 13)
    m.ExchangeID = readLen(r, 9)
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.ExchangeInstID = readLen(r, 31)
    m.TraderID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    binary.Read(r, binary.BigEndian, &m.OrderSubmitStatus)
    binary.Read(r, binary.BigEndian, &m.NotifySequence)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.SettlementID)
    m.OrderSysID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.OrderSource)
    binary.Read(r, binary.BigEndian, &m.OrderStatus)
    binary.Read(r, binary.BigEndian, &m.OrderType)
    binary.Read(r, binary.BigEndian, &m.VolumeTraded)
    binary.Read(r, binary.BigEndian, &m.VolumeTotal)
    m.InsertDate = readLen(r, 9)
    m.InsertTime = readLen(r, 9)
    m.ActiveTime = readLen(r, 9)
    m.SuspendTime = readLen(r, 9)
    m.UpdateTime = readLen(r, 9)
    m.CancelTime = readLen(r, 9)
    m.ActiveTraderID = readLen(r, 21)
    m.ClearingPartID = readLen(r, 11)
    binary.Read(r, binary.BigEndian, &m.SequenceNo)
    m.BranchID = readLen(r, 9)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
}

type  CThostFtdcExchangeOrderInsertErrorField struct {
    ExchangeID string;
    ParticipantID string;
    TraderID string;
    InstallID int32;
    OrderLocalID string;
    ErrorID int32;
    ErrorMsg string;
}

func (m* CThostFtdcExchangeOrderInsertErrorField) Marshal(w io.Writer) {
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.TraderID, 21)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    writeLen(w, m.OrderLocalID, 13)
    binary.Write(w, binary.BigEndian, &m.ErrorID)
    writeLen(w, m.ErrorMsg, 81)
}

func (m* CThostFtdcExchangeOrderInsertErrorField) Unmarshal(r io.Reader) {
    m.ExchangeID = readLen(r, 9)
    m.ParticipantID = readLen(r, 11)
    m.TraderID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    m.OrderLocalID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.ErrorID)
    m.ErrorMsg = readLen(r, 81)
}

type  CThostFtdcInputOrderActionField struct {
    BrokerID string;
    InvestorID string;
    OrderActionRef int32;
    OrderRef string;
    RequestID int32;
    FrontID int32;
    SessionID int32;
    ExchangeID string;
    OrderSysID string;
    ActionFlag uint8;
    LimitPrice float64;
    VolumeChange int32;
    UserID string;
    InstrumentID string;
    InvestUnitID string;
    IPAddress string;
    MacAddress string;
}

func (m* CThostFtdcInputOrderActionField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.OrderActionRef)
    writeLen(w, m.OrderRef, 13)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    binary.Write(w, binary.BigEndian, &m.FrontID)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.OrderSysID, 21)
    binary.Write(w, binary.BigEndian, &m.ActionFlag)
    binary.Write(w, binary.BigEndian, &m.LimitPrice)
    binary.Write(w, binary.BigEndian, &m.VolumeChange)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.InvestUnitID, 17)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
}

func (m* CThostFtdcInputOrderActionField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.OrderActionRef)
    m.OrderRef = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    binary.Read(r, binary.BigEndian, &m.FrontID)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.ExchangeID = readLen(r, 9)
    m.OrderSysID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.ActionFlag)
    binary.Read(r, binary.BigEndian, &m.LimitPrice)
    binary.Read(r, binary.BigEndian, &m.VolumeChange)
    m.UserID = readLen(r, 16)
    m.InstrumentID = readLen(r, 31)
    m.InvestUnitID = readLen(r, 17)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
}

type  CThostFtdcOrderActionField struct {
    BrokerID string;
    InvestorID string;
    OrderActionRef int32;
    OrderRef string;
    RequestID int32;
    FrontID int32;
    SessionID int32;
    ExchangeID string;
    OrderSysID string;
    ActionFlag uint8;
    LimitPrice float64;
    VolumeChange int32;
    ActionDate string;
    ActionTime string;
    TraderID string;
    InstallID int32;
    OrderLocalID string;
    ActionLocalID string;
    ParticipantID string;
    ClientID string;
    BusinessUnit string;
    OrderActionStatus uint8;
    UserID string;
    StatusMsg string;
    InstrumentID string;
    BranchID string;
    InvestUnitID string;
    IPAddress string;
    MacAddress string;
}

func (m* CThostFtdcOrderActionField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.OrderActionRef)
    writeLen(w, m.OrderRef, 13)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    binary.Write(w, binary.BigEndian, &m.FrontID)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.OrderSysID, 21)
    binary.Write(w, binary.BigEndian, &m.ActionFlag)
    binary.Write(w, binary.BigEndian, &m.LimitPrice)
    binary.Write(w, binary.BigEndian, &m.VolumeChange)
    writeLen(w, m.ActionDate, 9)
    writeLen(w, m.ActionTime, 9)
    writeLen(w, m.TraderID, 21)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    writeLen(w, m.OrderLocalID, 13)
    writeLen(w, m.ActionLocalID, 13)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.BusinessUnit, 21)
    binary.Write(w, binary.BigEndian, &m.OrderActionStatus)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.StatusMsg, 81)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.BranchID, 9)
    writeLen(w, m.InvestUnitID, 17)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
}

func (m* CThostFtdcOrderActionField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.OrderActionRef)
    m.OrderRef = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    binary.Read(r, binary.BigEndian, &m.FrontID)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.ExchangeID = readLen(r, 9)
    m.OrderSysID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.ActionFlag)
    binary.Read(r, binary.BigEndian, &m.LimitPrice)
    binary.Read(r, binary.BigEndian, &m.VolumeChange)
    m.ActionDate = readLen(r, 9)
    m.ActionTime = readLen(r, 9)
    m.TraderID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    m.OrderLocalID = readLen(r, 13)
    m.ActionLocalID = readLen(r, 13)
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.BusinessUnit = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.OrderActionStatus)
    m.UserID = readLen(r, 16)
    m.StatusMsg = readLen(r, 81)
    m.InstrumentID = readLen(r, 31)
    m.BranchID = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
}

type  CThostFtdcExchangeOrderActionField struct {
    ExchangeID string;
    OrderSysID string;
    ActionFlag uint8;
    LimitPrice float64;
    VolumeChange int32;
    ActionDate string;
    ActionTime string;
    TraderID string;
    InstallID int32;
    OrderLocalID string;
    ActionLocalID string;
    ParticipantID string;
    ClientID string;
    BusinessUnit string;
    OrderActionStatus uint8;
    UserID string;
    BranchID string;
    IPAddress string;
    MacAddress string;
}

func (m* CThostFtdcExchangeOrderActionField) Marshal(w io.Writer) {
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.OrderSysID, 21)
    binary.Write(w, binary.BigEndian, &m.ActionFlag)
    binary.Write(w, binary.BigEndian, &m.LimitPrice)
    binary.Write(w, binary.BigEndian, &m.VolumeChange)
    writeLen(w, m.ActionDate, 9)
    writeLen(w, m.ActionTime, 9)
    writeLen(w, m.TraderID, 21)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    writeLen(w, m.OrderLocalID, 13)
    writeLen(w, m.ActionLocalID, 13)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.BusinessUnit, 21)
    binary.Write(w, binary.BigEndian, &m.OrderActionStatus)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.BranchID, 9)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
}

func (m* CThostFtdcExchangeOrderActionField) Unmarshal(r io.Reader) {
    m.ExchangeID = readLen(r, 9)
    m.OrderSysID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.ActionFlag)
    binary.Read(r, binary.BigEndian, &m.LimitPrice)
    binary.Read(r, binary.BigEndian, &m.VolumeChange)
    m.ActionDate = readLen(r, 9)
    m.ActionTime = readLen(r, 9)
    m.TraderID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    m.OrderLocalID = readLen(r, 13)
    m.ActionLocalID = readLen(r, 13)
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.BusinessUnit = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.OrderActionStatus)
    m.UserID = readLen(r, 16)
    m.BranchID = readLen(r, 9)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
}

type  CThostFtdcExchangeOrderActionErrorField struct {
    ExchangeID string;
    OrderSysID string;
    TraderID string;
    InstallID int32;
    OrderLocalID string;
    ActionLocalID string;
    ErrorID int32;
    ErrorMsg string;
    BrokerID string;
    CancelTime string;
}

func (m* CThostFtdcExchangeOrderActionErrorField) Marshal(w io.Writer) {
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.OrderSysID, 21)
    writeLen(w, m.TraderID, 21)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    writeLen(w, m.OrderLocalID, 13)
    writeLen(w, m.ActionLocalID, 13)
    binary.Write(w, binary.BigEndian, &m.ErrorID)
    writeLen(w, m.ErrorMsg, 81)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.CancelTime, 9)
}

func (m* CThostFtdcExchangeOrderActionErrorField) Unmarshal(r io.Reader) {
    m.ExchangeID = readLen(r, 9)
    m.OrderSysID = readLen(r, 21)
    m.TraderID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    m.OrderLocalID = readLen(r, 13)
    m.ActionLocalID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.ErrorID)
    m.ErrorMsg = readLen(r, 81)
    m.BrokerID = readLen(r, 11)
    m.CancelTime = readLen(r, 9)
}

type  CThostFtdcExchangeTradeField struct {
    ExchangeID string;
    TradeID string;
    Direction uint8;
    OrderSysID string;
    ParticipantID string;
    ClientID string;
    TradingRole uint8;
    ExchangeInstID string;
    OffsetFlag uint8;
    HedgeFlag uint8;
    Price float64;
    Volume int32;
    TradeDate string;
    TradeTime string;
    TradeType uint8;
    PriceSource uint8;
    TraderID string;
    OrderLocalID string;
    ClearingPartID string;
    BusinessUnit string;
    SequenceNo int32;
    TradeSource uint8;
}

func (m* CThostFtdcExchangeTradeField) Marshal(w io.Writer) {
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.TradeID, 21)
    binary.Write(w, binary.BigEndian, &m.Direction)
    writeLen(w, m.OrderSysID, 21)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    binary.Write(w, binary.BigEndian, &m.TradingRole)
    writeLen(w, m.ExchangeInstID, 31)
    binary.Write(w, binary.BigEndian, &m.OffsetFlag)
    binary.Write(w, binary.BigEndian, &m.HedgeFlag)
    binary.Write(w, binary.BigEndian, &m.Price)
    binary.Write(w, binary.BigEndian, &m.Volume)
    writeLen(w, m.TradeDate, 9)
    writeLen(w, m.TradeTime, 9)
    binary.Write(w, binary.BigEndian, &m.TradeType)
    binary.Write(w, binary.BigEndian, &m.PriceSource)
    writeLen(w, m.TraderID, 21)
    writeLen(w, m.OrderLocalID, 13)
    writeLen(w, m.ClearingPartID, 11)
    writeLen(w, m.BusinessUnit, 21)
    binary.Write(w, binary.BigEndian, &m.SequenceNo)
    binary.Write(w, binary.BigEndian, &m.TradeSource)
}

func (m* CThostFtdcExchangeTradeField) Unmarshal(r io.Reader) {
    m.ExchangeID = readLen(r, 9)
    m.TradeID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.Direction)
    m.OrderSysID = readLen(r, 21)
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    binary.Read(r, binary.BigEndian, &m.TradingRole)
    m.ExchangeInstID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.OffsetFlag)
    binary.Read(r, binary.BigEndian, &m.HedgeFlag)
    binary.Read(r, binary.BigEndian, &m.Price)
    binary.Read(r, binary.BigEndian, &m.Volume)
    m.TradeDate = readLen(r, 9)
    m.TradeTime = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.TradeType)
    binary.Read(r, binary.BigEndian, &m.PriceSource)
    m.TraderID = readLen(r, 21)
    m.OrderLocalID = readLen(r, 13)
    m.ClearingPartID = readLen(r, 11)
    m.BusinessUnit = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.SequenceNo)
    binary.Read(r, binary.BigEndian, &m.TradeSource)
}

type  CThostFtdcTradeField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    OrderRef string;
    UserID string;
    ExchangeID string;
    TradeID string;
    Direction uint8;
    OrderSysID string;
    ParticipantID string;
    ClientID string;
    TradingRole uint8;
    ExchangeInstID string;
    OffsetFlag uint8;
    HedgeFlag uint8;
    Price float64;
    Volume int32;
    TradeDate string;
    TradeTime string;
    TradeType uint8;
    PriceSource uint8;
    TraderID string;
    OrderLocalID string;
    ClearingPartID string;
    BusinessUnit string;
    SequenceNo int32;
    TradingDay string;
    SettlementID int32;
    BrokerOrderSeq int32;
    TradeSource uint8;
    InvestUnitID string;
}

func (m* CThostFtdcTradeField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.OrderRef, 13)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.TradeID, 21)
    binary.Write(w, binary.BigEndian, &m.Direction)
    writeLen(w, m.OrderSysID, 21)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    binary.Write(w, binary.BigEndian, &m.TradingRole)
    writeLen(w, m.ExchangeInstID, 31)
    binary.Write(w, binary.BigEndian, &m.OffsetFlag)
    binary.Write(w, binary.BigEndian, &m.HedgeFlag)
    binary.Write(w, binary.BigEndian, &m.Price)
    binary.Write(w, binary.BigEndian, &m.Volume)
    writeLen(w, m.TradeDate, 9)
    writeLen(w, m.TradeTime, 9)
    binary.Write(w, binary.BigEndian, &m.TradeType)
    binary.Write(w, binary.BigEndian, &m.PriceSource)
    writeLen(w, m.TraderID, 21)
    writeLen(w, m.OrderLocalID, 13)
    writeLen(w, m.ClearingPartID, 11)
    writeLen(w, m.BusinessUnit, 21)
    binary.Write(w, binary.BigEndian, &m.SequenceNo)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.SettlementID)
    binary.Write(w, binary.BigEndian, &m.BrokerOrderSeq)
    binary.Write(w, binary.BigEndian, &m.TradeSource)
    writeLen(w, m.InvestUnitID, 17)
}

func (m* CThostFtdcTradeField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.OrderRef = readLen(r, 13)
    m.UserID = readLen(r, 16)
    m.ExchangeID = readLen(r, 9)
    m.TradeID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.Direction)
    m.OrderSysID = readLen(r, 21)
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    binary.Read(r, binary.BigEndian, &m.TradingRole)
    m.ExchangeInstID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.OffsetFlag)
    binary.Read(r, binary.BigEndian, &m.HedgeFlag)
    binary.Read(r, binary.BigEndian, &m.Price)
    binary.Read(r, binary.BigEndian, &m.Volume)
    m.TradeDate = readLen(r, 9)
    m.TradeTime = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.TradeType)
    binary.Read(r, binary.BigEndian, &m.PriceSource)
    m.TraderID = readLen(r, 21)
    m.OrderLocalID = readLen(r, 13)
    m.ClearingPartID = readLen(r, 11)
    m.BusinessUnit = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.SequenceNo)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.SettlementID)
    binary.Read(r, binary.BigEndian, &m.BrokerOrderSeq)
    binary.Read(r, binary.BigEndian, &m.TradeSource)
    m.InvestUnitID = readLen(r, 17)
}

type  CThostFtdcUserSessionField struct {
    FrontID int32;
    SessionID int32;
    BrokerID string;
    UserID string;
    LoginDate string;
    LoginTime string;
    IPAddress string;
    UserProductInfo string;
    InterfaceProductInfo string;
    ProtocolInfo string;
    MacAddress string;
    LoginRemark string;
}

func (m* CThostFtdcUserSessionField) Marshal(w io.Writer) {
    binary.Write(w, binary.BigEndian, &m.FrontID)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.LoginDate, 9)
    writeLen(w, m.LoginTime, 9)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.UserProductInfo, 11)
    writeLen(w, m.InterfaceProductInfo, 11)
    writeLen(w, m.ProtocolInfo, 11)
    writeLen(w, m.MacAddress, 21)
    writeLen(w, m.LoginRemark, 36)
}

func (m* CThostFtdcUserSessionField) Unmarshal(r io.Reader) {
    binary.Read(r, binary.BigEndian, &m.FrontID)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.BrokerID = readLen(r, 11)
    m.UserID = readLen(r, 16)
    m.LoginDate = readLen(r, 9)
    m.LoginTime = readLen(r, 9)
    m.IPAddress = readLen(r, 16)
    m.UserProductInfo = readLen(r, 11)
    m.InterfaceProductInfo = readLen(r, 11)
    m.ProtocolInfo = readLen(r, 11)
    m.MacAddress = readLen(r, 21)
    m.LoginRemark = readLen(r, 36)
}

type  CThostFtdcQueryMaxOrderVolumeField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    Direction uint8;
    OffsetFlag uint8;
    HedgeFlag uint8;
    MaxVolume int32;
    ExchangeID string;
    InvestUnitID string;
}

func (m* CThostFtdcQueryMaxOrderVolumeField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.Direction)
    binary.Write(w, binary.BigEndian, &m.OffsetFlag)
    binary.Write(w, binary.BigEndian, &m.HedgeFlag)
    binary.Write(w, binary.BigEndian, &m.MaxVolume)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.InvestUnitID, 17)
}

func (m* CThostFtdcQueryMaxOrderVolumeField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.Direction)
    binary.Read(r, binary.BigEndian, &m.OffsetFlag)
    binary.Read(r, binary.BigEndian, &m.HedgeFlag)
    binary.Read(r, binary.BigEndian, &m.MaxVolume)
    m.ExchangeID = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
}

type  CThostFtdcSettlementInfoConfirmField struct {
    BrokerID string;
    InvestorID string;
    ConfirmDate string;
    ConfirmTime string;
    SettlementID int32;
    AccountID string;
    CurrencyID string;
}

func (m* CThostFtdcSettlementInfoConfirmField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.ConfirmDate, 9)
    writeLen(w, m.ConfirmTime, 9)
    binary.Write(w, binary.BigEndian, &m.SettlementID)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.CurrencyID, 4)
}

func (m* CThostFtdcSettlementInfoConfirmField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.ConfirmDate = readLen(r, 9)
    m.ConfirmTime = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.SettlementID)
    m.AccountID = readLen(r, 13)
    m.CurrencyID = readLen(r, 4)
}

type  CThostFtdcSyncDepositField struct {
    DepositSeqNo string;
    BrokerID string;
    InvestorID string;
    Deposit float64;
    IsForce int32;
    CurrencyID string;
    BizType uint8;
}

func (m* CThostFtdcSyncDepositField) Marshal(w io.Writer) {
    writeLen(w, m.DepositSeqNo, 15)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.Deposit)
    binary.Write(w, binary.BigEndian, &m.IsForce)
    writeLen(w, m.CurrencyID, 4)
    binary.Write(w, binary.BigEndian, &m.BizType)
}

func (m* CThostFtdcSyncDepositField) Unmarshal(r io.Reader) {
    m.DepositSeqNo = readLen(r, 15)
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.Deposit)
    binary.Read(r, binary.BigEndian, &m.IsForce)
    m.CurrencyID = readLen(r, 4)
    binary.Read(r, binary.BigEndian, &m.BizType)
}

type  CThostFtdcSyncFundMortgageField struct {
    MortgageSeqNo string;
    BrokerID string;
    InvestorID string;
    FromCurrencyID string;
    MortgageAmount float64;
    ToCurrencyID string;
}

func (m* CThostFtdcSyncFundMortgageField) Marshal(w io.Writer) {
    writeLen(w, m.MortgageSeqNo, 15)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.FromCurrencyID, 4)
    binary.Write(w, binary.BigEndian, &m.MortgageAmount)
    writeLen(w, m.ToCurrencyID, 4)
}

func (m* CThostFtdcSyncFundMortgageField) Unmarshal(r io.Reader) {
    m.MortgageSeqNo = readLen(r, 15)
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.FromCurrencyID = readLen(r, 4)
    binary.Read(r, binary.BigEndian, &m.MortgageAmount)
    m.ToCurrencyID = readLen(r, 4)
}

type  CThostFtdcBrokerSyncField struct {
    BrokerID string;
}

func (m* CThostFtdcBrokerSyncField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
}

func (m* CThostFtdcBrokerSyncField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
}

type  CThostFtdcSyncingInvestorField struct {
    InvestorID string;
    BrokerID string;
    InvestorGroupID string;
    InvestorName string;
    IdentifiedCardType uint8;
    IdentifiedCardNo string;
    IsActive int32;
    Telephone string;
    Address string;
    OpenDate string;
    Mobile string;
    CommModelID string;
    MarginModelID string;
}

func (m* CThostFtdcSyncingInvestorField) Marshal(w io.Writer) {
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorGroupID, 13)
    writeLen(w, m.InvestorName, 81)
    binary.Write(w, binary.BigEndian, &m.IdentifiedCardType)
    writeLen(w, m.IdentifiedCardNo, 51)
    binary.Write(w, binary.BigEndian, &m.IsActive)
    writeLen(w, m.Telephone, 41)
    writeLen(w, m.Address, 101)
    writeLen(w, m.OpenDate, 9)
    writeLen(w, m.Mobile, 41)
    writeLen(w, m.CommModelID, 13)
    writeLen(w, m.MarginModelID, 13)
}

func (m* CThostFtdcSyncingInvestorField) Unmarshal(r io.Reader) {
    m.InvestorID = readLen(r, 13)
    m.BrokerID = readLen(r, 11)
    m.InvestorGroupID = readLen(r, 13)
    m.InvestorName = readLen(r, 81)
    binary.Read(r, binary.BigEndian, &m.IdentifiedCardType)
    m.IdentifiedCardNo = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.IsActive)
    m.Telephone = readLen(r, 41)
    m.Address = readLen(r, 101)
    m.OpenDate = readLen(r, 9)
    m.Mobile = readLen(r, 41)
    m.CommModelID = readLen(r, 13)
    m.MarginModelID = readLen(r, 13)
}

type  CThostFtdcSyncingTradingCodeField struct {
    InvestorID string;
    BrokerID string;
    ExchangeID string;
    ClientID string;
    IsActive int32;
    ClientIDType string;
    BranchID string;
}

func (m* CThostFtdcSyncingTradingCodeField) Marshal(w io.Writer) {
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ClientID, 11)
    binary.Write(w, binary.BigEndian, &m.IsActive)
    writeLen(w, m.ClientIDType, 11)
    writeLen(w, m.BranchID, 9)
}

func (m* CThostFtdcSyncingTradingCodeField) Unmarshal(r io.Reader) {
    m.InvestorID = readLen(r, 13)
    m.BrokerID = readLen(r, 11)
    m.ExchangeID = readLen(r, 9)
    m.ClientID = readLen(r, 11)
    binary.Read(r, binary.BigEndian, &m.IsActive)
    m.ClientIDType = readLen(r, 11)
    m.BranchID = readLen(r, 9)
}

type  CThostFtdcSyncingInvestorGroupField struct {
    BrokerID string;
    InvestorGroupID string;
    InvestorGroupName string;
}

func (m* CThostFtdcSyncingInvestorGroupField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorGroupID, 13)
    writeLen(w, m.InvestorGroupName, 41)
}

func (m* CThostFtdcSyncingInvestorGroupField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorGroupID = readLen(r, 13)
    m.InvestorGroupName = readLen(r, 41)
}

type  CThostFtdcSyncingTradingAccountField struct {
    BrokerID string;
    AccountID string;
    PreMortgage float64;
    PreCredit float64;
    PreDeposit float64;
    PreBalance float64;
    PreMargin float64;
    InterestBase float64;
    Interest float64;
    Deposit float64;
    Withdraw float64;
    FrozenMargin float64;
    FrozenCash float64;
    FrozenCommission float64;
    CurrMargin float64;
    CashIn float64;
    Commission float64;
    CloseProfit float64;
    PositionProfit float64;
    Balance float64;
    Available float64;
    WithdrawQuota float64;
    Reserve float64;
    TradingDay string;
    SettlementID int32;
    Credit float64;
    Mortgage float64;
    ExchangeMargin float64;
    DeliveryMargin float64;
    ExchangeDeliveryMargin float64;
    ReserveBalance float64;
    CurrencyID string;
    PreFundMortgageIn float64;
    PreFundMortgageOut float64;
    FundMortgageIn float64;
    FundMortgageOut float64;
    FundMortgageAvailable float64;
    MortgageableFund float64;
    SpecProductMargin float64;
    SpecProductFrozenMargin float64;
    SpecProductCommission float64;
    SpecProductFrozenCommission float64;
    SpecProductPositionProfit float64;
    SpecProductCloseProfit float64;
    SpecProductPositionProfitByAlg float64;
    SpecProductExchangeMargin float64;
    FrozenSwap float64;
    RemainSwap float64;
}

func (m* CThostFtdcSyncingTradingAccountField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.AccountID, 13)
    binary.Write(w, binary.BigEndian, &m.PreMortgage)
    binary.Write(w, binary.BigEndian, &m.PreCredit)
    binary.Write(w, binary.BigEndian, &m.PreDeposit)
    binary.Write(w, binary.BigEndian, &m.PreBalance)
    binary.Write(w, binary.BigEndian, &m.PreMargin)
    binary.Write(w, binary.BigEndian, &m.InterestBase)
    binary.Write(w, binary.BigEndian, &m.Interest)
    binary.Write(w, binary.BigEndian, &m.Deposit)
    binary.Write(w, binary.BigEndian, &m.Withdraw)
    binary.Write(w, binary.BigEndian, &m.FrozenMargin)
    binary.Write(w, binary.BigEndian, &m.FrozenCash)
    binary.Write(w, binary.BigEndian, &m.FrozenCommission)
    binary.Write(w, binary.BigEndian, &m.CurrMargin)
    binary.Write(w, binary.BigEndian, &m.CashIn)
    binary.Write(w, binary.BigEndian, &m.Commission)
    binary.Write(w, binary.BigEndian, &m.CloseProfit)
    binary.Write(w, binary.BigEndian, &m.PositionProfit)
    binary.Write(w, binary.BigEndian, &m.Balance)
    binary.Write(w, binary.BigEndian, &m.Available)
    binary.Write(w, binary.BigEndian, &m.WithdrawQuota)
    binary.Write(w, binary.BigEndian, &m.Reserve)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.SettlementID)
    binary.Write(w, binary.BigEndian, &m.Credit)
    binary.Write(w, binary.BigEndian, &m.Mortgage)
    binary.Write(w, binary.BigEndian, &m.ExchangeMargin)
    binary.Write(w, binary.BigEndian, &m.DeliveryMargin)
    binary.Write(w, binary.BigEndian, &m.ExchangeDeliveryMargin)
    binary.Write(w, binary.BigEndian, &m.ReserveBalance)
    writeLen(w, m.CurrencyID, 4)
    binary.Write(w, binary.BigEndian, &m.PreFundMortgageIn)
    binary.Write(w, binary.BigEndian, &m.PreFundMortgageOut)
    binary.Write(w, binary.BigEndian, &m.FundMortgageIn)
    binary.Write(w, binary.BigEndian, &m.FundMortgageOut)
    binary.Write(w, binary.BigEndian, &m.FundMortgageAvailable)
    binary.Write(w, binary.BigEndian, &m.MortgageableFund)
    binary.Write(w, binary.BigEndian, &m.SpecProductMargin)
    binary.Write(w, binary.BigEndian, &m.SpecProductFrozenMargin)
    binary.Write(w, binary.BigEndian, &m.SpecProductCommission)
    binary.Write(w, binary.BigEndian, &m.SpecProductFrozenCommission)
    binary.Write(w, binary.BigEndian, &m.SpecProductPositionProfit)
    binary.Write(w, binary.BigEndian, &m.SpecProductCloseProfit)
    binary.Write(w, binary.BigEndian, &m.SpecProductPositionProfitByAlg)
    binary.Write(w, binary.BigEndian, &m.SpecProductExchangeMargin)
    binary.Write(w, binary.BigEndian, &m.FrozenSwap)
    binary.Write(w, binary.BigEndian, &m.RemainSwap)
}

func (m* CThostFtdcSyncingTradingAccountField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.AccountID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.PreMortgage)
    binary.Read(r, binary.BigEndian, &m.PreCredit)
    binary.Read(r, binary.BigEndian, &m.PreDeposit)
    binary.Read(r, binary.BigEndian, &m.PreBalance)
    binary.Read(r, binary.BigEndian, &m.PreMargin)
    binary.Read(r, binary.BigEndian, &m.InterestBase)
    binary.Read(r, binary.BigEndian, &m.Interest)
    binary.Read(r, binary.BigEndian, &m.Deposit)
    binary.Read(r, binary.BigEndian, &m.Withdraw)
    binary.Read(r, binary.BigEndian, &m.FrozenMargin)
    binary.Read(r, binary.BigEndian, &m.FrozenCash)
    binary.Read(r, binary.BigEndian, &m.FrozenCommission)
    binary.Read(r, binary.BigEndian, &m.CurrMargin)
    binary.Read(r, binary.BigEndian, &m.CashIn)
    binary.Read(r, binary.BigEndian, &m.Commission)
    binary.Read(r, binary.BigEndian, &m.CloseProfit)
    binary.Read(r, binary.BigEndian, &m.PositionProfit)
    binary.Read(r, binary.BigEndian, &m.Balance)
    binary.Read(r, binary.BigEndian, &m.Available)
    binary.Read(r, binary.BigEndian, &m.WithdrawQuota)
    binary.Read(r, binary.BigEndian, &m.Reserve)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.SettlementID)
    binary.Read(r, binary.BigEndian, &m.Credit)
    binary.Read(r, binary.BigEndian, &m.Mortgage)
    binary.Read(r, binary.BigEndian, &m.ExchangeMargin)
    binary.Read(r, binary.BigEndian, &m.DeliveryMargin)
    binary.Read(r, binary.BigEndian, &m.ExchangeDeliveryMargin)
    binary.Read(r, binary.BigEndian, &m.ReserveBalance)
    m.CurrencyID = readLen(r, 4)
    binary.Read(r, binary.BigEndian, &m.PreFundMortgageIn)
    binary.Read(r, binary.BigEndian, &m.PreFundMortgageOut)
    binary.Read(r, binary.BigEndian, &m.FundMortgageIn)
    binary.Read(r, binary.BigEndian, &m.FundMortgageOut)
    binary.Read(r, binary.BigEndian, &m.FundMortgageAvailable)
    binary.Read(r, binary.BigEndian, &m.MortgageableFund)
    binary.Read(r, binary.BigEndian, &m.SpecProductMargin)
    binary.Read(r, binary.BigEndian, &m.SpecProductFrozenMargin)
    binary.Read(r, binary.BigEndian, &m.SpecProductCommission)
    binary.Read(r, binary.BigEndian, &m.SpecProductFrozenCommission)
    binary.Read(r, binary.BigEndian, &m.SpecProductPositionProfit)
    binary.Read(r, binary.BigEndian, &m.SpecProductCloseProfit)
    binary.Read(r, binary.BigEndian, &m.SpecProductPositionProfitByAlg)
    binary.Read(r, binary.BigEndian, &m.SpecProductExchangeMargin)
    binary.Read(r, binary.BigEndian, &m.FrozenSwap)
    binary.Read(r, binary.BigEndian, &m.RemainSwap)
}

type  CThostFtdcSyncingInvestorPositionField struct {
    InstrumentID string;
    BrokerID string;
    InvestorID string;
    PosiDirection uint8;
    HedgeFlag uint8;
    PositionDate uint8;
    YdPosition int32;
    Position int32;
    LongFrozen int32;
    ShortFrozen int32;
    LongFrozenAmount float64;
    ShortFrozenAmount float64;
    OpenVolume int32;
    CloseVolume int32;
    OpenAmount float64;
    CloseAmount float64;
    PositionCost float64;
    PreMargin float64;
    UseMargin float64;
    FrozenMargin float64;
    FrozenCash float64;
    FrozenCommission float64;
    CashIn float64;
    Commission float64;
    CloseProfit float64;
    PositionProfit float64;
    PreSettlementPrice float64;
    SettlementPrice float64;
    TradingDay string;
    SettlementID int32;
    OpenCost float64;
    ExchangeMargin float64;
    CombPosition int32;
    CombLongFrozen int32;
    CombShortFrozen int32;
    CloseProfitByDate float64;
    CloseProfitByTrade float64;
    TodayPosition int32;
    MarginRateByMoney float64;
    MarginRateByVolume float64;
    StrikeFrozen int32;
    StrikeFrozenAmount float64;
    AbandonFrozen int32;
    ExchangeID string;
    YdStrikeFrozen int32;
    InvestUnitID string;
}

func (m* CThostFtdcSyncingInvestorPositionField) Marshal(w io.Writer) {
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.PosiDirection)
    binary.Write(w, binary.BigEndian, &m.HedgeFlag)
    binary.Write(w, binary.BigEndian, &m.PositionDate)
    binary.Write(w, binary.BigEndian, &m.YdPosition)
    binary.Write(w, binary.BigEndian, &m.Position)
    binary.Write(w, binary.BigEndian, &m.LongFrozen)
    binary.Write(w, binary.BigEndian, &m.ShortFrozen)
    binary.Write(w, binary.BigEndian, &m.LongFrozenAmount)
    binary.Write(w, binary.BigEndian, &m.ShortFrozenAmount)
    binary.Write(w, binary.BigEndian, &m.OpenVolume)
    binary.Write(w, binary.BigEndian, &m.CloseVolume)
    binary.Write(w, binary.BigEndian, &m.OpenAmount)
    binary.Write(w, binary.BigEndian, &m.CloseAmount)
    binary.Write(w, binary.BigEndian, &m.PositionCost)
    binary.Write(w, binary.BigEndian, &m.PreMargin)
    binary.Write(w, binary.BigEndian, &m.UseMargin)
    binary.Write(w, binary.BigEndian, &m.FrozenMargin)
    binary.Write(w, binary.BigEndian, &m.FrozenCash)
    binary.Write(w, binary.BigEndian, &m.FrozenCommission)
    binary.Write(w, binary.BigEndian, &m.CashIn)
    binary.Write(w, binary.BigEndian, &m.Commission)
    binary.Write(w, binary.BigEndian, &m.CloseProfit)
    binary.Write(w, binary.BigEndian, &m.PositionProfit)
    binary.Write(w, binary.BigEndian, &m.PreSettlementPrice)
    binary.Write(w, binary.BigEndian, &m.SettlementPrice)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.SettlementID)
    binary.Write(w, binary.BigEndian, &m.OpenCost)
    binary.Write(w, binary.BigEndian, &m.ExchangeMargin)
    binary.Write(w, binary.BigEndian, &m.CombPosition)
    binary.Write(w, binary.BigEndian, &m.CombLongFrozen)
    binary.Write(w, binary.BigEndian, &m.CombShortFrozen)
    binary.Write(w, binary.BigEndian, &m.CloseProfitByDate)
    binary.Write(w, binary.BigEndian, &m.CloseProfitByTrade)
    binary.Write(w, binary.BigEndian, &m.TodayPosition)
    binary.Write(w, binary.BigEndian, &m.MarginRateByMoney)
    binary.Write(w, binary.BigEndian, &m.MarginRateByVolume)
    binary.Write(w, binary.BigEndian, &m.StrikeFrozen)
    binary.Write(w, binary.BigEndian, &m.StrikeFrozenAmount)
    binary.Write(w, binary.BigEndian, &m.AbandonFrozen)
    writeLen(w, m.ExchangeID, 9)
    binary.Write(w, binary.BigEndian, &m.YdStrikeFrozen)
    writeLen(w, m.InvestUnitID, 17)
}

func (m* CThostFtdcSyncingInvestorPositionField) Unmarshal(r io.Reader) {
    m.InstrumentID = readLen(r, 31)
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.PosiDirection)
    binary.Read(r, binary.BigEndian, &m.HedgeFlag)
    binary.Read(r, binary.BigEndian, &m.PositionDate)
    binary.Read(r, binary.BigEndian, &m.YdPosition)
    binary.Read(r, binary.BigEndian, &m.Position)
    binary.Read(r, binary.BigEndian, &m.LongFrozen)
    binary.Read(r, binary.BigEndian, &m.ShortFrozen)
    binary.Read(r, binary.BigEndian, &m.LongFrozenAmount)
    binary.Read(r, binary.BigEndian, &m.ShortFrozenAmount)
    binary.Read(r, binary.BigEndian, &m.OpenVolume)
    binary.Read(r, binary.BigEndian, &m.CloseVolume)
    binary.Read(r, binary.BigEndian, &m.OpenAmount)
    binary.Read(r, binary.BigEndian, &m.CloseAmount)
    binary.Read(r, binary.BigEndian, &m.PositionCost)
    binary.Read(r, binary.BigEndian, &m.PreMargin)
    binary.Read(r, binary.BigEndian, &m.UseMargin)
    binary.Read(r, binary.BigEndian, &m.FrozenMargin)
    binary.Read(r, binary.BigEndian, &m.FrozenCash)
    binary.Read(r, binary.BigEndian, &m.FrozenCommission)
    binary.Read(r, binary.BigEndian, &m.CashIn)
    binary.Read(r, binary.BigEndian, &m.Commission)
    binary.Read(r, binary.BigEndian, &m.CloseProfit)
    binary.Read(r, binary.BigEndian, &m.PositionProfit)
    binary.Read(r, binary.BigEndian, &m.PreSettlementPrice)
    binary.Read(r, binary.BigEndian, &m.SettlementPrice)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.SettlementID)
    binary.Read(r, binary.BigEndian, &m.OpenCost)
    binary.Read(r, binary.BigEndian, &m.ExchangeMargin)
    binary.Read(r, binary.BigEndian, &m.CombPosition)
    binary.Read(r, binary.BigEndian, &m.CombLongFrozen)
    binary.Read(r, binary.BigEndian, &m.CombShortFrozen)
    binary.Read(r, binary.BigEndian, &m.CloseProfitByDate)
    binary.Read(r, binary.BigEndian, &m.CloseProfitByTrade)
    binary.Read(r, binary.BigEndian, &m.TodayPosition)
    binary.Read(r, binary.BigEndian, &m.MarginRateByMoney)
    binary.Read(r, binary.BigEndian, &m.MarginRateByVolume)
    binary.Read(r, binary.BigEndian, &m.StrikeFrozen)
    binary.Read(r, binary.BigEndian, &m.StrikeFrozenAmount)
    binary.Read(r, binary.BigEndian, &m.AbandonFrozen)
    m.ExchangeID = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.YdStrikeFrozen)
    m.InvestUnitID = readLen(r, 17)
}

type  CThostFtdcSyncingInstrumentMarginRateField struct {
    InstrumentID string;
    InvestorRange uint8;
    BrokerID string;
    InvestorID string;
    HedgeFlag uint8;
    LongMarginRatioByMoney float64;
    LongMarginRatioByVolume float64;
    ShortMarginRatioByMoney float64;
    ShortMarginRatioByVolume float64;
    IsRelative int32;
}

func (m* CThostFtdcSyncingInstrumentMarginRateField) Marshal(w io.Writer) {
    writeLen(w, m.InstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.InvestorRange)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.HedgeFlag)
    binary.Write(w, binary.BigEndian, &m.LongMarginRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.LongMarginRatioByVolume)
    binary.Write(w, binary.BigEndian, &m.ShortMarginRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.ShortMarginRatioByVolume)
    binary.Write(w, binary.BigEndian, &m.IsRelative)
}

func (m* CThostFtdcSyncingInstrumentMarginRateField) Unmarshal(r io.Reader) {
    m.InstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.InvestorRange)
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.HedgeFlag)
    binary.Read(r, binary.BigEndian, &m.LongMarginRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.LongMarginRatioByVolume)
    binary.Read(r, binary.BigEndian, &m.ShortMarginRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.ShortMarginRatioByVolume)
    binary.Read(r, binary.BigEndian, &m.IsRelative)
}

type  CThostFtdcSyncingInstrumentCommissionRateField struct {
    InstrumentID string;
    InvestorRange uint8;
    BrokerID string;
    InvestorID string;
    OpenRatioByMoney float64;
    OpenRatioByVolume float64;
    CloseRatioByMoney float64;
    CloseRatioByVolume float64;
    CloseTodayRatioByMoney float64;
    CloseTodayRatioByVolume float64;
    ExchangeID string;
}

func (m* CThostFtdcSyncingInstrumentCommissionRateField) Marshal(w io.Writer) {
    writeLen(w, m.InstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.InvestorRange)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.OpenRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.OpenRatioByVolume)
    binary.Write(w, binary.BigEndian, &m.CloseRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.CloseRatioByVolume)
    binary.Write(w, binary.BigEndian, &m.CloseTodayRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.CloseTodayRatioByVolume)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcSyncingInstrumentCommissionRateField) Unmarshal(r io.Reader) {
    m.InstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.InvestorRange)
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.OpenRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.OpenRatioByVolume)
    binary.Read(r, binary.BigEndian, &m.CloseRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.CloseRatioByVolume)
    binary.Read(r, binary.BigEndian, &m.CloseTodayRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.CloseTodayRatioByVolume)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcSyncingInstrumentTradingRightField struct {
    InstrumentID string;
    InvestorRange uint8;
    BrokerID string;
    InvestorID string;
    TradingRight uint8;
    ExchangeID string;
}

func (m* CThostFtdcSyncingInstrumentTradingRightField) Marshal(w io.Writer) {
    writeLen(w, m.InstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.InvestorRange)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.TradingRight)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcSyncingInstrumentTradingRightField) Unmarshal(r io.Reader) {
    m.InstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.InvestorRange)
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.TradingRight)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcQryOrderField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    ExchangeID string;
    OrderSysID string;
    InsertTimeStart string;
    InsertTimeEnd string;
    InvestUnitID string;
}

func (m* CThostFtdcQryOrderField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.OrderSysID, 21)
    writeLen(w, m.InsertTimeStart, 9)
    writeLen(w, m.InsertTimeEnd, 9)
    writeLen(w, m.InvestUnitID, 17)
}

func (m* CThostFtdcQryOrderField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
    m.OrderSysID = readLen(r, 21)
    m.InsertTimeStart = readLen(r, 9)
    m.InsertTimeEnd = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
}

type  CThostFtdcQryTradeField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    ExchangeID string;
    TradeID string;
    TradeTimeStart string;
    TradeTimeEnd string;
    InvestUnitID string;
}

func (m* CThostFtdcQryTradeField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.TradeID, 21)
    writeLen(w, m.TradeTimeStart, 9)
    writeLen(w, m.TradeTimeEnd, 9)
    writeLen(w, m.InvestUnitID, 17)
}

func (m* CThostFtdcQryTradeField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
    m.TradeID = readLen(r, 21)
    m.TradeTimeStart = readLen(r, 9)
    m.TradeTimeEnd = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
}

type  CThostFtdcQryInvestorPositionField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    ExchangeID string;
    InvestUnitID string;
}

func (m* CThostFtdcQryInvestorPositionField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.InvestUnitID, 17)
}

func (m* CThostFtdcQryInvestorPositionField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
}

type  CThostFtdcQryTradingAccountField struct {
    BrokerID string;
    InvestorID string;
    CurrencyID string;
    BizType uint8;
    AccountID string;
}

func (m* CThostFtdcQryTradingAccountField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.CurrencyID, 4)
    binary.Write(w, binary.BigEndian, &m.BizType)
    writeLen(w, m.AccountID, 13)
}

func (m* CThostFtdcQryTradingAccountField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.CurrencyID = readLen(r, 4)
    binary.Read(r, binary.BigEndian, &m.BizType)
    m.AccountID = readLen(r, 13)
}

type  CThostFtdcQryInvestorField struct {
    BrokerID string;
    InvestorID string;
}

func (m* CThostFtdcQryInvestorField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
}

func (m* CThostFtdcQryInvestorField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
}

type  CThostFtdcQryTradingCodeField struct {
    BrokerID string;
    InvestorID string;
    ExchangeID string;
    ClientID string;
    ClientIDType string;
    InvestUnitID string;
}

func (m* CThostFtdcQryTradingCodeField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.ClientIDType, 11)
    writeLen(w, m.InvestUnitID, 17)
}

func (m* CThostFtdcQryTradingCodeField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.ExchangeID = readLen(r, 9)
    m.ClientID = readLen(r, 11)
    m.ClientIDType = readLen(r, 11)
    m.InvestUnitID = readLen(r, 17)
}

type  CThostFtdcQryInvestorGroupField struct {
    BrokerID string;
}

func (m* CThostFtdcQryInvestorGroupField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
}

func (m* CThostFtdcQryInvestorGroupField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
}

type  CThostFtdcQryInstrumentMarginRateField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    HedgeFlag uint8;
    ExchangeID string;
    InvestUnitID string;
}

func (m* CThostFtdcQryInstrumentMarginRateField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.HedgeFlag)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.InvestUnitID, 17)
}

func (m* CThostFtdcQryInstrumentMarginRateField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.HedgeFlag)
    m.ExchangeID = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
}

type  CThostFtdcQryInstrumentCommissionRateField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    ExchangeID string;
    InvestUnitID string;
}

func (m* CThostFtdcQryInstrumentCommissionRateField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.InvestUnitID, 17)
}

func (m* CThostFtdcQryInstrumentCommissionRateField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
}

type  CThostFtdcQryInstrumentTradingRightField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    ExchangeID string;
}

func (m* CThostFtdcQryInstrumentTradingRightField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcQryInstrumentTradingRightField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcQryBrokerField struct {
    BrokerID string;
}

func (m* CThostFtdcQryBrokerField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
}

func (m* CThostFtdcQryBrokerField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
}

type  CThostFtdcQryTraderField struct {
    ExchangeID string;
    ParticipantID string;
    TraderID string;
}

func (m* CThostFtdcQryTraderField) Marshal(w io.Writer) {
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.TraderID, 21)
}

func (m* CThostFtdcQryTraderField) Unmarshal(r io.Reader) {
    m.ExchangeID = readLen(r, 9)
    m.ParticipantID = readLen(r, 11)
    m.TraderID = readLen(r, 21)
}

type  CThostFtdcQrySuperUserFunctionField struct {
    UserID string;
}

func (m* CThostFtdcQrySuperUserFunctionField) Marshal(w io.Writer) {
    writeLen(w, m.UserID, 16)
}

func (m* CThostFtdcQrySuperUserFunctionField) Unmarshal(r io.Reader) {
    m.UserID = readLen(r, 16)
}

type  CThostFtdcQryUserSessionField struct {
    FrontID int32;
    SessionID int32;
    BrokerID string;
    UserID string;
}

func (m* CThostFtdcQryUserSessionField) Marshal(w io.Writer) {
    binary.Write(w, binary.BigEndian, &m.FrontID)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.UserID, 16)
}

func (m* CThostFtdcQryUserSessionField) Unmarshal(r io.Reader) {
    binary.Read(r, binary.BigEndian, &m.FrontID)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.BrokerID = readLen(r, 11)
    m.UserID = readLen(r, 16)
}

type  CThostFtdcQryPartBrokerField struct {
    ExchangeID string;
    BrokerID string;
    ParticipantID string;
}

func (m* CThostFtdcQryPartBrokerField) Marshal(w io.Writer) {
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.ParticipantID, 11)
}

func (m* CThostFtdcQryPartBrokerField) Unmarshal(r io.Reader) {
    m.ExchangeID = readLen(r, 9)
    m.BrokerID = readLen(r, 11)
    m.ParticipantID = readLen(r, 11)
}

type  CThostFtdcQryFrontStatusField struct {
    FrontID int32;
}

func (m* CThostFtdcQryFrontStatusField) Marshal(w io.Writer) {
    binary.Write(w, binary.BigEndian, &m.FrontID)
}

func (m* CThostFtdcQryFrontStatusField) Unmarshal(r io.Reader) {
    binary.Read(r, binary.BigEndian, &m.FrontID)
}

type  CThostFtdcQryExchangeOrderField struct {
    ParticipantID string;
    ClientID string;
    ExchangeInstID string;
    ExchangeID string;
    TraderID string;
}

func (m* CThostFtdcQryExchangeOrderField) Marshal(w io.Writer) {
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.ExchangeInstID, 31)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.TraderID, 21)
}

func (m* CThostFtdcQryExchangeOrderField) Unmarshal(r io.Reader) {
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.ExchangeInstID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
    m.TraderID = readLen(r, 21)
}

type  CThostFtdcQryOrderActionField struct {
    BrokerID string;
    InvestorID string;
    ExchangeID string;
}

func (m* CThostFtdcQryOrderActionField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcQryOrderActionField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcQryExchangeOrderActionField struct {
    ParticipantID string;
    ClientID string;
    ExchangeID string;
    TraderID string;
}

func (m* CThostFtdcQryExchangeOrderActionField) Marshal(w io.Writer) {
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.TraderID, 21)
}

func (m* CThostFtdcQryExchangeOrderActionField) Unmarshal(r io.Reader) {
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.ExchangeID = readLen(r, 9)
    m.TraderID = readLen(r, 21)
}

type  CThostFtdcQrySuperUserField struct {
    UserID string;
}

func (m* CThostFtdcQrySuperUserField) Marshal(w io.Writer) {
    writeLen(w, m.UserID, 16)
}

func (m* CThostFtdcQrySuperUserField) Unmarshal(r io.Reader) {
    m.UserID = readLen(r, 16)
}

type  CThostFtdcQryExchangeField struct {
    ExchangeID string;
}

func (m* CThostFtdcQryExchangeField) Marshal(w io.Writer) {
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcQryExchangeField) Unmarshal(r io.Reader) {
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcQryProductField struct {
    ProductID string;
    ProductClass uint8;
    ExchangeID string;
}

func (m* CThostFtdcQryProductField) Marshal(w io.Writer) {
    writeLen(w, m.ProductID, 31)
    binary.Write(w, binary.BigEndian, &m.ProductClass)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcQryProductField) Unmarshal(r io.Reader) {
    m.ProductID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.ProductClass)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcQryInstrumentField struct {
    InstrumentID string;
    ExchangeID string;
    ExchangeInstID string;
    ProductID string;
}

func (m* CThostFtdcQryInstrumentField) Marshal(w io.Writer) {
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ExchangeInstID, 31)
    writeLen(w, m.ProductID, 31)
}

func (m* CThostFtdcQryInstrumentField) Unmarshal(r io.Reader) {
    m.InstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
    m.ExchangeInstID = readLen(r, 31)
    m.ProductID = readLen(r, 31)
}

type  CThostFtdcQryDepthMarketDataField struct {
    InstrumentID string;
    ExchangeID string;
}

func (m* CThostFtdcQryDepthMarketDataField) Marshal(w io.Writer) {
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcQryDepthMarketDataField) Unmarshal(r io.Reader) {
    m.InstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcQryBrokerUserField struct {
    BrokerID string;
    UserID string;
}

func (m* CThostFtdcQryBrokerUserField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.UserID, 16)
}

func (m* CThostFtdcQryBrokerUserField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.UserID = readLen(r, 16)
}

type  CThostFtdcQryBrokerUserFunctionField struct {
    BrokerID string;
    UserID string;
}

func (m* CThostFtdcQryBrokerUserFunctionField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.UserID, 16)
}

func (m* CThostFtdcQryBrokerUserFunctionField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.UserID = readLen(r, 16)
}

type  CThostFtdcQryTraderOfferField struct {
    ExchangeID string;
    ParticipantID string;
    TraderID string;
}

func (m* CThostFtdcQryTraderOfferField) Marshal(w io.Writer) {
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.TraderID, 21)
}

func (m* CThostFtdcQryTraderOfferField) Unmarshal(r io.Reader) {
    m.ExchangeID = readLen(r, 9)
    m.ParticipantID = readLen(r, 11)
    m.TraderID = readLen(r, 21)
}

type  CThostFtdcQrySyncDepositField struct {
    BrokerID string;
    DepositSeqNo string;
}

func (m* CThostFtdcQrySyncDepositField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.DepositSeqNo, 15)
}

func (m* CThostFtdcQrySyncDepositField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.DepositSeqNo = readLen(r, 15)
}

type  CThostFtdcQrySettlementInfoField struct {
    BrokerID string;
    InvestorID string;
    TradingDay string;
    AccountID string;
    CurrencyID string;
}

func (m* CThostFtdcQrySettlementInfoField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.TradingDay, 9)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.CurrencyID, 4)
}

func (m* CThostFtdcQrySettlementInfoField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.TradingDay = readLen(r, 9)
    m.AccountID = readLen(r, 13)
    m.CurrencyID = readLen(r, 4)
}

type  CThostFtdcQryExchangeMarginRateField struct {
    BrokerID string;
    InstrumentID string;
    HedgeFlag uint8;
    ExchangeID string;
}

func (m* CThostFtdcQryExchangeMarginRateField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.HedgeFlag)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcQryExchangeMarginRateField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.HedgeFlag)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcQryExchangeMarginRateAdjustField struct {
    BrokerID string;
    InstrumentID string;
    HedgeFlag uint8;
}

func (m* CThostFtdcQryExchangeMarginRateAdjustField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.HedgeFlag)
}

func (m* CThostFtdcQryExchangeMarginRateAdjustField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.HedgeFlag)
}

type  CThostFtdcQryExchangeRateField struct {
    BrokerID string;
    FromCurrencyID string;
    ToCurrencyID string;
}

func (m* CThostFtdcQryExchangeRateField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.FromCurrencyID, 4)
    writeLen(w, m.ToCurrencyID, 4)
}

func (m* CThostFtdcQryExchangeRateField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.FromCurrencyID = readLen(r, 4)
    m.ToCurrencyID = readLen(r, 4)
}

type  CThostFtdcQrySyncFundMortgageField struct {
    BrokerID string;
    MortgageSeqNo string;
}

func (m* CThostFtdcQrySyncFundMortgageField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.MortgageSeqNo, 15)
}

func (m* CThostFtdcQrySyncFundMortgageField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.MortgageSeqNo = readLen(r, 15)
}

type  CThostFtdcQryHisOrderField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    ExchangeID string;
    OrderSysID string;
    InsertTimeStart string;
    InsertTimeEnd string;
    TradingDay string;
    SettlementID int32;
}

func (m* CThostFtdcQryHisOrderField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.OrderSysID, 21)
    writeLen(w, m.InsertTimeStart, 9)
    writeLen(w, m.InsertTimeEnd, 9)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.SettlementID)
}

func (m* CThostFtdcQryHisOrderField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
    m.OrderSysID = readLen(r, 21)
    m.InsertTimeStart = readLen(r, 9)
    m.InsertTimeEnd = readLen(r, 9)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.SettlementID)
}

type  CThostFtdcOptionInstrMiniMarginField struct {
    InstrumentID string;
    InvestorRange uint8;
    BrokerID string;
    InvestorID string;
    MinMargin float64;
    ValueMethod uint8;
    IsRelative int32;
    ExchangeID string;
}

func (m* CThostFtdcOptionInstrMiniMarginField) Marshal(w io.Writer) {
    writeLen(w, m.InstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.InvestorRange)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.MinMargin)
    binary.Write(w, binary.BigEndian, &m.ValueMethod)
    binary.Write(w, binary.BigEndian, &m.IsRelative)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcOptionInstrMiniMarginField) Unmarshal(r io.Reader) {
    m.InstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.InvestorRange)
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.MinMargin)
    binary.Read(r, binary.BigEndian, &m.ValueMethod)
    binary.Read(r, binary.BigEndian, &m.IsRelative)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcOptionInstrMarginAdjustField struct {
    InstrumentID string;
    InvestorRange uint8;
    BrokerID string;
    InvestorID string;
    SShortMarginRatioByMoney float64;
    SShortMarginRatioByVolume float64;
    HShortMarginRatioByMoney float64;
    HShortMarginRatioByVolume float64;
    AShortMarginRatioByMoney float64;
    AShortMarginRatioByVolume float64;
    IsRelative int32;
    ExchangeID string;
    MShortMarginRatioByMoney float64;
    MShortMarginRatioByVolume float64;
}

func (m* CThostFtdcOptionInstrMarginAdjustField) Marshal(w io.Writer) {
    writeLen(w, m.InstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.InvestorRange)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.SShortMarginRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.SShortMarginRatioByVolume)
    binary.Write(w, binary.BigEndian, &m.HShortMarginRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.HShortMarginRatioByVolume)
    binary.Write(w, binary.BigEndian, &m.AShortMarginRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.AShortMarginRatioByVolume)
    binary.Write(w, binary.BigEndian, &m.IsRelative)
    writeLen(w, m.ExchangeID, 9)
    binary.Write(w, binary.BigEndian, &m.MShortMarginRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.MShortMarginRatioByVolume)
}

func (m* CThostFtdcOptionInstrMarginAdjustField) Unmarshal(r io.Reader) {
    m.InstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.InvestorRange)
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.SShortMarginRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.SShortMarginRatioByVolume)
    binary.Read(r, binary.BigEndian, &m.HShortMarginRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.HShortMarginRatioByVolume)
    binary.Read(r, binary.BigEndian, &m.AShortMarginRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.AShortMarginRatioByVolume)
    binary.Read(r, binary.BigEndian, &m.IsRelative)
    m.ExchangeID = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.MShortMarginRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.MShortMarginRatioByVolume)
}

type  CThostFtdcOptionInstrCommRateField struct {
    InstrumentID string;
    InvestorRange uint8;
    BrokerID string;
    InvestorID string;
    OpenRatioByMoney float64;
    OpenRatioByVolume float64;
    CloseRatioByMoney float64;
    CloseRatioByVolume float64;
    CloseTodayRatioByMoney float64;
    CloseTodayRatioByVolume float64;
    StrikeRatioByMoney float64;
    StrikeRatioByVolume float64;
    ExchangeID string;
    InvestUnitID string;
}

func (m* CThostFtdcOptionInstrCommRateField) Marshal(w io.Writer) {
    writeLen(w, m.InstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.InvestorRange)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.OpenRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.OpenRatioByVolume)
    binary.Write(w, binary.BigEndian, &m.CloseRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.CloseRatioByVolume)
    binary.Write(w, binary.BigEndian, &m.CloseTodayRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.CloseTodayRatioByVolume)
    binary.Write(w, binary.BigEndian, &m.StrikeRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.StrikeRatioByVolume)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.InvestUnitID, 17)
}

func (m* CThostFtdcOptionInstrCommRateField) Unmarshal(r io.Reader) {
    m.InstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.InvestorRange)
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.OpenRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.OpenRatioByVolume)
    binary.Read(r, binary.BigEndian, &m.CloseRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.CloseRatioByVolume)
    binary.Read(r, binary.BigEndian, &m.CloseTodayRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.CloseTodayRatioByVolume)
    binary.Read(r, binary.BigEndian, &m.StrikeRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.StrikeRatioByVolume)
    m.ExchangeID = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
}

type  CThostFtdcOptionInstrTradeCostField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    HedgeFlag uint8;
    FixedMargin float64;
    MiniMargin float64;
    Royalty float64;
    ExchFixedMargin float64;
    ExchMiniMargin float64;
    ExchangeID string;
    InvestUnitID string;
}

func (m* CThostFtdcOptionInstrTradeCostField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.HedgeFlag)
    binary.Write(w, binary.BigEndian, &m.FixedMargin)
    binary.Write(w, binary.BigEndian, &m.MiniMargin)
    binary.Write(w, binary.BigEndian, &m.Royalty)
    binary.Write(w, binary.BigEndian, &m.ExchFixedMargin)
    binary.Write(w, binary.BigEndian, &m.ExchMiniMargin)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.InvestUnitID, 17)
}

func (m* CThostFtdcOptionInstrTradeCostField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.HedgeFlag)
    binary.Read(r, binary.BigEndian, &m.FixedMargin)
    binary.Read(r, binary.BigEndian, &m.MiniMargin)
    binary.Read(r, binary.BigEndian, &m.Royalty)
    binary.Read(r, binary.BigEndian, &m.ExchFixedMargin)
    binary.Read(r, binary.BigEndian, &m.ExchMiniMargin)
    m.ExchangeID = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
}

type  CThostFtdcQryOptionInstrTradeCostField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    HedgeFlag uint8;
    InputPrice float64;
    UnderlyingPrice float64;
    ExchangeID string;
    InvestUnitID string;
}

func (m* CThostFtdcQryOptionInstrTradeCostField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.HedgeFlag)
    binary.Write(w, binary.BigEndian, &m.InputPrice)
    binary.Write(w, binary.BigEndian, &m.UnderlyingPrice)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.InvestUnitID, 17)
}

func (m* CThostFtdcQryOptionInstrTradeCostField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.HedgeFlag)
    binary.Read(r, binary.BigEndian, &m.InputPrice)
    binary.Read(r, binary.BigEndian, &m.UnderlyingPrice)
    m.ExchangeID = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
}

type  CThostFtdcQryOptionInstrCommRateField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    ExchangeID string;
    InvestUnitID string;
}

func (m* CThostFtdcQryOptionInstrCommRateField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.InvestUnitID, 17)
}

func (m* CThostFtdcQryOptionInstrCommRateField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
}

type  CThostFtdcIndexPriceField struct {
    BrokerID string;
    InstrumentID string;
    ClosePrice float64;
    ExchangeID string;
}

func (m* CThostFtdcIndexPriceField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.ClosePrice)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcIndexPriceField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.ClosePrice)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcInputExecOrderField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    ExecOrderRef string;
    UserID string;
    Volume int32;
    RequestID int32;
    BusinessUnit string;
    OffsetFlag uint8;
    HedgeFlag uint8;
    ActionType uint8;
    PosiDirection uint8;
    ReservePositionFlag uint8;
    CloseFlag uint8;
    ExchangeID string;
    InvestUnitID string;
    AccountID string;
    CurrencyID string;
    ClientID string;
    IPAddress string;
    MacAddress string;
}

func (m* CThostFtdcInputExecOrderField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExecOrderRef, 13)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.Volume)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    writeLen(w, m.BusinessUnit, 21)
    binary.Write(w, binary.BigEndian, &m.OffsetFlag)
    binary.Write(w, binary.BigEndian, &m.HedgeFlag)
    binary.Write(w, binary.BigEndian, &m.ActionType)
    binary.Write(w, binary.BigEndian, &m.PosiDirection)
    binary.Write(w, binary.BigEndian, &m.ReservePositionFlag)
    binary.Write(w, binary.BigEndian, &m.CloseFlag)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.InvestUnitID, 17)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.CurrencyID, 4)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
}

func (m* CThostFtdcInputExecOrderField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.ExecOrderRef = readLen(r, 13)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.Volume)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    m.BusinessUnit = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.OffsetFlag)
    binary.Read(r, binary.BigEndian, &m.HedgeFlag)
    binary.Read(r, binary.BigEndian, &m.ActionType)
    binary.Read(r, binary.BigEndian, &m.PosiDirection)
    binary.Read(r, binary.BigEndian, &m.ReservePositionFlag)
    binary.Read(r, binary.BigEndian, &m.CloseFlag)
    m.ExchangeID = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
    m.AccountID = readLen(r, 13)
    m.CurrencyID = readLen(r, 4)
    m.ClientID = readLen(r, 11)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
}

type  CThostFtdcInputExecOrderActionField struct {
    BrokerID string;
    InvestorID string;
    ExecOrderActionRef int32;
    ExecOrderRef string;
    RequestID int32;
    FrontID int32;
    SessionID int32;
    ExchangeID string;
    ExecOrderSysID string;
    ActionFlag uint8;
    UserID string;
    InstrumentID string;
    InvestUnitID string;
    IPAddress string;
    MacAddress string;
}

func (m* CThostFtdcInputExecOrderActionField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.ExecOrderActionRef)
    writeLen(w, m.ExecOrderRef, 13)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    binary.Write(w, binary.BigEndian, &m.FrontID)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ExecOrderSysID, 21)
    binary.Write(w, binary.BigEndian, &m.ActionFlag)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.InvestUnitID, 17)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
}

func (m* CThostFtdcInputExecOrderActionField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.ExecOrderActionRef)
    m.ExecOrderRef = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    binary.Read(r, binary.BigEndian, &m.FrontID)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.ExchangeID = readLen(r, 9)
    m.ExecOrderSysID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.ActionFlag)
    m.UserID = readLen(r, 16)
    m.InstrumentID = readLen(r, 31)
    m.InvestUnitID = readLen(r, 17)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
}

type  CThostFtdcExecOrderField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    ExecOrderRef string;
    UserID string;
    Volume int32;
    RequestID int32;
    BusinessUnit string;
    OffsetFlag uint8;
    HedgeFlag uint8;
    ActionType uint8;
    PosiDirection uint8;
    ReservePositionFlag uint8;
    CloseFlag uint8;
    ExecOrderLocalID string;
    ExchangeID string;
    ParticipantID string;
    ClientID string;
    ExchangeInstID string;
    TraderID string;
    InstallID int32;
    OrderSubmitStatus uint8;
    NotifySequence int32;
    TradingDay string;
    SettlementID int32;
    ExecOrderSysID string;
    InsertDate string;
    InsertTime string;
    CancelTime string;
    ExecResult uint8;
    ClearingPartID string;
    SequenceNo int32;
    FrontID int32;
    SessionID int32;
    UserProductInfo string;
    StatusMsg string;
    ActiveUserID string;
    BrokerExecOrderSeq int32;
    BranchID string;
    InvestUnitID string;
    AccountID string;
    CurrencyID string;
    IPAddress string;
    MacAddress string;
}

func (m* CThostFtdcExecOrderField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExecOrderRef, 13)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.Volume)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    writeLen(w, m.BusinessUnit, 21)
    binary.Write(w, binary.BigEndian, &m.OffsetFlag)
    binary.Write(w, binary.BigEndian, &m.HedgeFlag)
    binary.Write(w, binary.BigEndian, &m.ActionType)
    binary.Write(w, binary.BigEndian, &m.PosiDirection)
    binary.Write(w, binary.BigEndian, &m.ReservePositionFlag)
    binary.Write(w, binary.BigEndian, &m.CloseFlag)
    writeLen(w, m.ExecOrderLocalID, 13)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.ExchangeInstID, 31)
    writeLen(w, m.TraderID, 21)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    binary.Write(w, binary.BigEndian, &m.OrderSubmitStatus)
    binary.Write(w, binary.BigEndian, &m.NotifySequence)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.SettlementID)
    writeLen(w, m.ExecOrderSysID, 21)
    writeLen(w, m.InsertDate, 9)
    writeLen(w, m.InsertTime, 9)
    writeLen(w, m.CancelTime, 9)
    binary.Write(w, binary.BigEndian, &m.ExecResult)
    writeLen(w, m.ClearingPartID, 11)
    binary.Write(w, binary.BigEndian, &m.SequenceNo)
    binary.Write(w, binary.BigEndian, &m.FrontID)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.UserProductInfo, 11)
    writeLen(w, m.StatusMsg, 81)
    writeLen(w, m.ActiveUserID, 16)
    binary.Write(w, binary.BigEndian, &m.BrokerExecOrderSeq)
    writeLen(w, m.BranchID, 9)
    writeLen(w, m.InvestUnitID, 17)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.CurrencyID, 4)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
}

func (m* CThostFtdcExecOrderField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.ExecOrderRef = readLen(r, 13)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.Volume)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    m.BusinessUnit = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.OffsetFlag)
    binary.Read(r, binary.BigEndian, &m.HedgeFlag)
    binary.Read(r, binary.BigEndian, &m.ActionType)
    binary.Read(r, binary.BigEndian, &m.PosiDirection)
    binary.Read(r, binary.BigEndian, &m.ReservePositionFlag)
    binary.Read(r, binary.BigEndian, &m.CloseFlag)
    m.ExecOrderLocalID = readLen(r, 13)
    m.ExchangeID = readLen(r, 9)
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.ExchangeInstID = readLen(r, 31)
    m.TraderID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    binary.Read(r, binary.BigEndian, &m.OrderSubmitStatus)
    binary.Read(r, binary.BigEndian, &m.NotifySequence)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.SettlementID)
    m.ExecOrderSysID = readLen(r, 21)
    m.InsertDate = readLen(r, 9)
    m.InsertTime = readLen(r, 9)
    m.CancelTime = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.ExecResult)
    m.ClearingPartID = readLen(r, 11)
    binary.Read(r, binary.BigEndian, &m.SequenceNo)
    binary.Read(r, binary.BigEndian, &m.FrontID)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.UserProductInfo = readLen(r, 11)
    m.StatusMsg = readLen(r, 81)
    m.ActiveUserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.BrokerExecOrderSeq)
    m.BranchID = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
    m.AccountID = readLen(r, 13)
    m.CurrencyID = readLen(r, 4)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
}

type  CThostFtdcExecOrderActionField struct {
    BrokerID string;
    InvestorID string;
    ExecOrderActionRef int32;
    ExecOrderRef string;
    RequestID int32;
    FrontID int32;
    SessionID int32;
    ExchangeID string;
    ExecOrderSysID string;
    ActionFlag uint8;
    ActionDate string;
    ActionTime string;
    TraderID string;
    InstallID int32;
    ExecOrderLocalID string;
    ActionLocalID string;
    ParticipantID string;
    ClientID string;
    BusinessUnit string;
    OrderActionStatus uint8;
    UserID string;
    ActionType uint8;
    StatusMsg string;
    InstrumentID string;
    BranchID string;
    InvestUnitID string;
    IPAddress string;
    MacAddress string;
}

func (m* CThostFtdcExecOrderActionField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.ExecOrderActionRef)
    writeLen(w, m.ExecOrderRef, 13)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    binary.Write(w, binary.BigEndian, &m.FrontID)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ExecOrderSysID, 21)
    binary.Write(w, binary.BigEndian, &m.ActionFlag)
    writeLen(w, m.ActionDate, 9)
    writeLen(w, m.ActionTime, 9)
    writeLen(w, m.TraderID, 21)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    writeLen(w, m.ExecOrderLocalID, 13)
    writeLen(w, m.ActionLocalID, 13)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.BusinessUnit, 21)
    binary.Write(w, binary.BigEndian, &m.OrderActionStatus)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.ActionType)
    writeLen(w, m.StatusMsg, 81)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.BranchID, 9)
    writeLen(w, m.InvestUnitID, 17)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
}

func (m* CThostFtdcExecOrderActionField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.ExecOrderActionRef)
    m.ExecOrderRef = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    binary.Read(r, binary.BigEndian, &m.FrontID)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.ExchangeID = readLen(r, 9)
    m.ExecOrderSysID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.ActionFlag)
    m.ActionDate = readLen(r, 9)
    m.ActionTime = readLen(r, 9)
    m.TraderID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    m.ExecOrderLocalID = readLen(r, 13)
    m.ActionLocalID = readLen(r, 13)
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.BusinessUnit = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.OrderActionStatus)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.ActionType)
    m.StatusMsg = readLen(r, 81)
    m.InstrumentID = readLen(r, 31)
    m.BranchID = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
}

type  CThostFtdcQryExecOrderField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    ExchangeID string;
    ExecOrderSysID string;
    InsertTimeStart string;
    InsertTimeEnd string;
}

func (m* CThostFtdcQryExecOrderField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ExecOrderSysID, 21)
    writeLen(w, m.InsertTimeStart, 9)
    writeLen(w, m.InsertTimeEnd, 9)
}

func (m* CThostFtdcQryExecOrderField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
    m.ExecOrderSysID = readLen(r, 21)
    m.InsertTimeStart = readLen(r, 9)
    m.InsertTimeEnd = readLen(r, 9)
}

type  CThostFtdcExchangeExecOrderField struct {
    Volume int32;
    RequestID int32;
    BusinessUnit string;
    OffsetFlag uint8;
    HedgeFlag uint8;
    ActionType uint8;
    PosiDirection uint8;
    ReservePositionFlag uint8;
    CloseFlag uint8;
    ExecOrderLocalID string;
    ExchangeID string;
    ParticipantID string;
    ClientID string;
    ExchangeInstID string;
    TraderID string;
    InstallID int32;
    OrderSubmitStatus uint8;
    NotifySequence int32;
    TradingDay string;
    SettlementID int32;
    ExecOrderSysID string;
    InsertDate string;
    InsertTime string;
    CancelTime string;
    ExecResult uint8;
    ClearingPartID string;
    SequenceNo int32;
    BranchID string;
    IPAddress string;
    MacAddress string;
}

func (m* CThostFtdcExchangeExecOrderField) Marshal(w io.Writer) {
    binary.Write(w, binary.BigEndian, &m.Volume)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    writeLen(w, m.BusinessUnit, 21)
    binary.Write(w, binary.BigEndian, &m.OffsetFlag)
    binary.Write(w, binary.BigEndian, &m.HedgeFlag)
    binary.Write(w, binary.BigEndian, &m.ActionType)
    binary.Write(w, binary.BigEndian, &m.PosiDirection)
    binary.Write(w, binary.BigEndian, &m.ReservePositionFlag)
    binary.Write(w, binary.BigEndian, &m.CloseFlag)
    writeLen(w, m.ExecOrderLocalID, 13)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.ExchangeInstID, 31)
    writeLen(w, m.TraderID, 21)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    binary.Write(w, binary.BigEndian, &m.OrderSubmitStatus)
    binary.Write(w, binary.BigEndian, &m.NotifySequence)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.SettlementID)
    writeLen(w, m.ExecOrderSysID, 21)
    writeLen(w, m.InsertDate, 9)
    writeLen(w, m.InsertTime, 9)
    writeLen(w, m.CancelTime, 9)
    binary.Write(w, binary.BigEndian, &m.ExecResult)
    writeLen(w, m.ClearingPartID, 11)
    binary.Write(w, binary.BigEndian, &m.SequenceNo)
    writeLen(w, m.BranchID, 9)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
}

func (m* CThostFtdcExchangeExecOrderField) Unmarshal(r io.Reader) {
    binary.Read(r, binary.BigEndian, &m.Volume)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    m.BusinessUnit = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.OffsetFlag)
    binary.Read(r, binary.BigEndian, &m.HedgeFlag)
    binary.Read(r, binary.BigEndian, &m.ActionType)
    binary.Read(r, binary.BigEndian, &m.PosiDirection)
    binary.Read(r, binary.BigEndian, &m.ReservePositionFlag)
    binary.Read(r, binary.BigEndian, &m.CloseFlag)
    m.ExecOrderLocalID = readLen(r, 13)
    m.ExchangeID = readLen(r, 9)
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.ExchangeInstID = readLen(r, 31)
    m.TraderID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    binary.Read(r, binary.BigEndian, &m.OrderSubmitStatus)
    binary.Read(r, binary.BigEndian, &m.NotifySequence)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.SettlementID)
    m.ExecOrderSysID = readLen(r, 21)
    m.InsertDate = readLen(r, 9)
    m.InsertTime = readLen(r, 9)
    m.CancelTime = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.ExecResult)
    m.ClearingPartID = readLen(r, 11)
    binary.Read(r, binary.BigEndian, &m.SequenceNo)
    m.BranchID = readLen(r, 9)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
}

type  CThostFtdcQryExchangeExecOrderField struct {
    ParticipantID string;
    ClientID string;
    ExchangeInstID string;
    ExchangeID string;
    TraderID string;
}

func (m* CThostFtdcQryExchangeExecOrderField) Marshal(w io.Writer) {
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.ExchangeInstID, 31)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.TraderID, 21)
}

func (m* CThostFtdcQryExchangeExecOrderField) Unmarshal(r io.Reader) {
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.ExchangeInstID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
    m.TraderID = readLen(r, 21)
}

type  CThostFtdcQryExecOrderActionField struct {
    BrokerID string;
    InvestorID string;
    ExchangeID string;
}

func (m* CThostFtdcQryExecOrderActionField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcQryExecOrderActionField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcExchangeExecOrderActionField struct {
    ExchangeID string;
    ExecOrderSysID string;
    ActionFlag uint8;
    ActionDate string;
    ActionTime string;
    TraderID string;
    InstallID int32;
    ExecOrderLocalID string;
    ActionLocalID string;
    ParticipantID string;
    ClientID string;
    BusinessUnit string;
    OrderActionStatus uint8;
    UserID string;
    ActionType uint8;
    BranchID string;
    IPAddress string;
    MacAddress string;
}

func (m* CThostFtdcExchangeExecOrderActionField) Marshal(w io.Writer) {
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ExecOrderSysID, 21)
    binary.Write(w, binary.BigEndian, &m.ActionFlag)
    writeLen(w, m.ActionDate, 9)
    writeLen(w, m.ActionTime, 9)
    writeLen(w, m.TraderID, 21)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    writeLen(w, m.ExecOrderLocalID, 13)
    writeLen(w, m.ActionLocalID, 13)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.BusinessUnit, 21)
    binary.Write(w, binary.BigEndian, &m.OrderActionStatus)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.ActionType)
    writeLen(w, m.BranchID, 9)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
}

func (m* CThostFtdcExchangeExecOrderActionField) Unmarshal(r io.Reader) {
    m.ExchangeID = readLen(r, 9)
    m.ExecOrderSysID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.ActionFlag)
    m.ActionDate = readLen(r, 9)
    m.ActionTime = readLen(r, 9)
    m.TraderID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    m.ExecOrderLocalID = readLen(r, 13)
    m.ActionLocalID = readLen(r, 13)
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.BusinessUnit = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.OrderActionStatus)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.ActionType)
    m.BranchID = readLen(r, 9)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
}

type  CThostFtdcQryExchangeExecOrderActionField struct {
    ParticipantID string;
    ClientID string;
    ExchangeID string;
    TraderID string;
}

func (m* CThostFtdcQryExchangeExecOrderActionField) Marshal(w io.Writer) {
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.TraderID, 21)
}

func (m* CThostFtdcQryExchangeExecOrderActionField) Unmarshal(r io.Reader) {
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.ExchangeID = readLen(r, 9)
    m.TraderID = readLen(r, 21)
}

type  CThostFtdcErrExecOrderField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    ExecOrderRef string;
    UserID string;
    Volume int32;
    RequestID int32;
    BusinessUnit string;
    OffsetFlag uint8;
    HedgeFlag uint8;
    ActionType uint8;
    PosiDirection uint8;
    ReservePositionFlag uint8;
    CloseFlag uint8;
    ExchangeID string;
    InvestUnitID string;
    AccountID string;
    CurrencyID string;
    ClientID string;
    IPAddress string;
    MacAddress string;
    ErrorID int32;
    ErrorMsg string;
}

func (m* CThostFtdcErrExecOrderField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExecOrderRef, 13)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.Volume)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    writeLen(w, m.BusinessUnit, 21)
    binary.Write(w, binary.BigEndian, &m.OffsetFlag)
    binary.Write(w, binary.BigEndian, &m.HedgeFlag)
    binary.Write(w, binary.BigEndian, &m.ActionType)
    binary.Write(w, binary.BigEndian, &m.PosiDirection)
    binary.Write(w, binary.BigEndian, &m.ReservePositionFlag)
    binary.Write(w, binary.BigEndian, &m.CloseFlag)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.InvestUnitID, 17)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.CurrencyID, 4)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
    binary.Write(w, binary.BigEndian, &m.ErrorID)
    writeLen(w, m.ErrorMsg, 81)
}

func (m* CThostFtdcErrExecOrderField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.ExecOrderRef = readLen(r, 13)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.Volume)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    m.BusinessUnit = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.OffsetFlag)
    binary.Read(r, binary.BigEndian, &m.HedgeFlag)
    binary.Read(r, binary.BigEndian, &m.ActionType)
    binary.Read(r, binary.BigEndian, &m.PosiDirection)
    binary.Read(r, binary.BigEndian, &m.ReservePositionFlag)
    binary.Read(r, binary.BigEndian, &m.CloseFlag)
    m.ExchangeID = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
    m.AccountID = readLen(r, 13)
    m.CurrencyID = readLen(r, 4)
    m.ClientID = readLen(r, 11)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.ErrorID)
    m.ErrorMsg = readLen(r, 81)
}

type  CThostFtdcQryErrExecOrderField struct {
    BrokerID string;
    InvestorID string;
}

func (m* CThostFtdcQryErrExecOrderField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
}

func (m* CThostFtdcQryErrExecOrderField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
}

type  CThostFtdcErrExecOrderActionField struct {
    BrokerID string;
    InvestorID string;
    ExecOrderActionRef int32;
    ExecOrderRef string;
    RequestID int32;
    FrontID int32;
    SessionID int32;
    ExchangeID string;
    ExecOrderSysID string;
    ActionFlag uint8;
    UserID string;
    InstrumentID string;
    InvestUnitID string;
    IPAddress string;
    MacAddress string;
    ErrorID int32;
    ErrorMsg string;
}

func (m* CThostFtdcErrExecOrderActionField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.ExecOrderActionRef)
    writeLen(w, m.ExecOrderRef, 13)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    binary.Write(w, binary.BigEndian, &m.FrontID)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ExecOrderSysID, 21)
    binary.Write(w, binary.BigEndian, &m.ActionFlag)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.InvestUnitID, 17)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
    binary.Write(w, binary.BigEndian, &m.ErrorID)
    writeLen(w, m.ErrorMsg, 81)
}

func (m* CThostFtdcErrExecOrderActionField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.ExecOrderActionRef)
    m.ExecOrderRef = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    binary.Read(r, binary.BigEndian, &m.FrontID)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.ExchangeID = readLen(r, 9)
    m.ExecOrderSysID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.ActionFlag)
    m.UserID = readLen(r, 16)
    m.InstrumentID = readLen(r, 31)
    m.InvestUnitID = readLen(r, 17)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.ErrorID)
    m.ErrorMsg = readLen(r, 81)
}

type  CThostFtdcQryErrExecOrderActionField struct {
    BrokerID string;
    InvestorID string;
}

func (m* CThostFtdcQryErrExecOrderActionField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
}

func (m* CThostFtdcQryErrExecOrderActionField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
}

type  CThostFtdcOptionInstrTradingRightField struct {
    InstrumentID string;
    InvestorRange uint8;
    BrokerID string;
    InvestorID string;
    Direction uint8;
    TradingRight uint8;
    ExchangeID string;
    HedgeFlag uint8;
}

func (m* CThostFtdcOptionInstrTradingRightField) Marshal(w io.Writer) {
    writeLen(w, m.InstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.InvestorRange)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.Direction)
    binary.Write(w, binary.BigEndian, &m.TradingRight)
    writeLen(w, m.ExchangeID, 9)
    binary.Write(w, binary.BigEndian, &m.HedgeFlag)
}

func (m* CThostFtdcOptionInstrTradingRightField) Unmarshal(r io.Reader) {
    m.InstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.InvestorRange)
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.Direction)
    binary.Read(r, binary.BigEndian, &m.TradingRight)
    m.ExchangeID = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.HedgeFlag)
}

type  CThostFtdcQryOptionInstrTradingRightField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    Direction uint8;
    ExchangeID string;
}

func (m* CThostFtdcQryOptionInstrTradingRightField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.Direction)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcQryOptionInstrTradingRightField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.Direction)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcInputForQuoteField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    ForQuoteRef string;
    UserID string;
    ExchangeID string;
    InvestUnitID string;
    IPAddress string;
    MacAddress string;
}

func (m* CThostFtdcInputForQuoteField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ForQuoteRef, 13)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.InvestUnitID, 17)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
}

func (m* CThostFtdcInputForQuoteField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.ForQuoteRef = readLen(r, 13)
    m.UserID = readLen(r, 16)
    m.ExchangeID = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
}

type  CThostFtdcForQuoteField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    ForQuoteRef string;
    UserID string;
    ForQuoteLocalID string;
    ExchangeID string;
    ParticipantID string;
    ClientID string;
    ExchangeInstID string;
    TraderID string;
    InstallID int32;
    InsertDate string;
    InsertTime string;
    ForQuoteStatus uint8;
    FrontID int32;
    SessionID int32;
    StatusMsg string;
    ActiveUserID string;
    BrokerForQutoSeq int32;
    InvestUnitID string;
    IPAddress string;
    MacAddress string;
    BranchID string;
}

func (m* CThostFtdcForQuoteField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ForQuoteRef, 13)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.ForQuoteLocalID, 13)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.ExchangeInstID, 31)
    writeLen(w, m.TraderID, 21)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    writeLen(w, m.InsertDate, 9)
    writeLen(w, m.InsertTime, 9)
    binary.Write(w, binary.BigEndian, &m.ForQuoteStatus)
    binary.Write(w, binary.BigEndian, &m.FrontID)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.StatusMsg, 81)
    writeLen(w, m.ActiveUserID, 16)
    binary.Write(w, binary.BigEndian, &m.BrokerForQutoSeq)
    writeLen(w, m.InvestUnitID, 17)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
    writeLen(w, m.BranchID, 9)
}

func (m* CThostFtdcForQuoteField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.ForQuoteRef = readLen(r, 13)
    m.UserID = readLen(r, 16)
    m.ForQuoteLocalID = readLen(r, 13)
    m.ExchangeID = readLen(r, 9)
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.ExchangeInstID = readLen(r, 31)
    m.TraderID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    m.InsertDate = readLen(r, 9)
    m.InsertTime = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.ForQuoteStatus)
    binary.Read(r, binary.BigEndian, &m.FrontID)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.StatusMsg = readLen(r, 81)
    m.ActiveUserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.BrokerForQutoSeq)
    m.InvestUnitID = readLen(r, 17)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
    m.BranchID = readLen(r, 9)
}

type  CThostFtdcQryForQuoteField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    ExchangeID string;
    InsertTimeStart string;
    InsertTimeEnd string;
    InvestUnitID string;
}

func (m* CThostFtdcQryForQuoteField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.InsertTimeStart, 9)
    writeLen(w, m.InsertTimeEnd, 9)
    writeLen(w, m.InvestUnitID, 17)
}

func (m* CThostFtdcQryForQuoteField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
    m.InsertTimeStart = readLen(r, 9)
    m.InsertTimeEnd = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
}

type  CThostFtdcExchangeForQuoteField struct {
    ForQuoteLocalID string;
    ExchangeID string;
    ParticipantID string;
    ClientID string;
    ExchangeInstID string;
    TraderID string;
    InstallID int32;
    InsertDate string;
    InsertTime string;
    ForQuoteStatus uint8;
    IPAddress string;
    MacAddress string;
    BranchID string;
}

func (m* CThostFtdcExchangeForQuoteField) Marshal(w io.Writer) {
    writeLen(w, m.ForQuoteLocalID, 13)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.ExchangeInstID, 31)
    writeLen(w, m.TraderID, 21)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    writeLen(w, m.InsertDate, 9)
    writeLen(w, m.InsertTime, 9)
    binary.Write(w, binary.BigEndian, &m.ForQuoteStatus)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
    writeLen(w, m.BranchID, 9)
}

func (m* CThostFtdcExchangeForQuoteField) Unmarshal(r io.Reader) {
    m.ForQuoteLocalID = readLen(r, 13)
    m.ExchangeID = readLen(r, 9)
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.ExchangeInstID = readLen(r, 31)
    m.TraderID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    m.InsertDate = readLen(r, 9)
    m.InsertTime = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.ForQuoteStatus)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
    m.BranchID = readLen(r, 9)
}

type  CThostFtdcQryExchangeForQuoteField struct {
    ParticipantID string;
    ClientID string;
    ExchangeInstID string;
    ExchangeID string;
    TraderID string;
}

func (m* CThostFtdcQryExchangeForQuoteField) Marshal(w io.Writer) {
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.ExchangeInstID, 31)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.TraderID, 21)
}

func (m* CThostFtdcQryExchangeForQuoteField) Unmarshal(r io.Reader) {
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.ExchangeInstID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
    m.TraderID = readLen(r, 21)
}

type  CThostFtdcInputQuoteField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    QuoteRef string;
    UserID string;
    AskPrice float64;
    BidPrice float64;
    AskVolume int32;
    BidVolume int32;
    RequestID int32;
    BusinessUnit string;
    AskOffsetFlag uint8;
    BidOffsetFlag uint8;
    AskHedgeFlag uint8;
    BidHedgeFlag uint8;
    AskOrderRef string;
    BidOrderRef string;
    ForQuoteSysID string;
    ExchangeID string;
    InvestUnitID string;
    ClientID string;
    IPAddress string;
    MacAddress string;
}

func (m* CThostFtdcInputQuoteField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.QuoteRef, 13)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.AskPrice)
    binary.Write(w, binary.BigEndian, &m.BidPrice)
    binary.Write(w, binary.BigEndian, &m.AskVolume)
    binary.Write(w, binary.BigEndian, &m.BidVolume)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    writeLen(w, m.BusinessUnit, 21)
    binary.Write(w, binary.BigEndian, &m.AskOffsetFlag)
    binary.Write(w, binary.BigEndian, &m.BidOffsetFlag)
    binary.Write(w, binary.BigEndian, &m.AskHedgeFlag)
    binary.Write(w, binary.BigEndian, &m.BidHedgeFlag)
    writeLen(w, m.AskOrderRef, 13)
    writeLen(w, m.BidOrderRef, 13)
    writeLen(w, m.ForQuoteSysID, 21)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.InvestUnitID, 17)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
}

func (m* CThostFtdcInputQuoteField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.QuoteRef = readLen(r, 13)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.AskPrice)
    binary.Read(r, binary.BigEndian, &m.BidPrice)
    binary.Read(r, binary.BigEndian, &m.AskVolume)
    binary.Read(r, binary.BigEndian, &m.BidVolume)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    m.BusinessUnit = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.AskOffsetFlag)
    binary.Read(r, binary.BigEndian, &m.BidOffsetFlag)
    binary.Read(r, binary.BigEndian, &m.AskHedgeFlag)
    binary.Read(r, binary.BigEndian, &m.BidHedgeFlag)
    m.AskOrderRef = readLen(r, 13)
    m.BidOrderRef = readLen(r, 13)
    m.ForQuoteSysID = readLen(r, 21)
    m.ExchangeID = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
    m.ClientID = readLen(r, 11)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
}

type  CThostFtdcInputQuoteActionField struct {
    BrokerID string;
    InvestorID string;
    QuoteActionRef int32;
    QuoteRef string;
    RequestID int32;
    FrontID int32;
    SessionID int32;
    ExchangeID string;
    QuoteSysID string;
    ActionFlag uint8;
    UserID string;
    InstrumentID string;
    InvestUnitID string;
    ClientID string;
    IPAddress string;
    MacAddress string;
}

func (m* CThostFtdcInputQuoteActionField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.QuoteActionRef)
    writeLen(w, m.QuoteRef, 13)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    binary.Write(w, binary.BigEndian, &m.FrontID)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.QuoteSysID, 21)
    binary.Write(w, binary.BigEndian, &m.ActionFlag)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.InvestUnitID, 17)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
}

func (m* CThostFtdcInputQuoteActionField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.QuoteActionRef)
    m.QuoteRef = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    binary.Read(r, binary.BigEndian, &m.FrontID)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.ExchangeID = readLen(r, 9)
    m.QuoteSysID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.ActionFlag)
    m.UserID = readLen(r, 16)
    m.InstrumentID = readLen(r, 31)
    m.InvestUnitID = readLen(r, 17)
    m.ClientID = readLen(r, 11)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
}

type  CThostFtdcQuoteField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    QuoteRef string;
    UserID string;
    AskPrice float64;
    BidPrice float64;
    AskVolume int32;
    BidVolume int32;
    RequestID int32;
    BusinessUnit string;
    AskOffsetFlag uint8;
    BidOffsetFlag uint8;
    AskHedgeFlag uint8;
    BidHedgeFlag uint8;
    QuoteLocalID string;
    ExchangeID string;
    ParticipantID string;
    ClientID string;
    ExchangeInstID string;
    TraderID string;
    InstallID int32;
    NotifySequence int32;
    OrderSubmitStatus uint8;
    TradingDay string;
    SettlementID int32;
    QuoteSysID string;
    InsertDate string;
    InsertTime string;
    CancelTime string;
    QuoteStatus uint8;
    ClearingPartID string;
    SequenceNo int32;
    AskOrderSysID string;
    BidOrderSysID string;
    FrontID int32;
    SessionID int32;
    UserProductInfo string;
    StatusMsg string;
    ActiveUserID string;
    BrokerQuoteSeq int32;
    AskOrderRef string;
    BidOrderRef string;
    ForQuoteSysID string;
    BranchID string;
    InvestUnitID string;
    AccountID string;
    CurrencyID string;
    IPAddress string;
    MacAddress string;
}

func (m* CThostFtdcQuoteField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.QuoteRef, 13)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.AskPrice)
    binary.Write(w, binary.BigEndian, &m.BidPrice)
    binary.Write(w, binary.BigEndian, &m.AskVolume)
    binary.Write(w, binary.BigEndian, &m.BidVolume)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    writeLen(w, m.BusinessUnit, 21)
    binary.Write(w, binary.BigEndian, &m.AskOffsetFlag)
    binary.Write(w, binary.BigEndian, &m.BidOffsetFlag)
    binary.Write(w, binary.BigEndian, &m.AskHedgeFlag)
    binary.Write(w, binary.BigEndian, &m.BidHedgeFlag)
    writeLen(w, m.QuoteLocalID, 13)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.ExchangeInstID, 31)
    writeLen(w, m.TraderID, 21)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    binary.Write(w, binary.BigEndian, &m.NotifySequence)
    binary.Write(w, binary.BigEndian, &m.OrderSubmitStatus)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.SettlementID)
    writeLen(w, m.QuoteSysID, 21)
    writeLen(w, m.InsertDate, 9)
    writeLen(w, m.InsertTime, 9)
    writeLen(w, m.CancelTime, 9)
    binary.Write(w, binary.BigEndian, &m.QuoteStatus)
    writeLen(w, m.ClearingPartID, 11)
    binary.Write(w, binary.BigEndian, &m.SequenceNo)
    writeLen(w, m.AskOrderSysID, 21)
    writeLen(w, m.BidOrderSysID, 21)
    binary.Write(w, binary.BigEndian, &m.FrontID)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.UserProductInfo, 11)
    writeLen(w, m.StatusMsg, 81)
    writeLen(w, m.ActiveUserID, 16)
    binary.Write(w, binary.BigEndian, &m.BrokerQuoteSeq)
    writeLen(w, m.AskOrderRef, 13)
    writeLen(w, m.BidOrderRef, 13)
    writeLen(w, m.ForQuoteSysID, 21)
    writeLen(w, m.BranchID, 9)
    writeLen(w, m.InvestUnitID, 17)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.CurrencyID, 4)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
}

func (m* CThostFtdcQuoteField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.QuoteRef = readLen(r, 13)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.AskPrice)
    binary.Read(r, binary.BigEndian, &m.BidPrice)
    binary.Read(r, binary.BigEndian, &m.AskVolume)
    binary.Read(r, binary.BigEndian, &m.BidVolume)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    m.BusinessUnit = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.AskOffsetFlag)
    binary.Read(r, binary.BigEndian, &m.BidOffsetFlag)
    binary.Read(r, binary.BigEndian, &m.AskHedgeFlag)
    binary.Read(r, binary.BigEndian, &m.BidHedgeFlag)
    m.QuoteLocalID = readLen(r, 13)
    m.ExchangeID = readLen(r, 9)
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.ExchangeInstID = readLen(r, 31)
    m.TraderID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    binary.Read(r, binary.BigEndian, &m.NotifySequence)
    binary.Read(r, binary.BigEndian, &m.OrderSubmitStatus)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.SettlementID)
    m.QuoteSysID = readLen(r, 21)
    m.InsertDate = readLen(r, 9)
    m.InsertTime = readLen(r, 9)
    m.CancelTime = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.QuoteStatus)
    m.ClearingPartID = readLen(r, 11)
    binary.Read(r, binary.BigEndian, &m.SequenceNo)
    m.AskOrderSysID = readLen(r, 21)
    m.BidOrderSysID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.FrontID)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.UserProductInfo = readLen(r, 11)
    m.StatusMsg = readLen(r, 81)
    m.ActiveUserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.BrokerQuoteSeq)
    m.AskOrderRef = readLen(r, 13)
    m.BidOrderRef = readLen(r, 13)
    m.ForQuoteSysID = readLen(r, 21)
    m.BranchID = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
    m.AccountID = readLen(r, 13)
    m.CurrencyID = readLen(r, 4)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
}

type  CThostFtdcQuoteActionField struct {
    BrokerID string;
    InvestorID string;
    QuoteActionRef int32;
    QuoteRef string;
    RequestID int32;
    FrontID int32;
    SessionID int32;
    ExchangeID string;
    QuoteSysID string;
    ActionFlag uint8;
    ActionDate string;
    ActionTime string;
    TraderID string;
    InstallID int32;
    QuoteLocalID string;
    ActionLocalID string;
    ParticipantID string;
    ClientID string;
    BusinessUnit string;
    OrderActionStatus uint8;
    UserID string;
    StatusMsg string;
    InstrumentID string;
    BranchID string;
    InvestUnitID string;
    IPAddress string;
    MacAddress string;
}

func (m* CThostFtdcQuoteActionField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.QuoteActionRef)
    writeLen(w, m.QuoteRef, 13)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    binary.Write(w, binary.BigEndian, &m.FrontID)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.QuoteSysID, 21)
    binary.Write(w, binary.BigEndian, &m.ActionFlag)
    writeLen(w, m.ActionDate, 9)
    writeLen(w, m.ActionTime, 9)
    writeLen(w, m.TraderID, 21)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    writeLen(w, m.QuoteLocalID, 13)
    writeLen(w, m.ActionLocalID, 13)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.BusinessUnit, 21)
    binary.Write(w, binary.BigEndian, &m.OrderActionStatus)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.StatusMsg, 81)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.BranchID, 9)
    writeLen(w, m.InvestUnitID, 17)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
}

func (m* CThostFtdcQuoteActionField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.QuoteActionRef)
    m.QuoteRef = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    binary.Read(r, binary.BigEndian, &m.FrontID)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.ExchangeID = readLen(r, 9)
    m.QuoteSysID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.ActionFlag)
    m.ActionDate = readLen(r, 9)
    m.ActionTime = readLen(r, 9)
    m.TraderID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    m.QuoteLocalID = readLen(r, 13)
    m.ActionLocalID = readLen(r, 13)
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.BusinessUnit = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.OrderActionStatus)
    m.UserID = readLen(r, 16)
    m.StatusMsg = readLen(r, 81)
    m.InstrumentID = readLen(r, 31)
    m.BranchID = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
}

type  CThostFtdcQryQuoteField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    ExchangeID string;
    QuoteSysID string;
    InsertTimeStart string;
    InsertTimeEnd string;
    InvestUnitID string;
}

func (m* CThostFtdcQryQuoteField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.QuoteSysID, 21)
    writeLen(w, m.InsertTimeStart, 9)
    writeLen(w, m.InsertTimeEnd, 9)
    writeLen(w, m.InvestUnitID, 17)
}

func (m* CThostFtdcQryQuoteField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
    m.QuoteSysID = readLen(r, 21)
    m.InsertTimeStart = readLen(r, 9)
    m.InsertTimeEnd = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
}

type  CThostFtdcExchangeQuoteField struct {
    AskPrice float64;
    BidPrice float64;
    AskVolume int32;
    BidVolume int32;
    RequestID int32;
    BusinessUnit string;
    AskOffsetFlag uint8;
    BidOffsetFlag uint8;
    AskHedgeFlag uint8;
    BidHedgeFlag uint8;
    QuoteLocalID string;
    ExchangeID string;
    ParticipantID string;
    ClientID string;
    ExchangeInstID string;
    TraderID string;
    InstallID int32;
    NotifySequence int32;
    OrderSubmitStatus uint8;
    TradingDay string;
    SettlementID int32;
    QuoteSysID string;
    InsertDate string;
    InsertTime string;
    CancelTime string;
    QuoteStatus uint8;
    ClearingPartID string;
    SequenceNo int32;
    AskOrderSysID string;
    BidOrderSysID string;
    ForQuoteSysID string;
    BranchID string;
    IPAddress string;
    MacAddress string;
}

func (m* CThostFtdcExchangeQuoteField) Marshal(w io.Writer) {
    binary.Write(w, binary.BigEndian, &m.AskPrice)
    binary.Write(w, binary.BigEndian, &m.BidPrice)
    binary.Write(w, binary.BigEndian, &m.AskVolume)
    binary.Write(w, binary.BigEndian, &m.BidVolume)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    writeLen(w, m.BusinessUnit, 21)
    binary.Write(w, binary.BigEndian, &m.AskOffsetFlag)
    binary.Write(w, binary.BigEndian, &m.BidOffsetFlag)
    binary.Write(w, binary.BigEndian, &m.AskHedgeFlag)
    binary.Write(w, binary.BigEndian, &m.BidHedgeFlag)
    writeLen(w, m.QuoteLocalID, 13)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.ExchangeInstID, 31)
    writeLen(w, m.TraderID, 21)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    binary.Write(w, binary.BigEndian, &m.NotifySequence)
    binary.Write(w, binary.BigEndian, &m.OrderSubmitStatus)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.SettlementID)
    writeLen(w, m.QuoteSysID, 21)
    writeLen(w, m.InsertDate, 9)
    writeLen(w, m.InsertTime, 9)
    writeLen(w, m.CancelTime, 9)
    binary.Write(w, binary.BigEndian, &m.QuoteStatus)
    writeLen(w, m.ClearingPartID, 11)
    binary.Write(w, binary.BigEndian, &m.SequenceNo)
    writeLen(w, m.AskOrderSysID, 21)
    writeLen(w, m.BidOrderSysID, 21)
    writeLen(w, m.ForQuoteSysID, 21)
    writeLen(w, m.BranchID, 9)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
}

func (m* CThostFtdcExchangeQuoteField) Unmarshal(r io.Reader) {
    binary.Read(r, binary.BigEndian, &m.AskPrice)
    binary.Read(r, binary.BigEndian, &m.BidPrice)
    binary.Read(r, binary.BigEndian, &m.AskVolume)
    binary.Read(r, binary.BigEndian, &m.BidVolume)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    m.BusinessUnit = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.AskOffsetFlag)
    binary.Read(r, binary.BigEndian, &m.BidOffsetFlag)
    binary.Read(r, binary.BigEndian, &m.AskHedgeFlag)
    binary.Read(r, binary.BigEndian, &m.BidHedgeFlag)
    m.QuoteLocalID = readLen(r, 13)
    m.ExchangeID = readLen(r, 9)
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.ExchangeInstID = readLen(r, 31)
    m.TraderID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    binary.Read(r, binary.BigEndian, &m.NotifySequence)
    binary.Read(r, binary.BigEndian, &m.OrderSubmitStatus)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.SettlementID)
    m.QuoteSysID = readLen(r, 21)
    m.InsertDate = readLen(r, 9)
    m.InsertTime = readLen(r, 9)
    m.CancelTime = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.QuoteStatus)
    m.ClearingPartID = readLen(r, 11)
    binary.Read(r, binary.BigEndian, &m.SequenceNo)
    m.AskOrderSysID = readLen(r, 21)
    m.BidOrderSysID = readLen(r, 21)
    m.ForQuoteSysID = readLen(r, 21)
    m.BranchID = readLen(r, 9)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
}

type  CThostFtdcQryExchangeQuoteField struct {
    ParticipantID string;
    ClientID string;
    ExchangeInstID string;
    ExchangeID string;
    TraderID string;
}

func (m* CThostFtdcQryExchangeQuoteField) Marshal(w io.Writer) {
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.ExchangeInstID, 31)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.TraderID, 21)
}

func (m* CThostFtdcQryExchangeQuoteField) Unmarshal(r io.Reader) {
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.ExchangeInstID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
    m.TraderID = readLen(r, 21)
}

type  CThostFtdcQryQuoteActionField struct {
    BrokerID string;
    InvestorID string;
    ExchangeID string;
}

func (m* CThostFtdcQryQuoteActionField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcQryQuoteActionField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcExchangeQuoteActionField struct {
    ExchangeID string;
    QuoteSysID string;
    ActionFlag uint8;
    ActionDate string;
    ActionTime string;
    TraderID string;
    InstallID int32;
    QuoteLocalID string;
    ActionLocalID string;
    ParticipantID string;
    ClientID string;
    BusinessUnit string;
    OrderActionStatus uint8;
    UserID string;
    IPAddress string;
    MacAddress string;
    BranchID string;
}

func (m* CThostFtdcExchangeQuoteActionField) Marshal(w io.Writer) {
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.QuoteSysID, 21)
    binary.Write(w, binary.BigEndian, &m.ActionFlag)
    writeLen(w, m.ActionDate, 9)
    writeLen(w, m.ActionTime, 9)
    writeLen(w, m.TraderID, 21)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    writeLen(w, m.QuoteLocalID, 13)
    writeLen(w, m.ActionLocalID, 13)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.BusinessUnit, 21)
    binary.Write(w, binary.BigEndian, &m.OrderActionStatus)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
    writeLen(w, m.BranchID, 9)
}

func (m* CThostFtdcExchangeQuoteActionField) Unmarshal(r io.Reader) {
    m.ExchangeID = readLen(r, 9)
    m.QuoteSysID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.ActionFlag)
    m.ActionDate = readLen(r, 9)
    m.ActionTime = readLen(r, 9)
    m.TraderID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    m.QuoteLocalID = readLen(r, 13)
    m.ActionLocalID = readLen(r, 13)
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.BusinessUnit = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.OrderActionStatus)
    m.UserID = readLen(r, 16)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
    m.BranchID = readLen(r, 9)
}

type  CThostFtdcQryExchangeQuoteActionField struct {
    ParticipantID string;
    ClientID string;
    ExchangeID string;
    TraderID string;
}

func (m* CThostFtdcQryExchangeQuoteActionField) Marshal(w io.Writer) {
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.TraderID, 21)
}

func (m* CThostFtdcQryExchangeQuoteActionField) Unmarshal(r io.Reader) {
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.ExchangeID = readLen(r, 9)
    m.TraderID = readLen(r, 21)
}

type  CThostFtdcOptionInstrDeltaField struct {
    InstrumentID string;
    InvestorRange uint8;
    BrokerID string;
    InvestorID string;
    Delta float64;
    ExchangeID string;
}

func (m* CThostFtdcOptionInstrDeltaField) Marshal(w io.Writer) {
    writeLen(w, m.InstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.InvestorRange)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.Delta)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcOptionInstrDeltaField) Unmarshal(r io.Reader) {
    m.InstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.InvestorRange)
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.Delta)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcForQuoteRspField struct {
    TradingDay string;
    InstrumentID string;
    ForQuoteSysID string;
    ForQuoteTime string;
    ActionDay string;
    ExchangeID string;
}

func (m* CThostFtdcForQuoteRspField) Marshal(w io.Writer) {
    writeLen(w, m.TradingDay, 9)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ForQuoteSysID, 21)
    writeLen(w, m.ForQuoteTime, 9)
    writeLen(w, m.ActionDay, 9)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcForQuoteRspField) Unmarshal(r io.Reader) {
    m.TradingDay = readLen(r, 9)
    m.InstrumentID = readLen(r, 31)
    m.ForQuoteSysID = readLen(r, 21)
    m.ForQuoteTime = readLen(r, 9)
    m.ActionDay = readLen(r, 9)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcStrikeOffsetField struct {
    InstrumentID string;
    InvestorRange uint8;
    BrokerID string;
    InvestorID string;
    Offset float64;
    OffsetType uint8;
    ExchangeID string;
}

func (m* CThostFtdcStrikeOffsetField) Marshal(w io.Writer) {
    writeLen(w, m.InstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.InvestorRange)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.Offset)
    binary.Write(w, binary.BigEndian, &m.OffsetType)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcStrikeOffsetField) Unmarshal(r io.Reader) {
    m.InstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.InvestorRange)
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.Offset)
    binary.Read(r, binary.BigEndian, &m.OffsetType)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcQryStrikeOffsetField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
}

func (m* CThostFtdcQryStrikeOffsetField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
}

func (m* CThostFtdcQryStrikeOffsetField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
}

type  CThostFtdcInputLockField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    LockRef string;
    UserID string;
    Volume int32;
    RequestID int32;
    BusinessUnit string;
    LockType uint8;
    ExchangeID string;
    IPAddress string;
    MacAddress string;
}

func (m* CThostFtdcInputLockField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.LockRef, 13)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.Volume)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    writeLen(w, m.BusinessUnit, 21)
    binary.Write(w, binary.BigEndian, &m.LockType)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
}

func (m* CThostFtdcInputLockField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.LockRef = readLen(r, 13)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.Volume)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    m.BusinessUnit = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.LockType)
    m.ExchangeID = readLen(r, 9)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
}

type  CThostFtdcLockField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    LockRef string;
    UserID string;
    Volume int32;
    RequestID int32;
    BusinessUnit string;
    LockType uint8;
    LockLocalID string;
    ExchangeID string;
    ParticipantID string;
    ClientID string;
    ExchangeInstID string;
    TraderID string;
    InstallID int32;
    OrderSubmitStatus uint8;
    NotifySequence int32;
    TradingDay string;
    SettlementID int32;
    LockSysID string;
    InsertDate string;
    InsertTime string;
    CancelTime string;
    LockStatus uint8;
    ClearingPartID string;
    SequenceNo int32;
    FrontID int32;
    SessionID int32;
    UserProductInfo string;
    StatusMsg string;
    ActiveUserID string;
    BrokerLockSeq int32;
    BranchID string;
    IPAddress string;
    MacAddress string;
}

func (m* CThostFtdcLockField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.LockRef, 13)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.Volume)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    writeLen(w, m.BusinessUnit, 21)
    binary.Write(w, binary.BigEndian, &m.LockType)
    writeLen(w, m.LockLocalID, 13)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.ExchangeInstID, 31)
    writeLen(w, m.TraderID, 21)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    binary.Write(w, binary.BigEndian, &m.OrderSubmitStatus)
    binary.Write(w, binary.BigEndian, &m.NotifySequence)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.SettlementID)
    writeLen(w, m.LockSysID, 21)
    writeLen(w, m.InsertDate, 9)
    writeLen(w, m.InsertTime, 9)
    writeLen(w, m.CancelTime, 9)
    binary.Write(w, binary.BigEndian, &m.LockStatus)
    writeLen(w, m.ClearingPartID, 11)
    binary.Write(w, binary.BigEndian, &m.SequenceNo)
    binary.Write(w, binary.BigEndian, &m.FrontID)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.UserProductInfo, 11)
    writeLen(w, m.StatusMsg, 81)
    writeLen(w, m.ActiveUserID, 16)
    binary.Write(w, binary.BigEndian, &m.BrokerLockSeq)
    writeLen(w, m.BranchID, 9)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
}

func (m* CThostFtdcLockField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.LockRef = readLen(r, 13)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.Volume)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    m.BusinessUnit = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.LockType)
    m.LockLocalID = readLen(r, 13)
    m.ExchangeID = readLen(r, 9)
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.ExchangeInstID = readLen(r, 31)
    m.TraderID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    binary.Read(r, binary.BigEndian, &m.OrderSubmitStatus)
    binary.Read(r, binary.BigEndian, &m.NotifySequence)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.SettlementID)
    m.LockSysID = readLen(r, 21)
    m.InsertDate = readLen(r, 9)
    m.InsertTime = readLen(r, 9)
    m.CancelTime = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.LockStatus)
    m.ClearingPartID = readLen(r, 11)
    binary.Read(r, binary.BigEndian, &m.SequenceNo)
    binary.Read(r, binary.BigEndian, &m.FrontID)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.UserProductInfo = readLen(r, 11)
    m.StatusMsg = readLen(r, 81)
    m.ActiveUserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.BrokerLockSeq)
    m.BranchID = readLen(r, 9)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
}

type  CThostFtdcQryLockField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    ExchangeID string;
    LockSysID string;
    InsertTimeStart string;
    InsertTimeEnd string;
}

func (m* CThostFtdcQryLockField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.LockSysID, 21)
    writeLen(w, m.InsertTimeStart, 9)
    writeLen(w, m.InsertTimeEnd, 9)
}

func (m* CThostFtdcQryLockField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
    m.LockSysID = readLen(r, 21)
    m.InsertTimeStart = readLen(r, 9)
    m.InsertTimeEnd = readLen(r, 9)
}

type  CThostFtdcLockPositionField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    ExchangeID string;
    Volume int32;
    FrozenVolume int32;
    TodayVolume int32;
}

func (m* CThostFtdcLockPositionField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
    binary.Write(w, binary.BigEndian, &m.Volume)
    binary.Write(w, binary.BigEndian, &m.FrozenVolume)
    binary.Write(w, binary.BigEndian, &m.TodayVolume)
}

func (m* CThostFtdcLockPositionField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.Volume)
    binary.Read(r, binary.BigEndian, &m.FrozenVolume)
    binary.Read(r, binary.BigEndian, &m.TodayVolume)
}

type  CThostFtdcQryLockPositionField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    ExchangeID string;
}

func (m* CThostFtdcQryLockPositionField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcQryLockPositionField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcETFOptionInstrCommRateField struct {
    InstrumentID string;
    InvestorRange uint8;
    BrokerID string;
    InvestorID string;
    OpenRatioByMoney float64;
    OpenRatioByVolume float64;
    CloseRatioByMoney float64;
    CloseRatioByVolume float64;
    CloseTodayRatioByMoney float64;
    CloseTodayRatioByVolume float64;
    StrikeRatioByMoney float64;
    StrikeRatioByVolume float64;
    ExchangeID string;
    HedgeFlag uint8;
    PosiDirection uint8;
}

func (m* CThostFtdcETFOptionInstrCommRateField) Marshal(w io.Writer) {
    writeLen(w, m.InstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.InvestorRange)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.OpenRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.OpenRatioByVolume)
    binary.Write(w, binary.BigEndian, &m.CloseRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.CloseRatioByVolume)
    binary.Write(w, binary.BigEndian, &m.CloseTodayRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.CloseTodayRatioByVolume)
    binary.Write(w, binary.BigEndian, &m.StrikeRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.StrikeRatioByVolume)
    writeLen(w, m.ExchangeID, 9)
    binary.Write(w, binary.BigEndian, &m.HedgeFlag)
    binary.Write(w, binary.BigEndian, &m.PosiDirection)
}

func (m* CThostFtdcETFOptionInstrCommRateField) Unmarshal(r io.Reader) {
    m.InstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.InvestorRange)
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.OpenRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.OpenRatioByVolume)
    binary.Read(r, binary.BigEndian, &m.CloseRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.CloseRatioByVolume)
    binary.Read(r, binary.BigEndian, &m.CloseTodayRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.CloseTodayRatioByVolume)
    binary.Read(r, binary.BigEndian, &m.StrikeRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.StrikeRatioByVolume)
    m.ExchangeID = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.HedgeFlag)
    binary.Read(r, binary.BigEndian, &m.PosiDirection)
}

type  CThostFtdcQryETFOptionInstrCommRateField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    ExchangeID string;
}

func (m* CThostFtdcQryETFOptionInstrCommRateField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcQryETFOptionInstrCommRateField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcPosiFreezeField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    ExchangeID string;
    OrderLocalID string;
    TraderID string;
    ParticipantID string;
    InstallID int32;
    Volume int32;
    FreezeReasonType uint8;
    FreezeType uint8;
}

func (m* CThostFtdcPosiFreezeField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.OrderLocalID, 13)
    writeLen(w, m.TraderID, 21)
    writeLen(w, m.ParticipantID, 11)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    binary.Write(w, binary.BigEndian, &m.Volume)
    binary.Write(w, binary.BigEndian, &m.FreezeReasonType)
    binary.Write(w, binary.BigEndian, &m.FreezeType)
}

func (m* CThostFtdcPosiFreezeField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
    m.OrderLocalID = readLen(r, 13)
    m.TraderID = readLen(r, 21)
    m.ParticipantID = readLen(r, 11)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    binary.Read(r, binary.BigEndian, &m.Volume)
    binary.Read(r, binary.BigEndian, &m.FreezeReasonType)
    binary.Read(r, binary.BigEndian, &m.FreezeType)
}

type  CThostFtdcQryExchangeLockField struct {
    ParticipantID string;
    ClientID string;
    ExchangeInstID string;
    ExchangeID string;
    TraderID string;
}

func (m* CThostFtdcQryExchangeLockField) Marshal(w io.Writer) {
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.ExchangeInstID, 31)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.TraderID, 21)
}

func (m* CThostFtdcQryExchangeLockField) Unmarshal(r io.Reader) {
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.ExchangeInstID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
    m.TraderID = readLen(r, 21)
}

type  CThostFtdcExchangeLockField struct {
    Volume int32;
    RequestID int32;
    BusinessUnit string;
    LockType uint8;
    LockLocalID string;
    ExchangeID string;
    ParticipantID string;
    ClientID string;
    ExchangeInstID string;
    TraderID string;
    InstallID int32;
    OrderSubmitStatus uint8;
    NotifySequence int32;
    TradingDay string;
    SettlementID int32;
    LockSysID string;
    InsertDate string;
    InsertTime string;
    CancelTime string;
    LockStatus uint8;
    ClearingPartID string;
    SequenceNo int32;
    BranchID string;
    IPAddress string;
    MacAddress string;
}

func (m* CThostFtdcExchangeLockField) Marshal(w io.Writer) {
    binary.Write(w, binary.BigEndian, &m.Volume)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    writeLen(w, m.BusinessUnit, 21)
    binary.Write(w, binary.BigEndian, &m.LockType)
    writeLen(w, m.LockLocalID, 13)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.ExchangeInstID, 31)
    writeLen(w, m.TraderID, 21)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    binary.Write(w, binary.BigEndian, &m.OrderSubmitStatus)
    binary.Write(w, binary.BigEndian, &m.NotifySequence)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.SettlementID)
    writeLen(w, m.LockSysID, 21)
    writeLen(w, m.InsertDate, 9)
    writeLen(w, m.InsertTime, 9)
    writeLen(w, m.CancelTime, 9)
    binary.Write(w, binary.BigEndian, &m.LockStatus)
    writeLen(w, m.ClearingPartID, 11)
    binary.Write(w, binary.BigEndian, &m.SequenceNo)
    writeLen(w, m.BranchID, 9)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
}

func (m* CThostFtdcExchangeLockField) Unmarshal(r io.Reader) {
    binary.Read(r, binary.BigEndian, &m.Volume)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    m.BusinessUnit = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.LockType)
    m.LockLocalID = readLen(r, 13)
    m.ExchangeID = readLen(r, 9)
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.ExchangeInstID = readLen(r, 31)
    m.TraderID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    binary.Read(r, binary.BigEndian, &m.OrderSubmitStatus)
    binary.Read(r, binary.BigEndian, &m.NotifySequence)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.SettlementID)
    m.LockSysID = readLen(r, 21)
    m.InsertDate = readLen(r, 9)
    m.InsertTime = readLen(r, 9)
    m.CancelTime = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.LockStatus)
    m.ClearingPartID = readLen(r, 11)
    binary.Read(r, binary.BigEndian, &m.SequenceNo)
    m.BranchID = readLen(r, 9)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
}

type  CThostFtdcExchangeExecOrderActionErrorField struct {
    ExchangeID string;
    ExecOrderSysID string;
    TraderID string;
    InstallID int32;
    ExecOrderLocalID string;
    ActionLocalID string;
    ErrorID int32;
    ErrorMsg string;
    BrokerID string;
}

func (m* CThostFtdcExchangeExecOrderActionErrorField) Marshal(w io.Writer) {
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ExecOrderSysID, 21)
    writeLen(w, m.TraderID, 21)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    writeLen(w, m.ExecOrderLocalID, 13)
    writeLen(w, m.ActionLocalID, 13)
    binary.Write(w, binary.BigEndian, &m.ErrorID)
    writeLen(w, m.ErrorMsg, 81)
    writeLen(w, m.BrokerID, 11)
}

func (m* CThostFtdcExchangeExecOrderActionErrorField) Unmarshal(r io.Reader) {
    m.ExchangeID = readLen(r, 9)
    m.ExecOrderSysID = readLen(r, 21)
    m.TraderID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    m.ExecOrderLocalID = readLen(r, 13)
    m.ActionLocalID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.ErrorID)
    m.ErrorMsg = readLen(r, 81)
    m.BrokerID = readLen(r, 11)
}

type  CThostFtdcInputBatchOrderActionField struct {
    BrokerID string;
    InvestorID string;
    OrderActionRef int32;
    RequestID int32;
    FrontID int32;
    SessionID int32;
    ExchangeID string;
    UserID string;
    InvestUnitID string;
    IPAddress string;
    MacAddress string;
}

func (m* CThostFtdcInputBatchOrderActionField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.OrderActionRef)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    binary.Write(w, binary.BigEndian, &m.FrontID)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.InvestUnitID, 17)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
}

func (m* CThostFtdcInputBatchOrderActionField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.OrderActionRef)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    binary.Read(r, binary.BigEndian, &m.FrontID)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.ExchangeID = readLen(r, 9)
    m.UserID = readLen(r, 16)
    m.InvestUnitID = readLen(r, 17)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
}

type  CThostFtdcBatchOrderActionField struct {
    BrokerID string;
    InvestorID string;
    OrderActionRef int32;
    RequestID int32;
    FrontID int32;
    SessionID int32;
    ExchangeID string;
    ActionDate string;
    ActionTime string;
    TraderID string;
    InstallID int32;
    ActionLocalID string;
    ParticipantID string;
    ClientID string;
    BusinessUnit string;
    OrderActionStatus uint8;
    UserID string;
    StatusMsg string;
    InvestUnitID string;
    IPAddress string;
    MacAddress string;
}

func (m* CThostFtdcBatchOrderActionField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.OrderActionRef)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    binary.Write(w, binary.BigEndian, &m.FrontID)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ActionDate, 9)
    writeLen(w, m.ActionTime, 9)
    writeLen(w, m.TraderID, 21)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    writeLen(w, m.ActionLocalID, 13)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.BusinessUnit, 21)
    binary.Write(w, binary.BigEndian, &m.OrderActionStatus)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.StatusMsg, 81)
    writeLen(w, m.InvestUnitID, 17)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
}

func (m* CThostFtdcBatchOrderActionField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.OrderActionRef)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    binary.Read(r, binary.BigEndian, &m.FrontID)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.ExchangeID = readLen(r, 9)
    m.ActionDate = readLen(r, 9)
    m.ActionTime = readLen(r, 9)
    m.TraderID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    m.ActionLocalID = readLen(r, 13)
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.BusinessUnit = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.OrderActionStatus)
    m.UserID = readLen(r, 16)
    m.StatusMsg = readLen(r, 81)
    m.InvestUnitID = readLen(r, 17)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
}

type  CThostFtdcExchangeBatchOrderActionField struct {
    ExchangeID string;
    ActionDate string;
    ActionTime string;
    TraderID string;
    InstallID int32;
    ActionLocalID string;
    ParticipantID string;
    ClientID string;
    BusinessUnit string;
    OrderActionStatus uint8;
    UserID string;
    IPAddress string;
    MacAddress string;
}

func (m* CThostFtdcExchangeBatchOrderActionField) Marshal(w io.Writer) {
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ActionDate, 9)
    writeLen(w, m.ActionTime, 9)
    writeLen(w, m.TraderID, 21)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    writeLen(w, m.ActionLocalID, 13)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.BusinessUnit, 21)
    binary.Write(w, binary.BigEndian, &m.OrderActionStatus)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
}

func (m* CThostFtdcExchangeBatchOrderActionField) Unmarshal(r io.Reader) {
    m.ExchangeID = readLen(r, 9)
    m.ActionDate = readLen(r, 9)
    m.ActionTime = readLen(r, 9)
    m.TraderID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    m.ActionLocalID = readLen(r, 13)
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.BusinessUnit = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.OrderActionStatus)
    m.UserID = readLen(r, 16)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
}

type  CThostFtdcQryBatchOrderActionField struct {
    BrokerID string;
    InvestorID string;
    ExchangeID string;
}

func (m* CThostFtdcQryBatchOrderActionField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcQryBatchOrderActionField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcLimitPosiField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    ExchangeID string;
    TotalVolume int32;
    LongVolume int32;
    OpenVolume int32;
    LongAmount float64;
    TotalVolumeFrozen int32;
    LongVolumeFrozen int32;
    OpenVolumeFrozen int32;
    LongAmountFrozen float64;
}

func (m* CThostFtdcLimitPosiField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
    binary.Write(w, binary.BigEndian, &m.TotalVolume)
    binary.Write(w, binary.BigEndian, &m.LongVolume)
    binary.Write(w, binary.BigEndian, &m.OpenVolume)
    binary.Write(w, binary.BigEndian, &m.LongAmount)
    binary.Write(w, binary.BigEndian, &m.TotalVolumeFrozen)
    binary.Write(w, binary.BigEndian, &m.LongVolumeFrozen)
    binary.Write(w, binary.BigEndian, &m.OpenVolumeFrozen)
    binary.Write(w, binary.BigEndian, &m.LongAmountFrozen)
}

func (m* CThostFtdcLimitPosiField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.TotalVolume)
    binary.Read(r, binary.BigEndian, &m.LongVolume)
    binary.Read(r, binary.BigEndian, &m.OpenVolume)
    binary.Read(r, binary.BigEndian, &m.LongAmount)
    binary.Read(r, binary.BigEndian, &m.TotalVolumeFrozen)
    binary.Read(r, binary.BigEndian, &m.LongVolumeFrozen)
    binary.Read(r, binary.BigEndian, &m.OpenVolumeFrozen)
    binary.Read(r, binary.BigEndian, &m.LongAmountFrozen)
}

type  CThostFtdcQryLimitPosiField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    ExchangeID string;
}

func (m* CThostFtdcQryLimitPosiField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcQryLimitPosiField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcBrokerLimitPosiField struct {
    BrokerID string;
    InstrumentID string;
    ExchangeID string;
    TotalVolume float64;
    LongVolume float64;
    TotalVolumeFrozen float64;
    LongVolumeFrozen float64;
}

func (m* CThostFtdcBrokerLimitPosiField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
    binary.Write(w, binary.BigEndian, &m.TotalVolume)
    binary.Write(w, binary.BigEndian, &m.LongVolume)
    binary.Write(w, binary.BigEndian, &m.TotalVolumeFrozen)
    binary.Write(w, binary.BigEndian, &m.LongVolumeFrozen)
}

func (m* CThostFtdcBrokerLimitPosiField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.TotalVolume)
    binary.Read(r, binary.BigEndian, &m.LongVolume)
    binary.Read(r, binary.BigEndian, &m.TotalVolumeFrozen)
    binary.Read(r, binary.BigEndian, &m.LongVolumeFrozen)
}

type  CThostFtdcQryBrokerLimitPosiField struct {
    BrokerID string;
    InstrumentID string;
    ExchangeID string;
}

func (m* CThostFtdcQryBrokerLimitPosiField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcQryBrokerLimitPosiField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcLimitPosiSField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    ExchangeID string;
    TotalVolume int32;
    OpenVolume int32;
    TotalVolumeFrozen int32;
    OpenVolumeFrozen int32;
}

func (m* CThostFtdcLimitPosiSField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
    binary.Write(w, binary.BigEndian, &m.TotalVolume)
    binary.Write(w, binary.BigEndian, &m.OpenVolume)
    binary.Write(w, binary.BigEndian, &m.TotalVolumeFrozen)
    binary.Write(w, binary.BigEndian, &m.OpenVolumeFrozen)
}

func (m* CThostFtdcLimitPosiSField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.TotalVolume)
    binary.Read(r, binary.BigEndian, &m.OpenVolume)
    binary.Read(r, binary.BigEndian, &m.TotalVolumeFrozen)
    binary.Read(r, binary.BigEndian, &m.OpenVolumeFrozen)
}

type  CThostFtdcQryLimitPosiSField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    ExchangeID string;
}

func (m* CThostFtdcQryLimitPosiSField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcQryLimitPosiSField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcLimitPosiParamField struct {
    InstrumentID string;
    InvestorRange uint8;
    BrokerID string;
    InvestorID string;
    ExchangeID string;
    TotalVolume int32;
    LongVolume int32;
    OpenVolume int32;
    LongAmount float64;
}

func (m* CThostFtdcLimitPosiParamField) Marshal(w io.Writer) {
    writeLen(w, m.InstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.InvestorRange)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.ExchangeID, 9)
    binary.Write(w, binary.BigEndian, &m.TotalVolume)
    binary.Write(w, binary.BigEndian, &m.LongVolume)
    binary.Write(w, binary.BigEndian, &m.OpenVolume)
    binary.Write(w, binary.BigEndian, &m.LongAmount)
}

func (m* CThostFtdcLimitPosiParamField) Unmarshal(r io.Reader) {
    m.InstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.InvestorRange)
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.ExchangeID = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.TotalVolume)
    binary.Read(r, binary.BigEndian, &m.LongVolume)
    binary.Read(r, binary.BigEndian, &m.OpenVolume)
    binary.Read(r, binary.BigEndian, &m.LongAmount)
}

type  CThostFtdcBrokerLimitPosiParamField struct {
    BrokerID string;
    InstrumentID string;
    ExchangeID string;
    TotalVolume float64;
    LongVolume float64;
}

func (m* CThostFtdcBrokerLimitPosiParamField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
    binary.Write(w, binary.BigEndian, &m.TotalVolume)
    binary.Write(w, binary.BigEndian, &m.LongVolume)
}

func (m* CThostFtdcBrokerLimitPosiParamField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.TotalVolume)
    binary.Read(r, binary.BigEndian, &m.LongVolume)
}

type  CThostFtdcLimitPosiParamSField struct {
    InstrumentID string;
    InvestorRange uint8;
    BrokerID string;
    InvestorID string;
    ExchangeID string;
    TotalVolume int32;
    OpenVolume int32;
}

func (m* CThostFtdcLimitPosiParamSField) Marshal(w io.Writer) {
    writeLen(w, m.InstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.InvestorRange)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.ExchangeID, 9)
    binary.Write(w, binary.BigEndian, &m.TotalVolume)
    binary.Write(w, binary.BigEndian, &m.OpenVolume)
}

func (m* CThostFtdcLimitPosiParamSField) Unmarshal(r io.Reader) {
    m.InstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.InvestorRange)
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.ExchangeID = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.TotalVolume)
    binary.Read(r, binary.BigEndian, &m.OpenVolume)
}

type  CThostFtdcInputStockDisposalActionField struct {
    BrokerID string;
    InvestorID string;
    StockDisposalActionRef int32;
    StockDisposalRef string;
    RequestID int32;
    FrontID int32;
    SessionID int32;
    ExchangeID string;
    StockDisposalSysID string;
    ActionFlag uint8;
    UserID string;
    InstrumentID string;
}

func (m* CThostFtdcInputStockDisposalActionField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.StockDisposalActionRef)
    writeLen(w, m.StockDisposalRef, 13)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    binary.Write(w, binary.BigEndian, &m.FrontID)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.StockDisposalSysID, 21)
    binary.Write(w, binary.BigEndian, &m.ActionFlag)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.InstrumentID, 31)
}

func (m* CThostFtdcInputStockDisposalActionField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.StockDisposalActionRef)
    m.StockDisposalRef = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    binary.Read(r, binary.BigEndian, &m.FrontID)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.ExchangeID = readLen(r, 9)
    m.StockDisposalSysID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.ActionFlag)
    m.UserID = readLen(r, 16)
    m.InstrumentID = readLen(r, 31)
}

type  CThostFtdcStockDisposalActionField struct {
    BrokerID string;
    InvestorID string;
    StockDisposalActionRef int32;
    StockDisposalRef string;
    RequestID int32;
    FrontID int32;
    SessionID int32;
    ExchangeID string;
    StockDisposalSysID string;
    ActionFlag uint8;
    ActionDate string;
    ActionTime string;
    TraderID string;
    InstallID int32;
    StockDisposalLocalID string;
    ActionLocalID string;
    ParticipantID string;
    ClientID string;
    BusinessUnit string;
    OrderActionStatus uint8;
    UserID string;
    ActionType uint8;
    StatusMsg string;
    InstrumentID string;
    BranchID string;
}

func (m* CThostFtdcStockDisposalActionField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.StockDisposalActionRef)
    writeLen(w, m.StockDisposalRef, 13)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    binary.Write(w, binary.BigEndian, &m.FrontID)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.StockDisposalSysID, 21)
    binary.Write(w, binary.BigEndian, &m.ActionFlag)
    writeLen(w, m.ActionDate, 9)
    writeLen(w, m.ActionTime, 9)
    writeLen(w, m.TraderID, 21)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    writeLen(w, m.StockDisposalLocalID, 13)
    writeLen(w, m.ActionLocalID, 13)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.BusinessUnit, 21)
    binary.Write(w, binary.BigEndian, &m.OrderActionStatus)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.ActionType)
    writeLen(w, m.StatusMsg, 81)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.BranchID, 9)
}

func (m* CThostFtdcStockDisposalActionField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.StockDisposalActionRef)
    m.StockDisposalRef = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    binary.Read(r, binary.BigEndian, &m.FrontID)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.ExchangeID = readLen(r, 9)
    m.StockDisposalSysID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.ActionFlag)
    m.ActionDate = readLen(r, 9)
    m.ActionTime = readLen(r, 9)
    m.TraderID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    m.StockDisposalLocalID = readLen(r, 13)
    m.ActionLocalID = readLen(r, 13)
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.BusinessUnit = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.OrderActionStatus)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.ActionType)
    m.StatusMsg = readLen(r, 81)
    m.InstrumentID = readLen(r, 31)
    m.BranchID = readLen(r, 9)
}

type  CThostFtdcQryStockDisposalActionField struct {
    BrokerID string;
    InvestorID string;
    ExchangeID string;
}

func (m* CThostFtdcQryStockDisposalActionField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcQryStockDisposalActionField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcExchangeStockDisposalActionField struct {
    ExchangeID string;
    StockDisposalSysID string;
    ActionFlag uint8;
    ActionDate string;
    ActionTime string;
    TraderID string;
    InstallID int32;
    StockDisposalLocalID string;
    ActionLocalID string;
    ParticipantID string;
    ClientID string;
    BusinessUnit string;
    OrderActionStatus uint8;
    UserID string;
    ActionType uint8;
    BranchID string;
}

func (m* CThostFtdcExchangeStockDisposalActionField) Marshal(w io.Writer) {
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.StockDisposalSysID, 21)
    binary.Write(w, binary.BigEndian, &m.ActionFlag)
    writeLen(w, m.ActionDate, 9)
    writeLen(w, m.ActionTime, 9)
    writeLen(w, m.TraderID, 21)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    writeLen(w, m.StockDisposalLocalID, 13)
    writeLen(w, m.ActionLocalID, 13)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.BusinessUnit, 21)
    binary.Write(w, binary.BigEndian, &m.OrderActionStatus)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.ActionType)
    writeLen(w, m.BranchID, 9)
}

func (m* CThostFtdcExchangeStockDisposalActionField) Unmarshal(r io.Reader) {
    m.ExchangeID = readLen(r, 9)
    m.StockDisposalSysID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.ActionFlag)
    m.ActionDate = readLen(r, 9)
    m.ActionTime = readLen(r, 9)
    m.TraderID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    m.StockDisposalLocalID = readLen(r, 13)
    m.ActionLocalID = readLen(r, 13)
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.BusinessUnit = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.OrderActionStatus)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.ActionType)
    m.BranchID = readLen(r, 9)
}

type  CThostFtdcQryExchangeStockDisposalActionField struct {
    ParticipantID string;
    ClientID string;
    ExchangeID string;
    TraderID string;
}

func (m* CThostFtdcQryExchangeStockDisposalActionField) Marshal(w io.Writer) {
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.TraderID, 21)
}

func (m* CThostFtdcQryExchangeStockDisposalActionField) Unmarshal(r io.Reader) {
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.ExchangeID = readLen(r, 9)
    m.TraderID = readLen(r, 21)
}

type  CThostFtdcQryErrStockDisposalActionField struct {
    BrokerID string;
    InvestorID string;
}

func (m* CThostFtdcQryErrStockDisposalActionField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
}

func (m* CThostFtdcQryErrStockDisposalActionField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
}

type  CThostFtdcExchangeStockDisposalActionErrorField struct {
    ExchangeID string;
    StockDisposalSysID string;
    TraderID string;
    InstallID int32;
    StockDisposalLocalID string;
    ActionLocalID string;
    ErrorID int32;
    ErrorMsg string;
    BrokerID string;
}

func (m* CThostFtdcExchangeStockDisposalActionErrorField) Marshal(w io.Writer) {
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.StockDisposalSysID, 21)
    writeLen(w, m.TraderID, 21)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    writeLen(w, m.StockDisposalLocalID, 13)
    writeLen(w, m.ActionLocalID, 13)
    binary.Write(w, binary.BigEndian, &m.ErrorID)
    writeLen(w, m.ErrorMsg, 81)
    writeLen(w, m.BrokerID, 11)
}

func (m* CThostFtdcExchangeStockDisposalActionErrorField) Unmarshal(r io.Reader) {
    m.ExchangeID = readLen(r, 9)
    m.StockDisposalSysID = readLen(r, 21)
    m.TraderID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    m.StockDisposalLocalID = readLen(r, 13)
    m.ActionLocalID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.ErrorID)
    m.ErrorMsg = readLen(r, 81)
    m.BrokerID = readLen(r, 11)
}

type  CThostFtdcErrStockDisposalActionField struct {
    BrokerID string;
    InvestorID string;
    StockDisposalActionRef int32;
    StockDisposalRef string;
    RequestID int32;
    FrontID int32;
    SessionID int32;
    ExchangeID string;
    StockDisposalSysID string;
    ActionFlag uint8;
    UserID string;
    InstrumentID string;
    ErrorID int32;
    ErrorMsg string;
}

func (m* CThostFtdcErrStockDisposalActionField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.StockDisposalActionRef)
    writeLen(w, m.StockDisposalRef, 13)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    binary.Write(w, binary.BigEndian, &m.FrontID)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.StockDisposalSysID, 21)
    binary.Write(w, binary.BigEndian, &m.ActionFlag)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.InstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.ErrorID)
    writeLen(w, m.ErrorMsg, 81)
}

func (m* CThostFtdcErrStockDisposalActionField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.StockDisposalActionRef)
    m.StockDisposalRef = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    binary.Read(r, binary.BigEndian, &m.FrontID)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.ExchangeID = readLen(r, 9)
    m.StockDisposalSysID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.ActionFlag)
    m.UserID = readLen(r, 16)
    m.InstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.ErrorID)
    m.ErrorMsg = readLen(r, 81)
}

type  CThostFtdcInvestorLevelField struct {
    BrokerID string;
    InvestorID string;
    ExchangeID string;
    LevelType uint8;
}

func (m* CThostFtdcInvestorLevelField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.ExchangeID, 9)
    binary.Write(w, binary.BigEndian, &m.LevelType)
}

func (m* CThostFtdcInvestorLevelField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.ExchangeID = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.LevelType)
}

type  CThostFtdcCombInstrumentGuardField struct {
    BrokerID string;
    InstrumentID string;
    GuarantRatio float64;
    ExchangeID string;
}

func (m* CThostFtdcCombInstrumentGuardField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.GuarantRatio)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcCombInstrumentGuardField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.GuarantRatio)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcQryCombInstrumentGuardField struct {
    BrokerID string;
    InstrumentID string;
    ExchangeID string;
}

func (m* CThostFtdcQryCombInstrumentGuardField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcQryCombInstrumentGuardField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcInputCombActionField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    CombActionRef string;
    UserID string;
    Direction uint8;
    Volume int32;
    CombDirection uint8;
    HedgeFlag uint8;
    ExchangeID string;
    IPAddress string;
    MacAddress string;
    InvestUnitID string;
}

func (m* CThostFtdcInputCombActionField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.CombActionRef, 13)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.Direction)
    binary.Write(w, binary.BigEndian, &m.Volume)
    binary.Write(w, binary.BigEndian, &m.CombDirection)
    binary.Write(w, binary.BigEndian, &m.HedgeFlag)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
    writeLen(w, m.InvestUnitID, 17)
}

func (m* CThostFtdcInputCombActionField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.CombActionRef = readLen(r, 13)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.Direction)
    binary.Read(r, binary.BigEndian, &m.Volume)
    binary.Read(r, binary.BigEndian, &m.CombDirection)
    binary.Read(r, binary.BigEndian, &m.HedgeFlag)
    m.ExchangeID = readLen(r, 9)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
    m.InvestUnitID = readLen(r, 17)
}

type  CThostFtdcCombActionField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    CombActionRef string;
    UserID string;
    Direction uint8;
    Volume int32;
    CombDirection uint8;
    HedgeFlag uint8;
    ActionLocalID string;
    ExchangeID string;
    ParticipantID string;
    ClientID string;
    ExchangeInstID string;
    TraderID string;
    InstallID int32;
    ActionStatus uint8;
    NotifySequence int32;
    TradingDay string;
    SettlementID int32;
    SequenceNo int32;
    FrontID int32;
    SessionID int32;
    UserProductInfo string;
    StatusMsg string;
    IPAddress string;
    MacAddress string;
    ComTradeID string;
    BranchID string;
    InvestUnitID string;
}

func (m* CThostFtdcCombActionField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.CombActionRef, 13)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.Direction)
    binary.Write(w, binary.BigEndian, &m.Volume)
    binary.Write(w, binary.BigEndian, &m.CombDirection)
    binary.Write(w, binary.BigEndian, &m.HedgeFlag)
    writeLen(w, m.ActionLocalID, 13)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.ExchangeInstID, 31)
    writeLen(w, m.TraderID, 21)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    binary.Write(w, binary.BigEndian, &m.ActionStatus)
    binary.Write(w, binary.BigEndian, &m.NotifySequence)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.SettlementID)
    binary.Write(w, binary.BigEndian, &m.SequenceNo)
    binary.Write(w, binary.BigEndian, &m.FrontID)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.UserProductInfo, 11)
    writeLen(w, m.StatusMsg, 81)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
    writeLen(w, m.ComTradeID, 21)
    writeLen(w, m.BranchID, 9)
    writeLen(w, m.InvestUnitID, 17)
}

func (m* CThostFtdcCombActionField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.CombActionRef = readLen(r, 13)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.Direction)
    binary.Read(r, binary.BigEndian, &m.Volume)
    binary.Read(r, binary.BigEndian, &m.CombDirection)
    binary.Read(r, binary.BigEndian, &m.HedgeFlag)
    m.ActionLocalID = readLen(r, 13)
    m.ExchangeID = readLen(r, 9)
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.ExchangeInstID = readLen(r, 31)
    m.TraderID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    binary.Read(r, binary.BigEndian, &m.ActionStatus)
    binary.Read(r, binary.BigEndian, &m.NotifySequence)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.SettlementID)
    binary.Read(r, binary.BigEndian, &m.SequenceNo)
    binary.Read(r, binary.BigEndian, &m.FrontID)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.UserProductInfo = readLen(r, 11)
    m.StatusMsg = readLen(r, 81)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
    m.ComTradeID = readLen(r, 21)
    m.BranchID = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
}

type  CThostFtdcQryCombActionField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    ExchangeID string;
    InvestUnitID string;
}

func (m* CThostFtdcQryCombActionField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.InvestUnitID, 17)
}

func (m* CThostFtdcQryCombActionField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
}

type  CThostFtdcExchangeCombActionField struct {
    Direction uint8;
    Volume int32;
    CombDirection uint8;
    HedgeFlag uint8;
    ActionLocalID string;
    ExchangeID string;
    ParticipantID string;
    ClientID string;
    ExchangeInstID string;
    TraderID string;
    InstallID int32;
    ActionStatus uint8;
    NotifySequence int32;
    TradingDay string;
    SettlementID int32;
    SequenceNo int32;
    IPAddress string;
    MacAddress string;
    ComTradeID string;
    BranchID string;
}

func (m* CThostFtdcExchangeCombActionField) Marshal(w io.Writer) {
    binary.Write(w, binary.BigEndian, &m.Direction)
    binary.Write(w, binary.BigEndian, &m.Volume)
    binary.Write(w, binary.BigEndian, &m.CombDirection)
    binary.Write(w, binary.BigEndian, &m.HedgeFlag)
    writeLen(w, m.ActionLocalID, 13)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.ExchangeInstID, 31)
    writeLen(w, m.TraderID, 21)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    binary.Write(w, binary.BigEndian, &m.ActionStatus)
    binary.Write(w, binary.BigEndian, &m.NotifySequence)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.SettlementID)
    binary.Write(w, binary.BigEndian, &m.SequenceNo)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
    writeLen(w, m.ComTradeID, 21)
    writeLen(w, m.BranchID, 9)
}

func (m* CThostFtdcExchangeCombActionField) Unmarshal(r io.Reader) {
    binary.Read(r, binary.BigEndian, &m.Direction)
    binary.Read(r, binary.BigEndian, &m.Volume)
    binary.Read(r, binary.BigEndian, &m.CombDirection)
    binary.Read(r, binary.BigEndian, &m.HedgeFlag)
    m.ActionLocalID = readLen(r, 13)
    m.ExchangeID = readLen(r, 9)
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.ExchangeInstID = readLen(r, 31)
    m.TraderID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    binary.Read(r, binary.BigEndian, &m.ActionStatus)
    binary.Read(r, binary.BigEndian, &m.NotifySequence)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.SettlementID)
    binary.Read(r, binary.BigEndian, &m.SequenceNo)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
    m.ComTradeID = readLen(r, 21)
    m.BranchID = readLen(r, 9)
}

type  CThostFtdcQryExchangeCombActionField struct {
    ParticipantID string;
    ClientID string;
    ExchangeInstID string;
    ExchangeID string;
    TraderID string;
}

func (m* CThostFtdcQryExchangeCombActionField) Marshal(w io.Writer) {
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.ExchangeInstID, 31)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.TraderID, 21)
}

func (m* CThostFtdcQryExchangeCombActionField) Unmarshal(r io.Reader) {
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.ExchangeInstID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
    m.TraderID = readLen(r, 21)
}

type  CThostFtdcProductExchRateField struct {
    ProductID string;
    QuoteCurrencyID string;
    ExchangeRate float64;
    ExchangeID string;
}

func (m* CThostFtdcProductExchRateField) Marshal(w io.Writer) {
    writeLen(w, m.ProductID, 31)
    writeLen(w, m.QuoteCurrencyID, 4)
    binary.Write(w, binary.BigEndian, &m.ExchangeRate)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcProductExchRateField) Unmarshal(r io.Reader) {
    m.ProductID = readLen(r, 31)
    m.QuoteCurrencyID = readLen(r, 4)
    binary.Read(r, binary.BigEndian, &m.ExchangeRate)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcQryProductExchRateField struct {
    ProductID string;
    ExchangeID string;
}

func (m* CThostFtdcQryProductExchRateField) Marshal(w io.Writer) {
    writeLen(w, m.ProductID, 31)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcQryProductExchRateField) Unmarshal(r io.Reader) {
    m.ProductID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcInputDesignateField struct {
    BrokerID string;
    InvestorID string;
    DesignateRef string;
    UserID string;
    DesignateType uint8;
    ExchangeID string;
    PBU string;
}

func (m* CThostFtdcInputDesignateField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.DesignateRef, 13)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.DesignateType)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.PBU, 21)
}

func (m* CThostFtdcInputDesignateField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.DesignateRef = readLen(r, 13)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.DesignateType)
    m.ExchangeID = readLen(r, 9)
    m.PBU = readLen(r, 21)
}

type  CThostFtdcDesignateField struct {
    BrokerID string;
    InvestorID string;
    DesignateRef string;
    UserID string;
    DesignateType uint8;
    DesignateLocalID string;
    ExchangeID string;
    ParticipantID string;
    ClientID string;
    TraderID string;
    InstallID int32;
    DesignateStatus uint8;
    NotifySequence int32;
    TradingDay string;
    SettlementID int32;
    InsertDate string;
    InsertTime string;
    FrontID int32;
    SessionID int32;
    UserProductInfo string;
    StatusMsg string;
    BranchID string;
    PBU string;
}

func (m* CThostFtdcDesignateField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.DesignateRef, 13)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.DesignateType)
    writeLen(w, m.DesignateLocalID, 13)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.TraderID, 21)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    binary.Write(w, binary.BigEndian, &m.DesignateStatus)
    binary.Write(w, binary.BigEndian, &m.NotifySequence)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.SettlementID)
    writeLen(w, m.InsertDate, 9)
    writeLen(w, m.InsertTime, 9)
    binary.Write(w, binary.BigEndian, &m.FrontID)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.UserProductInfo, 11)
    writeLen(w, m.StatusMsg, 81)
    writeLen(w, m.BranchID, 9)
    writeLen(w, m.PBU, 21)
}

func (m* CThostFtdcDesignateField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.DesignateRef = readLen(r, 13)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.DesignateType)
    m.DesignateLocalID = readLen(r, 13)
    m.ExchangeID = readLen(r, 9)
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.TraderID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    binary.Read(r, binary.BigEndian, &m.DesignateStatus)
    binary.Read(r, binary.BigEndian, &m.NotifySequence)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.SettlementID)
    m.InsertDate = readLen(r, 9)
    m.InsertTime = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.FrontID)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.UserProductInfo = readLen(r, 11)
    m.StatusMsg = readLen(r, 81)
    m.BranchID = readLen(r, 9)
    m.PBU = readLen(r, 21)
}

type  CThostFtdcQryDesignateField struct {
    BrokerID string;
    InvestorID string;
    ExchangeID string;
}

func (m* CThostFtdcQryDesignateField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcQryDesignateField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcExchangeDesignateField struct {
    DesignateType uint8;
    DesignateLocalID string;
    ExchangeID string;
    ParticipantID string;
    ClientID string;
    TraderID string;
    InstallID int32;
    DesignateStatus uint8;
    NotifySequence int32;
    TradingDay string;
    SettlementID int32;
    InsertDate string;
    InsertTime string;
    BranchID string;
}

func (m* CThostFtdcExchangeDesignateField) Marshal(w io.Writer) {
    binary.Write(w, binary.BigEndian, &m.DesignateType)
    writeLen(w, m.DesignateLocalID, 13)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.TraderID, 21)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    binary.Write(w, binary.BigEndian, &m.DesignateStatus)
    binary.Write(w, binary.BigEndian, &m.NotifySequence)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.SettlementID)
    writeLen(w, m.InsertDate, 9)
    writeLen(w, m.InsertTime, 9)
    writeLen(w, m.BranchID, 9)
}

func (m* CThostFtdcExchangeDesignateField) Unmarshal(r io.Reader) {
    binary.Read(r, binary.BigEndian, &m.DesignateType)
    m.DesignateLocalID = readLen(r, 13)
    m.ExchangeID = readLen(r, 9)
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.TraderID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    binary.Read(r, binary.BigEndian, &m.DesignateStatus)
    binary.Read(r, binary.BigEndian, &m.NotifySequence)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.SettlementID)
    m.InsertDate = readLen(r, 9)
    m.InsertTime = readLen(r, 9)
    m.BranchID = readLen(r, 9)
}

type  CThostFtdcInputStockDisposalField struct {
    BrokerID string;
    InvestorID string;
    StockDisposalRef string;
    UserID string;
    InstrumentID string;
    Volume int32;
    StockDisposalType uint8;
    ExchangeID string;
}

func (m* CThostFtdcInputStockDisposalField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.StockDisposalRef, 13)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.InstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.Volume)
    binary.Write(w, binary.BigEndian, &m.StockDisposalType)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcInputStockDisposalField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.StockDisposalRef = readLen(r, 13)
    m.UserID = readLen(r, 16)
    m.InstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.Volume)
    binary.Read(r, binary.BigEndian, &m.StockDisposalType)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcStockDisposalField struct {
    BrokerID string;
    InvestorID string;
    StockDisposalRef string;
    UserID string;
    InstrumentID string;
    Volume int32;
    StockDisposalType uint8;
    StockDisposalLocalID string;
    ExchangeID string;
    ExchangeInstID string;
    ParticipantID string;
    ClientID string;
    TraderID string;
    InstallID int32;
    StockDisposalStatus uint8;
    NotifySequence int32;
    TradingDay string;
    SettlementID int32;
    InsertDate string;
    InsertTime string;
    FrontID int32;
    SessionID int32;
    UserProductInfo string;
    StatusMsg string;
    BranchID string;
    StockDisposalSysID string;
    BusinessUnit string;
}

func (m* CThostFtdcStockDisposalField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.StockDisposalRef, 13)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.InstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.Volume)
    binary.Write(w, binary.BigEndian, &m.StockDisposalType)
    writeLen(w, m.StockDisposalLocalID, 13)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ExchangeInstID, 31)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.TraderID, 21)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    binary.Write(w, binary.BigEndian, &m.StockDisposalStatus)
    binary.Write(w, binary.BigEndian, &m.NotifySequence)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.SettlementID)
    writeLen(w, m.InsertDate, 9)
    writeLen(w, m.InsertTime, 9)
    binary.Write(w, binary.BigEndian, &m.FrontID)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.UserProductInfo, 11)
    writeLen(w, m.StatusMsg, 81)
    writeLen(w, m.BranchID, 9)
    writeLen(w, m.StockDisposalSysID, 21)
    writeLen(w, m.BusinessUnit, 21)
}

func (m* CThostFtdcStockDisposalField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.StockDisposalRef = readLen(r, 13)
    m.UserID = readLen(r, 16)
    m.InstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.Volume)
    binary.Read(r, binary.BigEndian, &m.StockDisposalType)
    m.StockDisposalLocalID = readLen(r, 13)
    m.ExchangeID = readLen(r, 9)
    m.ExchangeInstID = readLen(r, 31)
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.TraderID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    binary.Read(r, binary.BigEndian, &m.StockDisposalStatus)
    binary.Read(r, binary.BigEndian, &m.NotifySequence)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.SettlementID)
    m.InsertDate = readLen(r, 9)
    m.InsertTime = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.FrontID)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.UserProductInfo = readLen(r, 11)
    m.StatusMsg = readLen(r, 81)
    m.BranchID = readLen(r, 9)
    m.StockDisposalSysID = readLen(r, 21)
    m.BusinessUnit = readLen(r, 21)
}

type  CThostFtdcQryStockDisposalField struct {
    BrokerID string;
    InvestorID string;
    ExchangeID string;
}

func (m* CThostFtdcQryStockDisposalField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcQryStockDisposalField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcExchangeStockDisposalField struct {
    Volume int32;
    StockDisposalType uint8;
    StockDisposalLocalID string;
    ExchangeID string;
    ExchangeInstID string;
    ParticipantID string;
    ClientID string;
    TraderID string;
    InstallID int32;
    StockDisposalStatus uint8;
    NotifySequence int32;
    TradingDay string;
    SettlementID int32;
    InsertDate string;
    InsertTime string;
    BranchID string;
    StockDisposalSysID string;
    BusinessUnit string;
}

func (m* CThostFtdcExchangeStockDisposalField) Marshal(w io.Writer) {
    binary.Write(w, binary.BigEndian, &m.Volume)
    binary.Write(w, binary.BigEndian, &m.StockDisposalType)
    writeLen(w, m.StockDisposalLocalID, 13)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ExchangeInstID, 31)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.TraderID, 21)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    binary.Write(w, binary.BigEndian, &m.StockDisposalStatus)
    binary.Write(w, binary.BigEndian, &m.NotifySequence)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.SettlementID)
    writeLen(w, m.InsertDate, 9)
    writeLen(w, m.InsertTime, 9)
    writeLen(w, m.BranchID, 9)
    writeLen(w, m.StockDisposalSysID, 21)
    writeLen(w, m.BusinessUnit, 21)
}

func (m* CThostFtdcExchangeStockDisposalField) Unmarshal(r io.Reader) {
    binary.Read(r, binary.BigEndian, &m.Volume)
    binary.Read(r, binary.BigEndian, &m.StockDisposalType)
    m.StockDisposalLocalID = readLen(r, 13)
    m.ExchangeID = readLen(r, 9)
    m.ExchangeInstID = readLen(r, 31)
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.TraderID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    binary.Read(r, binary.BigEndian, &m.StockDisposalStatus)
    binary.Read(r, binary.BigEndian, &m.NotifySequence)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.SettlementID)
    m.InsertDate = readLen(r, 9)
    m.InsertTime = readLen(r, 9)
    m.BranchID = readLen(r, 9)
    m.StockDisposalSysID = readLen(r, 21)
    m.BusinessUnit = readLen(r, 21)
}

type  CThostFtdcQryInvestorLevelField struct {
    BrokerID string;
    InvestorID string;
    ExchangeID string;
}

func (m* CThostFtdcQryInvestorLevelField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcQryInvestorLevelField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcQryForQuoteParamField struct {
    BrokerID string;
    InstrumentID string;
    ExchangeID string;
}

func (m* CThostFtdcQryForQuoteParamField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcQryForQuoteParamField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcForQuoteParamField struct {
    BrokerID string;
    InstrumentID string;
    ExchangeID string;
    LastPrice float64;
    PriceInterval float64;
}

func (m* CThostFtdcForQuoteParamField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
    binary.Write(w, binary.BigEndian, &m.LastPrice)
    binary.Write(w, binary.BigEndian, &m.PriceInterval)
}

func (m* CThostFtdcForQuoteParamField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.LastPrice)
    binary.Read(r, binary.BigEndian, &m.PriceInterval)
}

type  CThostFtdcQryExecFreezeField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    ExchangeID string;
}

func (m* CThostFtdcQryExecFreezeField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcQryExecFreezeField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcExecFreezeField struct {
    InstrumentID string;
    ExchangeID string;
    BrokerID string;
    InvestorID string;
    PosiDirection uint8;
    OptionsType uint8;
    Volume int32;
    FrozenAmount float64;
}

func (m* CThostFtdcExecFreezeField) Marshal(w io.Writer) {
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.PosiDirection)
    binary.Write(w, binary.BigEndian, &m.OptionsType)
    binary.Write(w, binary.BigEndian, &m.Volume)
    binary.Write(w, binary.BigEndian, &m.FrozenAmount)
}

func (m* CThostFtdcExecFreezeField) Unmarshal(r io.Reader) {
    m.InstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.PosiDirection)
    binary.Read(r, binary.BigEndian, &m.OptionsType)
    binary.Read(r, binary.BigEndian, &m.Volume)
    binary.Read(r, binary.BigEndian, &m.FrozenAmount)
}

type  CThostFtdcMMOptionInstrCommRateField struct {
    InstrumentID string;
    InvestorRange uint8;
    BrokerID string;
    InvestorID string;
    OpenRatioByMoney float64;
    OpenRatioByVolume float64;
    CloseRatioByMoney float64;
    CloseRatioByVolume float64;
    CloseTodayRatioByMoney float64;
    CloseTodayRatioByVolume float64;
    StrikeRatioByMoney float64;
    StrikeRatioByVolume float64;
    ExchangeID string;
}

func (m* CThostFtdcMMOptionInstrCommRateField) Marshal(w io.Writer) {
    writeLen(w, m.InstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.InvestorRange)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.OpenRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.OpenRatioByVolume)
    binary.Write(w, binary.BigEndian, &m.CloseRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.CloseRatioByVolume)
    binary.Write(w, binary.BigEndian, &m.CloseTodayRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.CloseTodayRatioByVolume)
    binary.Write(w, binary.BigEndian, &m.StrikeRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.StrikeRatioByVolume)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcMMOptionInstrCommRateField) Unmarshal(r io.Reader) {
    m.InstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.InvestorRange)
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.OpenRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.OpenRatioByVolume)
    binary.Read(r, binary.BigEndian, &m.CloseRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.CloseRatioByVolume)
    binary.Read(r, binary.BigEndian, &m.CloseTodayRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.CloseTodayRatioByVolume)
    binary.Read(r, binary.BigEndian, &m.StrikeRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.StrikeRatioByVolume)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcQryMMOptionInstrCommRateField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    ExchangeID string;
}

func (m* CThostFtdcQryMMOptionInstrCommRateField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcQryMMOptionInstrCommRateField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcMMInstrumentCommissionRateField struct {
    InstrumentID string;
    InvestorRange uint8;
    BrokerID string;
    InvestorID string;
    OpenRatioByMoney float64;
    OpenRatioByVolume float64;
    CloseRatioByMoney float64;
    CloseRatioByVolume float64;
    CloseTodayRatioByMoney float64;
    CloseTodayRatioByVolume float64;
    ExchangeID string;
}

func (m* CThostFtdcMMInstrumentCommissionRateField) Marshal(w io.Writer) {
    writeLen(w, m.InstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.InvestorRange)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.OpenRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.OpenRatioByVolume)
    binary.Write(w, binary.BigEndian, &m.CloseRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.CloseRatioByVolume)
    binary.Write(w, binary.BigEndian, &m.CloseTodayRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.CloseTodayRatioByVolume)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcMMInstrumentCommissionRateField) Unmarshal(r io.Reader) {
    m.InstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.InvestorRange)
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.OpenRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.OpenRatioByVolume)
    binary.Read(r, binary.BigEndian, &m.CloseRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.CloseRatioByVolume)
    binary.Read(r, binary.BigEndian, &m.CloseTodayRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.CloseTodayRatioByVolume)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcQryMMInstrumentCommissionRateField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    ExchangeID string;
}

func (m* CThostFtdcQryMMInstrumentCommissionRateField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcQryMMInstrumentCommissionRateField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcInstrumentOrderCommRateField struct {
    InstrumentID string;
    InvestorRange uint8;
    BrokerID string;
    InvestorID string;
    HedgeFlag uint8;
    OrderCommByVolume float64;
    OrderActionCommByVolume float64;
    ExchangeID string;
    InvestUnitID string;
}

func (m* CThostFtdcInstrumentOrderCommRateField) Marshal(w io.Writer) {
    writeLen(w, m.InstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.InvestorRange)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.HedgeFlag)
    binary.Write(w, binary.BigEndian, &m.OrderCommByVolume)
    binary.Write(w, binary.BigEndian, &m.OrderActionCommByVolume)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.InvestUnitID, 17)
}

func (m* CThostFtdcInstrumentOrderCommRateField) Unmarshal(r io.Reader) {
    m.InstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.InvestorRange)
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.HedgeFlag)
    binary.Read(r, binary.BigEndian, &m.OrderCommByVolume)
    binary.Read(r, binary.BigEndian, &m.OrderActionCommByVolume)
    m.ExchangeID = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
}

type  CThostFtdcQryInstrumentOrderCommRateField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    ExchangeID string;
}

func (m* CThostFtdcQryInstrumentOrderCommRateField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcQryInstrumentOrderCommRateField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcLimitAmountField struct {
    BrokerID string;
    InvestorID string;
    ExchangeID string;
    LongAmount float64;
    LongAmountFrozen float64;
}

func (m* CThostFtdcLimitAmountField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.ExchangeID, 9)
    binary.Write(w, binary.BigEndian, &m.LongAmount)
    binary.Write(w, binary.BigEndian, &m.LongAmountFrozen)
}

func (m* CThostFtdcLimitAmountField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.ExchangeID = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.LongAmount)
    binary.Read(r, binary.BigEndian, &m.LongAmountFrozen)
}

type  CThostFtdcQryLimitAmountField struct {
    BrokerID string;
    InvestorID string;
    ExchangeID string;
}

func (m* CThostFtdcQryLimitAmountField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcQryLimitAmountField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcLimitAmountParamField struct {
    InvestorRange uint8;
    BrokerID string;
    InvestorID string;
    ExchangeID string;
    LongAmount float64;
}

func (m* CThostFtdcLimitAmountParamField) Marshal(w io.Writer) {
    binary.Write(w, binary.BigEndian, &m.InvestorRange)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.ExchangeID, 9)
    binary.Write(w, binary.BigEndian, &m.LongAmount)
}

func (m* CThostFtdcLimitAmountParamField) Unmarshal(r io.Reader) {
    binary.Read(r, binary.BigEndian, &m.InvestorRange)
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.ExchangeID = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.LongAmount)
}

type  CThostFtdcOptionInstrMarginGuardField struct {
    InstrumentID string;
    InvestorRange uint8;
    BrokerID string;
    InvestorID string;
    SShortMarginRatioByMoney float64;
    SShortMarginRatioByVolume float64;
    HShortMarginRatioByMoney float64;
    HShortMarginRatioByVolume float64;
    AShortMarginRatioByMoney float64;
    AShortMarginRatioByVolume float64;
    IsRelative int32;
    ExchangeID string;
}

func (m* CThostFtdcOptionInstrMarginGuardField) Marshal(w io.Writer) {
    writeLen(w, m.InstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.InvestorRange)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.SShortMarginRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.SShortMarginRatioByVolume)
    binary.Write(w, binary.BigEndian, &m.HShortMarginRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.HShortMarginRatioByVolume)
    binary.Write(w, binary.BigEndian, &m.AShortMarginRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.AShortMarginRatioByVolume)
    binary.Write(w, binary.BigEndian, &m.IsRelative)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcOptionInstrMarginGuardField) Unmarshal(r io.Reader) {
    m.InstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.InvestorRange)
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.SShortMarginRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.SShortMarginRatioByVolume)
    binary.Read(r, binary.BigEndian, &m.HShortMarginRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.HShortMarginRatioByVolume)
    binary.Read(r, binary.BigEndian, &m.AShortMarginRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.AShortMarginRatioByVolume)
    binary.Read(r, binary.BigEndian, &m.IsRelative)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcTradeParamField struct {
    BrokerID string;
    TradeParamID uint8;
    TradeParamValue string;
    Memo string;
}

func (m* CThostFtdcTradeParamField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    binary.Write(w, binary.BigEndian, &m.TradeParamID)
    writeLen(w, m.TradeParamValue, 256)
    writeLen(w, m.Memo, 161)
}

func (m* CThostFtdcTradeParamField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    binary.Read(r, binary.BigEndian, &m.TradeParamID)
    m.TradeParamValue = readLen(r, 256)
    m.Memo = readLen(r, 161)
}

type  CThostFtdcAuthenticationCodeField struct {
    BrokerID string;
    UserProductInfo string;
    AuthCode string;
    PreAuthCode string;
}

func (m* CThostFtdcAuthenticationCodeField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.UserProductInfo, 11)
    writeLen(w, m.AuthCode, 17)
    writeLen(w, m.PreAuthCode, 17)
}

func (m* CThostFtdcAuthenticationCodeField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.UserProductInfo = readLen(r, 11)
    m.AuthCode = readLen(r, 17)
    m.PreAuthCode = readLen(r, 17)
}

type  CThostFtdcReqSmsCodeGenerateField struct {
    BrokerID string;
    UserID string;
    Password string;
    UserProductInfo string;
    MacAddress string;
    ClientIPAddress string;
}

func (m* CThostFtdcReqSmsCodeGenerateField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.Password, 41)
    writeLen(w, m.UserProductInfo, 11)
    writeLen(w, m.MacAddress, 21)
    writeLen(w, m.ClientIPAddress, 16)
}

func (m* CThostFtdcReqSmsCodeGenerateField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.UserID = readLen(r, 16)
    m.Password = readLen(r, 41)
    m.UserProductInfo = readLen(r, 11)
    m.MacAddress = readLen(r, 21)
    m.ClientIPAddress = readLen(r, 16)
}

type  CThostFtdcRspSmsCodeGenerateField struct {
    BrokerID string;
    UserID string;
}

func (m* CThostFtdcRspSmsCodeGenerateField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.UserID, 16)
}

func (m* CThostFtdcRspSmsCodeGenerateField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.UserID = readLen(r, 16)
}

type  CThostFtdcBrokerUserAuthMethodField struct {
    BrokerID string;
    UserID string;
    IsSms int32;
}

func (m* CThostFtdcBrokerUserAuthMethodField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.IsSms)
}

func (m* CThostFtdcBrokerUserAuthMethodField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.IsSms)
}

type  CThostFtdcBrokerUserSmsCodeField struct {
    BrokerID string;
    UserID string;
    Mobile string;
    SmsCode string;
    SeqNo int32;
    Status uint8;
}

func (m* CThostFtdcBrokerUserSmsCodeField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.Mobile, 41)
    writeLen(w, m.SmsCode, 13)
    binary.Write(w, binary.BigEndian, &m.SeqNo)
    binary.Write(w, binary.BigEndian, &m.Status)
}

func (m* CThostFtdcBrokerUserSmsCodeField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.UserID = readLen(r, 16)
    m.Mobile = readLen(r, 41)
    m.SmsCode = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.SeqNo)
    binary.Read(r, binary.BigEndian, &m.Status)
}

type  CThostFtdcInstrumentMarginRateULField struct {
    InstrumentID string;
    InvestorRange uint8;
    BrokerID string;
    InvestorID string;
    HedgeFlag uint8;
    LongMarginRatioByMoney float64;
    LongMarginRatioByVolume float64;
    ShortMarginRatioByMoney float64;
    ShortMarginRatioByVolume float64;
}

func (m* CThostFtdcInstrumentMarginRateULField) Marshal(w io.Writer) {
    writeLen(w, m.InstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.InvestorRange)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.HedgeFlag)
    binary.Write(w, binary.BigEndian, &m.LongMarginRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.LongMarginRatioByVolume)
    binary.Write(w, binary.BigEndian, &m.ShortMarginRatioByMoney)
    binary.Write(w, binary.BigEndian, &m.ShortMarginRatioByVolume)
}

func (m* CThostFtdcInstrumentMarginRateULField) Unmarshal(r io.Reader) {
    m.InstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.InvestorRange)
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.HedgeFlag)
    binary.Read(r, binary.BigEndian, &m.LongMarginRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.LongMarginRatioByVolume)
    binary.Read(r, binary.BigEndian, &m.ShortMarginRatioByMoney)
    binary.Read(r, binary.BigEndian, &m.ShortMarginRatioByVolume)
}

type  CThostFtdcFutureLimitPosiParamField struct {
    InvestorRange uint8;
    BrokerID string;
    InvestorID string;
    ProductID string;
    ExchangeID string;
    SpecOpenVolume int32;
    ArbiOpenVolume int32;
    OpenVolume int32;
}

func (m* CThostFtdcFutureLimitPosiParamField) Marshal(w io.Writer) {
    binary.Write(w, binary.BigEndian, &m.InvestorRange)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.ProductID, 31)
    writeLen(w, m.ExchangeID, 9)
    binary.Write(w, binary.BigEndian, &m.SpecOpenVolume)
    binary.Write(w, binary.BigEndian, &m.ArbiOpenVolume)
    binary.Write(w, binary.BigEndian, &m.OpenVolume)
}

func (m* CThostFtdcFutureLimitPosiParamField) Unmarshal(r io.Reader) {
    binary.Read(r, binary.BigEndian, &m.InvestorRange)
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.ProductID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.SpecOpenVolume)
    binary.Read(r, binary.BigEndian, &m.ArbiOpenVolume)
    binary.Read(r, binary.BigEndian, &m.OpenVolume)
}

type  CThostFtdcLoginForbiddenIPField struct {
    IPAddress string;
}

func (m* CThostFtdcLoginForbiddenIPField) Marshal(w io.Writer) {
    writeLen(w, m.IPAddress, 16)
}

func (m* CThostFtdcLoginForbiddenIPField) Unmarshal(r io.Reader) {
    m.IPAddress = readLen(r, 16)
}

type  CThostFtdcIPListField struct {
    IPAddress string;
    IsWhite int32;
}

func (m* CThostFtdcIPListField) Marshal(w io.Writer) {
    writeLen(w, m.IPAddress, 16)
    binary.Write(w, binary.BigEndian, &m.IsWhite)
}

func (m* CThostFtdcIPListField) Unmarshal(r io.Reader) {
    m.IPAddress = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.IsWhite)
}

type  CThostFtdcInputOptionSelfCloseField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    OptionSelfCloseRef string;
    UserID string;
    Volume int32;
    RequestID int32;
    BusinessUnit string;
    HedgeFlag uint8;
    OptSelfCloseFlag uint8;
    ExchangeID string;
    InvestUnitID string;
    AccountID string;
    CurrencyID string;
    ClientID string;
    IPAddress string;
    MacAddress string;
}

func (m* CThostFtdcInputOptionSelfCloseField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.OptionSelfCloseRef, 13)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.Volume)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    writeLen(w, m.BusinessUnit, 21)
    binary.Write(w, binary.BigEndian, &m.HedgeFlag)
    binary.Write(w, binary.BigEndian, &m.OptSelfCloseFlag)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.InvestUnitID, 17)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.CurrencyID, 4)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
}

func (m* CThostFtdcInputOptionSelfCloseField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.OptionSelfCloseRef = readLen(r, 13)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.Volume)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    m.BusinessUnit = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.HedgeFlag)
    binary.Read(r, binary.BigEndian, &m.OptSelfCloseFlag)
    m.ExchangeID = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
    m.AccountID = readLen(r, 13)
    m.CurrencyID = readLen(r, 4)
    m.ClientID = readLen(r, 11)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
}

type  CThostFtdcInputOptionSelfCloseActionField struct {
    BrokerID string;
    InvestorID string;
    OptionSelfCloseActionRef int32;
    OptionSelfCloseRef string;
    RequestID int32;
    FrontID int32;
    SessionID int32;
    ExchangeID string;
    OptionSelfCloseSysID string;
    ActionFlag uint8;
    UserID string;
    InstrumentID string;
    InvestUnitID string;
    IPAddress string;
    MacAddress string;
}

func (m* CThostFtdcInputOptionSelfCloseActionField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.OptionSelfCloseActionRef)
    writeLen(w, m.OptionSelfCloseRef, 13)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    binary.Write(w, binary.BigEndian, &m.FrontID)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.OptionSelfCloseSysID, 21)
    binary.Write(w, binary.BigEndian, &m.ActionFlag)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.InvestUnitID, 17)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
}

func (m* CThostFtdcInputOptionSelfCloseActionField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.OptionSelfCloseActionRef)
    m.OptionSelfCloseRef = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    binary.Read(r, binary.BigEndian, &m.FrontID)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.ExchangeID = readLen(r, 9)
    m.OptionSelfCloseSysID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.ActionFlag)
    m.UserID = readLen(r, 16)
    m.InstrumentID = readLen(r, 31)
    m.InvestUnitID = readLen(r, 17)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
}

type  CThostFtdcOptionSelfCloseField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    OptionSelfCloseRef string;
    UserID string;
    Volume int32;
    RequestID int32;
    BusinessUnit string;
    HedgeFlag uint8;
    OptSelfCloseFlag uint8;
    OptionSelfCloseLocalID string;
    ExchangeID string;
    ParticipantID string;
    ClientID string;
    ExchangeInstID string;
    TraderID string;
    InstallID int32;
    OrderSubmitStatus uint8;
    NotifySequence int32;
    TradingDay string;
    SettlementID int32;
    OptionSelfCloseSysID string;
    InsertDate string;
    InsertTime string;
    CancelTime string;
    OptionSelfCloseStatus uint8;
    ClearingPartID string;
    SequenceNo int32;
    FrontID int32;
    SessionID int32;
    UserProductInfo string;
    StatusMsg string;
    ActiveUserID string;
    BrokerOptionSelfCloseSeq int32;
    BranchID string;
    InvestUnitID string;
    AccountID string;
    CurrencyID string;
    IPAddress string;
    MacAddress string;
}

func (m* CThostFtdcOptionSelfCloseField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.OptionSelfCloseRef, 13)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.Volume)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    writeLen(w, m.BusinessUnit, 21)
    binary.Write(w, binary.BigEndian, &m.HedgeFlag)
    binary.Write(w, binary.BigEndian, &m.OptSelfCloseFlag)
    writeLen(w, m.OptionSelfCloseLocalID, 13)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.ExchangeInstID, 31)
    writeLen(w, m.TraderID, 21)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    binary.Write(w, binary.BigEndian, &m.OrderSubmitStatus)
    binary.Write(w, binary.BigEndian, &m.NotifySequence)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.SettlementID)
    writeLen(w, m.OptionSelfCloseSysID, 21)
    writeLen(w, m.InsertDate, 9)
    writeLen(w, m.InsertTime, 9)
    writeLen(w, m.CancelTime, 9)
    binary.Write(w, binary.BigEndian, &m.OptionSelfCloseStatus)
    writeLen(w, m.ClearingPartID, 11)
    binary.Write(w, binary.BigEndian, &m.SequenceNo)
    binary.Write(w, binary.BigEndian, &m.FrontID)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.UserProductInfo, 11)
    writeLen(w, m.StatusMsg, 81)
    writeLen(w, m.ActiveUserID, 16)
    binary.Write(w, binary.BigEndian, &m.BrokerOptionSelfCloseSeq)
    writeLen(w, m.BranchID, 9)
    writeLen(w, m.InvestUnitID, 17)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.CurrencyID, 4)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
}

func (m* CThostFtdcOptionSelfCloseField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.OptionSelfCloseRef = readLen(r, 13)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.Volume)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    m.BusinessUnit = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.HedgeFlag)
    binary.Read(r, binary.BigEndian, &m.OptSelfCloseFlag)
    m.OptionSelfCloseLocalID = readLen(r, 13)
    m.ExchangeID = readLen(r, 9)
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.ExchangeInstID = readLen(r, 31)
    m.TraderID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    binary.Read(r, binary.BigEndian, &m.OrderSubmitStatus)
    binary.Read(r, binary.BigEndian, &m.NotifySequence)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.SettlementID)
    m.OptionSelfCloseSysID = readLen(r, 21)
    m.InsertDate = readLen(r, 9)
    m.InsertTime = readLen(r, 9)
    m.CancelTime = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.OptionSelfCloseStatus)
    m.ClearingPartID = readLen(r, 11)
    binary.Read(r, binary.BigEndian, &m.SequenceNo)
    binary.Read(r, binary.BigEndian, &m.FrontID)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.UserProductInfo = readLen(r, 11)
    m.StatusMsg = readLen(r, 81)
    m.ActiveUserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.BrokerOptionSelfCloseSeq)
    m.BranchID = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
    m.AccountID = readLen(r, 13)
    m.CurrencyID = readLen(r, 4)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
}

type  CThostFtdcOptionSelfCloseActionField struct {
    BrokerID string;
    InvestorID string;
    OptionSelfCloseActionRef int32;
    OptionSelfCloseRef string;
    RequestID int32;
    FrontID int32;
    SessionID int32;
    ExchangeID string;
    OptionSelfCloseSysID string;
    ActionFlag uint8;
    ActionDate string;
    ActionTime string;
    TraderID string;
    InstallID int32;
    OptionSelfCloseLocalID string;
    ActionLocalID string;
    ParticipantID string;
    ClientID string;
    BusinessUnit string;
    OrderActionStatus uint8;
    UserID string;
    StatusMsg string;
    InstrumentID string;
    BranchID string;
    InvestUnitID string;
    IPAddress string;
    MacAddress string;
}

func (m* CThostFtdcOptionSelfCloseActionField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.OptionSelfCloseActionRef)
    writeLen(w, m.OptionSelfCloseRef, 13)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    binary.Write(w, binary.BigEndian, &m.FrontID)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.OptionSelfCloseSysID, 21)
    binary.Write(w, binary.BigEndian, &m.ActionFlag)
    writeLen(w, m.ActionDate, 9)
    writeLen(w, m.ActionTime, 9)
    writeLen(w, m.TraderID, 21)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    writeLen(w, m.OptionSelfCloseLocalID, 13)
    writeLen(w, m.ActionLocalID, 13)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.BusinessUnit, 21)
    binary.Write(w, binary.BigEndian, &m.OrderActionStatus)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.StatusMsg, 81)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.BranchID, 9)
    writeLen(w, m.InvestUnitID, 17)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
}

func (m* CThostFtdcOptionSelfCloseActionField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.OptionSelfCloseActionRef)
    m.OptionSelfCloseRef = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    binary.Read(r, binary.BigEndian, &m.FrontID)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.ExchangeID = readLen(r, 9)
    m.OptionSelfCloseSysID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.ActionFlag)
    m.ActionDate = readLen(r, 9)
    m.ActionTime = readLen(r, 9)
    m.TraderID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    m.OptionSelfCloseLocalID = readLen(r, 13)
    m.ActionLocalID = readLen(r, 13)
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.BusinessUnit = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.OrderActionStatus)
    m.UserID = readLen(r, 16)
    m.StatusMsg = readLen(r, 81)
    m.InstrumentID = readLen(r, 31)
    m.BranchID = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
}

type  CThostFtdcQryOptionSelfCloseField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    ExchangeID string;
    OptionSelfCloseSysID string;
    InsertTimeStart string;
    InsertTimeEnd string;
}

func (m* CThostFtdcQryOptionSelfCloseField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.OptionSelfCloseSysID, 21)
    writeLen(w, m.InsertTimeStart, 9)
    writeLen(w, m.InsertTimeEnd, 9)
}

func (m* CThostFtdcQryOptionSelfCloseField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
    m.OptionSelfCloseSysID = readLen(r, 21)
    m.InsertTimeStart = readLen(r, 9)
    m.InsertTimeEnd = readLen(r, 9)
}

type  CThostFtdcExchangeOptionSelfCloseField struct {
    Volume int32;
    RequestID int32;
    BusinessUnit string;
    HedgeFlag uint8;
    OptSelfCloseFlag uint8;
    OptionSelfCloseLocalID string;
    ExchangeID string;
    ParticipantID string;
    ClientID string;
    ExchangeInstID string;
    TraderID string;
    InstallID int32;
    OrderSubmitStatus uint8;
    NotifySequence int32;
    TradingDay string;
    SettlementID int32;
    OptionSelfCloseSysID string;
    InsertDate string;
    InsertTime string;
    CancelTime string;
    OptionSelfCloseStatus uint8;
    ClearingPartID string;
    SequenceNo int32;
    BranchID string;
    IPAddress string;
    MacAddress string;
}

func (m* CThostFtdcExchangeOptionSelfCloseField) Marshal(w io.Writer) {
    binary.Write(w, binary.BigEndian, &m.Volume)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    writeLen(w, m.BusinessUnit, 21)
    binary.Write(w, binary.BigEndian, &m.HedgeFlag)
    binary.Write(w, binary.BigEndian, &m.OptSelfCloseFlag)
    writeLen(w, m.OptionSelfCloseLocalID, 13)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.ExchangeInstID, 31)
    writeLen(w, m.TraderID, 21)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    binary.Write(w, binary.BigEndian, &m.OrderSubmitStatus)
    binary.Write(w, binary.BigEndian, &m.NotifySequence)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.SettlementID)
    writeLen(w, m.OptionSelfCloseSysID, 21)
    writeLen(w, m.InsertDate, 9)
    writeLen(w, m.InsertTime, 9)
    writeLen(w, m.CancelTime, 9)
    binary.Write(w, binary.BigEndian, &m.OptionSelfCloseStatus)
    writeLen(w, m.ClearingPartID, 11)
    binary.Write(w, binary.BigEndian, &m.SequenceNo)
    writeLen(w, m.BranchID, 9)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
}

func (m* CThostFtdcExchangeOptionSelfCloseField) Unmarshal(r io.Reader) {
    binary.Read(r, binary.BigEndian, &m.Volume)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    m.BusinessUnit = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.HedgeFlag)
    binary.Read(r, binary.BigEndian, &m.OptSelfCloseFlag)
    m.OptionSelfCloseLocalID = readLen(r, 13)
    m.ExchangeID = readLen(r, 9)
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.ExchangeInstID = readLen(r, 31)
    m.TraderID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    binary.Read(r, binary.BigEndian, &m.OrderSubmitStatus)
    binary.Read(r, binary.BigEndian, &m.NotifySequence)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.SettlementID)
    m.OptionSelfCloseSysID = readLen(r, 21)
    m.InsertDate = readLen(r, 9)
    m.InsertTime = readLen(r, 9)
    m.CancelTime = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.OptionSelfCloseStatus)
    m.ClearingPartID = readLen(r, 11)
    binary.Read(r, binary.BigEndian, &m.SequenceNo)
    m.BranchID = readLen(r, 9)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
}

type  CThostFtdcQryOptionSelfCloseActionField struct {
    BrokerID string;
    InvestorID string;
    ExchangeID string;
}

func (m* CThostFtdcQryOptionSelfCloseActionField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcQryOptionSelfCloseActionField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcExchangeOptionSelfCloseActionField struct {
    ExchangeID string;
    OptionSelfCloseSysID string;
    ActionFlag uint8;
    ActionDate string;
    ActionTime string;
    TraderID string;
    InstallID int32;
    OptionSelfCloseLocalID string;
    ActionLocalID string;
    ParticipantID string;
    ClientID string;
    BusinessUnit string;
    OrderActionStatus uint8;
    UserID string;
    BranchID string;
    IPAddress string;
    MacAddress string;
}

func (m* CThostFtdcExchangeOptionSelfCloseActionField) Marshal(w io.Writer) {
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.OptionSelfCloseSysID, 21)
    binary.Write(w, binary.BigEndian, &m.ActionFlag)
    writeLen(w, m.ActionDate, 9)
    writeLen(w, m.ActionTime, 9)
    writeLen(w, m.TraderID, 21)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    writeLen(w, m.OptionSelfCloseLocalID, 13)
    writeLen(w, m.ActionLocalID, 13)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.BusinessUnit, 21)
    binary.Write(w, binary.BigEndian, &m.OrderActionStatus)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.BranchID, 9)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
}

func (m* CThostFtdcExchangeOptionSelfCloseActionField) Unmarshal(r io.Reader) {
    m.ExchangeID = readLen(r, 9)
    m.OptionSelfCloseSysID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.ActionFlag)
    m.ActionDate = readLen(r, 9)
    m.ActionTime = readLen(r, 9)
    m.TraderID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    m.OptionSelfCloseLocalID = readLen(r, 13)
    m.ActionLocalID = readLen(r, 13)
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.BusinessUnit = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.OrderActionStatus)
    m.UserID = readLen(r, 16)
    m.BranchID = readLen(r, 9)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
}

type  CThostFtdcSyncDelaySwapField struct {
    DelaySwapSeqNo string;
    BrokerID string;
    InvestorID string;
    FromCurrencyID string;
    FromAmount float64;
    FromFrozenSwap float64;
    ToCurrencyID string;
    ToAmount float64;
    ToFrozenSwap float64;
}

func (m* CThostFtdcSyncDelaySwapField) Marshal(w io.Writer) {
    writeLen(w, m.DelaySwapSeqNo, 15)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.FromCurrencyID, 4)
    binary.Write(w, binary.BigEndian, &m.FromAmount)
    binary.Write(w, binary.BigEndian, &m.FromFrozenSwap)
    writeLen(w, m.ToCurrencyID, 4)
    binary.Write(w, binary.BigEndian, &m.ToAmount)
    binary.Write(w, binary.BigEndian, &m.ToFrozenSwap)
}

func (m* CThostFtdcSyncDelaySwapField) Unmarshal(r io.Reader) {
    m.DelaySwapSeqNo = readLen(r, 15)
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.FromCurrencyID = readLen(r, 4)
    binary.Read(r, binary.BigEndian, &m.FromAmount)
    binary.Read(r, binary.BigEndian, &m.FromFrozenSwap)
    m.ToCurrencyID = readLen(r, 4)
    binary.Read(r, binary.BigEndian, &m.ToAmount)
    binary.Read(r, binary.BigEndian, &m.ToFrozenSwap)
}

type  CThostFtdcQrySyncDelaySwapField struct {
    BrokerID string;
    DelaySwapSeqNo string;
}

func (m* CThostFtdcQrySyncDelaySwapField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.DelaySwapSeqNo, 15)
}

func (m* CThostFtdcQrySyncDelaySwapField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.DelaySwapSeqNo = readLen(r, 15)
}

type  CThostFtdcInvestUnitField struct {
    BrokerID string;
    InvestorID string;
    InvestUnitID string;
    InvestorUnitName string;
    InvestorGroupID string;
    CommModelID string;
    MarginModelID string;
    AccountID string;
    CurrencyID string;
}

func (m* CThostFtdcInvestUnitField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InvestUnitID, 17)
    writeLen(w, m.InvestorUnitName, 81)
    writeLen(w, m.InvestorGroupID, 13)
    writeLen(w, m.CommModelID, 13)
    writeLen(w, m.MarginModelID, 13)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.CurrencyID, 4)
}

func (m* CThostFtdcInvestUnitField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InvestUnitID = readLen(r, 17)
    m.InvestorUnitName = readLen(r, 81)
    m.InvestorGroupID = readLen(r, 13)
    m.CommModelID = readLen(r, 13)
    m.MarginModelID = readLen(r, 13)
    m.AccountID = readLen(r, 13)
    m.CurrencyID = readLen(r, 4)
}

type  CThostFtdcQryInvestUnitField struct {
    BrokerID string;
    InvestorID string;
    InvestUnitID string;
}

func (m* CThostFtdcQryInvestUnitField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InvestUnitID, 17)
}

func (m* CThostFtdcQryInvestUnitField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InvestUnitID = readLen(r, 17)
}

type  CThostFtdcMarketDataField struct {
    TradingDay string;
    InstrumentID string;
    ExchangeID string;
    ExchangeInstID string;
    LastPrice float64;
    PreSettlementPrice float64;
    PreClosePrice float64;
    PreOpenInterest float64;
    OpenPrice float64;
    HighestPrice float64;
    LowestPrice float64;
    Volume int32;
    Turnover float64;
    OpenInterest float64;
    ClosePrice float64;
    SettlementPrice float64;
    UpperLimitPrice float64;
    LowerLimitPrice float64;
    PreDelta float64;
    CurrDelta float64;
    UpdateTime string;
    UpdateMillisec int32;
    ActionDay string;
}

func (m* CThostFtdcMarketDataField) Marshal(w io.Writer) {
    writeLen(w, m.TradingDay, 9)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ExchangeInstID, 31)
    binary.Write(w, binary.BigEndian, &m.LastPrice)
    binary.Write(w, binary.BigEndian, &m.PreSettlementPrice)
    binary.Write(w, binary.BigEndian, &m.PreClosePrice)
    binary.Write(w, binary.BigEndian, &m.PreOpenInterest)
    binary.Write(w, binary.BigEndian, &m.OpenPrice)
    binary.Write(w, binary.BigEndian, &m.HighestPrice)
    binary.Write(w, binary.BigEndian, &m.LowestPrice)
    binary.Write(w, binary.BigEndian, &m.Volume)
    binary.Write(w, binary.BigEndian, &m.Turnover)
    binary.Write(w, binary.BigEndian, &m.OpenInterest)
    binary.Write(w, binary.BigEndian, &m.ClosePrice)
    binary.Write(w, binary.BigEndian, &m.SettlementPrice)
    binary.Write(w, binary.BigEndian, &m.UpperLimitPrice)
    binary.Write(w, binary.BigEndian, &m.LowerLimitPrice)
    binary.Write(w, binary.BigEndian, &m.PreDelta)
    binary.Write(w, binary.BigEndian, &m.CurrDelta)
    writeLen(w, m.UpdateTime, 9)
    binary.Write(w, binary.BigEndian, &m.UpdateMillisec)
    writeLen(w, m.ActionDay, 9)
}

func (m* CThostFtdcMarketDataField) Unmarshal(r io.Reader) {
    m.TradingDay = readLen(r, 9)
    m.InstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
    m.ExchangeInstID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.LastPrice)
    binary.Read(r, binary.BigEndian, &m.PreSettlementPrice)
    binary.Read(r, binary.BigEndian, &m.PreClosePrice)
    binary.Read(r, binary.BigEndian, &m.PreOpenInterest)
    binary.Read(r, binary.BigEndian, &m.OpenPrice)
    binary.Read(r, binary.BigEndian, &m.HighestPrice)
    binary.Read(r, binary.BigEndian, &m.LowestPrice)
    binary.Read(r, binary.BigEndian, &m.Volume)
    binary.Read(r, binary.BigEndian, &m.Turnover)
    binary.Read(r, binary.BigEndian, &m.OpenInterest)
    binary.Read(r, binary.BigEndian, &m.ClosePrice)
    binary.Read(r, binary.BigEndian, &m.SettlementPrice)
    binary.Read(r, binary.BigEndian, &m.UpperLimitPrice)
    binary.Read(r, binary.BigEndian, &m.LowerLimitPrice)
    binary.Read(r, binary.BigEndian, &m.PreDelta)
    binary.Read(r, binary.BigEndian, &m.CurrDelta)
    m.UpdateTime = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.UpdateMillisec)
    m.ActionDay = readLen(r, 9)
}

type  CThostFtdcMarketDataBaseField struct {
    TradingDay string;
    PreSettlementPrice float64;
    PreClosePrice float64;
    PreOpenInterest float64;
    PreDelta float64;
}

func (m* CThostFtdcMarketDataBaseField) Marshal(w io.Writer) {
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.PreSettlementPrice)
    binary.Write(w, binary.BigEndian, &m.PreClosePrice)
    binary.Write(w, binary.BigEndian, &m.PreOpenInterest)
    binary.Write(w, binary.BigEndian, &m.PreDelta)
}

func (m* CThostFtdcMarketDataBaseField) Unmarshal(r io.Reader) {
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.PreSettlementPrice)
    binary.Read(r, binary.BigEndian, &m.PreClosePrice)
    binary.Read(r, binary.BigEndian, &m.PreOpenInterest)
    binary.Read(r, binary.BigEndian, &m.PreDelta)
}

type  CThostFtdcMarketDataStaticField struct {
    OpenPrice float64;
    HighestPrice float64;
    LowestPrice float64;
    ClosePrice float64;
    UpperLimitPrice float64;
    LowerLimitPrice float64;
    SettlementPrice float64;
    CurrDelta float64;
}

func (m* CThostFtdcMarketDataStaticField) Marshal(w io.Writer) {
    binary.Write(w, binary.BigEndian, &m.OpenPrice)
    binary.Write(w, binary.BigEndian, &m.HighestPrice)
    binary.Write(w, binary.BigEndian, &m.LowestPrice)
    binary.Write(w, binary.BigEndian, &m.ClosePrice)
    binary.Write(w, binary.BigEndian, &m.UpperLimitPrice)
    binary.Write(w, binary.BigEndian, &m.LowerLimitPrice)
    binary.Write(w, binary.BigEndian, &m.SettlementPrice)
    binary.Write(w, binary.BigEndian, &m.CurrDelta)
}

func (m* CThostFtdcMarketDataStaticField) Unmarshal(r io.Reader) {
    binary.Read(r, binary.BigEndian, &m.OpenPrice)
    binary.Read(r, binary.BigEndian, &m.HighestPrice)
    binary.Read(r, binary.BigEndian, &m.LowestPrice)
    binary.Read(r, binary.BigEndian, &m.ClosePrice)
    binary.Read(r, binary.BigEndian, &m.UpperLimitPrice)
    binary.Read(r, binary.BigEndian, &m.LowerLimitPrice)
    binary.Read(r, binary.BigEndian, &m.SettlementPrice)
    binary.Read(r, binary.BigEndian, &m.CurrDelta)
}

type  CThostFtdcMarketDataLastMatchField struct {
    LastPrice float64;
    Volume int32;
    Turnover float64;
    OpenInterest float64;
}

func (m* CThostFtdcMarketDataLastMatchField) Marshal(w io.Writer) {
    binary.Write(w, binary.BigEndian, &m.LastPrice)
    binary.Write(w, binary.BigEndian, &m.Volume)
    binary.Write(w, binary.BigEndian, &m.Turnover)
    binary.Write(w, binary.BigEndian, &m.OpenInterest)
}

func (m* CThostFtdcMarketDataLastMatchField) Unmarshal(r io.Reader) {
    binary.Read(r, binary.BigEndian, &m.LastPrice)
    binary.Read(r, binary.BigEndian, &m.Volume)
    binary.Read(r, binary.BigEndian, &m.Turnover)
    binary.Read(r, binary.BigEndian, &m.OpenInterest)
}

type  CThostFtdcMarketDataBestPriceField struct {
    BidPrice1 float64;
    BidVolume1 int32;
    AskPrice1 float64;
    AskVolume1 int32;
}

func (m* CThostFtdcMarketDataBestPriceField) Marshal(w io.Writer) {
    binary.Write(w, binary.BigEndian, &m.BidPrice1)
    binary.Write(w, binary.BigEndian, &m.BidVolume1)
    binary.Write(w, binary.BigEndian, &m.AskPrice1)
    binary.Write(w, binary.BigEndian, &m.AskVolume1)
}

func (m* CThostFtdcMarketDataBestPriceField) Unmarshal(r io.Reader) {
    binary.Read(r, binary.BigEndian, &m.BidPrice1)
    binary.Read(r, binary.BigEndian, &m.BidVolume1)
    binary.Read(r, binary.BigEndian, &m.AskPrice1)
    binary.Read(r, binary.BigEndian, &m.AskVolume1)
}

type  CThostFtdcMarketDataBid23Field struct {
    BidPrice2 float64;
    BidVolume2 int32;
    BidPrice3 float64;
    BidVolume3 int32;
}

func (m* CThostFtdcMarketDataBid23Field) Marshal(w io.Writer) {
    binary.Write(w, binary.BigEndian, &m.BidPrice2)
    binary.Write(w, binary.BigEndian, &m.BidVolume2)
    binary.Write(w, binary.BigEndian, &m.BidPrice3)
    binary.Write(w, binary.BigEndian, &m.BidVolume3)
}

func (m* CThostFtdcMarketDataBid23Field) Unmarshal(r io.Reader) {
    binary.Read(r, binary.BigEndian, &m.BidPrice2)
    binary.Read(r, binary.BigEndian, &m.BidVolume2)
    binary.Read(r, binary.BigEndian, &m.BidPrice3)
    binary.Read(r, binary.BigEndian, &m.BidVolume3)
}

type  CThostFtdcMarketDataAsk23Field struct {
    AskPrice2 float64;
    AskVolume2 int32;
    AskPrice3 float64;
    AskVolume3 int32;
}

func (m* CThostFtdcMarketDataAsk23Field) Marshal(w io.Writer) {
    binary.Write(w, binary.BigEndian, &m.AskPrice2)
    binary.Write(w, binary.BigEndian, &m.AskVolume2)
    binary.Write(w, binary.BigEndian, &m.AskPrice3)
    binary.Write(w, binary.BigEndian, &m.AskVolume3)
}

func (m* CThostFtdcMarketDataAsk23Field) Unmarshal(r io.Reader) {
    binary.Read(r, binary.BigEndian, &m.AskPrice2)
    binary.Read(r, binary.BigEndian, &m.AskVolume2)
    binary.Read(r, binary.BigEndian, &m.AskPrice3)
    binary.Read(r, binary.BigEndian, &m.AskVolume3)
}

type  CThostFtdcMarketDataBid45Field struct {
    BidPrice4 float64;
    BidVolume4 int32;
    BidPrice5 float64;
    BidVolume5 int32;
}

func (m* CThostFtdcMarketDataBid45Field) Marshal(w io.Writer) {
    binary.Write(w, binary.BigEndian, &m.BidPrice4)
    binary.Write(w, binary.BigEndian, &m.BidVolume4)
    binary.Write(w, binary.BigEndian, &m.BidPrice5)
    binary.Write(w, binary.BigEndian, &m.BidVolume5)
}

func (m* CThostFtdcMarketDataBid45Field) Unmarshal(r io.Reader) {
    binary.Read(r, binary.BigEndian, &m.BidPrice4)
    binary.Read(r, binary.BigEndian, &m.BidVolume4)
    binary.Read(r, binary.BigEndian, &m.BidPrice5)
    binary.Read(r, binary.BigEndian, &m.BidVolume5)
}

type  CThostFtdcMarketDataAsk45Field struct {
    AskPrice4 float64;
    AskVolume4 int32;
    AskPrice5 float64;
    AskVolume5 int32;
}

func (m* CThostFtdcMarketDataAsk45Field) Marshal(w io.Writer) {
    binary.Write(w, binary.BigEndian, &m.AskPrice4)
    binary.Write(w, binary.BigEndian, &m.AskVolume4)
    binary.Write(w, binary.BigEndian, &m.AskPrice5)
    binary.Write(w, binary.BigEndian, &m.AskVolume5)
}

func (m* CThostFtdcMarketDataAsk45Field) Unmarshal(r io.Reader) {
    binary.Read(r, binary.BigEndian, &m.AskPrice4)
    binary.Read(r, binary.BigEndian, &m.AskVolume4)
    binary.Read(r, binary.BigEndian, &m.AskPrice5)
    binary.Read(r, binary.BigEndian, &m.AskVolume5)
}

type  CThostFtdcMarketDataUpdateTimeField struct {
    InstrumentID string;
    UpdateTime string;
    UpdateMillisec int32;
    ActionDay string;
    ExchangeID string;
}

func (m* CThostFtdcMarketDataUpdateTimeField) Marshal(w io.Writer) {
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.UpdateTime, 9)
    binary.Write(w, binary.BigEndian, &m.UpdateMillisec)
    writeLen(w, m.ActionDay, 9)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcMarketDataUpdateTimeField) Unmarshal(r io.Reader) {
    m.InstrumentID = readLen(r, 31)
    m.UpdateTime = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.UpdateMillisec)
    m.ActionDay = readLen(r, 9)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcMarketDataExchangeField struct {
    ExchangeID string;
}

func (m* CThostFtdcMarketDataExchangeField) Marshal(w io.Writer) {
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcMarketDataExchangeField) Unmarshal(r io.Reader) {
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcSpecificInstrumentField struct {
    InstrumentID string;
}

func (m* CThostFtdcSpecificInstrumentField) Marshal(w io.Writer) {
    writeLen(w, m.InstrumentID, 31)
}

func (m* CThostFtdcSpecificInstrumentField) Unmarshal(r io.Reader) {
    m.InstrumentID = readLen(r, 31)
}

type  CThostFtdcInstrumentStatusField struct {
    ExchangeID string;
    ExchangeInstID string;
    SettlementGroupID string;
    InstrumentID string;
    InstrumentStatus uint8;
    TradingSegmentSN int32;
    EnterTime string;
    EnterReason uint8;
}

func (m* CThostFtdcInstrumentStatusField) Marshal(w io.Writer) {
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ExchangeInstID, 31)
    writeLen(w, m.SettlementGroupID, 9)
    writeLen(w, m.InstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.InstrumentStatus)
    binary.Write(w, binary.BigEndian, &m.TradingSegmentSN)
    writeLen(w, m.EnterTime, 9)
    binary.Write(w, binary.BigEndian, &m.EnterReason)
}

func (m* CThostFtdcInstrumentStatusField) Unmarshal(r io.Reader) {
    m.ExchangeID = readLen(r, 9)
    m.ExchangeInstID = readLen(r, 31)
    m.SettlementGroupID = readLen(r, 9)
    m.InstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.InstrumentStatus)
    binary.Read(r, binary.BigEndian, &m.TradingSegmentSN)
    m.EnterTime = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.EnterReason)
}

type  CThostFtdcQryInstrumentStatusField struct {
    ExchangeID string;
    ExchangeInstID string;
}

func (m* CThostFtdcQryInstrumentStatusField) Marshal(w io.Writer) {
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ExchangeInstID, 31)
}

func (m* CThostFtdcQryInstrumentStatusField) Unmarshal(r io.Reader) {
    m.ExchangeID = readLen(r, 9)
    m.ExchangeInstID = readLen(r, 31)
}

type  CThostFtdcInvestorAccountField struct {
    BrokerID string;
    InvestorID string;
    AccountID string;
    CurrencyID string;
    BizType uint8;
}

func (m* CThostFtdcInvestorAccountField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.CurrencyID, 4)
    binary.Write(w, binary.BigEndian, &m.BizType)
}

func (m* CThostFtdcInvestorAccountField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.AccountID = readLen(r, 13)
    m.CurrencyID = readLen(r, 4)
    binary.Read(r, binary.BigEndian, &m.BizType)
}

type  CThostFtdcPositionProfitAlgorithmField struct {
    BrokerID string;
    AccountID string;
    Algorithm uint8;
    Memo string;
    CurrencyID string;
}

func (m* CThostFtdcPositionProfitAlgorithmField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.AccountID, 13)
    binary.Write(w, binary.BigEndian, &m.Algorithm)
    writeLen(w, m.Memo, 161)
    writeLen(w, m.CurrencyID, 4)
}

func (m* CThostFtdcPositionProfitAlgorithmField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.AccountID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.Algorithm)
    m.Memo = readLen(r, 161)
    m.CurrencyID = readLen(r, 4)
}

type  CThostFtdcDiscountField struct {
    BrokerID string;
    InvestorRange uint8;
    InvestorID string;
    Discount float64;
}

func (m* CThostFtdcDiscountField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    binary.Write(w, binary.BigEndian, &m.InvestorRange)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.Discount)
}

func (m* CThostFtdcDiscountField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    binary.Read(r, binary.BigEndian, &m.InvestorRange)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.Discount)
}

type  CThostFtdcQryTransferBankField struct {
    BankID string;
    BankBrchID string;
}

func (m* CThostFtdcQryTransferBankField) Marshal(w io.Writer) {
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBrchID, 5)
}

func (m* CThostFtdcQryTransferBankField) Unmarshal(r io.Reader) {
    m.BankID = readLen(r, 4)
    m.BankBrchID = readLen(r, 5)
}

type  CThostFtdcTransferBankField struct {
    BankID string;
    BankBrchID string;
    BankName string;
    IsActive int32;
}

func (m* CThostFtdcTransferBankField) Marshal(w io.Writer) {
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBrchID, 5)
    writeLen(w, m.BankName, 101)
    binary.Write(w, binary.BigEndian, &m.IsActive)
}

func (m* CThostFtdcTransferBankField) Unmarshal(r io.Reader) {
    m.BankID = readLen(r, 4)
    m.BankBrchID = readLen(r, 5)
    m.BankName = readLen(r, 101)
    binary.Read(r, binary.BigEndian, &m.IsActive)
}

type  CThostFtdcQryInvestorPositionDetailField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    ExchangeID string;
    InvestUnitID string;
}

func (m* CThostFtdcQryInvestorPositionDetailField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.InvestUnitID, 17)
}

func (m* CThostFtdcQryInvestorPositionDetailField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
}

type  CThostFtdcInvestorPositionDetailField struct {
    InstrumentID string;
    BrokerID string;
    InvestorID string;
    HedgeFlag uint8;
    Direction uint8;
    OpenDate string;
    TradeID string;
    Volume int32;
    OpenPrice float64;
    TradingDay string;
    SettlementID int32;
    TradeType uint8;
    CombInstrumentID string;
    ExchangeID string;
    CloseProfitByDate float64;
    CloseProfitByTrade float64;
    PositionProfitByDate float64;
    PositionProfitByTrade float64;
    Margin float64;
    ExchMargin float64;
    MarginRateByMoney float64;
    MarginRateByVolume float64;
    LastSettlementPrice float64;
    SettlementPrice float64;
    CloseVolume int32;
    CloseAmount float64;
    InvestUnitID string;
}

func (m* CThostFtdcInvestorPositionDetailField) Marshal(w io.Writer) {
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.HedgeFlag)
    binary.Write(w, binary.BigEndian, &m.Direction)
    writeLen(w, m.OpenDate, 9)
    writeLen(w, m.TradeID, 21)
    binary.Write(w, binary.BigEndian, &m.Volume)
    binary.Write(w, binary.BigEndian, &m.OpenPrice)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.SettlementID)
    binary.Write(w, binary.BigEndian, &m.TradeType)
    writeLen(w, m.CombInstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
    binary.Write(w, binary.BigEndian, &m.CloseProfitByDate)
    binary.Write(w, binary.BigEndian, &m.CloseProfitByTrade)
    binary.Write(w, binary.BigEndian, &m.PositionProfitByDate)
    binary.Write(w, binary.BigEndian, &m.PositionProfitByTrade)
    binary.Write(w, binary.BigEndian, &m.Margin)
    binary.Write(w, binary.BigEndian, &m.ExchMargin)
    binary.Write(w, binary.BigEndian, &m.MarginRateByMoney)
    binary.Write(w, binary.BigEndian, &m.MarginRateByVolume)
    binary.Write(w, binary.BigEndian, &m.LastSettlementPrice)
    binary.Write(w, binary.BigEndian, &m.SettlementPrice)
    binary.Write(w, binary.BigEndian, &m.CloseVolume)
    binary.Write(w, binary.BigEndian, &m.CloseAmount)
    writeLen(w, m.InvestUnitID, 17)
}

func (m* CThostFtdcInvestorPositionDetailField) Unmarshal(r io.Reader) {
    m.InstrumentID = readLen(r, 31)
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.HedgeFlag)
    binary.Read(r, binary.BigEndian, &m.Direction)
    m.OpenDate = readLen(r, 9)
    m.TradeID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.Volume)
    binary.Read(r, binary.BigEndian, &m.OpenPrice)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.SettlementID)
    binary.Read(r, binary.BigEndian, &m.TradeType)
    m.CombInstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.CloseProfitByDate)
    binary.Read(r, binary.BigEndian, &m.CloseProfitByTrade)
    binary.Read(r, binary.BigEndian, &m.PositionProfitByDate)
    binary.Read(r, binary.BigEndian, &m.PositionProfitByTrade)
    binary.Read(r, binary.BigEndian, &m.Margin)
    binary.Read(r, binary.BigEndian, &m.ExchMargin)
    binary.Read(r, binary.BigEndian, &m.MarginRateByMoney)
    binary.Read(r, binary.BigEndian, &m.MarginRateByVolume)
    binary.Read(r, binary.BigEndian, &m.LastSettlementPrice)
    binary.Read(r, binary.BigEndian, &m.SettlementPrice)
    binary.Read(r, binary.BigEndian, &m.CloseVolume)
    binary.Read(r, binary.BigEndian, &m.CloseAmount)
    m.InvestUnitID = readLen(r, 17)
}

type  CThostFtdcTradingAccountPasswordField struct {
    BrokerID string;
    AccountID string;
    Password string;
    CurrencyID string;
}

func (m* CThostFtdcTradingAccountPasswordField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.Password, 41)
    writeLen(w, m.CurrencyID, 4)
}

func (m* CThostFtdcTradingAccountPasswordField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.AccountID = readLen(r, 13)
    m.Password = readLen(r, 41)
    m.CurrencyID = readLen(r, 4)
}

type  CThostFtdcMDTraderOfferField struct {
    ExchangeID string;
    TraderID string;
    ParticipantID string;
    Password string;
    InstallID int32;
    OrderLocalID string;
    TraderConnectStatus uint8;
    ConnectRequestDate string;
    ConnectRequestTime string;
    LastReportDate string;
    LastReportTime string;
    ConnectDate string;
    ConnectTime string;
    StartDate string;
    StartTime string;
    TradingDay string;
    BrokerID string;
    MaxTradeID string;
    MaxOrderMessageReference string;
    BizType uint8;
}

func (m* CThostFtdcMDTraderOfferField) Marshal(w io.Writer) {
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.TraderID, 21)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.Password, 41)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    writeLen(w, m.OrderLocalID, 13)
    binary.Write(w, binary.BigEndian, &m.TraderConnectStatus)
    writeLen(w, m.ConnectRequestDate, 9)
    writeLen(w, m.ConnectRequestTime, 9)
    writeLen(w, m.LastReportDate, 9)
    writeLen(w, m.LastReportTime, 9)
    writeLen(w, m.ConnectDate, 9)
    writeLen(w, m.ConnectTime, 9)
    writeLen(w, m.StartDate, 9)
    writeLen(w, m.StartTime, 9)
    writeLen(w, m.TradingDay, 9)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.MaxTradeID, 21)
    writeLen(w, m.MaxOrderMessageReference, 7)
    binary.Write(w, binary.BigEndian, &m.BizType)
}

func (m* CThostFtdcMDTraderOfferField) Unmarshal(r io.Reader) {
    m.ExchangeID = readLen(r, 9)
    m.TraderID = readLen(r, 21)
    m.ParticipantID = readLen(r, 11)
    m.Password = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    m.OrderLocalID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.TraderConnectStatus)
    m.ConnectRequestDate = readLen(r, 9)
    m.ConnectRequestTime = readLen(r, 9)
    m.LastReportDate = readLen(r, 9)
    m.LastReportTime = readLen(r, 9)
    m.ConnectDate = readLen(r, 9)
    m.ConnectTime = readLen(r, 9)
    m.StartDate = readLen(r, 9)
    m.StartTime = readLen(r, 9)
    m.TradingDay = readLen(r, 9)
    m.BrokerID = readLen(r, 11)
    m.MaxTradeID = readLen(r, 21)
    m.MaxOrderMessageReference = readLen(r, 7)
    binary.Read(r, binary.BigEndian, &m.BizType)
}

type  CThostFtdcQryMDTraderOfferField struct {
    ExchangeID string;
    ParticipantID string;
    TraderID string;
}

func (m* CThostFtdcQryMDTraderOfferField) Marshal(w io.Writer) {
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.TraderID, 21)
}

func (m* CThostFtdcQryMDTraderOfferField) Unmarshal(r io.Reader) {
    m.ExchangeID = readLen(r, 9)
    m.ParticipantID = readLen(r, 11)
    m.TraderID = readLen(r, 21)
}

type  CThostFtdcQryNoticeField struct {
    BrokerID string;
}

func (m* CThostFtdcQryNoticeField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
}

func (m* CThostFtdcQryNoticeField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
}

type  CThostFtdcNoticeField struct {
    BrokerID string;
    Content string;
    SequenceLabel string;
}

func (m* CThostFtdcNoticeField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.Content, 501)
    writeLen(w, m.SequenceLabel, 2)
}

func (m* CThostFtdcNoticeField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.Content = readLen(r, 501)
    m.SequenceLabel = readLen(r, 2)
}

type  CThostFtdcUserRightField struct {
    BrokerID string;
    UserID string;
    UserRightType uint8;
    IsForbidden int32;
}

func (m* CThostFtdcUserRightField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.UserRightType)
    binary.Write(w, binary.BigEndian, &m.IsForbidden)
}

func (m* CThostFtdcUserRightField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.UserRightType)
    binary.Read(r, binary.BigEndian, &m.IsForbidden)
}

type  CThostFtdcQrySettlementInfoConfirmField struct {
    BrokerID string;
    InvestorID string;
    AccountID string;
    CurrencyID string;
}

func (m* CThostFtdcQrySettlementInfoConfirmField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.CurrencyID, 4)
}

func (m* CThostFtdcQrySettlementInfoConfirmField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.AccountID = readLen(r, 13)
    m.CurrencyID = readLen(r, 4)
}

type  CThostFtdcLoadSettlementInfoField struct {
    BrokerID string;
}

func (m* CThostFtdcLoadSettlementInfoField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
}

func (m* CThostFtdcLoadSettlementInfoField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
}

type  CThostFtdcBrokerWithdrawAlgorithmField struct {
    BrokerID string;
    WithdrawAlgorithm uint8;
    UsingRatio float64;
    IncludeCloseProfit uint8;
    AllWithoutTrade uint8;
    AvailIncludeCloseProfit uint8;
    IsBrokerUserEvent int32;
    CurrencyID string;
    FundMortgageRatio float64;
    BalanceAlgorithm uint8;
}

func (m* CThostFtdcBrokerWithdrawAlgorithmField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    binary.Write(w, binary.BigEndian, &m.WithdrawAlgorithm)
    binary.Write(w, binary.BigEndian, &m.UsingRatio)
    binary.Write(w, binary.BigEndian, &m.IncludeCloseProfit)
    binary.Write(w, binary.BigEndian, &m.AllWithoutTrade)
    binary.Write(w, binary.BigEndian, &m.AvailIncludeCloseProfit)
    binary.Write(w, binary.BigEndian, &m.IsBrokerUserEvent)
    writeLen(w, m.CurrencyID, 4)
    binary.Write(w, binary.BigEndian, &m.FundMortgageRatio)
    binary.Write(w, binary.BigEndian, &m.BalanceAlgorithm)
}

func (m* CThostFtdcBrokerWithdrawAlgorithmField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    binary.Read(r, binary.BigEndian, &m.WithdrawAlgorithm)
    binary.Read(r, binary.BigEndian, &m.UsingRatio)
    binary.Read(r, binary.BigEndian, &m.IncludeCloseProfit)
    binary.Read(r, binary.BigEndian, &m.AllWithoutTrade)
    binary.Read(r, binary.BigEndian, &m.AvailIncludeCloseProfit)
    binary.Read(r, binary.BigEndian, &m.IsBrokerUserEvent)
    m.CurrencyID = readLen(r, 4)
    binary.Read(r, binary.BigEndian, &m.FundMortgageRatio)
    binary.Read(r, binary.BigEndian, &m.BalanceAlgorithm)
}

type  CThostFtdcTradingAccountPasswordUpdateV1Field struct {
    BrokerID string;
    InvestorID string;
    OldPassword string;
    NewPassword string;
}

func (m* CThostFtdcTradingAccountPasswordUpdateV1Field) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.OldPassword, 41)
    writeLen(w, m.NewPassword, 41)
}

func (m* CThostFtdcTradingAccountPasswordUpdateV1Field) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.OldPassword = readLen(r, 41)
    m.NewPassword = readLen(r, 41)
}

type  CThostFtdcTradingAccountPasswordUpdateField struct {
    BrokerID string;
    AccountID string;
    OldPassword string;
    NewPassword string;
    CurrencyID string;
}

func (m* CThostFtdcTradingAccountPasswordUpdateField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.OldPassword, 41)
    writeLen(w, m.NewPassword, 41)
    writeLen(w, m.CurrencyID, 4)
}

func (m* CThostFtdcTradingAccountPasswordUpdateField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.AccountID = readLen(r, 13)
    m.OldPassword = readLen(r, 41)
    m.NewPassword = readLen(r, 41)
    m.CurrencyID = readLen(r, 4)
}

type  CThostFtdcQryCombinationLegField struct {
    CombInstrumentID string;
    LegID int32;
    LegInstrumentID string;
}

func (m* CThostFtdcQryCombinationLegField) Marshal(w io.Writer) {
    writeLen(w, m.CombInstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.LegID)
    writeLen(w, m.LegInstrumentID, 31)
}

func (m* CThostFtdcQryCombinationLegField) Unmarshal(r io.Reader) {
    m.CombInstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.LegID)
    m.LegInstrumentID = readLen(r, 31)
}

type  CThostFtdcQrySyncStatusField struct {
    TradingDay string;
}

func (m* CThostFtdcQrySyncStatusField) Marshal(w io.Writer) {
    writeLen(w, m.TradingDay, 9)
}

func (m* CThostFtdcQrySyncStatusField) Unmarshal(r io.Reader) {
    m.TradingDay = readLen(r, 9)
}

type  CThostFtdcCombinationLegField struct {
    CombInstrumentID string;
    LegID int32;
    LegInstrumentID string;
    Direction uint8;
    LegMultiple int32;
    ImplyLevel int32;
}

func (m* CThostFtdcCombinationLegField) Marshal(w io.Writer) {
    writeLen(w, m.CombInstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.LegID)
    writeLen(w, m.LegInstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.Direction)
    binary.Write(w, binary.BigEndian, &m.LegMultiple)
    binary.Write(w, binary.BigEndian, &m.ImplyLevel)
}

func (m* CThostFtdcCombinationLegField) Unmarshal(r io.Reader) {
    m.CombInstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.LegID)
    m.LegInstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.Direction)
    binary.Read(r, binary.BigEndian, &m.LegMultiple)
    binary.Read(r, binary.BigEndian, &m.ImplyLevel)
}

type  CThostFtdcSyncStatusField struct {
    TradingDay string;
    DataSyncStatus uint8;
}

func (m* CThostFtdcSyncStatusField) Marshal(w io.Writer) {
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.DataSyncStatus)
}

func (m* CThostFtdcSyncStatusField) Unmarshal(r io.Reader) {
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.DataSyncStatus)
}

type  CThostFtdcQryLinkManField struct {
    BrokerID string;
    InvestorID string;
}

func (m* CThostFtdcQryLinkManField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
}

func (m* CThostFtdcQryLinkManField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
}

type  CThostFtdcLinkManField struct {
    BrokerID string;
    InvestorID string;
    PersonType uint8;
    IdentifiedCardType uint8;
    IdentifiedCardNo string;
    PersonName string;
    Telephone string;
    Address string;
    ZipCode string;
    Priority int32;
    UOAZipCode string;
    PersonFullName string;
}

func (m* CThostFtdcLinkManField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.PersonType)
    binary.Write(w, binary.BigEndian, &m.IdentifiedCardType)
    writeLen(w, m.IdentifiedCardNo, 51)
    writeLen(w, m.PersonName, 81)
    writeLen(w, m.Telephone, 41)
    writeLen(w, m.Address, 101)
    writeLen(w, m.ZipCode, 7)
    binary.Write(w, binary.BigEndian, &m.Priority)
    writeLen(w, m.UOAZipCode, 11)
    writeLen(w, m.PersonFullName, 101)
}

func (m* CThostFtdcLinkManField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.PersonType)
    binary.Read(r, binary.BigEndian, &m.IdentifiedCardType)
    m.IdentifiedCardNo = readLen(r, 51)
    m.PersonName = readLen(r, 81)
    m.Telephone = readLen(r, 41)
    m.Address = readLen(r, 101)
    m.ZipCode = readLen(r, 7)
    binary.Read(r, binary.BigEndian, &m.Priority)
    m.UOAZipCode = readLen(r, 11)
    m.PersonFullName = readLen(r, 101)
}

type  CThostFtdcQryBrokerUserEventField struct {
    BrokerID string;
    UserID string;
    UserEventType uint8;
}

func (m* CThostFtdcQryBrokerUserEventField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.UserEventType)
}

func (m* CThostFtdcQryBrokerUserEventField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.UserEventType)
}

type  CThostFtdcBrokerUserEventField struct {
    BrokerID string;
    UserID string;
    UserEventType uint8;
    EventSequenceNo int32;
    EventDate string;
    EventTime string;
    UserEventInfo string;
    InvestorID string;
    InstrumentID string;
    ExchangeID string;
}

func (m* CThostFtdcBrokerUserEventField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.UserEventType)
    binary.Write(w, binary.BigEndian, &m.EventSequenceNo)
    writeLen(w, m.EventDate, 9)
    writeLen(w, m.EventTime, 9)
    writeLen(w, m.UserEventInfo, 1025)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcBrokerUserEventField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.UserEventType)
    binary.Read(r, binary.BigEndian, &m.EventSequenceNo)
    m.EventDate = readLen(r, 9)
    m.EventTime = readLen(r, 9)
    m.UserEventInfo = readLen(r, 1025)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcQryContractBankField struct {
    BrokerID string;
    BankID string;
    BankBrchID string;
}

func (m* CThostFtdcQryContractBankField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBrchID, 5)
}

func (m* CThostFtdcQryContractBankField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.BankID = readLen(r, 4)
    m.BankBrchID = readLen(r, 5)
}

type  CThostFtdcContractBankField struct {
    BrokerID string;
    BankID string;
    BankBrchID string;
    BankName string;
}

func (m* CThostFtdcContractBankField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBrchID, 5)
    writeLen(w, m.BankName, 101)
}

func (m* CThostFtdcContractBankField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.BankID = readLen(r, 4)
    m.BankBrchID = readLen(r, 5)
    m.BankName = readLen(r, 101)
}

type  CThostFtdcInvestorPositionCombineDetailField struct {
    TradingDay string;
    OpenDate string;
    ExchangeID string;
    SettlementID int32;
    BrokerID string;
    InvestorID string;
    ComTradeID string;
    TradeID string;
    InstrumentID string;
    HedgeFlag uint8;
    Direction uint8;
    TotalAmt int32;
    Margin float64;
    ExchMargin float64;
    MarginRateByMoney float64;
    MarginRateByVolume float64;
    LegID int32;
    LegMultiple int32;
    CombInstrumentID string;
    TradeGroupID int32;
    InvestUnitID string;
}

func (m* CThostFtdcInvestorPositionCombineDetailField) Marshal(w io.Writer) {
    writeLen(w, m.TradingDay, 9)
    writeLen(w, m.OpenDate, 9)
    writeLen(w, m.ExchangeID, 9)
    binary.Write(w, binary.BigEndian, &m.SettlementID)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.ComTradeID, 21)
    writeLen(w, m.TradeID, 21)
    writeLen(w, m.InstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.HedgeFlag)
    binary.Write(w, binary.BigEndian, &m.Direction)
    binary.Write(w, binary.BigEndian, &m.TotalAmt)
    binary.Write(w, binary.BigEndian, &m.Margin)
    binary.Write(w, binary.BigEndian, &m.ExchMargin)
    binary.Write(w, binary.BigEndian, &m.MarginRateByMoney)
    binary.Write(w, binary.BigEndian, &m.MarginRateByVolume)
    binary.Write(w, binary.BigEndian, &m.LegID)
    binary.Write(w, binary.BigEndian, &m.LegMultiple)
    writeLen(w, m.CombInstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.TradeGroupID)
    writeLen(w, m.InvestUnitID, 17)
}

func (m* CThostFtdcInvestorPositionCombineDetailField) Unmarshal(r io.Reader) {
    m.TradingDay = readLen(r, 9)
    m.OpenDate = readLen(r, 9)
    m.ExchangeID = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.SettlementID)
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.ComTradeID = readLen(r, 21)
    m.TradeID = readLen(r, 21)
    m.InstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.HedgeFlag)
    binary.Read(r, binary.BigEndian, &m.Direction)
    binary.Read(r, binary.BigEndian, &m.TotalAmt)
    binary.Read(r, binary.BigEndian, &m.Margin)
    binary.Read(r, binary.BigEndian, &m.ExchMargin)
    binary.Read(r, binary.BigEndian, &m.MarginRateByMoney)
    binary.Read(r, binary.BigEndian, &m.MarginRateByVolume)
    binary.Read(r, binary.BigEndian, &m.LegID)
    binary.Read(r, binary.BigEndian, &m.LegMultiple)
    m.CombInstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.TradeGroupID)
    m.InvestUnitID = readLen(r, 17)
}

type  CThostFtdcParkedOrderField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    OrderRef string;
    UserID string;
    OrderPriceType uint8;
    Direction uint8;
    CombOffsetFlag string;
    CombHedgeFlag string;
    LimitPrice float64;
    VolumeTotalOriginal int32;
    TimeCondition uint8;
    GTDDate string;
    VolumeCondition uint8;
    MinVolume int32;
    ContingentCondition uint8;
    StopPrice float64;
    ForceCloseReason uint8;
    IsAutoSuspend int32;
    BusinessUnit string;
    RequestID int32;
    UserForceClose int32;
    ExchangeID string;
    ParkedOrderID string;
    UserType uint8;
    Status uint8;
    ErrorID int32;
    ErrorMsg string;
    IsSwapOrder int32;
    AccountID string;
    CurrencyID string;
    ClientID string;
    InvestUnitID string;
    IPAddress string;
    MacAddress string;
}

func (m* CThostFtdcParkedOrderField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.OrderRef, 13)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.OrderPriceType)
    binary.Write(w, binary.BigEndian, &m.Direction)
    writeLen(w, m.CombOffsetFlag, 5)
    writeLen(w, m.CombHedgeFlag, 5)
    binary.Write(w, binary.BigEndian, &m.LimitPrice)
    binary.Write(w, binary.BigEndian, &m.VolumeTotalOriginal)
    binary.Write(w, binary.BigEndian, &m.TimeCondition)
    writeLen(w, m.GTDDate, 9)
    binary.Write(w, binary.BigEndian, &m.VolumeCondition)
    binary.Write(w, binary.BigEndian, &m.MinVolume)
    binary.Write(w, binary.BigEndian, &m.ContingentCondition)
    binary.Write(w, binary.BigEndian, &m.StopPrice)
    binary.Write(w, binary.BigEndian, &m.ForceCloseReason)
    binary.Write(w, binary.BigEndian, &m.IsAutoSuspend)
    writeLen(w, m.BusinessUnit, 21)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    binary.Write(w, binary.BigEndian, &m.UserForceClose)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ParkedOrderID, 13)
    binary.Write(w, binary.BigEndian, &m.UserType)
    binary.Write(w, binary.BigEndian, &m.Status)
    binary.Write(w, binary.BigEndian, &m.ErrorID)
    writeLen(w, m.ErrorMsg, 81)
    binary.Write(w, binary.BigEndian, &m.IsSwapOrder)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.CurrencyID, 4)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.InvestUnitID, 17)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
}

func (m* CThostFtdcParkedOrderField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.OrderRef = readLen(r, 13)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.OrderPriceType)
    binary.Read(r, binary.BigEndian, &m.Direction)
    m.CombOffsetFlag = readLen(r, 5)
    m.CombHedgeFlag = readLen(r, 5)
    binary.Read(r, binary.BigEndian, &m.LimitPrice)
    binary.Read(r, binary.BigEndian, &m.VolumeTotalOriginal)
    binary.Read(r, binary.BigEndian, &m.TimeCondition)
    m.GTDDate = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.VolumeCondition)
    binary.Read(r, binary.BigEndian, &m.MinVolume)
    binary.Read(r, binary.BigEndian, &m.ContingentCondition)
    binary.Read(r, binary.BigEndian, &m.StopPrice)
    binary.Read(r, binary.BigEndian, &m.ForceCloseReason)
    binary.Read(r, binary.BigEndian, &m.IsAutoSuspend)
    m.BusinessUnit = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    binary.Read(r, binary.BigEndian, &m.UserForceClose)
    m.ExchangeID = readLen(r, 9)
    m.ParkedOrderID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.UserType)
    binary.Read(r, binary.BigEndian, &m.Status)
    binary.Read(r, binary.BigEndian, &m.ErrorID)
    m.ErrorMsg = readLen(r, 81)
    binary.Read(r, binary.BigEndian, &m.IsSwapOrder)
    m.AccountID = readLen(r, 13)
    m.CurrencyID = readLen(r, 4)
    m.ClientID = readLen(r, 11)
    m.InvestUnitID = readLen(r, 17)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
}

type  CThostFtdcParkedOrderActionField struct {
    BrokerID string;
    InvestorID string;
    OrderActionRef int32;
    OrderRef string;
    RequestID int32;
    FrontID int32;
    SessionID int32;
    ExchangeID string;
    OrderSysID string;
    ActionFlag uint8;
    LimitPrice float64;
    VolumeChange int32;
    UserID string;
    InstrumentID string;
    ParkedOrderActionID string;
    UserType uint8;
    Status uint8;
    ErrorID int32;
    ErrorMsg string;
    InvestUnitID string;
    IPAddress string;
    MacAddress string;
}

func (m* CThostFtdcParkedOrderActionField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.OrderActionRef)
    writeLen(w, m.OrderRef, 13)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    binary.Write(w, binary.BigEndian, &m.FrontID)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.OrderSysID, 21)
    binary.Write(w, binary.BigEndian, &m.ActionFlag)
    binary.Write(w, binary.BigEndian, &m.LimitPrice)
    binary.Write(w, binary.BigEndian, &m.VolumeChange)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ParkedOrderActionID, 13)
    binary.Write(w, binary.BigEndian, &m.UserType)
    binary.Write(w, binary.BigEndian, &m.Status)
    binary.Write(w, binary.BigEndian, &m.ErrorID)
    writeLen(w, m.ErrorMsg, 81)
    writeLen(w, m.InvestUnitID, 17)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
}

func (m* CThostFtdcParkedOrderActionField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.OrderActionRef)
    m.OrderRef = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    binary.Read(r, binary.BigEndian, &m.FrontID)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.ExchangeID = readLen(r, 9)
    m.OrderSysID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.ActionFlag)
    binary.Read(r, binary.BigEndian, &m.LimitPrice)
    binary.Read(r, binary.BigEndian, &m.VolumeChange)
    m.UserID = readLen(r, 16)
    m.InstrumentID = readLen(r, 31)
    m.ParkedOrderActionID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.UserType)
    binary.Read(r, binary.BigEndian, &m.Status)
    binary.Read(r, binary.BigEndian, &m.ErrorID)
    m.ErrorMsg = readLen(r, 81)
    m.InvestUnitID = readLen(r, 17)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
}

type  CThostFtdcQryParkedOrderField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    ExchangeID string;
    InvestUnitID string;
}

func (m* CThostFtdcQryParkedOrderField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.InvestUnitID, 17)
}

func (m* CThostFtdcQryParkedOrderField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
}

type  CThostFtdcQryParkedOrderActionField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    ExchangeID string;
    InvestUnitID string;
}

func (m* CThostFtdcQryParkedOrderActionField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.InvestUnitID, 17)
}

func (m* CThostFtdcQryParkedOrderActionField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
}

type  CThostFtdcRemoveParkedOrderField struct {
    BrokerID string;
    InvestorID string;
    ParkedOrderID string;
    InvestUnitID string;
}

func (m* CThostFtdcRemoveParkedOrderField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.ParkedOrderID, 13)
    writeLen(w, m.InvestUnitID, 17)
}

func (m* CThostFtdcRemoveParkedOrderField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.ParkedOrderID = readLen(r, 13)
    m.InvestUnitID = readLen(r, 17)
}

type  CThostFtdcRemoveParkedOrderActionField struct {
    BrokerID string;
    InvestorID string;
    ParkedOrderActionID string;
    InvestUnitID string;
}

func (m* CThostFtdcRemoveParkedOrderActionField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.ParkedOrderActionID, 13)
    writeLen(w, m.InvestUnitID, 17)
}

func (m* CThostFtdcRemoveParkedOrderActionField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.ParkedOrderActionID = readLen(r, 13)
    m.InvestUnitID = readLen(r, 17)
}

type  CThostFtdcInvestorWithdrawAlgorithmField struct {
    BrokerID string;
    InvestorRange uint8;
    InvestorID string;
    UsingRatio float64;
    CurrencyID string;
    FundMortgageRatio float64;
}

func (m* CThostFtdcInvestorWithdrawAlgorithmField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    binary.Write(w, binary.BigEndian, &m.InvestorRange)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.UsingRatio)
    writeLen(w, m.CurrencyID, 4)
    binary.Write(w, binary.BigEndian, &m.FundMortgageRatio)
}

func (m* CThostFtdcInvestorWithdrawAlgorithmField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    binary.Read(r, binary.BigEndian, &m.InvestorRange)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.UsingRatio)
    m.CurrencyID = readLen(r, 4)
    binary.Read(r, binary.BigEndian, &m.FundMortgageRatio)
}

type  CThostFtdcQryInvestorPositionCombineDetailField struct {
    BrokerID string;
    InvestorID string;
    CombInstrumentID string;
    ExchangeID string;
    InvestUnitID string;
}

func (m* CThostFtdcQryInvestorPositionCombineDetailField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.CombInstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.InvestUnitID, 17)
}

func (m* CThostFtdcQryInvestorPositionCombineDetailField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.CombInstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
}

type  CThostFtdcMarketDataAveragePriceField struct {
    AveragePrice float64;
}

func (m* CThostFtdcMarketDataAveragePriceField) Marshal(w io.Writer) {
    binary.Write(w, binary.BigEndian, &m.AveragePrice)
}

func (m* CThostFtdcMarketDataAveragePriceField) Unmarshal(r io.Reader) {
    binary.Read(r, binary.BigEndian, &m.AveragePrice)
}

type  CThostFtdcVerifyInvestorPasswordField struct {
    BrokerID string;
    InvestorID string;
    Password string;
}

func (m* CThostFtdcVerifyInvestorPasswordField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.Password, 41)
}

func (m* CThostFtdcVerifyInvestorPasswordField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.Password = readLen(r, 41)
}

type  CThostFtdcUserIPField struct {
    BrokerID string;
    UserID string;
    IPAddress string;
    IPMask string;
    MacAddress string;
}

func (m* CThostFtdcUserIPField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.IPMask, 16)
    writeLen(w, m.MacAddress, 21)
}

func (m* CThostFtdcUserIPField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.UserID = readLen(r, 16)
    m.IPAddress = readLen(r, 16)
    m.IPMask = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
}

type  CThostFtdcTradingNoticeInfoField struct {
    BrokerID string;
    InvestorID string;
    SendTime string;
    FieldContent string;
    SequenceSeries int16;
    SequenceNo int32;
    InvestUnitID string;
}

func (m* CThostFtdcTradingNoticeInfoField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.SendTime, 9)
    writeLen(w, m.FieldContent, 501)
    binary.Write(w, binary.BigEndian, &m.SequenceSeries)
    binary.Write(w, binary.BigEndian, &m.SequenceNo)
    writeLen(w, m.InvestUnitID, 17)
}

func (m* CThostFtdcTradingNoticeInfoField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.SendTime = readLen(r, 9)
    m.FieldContent = readLen(r, 501)
    binary.Read(r, binary.BigEndian, &m.SequenceSeries)
    binary.Read(r, binary.BigEndian, &m.SequenceNo)
    m.InvestUnitID = readLen(r, 17)
}

type  CThostFtdcTradingNoticeField struct {
    BrokerID string;
    InvestorRange uint8;
    InvestorID string;
    SequenceSeries int16;
    UserID string;
    SendTime string;
    SequenceNo int32;
    FieldContent string;
    InvestUnitID string;
}

func (m* CThostFtdcTradingNoticeField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    binary.Write(w, binary.BigEndian, &m.InvestorRange)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.SequenceSeries)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.SendTime, 9)
    binary.Write(w, binary.BigEndian, &m.SequenceNo)
    writeLen(w, m.FieldContent, 501)
    writeLen(w, m.InvestUnitID, 17)
}

func (m* CThostFtdcTradingNoticeField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    binary.Read(r, binary.BigEndian, &m.InvestorRange)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.SequenceSeries)
    m.UserID = readLen(r, 16)
    m.SendTime = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.SequenceNo)
    m.FieldContent = readLen(r, 501)
    m.InvestUnitID = readLen(r, 17)
}

type  CThostFtdcQryTradingNoticeField struct {
    BrokerID string;
    InvestorID string;
    InvestUnitID string;
}

func (m* CThostFtdcQryTradingNoticeField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InvestUnitID, 17)
}

func (m* CThostFtdcQryTradingNoticeField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InvestUnitID = readLen(r, 17)
}

type  CThostFtdcQryErrOrderField struct {
    BrokerID string;
    InvestorID string;
}

func (m* CThostFtdcQryErrOrderField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
}

func (m* CThostFtdcQryErrOrderField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
}

type  CThostFtdcErrOrderField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    OrderRef string;
    UserID string;
    OrderPriceType uint8;
    Direction uint8;
    CombOffsetFlag string;
    CombHedgeFlag string;
    LimitPrice float64;
    VolumeTotalOriginal int32;
    TimeCondition uint8;
    GTDDate string;
    VolumeCondition uint8;
    MinVolume int32;
    ContingentCondition uint8;
    StopPrice float64;
    ForceCloseReason uint8;
    IsAutoSuspend int32;
    BusinessUnit string;
    RequestID int32;
    UserForceClose int32;
    ErrorID int32;
    ErrorMsg string;
    IsSwapOrder int32;
    ExchangeID string;
    InvestUnitID string;
    AccountID string;
    CurrencyID string;
    ClientID string;
    IPAddress string;
    MacAddress string;
}

func (m* CThostFtdcErrOrderField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.OrderRef, 13)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.OrderPriceType)
    binary.Write(w, binary.BigEndian, &m.Direction)
    writeLen(w, m.CombOffsetFlag, 5)
    writeLen(w, m.CombHedgeFlag, 5)
    binary.Write(w, binary.BigEndian, &m.LimitPrice)
    binary.Write(w, binary.BigEndian, &m.VolumeTotalOriginal)
    binary.Write(w, binary.BigEndian, &m.TimeCondition)
    writeLen(w, m.GTDDate, 9)
    binary.Write(w, binary.BigEndian, &m.VolumeCondition)
    binary.Write(w, binary.BigEndian, &m.MinVolume)
    binary.Write(w, binary.BigEndian, &m.ContingentCondition)
    binary.Write(w, binary.BigEndian, &m.StopPrice)
    binary.Write(w, binary.BigEndian, &m.ForceCloseReason)
    binary.Write(w, binary.BigEndian, &m.IsAutoSuspend)
    writeLen(w, m.BusinessUnit, 21)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    binary.Write(w, binary.BigEndian, &m.UserForceClose)
    binary.Write(w, binary.BigEndian, &m.ErrorID)
    writeLen(w, m.ErrorMsg, 81)
    binary.Write(w, binary.BigEndian, &m.IsSwapOrder)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.InvestUnitID, 17)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.CurrencyID, 4)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
}

func (m* CThostFtdcErrOrderField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.OrderRef = readLen(r, 13)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.OrderPriceType)
    binary.Read(r, binary.BigEndian, &m.Direction)
    m.CombOffsetFlag = readLen(r, 5)
    m.CombHedgeFlag = readLen(r, 5)
    binary.Read(r, binary.BigEndian, &m.LimitPrice)
    binary.Read(r, binary.BigEndian, &m.VolumeTotalOriginal)
    binary.Read(r, binary.BigEndian, &m.TimeCondition)
    m.GTDDate = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.VolumeCondition)
    binary.Read(r, binary.BigEndian, &m.MinVolume)
    binary.Read(r, binary.BigEndian, &m.ContingentCondition)
    binary.Read(r, binary.BigEndian, &m.StopPrice)
    binary.Read(r, binary.BigEndian, &m.ForceCloseReason)
    binary.Read(r, binary.BigEndian, &m.IsAutoSuspend)
    m.BusinessUnit = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    binary.Read(r, binary.BigEndian, &m.UserForceClose)
    binary.Read(r, binary.BigEndian, &m.ErrorID)
    m.ErrorMsg = readLen(r, 81)
    binary.Read(r, binary.BigEndian, &m.IsSwapOrder)
    m.ExchangeID = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
    m.AccountID = readLen(r, 13)
    m.CurrencyID = readLen(r, 4)
    m.ClientID = readLen(r, 11)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
}

type  CThostFtdcErrorConditionalOrderField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    OrderRef string;
    UserID string;
    OrderPriceType uint8;
    Direction uint8;
    CombOffsetFlag string;
    CombHedgeFlag string;
    LimitPrice float64;
    VolumeTotalOriginal int32;
    TimeCondition uint8;
    GTDDate string;
    VolumeCondition uint8;
    MinVolume int32;
    ContingentCondition uint8;
    StopPrice float64;
    ForceCloseReason uint8;
    IsAutoSuspend int32;
    BusinessUnit string;
    RequestID int32;
    OrderLocalID string;
    ExchangeID string;
    ParticipantID string;
    ClientID string;
    ExchangeInstID string;
    TraderID string;
    InstallID int32;
    OrderSubmitStatus uint8;
    NotifySequence int32;
    TradingDay string;
    SettlementID int32;
    OrderSysID string;
    OrderSource uint8;
    OrderStatus uint8;
    OrderType uint8;
    VolumeTraded int32;
    VolumeTotal int32;
    InsertDate string;
    InsertTime string;
    ActiveTime string;
    SuspendTime string;
    UpdateTime string;
    CancelTime string;
    ActiveTraderID string;
    ClearingPartID string;
    SequenceNo int32;
    FrontID int32;
    SessionID int32;
    UserProductInfo string;
    StatusMsg string;
    UserForceClose int32;
    ActiveUserID string;
    BrokerOrderSeq int32;
    RelativeOrderSysID string;
    ZCETotalTradedVolume int32;
    ErrorID int32;
    ErrorMsg string;
    IsSwapOrder int32;
    BranchID string;
    InvestUnitID string;
    AccountID string;
    CurrencyID string;
    IPAddress string;
    MacAddress string;
}

func (m* CThostFtdcErrorConditionalOrderField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.OrderRef, 13)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.OrderPriceType)
    binary.Write(w, binary.BigEndian, &m.Direction)
    writeLen(w, m.CombOffsetFlag, 5)
    writeLen(w, m.CombHedgeFlag, 5)
    binary.Write(w, binary.BigEndian, &m.LimitPrice)
    binary.Write(w, binary.BigEndian, &m.VolumeTotalOriginal)
    binary.Write(w, binary.BigEndian, &m.TimeCondition)
    writeLen(w, m.GTDDate, 9)
    binary.Write(w, binary.BigEndian, &m.VolumeCondition)
    binary.Write(w, binary.BigEndian, &m.MinVolume)
    binary.Write(w, binary.BigEndian, &m.ContingentCondition)
    binary.Write(w, binary.BigEndian, &m.StopPrice)
    binary.Write(w, binary.BigEndian, &m.ForceCloseReason)
    binary.Write(w, binary.BigEndian, &m.IsAutoSuspend)
    writeLen(w, m.BusinessUnit, 21)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    writeLen(w, m.OrderLocalID, 13)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.ExchangeInstID, 31)
    writeLen(w, m.TraderID, 21)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    binary.Write(w, binary.BigEndian, &m.OrderSubmitStatus)
    binary.Write(w, binary.BigEndian, &m.NotifySequence)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.SettlementID)
    writeLen(w, m.OrderSysID, 21)
    binary.Write(w, binary.BigEndian, &m.OrderSource)
    binary.Write(w, binary.BigEndian, &m.OrderStatus)
    binary.Write(w, binary.BigEndian, &m.OrderType)
    binary.Write(w, binary.BigEndian, &m.VolumeTraded)
    binary.Write(w, binary.BigEndian, &m.VolumeTotal)
    writeLen(w, m.InsertDate, 9)
    writeLen(w, m.InsertTime, 9)
    writeLen(w, m.ActiveTime, 9)
    writeLen(w, m.SuspendTime, 9)
    writeLen(w, m.UpdateTime, 9)
    writeLen(w, m.CancelTime, 9)
    writeLen(w, m.ActiveTraderID, 21)
    writeLen(w, m.ClearingPartID, 11)
    binary.Write(w, binary.BigEndian, &m.SequenceNo)
    binary.Write(w, binary.BigEndian, &m.FrontID)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.UserProductInfo, 11)
    writeLen(w, m.StatusMsg, 81)
    binary.Write(w, binary.BigEndian, &m.UserForceClose)
    writeLen(w, m.ActiveUserID, 16)
    binary.Write(w, binary.BigEndian, &m.BrokerOrderSeq)
    writeLen(w, m.RelativeOrderSysID, 21)
    binary.Write(w, binary.BigEndian, &m.ZCETotalTradedVolume)
    binary.Write(w, binary.BigEndian, &m.ErrorID)
    writeLen(w, m.ErrorMsg, 81)
    binary.Write(w, binary.BigEndian, &m.IsSwapOrder)
    writeLen(w, m.BranchID, 9)
    writeLen(w, m.InvestUnitID, 17)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.CurrencyID, 4)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
}

func (m* CThostFtdcErrorConditionalOrderField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    m.OrderRef = readLen(r, 13)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.OrderPriceType)
    binary.Read(r, binary.BigEndian, &m.Direction)
    m.CombOffsetFlag = readLen(r, 5)
    m.CombHedgeFlag = readLen(r, 5)
    binary.Read(r, binary.BigEndian, &m.LimitPrice)
    binary.Read(r, binary.BigEndian, &m.VolumeTotalOriginal)
    binary.Read(r, binary.BigEndian, &m.TimeCondition)
    m.GTDDate = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.VolumeCondition)
    binary.Read(r, binary.BigEndian, &m.MinVolume)
    binary.Read(r, binary.BigEndian, &m.ContingentCondition)
    binary.Read(r, binary.BigEndian, &m.StopPrice)
    binary.Read(r, binary.BigEndian, &m.ForceCloseReason)
    binary.Read(r, binary.BigEndian, &m.IsAutoSuspend)
    m.BusinessUnit = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    m.OrderLocalID = readLen(r, 13)
    m.ExchangeID = readLen(r, 9)
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.ExchangeInstID = readLen(r, 31)
    m.TraderID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    binary.Read(r, binary.BigEndian, &m.OrderSubmitStatus)
    binary.Read(r, binary.BigEndian, &m.NotifySequence)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.SettlementID)
    m.OrderSysID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.OrderSource)
    binary.Read(r, binary.BigEndian, &m.OrderStatus)
    binary.Read(r, binary.BigEndian, &m.OrderType)
    binary.Read(r, binary.BigEndian, &m.VolumeTraded)
    binary.Read(r, binary.BigEndian, &m.VolumeTotal)
    m.InsertDate = readLen(r, 9)
    m.InsertTime = readLen(r, 9)
    m.ActiveTime = readLen(r, 9)
    m.SuspendTime = readLen(r, 9)
    m.UpdateTime = readLen(r, 9)
    m.CancelTime = readLen(r, 9)
    m.ActiveTraderID = readLen(r, 21)
    m.ClearingPartID = readLen(r, 11)
    binary.Read(r, binary.BigEndian, &m.SequenceNo)
    binary.Read(r, binary.BigEndian, &m.FrontID)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.UserProductInfo = readLen(r, 11)
    m.StatusMsg = readLen(r, 81)
    binary.Read(r, binary.BigEndian, &m.UserForceClose)
    m.ActiveUserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.BrokerOrderSeq)
    m.RelativeOrderSysID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.ZCETotalTradedVolume)
    binary.Read(r, binary.BigEndian, &m.ErrorID)
    m.ErrorMsg = readLen(r, 81)
    binary.Read(r, binary.BigEndian, &m.IsSwapOrder)
    m.BranchID = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
    m.AccountID = readLen(r, 13)
    m.CurrencyID = readLen(r, 4)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
}

type  CThostFtdcQryErrOrderActionField struct {
    BrokerID string;
    InvestorID string;
}

func (m* CThostFtdcQryErrOrderActionField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
}

func (m* CThostFtdcQryErrOrderActionField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
}

type  CThostFtdcErrOrderActionField struct {
    BrokerID string;
    InvestorID string;
    OrderActionRef int32;
    OrderRef string;
    RequestID int32;
    FrontID int32;
    SessionID int32;
    ExchangeID string;
    OrderSysID string;
    ActionFlag uint8;
    LimitPrice float64;
    VolumeChange int32;
    ActionDate string;
    ActionTime string;
    TraderID string;
    InstallID int32;
    OrderLocalID string;
    ActionLocalID string;
    ParticipantID string;
    ClientID string;
    BusinessUnit string;
    OrderActionStatus uint8;
    UserID string;
    StatusMsg string;
    InstrumentID string;
    BranchID string;
    InvestUnitID string;
    IPAddress string;
    MacAddress string;
    ErrorID int32;
    ErrorMsg string;
}

func (m* CThostFtdcErrOrderActionField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.OrderActionRef)
    writeLen(w, m.OrderRef, 13)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    binary.Write(w, binary.BigEndian, &m.FrontID)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.OrderSysID, 21)
    binary.Write(w, binary.BigEndian, &m.ActionFlag)
    binary.Write(w, binary.BigEndian, &m.LimitPrice)
    binary.Write(w, binary.BigEndian, &m.VolumeChange)
    writeLen(w, m.ActionDate, 9)
    writeLen(w, m.ActionTime, 9)
    writeLen(w, m.TraderID, 21)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    writeLen(w, m.OrderLocalID, 13)
    writeLen(w, m.ActionLocalID, 13)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.BusinessUnit, 21)
    binary.Write(w, binary.BigEndian, &m.OrderActionStatus)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.StatusMsg, 81)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.BranchID, 9)
    writeLen(w, m.InvestUnitID, 17)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
    binary.Write(w, binary.BigEndian, &m.ErrorID)
    writeLen(w, m.ErrorMsg, 81)
}

func (m* CThostFtdcErrOrderActionField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.OrderActionRef)
    m.OrderRef = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    binary.Read(r, binary.BigEndian, &m.FrontID)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.ExchangeID = readLen(r, 9)
    m.OrderSysID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.ActionFlag)
    binary.Read(r, binary.BigEndian, &m.LimitPrice)
    binary.Read(r, binary.BigEndian, &m.VolumeChange)
    m.ActionDate = readLen(r, 9)
    m.ActionTime = readLen(r, 9)
    m.TraderID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    m.OrderLocalID = readLen(r, 13)
    m.ActionLocalID = readLen(r, 13)
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.BusinessUnit = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.OrderActionStatus)
    m.UserID = readLen(r, 16)
    m.StatusMsg = readLen(r, 81)
    m.InstrumentID = readLen(r, 31)
    m.BranchID = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.ErrorID)
    m.ErrorMsg = readLen(r, 81)
}

type  CThostFtdcQryExchangeSequenceField struct {
    ExchangeID string;
}

func (m* CThostFtdcQryExchangeSequenceField) Marshal(w io.Writer) {
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcQryExchangeSequenceField) Unmarshal(r io.Reader) {
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcExchangeSequenceField struct {
    ExchangeID string;
    SequenceNo int32;
    MarketStatus uint8;
}

func (m* CThostFtdcExchangeSequenceField) Marshal(w io.Writer) {
    writeLen(w, m.ExchangeID, 9)
    binary.Write(w, binary.BigEndian, &m.SequenceNo)
    binary.Write(w, binary.BigEndian, &m.MarketStatus)
}

func (m* CThostFtdcExchangeSequenceField) Unmarshal(r io.Reader) {
    m.ExchangeID = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.SequenceNo)
    binary.Read(r, binary.BigEndian, &m.MarketStatus)
}

type  CThostFtdcQueryMaxOrderVolumeWithPriceField struct {
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    Direction uint8;
    OffsetFlag uint8;
    HedgeFlag uint8;
    MaxVolume int32;
    Price float64;
    ExchangeID string;
    InvestUnitID string;
}

func (m* CThostFtdcQueryMaxOrderVolumeWithPriceField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.Direction)
    binary.Write(w, binary.BigEndian, &m.OffsetFlag)
    binary.Write(w, binary.BigEndian, &m.HedgeFlag)
    binary.Write(w, binary.BigEndian, &m.MaxVolume)
    binary.Write(w, binary.BigEndian, &m.Price)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.InvestUnitID, 17)
}

func (m* CThostFtdcQueryMaxOrderVolumeWithPriceField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.Direction)
    binary.Read(r, binary.BigEndian, &m.OffsetFlag)
    binary.Read(r, binary.BigEndian, &m.HedgeFlag)
    binary.Read(r, binary.BigEndian, &m.MaxVolume)
    binary.Read(r, binary.BigEndian, &m.Price)
    m.ExchangeID = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
}

type  CThostFtdcQryBrokerTradingParamsField struct {
    BrokerID string;
    InvestorID string;
    CurrencyID string;
    AccountID string;
}

func (m* CThostFtdcQryBrokerTradingParamsField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.CurrencyID, 4)
    writeLen(w, m.AccountID, 13)
}

func (m* CThostFtdcQryBrokerTradingParamsField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.CurrencyID = readLen(r, 4)
    m.AccountID = readLen(r, 13)
}

type  CThostFtdcBrokerTradingParamsField struct {
    BrokerID string;
    InvestorID string;
    MarginPriceType uint8;
    Algorithm uint8;
    AvailIncludeCloseProfit uint8;
    CurrencyID string;
    OptionRoyaltyPriceType uint8;
    AccountID string;
}

func (m* CThostFtdcBrokerTradingParamsField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.MarginPriceType)
    binary.Write(w, binary.BigEndian, &m.Algorithm)
    binary.Write(w, binary.BigEndian, &m.AvailIncludeCloseProfit)
    writeLen(w, m.CurrencyID, 4)
    binary.Write(w, binary.BigEndian, &m.OptionRoyaltyPriceType)
    writeLen(w, m.AccountID, 13)
}

func (m* CThostFtdcBrokerTradingParamsField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.MarginPriceType)
    binary.Read(r, binary.BigEndian, &m.Algorithm)
    binary.Read(r, binary.BigEndian, &m.AvailIncludeCloseProfit)
    m.CurrencyID = readLen(r, 4)
    binary.Read(r, binary.BigEndian, &m.OptionRoyaltyPriceType)
    m.AccountID = readLen(r, 13)
}

type  CThostFtdcQryBrokerTradingAlgosField struct {
    BrokerID string;
    ExchangeID string;
    InstrumentID string;
}

func (m* CThostFtdcQryBrokerTradingAlgosField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.InstrumentID, 31)
}

func (m* CThostFtdcQryBrokerTradingAlgosField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.ExchangeID = readLen(r, 9)
    m.InstrumentID = readLen(r, 31)
}

type  CThostFtdcBrokerTradingAlgosField struct {
    BrokerID string;
    ExchangeID string;
    InstrumentID string;
    HandlePositionAlgoID uint8;
    FindMarginRateAlgoID uint8;
    HandleTradingAccountAlgoID uint8;
}

func (m* CThostFtdcBrokerTradingAlgosField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.InstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.HandlePositionAlgoID)
    binary.Write(w, binary.BigEndian, &m.FindMarginRateAlgoID)
    binary.Write(w, binary.BigEndian, &m.HandleTradingAccountAlgoID)
}

func (m* CThostFtdcBrokerTradingAlgosField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.ExchangeID = readLen(r, 9)
    m.InstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.HandlePositionAlgoID)
    binary.Read(r, binary.BigEndian, &m.FindMarginRateAlgoID)
    binary.Read(r, binary.BigEndian, &m.HandleTradingAccountAlgoID)
}

type  CThostFtdcQueryBrokerDepositField struct {
    BrokerID string;
    ExchangeID string;
}

func (m* CThostFtdcQueryBrokerDepositField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcQueryBrokerDepositField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcBrokerDepositField struct {
    TradingDay string;
    BrokerID string;
    ParticipantID string;
    ExchangeID string;
    PreBalance float64;
    CurrMargin float64;
    CloseProfit float64;
    Balance float64;
    Deposit float64;
    Withdraw float64;
    Available float64;
    Reserve float64;
    FrozenMargin float64;
}

func (m* CThostFtdcBrokerDepositField) Marshal(w io.Writer) {
    writeLen(w, m.TradingDay, 9)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ExchangeID, 9)
    binary.Write(w, binary.BigEndian, &m.PreBalance)
    binary.Write(w, binary.BigEndian, &m.CurrMargin)
    binary.Write(w, binary.BigEndian, &m.CloseProfit)
    binary.Write(w, binary.BigEndian, &m.Balance)
    binary.Write(w, binary.BigEndian, &m.Deposit)
    binary.Write(w, binary.BigEndian, &m.Withdraw)
    binary.Write(w, binary.BigEndian, &m.Available)
    binary.Write(w, binary.BigEndian, &m.Reserve)
    binary.Write(w, binary.BigEndian, &m.FrozenMargin)
}

func (m* CThostFtdcBrokerDepositField) Unmarshal(r io.Reader) {
    m.TradingDay = readLen(r, 9)
    m.BrokerID = readLen(r, 11)
    m.ParticipantID = readLen(r, 11)
    m.ExchangeID = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.PreBalance)
    binary.Read(r, binary.BigEndian, &m.CurrMargin)
    binary.Read(r, binary.BigEndian, &m.CloseProfit)
    binary.Read(r, binary.BigEndian, &m.Balance)
    binary.Read(r, binary.BigEndian, &m.Deposit)
    binary.Read(r, binary.BigEndian, &m.Withdraw)
    binary.Read(r, binary.BigEndian, &m.Available)
    binary.Read(r, binary.BigEndian, &m.Reserve)
    binary.Read(r, binary.BigEndian, &m.FrozenMargin)
}

type  CThostFtdcQryCFMMCBrokerKeyField struct {
    BrokerID string;
}

func (m* CThostFtdcQryCFMMCBrokerKeyField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
}

func (m* CThostFtdcQryCFMMCBrokerKeyField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
}

type  CThostFtdcCFMMCBrokerKeyField struct {
    BrokerID string;
    ParticipantID string;
    CreateDate string;
    CreateTime string;
    KeyID int32;
    CurrentKey string;
    KeyKind uint8;
}

func (m* CThostFtdcCFMMCBrokerKeyField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.CreateDate, 9)
    writeLen(w, m.CreateTime, 9)
    binary.Write(w, binary.BigEndian, &m.KeyID)
    writeLen(w, m.CurrentKey, 21)
    binary.Write(w, binary.BigEndian, &m.KeyKind)
}

func (m* CThostFtdcCFMMCBrokerKeyField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.ParticipantID = readLen(r, 11)
    m.CreateDate = readLen(r, 9)
    m.CreateTime = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.KeyID)
    m.CurrentKey = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.KeyKind)
}

type  CThostFtdcCFMMCTradingAccountKeyField struct {
    BrokerID string;
    ParticipantID string;
    AccountID string;
    KeyID int32;
    CurrentKey string;
}

func (m* CThostFtdcCFMMCTradingAccountKeyField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.AccountID, 13)
    binary.Write(w, binary.BigEndian, &m.KeyID)
    writeLen(w, m.CurrentKey, 21)
}

func (m* CThostFtdcCFMMCTradingAccountKeyField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.ParticipantID = readLen(r, 11)
    m.AccountID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.KeyID)
    m.CurrentKey = readLen(r, 21)
}

type  CThostFtdcQryCFMMCTradingAccountKeyField struct {
    BrokerID string;
    InvestorID string;
}

func (m* CThostFtdcQryCFMMCTradingAccountKeyField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
}

func (m* CThostFtdcQryCFMMCTradingAccountKeyField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
}

type  CThostFtdcBrokerUserOTPParamField struct {
    BrokerID string;
    UserID string;
    OTPVendorsID string;
    SerialNumber string;
    AuthKey string;
    LastDrift int32;
    LastSuccess int32;
    OTPType uint8;
}

func (m* CThostFtdcBrokerUserOTPParamField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.OTPVendorsID, 2)
    writeLen(w, m.SerialNumber, 17)
    writeLen(w, m.AuthKey, 41)
    binary.Write(w, binary.BigEndian, &m.LastDrift)
    binary.Write(w, binary.BigEndian, &m.LastSuccess)
    binary.Write(w, binary.BigEndian, &m.OTPType)
}

func (m* CThostFtdcBrokerUserOTPParamField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.UserID = readLen(r, 16)
    m.OTPVendorsID = readLen(r, 2)
    m.SerialNumber = readLen(r, 17)
    m.AuthKey = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.LastDrift)
    binary.Read(r, binary.BigEndian, &m.LastSuccess)
    binary.Read(r, binary.BigEndian, &m.OTPType)
}

type  CThostFtdcManualSyncBrokerUserOTPField struct {
    BrokerID string;
    UserID string;
    OTPType uint8;
    FirstOTP string;
    SecondOTP string;
}

func (m* CThostFtdcManualSyncBrokerUserOTPField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.OTPType)
    writeLen(w, m.FirstOTP, 41)
    writeLen(w, m.SecondOTP, 41)
}

func (m* CThostFtdcManualSyncBrokerUserOTPField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.OTPType)
    m.FirstOTP = readLen(r, 41)
    m.SecondOTP = readLen(r, 41)
}

type  CThostFtdcCommRateModelField struct {
    BrokerID string;
    CommModelID string;
    CommModelName string;
}

func (m* CThostFtdcCommRateModelField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.CommModelID, 13)
    writeLen(w, m.CommModelName, 161)
}

func (m* CThostFtdcCommRateModelField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.CommModelID = readLen(r, 13)
    m.CommModelName = readLen(r, 161)
}

type  CThostFtdcQryCommRateModelField struct {
    BrokerID string;
    CommModelID string;
}

func (m* CThostFtdcQryCommRateModelField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.CommModelID, 13)
}

func (m* CThostFtdcQryCommRateModelField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.CommModelID = readLen(r, 13)
}

type  CThostFtdcMarginModelField struct {
    BrokerID string;
    MarginModelID string;
    MarginModelName string;
}

func (m* CThostFtdcMarginModelField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.MarginModelID, 13)
    writeLen(w, m.MarginModelName, 161)
}

func (m* CThostFtdcMarginModelField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.MarginModelID = readLen(r, 13)
    m.MarginModelName = readLen(r, 161)
}

type  CThostFtdcQryMarginModelField struct {
    BrokerID string;
    MarginModelID string;
}

func (m* CThostFtdcQryMarginModelField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.MarginModelID, 13)
}

func (m* CThostFtdcQryMarginModelField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.MarginModelID = readLen(r, 13)
}

type  CThostFtdcEWarrantOffsetField struct {
    TradingDay string;
    BrokerID string;
    InvestorID string;
    ExchangeID string;
    InstrumentID string;
    Direction uint8;
    HedgeFlag uint8;
    Volume int32;
    InvestUnitID string;
}

func (m* CThostFtdcEWarrantOffsetField) Marshal(w io.Writer) {
    writeLen(w, m.TradingDay, 9)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.InstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.Direction)
    binary.Write(w, binary.BigEndian, &m.HedgeFlag)
    binary.Write(w, binary.BigEndian, &m.Volume)
    writeLen(w, m.InvestUnitID, 17)
}

func (m* CThostFtdcEWarrantOffsetField) Unmarshal(r io.Reader) {
    m.TradingDay = readLen(r, 9)
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.ExchangeID = readLen(r, 9)
    m.InstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.Direction)
    binary.Read(r, binary.BigEndian, &m.HedgeFlag)
    binary.Read(r, binary.BigEndian, &m.Volume)
    m.InvestUnitID = readLen(r, 17)
}

type  CThostFtdcQryEWarrantOffsetField struct {
    BrokerID string;
    InvestorID string;
    ExchangeID string;
    InstrumentID string;
    InvestUnitID string;
}

func (m* CThostFtdcQryEWarrantOffsetField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.InstrumentID, 31)
    writeLen(w, m.InvestUnitID, 17)
}

func (m* CThostFtdcQryEWarrantOffsetField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.ExchangeID = readLen(r, 9)
    m.InstrumentID = readLen(r, 31)
    m.InvestUnitID = readLen(r, 17)
}

type  CThostFtdcQryInvestorProductGroupMarginField struct {
    BrokerID string;
    InvestorID string;
    ProductGroupID string;
    HedgeFlag uint8;
    ExchangeID string;
    InvestUnitID string;
}

func (m* CThostFtdcQryInvestorProductGroupMarginField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.ProductGroupID, 31)
    binary.Write(w, binary.BigEndian, &m.HedgeFlag)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.InvestUnitID, 17)
}

func (m* CThostFtdcQryInvestorProductGroupMarginField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.ProductGroupID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.HedgeFlag)
    m.ExchangeID = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
}

type  CThostFtdcInvestorProductGroupMarginField struct {
    ProductGroupID string;
    BrokerID string;
    InvestorID string;
    TradingDay string;
    SettlementID int32;
    FrozenMargin float64;
    LongFrozenMargin float64;
    ShortFrozenMargin float64;
    UseMargin float64;
    LongUseMargin float64;
    ShortUseMargin float64;
    ExchMargin float64;
    LongExchMargin float64;
    ShortExchMargin float64;
    CloseProfit float64;
    FrozenCommission float64;
    Commission float64;
    FrozenCash float64;
    CashIn float64;
    PositionProfit float64;
    OffsetAmount float64;
    LongOffsetAmount float64;
    ShortOffsetAmount float64;
    ExchOffsetAmount float64;
    LongExchOffsetAmount float64;
    ShortExchOffsetAmount float64;
    HedgeFlag uint8;
    ExchangeID string;
    InvestUnitID string;
}

func (m* CThostFtdcInvestorProductGroupMarginField) Marshal(w io.Writer) {
    writeLen(w, m.ProductGroupID, 31)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.SettlementID)
    binary.Write(w, binary.BigEndian, &m.FrozenMargin)
    binary.Write(w, binary.BigEndian, &m.LongFrozenMargin)
    binary.Write(w, binary.BigEndian, &m.ShortFrozenMargin)
    binary.Write(w, binary.BigEndian, &m.UseMargin)
    binary.Write(w, binary.BigEndian, &m.LongUseMargin)
    binary.Write(w, binary.BigEndian, &m.ShortUseMargin)
    binary.Write(w, binary.BigEndian, &m.ExchMargin)
    binary.Write(w, binary.BigEndian, &m.LongExchMargin)
    binary.Write(w, binary.BigEndian, &m.ShortExchMargin)
    binary.Write(w, binary.BigEndian, &m.CloseProfit)
    binary.Write(w, binary.BigEndian, &m.FrozenCommission)
    binary.Write(w, binary.BigEndian, &m.Commission)
    binary.Write(w, binary.BigEndian, &m.FrozenCash)
    binary.Write(w, binary.BigEndian, &m.CashIn)
    binary.Write(w, binary.BigEndian, &m.PositionProfit)
    binary.Write(w, binary.BigEndian, &m.OffsetAmount)
    binary.Write(w, binary.BigEndian, &m.LongOffsetAmount)
    binary.Write(w, binary.BigEndian, &m.ShortOffsetAmount)
    binary.Write(w, binary.BigEndian, &m.ExchOffsetAmount)
    binary.Write(w, binary.BigEndian, &m.LongExchOffsetAmount)
    binary.Write(w, binary.BigEndian, &m.ShortExchOffsetAmount)
    binary.Write(w, binary.BigEndian, &m.HedgeFlag)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.InvestUnitID, 17)
}

func (m* CThostFtdcInvestorProductGroupMarginField) Unmarshal(r io.Reader) {
    m.ProductGroupID = readLen(r, 31)
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.SettlementID)
    binary.Read(r, binary.BigEndian, &m.FrozenMargin)
    binary.Read(r, binary.BigEndian, &m.LongFrozenMargin)
    binary.Read(r, binary.BigEndian, &m.ShortFrozenMargin)
    binary.Read(r, binary.BigEndian, &m.UseMargin)
    binary.Read(r, binary.BigEndian, &m.LongUseMargin)
    binary.Read(r, binary.BigEndian, &m.ShortUseMargin)
    binary.Read(r, binary.BigEndian, &m.ExchMargin)
    binary.Read(r, binary.BigEndian, &m.LongExchMargin)
    binary.Read(r, binary.BigEndian, &m.ShortExchMargin)
    binary.Read(r, binary.BigEndian, &m.CloseProfit)
    binary.Read(r, binary.BigEndian, &m.FrozenCommission)
    binary.Read(r, binary.BigEndian, &m.Commission)
    binary.Read(r, binary.BigEndian, &m.FrozenCash)
    binary.Read(r, binary.BigEndian, &m.CashIn)
    binary.Read(r, binary.BigEndian, &m.PositionProfit)
    binary.Read(r, binary.BigEndian, &m.OffsetAmount)
    binary.Read(r, binary.BigEndian, &m.LongOffsetAmount)
    binary.Read(r, binary.BigEndian, &m.ShortOffsetAmount)
    binary.Read(r, binary.BigEndian, &m.ExchOffsetAmount)
    binary.Read(r, binary.BigEndian, &m.LongExchOffsetAmount)
    binary.Read(r, binary.BigEndian, &m.ShortExchOffsetAmount)
    binary.Read(r, binary.BigEndian, &m.HedgeFlag)
    m.ExchangeID = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
}

type  CThostFtdcQueryCFMMCTradingAccountTokenField struct {
    BrokerID string;
    InvestorID string;
    InvestUnitID string;
}

func (m* CThostFtdcQueryCFMMCTradingAccountTokenField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InvestUnitID, 17)
}

func (m* CThostFtdcQueryCFMMCTradingAccountTokenField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InvestUnitID = readLen(r, 17)
}

type  CThostFtdcCFMMCTradingAccountTokenField struct {
    BrokerID string;
    ParticipantID string;
    AccountID string;
    KeyID int32;
    Token string;
}

func (m* CThostFtdcCFMMCTradingAccountTokenField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.AccountID, 13)
    binary.Write(w, binary.BigEndian, &m.KeyID)
    writeLen(w, m.Token, 21)
}

func (m* CThostFtdcCFMMCTradingAccountTokenField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.ParticipantID = readLen(r, 11)
    m.AccountID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.KeyID)
    m.Token = readLen(r, 21)
}

type  CThostFtdcInionRightField struct {
    BrokerID string;
    ExchangeID string;
    InvestorID string;
    InstructionRight uint8;
    IsForbidden int32;
}

func (m* CThostFtdcInionRightField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.InstructionRight)
    binary.Write(w, binary.BigEndian, &m.IsForbidden)
}

func (m* CThostFtdcInionRightField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.ExchangeID = readLen(r, 9)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.InstructionRight)
    binary.Read(r, binary.BigEndian, &m.IsForbidden)
}

type  CThostFtdcQryProductGroupField struct {
    ProductID string;
    ExchangeID string;
}

func (m* CThostFtdcQryProductGroupField) Marshal(w io.Writer) {
    writeLen(w, m.ProductID, 31)
    writeLen(w, m.ExchangeID, 9)
}

func (m* CThostFtdcQryProductGroupField) Unmarshal(r io.Reader) {
    m.ProductID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
}

type  CThostFtdcProductGroupField struct {
    ProductID string;
    ExchangeID string;
    ProductGroupID string;
}

func (m* CThostFtdcProductGroupField) Marshal(w io.Writer) {
    writeLen(w, m.ProductID, 31)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ProductGroupID, 31)
}

func (m* CThostFtdcProductGroupField) Unmarshal(r io.Reader) {
    m.ProductID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
    m.ProductGroupID = readLen(r, 31)
}

type  CThostFtdcBulletinField struct {
    ExchangeID string;
    TradingDay string;
    BulletinID int32;
    SequenceNo int32;
    NewsType string;
    NewsUrgency uint8;
    SendTime string;
    Abstract string;
    ComeFrom string;
    Content string;
    URLLink string;
    MarketID string;
}

func (m* CThostFtdcBulletinField) Marshal(w io.Writer) {
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.BulletinID)
    binary.Write(w, binary.BigEndian, &m.SequenceNo)
    writeLen(w, m.NewsType, 3)
    binary.Write(w, binary.BigEndian, &m.NewsUrgency)
    writeLen(w, m.SendTime, 9)
    writeLen(w, m.Abstract, 81)
    writeLen(w, m.ComeFrom, 21)
    writeLen(w, m.Content, 501)
    writeLen(w, m.URLLink, 201)
    writeLen(w, m.MarketID, 31)
}

func (m* CThostFtdcBulletinField) Unmarshal(r io.Reader) {
    m.ExchangeID = readLen(r, 9)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.BulletinID)
    binary.Read(r, binary.BigEndian, &m.SequenceNo)
    m.NewsType = readLen(r, 3)
    binary.Read(r, binary.BigEndian, &m.NewsUrgency)
    m.SendTime = readLen(r, 9)
    m.Abstract = readLen(r, 81)
    m.ComeFrom = readLen(r, 21)
    m.Content = readLen(r, 501)
    m.URLLink = readLen(r, 201)
    m.MarketID = readLen(r, 31)
}

type  CThostFtdcQryBulletinField struct {
    ExchangeID string;
    BulletinID int32;
    SequenceNo int32;
    NewsType string;
    NewsUrgency uint8;
}

func (m* CThostFtdcQryBulletinField) Marshal(w io.Writer) {
    writeLen(w, m.ExchangeID, 9)
    binary.Write(w, binary.BigEndian, &m.BulletinID)
    binary.Write(w, binary.BigEndian, &m.SequenceNo)
    writeLen(w, m.NewsType, 3)
    binary.Write(w, binary.BigEndian, &m.NewsUrgency)
}

func (m* CThostFtdcQryBulletinField) Unmarshal(r io.Reader) {
    m.ExchangeID = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.BulletinID)
    binary.Read(r, binary.BigEndian, &m.SequenceNo)
    m.NewsType = readLen(r, 3)
    binary.Read(r, binary.BigEndian, &m.NewsUrgency)
}

type  CThostFtdcReqOpenAccountField struct {
    TradeCode string;
    BankID string;
    BankBranchID string;
    BrokerID string;
    BrokerBranchID string;
    TradeDate string;
    TradeTime string;
    BankSerial string;
    TradingDay string;
    PlateSerial int32;
    LastFragment uint8;
    SessionID int32;
    CustomerName string;
    IdCardType uint8;
    IdentifiedCardNo string;
    Gender uint8;
    CountryCode string;
    CustType uint8;
    Address string;
    ZipCode string;
    Telephone string;
    MobilePhone string;
    Fax string;
    EMail string;
    MoneyAccountStatus uint8;
    BankAccount string;
    BankPassWord string;
    AccountID string;
    Password string;
    InstallID int32;
    VerifyCertNoFlag uint8;
    CurrencyID string;
    CashExchangeCode uint8;
    Digest string;
    BankAccType uint8;
    DeviceID string;
    BankSecuAccType uint8;
    BrokerIDByBank string;
    BankSecuAcc string;
    BankPwdFlag uint8;
    SecuPwdFlag uint8;
    OperNo string;
    TID int32;
    UserID string;
    SecDaBeBl float64;
    BankChal uint8;
}

func (m* CThostFtdcReqOpenAccountField) Marshal(w io.Writer) {
    writeLen(w, m.TradeCode, 7)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBranchID, 5)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.BrokerBranchID, 31)
    writeLen(w, m.TradeDate, 9)
    writeLen(w, m.TradeTime, 9)
    writeLen(w, m.BankSerial, 13)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.PlateSerial)
    binary.Write(w, binary.BigEndian, &m.LastFragment)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.CustomerName, 51)
    binary.Write(w, binary.BigEndian, &m.IdCardType)
    writeLen(w, m.IdentifiedCardNo, 51)
    binary.Write(w, binary.BigEndian, &m.Gender)
    writeLen(w, m.CountryCode, 21)
    binary.Write(w, binary.BigEndian, &m.CustType)
    writeLen(w, m.Address, 101)
    writeLen(w, m.ZipCode, 7)
    writeLen(w, m.Telephone, 41)
    writeLen(w, m.MobilePhone, 21)
    writeLen(w, m.Fax, 41)
    writeLen(w, m.EMail, 41)
    binary.Write(w, binary.BigEndian, &m.MoneyAccountStatus)
    writeLen(w, m.BankAccount, 41)
    writeLen(w, m.BankPassWord, 41)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.Password, 41)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    binary.Write(w, binary.BigEndian, &m.VerifyCertNoFlag)
    writeLen(w, m.CurrencyID, 4)
    binary.Write(w, binary.BigEndian, &m.CashExchangeCode)
    writeLen(w, m.Digest, 36)
    binary.Write(w, binary.BigEndian, &m.BankAccType)
    writeLen(w, m.DeviceID, 3)
    binary.Write(w, binary.BigEndian, &m.BankSecuAccType)
    writeLen(w, m.BrokerIDByBank, 33)
    writeLen(w, m.BankSecuAcc, 41)
    binary.Write(w, binary.BigEndian, &m.BankPwdFlag)
    binary.Write(w, binary.BigEndian, &m.SecuPwdFlag)
    writeLen(w, m.OperNo, 17)
    binary.Write(w, binary.BigEndian, &m.TID)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.SecDaBeBl)
    binary.Write(w, binary.BigEndian, &m.BankChal)
}

func (m* CThostFtdcReqOpenAccountField) Unmarshal(r io.Reader) {
    m.TradeCode = readLen(r, 7)
    m.BankID = readLen(r, 4)
    m.BankBranchID = readLen(r, 5)
    m.BrokerID = readLen(r, 11)
    m.BrokerBranchID = readLen(r, 31)
    m.TradeDate = readLen(r, 9)
    m.TradeTime = readLen(r, 9)
    m.BankSerial = readLen(r, 13)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.PlateSerial)
    binary.Read(r, binary.BigEndian, &m.LastFragment)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.CustomerName = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.IdCardType)
    m.IdentifiedCardNo = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.Gender)
    m.CountryCode = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.CustType)
    m.Address = readLen(r, 101)
    m.ZipCode = readLen(r, 7)
    m.Telephone = readLen(r, 41)
    m.MobilePhone = readLen(r, 21)
    m.Fax = readLen(r, 41)
    m.EMail = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.MoneyAccountStatus)
    m.BankAccount = readLen(r, 41)
    m.BankPassWord = readLen(r, 41)
    m.AccountID = readLen(r, 13)
    m.Password = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    binary.Read(r, binary.BigEndian, &m.VerifyCertNoFlag)
    m.CurrencyID = readLen(r, 4)
    binary.Read(r, binary.BigEndian, &m.CashExchangeCode)
    m.Digest = readLen(r, 36)
    binary.Read(r, binary.BigEndian, &m.BankAccType)
    m.DeviceID = readLen(r, 3)
    binary.Read(r, binary.BigEndian, &m.BankSecuAccType)
    m.BrokerIDByBank = readLen(r, 33)
    m.BankSecuAcc = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.BankPwdFlag)
    binary.Read(r, binary.BigEndian, &m.SecuPwdFlag)
    m.OperNo = readLen(r, 17)
    binary.Read(r, binary.BigEndian, &m.TID)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.SecDaBeBl)
    binary.Read(r, binary.BigEndian, &m.BankChal)
}

type  CThostFtdcReqCancelAccountField struct {
    TradeCode string;
    BankID string;
    BankBranchID string;
    BrokerID string;
    BrokerBranchID string;
    TradeDate string;
    TradeTime string;
    BankSerial string;
    TradingDay string;
    PlateSerial int32;
    LastFragment uint8;
    SessionID int32;
    CustomerName string;
    IdCardType uint8;
    IdentifiedCardNo string;
    Gender uint8;
    CountryCode string;
    CustType uint8;
    Address string;
    ZipCode string;
    Telephone string;
    MobilePhone string;
    Fax string;
    EMail string;
    MoneyAccountStatus uint8;
    BankAccount string;
    BankPassWord string;
    AccountID string;
    Password string;
    InstallID int32;
    VerifyCertNoFlag uint8;
    CurrencyID string;
    CashExchangeCode uint8;
    Digest string;
    BankAccType uint8;
    DeviceID string;
    BankSecuAccType uint8;
    BrokerIDByBank string;
    BankSecuAcc string;
    BankPwdFlag uint8;
    SecuPwdFlag uint8;
    OperNo string;
    TID int32;
    UserID string;
    SecDaBeBl float64;
    BankChal uint8;
}

func (m* CThostFtdcReqCancelAccountField) Marshal(w io.Writer) {
    writeLen(w, m.TradeCode, 7)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBranchID, 5)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.BrokerBranchID, 31)
    writeLen(w, m.TradeDate, 9)
    writeLen(w, m.TradeTime, 9)
    writeLen(w, m.BankSerial, 13)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.PlateSerial)
    binary.Write(w, binary.BigEndian, &m.LastFragment)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.CustomerName, 51)
    binary.Write(w, binary.BigEndian, &m.IdCardType)
    writeLen(w, m.IdentifiedCardNo, 51)
    binary.Write(w, binary.BigEndian, &m.Gender)
    writeLen(w, m.CountryCode, 21)
    binary.Write(w, binary.BigEndian, &m.CustType)
    writeLen(w, m.Address, 101)
    writeLen(w, m.ZipCode, 7)
    writeLen(w, m.Telephone, 41)
    writeLen(w, m.MobilePhone, 21)
    writeLen(w, m.Fax, 41)
    writeLen(w, m.EMail, 41)
    binary.Write(w, binary.BigEndian, &m.MoneyAccountStatus)
    writeLen(w, m.BankAccount, 41)
    writeLen(w, m.BankPassWord, 41)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.Password, 41)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    binary.Write(w, binary.BigEndian, &m.VerifyCertNoFlag)
    writeLen(w, m.CurrencyID, 4)
    binary.Write(w, binary.BigEndian, &m.CashExchangeCode)
    writeLen(w, m.Digest, 36)
    binary.Write(w, binary.BigEndian, &m.BankAccType)
    writeLen(w, m.DeviceID, 3)
    binary.Write(w, binary.BigEndian, &m.BankSecuAccType)
    writeLen(w, m.BrokerIDByBank, 33)
    writeLen(w, m.BankSecuAcc, 41)
    binary.Write(w, binary.BigEndian, &m.BankPwdFlag)
    binary.Write(w, binary.BigEndian, &m.SecuPwdFlag)
    writeLen(w, m.OperNo, 17)
    binary.Write(w, binary.BigEndian, &m.TID)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.SecDaBeBl)
    binary.Write(w, binary.BigEndian, &m.BankChal)
}

func (m* CThostFtdcReqCancelAccountField) Unmarshal(r io.Reader) {
    m.TradeCode = readLen(r, 7)
    m.BankID = readLen(r, 4)
    m.BankBranchID = readLen(r, 5)
    m.BrokerID = readLen(r, 11)
    m.BrokerBranchID = readLen(r, 31)
    m.TradeDate = readLen(r, 9)
    m.TradeTime = readLen(r, 9)
    m.BankSerial = readLen(r, 13)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.PlateSerial)
    binary.Read(r, binary.BigEndian, &m.LastFragment)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.CustomerName = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.IdCardType)
    m.IdentifiedCardNo = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.Gender)
    m.CountryCode = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.CustType)
    m.Address = readLen(r, 101)
    m.ZipCode = readLen(r, 7)
    m.Telephone = readLen(r, 41)
    m.MobilePhone = readLen(r, 21)
    m.Fax = readLen(r, 41)
    m.EMail = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.MoneyAccountStatus)
    m.BankAccount = readLen(r, 41)
    m.BankPassWord = readLen(r, 41)
    m.AccountID = readLen(r, 13)
    m.Password = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    binary.Read(r, binary.BigEndian, &m.VerifyCertNoFlag)
    m.CurrencyID = readLen(r, 4)
    binary.Read(r, binary.BigEndian, &m.CashExchangeCode)
    m.Digest = readLen(r, 36)
    binary.Read(r, binary.BigEndian, &m.BankAccType)
    m.DeviceID = readLen(r, 3)
    binary.Read(r, binary.BigEndian, &m.BankSecuAccType)
    m.BrokerIDByBank = readLen(r, 33)
    m.BankSecuAcc = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.BankPwdFlag)
    binary.Read(r, binary.BigEndian, &m.SecuPwdFlag)
    m.OperNo = readLen(r, 17)
    binary.Read(r, binary.BigEndian, &m.TID)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.SecDaBeBl)
    binary.Read(r, binary.BigEndian, &m.BankChal)
}

type  CThostFtdcReqChangeAccountField struct {
    TradeCode string;
    BankID string;
    BankBranchID string;
    BrokerID string;
    BrokerBranchID string;
    TradeDate string;
    TradeTime string;
    BankSerial string;
    TradingDay string;
    PlateSerial int32;
    LastFragment uint8;
    SessionID int32;
    CustomerName string;
    IdCardType uint8;
    IdentifiedCardNo string;
    Gender uint8;
    CountryCode string;
    CustType uint8;
    Address string;
    ZipCode string;
    Telephone string;
    MobilePhone string;
    Fax string;
    EMail string;
    MoneyAccountStatus uint8;
    BankAccount string;
    BankPassWord string;
    NewBankAccount string;
    NewBankPassWord string;
    AccountID string;
    Password string;
    BankAccType uint8;
    InstallID int32;
    VerifyCertNoFlag uint8;
    CurrencyID string;
    BrokerIDByBank string;
    BankPwdFlag uint8;
    SecuPwdFlag uint8;
    TID int32;
    Digest string;
}

func (m* CThostFtdcReqChangeAccountField) Marshal(w io.Writer) {
    writeLen(w, m.TradeCode, 7)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBranchID, 5)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.BrokerBranchID, 31)
    writeLen(w, m.TradeDate, 9)
    writeLen(w, m.TradeTime, 9)
    writeLen(w, m.BankSerial, 13)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.PlateSerial)
    binary.Write(w, binary.BigEndian, &m.LastFragment)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.CustomerName, 51)
    binary.Write(w, binary.BigEndian, &m.IdCardType)
    writeLen(w, m.IdentifiedCardNo, 51)
    binary.Write(w, binary.BigEndian, &m.Gender)
    writeLen(w, m.CountryCode, 21)
    binary.Write(w, binary.BigEndian, &m.CustType)
    writeLen(w, m.Address, 101)
    writeLen(w, m.ZipCode, 7)
    writeLen(w, m.Telephone, 41)
    writeLen(w, m.MobilePhone, 21)
    writeLen(w, m.Fax, 41)
    writeLen(w, m.EMail, 41)
    binary.Write(w, binary.BigEndian, &m.MoneyAccountStatus)
    writeLen(w, m.BankAccount, 41)
    writeLen(w, m.BankPassWord, 41)
    writeLen(w, m.NewBankAccount, 41)
    writeLen(w, m.NewBankPassWord, 41)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.Password, 41)
    binary.Write(w, binary.BigEndian, &m.BankAccType)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    binary.Write(w, binary.BigEndian, &m.VerifyCertNoFlag)
    writeLen(w, m.CurrencyID, 4)
    writeLen(w, m.BrokerIDByBank, 33)
    binary.Write(w, binary.BigEndian, &m.BankPwdFlag)
    binary.Write(w, binary.BigEndian, &m.SecuPwdFlag)
    binary.Write(w, binary.BigEndian, &m.TID)
    writeLen(w, m.Digest, 36)
}

func (m* CThostFtdcReqChangeAccountField) Unmarshal(r io.Reader) {
    m.TradeCode = readLen(r, 7)
    m.BankID = readLen(r, 4)
    m.BankBranchID = readLen(r, 5)
    m.BrokerID = readLen(r, 11)
    m.BrokerBranchID = readLen(r, 31)
    m.TradeDate = readLen(r, 9)
    m.TradeTime = readLen(r, 9)
    m.BankSerial = readLen(r, 13)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.PlateSerial)
    binary.Read(r, binary.BigEndian, &m.LastFragment)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.CustomerName = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.IdCardType)
    m.IdentifiedCardNo = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.Gender)
    m.CountryCode = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.CustType)
    m.Address = readLen(r, 101)
    m.ZipCode = readLen(r, 7)
    m.Telephone = readLen(r, 41)
    m.MobilePhone = readLen(r, 21)
    m.Fax = readLen(r, 41)
    m.EMail = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.MoneyAccountStatus)
    m.BankAccount = readLen(r, 41)
    m.BankPassWord = readLen(r, 41)
    m.NewBankAccount = readLen(r, 41)
    m.NewBankPassWord = readLen(r, 41)
    m.AccountID = readLen(r, 13)
    m.Password = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.BankAccType)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    binary.Read(r, binary.BigEndian, &m.VerifyCertNoFlag)
    m.CurrencyID = readLen(r, 4)
    m.BrokerIDByBank = readLen(r, 33)
    binary.Read(r, binary.BigEndian, &m.BankPwdFlag)
    binary.Read(r, binary.BigEndian, &m.SecuPwdFlag)
    binary.Read(r, binary.BigEndian, &m.TID)
    m.Digest = readLen(r, 36)
}

type  CThostFtdcReqTransferField struct {
    TradeCode string;
    BankID string;
    BankBranchID string;
    BrokerID string;
    BrokerBranchID string;
    TradeDate string;
    TradeTime string;
    BankSerial string;
    TradingDay string;
    PlateSerial int32;
    LastFragment uint8;
    SessionID int32;
    CustomerName string;
    IdCardType uint8;
    IdentifiedCardNo string;
    CustType uint8;
    BankAccount string;
    BankPassWord string;
    AccountID string;
    Password string;
    InstallID int32;
    FutureSerial int32;
    UserID string;
    VerifyCertNoFlag uint8;
    CurrencyID string;
    TradeAmount float64;
    FutureFetchAmount float64;
    FeePayFlag uint8;
    CustFee float64;
    BrokerFee float64;
    Message string;
    Digest string;
    BankAccType uint8;
    DeviceID string;
    BankSecuAccType uint8;
    BrokerIDByBank string;
    BankSecuAcc string;
    BankPwdFlag uint8;
    SecuPwdFlag uint8;
    OperNo string;
    RequestID int32;
    TID int32;
    TransferStatus uint8;
}

func (m* CThostFtdcReqTransferField) Marshal(w io.Writer) {
    writeLen(w, m.TradeCode, 7)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBranchID, 5)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.BrokerBranchID, 31)
    writeLen(w, m.TradeDate, 9)
    writeLen(w, m.TradeTime, 9)
    writeLen(w, m.BankSerial, 13)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.PlateSerial)
    binary.Write(w, binary.BigEndian, &m.LastFragment)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.CustomerName, 51)
    binary.Write(w, binary.BigEndian, &m.IdCardType)
    writeLen(w, m.IdentifiedCardNo, 51)
    binary.Write(w, binary.BigEndian, &m.CustType)
    writeLen(w, m.BankAccount, 41)
    writeLen(w, m.BankPassWord, 41)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.Password, 41)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    binary.Write(w, binary.BigEndian, &m.FutureSerial)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.VerifyCertNoFlag)
    writeLen(w, m.CurrencyID, 4)
    binary.Write(w, binary.BigEndian, &m.TradeAmount)
    binary.Write(w, binary.BigEndian, &m.FutureFetchAmount)
    binary.Write(w, binary.BigEndian, &m.FeePayFlag)
    binary.Write(w, binary.BigEndian, &m.CustFee)
    binary.Write(w, binary.BigEndian, &m.BrokerFee)
    writeLen(w, m.Message, 129)
    writeLen(w, m.Digest, 36)
    binary.Write(w, binary.BigEndian, &m.BankAccType)
    writeLen(w, m.DeviceID, 3)
    binary.Write(w, binary.BigEndian, &m.BankSecuAccType)
    writeLen(w, m.BrokerIDByBank, 33)
    writeLen(w, m.BankSecuAcc, 41)
    binary.Write(w, binary.BigEndian, &m.BankPwdFlag)
    binary.Write(w, binary.BigEndian, &m.SecuPwdFlag)
    writeLen(w, m.OperNo, 17)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    binary.Write(w, binary.BigEndian, &m.TID)
    binary.Write(w, binary.BigEndian, &m.TransferStatus)
}

func (m* CThostFtdcReqTransferField) Unmarshal(r io.Reader) {
    m.TradeCode = readLen(r, 7)
    m.BankID = readLen(r, 4)
    m.BankBranchID = readLen(r, 5)
    m.BrokerID = readLen(r, 11)
    m.BrokerBranchID = readLen(r, 31)
    m.TradeDate = readLen(r, 9)
    m.TradeTime = readLen(r, 9)
    m.BankSerial = readLen(r, 13)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.PlateSerial)
    binary.Read(r, binary.BigEndian, &m.LastFragment)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.CustomerName = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.IdCardType)
    m.IdentifiedCardNo = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.CustType)
    m.BankAccount = readLen(r, 41)
    m.BankPassWord = readLen(r, 41)
    m.AccountID = readLen(r, 13)
    m.Password = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    binary.Read(r, binary.BigEndian, &m.FutureSerial)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.VerifyCertNoFlag)
    m.CurrencyID = readLen(r, 4)
    binary.Read(r, binary.BigEndian, &m.TradeAmount)
    binary.Read(r, binary.BigEndian, &m.FutureFetchAmount)
    binary.Read(r, binary.BigEndian, &m.FeePayFlag)
    binary.Read(r, binary.BigEndian, &m.CustFee)
    binary.Read(r, binary.BigEndian, &m.BrokerFee)
    m.Message = readLen(r, 129)
    m.Digest = readLen(r, 36)
    binary.Read(r, binary.BigEndian, &m.BankAccType)
    m.DeviceID = readLen(r, 3)
    binary.Read(r, binary.BigEndian, &m.BankSecuAccType)
    m.BrokerIDByBank = readLen(r, 33)
    m.BankSecuAcc = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.BankPwdFlag)
    binary.Read(r, binary.BigEndian, &m.SecuPwdFlag)
    m.OperNo = readLen(r, 17)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    binary.Read(r, binary.BigEndian, &m.TID)
    binary.Read(r, binary.BigEndian, &m.TransferStatus)
}

type  CThostFtdcRspTransferField struct {
    TradeCode string;
    BankID string;
    BankBranchID string;
    BrokerID string;
    BrokerBranchID string;
    TradeDate string;
    TradeTime string;
    BankSerial string;
    TradingDay string;
    PlateSerial int32;
    LastFragment uint8;
    SessionID int32;
    CustomerName string;
    IdCardType uint8;
    IdentifiedCardNo string;
    CustType uint8;
    BankAccount string;
    BankPassWord string;
    AccountID string;
    Password string;
    InstallID int32;
    FutureSerial int32;
    UserID string;
    VerifyCertNoFlag uint8;
    CurrencyID string;
    TradeAmount float64;
    FutureFetchAmount float64;
    FeePayFlag uint8;
    CustFee float64;
    BrokerFee float64;
    Message string;
    Digest string;
    BankAccType uint8;
    DeviceID string;
    BankSecuAccType uint8;
    BrokerIDByBank string;
    BankSecuAcc string;
    BankPwdFlag uint8;
    SecuPwdFlag uint8;
    OperNo string;
    RequestID int32;
    TID int32;
    TransferStatus uint8;
    ErrorID int32;
    ErrorMsg string;
}

func (m* CThostFtdcRspTransferField) Marshal(w io.Writer) {
    writeLen(w, m.TradeCode, 7)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBranchID, 5)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.BrokerBranchID, 31)
    writeLen(w, m.TradeDate, 9)
    writeLen(w, m.TradeTime, 9)
    writeLen(w, m.BankSerial, 13)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.PlateSerial)
    binary.Write(w, binary.BigEndian, &m.LastFragment)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.CustomerName, 51)
    binary.Write(w, binary.BigEndian, &m.IdCardType)
    writeLen(w, m.IdentifiedCardNo, 51)
    binary.Write(w, binary.BigEndian, &m.CustType)
    writeLen(w, m.BankAccount, 41)
    writeLen(w, m.BankPassWord, 41)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.Password, 41)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    binary.Write(w, binary.BigEndian, &m.FutureSerial)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.VerifyCertNoFlag)
    writeLen(w, m.CurrencyID, 4)
    binary.Write(w, binary.BigEndian, &m.TradeAmount)
    binary.Write(w, binary.BigEndian, &m.FutureFetchAmount)
    binary.Write(w, binary.BigEndian, &m.FeePayFlag)
    binary.Write(w, binary.BigEndian, &m.CustFee)
    binary.Write(w, binary.BigEndian, &m.BrokerFee)
    writeLen(w, m.Message, 129)
    writeLen(w, m.Digest, 36)
    binary.Write(w, binary.BigEndian, &m.BankAccType)
    writeLen(w, m.DeviceID, 3)
    binary.Write(w, binary.BigEndian, &m.BankSecuAccType)
    writeLen(w, m.BrokerIDByBank, 33)
    writeLen(w, m.BankSecuAcc, 41)
    binary.Write(w, binary.BigEndian, &m.BankPwdFlag)
    binary.Write(w, binary.BigEndian, &m.SecuPwdFlag)
    writeLen(w, m.OperNo, 17)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    binary.Write(w, binary.BigEndian, &m.TID)
    binary.Write(w, binary.BigEndian, &m.TransferStatus)
    binary.Write(w, binary.BigEndian, &m.ErrorID)
    writeLen(w, m.ErrorMsg, 81)
}

func (m* CThostFtdcRspTransferField) Unmarshal(r io.Reader) {
    m.TradeCode = readLen(r, 7)
    m.BankID = readLen(r, 4)
    m.BankBranchID = readLen(r, 5)
    m.BrokerID = readLen(r, 11)
    m.BrokerBranchID = readLen(r, 31)
    m.TradeDate = readLen(r, 9)
    m.TradeTime = readLen(r, 9)
    m.BankSerial = readLen(r, 13)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.PlateSerial)
    binary.Read(r, binary.BigEndian, &m.LastFragment)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.CustomerName = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.IdCardType)
    m.IdentifiedCardNo = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.CustType)
    m.BankAccount = readLen(r, 41)
    m.BankPassWord = readLen(r, 41)
    m.AccountID = readLen(r, 13)
    m.Password = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    binary.Read(r, binary.BigEndian, &m.FutureSerial)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.VerifyCertNoFlag)
    m.CurrencyID = readLen(r, 4)
    binary.Read(r, binary.BigEndian, &m.TradeAmount)
    binary.Read(r, binary.BigEndian, &m.FutureFetchAmount)
    binary.Read(r, binary.BigEndian, &m.FeePayFlag)
    binary.Read(r, binary.BigEndian, &m.CustFee)
    binary.Read(r, binary.BigEndian, &m.BrokerFee)
    m.Message = readLen(r, 129)
    m.Digest = readLen(r, 36)
    binary.Read(r, binary.BigEndian, &m.BankAccType)
    m.DeviceID = readLen(r, 3)
    binary.Read(r, binary.BigEndian, &m.BankSecuAccType)
    m.BrokerIDByBank = readLen(r, 33)
    m.BankSecuAcc = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.BankPwdFlag)
    binary.Read(r, binary.BigEndian, &m.SecuPwdFlag)
    m.OperNo = readLen(r, 17)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    binary.Read(r, binary.BigEndian, &m.TID)
    binary.Read(r, binary.BigEndian, &m.TransferStatus)
    binary.Read(r, binary.BigEndian, &m.ErrorID)
    m.ErrorMsg = readLen(r, 81)
}

type  CThostFtdcReqRepealField struct {
    RepealTimeInterval int32;
    RepealedTimes int32;
    BankRepealFlag uint8;
    BrokerRepealFlag uint8;
    PlateRepealSerial int32;
    BankRepealSerial string;
    FutureRepealSerial int32;
    TradeCode string;
    BankID string;
    BankBranchID string;
    BrokerID string;
    BrokerBranchID string;
    TradeDate string;
    TradeTime string;
    BankSerial string;
    TradingDay string;
    PlateSerial int32;
    LastFragment uint8;
    SessionID int32;
    CustomerName string;
    IdCardType uint8;
    IdentifiedCardNo string;
    CustType uint8;
    BankAccount string;
    BankPassWord string;
    AccountID string;
    Password string;
    InstallID int32;
    FutureSerial int32;
    UserID string;
    VerifyCertNoFlag uint8;
    CurrencyID string;
    TradeAmount float64;
    FutureFetchAmount float64;
    FeePayFlag uint8;
    CustFee float64;
    BrokerFee float64;
    Message string;
    Digest string;
    BankAccType uint8;
    DeviceID string;
    BankSecuAccType uint8;
    BrokerIDByBank string;
    BankSecuAcc string;
    BankPwdFlag uint8;
    SecuPwdFlag uint8;
    OperNo string;
    RequestID int32;
    TID int32;
    TransferStatus uint8;
}

func (m* CThostFtdcReqRepealField) Marshal(w io.Writer) {
    binary.Write(w, binary.BigEndian, &m.RepealTimeInterval)
    binary.Write(w, binary.BigEndian, &m.RepealedTimes)
    binary.Write(w, binary.BigEndian, &m.BankRepealFlag)
    binary.Write(w, binary.BigEndian, &m.BrokerRepealFlag)
    binary.Write(w, binary.BigEndian, &m.PlateRepealSerial)
    writeLen(w, m.BankRepealSerial, 13)
    binary.Write(w, binary.BigEndian, &m.FutureRepealSerial)
    writeLen(w, m.TradeCode, 7)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBranchID, 5)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.BrokerBranchID, 31)
    writeLen(w, m.TradeDate, 9)
    writeLen(w, m.TradeTime, 9)
    writeLen(w, m.BankSerial, 13)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.PlateSerial)
    binary.Write(w, binary.BigEndian, &m.LastFragment)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.CustomerName, 51)
    binary.Write(w, binary.BigEndian, &m.IdCardType)
    writeLen(w, m.IdentifiedCardNo, 51)
    binary.Write(w, binary.BigEndian, &m.CustType)
    writeLen(w, m.BankAccount, 41)
    writeLen(w, m.BankPassWord, 41)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.Password, 41)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    binary.Write(w, binary.BigEndian, &m.FutureSerial)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.VerifyCertNoFlag)
    writeLen(w, m.CurrencyID, 4)
    binary.Write(w, binary.BigEndian, &m.TradeAmount)
    binary.Write(w, binary.BigEndian, &m.FutureFetchAmount)
    binary.Write(w, binary.BigEndian, &m.FeePayFlag)
    binary.Write(w, binary.BigEndian, &m.CustFee)
    binary.Write(w, binary.BigEndian, &m.BrokerFee)
    writeLen(w, m.Message, 129)
    writeLen(w, m.Digest, 36)
    binary.Write(w, binary.BigEndian, &m.BankAccType)
    writeLen(w, m.DeviceID, 3)
    binary.Write(w, binary.BigEndian, &m.BankSecuAccType)
    writeLen(w, m.BrokerIDByBank, 33)
    writeLen(w, m.BankSecuAcc, 41)
    binary.Write(w, binary.BigEndian, &m.BankPwdFlag)
    binary.Write(w, binary.BigEndian, &m.SecuPwdFlag)
    writeLen(w, m.OperNo, 17)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    binary.Write(w, binary.BigEndian, &m.TID)
    binary.Write(w, binary.BigEndian, &m.TransferStatus)
}

func (m* CThostFtdcReqRepealField) Unmarshal(r io.Reader) {
    binary.Read(r, binary.BigEndian, &m.RepealTimeInterval)
    binary.Read(r, binary.BigEndian, &m.RepealedTimes)
    binary.Read(r, binary.BigEndian, &m.BankRepealFlag)
    binary.Read(r, binary.BigEndian, &m.BrokerRepealFlag)
    binary.Read(r, binary.BigEndian, &m.PlateRepealSerial)
    m.BankRepealSerial = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.FutureRepealSerial)
    m.TradeCode = readLen(r, 7)
    m.BankID = readLen(r, 4)
    m.BankBranchID = readLen(r, 5)
    m.BrokerID = readLen(r, 11)
    m.BrokerBranchID = readLen(r, 31)
    m.TradeDate = readLen(r, 9)
    m.TradeTime = readLen(r, 9)
    m.BankSerial = readLen(r, 13)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.PlateSerial)
    binary.Read(r, binary.BigEndian, &m.LastFragment)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.CustomerName = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.IdCardType)
    m.IdentifiedCardNo = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.CustType)
    m.BankAccount = readLen(r, 41)
    m.BankPassWord = readLen(r, 41)
    m.AccountID = readLen(r, 13)
    m.Password = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    binary.Read(r, binary.BigEndian, &m.FutureSerial)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.VerifyCertNoFlag)
    m.CurrencyID = readLen(r, 4)
    binary.Read(r, binary.BigEndian, &m.TradeAmount)
    binary.Read(r, binary.BigEndian, &m.FutureFetchAmount)
    binary.Read(r, binary.BigEndian, &m.FeePayFlag)
    binary.Read(r, binary.BigEndian, &m.CustFee)
    binary.Read(r, binary.BigEndian, &m.BrokerFee)
    m.Message = readLen(r, 129)
    m.Digest = readLen(r, 36)
    binary.Read(r, binary.BigEndian, &m.BankAccType)
    m.DeviceID = readLen(r, 3)
    binary.Read(r, binary.BigEndian, &m.BankSecuAccType)
    m.BrokerIDByBank = readLen(r, 33)
    m.BankSecuAcc = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.BankPwdFlag)
    binary.Read(r, binary.BigEndian, &m.SecuPwdFlag)
    m.OperNo = readLen(r, 17)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    binary.Read(r, binary.BigEndian, &m.TID)
    binary.Read(r, binary.BigEndian, &m.TransferStatus)
}

type  CThostFtdcRspRepealField struct {
    RepealTimeInterval int32;
    RepealedTimes int32;
    BankRepealFlag uint8;
    BrokerRepealFlag uint8;
    PlateRepealSerial int32;
    BankRepealSerial string;
    FutureRepealSerial int32;
    TradeCode string;
    BankID string;
    BankBranchID string;
    BrokerID string;
    BrokerBranchID string;
    TradeDate string;
    TradeTime string;
    BankSerial string;
    TradingDay string;
    PlateSerial int32;
    LastFragment uint8;
    SessionID int32;
    CustomerName string;
    IdCardType uint8;
    IdentifiedCardNo string;
    CustType uint8;
    BankAccount string;
    BankPassWord string;
    AccountID string;
    Password string;
    InstallID int32;
    FutureSerial int32;
    UserID string;
    VerifyCertNoFlag uint8;
    CurrencyID string;
    TradeAmount float64;
    FutureFetchAmount float64;
    FeePayFlag uint8;
    CustFee float64;
    BrokerFee float64;
    Message string;
    Digest string;
    BankAccType uint8;
    DeviceID string;
    BankSecuAccType uint8;
    BrokerIDByBank string;
    BankSecuAcc string;
    BankPwdFlag uint8;
    SecuPwdFlag uint8;
    OperNo string;
    RequestID int32;
    TID int32;
    TransferStatus uint8;
    ErrorID int32;
    ErrorMsg string;
}

func (m* CThostFtdcRspRepealField) Marshal(w io.Writer) {
    binary.Write(w, binary.BigEndian, &m.RepealTimeInterval)
    binary.Write(w, binary.BigEndian, &m.RepealedTimes)
    binary.Write(w, binary.BigEndian, &m.BankRepealFlag)
    binary.Write(w, binary.BigEndian, &m.BrokerRepealFlag)
    binary.Write(w, binary.BigEndian, &m.PlateRepealSerial)
    writeLen(w, m.BankRepealSerial, 13)
    binary.Write(w, binary.BigEndian, &m.FutureRepealSerial)
    writeLen(w, m.TradeCode, 7)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBranchID, 5)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.BrokerBranchID, 31)
    writeLen(w, m.TradeDate, 9)
    writeLen(w, m.TradeTime, 9)
    writeLen(w, m.BankSerial, 13)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.PlateSerial)
    binary.Write(w, binary.BigEndian, &m.LastFragment)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.CustomerName, 51)
    binary.Write(w, binary.BigEndian, &m.IdCardType)
    writeLen(w, m.IdentifiedCardNo, 51)
    binary.Write(w, binary.BigEndian, &m.CustType)
    writeLen(w, m.BankAccount, 41)
    writeLen(w, m.BankPassWord, 41)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.Password, 41)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    binary.Write(w, binary.BigEndian, &m.FutureSerial)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.VerifyCertNoFlag)
    writeLen(w, m.CurrencyID, 4)
    binary.Write(w, binary.BigEndian, &m.TradeAmount)
    binary.Write(w, binary.BigEndian, &m.FutureFetchAmount)
    binary.Write(w, binary.BigEndian, &m.FeePayFlag)
    binary.Write(w, binary.BigEndian, &m.CustFee)
    binary.Write(w, binary.BigEndian, &m.BrokerFee)
    writeLen(w, m.Message, 129)
    writeLen(w, m.Digest, 36)
    binary.Write(w, binary.BigEndian, &m.BankAccType)
    writeLen(w, m.DeviceID, 3)
    binary.Write(w, binary.BigEndian, &m.BankSecuAccType)
    writeLen(w, m.BrokerIDByBank, 33)
    writeLen(w, m.BankSecuAcc, 41)
    binary.Write(w, binary.BigEndian, &m.BankPwdFlag)
    binary.Write(w, binary.BigEndian, &m.SecuPwdFlag)
    writeLen(w, m.OperNo, 17)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    binary.Write(w, binary.BigEndian, &m.TID)
    binary.Write(w, binary.BigEndian, &m.TransferStatus)
    binary.Write(w, binary.BigEndian, &m.ErrorID)
    writeLen(w, m.ErrorMsg, 81)
}

func (m* CThostFtdcRspRepealField) Unmarshal(r io.Reader) {
    binary.Read(r, binary.BigEndian, &m.RepealTimeInterval)
    binary.Read(r, binary.BigEndian, &m.RepealedTimes)
    binary.Read(r, binary.BigEndian, &m.BankRepealFlag)
    binary.Read(r, binary.BigEndian, &m.BrokerRepealFlag)
    binary.Read(r, binary.BigEndian, &m.PlateRepealSerial)
    m.BankRepealSerial = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.FutureRepealSerial)
    m.TradeCode = readLen(r, 7)
    m.BankID = readLen(r, 4)
    m.BankBranchID = readLen(r, 5)
    m.BrokerID = readLen(r, 11)
    m.BrokerBranchID = readLen(r, 31)
    m.TradeDate = readLen(r, 9)
    m.TradeTime = readLen(r, 9)
    m.BankSerial = readLen(r, 13)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.PlateSerial)
    binary.Read(r, binary.BigEndian, &m.LastFragment)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.CustomerName = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.IdCardType)
    m.IdentifiedCardNo = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.CustType)
    m.BankAccount = readLen(r, 41)
    m.BankPassWord = readLen(r, 41)
    m.AccountID = readLen(r, 13)
    m.Password = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    binary.Read(r, binary.BigEndian, &m.FutureSerial)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.VerifyCertNoFlag)
    m.CurrencyID = readLen(r, 4)
    binary.Read(r, binary.BigEndian, &m.TradeAmount)
    binary.Read(r, binary.BigEndian, &m.FutureFetchAmount)
    binary.Read(r, binary.BigEndian, &m.FeePayFlag)
    binary.Read(r, binary.BigEndian, &m.CustFee)
    binary.Read(r, binary.BigEndian, &m.BrokerFee)
    m.Message = readLen(r, 129)
    m.Digest = readLen(r, 36)
    binary.Read(r, binary.BigEndian, &m.BankAccType)
    m.DeviceID = readLen(r, 3)
    binary.Read(r, binary.BigEndian, &m.BankSecuAccType)
    m.BrokerIDByBank = readLen(r, 33)
    m.BankSecuAcc = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.BankPwdFlag)
    binary.Read(r, binary.BigEndian, &m.SecuPwdFlag)
    m.OperNo = readLen(r, 17)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    binary.Read(r, binary.BigEndian, &m.TID)
    binary.Read(r, binary.BigEndian, &m.TransferStatus)
    binary.Read(r, binary.BigEndian, &m.ErrorID)
    m.ErrorMsg = readLen(r, 81)
}

type  CThostFtdcReqQueryAccountField struct {
    TradeCode string;
    BankID string;
    BankBranchID string;
    BrokerID string;
    BrokerBranchID string;
    TradeDate string;
    TradeTime string;
    BankSerial string;
    TradingDay string;
    PlateSerial int32;
    LastFragment uint8;
    SessionID int32;
    CustomerName string;
    IdCardType uint8;
    IdentifiedCardNo string;
    CustType uint8;
    BankAccount string;
    BankPassWord string;
    AccountID string;
    Password string;
    FutureSerial int32;
    InstallID int32;
    UserID string;
    VerifyCertNoFlag uint8;
    CurrencyID string;
    Digest string;
    BankAccType uint8;
    DeviceID string;
    BankSecuAccType uint8;
    BrokerIDByBank string;
    BankSecuAcc string;
    BankPwdFlag uint8;
    SecuPwdFlag uint8;
    OperNo string;
    RequestID int32;
    TID int32;
}

func (m* CThostFtdcReqQueryAccountField) Marshal(w io.Writer) {
    writeLen(w, m.TradeCode, 7)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBranchID, 5)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.BrokerBranchID, 31)
    writeLen(w, m.TradeDate, 9)
    writeLen(w, m.TradeTime, 9)
    writeLen(w, m.BankSerial, 13)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.PlateSerial)
    binary.Write(w, binary.BigEndian, &m.LastFragment)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.CustomerName, 51)
    binary.Write(w, binary.BigEndian, &m.IdCardType)
    writeLen(w, m.IdentifiedCardNo, 51)
    binary.Write(w, binary.BigEndian, &m.CustType)
    writeLen(w, m.BankAccount, 41)
    writeLen(w, m.BankPassWord, 41)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.Password, 41)
    binary.Write(w, binary.BigEndian, &m.FutureSerial)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.VerifyCertNoFlag)
    writeLen(w, m.CurrencyID, 4)
    writeLen(w, m.Digest, 36)
    binary.Write(w, binary.BigEndian, &m.BankAccType)
    writeLen(w, m.DeviceID, 3)
    binary.Write(w, binary.BigEndian, &m.BankSecuAccType)
    writeLen(w, m.BrokerIDByBank, 33)
    writeLen(w, m.BankSecuAcc, 41)
    binary.Write(w, binary.BigEndian, &m.BankPwdFlag)
    binary.Write(w, binary.BigEndian, &m.SecuPwdFlag)
    writeLen(w, m.OperNo, 17)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    binary.Write(w, binary.BigEndian, &m.TID)
}

func (m* CThostFtdcReqQueryAccountField) Unmarshal(r io.Reader) {
    m.TradeCode = readLen(r, 7)
    m.BankID = readLen(r, 4)
    m.BankBranchID = readLen(r, 5)
    m.BrokerID = readLen(r, 11)
    m.BrokerBranchID = readLen(r, 31)
    m.TradeDate = readLen(r, 9)
    m.TradeTime = readLen(r, 9)
    m.BankSerial = readLen(r, 13)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.PlateSerial)
    binary.Read(r, binary.BigEndian, &m.LastFragment)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.CustomerName = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.IdCardType)
    m.IdentifiedCardNo = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.CustType)
    m.BankAccount = readLen(r, 41)
    m.BankPassWord = readLen(r, 41)
    m.AccountID = readLen(r, 13)
    m.Password = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.FutureSerial)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.VerifyCertNoFlag)
    m.CurrencyID = readLen(r, 4)
    m.Digest = readLen(r, 36)
    binary.Read(r, binary.BigEndian, &m.BankAccType)
    m.DeviceID = readLen(r, 3)
    binary.Read(r, binary.BigEndian, &m.BankSecuAccType)
    m.BrokerIDByBank = readLen(r, 33)
    m.BankSecuAcc = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.BankPwdFlag)
    binary.Read(r, binary.BigEndian, &m.SecuPwdFlag)
    m.OperNo = readLen(r, 17)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    binary.Read(r, binary.BigEndian, &m.TID)
}

type  CThostFtdcRspQueryAccountField struct {
    TradeCode string;
    BankID string;
    BankBranchID string;
    BrokerID string;
    BrokerBranchID string;
    TradeDate string;
    TradeTime string;
    BankSerial string;
    TradingDay string;
    PlateSerial int32;
    LastFragment uint8;
    SessionID int32;
    CustomerName string;
    IdCardType uint8;
    IdentifiedCardNo string;
    CustType uint8;
    BankAccount string;
    BankPassWord string;
    AccountID string;
    Password string;
    FutureSerial int32;
    InstallID int32;
    UserID string;
    VerifyCertNoFlag uint8;
    CurrencyID string;
    Digest string;
    BankAccType uint8;
    DeviceID string;
    BankSecuAccType uint8;
    BrokerIDByBank string;
    BankSecuAcc string;
    BankPwdFlag uint8;
    SecuPwdFlag uint8;
    OperNo string;
    RequestID int32;
    TID int32;
    BankUseAmount float64;
    BankFetchAmount float64;
}

func (m* CThostFtdcRspQueryAccountField) Marshal(w io.Writer) {
    writeLen(w, m.TradeCode, 7)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBranchID, 5)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.BrokerBranchID, 31)
    writeLen(w, m.TradeDate, 9)
    writeLen(w, m.TradeTime, 9)
    writeLen(w, m.BankSerial, 13)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.PlateSerial)
    binary.Write(w, binary.BigEndian, &m.LastFragment)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.CustomerName, 51)
    binary.Write(w, binary.BigEndian, &m.IdCardType)
    writeLen(w, m.IdentifiedCardNo, 51)
    binary.Write(w, binary.BigEndian, &m.CustType)
    writeLen(w, m.BankAccount, 41)
    writeLen(w, m.BankPassWord, 41)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.Password, 41)
    binary.Write(w, binary.BigEndian, &m.FutureSerial)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.VerifyCertNoFlag)
    writeLen(w, m.CurrencyID, 4)
    writeLen(w, m.Digest, 36)
    binary.Write(w, binary.BigEndian, &m.BankAccType)
    writeLen(w, m.DeviceID, 3)
    binary.Write(w, binary.BigEndian, &m.BankSecuAccType)
    writeLen(w, m.BrokerIDByBank, 33)
    writeLen(w, m.BankSecuAcc, 41)
    binary.Write(w, binary.BigEndian, &m.BankPwdFlag)
    binary.Write(w, binary.BigEndian, &m.SecuPwdFlag)
    writeLen(w, m.OperNo, 17)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    binary.Write(w, binary.BigEndian, &m.TID)
    binary.Write(w, binary.BigEndian, &m.BankUseAmount)
    binary.Write(w, binary.BigEndian, &m.BankFetchAmount)
}

func (m* CThostFtdcRspQueryAccountField) Unmarshal(r io.Reader) {
    m.TradeCode = readLen(r, 7)
    m.BankID = readLen(r, 4)
    m.BankBranchID = readLen(r, 5)
    m.BrokerID = readLen(r, 11)
    m.BrokerBranchID = readLen(r, 31)
    m.TradeDate = readLen(r, 9)
    m.TradeTime = readLen(r, 9)
    m.BankSerial = readLen(r, 13)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.PlateSerial)
    binary.Read(r, binary.BigEndian, &m.LastFragment)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.CustomerName = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.IdCardType)
    m.IdentifiedCardNo = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.CustType)
    m.BankAccount = readLen(r, 41)
    m.BankPassWord = readLen(r, 41)
    m.AccountID = readLen(r, 13)
    m.Password = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.FutureSerial)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.VerifyCertNoFlag)
    m.CurrencyID = readLen(r, 4)
    m.Digest = readLen(r, 36)
    binary.Read(r, binary.BigEndian, &m.BankAccType)
    m.DeviceID = readLen(r, 3)
    binary.Read(r, binary.BigEndian, &m.BankSecuAccType)
    m.BrokerIDByBank = readLen(r, 33)
    m.BankSecuAcc = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.BankPwdFlag)
    binary.Read(r, binary.BigEndian, &m.SecuPwdFlag)
    m.OperNo = readLen(r, 17)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    binary.Read(r, binary.BigEndian, &m.TID)
    binary.Read(r, binary.BigEndian, &m.BankUseAmount)
    binary.Read(r, binary.BigEndian, &m.BankFetchAmount)
}

type  CThostFtdcFutureSignIOField struct {
    TradeCode string;
    BankID string;
    BankBranchID string;
    BrokerID string;
    BrokerBranchID string;
    TradeDate string;
    TradeTime string;
    BankSerial string;
    TradingDay string;
    PlateSerial int32;
    LastFragment uint8;
    SessionID int32;
    InstallID int32;
    UserID string;
    Digest string;
    CurrencyID string;
    DeviceID string;
    BrokerIDByBank string;
    OperNo string;
    RequestID int32;
    TID int32;
}

func (m* CThostFtdcFutureSignIOField) Marshal(w io.Writer) {
    writeLen(w, m.TradeCode, 7)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBranchID, 5)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.BrokerBranchID, 31)
    writeLen(w, m.TradeDate, 9)
    writeLen(w, m.TradeTime, 9)
    writeLen(w, m.BankSerial, 13)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.PlateSerial)
    binary.Write(w, binary.BigEndian, &m.LastFragment)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.Digest, 36)
    writeLen(w, m.CurrencyID, 4)
    writeLen(w, m.DeviceID, 3)
    writeLen(w, m.BrokerIDByBank, 33)
    writeLen(w, m.OperNo, 17)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    binary.Write(w, binary.BigEndian, &m.TID)
}

func (m* CThostFtdcFutureSignIOField) Unmarshal(r io.Reader) {
    m.TradeCode = readLen(r, 7)
    m.BankID = readLen(r, 4)
    m.BankBranchID = readLen(r, 5)
    m.BrokerID = readLen(r, 11)
    m.BrokerBranchID = readLen(r, 31)
    m.TradeDate = readLen(r, 9)
    m.TradeTime = readLen(r, 9)
    m.BankSerial = readLen(r, 13)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.PlateSerial)
    binary.Read(r, binary.BigEndian, &m.LastFragment)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    m.UserID = readLen(r, 16)
    m.Digest = readLen(r, 36)
    m.CurrencyID = readLen(r, 4)
    m.DeviceID = readLen(r, 3)
    m.BrokerIDByBank = readLen(r, 33)
    m.OperNo = readLen(r, 17)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    binary.Read(r, binary.BigEndian, &m.TID)
}

type  CThostFtdcRspFutureSignInField struct {
    TradeCode string;
    BankID string;
    BankBranchID string;
    BrokerID string;
    BrokerBranchID string;
    TradeDate string;
    TradeTime string;
    BankSerial string;
    TradingDay string;
    PlateSerial int32;
    LastFragment uint8;
    SessionID int32;
    InstallID int32;
    UserID string;
    Digest string;
    CurrencyID string;
    DeviceID string;
    BrokerIDByBank string;
    OperNo string;
    RequestID int32;
    TID int32;
    ErrorID int32;
    ErrorMsg string;
    PinKey string;
    MacKey string;
}

func (m* CThostFtdcRspFutureSignInField) Marshal(w io.Writer) {
    writeLen(w, m.TradeCode, 7)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBranchID, 5)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.BrokerBranchID, 31)
    writeLen(w, m.TradeDate, 9)
    writeLen(w, m.TradeTime, 9)
    writeLen(w, m.BankSerial, 13)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.PlateSerial)
    binary.Write(w, binary.BigEndian, &m.LastFragment)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.Digest, 36)
    writeLen(w, m.CurrencyID, 4)
    writeLen(w, m.DeviceID, 3)
    writeLen(w, m.BrokerIDByBank, 33)
    writeLen(w, m.OperNo, 17)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    binary.Write(w, binary.BigEndian, &m.TID)
    binary.Write(w, binary.BigEndian, &m.ErrorID)
    writeLen(w, m.ErrorMsg, 81)
    writeLen(w, m.PinKey, 129)
    writeLen(w, m.MacKey, 129)
}

func (m* CThostFtdcRspFutureSignInField) Unmarshal(r io.Reader) {
    m.TradeCode = readLen(r, 7)
    m.BankID = readLen(r, 4)
    m.BankBranchID = readLen(r, 5)
    m.BrokerID = readLen(r, 11)
    m.BrokerBranchID = readLen(r, 31)
    m.TradeDate = readLen(r, 9)
    m.TradeTime = readLen(r, 9)
    m.BankSerial = readLen(r, 13)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.PlateSerial)
    binary.Read(r, binary.BigEndian, &m.LastFragment)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    m.UserID = readLen(r, 16)
    m.Digest = readLen(r, 36)
    m.CurrencyID = readLen(r, 4)
    m.DeviceID = readLen(r, 3)
    m.BrokerIDByBank = readLen(r, 33)
    m.OperNo = readLen(r, 17)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    binary.Read(r, binary.BigEndian, &m.TID)
    binary.Read(r, binary.BigEndian, &m.ErrorID)
    m.ErrorMsg = readLen(r, 81)
    m.PinKey = readLen(r, 129)
    m.MacKey = readLen(r, 129)
}

type  CThostFtdcReqFutureSignOutField struct {
    TradeCode string;
    BankID string;
    BankBranchID string;
    BrokerID string;
    BrokerBranchID string;
    TradeDate string;
    TradeTime string;
    BankSerial string;
    TradingDay string;
    PlateSerial int32;
    LastFragment uint8;
    SessionID int32;
    InstallID int32;
    UserID string;
    Digest string;
    CurrencyID string;
    DeviceID string;
    BrokerIDByBank string;
    OperNo string;
    RequestID int32;
    TID int32;
}

func (m* CThostFtdcReqFutureSignOutField) Marshal(w io.Writer) {
    writeLen(w, m.TradeCode, 7)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBranchID, 5)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.BrokerBranchID, 31)
    writeLen(w, m.TradeDate, 9)
    writeLen(w, m.TradeTime, 9)
    writeLen(w, m.BankSerial, 13)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.PlateSerial)
    binary.Write(w, binary.BigEndian, &m.LastFragment)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.Digest, 36)
    writeLen(w, m.CurrencyID, 4)
    writeLen(w, m.DeviceID, 3)
    writeLen(w, m.BrokerIDByBank, 33)
    writeLen(w, m.OperNo, 17)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    binary.Write(w, binary.BigEndian, &m.TID)
}

func (m* CThostFtdcReqFutureSignOutField) Unmarshal(r io.Reader) {
    m.TradeCode = readLen(r, 7)
    m.BankID = readLen(r, 4)
    m.BankBranchID = readLen(r, 5)
    m.BrokerID = readLen(r, 11)
    m.BrokerBranchID = readLen(r, 31)
    m.TradeDate = readLen(r, 9)
    m.TradeTime = readLen(r, 9)
    m.BankSerial = readLen(r, 13)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.PlateSerial)
    binary.Read(r, binary.BigEndian, &m.LastFragment)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    m.UserID = readLen(r, 16)
    m.Digest = readLen(r, 36)
    m.CurrencyID = readLen(r, 4)
    m.DeviceID = readLen(r, 3)
    m.BrokerIDByBank = readLen(r, 33)
    m.OperNo = readLen(r, 17)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    binary.Read(r, binary.BigEndian, &m.TID)
}

type  CThostFtdcRspFutureSignOutField struct {
    TradeCode string;
    BankID string;
    BankBranchID string;
    BrokerID string;
    BrokerBranchID string;
    TradeDate string;
    TradeTime string;
    BankSerial string;
    TradingDay string;
    PlateSerial int32;
    LastFragment uint8;
    SessionID int32;
    InstallID int32;
    UserID string;
    Digest string;
    CurrencyID string;
    DeviceID string;
    BrokerIDByBank string;
    OperNo string;
    RequestID int32;
    TID int32;
    ErrorID int32;
    ErrorMsg string;
}

func (m* CThostFtdcRspFutureSignOutField) Marshal(w io.Writer) {
    writeLen(w, m.TradeCode, 7)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBranchID, 5)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.BrokerBranchID, 31)
    writeLen(w, m.TradeDate, 9)
    writeLen(w, m.TradeTime, 9)
    writeLen(w, m.BankSerial, 13)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.PlateSerial)
    binary.Write(w, binary.BigEndian, &m.LastFragment)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.Digest, 36)
    writeLen(w, m.CurrencyID, 4)
    writeLen(w, m.DeviceID, 3)
    writeLen(w, m.BrokerIDByBank, 33)
    writeLen(w, m.OperNo, 17)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    binary.Write(w, binary.BigEndian, &m.TID)
    binary.Write(w, binary.BigEndian, &m.ErrorID)
    writeLen(w, m.ErrorMsg, 81)
}

func (m* CThostFtdcRspFutureSignOutField) Unmarshal(r io.Reader) {
    m.TradeCode = readLen(r, 7)
    m.BankID = readLen(r, 4)
    m.BankBranchID = readLen(r, 5)
    m.BrokerID = readLen(r, 11)
    m.BrokerBranchID = readLen(r, 31)
    m.TradeDate = readLen(r, 9)
    m.TradeTime = readLen(r, 9)
    m.BankSerial = readLen(r, 13)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.PlateSerial)
    binary.Read(r, binary.BigEndian, &m.LastFragment)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    m.UserID = readLen(r, 16)
    m.Digest = readLen(r, 36)
    m.CurrencyID = readLen(r, 4)
    m.DeviceID = readLen(r, 3)
    m.BrokerIDByBank = readLen(r, 33)
    m.OperNo = readLen(r, 17)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    binary.Read(r, binary.BigEndian, &m.TID)
    binary.Read(r, binary.BigEndian, &m.ErrorID)
    m.ErrorMsg = readLen(r, 81)
}

type  CThostFtdcReqQueryTradeResultBySerialField struct {
    TradeCode string;
    BankID string;
    BankBranchID string;
    BrokerID string;
    BrokerBranchID string;
    TradeDate string;
    TradeTime string;
    BankSerial string;
    TradingDay string;
    PlateSerial int32;
    LastFragment uint8;
    SessionID int32;
    Reference int32;
    RefrenceIssureType uint8;
    RefrenceIssure string;
    CustomerName string;
    IdCardType uint8;
    IdentifiedCardNo string;
    CustType uint8;
    BankAccount string;
    BankPassWord string;
    AccountID string;
    Password string;
    CurrencyID string;
    TradeAmount float64;
    Digest string;
}

func (m* CThostFtdcReqQueryTradeResultBySerialField) Marshal(w io.Writer) {
    writeLen(w, m.TradeCode, 7)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBranchID, 5)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.BrokerBranchID, 31)
    writeLen(w, m.TradeDate, 9)
    writeLen(w, m.TradeTime, 9)
    writeLen(w, m.BankSerial, 13)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.PlateSerial)
    binary.Write(w, binary.BigEndian, &m.LastFragment)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    binary.Write(w, binary.BigEndian, &m.Reference)
    binary.Write(w, binary.BigEndian, &m.RefrenceIssureType)
    writeLen(w, m.RefrenceIssure, 36)
    writeLen(w, m.CustomerName, 51)
    binary.Write(w, binary.BigEndian, &m.IdCardType)
    writeLen(w, m.IdentifiedCardNo, 51)
    binary.Write(w, binary.BigEndian, &m.CustType)
    writeLen(w, m.BankAccount, 41)
    writeLen(w, m.BankPassWord, 41)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.Password, 41)
    writeLen(w, m.CurrencyID, 4)
    binary.Write(w, binary.BigEndian, &m.TradeAmount)
    writeLen(w, m.Digest, 36)
}

func (m* CThostFtdcReqQueryTradeResultBySerialField) Unmarshal(r io.Reader) {
    m.TradeCode = readLen(r, 7)
    m.BankID = readLen(r, 4)
    m.BankBranchID = readLen(r, 5)
    m.BrokerID = readLen(r, 11)
    m.BrokerBranchID = readLen(r, 31)
    m.TradeDate = readLen(r, 9)
    m.TradeTime = readLen(r, 9)
    m.BankSerial = readLen(r, 13)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.PlateSerial)
    binary.Read(r, binary.BigEndian, &m.LastFragment)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    binary.Read(r, binary.BigEndian, &m.Reference)
    binary.Read(r, binary.BigEndian, &m.RefrenceIssureType)
    m.RefrenceIssure = readLen(r, 36)
    m.CustomerName = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.IdCardType)
    m.IdentifiedCardNo = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.CustType)
    m.BankAccount = readLen(r, 41)
    m.BankPassWord = readLen(r, 41)
    m.AccountID = readLen(r, 13)
    m.Password = readLen(r, 41)
    m.CurrencyID = readLen(r, 4)
    binary.Read(r, binary.BigEndian, &m.TradeAmount)
    m.Digest = readLen(r, 36)
}

type  CThostFtdcRspQueryTradeResultBySerialField struct {
    TradeCode string;
    BankID string;
    BankBranchID string;
    BrokerID string;
    BrokerBranchID string;
    TradeDate string;
    TradeTime string;
    BankSerial string;
    TradingDay string;
    PlateSerial int32;
    LastFragment uint8;
    SessionID int32;
    ErrorID int32;
    ErrorMsg string;
    Reference int32;
    RefrenceIssureType uint8;
    RefrenceIssure string;
    OriginReturnCode string;
    OriginDescrInfoForReturnCode string;
    BankAccount string;
    BankPassWord string;
    AccountID string;
    Password string;
    CurrencyID string;
    TradeAmount float64;
    Digest string;
}

func (m* CThostFtdcRspQueryTradeResultBySerialField) Marshal(w io.Writer) {
    writeLen(w, m.TradeCode, 7)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBranchID, 5)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.BrokerBranchID, 31)
    writeLen(w, m.TradeDate, 9)
    writeLen(w, m.TradeTime, 9)
    writeLen(w, m.BankSerial, 13)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.PlateSerial)
    binary.Write(w, binary.BigEndian, &m.LastFragment)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    binary.Write(w, binary.BigEndian, &m.ErrorID)
    writeLen(w, m.ErrorMsg, 81)
    binary.Write(w, binary.BigEndian, &m.Reference)
    binary.Write(w, binary.BigEndian, &m.RefrenceIssureType)
    writeLen(w, m.RefrenceIssure, 36)
    writeLen(w, m.OriginReturnCode, 7)
    writeLen(w, m.OriginDescrInfoForReturnCode, 129)
    writeLen(w, m.BankAccount, 41)
    writeLen(w, m.BankPassWord, 41)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.Password, 41)
    writeLen(w, m.CurrencyID, 4)
    binary.Write(w, binary.BigEndian, &m.TradeAmount)
    writeLen(w, m.Digest, 36)
}

func (m* CThostFtdcRspQueryTradeResultBySerialField) Unmarshal(r io.Reader) {
    m.TradeCode = readLen(r, 7)
    m.BankID = readLen(r, 4)
    m.BankBranchID = readLen(r, 5)
    m.BrokerID = readLen(r, 11)
    m.BrokerBranchID = readLen(r, 31)
    m.TradeDate = readLen(r, 9)
    m.TradeTime = readLen(r, 9)
    m.BankSerial = readLen(r, 13)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.PlateSerial)
    binary.Read(r, binary.BigEndian, &m.LastFragment)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    binary.Read(r, binary.BigEndian, &m.ErrorID)
    m.ErrorMsg = readLen(r, 81)
    binary.Read(r, binary.BigEndian, &m.Reference)
    binary.Read(r, binary.BigEndian, &m.RefrenceIssureType)
    m.RefrenceIssure = readLen(r, 36)
    m.OriginReturnCode = readLen(r, 7)
    m.OriginDescrInfoForReturnCode = readLen(r, 129)
    m.BankAccount = readLen(r, 41)
    m.BankPassWord = readLen(r, 41)
    m.AccountID = readLen(r, 13)
    m.Password = readLen(r, 41)
    m.CurrencyID = readLen(r, 4)
    binary.Read(r, binary.BigEndian, &m.TradeAmount)
    m.Digest = readLen(r, 36)
}

type  CThostFtdcReqDayEndFileReadyField struct {
    TradeCode string;
    BankID string;
    BankBranchID string;
    BrokerID string;
    BrokerBranchID string;
    TradeDate string;
    TradeTime string;
    BankSerial string;
    TradingDay string;
    PlateSerial int32;
    LastFragment uint8;
    SessionID int32;
    FileBusinessCode uint8;
    InstallID int32;
    UserID string;
    Digest string;
    CurrencyID string;
    DeviceID string;
    BrokerIDByBank string;
    OperNo string;
    RequestID int32;
    TID int32;
}

func (m* CThostFtdcReqDayEndFileReadyField) Marshal(w io.Writer) {
    writeLen(w, m.TradeCode, 7)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBranchID, 5)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.BrokerBranchID, 31)
    writeLen(w, m.TradeDate, 9)
    writeLen(w, m.TradeTime, 9)
    writeLen(w, m.BankSerial, 13)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.PlateSerial)
    binary.Write(w, binary.BigEndian, &m.LastFragment)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    binary.Write(w, binary.BigEndian, &m.FileBusinessCode)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.Digest, 36)
    writeLen(w, m.CurrencyID, 4)
    writeLen(w, m.DeviceID, 3)
    writeLen(w, m.BrokerIDByBank, 33)
    writeLen(w, m.OperNo, 17)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    binary.Write(w, binary.BigEndian, &m.TID)
}

func (m* CThostFtdcReqDayEndFileReadyField) Unmarshal(r io.Reader) {
    m.TradeCode = readLen(r, 7)
    m.BankID = readLen(r, 4)
    m.BankBranchID = readLen(r, 5)
    m.BrokerID = readLen(r, 11)
    m.BrokerBranchID = readLen(r, 31)
    m.TradeDate = readLen(r, 9)
    m.TradeTime = readLen(r, 9)
    m.BankSerial = readLen(r, 13)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.PlateSerial)
    binary.Read(r, binary.BigEndian, &m.LastFragment)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    binary.Read(r, binary.BigEndian, &m.FileBusinessCode)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    m.UserID = readLen(r, 16)
    m.Digest = readLen(r, 36)
    m.CurrencyID = readLen(r, 4)
    m.DeviceID = readLen(r, 3)
    m.BrokerIDByBank = readLen(r, 33)
    m.OperNo = readLen(r, 17)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    binary.Read(r, binary.BigEndian, &m.TID)
}

type  CThostFtdcRspDayEndFileReadyField struct {
    TradeCode string;
    BankID string;
    BankBranchID string;
    BrokerID string;
    BrokerBranchID string;
    TradeDate string;
    TradeTime string;
    BankSerial string;
    TradingDay string;
    PlateSerial int32;
    LastFragment uint8;
    SessionID int32;
    FileBusinessCode uint8;
    InstallID int32;
    UserID string;
    Digest string;
    CurrencyID string;
    DeviceID string;
    BrokerIDByBank string;
    OperNo string;
    RequestID int32;
    TID int32;
    ErrorID int32;
    ErrorMsg string;
}

func (m* CThostFtdcRspDayEndFileReadyField) Marshal(w io.Writer) {
    writeLen(w, m.TradeCode, 7)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBranchID, 5)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.BrokerBranchID, 31)
    writeLen(w, m.TradeDate, 9)
    writeLen(w, m.TradeTime, 9)
    writeLen(w, m.BankSerial, 13)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.PlateSerial)
    binary.Write(w, binary.BigEndian, &m.LastFragment)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    binary.Write(w, binary.BigEndian, &m.FileBusinessCode)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.Digest, 36)
    writeLen(w, m.CurrencyID, 4)
    writeLen(w, m.DeviceID, 3)
    writeLen(w, m.BrokerIDByBank, 33)
    writeLen(w, m.OperNo, 17)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    binary.Write(w, binary.BigEndian, &m.TID)
    binary.Write(w, binary.BigEndian, &m.ErrorID)
    writeLen(w, m.ErrorMsg, 81)
}

func (m* CThostFtdcRspDayEndFileReadyField) Unmarshal(r io.Reader) {
    m.TradeCode = readLen(r, 7)
    m.BankID = readLen(r, 4)
    m.BankBranchID = readLen(r, 5)
    m.BrokerID = readLen(r, 11)
    m.BrokerBranchID = readLen(r, 31)
    m.TradeDate = readLen(r, 9)
    m.TradeTime = readLen(r, 9)
    m.BankSerial = readLen(r, 13)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.PlateSerial)
    binary.Read(r, binary.BigEndian, &m.LastFragment)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    binary.Read(r, binary.BigEndian, &m.FileBusinessCode)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    m.UserID = readLen(r, 16)
    m.Digest = readLen(r, 36)
    m.CurrencyID = readLen(r, 4)
    m.DeviceID = readLen(r, 3)
    m.BrokerIDByBank = readLen(r, 33)
    m.OperNo = readLen(r, 17)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    binary.Read(r, binary.BigEndian, &m.TID)
    binary.Read(r, binary.BigEndian, &m.ErrorID)
    m.ErrorMsg = readLen(r, 81)
}

type  CThostFtdcDayEndFileReadyField struct {
    TradeCode string;
    BankID string;
    BankBranchID string;
    BrokerID string;
    BrokerBranchID string;
    TradeDate string;
    TradeTime string;
    BankSerial string;
    TradingDay string;
    PlateSerial int32;
    LastFragment uint8;
    SessionID int32;
    FileBusinessCode uint8;
    InstallID int32;
    UserID string;
    Digest string;
    CurrencyID string;
    DeviceID string;
    BrokerIDByBank string;
    OperNo string;
    RequestID int32;
    TID int32;
    ErrorID int32;
    ErrorMsg string;
}

func (m* CThostFtdcDayEndFileReadyField) Marshal(w io.Writer) {
    writeLen(w, m.TradeCode, 7)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBranchID, 5)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.BrokerBranchID, 31)
    writeLen(w, m.TradeDate, 9)
    writeLen(w, m.TradeTime, 9)
    writeLen(w, m.BankSerial, 13)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.PlateSerial)
    binary.Write(w, binary.BigEndian, &m.LastFragment)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    binary.Write(w, binary.BigEndian, &m.FileBusinessCode)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.Digest, 36)
    writeLen(w, m.CurrencyID, 4)
    writeLen(w, m.DeviceID, 3)
    writeLen(w, m.BrokerIDByBank, 33)
    writeLen(w, m.OperNo, 17)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    binary.Write(w, binary.BigEndian, &m.TID)
    binary.Write(w, binary.BigEndian, &m.ErrorID)
    writeLen(w, m.ErrorMsg, 81)
}

func (m* CThostFtdcDayEndFileReadyField) Unmarshal(r io.Reader) {
    m.TradeCode = readLen(r, 7)
    m.BankID = readLen(r, 4)
    m.BankBranchID = readLen(r, 5)
    m.BrokerID = readLen(r, 11)
    m.BrokerBranchID = readLen(r, 31)
    m.TradeDate = readLen(r, 9)
    m.TradeTime = readLen(r, 9)
    m.BankSerial = readLen(r, 13)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.PlateSerial)
    binary.Read(r, binary.BigEndian, &m.LastFragment)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    binary.Read(r, binary.BigEndian, &m.FileBusinessCode)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    m.UserID = readLen(r, 16)
    m.Digest = readLen(r, 36)
    m.CurrencyID = readLen(r, 4)
    m.DeviceID = readLen(r, 3)
    m.BrokerIDByBank = readLen(r, 33)
    m.OperNo = readLen(r, 17)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    binary.Read(r, binary.BigEndian, &m.TID)
    binary.Read(r, binary.BigEndian, &m.ErrorID)
    m.ErrorMsg = readLen(r, 81)
}

type  CThostFtdcReturnResultField struct {
    ReturnCode string;
    DescrInfoForReturnCode string;
}

func (m* CThostFtdcReturnResultField) Marshal(w io.Writer) {
    writeLen(w, m.ReturnCode, 7)
    writeLen(w, m.DescrInfoForReturnCode, 129)
}

func (m* CThostFtdcReturnResultField) Unmarshal(r io.Reader) {
    m.ReturnCode = readLen(r, 7)
    m.DescrInfoForReturnCode = readLen(r, 129)
}

type  CThostFtdcVerifyFuturePasswordField struct {
    TradeCode string;
    BankID string;
    BankBranchID string;
    BrokerID string;
    BrokerBranchID string;
    TradeDate string;
    TradeTime string;
    BankSerial string;
    TradingDay string;
    PlateSerial int32;
    LastFragment uint8;
    SessionID int32;
    AccountID string;
    Password string;
    BankAccount string;
    BankPassWord string;
    InstallID int32;
    TID int32;
    CurrencyID string;
}

func (m* CThostFtdcVerifyFuturePasswordField) Marshal(w io.Writer) {
    writeLen(w, m.TradeCode, 7)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBranchID, 5)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.BrokerBranchID, 31)
    writeLen(w, m.TradeDate, 9)
    writeLen(w, m.TradeTime, 9)
    writeLen(w, m.BankSerial, 13)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.PlateSerial)
    binary.Write(w, binary.BigEndian, &m.LastFragment)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.Password, 41)
    writeLen(w, m.BankAccount, 41)
    writeLen(w, m.BankPassWord, 41)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    binary.Write(w, binary.BigEndian, &m.TID)
    writeLen(w, m.CurrencyID, 4)
}

func (m* CThostFtdcVerifyFuturePasswordField) Unmarshal(r io.Reader) {
    m.TradeCode = readLen(r, 7)
    m.BankID = readLen(r, 4)
    m.BankBranchID = readLen(r, 5)
    m.BrokerID = readLen(r, 11)
    m.BrokerBranchID = readLen(r, 31)
    m.TradeDate = readLen(r, 9)
    m.TradeTime = readLen(r, 9)
    m.BankSerial = readLen(r, 13)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.PlateSerial)
    binary.Read(r, binary.BigEndian, &m.LastFragment)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.AccountID = readLen(r, 13)
    m.Password = readLen(r, 41)
    m.BankAccount = readLen(r, 41)
    m.BankPassWord = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    binary.Read(r, binary.BigEndian, &m.TID)
    m.CurrencyID = readLen(r, 4)
}

type  CThostFtdcVerifyCustInfoField struct {
    CustomerName string;
    IdCardType uint8;
    IdentifiedCardNo string;
    CustType uint8;
}

func (m* CThostFtdcVerifyCustInfoField) Marshal(w io.Writer) {
    writeLen(w, m.CustomerName, 51)
    binary.Write(w, binary.BigEndian, &m.IdCardType)
    writeLen(w, m.IdentifiedCardNo, 51)
    binary.Write(w, binary.BigEndian, &m.CustType)
}

func (m* CThostFtdcVerifyCustInfoField) Unmarshal(r io.Reader) {
    m.CustomerName = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.IdCardType)
    m.IdentifiedCardNo = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.CustType)
}

type  CThostFtdcVerifyFuturePasswordAndCustInfoField struct {
    CustomerName string;
    IdCardType uint8;
    IdentifiedCardNo string;
    CustType uint8;
    AccountID string;
    Password string;
    CurrencyID string;
}

func (m* CThostFtdcVerifyFuturePasswordAndCustInfoField) Marshal(w io.Writer) {
    writeLen(w, m.CustomerName, 51)
    binary.Write(w, binary.BigEndian, &m.IdCardType)
    writeLen(w, m.IdentifiedCardNo, 51)
    binary.Write(w, binary.BigEndian, &m.CustType)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.Password, 41)
    writeLen(w, m.CurrencyID, 4)
}

func (m* CThostFtdcVerifyFuturePasswordAndCustInfoField) Unmarshal(r io.Reader) {
    m.CustomerName = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.IdCardType)
    m.IdentifiedCardNo = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.CustType)
    m.AccountID = readLen(r, 13)
    m.Password = readLen(r, 41)
    m.CurrencyID = readLen(r, 4)
}

type  CThostFtdcDepositResultInformField struct {
    DepositSeqNo string;
    BrokerID string;
    InvestorID string;
    Deposit float64;
    RequestID int32;
    ReturnCode string;
    DescrInfoForReturnCode string;
}

func (m* CThostFtdcDepositResultInformField) Marshal(w io.Writer) {
    writeLen(w, m.DepositSeqNo, 15)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.Deposit)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    writeLen(w, m.ReturnCode, 7)
    writeLen(w, m.DescrInfoForReturnCode, 129)
}

func (m* CThostFtdcDepositResultInformField) Unmarshal(r io.Reader) {
    m.DepositSeqNo = readLen(r, 15)
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.Deposit)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    m.ReturnCode = readLen(r, 7)
    m.DescrInfoForReturnCode = readLen(r, 129)
}

type  CThostFtdcReqSyncKeyField struct {
    TradeCode string;
    BankID string;
    BankBranchID string;
    BrokerID string;
    BrokerBranchID string;
    TradeDate string;
    TradeTime string;
    BankSerial string;
    TradingDay string;
    PlateSerial int32;
    LastFragment uint8;
    SessionID int32;
    InstallID int32;
    UserID string;
    Message string;
    DeviceID string;
    BrokerIDByBank string;
    OperNo string;
    RequestID int32;
    TID int32;
}

func (m* CThostFtdcReqSyncKeyField) Marshal(w io.Writer) {
    writeLen(w, m.TradeCode, 7)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBranchID, 5)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.BrokerBranchID, 31)
    writeLen(w, m.TradeDate, 9)
    writeLen(w, m.TradeTime, 9)
    writeLen(w, m.BankSerial, 13)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.PlateSerial)
    binary.Write(w, binary.BigEndian, &m.LastFragment)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.Message, 129)
    writeLen(w, m.DeviceID, 3)
    writeLen(w, m.BrokerIDByBank, 33)
    writeLen(w, m.OperNo, 17)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    binary.Write(w, binary.BigEndian, &m.TID)
}

func (m* CThostFtdcReqSyncKeyField) Unmarshal(r io.Reader) {
    m.TradeCode = readLen(r, 7)
    m.BankID = readLen(r, 4)
    m.BankBranchID = readLen(r, 5)
    m.BrokerID = readLen(r, 11)
    m.BrokerBranchID = readLen(r, 31)
    m.TradeDate = readLen(r, 9)
    m.TradeTime = readLen(r, 9)
    m.BankSerial = readLen(r, 13)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.PlateSerial)
    binary.Read(r, binary.BigEndian, &m.LastFragment)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    m.UserID = readLen(r, 16)
    m.Message = readLen(r, 129)
    m.DeviceID = readLen(r, 3)
    m.BrokerIDByBank = readLen(r, 33)
    m.OperNo = readLen(r, 17)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    binary.Read(r, binary.BigEndian, &m.TID)
}

type  CThostFtdcRspSyncKeyField struct {
    TradeCode string;
    BankID string;
    BankBranchID string;
    BrokerID string;
    BrokerBranchID string;
    TradeDate string;
    TradeTime string;
    BankSerial string;
    TradingDay string;
    PlateSerial int32;
    LastFragment uint8;
    SessionID int32;
    InstallID int32;
    UserID string;
    Message string;
    DeviceID string;
    BrokerIDByBank string;
    OperNo string;
    RequestID int32;
    TID int32;
    ErrorID int32;
    ErrorMsg string;
}

func (m* CThostFtdcRspSyncKeyField) Marshal(w io.Writer) {
    writeLen(w, m.TradeCode, 7)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBranchID, 5)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.BrokerBranchID, 31)
    writeLen(w, m.TradeDate, 9)
    writeLen(w, m.TradeTime, 9)
    writeLen(w, m.BankSerial, 13)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.PlateSerial)
    binary.Write(w, binary.BigEndian, &m.LastFragment)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.Message, 129)
    writeLen(w, m.DeviceID, 3)
    writeLen(w, m.BrokerIDByBank, 33)
    writeLen(w, m.OperNo, 17)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    binary.Write(w, binary.BigEndian, &m.TID)
    binary.Write(w, binary.BigEndian, &m.ErrorID)
    writeLen(w, m.ErrorMsg, 81)
}

func (m* CThostFtdcRspSyncKeyField) Unmarshal(r io.Reader) {
    m.TradeCode = readLen(r, 7)
    m.BankID = readLen(r, 4)
    m.BankBranchID = readLen(r, 5)
    m.BrokerID = readLen(r, 11)
    m.BrokerBranchID = readLen(r, 31)
    m.TradeDate = readLen(r, 9)
    m.TradeTime = readLen(r, 9)
    m.BankSerial = readLen(r, 13)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.PlateSerial)
    binary.Read(r, binary.BigEndian, &m.LastFragment)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    m.UserID = readLen(r, 16)
    m.Message = readLen(r, 129)
    m.DeviceID = readLen(r, 3)
    m.BrokerIDByBank = readLen(r, 33)
    m.OperNo = readLen(r, 17)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    binary.Read(r, binary.BigEndian, &m.TID)
    binary.Read(r, binary.BigEndian, &m.ErrorID)
    m.ErrorMsg = readLen(r, 81)
}

type  CThostFtdcNotifyQueryAccountField struct {
    TradeCode string;
    BankID string;
    BankBranchID string;
    BrokerID string;
    BrokerBranchID string;
    TradeDate string;
    TradeTime string;
    BankSerial string;
    TradingDay string;
    PlateSerial int32;
    LastFragment uint8;
    SessionID int32;
    CustomerName string;
    IdCardType uint8;
    IdentifiedCardNo string;
    CustType uint8;
    BankAccount string;
    BankPassWord string;
    AccountID string;
    Password string;
    FutureSerial int32;
    InstallID int32;
    UserID string;
    VerifyCertNoFlag uint8;
    CurrencyID string;
    Digest string;
    BankAccType uint8;
    DeviceID string;
    BankSecuAccType uint8;
    BrokerIDByBank string;
    BankSecuAcc string;
    BankPwdFlag uint8;
    SecuPwdFlag uint8;
    OperNo string;
    RequestID int32;
    TID int32;
    BankUseAmount float64;
    BankFetchAmount float64;
    ErrorID int32;
    ErrorMsg string;
}

func (m* CThostFtdcNotifyQueryAccountField) Marshal(w io.Writer) {
    writeLen(w, m.TradeCode, 7)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBranchID, 5)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.BrokerBranchID, 31)
    writeLen(w, m.TradeDate, 9)
    writeLen(w, m.TradeTime, 9)
    writeLen(w, m.BankSerial, 13)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.PlateSerial)
    binary.Write(w, binary.BigEndian, &m.LastFragment)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.CustomerName, 51)
    binary.Write(w, binary.BigEndian, &m.IdCardType)
    writeLen(w, m.IdentifiedCardNo, 51)
    binary.Write(w, binary.BigEndian, &m.CustType)
    writeLen(w, m.BankAccount, 41)
    writeLen(w, m.BankPassWord, 41)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.Password, 41)
    binary.Write(w, binary.BigEndian, &m.FutureSerial)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.VerifyCertNoFlag)
    writeLen(w, m.CurrencyID, 4)
    writeLen(w, m.Digest, 36)
    binary.Write(w, binary.BigEndian, &m.BankAccType)
    writeLen(w, m.DeviceID, 3)
    binary.Write(w, binary.BigEndian, &m.BankSecuAccType)
    writeLen(w, m.BrokerIDByBank, 33)
    writeLen(w, m.BankSecuAcc, 41)
    binary.Write(w, binary.BigEndian, &m.BankPwdFlag)
    binary.Write(w, binary.BigEndian, &m.SecuPwdFlag)
    writeLen(w, m.OperNo, 17)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    binary.Write(w, binary.BigEndian, &m.TID)
    binary.Write(w, binary.BigEndian, &m.BankUseAmount)
    binary.Write(w, binary.BigEndian, &m.BankFetchAmount)
    binary.Write(w, binary.BigEndian, &m.ErrorID)
    writeLen(w, m.ErrorMsg, 81)
}

func (m* CThostFtdcNotifyQueryAccountField) Unmarshal(r io.Reader) {
    m.TradeCode = readLen(r, 7)
    m.BankID = readLen(r, 4)
    m.BankBranchID = readLen(r, 5)
    m.BrokerID = readLen(r, 11)
    m.BrokerBranchID = readLen(r, 31)
    m.TradeDate = readLen(r, 9)
    m.TradeTime = readLen(r, 9)
    m.BankSerial = readLen(r, 13)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.PlateSerial)
    binary.Read(r, binary.BigEndian, &m.LastFragment)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.CustomerName = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.IdCardType)
    m.IdentifiedCardNo = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.CustType)
    m.BankAccount = readLen(r, 41)
    m.BankPassWord = readLen(r, 41)
    m.AccountID = readLen(r, 13)
    m.Password = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.FutureSerial)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.VerifyCertNoFlag)
    m.CurrencyID = readLen(r, 4)
    m.Digest = readLen(r, 36)
    binary.Read(r, binary.BigEndian, &m.BankAccType)
    m.DeviceID = readLen(r, 3)
    binary.Read(r, binary.BigEndian, &m.BankSecuAccType)
    m.BrokerIDByBank = readLen(r, 33)
    m.BankSecuAcc = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.BankPwdFlag)
    binary.Read(r, binary.BigEndian, &m.SecuPwdFlag)
    m.OperNo = readLen(r, 17)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    binary.Read(r, binary.BigEndian, &m.TID)
    binary.Read(r, binary.BigEndian, &m.BankUseAmount)
    binary.Read(r, binary.BigEndian, &m.BankFetchAmount)
    binary.Read(r, binary.BigEndian, &m.ErrorID)
    m.ErrorMsg = readLen(r, 81)
}

type  CThostFtdcTransferSerialField struct {
    PlateSerial int32;
    TradeDate string;
    TradingDay string;
    TradeTime string;
    TradeCode string;
    SessionID int32;
    BankID string;
    BankBranchID string;
    BankAccType uint8;
    BankAccount string;
    BankSerial string;
    BrokerID string;
    BrokerBranchID string;
    FutureAccType uint8;
    AccountID string;
    InvestorID string;
    FutureSerial int32;
    IdCardType uint8;
    IdentifiedCardNo string;
    CurrencyID string;
    TradeAmount float64;
    CustFee float64;
    BrokerFee float64;
    AvailabilityFlag uint8;
    OperatorCode string;
    BankNewAccount string;
    ErrorID int32;
    ErrorMsg string;
}

func (m* CThostFtdcTransferSerialField) Marshal(w io.Writer) {
    binary.Write(w, binary.BigEndian, &m.PlateSerial)
    writeLen(w, m.TradeDate, 9)
    writeLen(w, m.TradingDay, 9)
    writeLen(w, m.TradeTime, 9)
    writeLen(w, m.TradeCode, 7)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBranchID, 5)
    binary.Write(w, binary.BigEndian, &m.BankAccType)
    writeLen(w, m.BankAccount, 41)
    writeLen(w, m.BankSerial, 13)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.BrokerBranchID, 31)
    binary.Write(w, binary.BigEndian, &m.FutureAccType)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.FutureSerial)
    binary.Write(w, binary.BigEndian, &m.IdCardType)
    writeLen(w, m.IdentifiedCardNo, 51)
    writeLen(w, m.CurrencyID, 4)
    binary.Write(w, binary.BigEndian, &m.TradeAmount)
    binary.Write(w, binary.BigEndian, &m.CustFee)
    binary.Write(w, binary.BigEndian, &m.BrokerFee)
    binary.Write(w, binary.BigEndian, &m.AvailabilityFlag)
    writeLen(w, m.OperatorCode, 17)
    writeLen(w, m.BankNewAccount, 41)
    binary.Write(w, binary.BigEndian, &m.ErrorID)
    writeLen(w, m.ErrorMsg, 81)
}

func (m* CThostFtdcTransferSerialField) Unmarshal(r io.Reader) {
    binary.Read(r, binary.BigEndian, &m.PlateSerial)
    m.TradeDate = readLen(r, 9)
    m.TradingDay = readLen(r, 9)
    m.TradeTime = readLen(r, 9)
    m.TradeCode = readLen(r, 7)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.BankID = readLen(r, 4)
    m.BankBranchID = readLen(r, 5)
    binary.Read(r, binary.BigEndian, &m.BankAccType)
    m.BankAccount = readLen(r, 41)
    m.BankSerial = readLen(r, 13)
    m.BrokerID = readLen(r, 11)
    m.BrokerBranchID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.FutureAccType)
    m.AccountID = readLen(r, 13)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.FutureSerial)
    binary.Read(r, binary.BigEndian, &m.IdCardType)
    m.IdentifiedCardNo = readLen(r, 51)
    m.CurrencyID = readLen(r, 4)
    binary.Read(r, binary.BigEndian, &m.TradeAmount)
    binary.Read(r, binary.BigEndian, &m.CustFee)
    binary.Read(r, binary.BigEndian, &m.BrokerFee)
    binary.Read(r, binary.BigEndian, &m.AvailabilityFlag)
    m.OperatorCode = readLen(r, 17)
    m.BankNewAccount = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.ErrorID)
    m.ErrorMsg = readLen(r, 81)
}

type  CThostFtdcQryTransferSerialField struct {
    BrokerID string;
    AccountID string;
    BankID string;
    CurrencyID string;
}

func (m* CThostFtdcQryTransferSerialField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.CurrencyID, 4)
}

func (m* CThostFtdcQryTransferSerialField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.AccountID = readLen(r, 13)
    m.BankID = readLen(r, 4)
    m.CurrencyID = readLen(r, 4)
}

type  CThostFtdcNotifyFutureSignInField struct {
    TradeCode string;
    BankID string;
    BankBranchID string;
    BrokerID string;
    BrokerBranchID string;
    TradeDate string;
    TradeTime string;
    BankSerial string;
    TradingDay string;
    PlateSerial int32;
    LastFragment uint8;
    SessionID int32;
    InstallID int32;
    UserID string;
    Digest string;
    CurrencyID string;
    DeviceID string;
    BrokerIDByBank string;
    OperNo string;
    RequestID int32;
    TID int32;
    ErrorID int32;
    ErrorMsg string;
    PinKey string;
    MacKey string;
}

func (m* CThostFtdcNotifyFutureSignInField) Marshal(w io.Writer) {
    writeLen(w, m.TradeCode, 7)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBranchID, 5)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.BrokerBranchID, 31)
    writeLen(w, m.TradeDate, 9)
    writeLen(w, m.TradeTime, 9)
    writeLen(w, m.BankSerial, 13)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.PlateSerial)
    binary.Write(w, binary.BigEndian, &m.LastFragment)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.Digest, 36)
    writeLen(w, m.CurrencyID, 4)
    writeLen(w, m.DeviceID, 3)
    writeLen(w, m.BrokerIDByBank, 33)
    writeLen(w, m.OperNo, 17)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    binary.Write(w, binary.BigEndian, &m.TID)
    binary.Write(w, binary.BigEndian, &m.ErrorID)
    writeLen(w, m.ErrorMsg, 81)
    writeLen(w, m.PinKey, 129)
    writeLen(w, m.MacKey, 129)
}

func (m* CThostFtdcNotifyFutureSignInField) Unmarshal(r io.Reader) {
    m.TradeCode = readLen(r, 7)
    m.BankID = readLen(r, 4)
    m.BankBranchID = readLen(r, 5)
    m.BrokerID = readLen(r, 11)
    m.BrokerBranchID = readLen(r, 31)
    m.TradeDate = readLen(r, 9)
    m.TradeTime = readLen(r, 9)
    m.BankSerial = readLen(r, 13)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.PlateSerial)
    binary.Read(r, binary.BigEndian, &m.LastFragment)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    m.UserID = readLen(r, 16)
    m.Digest = readLen(r, 36)
    m.CurrencyID = readLen(r, 4)
    m.DeviceID = readLen(r, 3)
    m.BrokerIDByBank = readLen(r, 33)
    m.OperNo = readLen(r, 17)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    binary.Read(r, binary.BigEndian, &m.TID)
    binary.Read(r, binary.BigEndian, &m.ErrorID)
    m.ErrorMsg = readLen(r, 81)
    m.PinKey = readLen(r, 129)
    m.MacKey = readLen(r, 129)
}

type  CThostFtdcNotifyFutureSignOutField struct {
    TradeCode string;
    BankID string;
    BankBranchID string;
    BrokerID string;
    BrokerBranchID string;
    TradeDate string;
    TradeTime string;
    BankSerial string;
    TradingDay string;
    PlateSerial int32;
    LastFragment uint8;
    SessionID int32;
    InstallID int32;
    UserID string;
    Digest string;
    CurrencyID string;
    DeviceID string;
    BrokerIDByBank string;
    OperNo string;
    RequestID int32;
    TID int32;
    ErrorID int32;
    ErrorMsg string;
}

func (m* CThostFtdcNotifyFutureSignOutField) Marshal(w io.Writer) {
    writeLen(w, m.TradeCode, 7)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBranchID, 5)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.BrokerBranchID, 31)
    writeLen(w, m.TradeDate, 9)
    writeLen(w, m.TradeTime, 9)
    writeLen(w, m.BankSerial, 13)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.PlateSerial)
    binary.Write(w, binary.BigEndian, &m.LastFragment)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.Digest, 36)
    writeLen(w, m.CurrencyID, 4)
    writeLen(w, m.DeviceID, 3)
    writeLen(w, m.BrokerIDByBank, 33)
    writeLen(w, m.OperNo, 17)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    binary.Write(w, binary.BigEndian, &m.TID)
    binary.Write(w, binary.BigEndian, &m.ErrorID)
    writeLen(w, m.ErrorMsg, 81)
}

func (m* CThostFtdcNotifyFutureSignOutField) Unmarshal(r io.Reader) {
    m.TradeCode = readLen(r, 7)
    m.BankID = readLen(r, 4)
    m.BankBranchID = readLen(r, 5)
    m.BrokerID = readLen(r, 11)
    m.BrokerBranchID = readLen(r, 31)
    m.TradeDate = readLen(r, 9)
    m.TradeTime = readLen(r, 9)
    m.BankSerial = readLen(r, 13)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.PlateSerial)
    binary.Read(r, binary.BigEndian, &m.LastFragment)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    m.UserID = readLen(r, 16)
    m.Digest = readLen(r, 36)
    m.CurrencyID = readLen(r, 4)
    m.DeviceID = readLen(r, 3)
    m.BrokerIDByBank = readLen(r, 33)
    m.OperNo = readLen(r, 17)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    binary.Read(r, binary.BigEndian, &m.TID)
    binary.Read(r, binary.BigEndian, &m.ErrorID)
    m.ErrorMsg = readLen(r, 81)
}

type  CThostFtdcNotifySyncKeyField struct {
    TradeCode string;
    BankID string;
    BankBranchID string;
    BrokerID string;
    BrokerBranchID string;
    TradeDate string;
    TradeTime string;
    BankSerial string;
    TradingDay string;
    PlateSerial int32;
    LastFragment uint8;
    SessionID int32;
    InstallID int32;
    UserID string;
    Message string;
    DeviceID string;
    BrokerIDByBank string;
    OperNo string;
    RequestID int32;
    TID int32;
    ErrorID int32;
    ErrorMsg string;
}

func (m* CThostFtdcNotifySyncKeyField) Marshal(w io.Writer) {
    writeLen(w, m.TradeCode, 7)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBranchID, 5)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.BrokerBranchID, 31)
    writeLen(w, m.TradeDate, 9)
    writeLen(w, m.TradeTime, 9)
    writeLen(w, m.BankSerial, 13)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.PlateSerial)
    binary.Write(w, binary.BigEndian, &m.LastFragment)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.Message, 129)
    writeLen(w, m.DeviceID, 3)
    writeLen(w, m.BrokerIDByBank, 33)
    writeLen(w, m.OperNo, 17)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    binary.Write(w, binary.BigEndian, &m.TID)
    binary.Write(w, binary.BigEndian, &m.ErrorID)
    writeLen(w, m.ErrorMsg, 81)
}

func (m* CThostFtdcNotifySyncKeyField) Unmarshal(r io.Reader) {
    m.TradeCode = readLen(r, 7)
    m.BankID = readLen(r, 4)
    m.BankBranchID = readLen(r, 5)
    m.BrokerID = readLen(r, 11)
    m.BrokerBranchID = readLen(r, 31)
    m.TradeDate = readLen(r, 9)
    m.TradeTime = readLen(r, 9)
    m.BankSerial = readLen(r, 13)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.PlateSerial)
    binary.Read(r, binary.BigEndian, &m.LastFragment)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    m.UserID = readLen(r, 16)
    m.Message = readLen(r, 129)
    m.DeviceID = readLen(r, 3)
    m.BrokerIDByBank = readLen(r, 33)
    m.OperNo = readLen(r, 17)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    binary.Read(r, binary.BigEndian, &m.TID)
    binary.Read(r, binary.BigEndian, &m.ErrorID)
    m.ErrorMsg = readLen(r, 81)
}

type  CThostFtdcQryAccountregisterField struct {
    BrokerID string;
    AccountID string;
    BankID string;
    BankBranchID string;
    CurrencyID string;
}

func (m* CThostFtdcQryAccountregisterField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBranchID, 5)
    writeLen(w, m.CurrencyID, 4)
}

func (m* CThostFtdcQryAccountregisterField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.AccountID = readLen(r, 13)
    m.BankID = readLen(r, 4)
    m.BankBranchID = readLen(r, 5)
    m.CurrencyID = readLen(r, 4)
}

type  CThostFtdcAccountregisterField struct {
    TradeDay string;
    BankID string;
    BankBranchID string;
    BankAccount string;
    BrokerID string;
    BrokerBranchID string;
    AccountID string;
    IdCardType uint8;
    IdentifiedCardNo string;
    CustomerName string;
    CurrencyID string;
    OpenOrDestroy uint8;
    RegDate string;
    OutDate string;
    TID int32;
    CustType uint8;
    BankAccType uint8;
}

func (m* CThostFtdcAccountregisterField) Marshal(w io.Writer) {
    writeLen(w, m.TradeDay, 9)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBranchID, 5)
    writeLen(w, m.BankAccount, 41)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.BrokerBranchID, 31)
    writeLen(w, m.AccountID, 13)
    binary.Write(w, binary.BigEndian, &m.IdCardType)
    writeLen(w, m.IdentifiedCardNo, 51)
    writeLen(w, m.CustomerName, 51)
    writeLen(w, m.CurrencyID, 4)
    binary.Write(w, binary.BigEndian, &m.OpenOrDestroy)
    writeLen(w, m.RegDate, 9)
    writeLen(w, m.OutDate, 9)
    binary.Write(w, binary.BigEndian, &m.TID)
    binary.Write(w, binary.BigEndian, &m.CustType)
    binary.Write(w, binary.BigEndian, &m.BankAccType)
}

func (m* CThostFtdcAccountregisterField) Unmarshal(r io.Reader) {
    m.TradeDay = readLen(r, 9)
    m.BankID = readLen(r, 4)
    m.BankBranchID = readLen(r, 5)
    m.BankAccount = readLen(r, 41)
    m.BrokerID = readLen(r, 11)
    m.BrokerBranchID = readLen(r, 31)
    m.AccountID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.IdCardType)
    m.IdentifiedCardNo = readLen(r, 51)
    m.CustomerName = readLen(r, 51)
    m.CurrencyID = readLen(r, 4)
    binary.Read(r, binary.BigEndian, &m.OpenOrDestroy)
    m.RegDate = readLen(r, 9)
    m.OutDate = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.TID)
    binary.Read(r, binary.BigEndian, &m.CustType)
    binary.Read(r, binary.BigEndian, &m.BankAccType)
}

type  CThostFtdcOpenAccountField struct {
    TradeCode string;
    BankID string;
    BankBranchID string;
    BrokerID string;
    BrokerBranchID string;
    TradeDate string;
    TradeTime string;
    BankSerial string;
    TradingDay string;
    PlateSerial int32;
    LastFragment uint8;
    SessionID int32;
    CustomerName string;
    IdCardType uint8;
    IdentifiedCardNo string;
    Gender uint8;
    CountryCode string;
    CustType uint8;
    Address string;
    ZipCode string;
    Telephone string;
    MobilePhone string;
    Fax string;
    EMail string;
    MoneyAccountStatus uint8;
    BankAccount string;
    BankPassWord string;
    AccountID string;
    Password string;
    InstallID int32;
    VerifyCertNoFlag uint8;
    CurrencyID string;
    CashExchangeCode uint8;
    Digest string;
    BankAccType uint8;
    DeviceID string;
    BankSecuAccType uint8;
    BrokerIDByBank string;
    BankSecuAcc string;
    BankPwdFlag uint8;
    SecuPwdFlag uint8;
    OperNo string;
    TID int32;
    UserID string;
    ErrorID int32;
    ErrorMsg string;
    SecDaBeBl float64;
    BankChal uint8;
}

func (m* CThostFtdcOpenAccountField) Marshal(w io.Writer) {
    writeLen(w, m.TradeCode, 7)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBranchID, 5)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.BrokerBranchID, 31)
    writeLen(w, m.TradeDate, 9)
    writeLen(w, m.TradeTime, 9)
    writeLen(w, m.BankSerial, 13)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.PlateSerial)
    binary.Write(w, binary.BigEndian, &m.LastFragment)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.CustomerName, 51)
    binary.Write(w, binary.BigEndian, &m.IdCardType)
    writeLen(w, m.IdentifiedCardNo, 51)
    binary.Write(w, binary.BigEndian, &m.Gender)
    writeLen(w, m.CountryCode, 21)
    binary.Write(w, binary.BigEndian, &m.CustType)
    writeLen(w, m.Address, 101)
    writeLen(w, m.ZipCode, 7)
    writeLen(w, m.Telephone, 41)
    writeLen(w, m.MobilePhone, 21)
    writeLen(w, m.Fax, 41)
    writeLen(w, m.EMail, 41)
    binary.Write(w, binary.BigEndian, &m.MoneyAccountStatus)
    writeLen(w, m.BankAccount, 41)
    writeLen(w, m.BankPassWord, 41)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.Password, 41)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    binary.Write(w, binary.BigEndian, &m.VerifyCertNoFlag)
    writeLen(w, m.CurrencyID, 4)
    binary.Write(w, binary.BigEndian, &m.CashExchangeCode)
    writeLen(w, m.Digest, 36)
    binary.Write(w, binary.BigEndian, &m.BankAccType)
    writeLen(w, m.DeviceID, 3)
    binary.Write(w, binary.BigEndian, &m.BankSecuAccType)
    writeLen(w, m.BrokerIDByBank, 33)
    writeLen(w, m.BankSecuAcc, 41)
    binary.Write(w, binary.BigEndian, &m.BankPwdFlag)
    binary.Write(w, binary.BigEndian, &m.SecuPwdFlag)
    writeLen(w, m.OperNo, 17)
    binary.Write(w, binary.BigEndian, &m.TID)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.ErrorID)
    writeLen(w, m.ErrorMsg, 81)
    binary.Write(w, binary.BigEndian, &m.SecDaBeBl)
    binary.Write(w, binary.BigEndian, &m.BankChal)
}

func (m* CThostFtdcOpenAccountField) Unmarshal(r io.Reader) {
    m.TradeCode = readLen(r, 7)
    m.BankID = readLen(r, 4)
    m.BankBranchID = readLen(r, 5)
    m.BrokerID = readLen(r, 11)
    m.BrokerBranchID = readLen(r, 31)
    m.TradeDate = readLen(r, 9)
    m.TradeTime = readLen(r, 9)
    m.BankSerial = readLen(r, 13)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.PlateSerial)
    binary.Read(r, binary.BigEndian, &m.LastFragment)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.CustomerName = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.IdCardType)
    m.IdentifiedCardNo = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.Gender)
    m.CountryCode = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.CustType)
    m.Address = readLen(r, 101)
    m.ZipCode = readLen(r, 7)
    m.Telephone = readLen(r, 41)
    m.MobilePhone = readLen(r, 21)
    m.Fax = readLen(r, 41)
    m.EMail = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.MoneyAccountStatus)
    m.BankAccount = readLen(r, 41)
    m.BankPassWord = readLen(r, 41)
    m.AccountID = readLen(r, 13)
    m.Password = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    binary.Read(r, binary.BigEndian, &m.VerifyCertNoFlag)
    m.CurrencyID = readLen(r, 4)
    binary.Read(r, binary.BigEndian, &m.CashExchangeCode)
    m.Digest = readLen(r, 36)
    binary.Read(r, binary.BigEndian, &m.BankAccType)
    m.DeviceID = readLen(r, 3)
    binary.Read(r, binary.BigEndian, &m.BankSecuAccType)
    m.BrokerIDByBank = readLen(r, 33)
    m.BankSecuAcc = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.BankPwdFlag)
    binary.Read(r, binary.BigEndian, &m.SecuPwdFlag)
    m.OperNo = readLen(r, 17)
    binary.Read(r, binary.BigEndian, &m.TID)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.ErrorID)
    m.ErrorMsg = readLen(r, 81)
    binary.Read(r, binary.BigEndian, &m.SecDaBeBl)
    binary.Read(r, binary.BigEndian, &m.BankChal)
}

type  CThostFtdcCancelAccountField struct {
    TradeCode string;
    BankID string;
    BankBranchID string;
    BrokerID string;
    BrokerBranchID string;
    TradeDate string;
    TradeTime string;
    BankSerial string;
    TradingDay string;
    PlateSerial int32;
    LastFragment uint8;
    SessionID int32;
    CustomerName string;
    IdCardType uint8;
    IdentifiedCardNo string;
    Gender uint8;
    CountryCode string;
    CustType uint8;
    Address string;
    ZipCode string;
    Telephone string;
    MobilePhone string;
    Fax string;
    EMail string;
    MoneyAccountStatus uint8;
    BankAccount string;
    BankPassWord string;
    AccountID string;
    Password string;
    InstallID int32;
    VerifyCertNoFlag uint8;
    CurrencyID string;
    CashExchangeCode uint8;
    Digest string;
    BankAccType uint8;
    DeviceID string;
    BankSecuAccType uint8;
    BrokerIDByBank string;
    BankSecuAcc string;
    BankPwdFlag uint8;
    SecuPwdFlag uint8;
    OperNo string;
    TID int32;
    UserID string;
    SecDaBeBl float64;
    BankChal uint8;
    ErrorID int32;
    ErrorMsg string;
}

func (m* CThostFtdcCancelAccountField) Marshal(w io.Writer) {
    writeLen(w, m.TradeCode, 7)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBranchID, 5)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.BrokerBranchID, 31)
    writeLen(w, m.TradeDate, 9)
    writeLen(w, m.TradeTime, 9)
    writeLen(w, m.BankSerial, 13)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.PlateSerial)
    binary.Write(w, binary.BigEndian, &m.LastFragment)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.CustomerName, 51)
    binary.Write(w, binary.BigEndian, &m.IdCardType)
    writeLen(w, m.IdentifiedCardNo, 51)
    binary.Write(w, binary.BigEndian, &m.Gender)
    writeLen(w, m.CountryCode, 21)
    binary.Write(w, binary.BigEndian, &m.CustType)
    writeLen(w, m.Address, 101)
    writeLen(w, m.ZipCode, 7)
    writeLen(w, m.Telephone, 41)
    writeLen(w, m.MobilePhone, 21)
    writeLen(w, m.Fax, 41)
    writeLen(w, m.EMail, 41)
    binary.Write(w, binary.BigEndian, &m.MoneyAccountStatus)
    writeLen(w, m.BankAccount, 41)
    writeLen(w, m.BankPassWord, 41)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.Password, 41)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    binary.Write(w, binary.BigEndian, &m.VerifyCertNoFlag)
    writeLen(w, m.CurrencyID, 4)
    binary.Write(w, binary.BigEndian, &m.CashExchangeCode)
    writeLen(w, m.Digest, 36)
    binary.Write(w, binary.BigEndian, &m.BankAccType)
    writeLen(w, m.DeviceID, 3)
    binary.Write(w, binary.BigEndian, &m.BankSecuAccType)
    writeLen(w, m.BrokerIDByBank, 33)
    writeLen(w, m.BankSecuAcc, 41)
    binary.Write(w, binary.BigEndian, &m.BankPwdFlag)
    binary.Write(w, binary.BigEndian, &m.SecuPwdFlag)
    writeLen(w, m.OperNo, 17)
    binary.Write(w, binary.BigEndian, &m.TID)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.SecDaBeBl)
    binary.Write(w, binary.BigEndian, &m.BankChal)
    binary.Write(w, binary.BigEndian, &m.ErrorID)
    writeLen(w, m.ErrorMsg, 81)
}

func (m* CThostFtdcCancelAccountField) Unmarshal(r io.Reader) {
    m.TradeCode = readLen(r, 7)
    m.BankID = readLen(r, 4)
    m.BankBranchID = readLen(r, 5)
    m.BrokerID = readLen(r, 11)
    m.BrokerBranchID = readLen(r, 31)
    m.TradeDate = readLen(r, 9)
    m.TradeTime = readLen(r, 9)
    m.BankSerial = readLen(r, 13)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.PlateSerial)
    binary.Read(r, binary.BigEndian, &m.LastFragment)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.CustomerName = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.IdCardType)
    m.IdentifiedCardNo = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.Gender)
    m.CountryCode = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.CustType)
    m.Address = readLen(r, 101)
    m.ZipCode = readLen(r, 7)
    m.Telephone = readLen(r, 41)
    m.MobilePhone = readLen(r, 21)
    m.Fax = readLen(r, 41)
    m.EMail = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.MoneyAccountStatus)
    m.BankAccount = readLen(r, 41)
    m.BankPassWord = readLen(r, 41)
    m.AccountID = readLen(r, 13)
    m.Password = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    binary.Read(r, binary.BigEndian, &m.VerifyCertNoFlag)
    m.CurrencyID = readLen(r, 4)
    binary.Read(r, binary.BigEndian, &m.CashExchangeCode)
    m.Digest = readLen(r, 36)
    binary.Read(r, binary.BigEndian, &m.BankAccType)
    m.DeviceID = readLen(r, 3)
    binary.Read(r, binary.BigEndian, &m.BankSecuAccType)
    m.BrokerIDByBank = readLen(r, 33)
    m.BankSecuAcc = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.BankPwdFlag)
    binary.Read(r, binary.BigEndian, &m.SecuPwdFlag)
    m.OperNo = readLen(r, 17)
    binary.Read(r, binary.BigEndian, &m.TID)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.SecDaBeBl)
    binary.Read(r, binary.BigEndian, &m.BankChal)
    binary.Read(r, binary.BigEndian, &m.ErrorID)
    m.ErrorMsg = readLen(r, 81)
}

type  CThostFtdcChangeAccountField struct {
    TradeCode string;
    BankID string;
    BankBranchID string;
    BrokerID string;
    BrokerBranchID string;
    TradeDate string;
    TradeTime string;
    BankSerial string;
    TradingDay string;
    PlateSerial int32;
    LastFragment uint8;
    SessionID int32;
    CustomerName string;
    IdCardType uint8;
    IdentifiedCardNo string;
    Gender uint8;
    CountryCode string;
    CustType uint8;
    Address string;
    ZipCode string;
    Telephone string;
    MobilePhone string;
    Fax string;
    EMail string;
    MoneyAccountStatus uint8;
    BankAccount string;
    BankPassWord string;
    NewBankAccount string;
    NewBankPassWord string;
    AccountID string;
    Password string;
    BankAccType uint8;
    InstallID int32;
    VerifyCertNoFlag uint8;
    CurrencyID string;
    BrokerIDByBank string;
    BankPwdFlag uint8;
    SecuPwdFlag uint8;
    TID int32;
    Digest string;
    ErrorID int32;
    ErrorMsg string;
}

func (m* CThostFtdcChangeAccountField) Marshal(w io.Writer) {
    writeLen(w, m.TradeCode, 7)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBranchID, 5)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.BrokerBranchID, 31)
    writeLen(w, m.TradeDate, 9)
    writeLen(w, m.TradeTime, 9)
    writeLen(w, m.BankSerial, 13)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.PlateSerial)
    binary.Write(w, binary.BigEndian, &m.LastFragment)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.CustomerName, 51)
    binary.Write(w, binary.BigEndian, &m.IdCardType)
    writeLen(w, m.IdentifiedCardNo, 51)
    binary.Write(w, binary.BigEndian, &m.Gender)
    writeLen(w, m.CountryCode, 21)
    binary.Write(w, binary.BigEndian, &m.CustType)
    writeLen(w, m.Address, 101)
    writeLen(w, m.ZipCode, 7)
    writeLen(w, m.Telephone, 41)
    writeLen(w, m.MobilePhone, 21)
    writeLen(w, m.Fax, 41)
    writeLen(w, m.EMail, 41)
    binary.Write(w, binary.BigEndian, &m.MoneyAccountStatus)
    writeLen(w, m.BankAccount, 41)
    writeLen(w, m.BankPassWord, 41)
    writeLen(w, m.NewBankAccount, 41)
    writeLen(w, m.NewBankPassWord, 41)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.Password, 41)
    binary.Write(w, binary.BigEndian, &m.BankAccType)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    binary.Write(w, binary.BigEndian, &m.VerifyCertNoFlag)
    writeLen(w, m.CurrencyID, 4)
    writeLen(w, m.BrokerIDByBank, 33)
    binary.Write(w, binary.BigEndian, &m.BankPwdFlag)
    binary.Write(w, binary.BigEndian, &m.SecuPwdFlag)
    binary.Write(w, binary.BigEndian, &m.TID)
    writeLen(w, m.Digest, 36)
    binary.Write(w, binary.BigEndian, &m.ErrorID)
    writeLen(w, m.ErrorMsg, 81)
}

func (m* CThostFtdcChangeAccountField) Unmarshal(r io.Reader) {
    m.TradeCode = readLen(r, 7)
    m.BankID = readLen(r, 4)
    m.BankBranchID = readLen(r, 5)
    m.BrokerID = readLen(r, 11)
    m.BrokerBranchID = readLen(r, 31)
    m.TradeDate = readLen(r, 9)
    m.TradeTime = readLen(r, 9)
    m.BankSerial = readLen(r, 13)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.PlateSerial)
    binary.Read(r, binary.BigEndian, &m.LastFragment)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.CustomerName = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.IdCardType)
    m.IdentifiedCardNo = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.Gender)
    m.CountryCode = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.CustType)
    m.Address = readLen(r, 101)
    m.ZipCode = readLen(r, 7)
    m.Telephone = readLen(r, 41)
    m.MobilePhone = readLen(r, 21)
    m.Fax = readLen(r, 41)
    m.EMail = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.MoneyAccountStatus)
    m.BankAccount = readLen(r, 41)
    m.BankPassWord = readLen(r, 41)
    m.NewBankAccount = readLen(r, 41)
    m.NewBankPassWord = readLen(r, 41)
    m.AccountID = readLen(r, 13)
    m.Password = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.BankAccType)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    binary.Read(r, binary.BigEndian, &m.VerifyCertNoFlag)
    m.CurrencyID = readLen(r, 4)
    m.BrokerIDByBank = readLen(r, 33)
    binary.Read(r, binary.BigEndian, &m.BankPwdFlag)
    binary.Read(r, binary.BigEndian, &m.SecuPwdFlag)
    binary.Read(r, binary.BigEndian, &m.TID)
    m.Digest = readLen(r, 36)
    binary.Read(r, binary.BigEndian, &m.ErrorID)
    m.ErrorMsg = readLen(r, 81)
}

type  CThostFtdcSecAgentACIDMapField struct {
    BrokerID string;
    UserID string;
    AccountID string;
    CurrencyID string;
    BrokerSecAgentID string;
}

func (m* CThostFtdcSecAgentACIDMapField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.CurrencyID, 4)
    writeLen(w, m.BrokerSecAgentID, 13)
}

func (m* CThostFtdcSecAgentACIDMapField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.UserID = readLen(r, 16)
    m.AccountID = readLen(r, 13)
    m.CurrencyID = readLen(r, 4)
    m.BrokerSecAgentID = readLen(r, 13)
}

type  CThostFtdcQrySecAgentACIDMapField struct {
    BrokerID string;
    UserID string;
    AccountID string;
    CurrencyID string;
}

func (m* CThostFtdcQrySecAgentACIDMapField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.CurrencyID, 4)
}

func (m* CThostFtdcQrySecAgentACIDMapField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.UserID = readLen(r, 16)
    m.AccountID = readLen(r, 13)
    m.CurrencyID = readLen(r, 4)
}

type  CThostFtdcUserRightsAssignField struct {
    BrokerID string;
    UserID string;
    DRIdentityID int32;
}

func (m* CThostFtdcUserRightsAssignField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.DRIdentityID)
}

func (m* CThostFtdcUserRightsAssignField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.DRIdentityID)
}

type  CThostFtdcBrokerUserRightAssignField struct {
    BrokerID string;
    DRIdentityID int32;
    Tradeable int32;
}

func (m* CThostFtdcBrokerUserRightAssignField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    binary.Write(w, binary.BigEndian, &m.DRIdentityID)
    binary.Write(w, binary.BigEndian, &m.Tradeable)
}

func (m* CThostFtdcBrokerUserRightAssignField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    binary.Read(r, binary.BigEndian, &m.DRIdentityID)
    binary.Read(r, binary.BigEndian, &m.Tradeable)
}

type  CThostFtdcDRTransferField struct {
    OrigDRIdentityID int32;
    DestDRIdentityID int32;
    OrigBrokerID string;
    DestBrokerID string;
}

func (m* CThostFtdcDRTransferField) Marshal(w io.Writer) {
    binary.Write(w, binary.BigEndian, &m.OrigDRIdentityID)
    binary.Write(w, binary.BigEndian, &m.DestDRIdentityID)
    writeLen(w, m.OrigBrokerID, 11)
    writeLen(w, m.DestBrokerID, 11)
}

func (m* CThostFtdcDRTransferField) Unmarshal(r io.Reader) {
    binary.Read(r, binary.BigEndian, &m.OrigDRIdentityID)
    binary.Read(r, binary.BigEndian, &m.DestDRIdentityID)
    m.OrigBrokerID = readLen(r, 11)
    m.DestBrokerID = readLen(r, 11)
}

type  CThostFtdcFensUserInfoField struct {
    BrokerID string;
    UserID string;
    LoginMode uint8;
}

func (m* CThostFtdcFensUserInfoField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.LoginMode)
}

func (m* CThostFtdcFensUserInfoField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.LoginMode)
}

type  CThostFtdcCurrTransferIdentityField struct {
    IdentityID int32;
}

func (m* CThostFtdcCurrTransferIdentityField) Marshal(w io.Writer) {
    binary.Write(w, binary.BigEndian, &m.IdentityID)
}

func (m* CThostFtdcCurrTransferIdentityField) Unmarshal(r io.Reader) {
    binary.Read(r, binary.BigEndian, &m.IdentityID)
}

type  CThostFtdcLoginForbiddenUserField struct {
    BrokerID string;
    UserID string;
    IPAddress string;
}

func (m* CThostFtdcLoginForbiddenUserField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.IPAddress, 16)
}

func (m* CThostFtdcLoginForbiddenUserField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.UserID = readLen(r, 16)
    m.IPAddress = readLen(r, 16)
}

type  CThostFtdcQryLoginForbiddenUserField struct {
    BrokerID string;
    UserID string;
}

func (m* CThostFtdcQryLoginForbiddenUserField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.UserID, 16)
}

func (m* CThostFtdcQryLoginForbiddenUserField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.UserID = readLen(r, 16)
}

type  CThostFtdcMulticastGroupInfoField struct {
    GroupIP string;
    GroupPort int32;
    SourceIP string;
}

func (m* CThostFtdcMulticastGroupInfoField) Marshal(w io.Writer) {
    writeLen(w, m.GroupIP, 16)
    binary.Write(w, binary.BigEndian, &m.GroupPort)
    writeLen(w, m.SourceIP, 16)
}

func (m* CThostFtdcMulticastGroupInfoField) Unmarshal(r io.Reader) {
    m.GroupIP = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.GroupPort)
    m.SourceIP = readLen(r, 16)
}

type  CThostFtdcTradingAccountReserveField struct {
    BrokerID string;
    AccountID string;
    Reserve float64;
    CurrencyID string;
}

func (m* CThostFtdcTradingAccountReserveField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.AccountID, 13)
    binary.Write(w, binary.BigEndian, &m.Reserve)
    writeLen(w, m.CurrencyID, 4)
}

func (m* CThostFtdcTradingAccountReserveField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.AccountID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.Reserve)
    m.CurrencyID = readLen(r, 4)
}

type  CThostFtdcQryLoginForbiddenIPField struct {
    IPAddress string;
}

func (m* CThostFtdcQryLoginForbiddenIPField) Marshal(w io.Writer) {
    writeLen(w, m.IPAddress, 16)
}

func (m* CThostFtdcQryLoginForbiddenIPField) Unmarshal(r io.Reader) {
    m.IPAddress = readLen(r, 16)
}

type  CThostFtdcQryIPListField struct {
    IPAddress string;
}

func (m* CThostFtdcQryIPListField) Marshal(w io.Writer) {
    writeLen(w, m.IPAddress, 16)
}

func (m* CThostFtdcQryIPListField) Unmarshal(r io.Reader) {
    m.IPAddress = readLen(r, 16)
}

type  CThostFtdcQryUserRightsAssignField struct {
    BrokerID string;
    UserID string;
}

func (m* CThostFtdcQryUserRightsAssignField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.UserID, 16)
}

func (m* CThostFtdcQryUserRightsAssignField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.UserID = readLen(r, 16)
}

type  CThostFtdcDBFRecordField struct {
    DBFComdType string;
    DBFComTime string;
    DBFOComNo string;
    DBFComNo string;
    DBFFdName1 string;
    DBFFdContent1 string;
    DBFFdName2 string;
    DBFFdContent2 string;
    DBFFdName3 string;
    DBFFdContent3 string;
    DBFFdName4 string;
    DBFFdContent4 string;
}

func (m* CThostFtdcDBFRecordField) Marshal(w io.Writer) {
    writeLen(w, m.DBFComdType, 65)
    writeLen(w, m.DBFComTime, 65)
    writeLen(w, m.DBFOComNo, 17)
    writeLen(w, m.DBFComNo, 17)
    writeLen(w, m.DBFFdName1, 256)
    writeLen(w, m.DBFFdContent1, 256)
    writeLen(w, m.DBFFdName2, 256)
    writeLen(w, m.DBFFdContent2, 256)
    writeLen(w, m.DBFFdName3, 256)
    writeLen(w, m.DBFFdContent3, 256)
    writeLen(w, m.DBFFdName4, 256)
    writeLen(w, m.DBFFdContent4, 256)
}

func (m* CThostFtdcDBFRecordField) Unmarshal(r io.Reader) {
    m.DBFComdType = readLen(r, 65)
    m.DBFComTime = readLen(r, 65)
    m.DBFOComNo = readLen(r, 17)
    m.DBFComNo = readLen(r, 17)
    m.DBFFdName1 = readLen(r, 256)
    m.DBFFdContent1 = readLen(r, 256)
    m.DBFFdName2 = readLen(r, 256)
    m.DBFFdContent2 = readLen(r, 256)
    m.DBFFdName3 = readLen(r, 256)
    m.DBFFdContent3 = readLen(r, 256)
    m.DBFFdName4 = readLen(r, 256)
    m.DBFFdContent4 = readLen(r, 256)
}

type  CThostFtdcAccountPropertyField struct {
    BrokerID string;
    AccountID string;
    BankID string;
    BankAccount string;
    OpenName string;
    OpenBank string;
    IsActive int32;
    AccountSourceType uint8;
    OpenDate string;
    CancelDate string;
    OperatorID string;
    OperateDate string;
    OperateTime string;
    CurrencyID string;
}

func (m* CThostFtdcAccountPropertyField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankAccount, 41)
    writeLen(w, m.OpenName, 101)
    writeLen(w, m.OpenBank, 101)
    binary.Write(w, binary.BigEndian, &m.IsActive)
    binary.Write(w, binary.BigEndian, &m.AccountSourceType)
    writeLen(w, m.OpenDate, 9)
    writeLen(w, m.CancelDate, 9)
    writeLen(w, m.OperatorID, 65)
    writeLen(w, m.OperateDate, 9)
    writeLen(w, m.OperateTime, 9)
    writeLen(w, m.CurrencyID, 4)
}

func (m* CThostFtdcAccountPropertyField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.AccountID = readLen(r, 13)
    m.BankID = readLen(r, 4)
    m.BankAccount = readLen(r, 41)
    m.OpenName = readLen(r, 101)
    m.OpenBank = readLen(r, 101)
    binary.Read(r, binary.BigEndian, &m.IsActive)
    binary.Read(r, binary.BigEndian, &m.AccountSourceType)
    m.OpenDate = readLen(r, 9)
    m.CancelDate = readLen(r, 9)
    m.OperatorID = readLen(r, 65)
    m.OperateDate = readLen(r, 9)
    m.OperateTime = readLen(r, 9)
    m.CurrencyID = readLen(r, 4)
}

type  CThostFtdcQryCurrDRIdentityField struct {
    DRIdentityID int32;
}

func (m* CThostFtdcQryCurrDRIdentityField) Marshal(w io.Writer) {
    binary.Write(w, binary.BigEndian, &m.DRIdentityID)
}

func (m* CThostFtdcQryCurrDRIdentityField) Unmarshal(r io.Reader) {
    binary.Read(r, binary.BigEndian, &m.DRIdentityID)
}

type  CThostFtdcCurrDRIdentityField struct {
    DRIdentityID int32;
}

func (m* CThostFtdcCurrDRIdentityField) Marshal(w io.Writer) {
    binary.Write(w, binary.BigEndian, &m.DRIdentityID)
}

func (m* CThostFtdcCurrDRIdentityField) Unmarshal(r io.Reader) {
    binary.Read(r, binary.BigEndian, &m.DRIdentityID)
}

type  CThostFtdcReqReserveOpenAccountTpdField struct {
    TradeCode string;
    BankID string;
    BankBranchID string;
    BrokerID string;
    BrokerBranchID string;
    TradeDate string;
    TradeTime string;
    BankSerial string;
    TradingDay string;
    PlateSerial int32;
    LastFragment uint8;
    SessionID int32;
    CustomerName string;
    IdCardType uint8;
    IdentifiedCardNo string;
    Gender uint8;
    CountryCode string;
    CustType uint8;
    Address string;
    ZipCode string;
    Telephone string;
    MobilePhone string;
    Fax string;
    EMail string;
    MoneyAccountStatus uint8;
    BankAccount string;
    BankPassWord string;
    AccountID string;
    Password string;
    InstallID int32;
    VerifyCertNoFlag uint8;
    CurrencyID string;
    CashExchangeCode uint8;
    Digest string;
    BankAccType uint8;
    DeviceID string;
    BankSecuAccType uint8;
    BrokerIDByBank string;
    BankSecuAcc string;
    BankPwdFlag uint8;
    SecuPwdFlag uint8;
    OperNo string;
    TID int32;
    UserID string;
    SecDaBeBl float64;
    BankChal uint8;
}

func (m* CThostFtdcReqReserveOpenAccountTpdField) Marshal(w io.Writer) {
    writeLen(w, m.TradeCode, 7)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBranchID, 5)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.BrokerBranchID, 31)
    writeLen(w, m.TradeDate, 9)
    writeLen(w, m.TradeTime, 9)
    writeLen(w, m.BankSerial, 13)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.PlateSerial)
    binary.Write(w, binary.BigEndian, &m.LastFragment)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.CustomerName, 51)
    binary.Write(w, binary.BigEndian, &m.IdCardType)
    writeLen(w, m.IdentifiedCardNo, 51)
    binary.Write(w, binary.BigEndian, &m.Gender)
    writeLen(w, m.CountryCode, 21)
    binary.Write(w, binary.BigEndian, &m.CustType)
    writeLen(w, m.Address, 101)
    writeLen(w, m.ZipCode, 7)
    writeLen(w, m.Telephone, 41)
    writeLen(w, m.MobilePhone, 21)
    writeLen(w, m.Fax, 41)
    writeLen(w, m.EMail, 41)
    binary.Write(w, binary.BigEndian, &m.MoneyAccountStatus)
    writeLen(w, m.BankAccount, 41)
    writeLen(w, m.BankPassWord, 41)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.Password, 41)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    binary.Write(w, binary.BigEndian, &m.VerifyCertNoFlag)
    writeLen(w, m.CurrencyID, 4)
    binary.Write(w, binary.BigEndian, &m.CashExchangeCode)
    writeLen(w, m.Digest, 36)
    binary.Write(w, binary.BigEndian, &m.BankAccType)
    writeLen(w, m.DeviceID, 3)
    binary.Write(w, binary.BigEndian, &m.BankSecuAccType)
    writeLen(w, m.BrokerIDByBank, 33)
    writeLen(w, m.BankSecuAcc, 41)
    binary.Write(w, binary.BigEndian, &m.BankPwdFlag)
    binary.Write(w, binary.BigEndian, &m.SecuPwdFlag)
    writeLen(w, m.OperNo, 17)
    binary.Write(w, binary.BigEndian, &m.TID)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.SecDaBeBl)
    binary.Write(w, binary.BigEndian, &m.BankChal)
}

func (m* CThostFtdcReqReserveOpenAccountTpdField) Unmarshal(r io.Reader) {
    m.TradeCode = readLen(r, 7)
    m.BankID = readLen(r, 4)
    m.BankBranchID = readLen(r, 5)
    m.BrokerID = readLen(r, 11)
    m.BrokerBranchID = readLen(r, 31)
    m.TradeDate = readLen(r, 9)
    m.TradeTime = readLen(r, 9)
    m.BankSerial = readLen(r, 13)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.PlateSerial)
    binary.Read(r, binary.BigEndian, &m.LastFragment)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.CustomerName = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.IdCardType)
    m.IdentifiedCardNo = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.Gender)
    m.CountryCode = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.CustType)
    m.Address = readLen(r, 101)
    m.ZipCode = readLen(r, 7)
    m.Telephone = readLen(r, 41)
    m.MobilePhone = readLen(r, 21)
    m.Fax = readLen(r, 41)
    m.EMail = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.MoneyAccountStatus)
    m.BankAccount = readLen(r, 41)
    m.BankPassWord = readLen(r, 41)
    m.AccountID = readLen(r, 13)
    m.Password = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    binary.Read(r, binary.BigEndian, &m.VerifyCertNoFlag)
    m.CurrencyID = readLen(r, 4)
    binary.Read(r, binary.BigEndian, &m.CashExchangeCode)
    m.Digest = readLen(r, 36)
    binary.Read(r, binary.BigEndian, &m.BankAccType)
    m.DeviceID = readLen(r, 3)
    binary.Read(r, binary.BigEndian, &m.BankSecuAccType)
    m.BrokerIDByBank = readLen(r, 33)
    m.BankSecuAcc = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.BankPwdFlag)
    binary.Read(r, binary.BigEndian, &m.SecuPwdFlag)
    m.OperNo = readLen(r, 17)
    binary.Read(r, binary.BigEndian, &m.TID)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.SecDaBeBl)
    binary.Read(r, binary.BigEndian, &m.BankChal)
}

type  CThostFtdcReserveOpenAccountTpdField struct {
    TradeCode string;
    BankID string;
    BankBranchID string;
    BrokerID string;
    BrokerBranchID string;
    TradeDate string;
    TradeTime string;
    BankSerial string;
    TradingDay string;
    PlateSerial int32;
    LastFragment uint8;
    SessionID int32;
    CustomerName string;
    IdCardType uint8;
    IdentifiedCardNo string;
    Gender uint8;
    CountryCode string;
    CustType uint8;
    Address string;
    ZipCode string;
    Telephone string;
    MobilePhone string;
    Fax string;
    EMail string;
    MoneyAccountStatus uint8;
    BankAccount string;
    BankPassWord string;
    AccountID string;
    Password string;
    InstallID int32;
    VerifyCertNoFlag uint8;
    CurrencyID string;
    CashExchangeCode uint8;
    Digest string;
    BankAccType uint8;
    DeviceID string;
    BankSecuAccType uint8;
    BrokerIDByBank string;
    BankSecuAcc string;
    BankPwdFlag uint8;
    SecuPwdFlag uint8;
    OperNo string;
    TID int32;
    UserID string;
    SecDaBeBl float64;
    BankChal uint8;
    ReserveOpenAccStas uint8;
    ErrorID int32;
    ErrorMsg string;
}

func (m* CThostFtdcReserveOpenAccountTpdField) Marshal(w io.Writer) {
    writeLen(w, m.TradeCode, 7)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBranchID, 5)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.BrokerBranchID, 31)
    writeLen(w, m.TradeDate, 9)
    writeLen(w, m.TradeTime, 9)
    writeLen(w, m.BankSerial, 13)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.PlateSerial)
    binary.Write(w, binary.BigEndian, &m.LastFragment)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.CustomerName, 51)
    binary.Write(w, binary.BigEndian, &m.IdCardType)
    writeLen(w, m.IdentifiedCardNo, 51)
    binary.Write(w, binary.BigEndian, &m.Gender)
    writeLen(w, m.CountryCode, 21)
    binary.Write(w, binary.BigEndian, &m.CustType)
    writeLen(w, m.Address, 101)
    writeLen(w, m.ZipCode, 7)
    writeLen(w, m.Telephone, 41)
    writeLen(w, m.MobilePhone, 21)
    writeLen(w, m.Fax, 41)
    writeLen(w, m.EMail, 41)
    binary.Write(w, binary.BigEndian, &m.MoneyAccountStatus)
    writeLen(w, m.BankAccount, 41)
    writeLen(w, m.BankPassWord, 41)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.Password, 41)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    binary.Write(w, binary.BigEndian, &m.VerifyCertNoFlag)
    writeLen(w, m.CurrencyID, 4)
    binary.Write(w, binary.BigEndian, &m.CashExchangeCode)
    writeLen(w, m.Digest, 36)
    binary.Write(w, binary.BigEndian, &m.BankAccType)
    writeLen(w, m.DeviceID, 3)
    binary.Write(w, binary.BigEndian, &m.BankSecuAccType)
    writeLen(w, m.BrokerIDByBank, 33)
    writeLen(w, m.BankSecuAcc, 41)
    binary.Write(w, binary.BigEndian, &m.BankPwdFlag)
    binary.Write(w, binary.BigEndian, &m.SecuPwdFlag)
    writeLen(w, m.OperNo, 17)
    binary.Write(w, binary.BigEndian, &m.TID)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.SecDaBeBl)
    binary.Write(w, binary.BigEndian, &m.BankChal)
    binary.Write(w, binary.BigEndian, &m.ReserveOpenAccStas)
    binary.Write(w, binary.BigEndian, &m.ErrorID)
    writeLen(w, m.ErrorMsg, 81)
}

func (m* CThostFtdcReserveOpenAccountTpdField) Unmarshal(r io.Reader) {
    m.TradeCode = readLen(r, 7)
    m.BankID = readLen(r, 4)
    m.BankBranchID = readLen(r, 5)
    m.BrokerID = readLen(r, 11)
    m.BrokerBranchID = readLen(r, 31)
    m.TradeDate = readLen(r, 9)
    m.TradeTime = readLen(r, 9)
    m.BankSerial = readLen(r, 13)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.PlateSerial)
    binary.Read(r, binary.BigEndian, &m.LastFragment)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.CustomerName = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.IdCardType)
    m.IdentifiedCardNo = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.Gender)
    m.CountryCode = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.CustType)
    m.Address = readLen(r, 101)
    m.ZipCode = readLen(r, 7)
    m.Telephone = readLen(r, 41)
    m.MobilePhone = readLen(r, 21)
    m.Fax = readLen(r, 41)
    m.EMail = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.MoneyAccountStatus)
    m.BankAccount = readLen(r, 41)
    m.BankPassWord = readLen(r, 41)
    m.AccountID = readLen(r, 13)
    m.Password = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    binary.Read(r, binary.BigEndian, &m.VerifyCertNoFlag)
    m.CurrencyID = readLen(r, 4)
    binary.Read(r, binary.BigEndian, &m.CashExchangeCode)
    m.Digest = readLen(r, 36)
    binary.Read(r, binary.BigEndian, &m.BankAccType)
    m.DeviceID = readLen(r, 3)
    binary.Read(r, binary.BigEndian, &m.BankSecuAccType)
    m.BrokerIDByBank = readLen(r, 33)
    m.BankSecuAcc = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.BankPwdFlag)
    binary.Read(r, binary.BigEndian, &m.SecuPwdFlag)
    m.OperNo = readLen(r, 17)
    binary.Read(r, binary.BigEndian, &m.TID)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.SecDaBeBl)
    binary.Read(r, binary.BigEndian, &m.BankChal)
    binary.Read(r, binary.BigEndian, &m.ReserveOpenAccStas)
    binary.Read(r, binary.BigEndian, &m.ErrorID)
    m.ErrorMsg = readLen(r, 81)
}

type  CThostFtdcReqResOpenAccConfirmTpdField struct {
    TradeCode string;
    BankID string;
    BankBranchID string;
    BrokerID string;
    BrokerBranchID string;
    TradeDate string;
    TradeTime string;
    BankSerial string;
    TradingDay string;
    PlateSerial int32;
    LastFragment uint8;
    SessionID int32;
    CustomerName string;
    IdCardType uint8;
    IdentifiedCardNo string;
    Gender uint8;
    CountryCode string;
    CustType uint8;
    Address string;
    ZipCode string;
    Telephone string;
    MobilePhone string;
    Fax string;
    EMail string;
    MoneyAccountStatus uint8;
    BankAccount string;
    BankPassWord string;
    AccountID string;
    Password string;
    BankAccType uint8;
    InstallID int32;
    VerifyCertNoFlag uint8;
    CurrencyID string;
    CashExchangeCode uint8;
    BrokerIDByBank string;
    BankSecuAccType uint8;
    BankSecuAcc string;
    BankPwdFlag uint8;
    SecuPwdFlag uint8;
    DeviceID string;
    OperNo string;
    UserID string;
    TID int32;
    Digest string;
}

func (m* CThostFtdcReqResOpenAccConfirmTpdField) Marshal(w io.Writer) {
    writeLen(w, m.TradeCode, 7)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBranchID, 5)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.BrokerBranchID, 31)
    writeLen(w, m.TradeDate, 9)
    writeLen(w, m.TradeTime, 9)
    writeLen(w, m.BankSerial, 13)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.PlateSerial)
    binary.Write(w, binary.BigEndian, &m.LastFragment)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.CustomerName, 51)
    binary.Write(w, binary.BigEndian, &m.IdCardType)
    writeLen(w, m.IdentifiedCardNo, 51)
    binary.Write(w, binary.BigEndian, &m.Gender)
    writeLen(w, m.CountryCode, 21)
    binary.Write(w, binary.BigEndian, &m.CustType)
    writeLen(w, m.Address, 101)
    writeLen(w, m.ZipCode, 7)
    writeLen(w, m.Telephone, 41)
    writeLen(w, m.MobilePhone, 21)
    writeLen(w, m.Fax, 41)
    writeLen(w, m.EMail, 41)
    binary.Write(w, binary.BigEndian, &m.MoneyAccountStatus)
    writeLen(w, m.BankAccount, 41)
    writeLen(w, m.BankPassWord, 41)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.Password, 41)
    binary.Write(w, binary.BigEndian, &m.BankAccType)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    binary.Write(w, binary.BigEndian, &m.VerifyCertNoFlag)
    writeLen(w, m.CurrencyID, 4)
    binary.Write(w, binary.BigEndian, &m.CashExchangeCode)
    writeLen(w, m.BrokerIDByBank, 33)
    binary.Write(w, binary.BigEndian, &m.BankSecuAccType)
    writeLen(w, m.BankSecuAcc, 41)
    binary.Write(w, binary.BigEndian, &m.BankPwdFlag)
    binary.Write(w, binary.BigEndian, &m.SecuPwdFlag)
    writeLen(w, m.DeviceID, 3)
    writeLen(w, m.OperNo, 17)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.TID)
    writeLen(w, m.Digest, 36)
}

func (m* CThostFtdcReqResOpenAccConfirmTpdField) Unmarshal(r io.Reader) {
    m.TradeCode = readLen(r, 7)
    m.BankID = readLen(r, 4)
    m.BankBranchID = readLen(r, 5)
    m.BrokerID = readLen(r, 11)
    m.BrokerBranchID = readLen(r, 31)
    m.TradeDate = readLen(r, 9)
    m.TradeTime = readLen(r, 9)
    m.BankSerial = readLen(r, 13)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.PlateSerial)
    binary.Read(r, binary.BigEndian, &m.LastFragment)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.CustomerName = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.IdCardType)
    m.IdentifiedCardNo = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.Gender)
    m.CountryCode = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.CustType)
    m.Address = readLen(r, 101)
    m.ZipCode = readLen(r, 7)
    m.Telephone = readLen(r, 41)
    m.MobilePhone = readLen(r, 21)
    m.Fax = readLen(r, 41)
    m.EMail = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.MoneyAccountStatus)
    m.BankAccount = readLen(r, 41)
    m.BankPassWord = readLen(r, 41)
    m.AccountID = readLen(r, 13)
    m.Password = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.BankAccType)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    binary.Read(r, binary.BigEndian, &m.VerifyCertNoFlag)
    m.CurrencyID = readLen(r, 4)
    binary.Read(r, binary.BigEndian, &m.CashExchangeCode)
    m.BrokerIDByBank = readLen(r, 33)
    binary.Read(r, binary.BigEndian, &m.BankSecuAccType)
    m.BankSecuAcc = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.BankPwdFlag)
    binary.Read(r, binary.BigEndian, &m.SecuPwdFlag)
    m.DeviceID = readLen(r, 3)
    m.OperNo = readLen(r, 17)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.TID)
    m.Digest = readLen(r, 36)
}

type  CThostFtdcResOpenAccConfirmTpdField struct {
    TradeCode string;
    BankID string;
    BankBranchID string;
    BrokerID string;
    BrokerBranchID string;
    TradeDate string;
    TradeTime string;
    BankSerial string;
    TradingDay string;
    PlateSerial int32;
    LastFragment uint8;
    SessionID int32;
    CustomerName string;
    IdCardType uint8;
    IdentifiedCardNo string;
    Gender uint8;
    CountryCode string;
    CustType uint8;
    Address string;
    ZipCode string;
    Telephone string;
    MobilePhone string;
    Fax string;
    EMail string;
    MoneyAccountStatus uint8;
    BankAccount string;
    BankPassWord string;
    AccountID string;
    Password string;
    BankAccType uint8;
    InstallID int32;
    VerifyCertNoFlag uint8;
    CurrencyID string;
    CashExchangeCode uint8;
    BrokerIDByBank string;
    BankSecuAccType uint8;
    BankSecuAcc string;
    BankPwdFlag uint8;
    SecuPwdFlag uint8;
    DeviceID string;
    OperNo string;
    UserID string;
    TID int32;
    Digest string;
    ErrorID int32;
    ErrorMsg string;
}

func (m* CThostFtdcResOpenAccConfirmTpdField) Marshal(w io.Writer) {
    writeLen(w, m.TradeCode, 7)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBranchID, 5)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.BrokerBranchID, 31)
    writeLen(w, m.TradeDate, 9)
    writeLen(w, m.TradeTime, 9)
    writeLen(w, m.BankSerial, 13)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.PlateSerial)
    binary.Write(w, binary.BigEndian, &m.LastFragment)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.CustomerName, 51)
    binary.Write(w, binary.BigEndian, &m.IdCardType)
    writeLen(w, m.IdentifiedCardNo, 51)
    binary.Write(w, binary.BigEndian, &m.Gender)
    writeLen(w, m.CountryCode, 21)
    binary.Write(w, binary.BigEndian, &m.CustType)
    writeLen(w, m.Address, 101)
    writeLen(w, m.ZipCode, 7)
    writeLen(w, m.Telephone, 41)
    writeLen(w, m.MobilePhone, 21)
    writeLen(w, m.Fax, 41)
    writeLen(w, m.EMail, 41)
    binary.Write(w, binary.BigEndian, &m.MoneyAccountStatus)
    writeLen(w, m.BankAccount, 41)
    writeLen(w, m.BankPassWord, 41)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.Password, 41)
    binary.Write(w, binary.BigEndian, &m.BankAccType)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    binary.Write(w, binary.BigEndian, &m.VerifyCertNoFlag)
    writeLen(w, m.CurrencyID, 4)
    binary.Write(w, binary.BigEndian, &m.CashExchangeCode)
    writeLen(w, m.BrokerIDByBank, 33)
    binary.Write(w, binary.BigEndian, &m.BankSecuAccType)
    writeLen(w, m.BankSecuAcc, 41)
    binary.Write(w, binary.BigEndian, &m.BankPwdFlag)
    binary.Write(w, binary.BigEndian, &m.SecuPwdFlag)
    writeLen(w, m.DeviceID, 3)
    writeLen(w, m.OperNo, 17)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.TID)
    writeLen(w, m.Digest, 36)
    binary.Write(w, binary.BigEndian, &m.ErrorID)
    writeLen(w, m.ErrorMsg, 81)
}

func (m* CThostFtdcResOpenAccConfirmTpdField) Unmarshal(r io.Reader) {
    m.TradeCode = readLen(r, 7)
    m.BankID = readLen(r, 4)
    m.BankBranchID = readLen(r, 5)
    m.BrokerID = readLen(r, 11)
    m.BrokerBranchID = readLen(r, 31)
    m.TradeDate = readLen(r, 9)
    m.TradeTime = readLen(r, 9)
    m.BankSerial = readLen(r, 13)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.PlateSerial)
    binary.Read(r, binary.BigEndian, &m.LastFragment)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.CustomerName = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.IdCardType)
    m.IdentifiedCardNo = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.Gender)
    m.CountryCode = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.CustType)
    m.Address = readLen(r, 101)
    m.ZipCode = readLen(r, 7)
    m.Telephone = readLen(r, 41)
    m.MobilePhone = readLen(r, 21)
    m.Fax = readLen(r, 41)
    m.EMail = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.MoneyAccountStatus)
    m.BankAccount = readLen(r, 41)
    m.BankPassWord = readLen(r, 41)
    m.AccountID = readLen(r, 13)
    m.Password = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.BankAccType)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    binary.Read(r, binary.BigEndian, &m.VerifyCertNoFlag)
    m.CurrencyID = readLen(r, 4)
    binary.Read(r, binary.BigEndian, &m.CashExchangeCode)
    m.BrokerIDByBank = readLen(r, 33)
    binary.Read(r, binary.BigEndian, &m.BankSecuAccType)
    m.BankSecuAcc = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.BankPwdFlag)
    binary.Read(r, binary.BigEndian, &m.SecuPwdFlag)
    m.DeviceID = readLen(r, 3)
    m.OperNo = readLen(r, 17)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.TID)
    m.Digest = readLen(r, 36)
    binary.Read(r, binary.BigEndian, &m.ErrorID)
    m.ErrorMsg = readLen(r, 81)
}

type  CThostFtdcReqSecuritiesDepositInterestField struct {
    TradeCode string;
    BankID string;
    BankBranchID string;
    BrokerID string;
    BrokerBranchID string;
    TradeDate string;
    TradeTime string;
    BankSerial string;
    TradingDay string;
    PlateSerial int32;
    LastFragment uint8;
    SessionID int32;
    CustomerName string;
    IdCardType uint8;
    IdentifiedCardNo string;
    Gender uint8;
    CountryCode string;
    CustType uint8;
    Address string;
    ZipCode string;
    Telephone string;
    MobilePhone string;
    Fax string;
    EMail string;
    MoneyAccountStatus uint8;
    BankAccount string;
    BankPassWord string;
    AccountID string;
    Password string;
    InstallID int32;
    FutureSerial int32;
    VerifyCertNoFlag uint8;
    CurrencyID string;
    CashExchangeCode uint8;
    Digest string;
    BankAccType uint8;
    DeviceID string;
    BankSecuAccType uint8;
    BrokerIDByBank string;
    BankSecuAcc string;
    BankPwdFlag uint8;
    SecuPwdFlag uint8;
    OperNo string;
    TID int32;
    UserID string;
    DepositInterest float64;
}

func (m* CThostFtdcReqSecuritiesDepositInterestField) Marshal(w io.Writer) {
    writeLen(w, m.TradeCode, 7)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBranchID, 5)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.BrokerBranchID, 31)
    writeLen(w, m.TradeDate, 9)
    writeLen(w, m.TradeTime, 9)
    writeLen(w, m.BankSerial, 13)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.PlateSerial)
    binary.Write(w, binary.BigEndian, &m.LastFragment)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.CustomerName, 51)
    binary.Write(w, binary.BigEndian, &m.IdCardType)
    writeLen(w, m.IdentifiedCardNo, 51)
    binary.Write(w, binary.BigEndian, &m.Gender)
    writeLen(w, m.CountryCode, 21)
    binary.Write(w, binary.BigEndian, &m.CustType)
    writeLen(w, m.Address, 101)
    writeLen(w, m.ZipCode, 7)
    writeLen(w, m.Telephone, 41)
    writeLen(w, m.MobilePhone, 21)
    writeLen(w, m.Fax, 41)
    writeLen(w, m.EMail, 41)
    binary.Write(w, binary.BigEndian, &m.MoneyAccountStatus)
    writeLen(w, m.BankAccount, 41)
    writeLen(w, m.BankPassWord, 41)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.Password, 41)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    binary.Write(w, binary.BigEndian, &m.FutureSerial)
    binary.Write(w, binary.BigEndian, &m.VerifyCertNoFlag)
    writeLen(w, m.CurrencyID, 4)
    binary.Write(w, binary.BigEndian, &m.CashExchangeCode)
    writeLen(w, m.Digest, 36)
    binary.Write(w, binary.BigEndian, &m.BankAccType)
    writeLen(w, m.DeviceID, 3)
    binary.Write(w, binary.BigEndian, &m.BankSecuAccType)
    writeLen(w, m.BrokerIDByBank, 33)
    writeLen(w, m.BankSecuAcc, 41)
    binary.Write(w, binary.BigEndian, &m.BankPwdFlag)
    binary.Write(w, binary.BigEndian, &m.SecuPwdFlag)
    writeLen(w, m.OperNo, 17)
    binary.Write(w, binary.BigEndian, &m.TID)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.DepositInterest)
}

func (m* CThostFtdcReqSecuritiesDepositInterestField) Unmarshal(r io.Reader) {
    m.TradeCode = readLen(r, 7)
    m.BankID = readLen(r, 4)
    m.BankBranchID = readLen(r, 5)
    m.BrokerID = readLen(r, 11)
    m.BrokerBranchID = readLen(r, 31)
    m.TradeDate = readLen(r, 9)
    m.TradeTime = readLen(r, 9)
    m.BankSerial = readLen(r, 13)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.PlateSerial)
    binary.Read(r, binary.BigEndian, &m.LastFragment)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.CustomerName = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.IdCardType)
    m.IdentifiedCardNo = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.Gender)
    m.CountryCode = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.CustType)
    m.Address = readLen(r, 101)
    m.ZipCode = readLen(r, 7)
    m.Telephone = readLen(r, 41)
    m.MobilePhone = readLen(r, 21)
    m.Fax = readLen(r, 41)
    m.EMail = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.MoneyAccountStatus)
    m.BankAccount = readLen(r, 41)
    m.BankPassWord = readLen(r, 41)
    m.AccountID = readLen(r, 13)
    m.Password = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    binary.Read(r, binary.BigEndian, &m.FutureSerial)
    binary.Read(r, binary.BigEndian, &m.VerifyCertNoFlag)
    m.CurrencyID = readLen(r, 4)
    binary.Read(r, binary.BigEndian, &m.CashExchangeCode)
    m.Digest = readLen(r, 36)
    binary.Read(r, binary.BigEndian, &m.BankAccType)
    m.DeviceID = readLen(r, 3)
    binary.Read(r, binary.BigEndian, &m.BankSecuAccType)
    m.BrokerIDByBank = readLen(r, 33)
    m.BankSecuAcc = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.BankPwdFlag)
    binary.Read(r, binary.BigEndian, &m.SecuPwdFlag)
    m.OperNo = readLen(r, 17)
    binary.Read(r, binary.BigEndian, &m.TID)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.DepositInterest)
}

type  CThostFtdcSecuritiesDepositInterestField struct {
    TradeCode string;
    BankID string;
    BankBranchID string;
    BrokerID string;
    BrokerBranchID string;
    TradeDate string;
    TradeTime string;
    BankSerial string;
    TradingDay string;
    PlateSerial int32;
    LastFragment uint8;
    SessionID int32;
    CustomerName string;
    IdCardType uint8;
    IdentifiedCardNo string;
    Gender uint8;
    CountryCode string;
    CustType uint8;
    Address string;
    ZipCode string;
    Telephone string;
    MobilePhone string;
    Fax string;
    EMail string;
    MoneyAccountStatus uint8;
    BankAccount string;
    BankPassWord string;
    AccountID string;
    Password string;
    InstallID int32;
    FutureSerial int32;
    VerifyCertNoFlag uint8;
    CurrencyID string;
    CashExchangeCode uint8;
    Digest string;
    BankAccType uint8;
    DeviceID string;
    BankSecuAccType uint8;
    BrokerIDByBank string;
    BankSecuAcc string;
    BankPwdFlag uint8;
    SecuPwdFlag uint8;
    OperNo string;
    TID int32;
    UserID string;
    DepositInterest float64;
    InterestType uint8;
    ErrorID int32;
    ErrorMsg string;
}

func (m* CThostFtdcSecuritiesDepositInterestField) Marshal(w io.Writer) {
    writeLen(w, m.TradeCode, 7)
    writeLen(w, m.BankID, 4)
    writeLen(w, m.BankBranchID, 5)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.BrokerBranchID, 31)
    writeLen(w, m.TradeDate, 9)
    writeLen(w, m.TradeTime, 9)
    writeLen(w, m.BankSerial, 13)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.PlateSerial)
    binary.Write(w, binary.BigEndian, &m.LastFragment)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.CustomerName, 51)
    binary.Write(w, binary.BigEndian, &m.IdCardType)
    writeLen(w, m.IdentifiedCardNo, 51)
    binary.Write(w, binary.BigEndian, &m.Gender)
    writeLen(w, m.CountryCode, 21)
    binary.Write(w, binary.BigEndian, &m.CustType)
    writeLen(w, m.Address, 101)
    writeLen(w, m.ZipCode, 7)
    writeLen(w, m.Telephone, 41)
    writeLen(w, m.MobilePhone, 21)
    writeLen(w, m.Fax, 41)
    writeLen(w, m.EMail, 41)
    binary.Write(w, binary.BigEndian, &m.MoneyAccountStatus)
    writeLen(w, m.BankAccount, 41)
    writeLen(w, m.BankPassWord, 41)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.Password, 41)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    binary.Write(w, binary.BigEndian, &m.FutureSerial)
    binary.Write(w, binary.BigEndian, &m.VerifyCertNoFlag)
    writeLen(w, m.CurrencyID, 4)
    binary.Write(w, binary.BigEndian, &m.CashExchangeCode)
    writeLen(w, m.Digest, 36)
    binary.Write(w, binary.BigEndian, &m.BankAccType)
    writeLen(w, m.DeviceID, 3)
    binary.Write(w, binary.BigEndian, &m.BankSecuAccType)
    writeLen(w, m.BrokerIDByBank, 33)
    writeLen(w, m.BankSecuAcc, 41)
    binary.Write(w, binary.BigEndian, &m.BankPwdFlag)
    binary.Write(w, binary.BigEndian, &m.SecuPwdFlag)
    writeLen(w, m.OperNo, 17)
    binary.Write(w, binary.BigEndian, &m.TID)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.DepositInterest)
    binary.Write(w, binary.BigEndian, &m.InterestType)
    binary.Write(w, binary.BigEndian, &m.ErrorID)
    writeLen(w, m.ErrorMsg, 81)
}

func (m* CThostFtdcSecuritiesDepositInterestField) Unmarshal(r io.Reader) {
    m.TradeCode = readLen(r, 7)
    m.BankID = readLen(r, 4)
    m.BankBranchID = readLen(r, 5)
    m.BrokerID = readLen(r, 11)
    m.BrokerBranchID = readLen(r, 31)
    m.TradeDate = readLen(r, 9)
    m.TradeTime = readLen(r, 9)
    m.BankSerial = readLen(r, 13)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.PlateSerial)
    binary.Read(r, binary.BigEndian, &m.LastFragment)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.CustomerName = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.IdCardType)
    m.IdentifiedCardNo = readLen(r, 51)
    binary.Read(r, binary.BigEndian, &m.Gender)
    m.CountryCode = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.CustType)
    m.Address = readLen(r, 101)
    m.ZipCode = readLen(r, 7)
    m.Telephone = readLen(r, 41)
    m.MobilePhone = readLen(r, 21)
    m.Fax = readLen(r, 41)
    m.EMail = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.MoneyAccountStatus)
    m.BankAccount = readLen(r, 41)
    m.BankPassWord = readLen(r, 41)
    m.AccountID = readLen(r, 13)
    m.Password = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    binary.Read(r, binary.BigEndian, &m.FutureSerial)
    binary.Read(r, binary.BigEndian, &m.VerifyCertNoFlag)
    m.CurrencyID = readLen(r, 4)
    binary.Read(r, binary.BigEndian, &m.CashExchangeCode)
    m.Digest = readLen(r, 36)
    binary.Read(r, binary.BigEndian, &m.BankAccType)
    m.DeviceID = readLen(r, 3)
    binary.Read(r, binary.BigEndian, &m.BankSecuAccType)
    m.BrokerIDByBank = readLen(r, 33)
    m.BankSecuAcc = readLen(r, 41)
    binary.Read(r, binary.BigEndian, &m.BankPwdFlag)
    binary.Read(r, binary.BigEndian, &m.SecuPwdFlag)
    m.OperNo = readLen(r, 17)
    binary.Read(r, binary.BigEndian, &m.TID)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.DepositInterest)
    binary.Read(r, binary.BigEndian, &m.InterestType)
    binary.Read(r, binary.BigEndian, &m.ErrorID)
    m.ErrorMsg = readLen(r, 81)
}

type  CThostFtdcBrokerBreakSectionField struct {
    ExchangeID string;
    BrokerID string;
    TimeStart string;
    TimeEnd string;
    SequenceNo int32;
}

func (m* CThostFtdcBrokerBreakSectionField) Marshal(w io.Writer) {
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.TimeStart, 9)
    writeLen(w, m.TimeEnd, 9)
    binary.Write(w, binary.BigEndian, &m.SequenceNo)
}

func (m* CThostFtdcBrokerBreakSectionField) Unmarshal(r io.Reader) {
    m.ExchangeID = readLen(r, 9)
    m.BrokerID = readLen(r, 11)
    m.TimeStart = readLen(r, 9)
    m.TimeEnd = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.SequenceNo)
}

type  CThostFtdcBrokerLockInvestorStockField struct {
    ExchangeID string;
    BrokerID string;
    InvestorID string;
    InstrumentID string;
    Volume int32;
}

func (m* CThostFtdcBrokerLockInvestorStockField) Marshal(w io.Writer) {
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
    binary.Write(w, binary.BigEndian, &m.Volume)
}

func (m* CThostFtdcBrokerLockInvestorStockField) Unmarshal(r io.Reader) {
    m.ExchangeID = readLen(r, 9)
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.Volume)
}

type  CThostFtdcQryBrokerLockInvestorStockField struct {
    ExchangeID string;
    BrokerID string;
    InvestorID string;
    InstrumentID string;
}

func (m* CThostFtdcQryBrokerLockInvestorStockField) Marshal(w io.Writer) {
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.InstrumentID, 31)
}

func (m* CThostFtdcQryBrokerLockInvestorStockField) Unmarshal(r io.Reader) {
    m.ExchangeID = readLen(r, 9)
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.InstrumentID = readLen(r, 31)
}

type  CThostFtdcExecIsCheckUnderlyingField struct {
    ExchangeID string;
    UnderlyingInstrID string;
    IsCheck int32;
}

func (m* CThostFtdcExecIsCheckUnderlyingField) Marshal(w io.Writer) {
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.UnderlyingInstrID, 31)
    binary.Write(w, binary.BigEndian, &m.IsCheck)
}

func (m* CThostFtdcExecIsCheckUnderlyingField) Unmarshal(r io.Reader) {
    m.ExchangeID = readLen(r, 9)
    m.UnderlyingInstrID = readLen(r, 31)
    binary.Read(r, binary.BigEndian, &m.IsCheck)
}

type  CThostFtdcQryExecIsCheckUnderlyingField struct {
    ExchangeID string;
    UnderlyingInstrID string;
}

func (m* CThostFtdcQryExecIsCheckUnderlyingField) Marshal(w io.Writer) {
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.UnderlyingInstrID, 31)
}

func (m* CThostFtdcQryExecIsCheckUnderlyingField) Unmarshal(r io.Reader) {
    m.ExchangeID = readLen(r, 9)
    m.UnderlyingInstrID = readLen(r, 31)
}

type  CThostFtdcPBUInvestorMapField struct {
    BrokerID string;
    ExchangeID string;
    InvestorID string;
    PBU string;
    ClientID string;
    OperationDate string;
    OperationTime string;
}

func (m* CThostFtdcPBUInvestorMapField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.PBU, 21)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.OperationDate, 9)
    writeLen(w, m.OperationTime, 9)
}

func (m* CThostFtdcPBUInvestorMapField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.ExchangeID = readLen(r, 9)
    m.InvestorID = readLen(r, 13)
    m.PBU = readLen(r, 21)
    m.ClientID = readLen(r, 11)
    m.OperationDate = readLen(r, 9)
    m.OperationTime = readLen(r, 9)
}

type  CThostFtdcInputExecCombineOrderField struct {
    BrokerID string;
    InvestorID string;
    CallInstrumentID string;
    PutInstrumentID string;
    ExecCombineOrderRef string;
    UserID string;
    Volume int32;
    RequestID int32;
    BusinessUnit string;
    ActionType uint8;
    ExchangeID string;
    InvestUnitID string;
    ClientID string;
    IPAddress string;
    MacAddress string;
}

func (m* CThostFtdcInputExecCombineOrderField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.CallInstrumentID, 31)
    writeLen(w, m.PutInstrumentID, 31)
    writeLen(w, m.ExecCombineOrderRef, 13)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.Volume)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    writeLen(w, m.BusinessUnit, 21)
    binary.Write(w, binary.BigEndian, &m.ActionType)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.InvestUnitID, 17)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
}

func (m* CThostFtdcInputExecCombineOrderField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.CallInstrumentID = readLen(r, 31)
    m.PutInstrumentID = readLen(r, 31)
    m.ExecCombineOrderRef = readLen(r, 13)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.Volume)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    m.BusinessUnit = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.ActionType)
    m.ExchangeID = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
    m.ClientID = readLen(r, 11)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
}

type  CThostFtdcInputExecCombineOrderActionField struct {
    BrokerID string;
    InvestorID string;
    ExecCombineOrderActionRef int32;
    ExecCombineOrderRef string;
    RequestID int32;
    FrontID int32;
    SessionID int32;
    ExchangeID string;
    ExecCombineOrderSysID string;
    ActionFlag uint8;
    UserID string;
    InvestUnitID string;
    IPAddress string;
    MacAddress string;
}

func (m* CThostFtdcInputExecCombineOrderActionField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.ExecCombineOrderActionRef)
    writeLen(w, m.ExecCombineOrderRef, 13)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    binary.Write(w, binary.BigEndian, &m.FrontID)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ExecCombineOrderSysID, 21)
    binary.Write(w, binary.BigEndian, &m.ActionFlag)
    writeLen(w, m.UserID, 16)
    writeLen(w, m.InvestUnitID, 17)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
}

func (m* CThostFtdcInputExecCombineOrderActionField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.ExecCombineOrderActionRef)
    m.ExecCombineOrderRef = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    binary.Read(r, binary.BigEndian, &m.FrontID)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.ExchangeID = readLen(r, 9)
    m.ExecCombineOrderSysID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.ActionFlag)
    m.UserID = readLen(r, 16)
    m.InvestUnitID = readLen(r, 17)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
}

type  CThostFtdcExecCombineOrderField struct {
    BrokerID string;
    InvestorID string;
    CallInstrumentID string;
    PutInstrumentID string;
    ExecCombineOrderRef string;
    UserID string;
    Volume int32;
    RequestID int32;
    BusinessUnit string;
    ActionType uint8;
    ExecCombineOrderLocalID string;
    ExchangeID string;
    ParticipantID string;
    ClientID string;
    UnderlyingInstID string;
    ExchangeCallInstID string;
    ExchangePutInstID string;
    TraderID string;
    InstallID int32;
    OrderSubmitStatus uint8;
    NotifySequence int32;
    TradingDay string;
    SettlementID int32;
    ExecCombineOrderSysID string;
    InsertDate string;
    InsertTime string;
    CancelTime string;
    ExecResult uint8;
    ClearingPartID string;
    SequenceNo int32;
    FrontID int32;
    SessionID int32;
    UserProductInfo string;
    StatusMsg string;
    ActiveUserID string;
    BrokerExecCombineOrderSeq int32;
    BranchID string;
    InvestUnitID string;
    AccountID string;
    CurrencyID string;
    IPAddress string;
    MacAddress string;
}

func (m* CThostFtdcExecCombineOrderField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.CallInstrumentID, 31)
    writeLen(w, m.PutInstrumentID, 31)
    writeLen(w, m.ExecCombineOrderRef, 13)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.Volume)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    writeLen(w, m.BusinessUnit, 21)
    binary.Write(w, binary.BigEndian, &m.ActionType)
    writeLen(w, m.ExecCombineOrderLocalID, 13)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.UnderlyingInstID, 31)
    writeLen(w, m.ExchangeCallInstID, 31)
    writeLen(w, m.ExchangePutInstID, 31)
    writeLen(w, m.TraderID, 21)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    binary.Write(w, binary.BigEndian, &m.OrderSubmitStatus)
    binary.Write(w, binary.BigEndian, &m.NotifySequence)
    writeLen(w, m.TradingDay, 9)
    binary.Write(w, binary.BigEndian, &m.SettlementID)
    writeLen(w, m.ExecCombineOrderSysID, 21)
    writeLen(w, m.InsertDate, 9)
    writeLen(w, m.InsertTime, 9)
    writeLen(w, m.CancelTime, 9)
    binary.Write(w, binary.BigEndian, &m.ExecResult)
    writeLen(w, m.ClearingPartID, 11)
    binary.Write(w, binary.BigEndian, &m.SequenceNo)
    binary.Write(w, binary.BigEndian, &m.FrontID)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.UserProductInfo, 11)
    writeLen(w, m.StatusMsg, 81)
    writeLen(w, m.ActiveUserID, 16)
    binary.Write(w, binary.BigEndian, &m.BrokerExecCombineOrderSeq)
    writeLen(w, m.BranchID, 9)
    writeLen(w, m.InvestUnitID, 17)
    writeLen(w, m.AccountID, 13)
    writeLen(w, m.CurrencyID, 4)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
}

func (m* CThostFtdcExecCombineOrderField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.CallInstrumentID = readLen(r, 31)
    m.PutInstrumentID = readLen(r, 31)
    m.ExecCombineOrderRef = readLen(r, 13)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.Volume)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    m.BusinessUnit = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.ActionType)
    m.ExecCombineOrderLocalID = readLen(r, 13)
    m.ExchangeID = readLen(r, 9)
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.UnderlyingInstID = readLen(r, 31)
    m.ExchangeCallInstID = readLen(r, 31)
    m.ExchangePutInstID = readLen(r, 31)
    m.TraderID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    binary.Read(r, binary.BigEndian, &m.OrderSubmitStatus)
    binary.Read(r, binary.BigEndian, &m.NotifySequence)
    m.TradingDay = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.SettlementID)
    m.ExecCombineOrderSysID = readLen(r, 21)
    m.InsertDate = readLen(r, 9)
    m.InsertTime = readLen(r, 9)
    m.CancelTime = readLen(r, 9)
    binary.Read(r, binary.BigEndian, &m.ExecResult)
    m.ClearingPartID = readLen(r, 11)
    binary.Read(r, binary.BigEndian, &m.SequenceNo)
    binary.Read(r, binary.BigEndian, &m.FrontID)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.UserProductInfo = readLen(r, 11)
    m.StatusMsg = readLen(r, 81)
    m.ActiveUserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.BrokerExecCombineOrderSeq)
    m.BranchID = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
    m.AccountID = readLen(r, 13)
    m.CurrencyID = readLen(r, 4)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
}

type  CThostFtdcExecCombineOrderActionField struct {
    BrokerID string;
    InvestorID string;
    ExecCombineOrderActionRef int32;
    ExecCombineOrderRef string;
    RequestID int32;
    FrontID int32;
    SessionID int32;
    ExchangeID string;
    ExecCombineOrderSysID string;
    ActionFlag uint8;
    ActionDate string;
    ActionTime string;
    TraderID string;
    InstallID int32;
    ExecCombineOrderLocalID string;
    ActionLocalID string;
    ParticipantID string;
    ClientID string;
    BusinessUnit string;
    OrderActionStatus uint8;
    UserID string;
    ActionType uint8;
    StatusMsg string;
    CallInstrumentID string;
    PutInstrumentID string;
    BranchID string;
    InvestUnitID string;
    IPAddress string;
    MacAddress string;
}

func (m* CThostFtdcExecCombineOrderActionField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    binary.Write(w, binary.BigEndian, &m.ExecCombineOrderActionRef)
    writeLen(w, m.ExecCombineOrderRef, 13)
    binary.Write(w, binary.BigEndian, &m.RequestID)
    binary.Write(w, binary.BigEndian, &m.FrontID)
    binary.Write(w, binary.BigEndian, &m.SessionID)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ExecCombineOrderSysID, 21)
    binary.Write(w, binary.BigEndian, &m.ActionFlag)
    writeLen(w, m.ActionDate, 9)
    writeLen(w, m.ActionTime, 9)
    writeLen(w, m.TraderID, 21)
    binary.Write(w, binary.BigEndian, &m.InstallID)
    writeLen(w, m.ExecCombineOrderLocalID, 13)
    writeLen(w, m.ActionLocalID, 13)
    writeLen(w, m.ParticipantID, 11)
    writeLen(w, m.ClientID, 11)
    writeLen(w, m.BusinessUnit, 21)
    binary.Write(w, binary.BigEndian, &m.OrderActionStatus)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.ActionType)
    writeLen(w, m.StatusMsg, 81)
    writeLen(w, m.CallInstrumentID, 31)
    writeLen(w, m.PutInstrumentID, 31)
    writeLen(w, m.BranchID, 9)
    writeLen(w, m.InvestUnitID, 17)
    writeLen(w, m.IPAddress, 16)
    writeLen(w, m.MacAddress, 21)
}

func (m* CThostFtdcExecCombineOrderActionField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.ExecCombineOrderActionRef)
    m.ExecCombineOrderRef = readLen(r, 13)
    binary.Read(r, binary.BigEndian, &m.RequestID)
    binary.Read(r, binary.BigEndian, &m.FrontID)
    binary.Read(r, binary.BigEndian, &m.SessionID)
    m.ExchangeID = readLen(r, 9)
    m.ExecCombineOrderSysID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.ActionFlag)
    m.ActionDate = readLen(r, 9)
    m.ActionTime = readLen(r, 9)
    m.TraderID = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.InstallID)
    m.ExecCombineOrderLocalID = readLen(r, 13)
    m.ActionLocalID = readLen(r, 13)
    m.ParticipantID = readLen(r, 11)
    m.ClientID = readLen(r, 11)
    m.BusinessUnit = readLen(r, 21)
    binary.Read(r, binary.BigEndian, &m.OrderActionStatus)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.ActionType)
    m.StatusMsg = readLen(r, 81)
    m.CallInstrumentID = readLen(r, 31)
    m.PutInstrumentID = readLen(r, 31)
    m.BranchID = readLen(r, 9)
    m.InvestUnitID = readLen(r, 17)
    m.IPAddress = readLen(r, 16)
    m.MacAddress = readLen(r, 21)
}

type  CThostFtdcQryExecCombineOrderField struct {
    BrokerID string;
    InvestorID string;
    CallInstrumentID string;
    PutInstrumentID string;
    ExchangeID string;
    ExecCombineOrderSysID string;
    InsertTimeStart string;
    InsertTimeEnd string;
}

func (m* CThostFtdcQryExecCombineOrderField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.InvestorID, 13)
    writeLen(w, m.CallInstrumentID, 31)
    writeLen(w, m.PutInstrumentID, 31)
    writeLen(w, m.ExchangeID, 9)
    writeLen(w, m.ExecCombineOrderSysID, 21)
    writeLen(w, m.InsertTimeStart, 9)
    writeLen(w, m.InsertTimeEnd, 9)
}

func (m* CThostFtdcQryExecCombineOrderField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.InvestorID = readLen(r, 13)
    m.CallInstrumentID = readLen(r, 31)
    m.PutInstrumentID = readLen(r, 31)
    m.ExchangeID = readLen(r, 9)
    m.ExecCombineOrderSysID = readLen(r, 21)
    m.InsertTimeStart = readLen(r, 9)
    m.InsertTimeEnd = readLen(r, 9)
}

type  CThostFtdcUserSystemInfoField struct {
    BrokerID string;
    UserID string;
    ClientSystemInfoLen int32;
    ClientSystemInfo string;
    ClientPublicIP string;
    ClientIPPort int32;
    ClientLoginTime string;
    ClientAppID string;
}

func (m* CThostFtdcUserSystemInfoField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.UserID, 16)
    binary.Write(w, binary.BigEndian, &m.ClientSystemInfoLen)
    writeLen(w, m.ClientSystemInfo, 273)
    writeLen(w, m.ClientPublicIP, 16)
    binary.Write(w, binary.BigEndian, &m.ClientIPPort)
    writeLen(w, m.ClientLoginTime, 9)
    writeLen(w, m.ClientAppID, 33)
}

func (m* CThostFtdcUserSystemInfoField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.UserID = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.ClientSystemInfoLen)
    m.ClientSystemInfo = readLen(r, 273)
    m.ClientPublicIP = readLen(r, 16)
    binary.Read(r, binary.BigEndian, &m.ClientIPPort)
    m.ClientLoginTime = readLen(r, 9)
    m.ClientAppID = readLen(r, 33)
}

type  CThostFtdcReqApiHandshakeField struct {
    CryptoKeyVersion string;
}

func (m* CThostFtdcReqApiHandshakeField) Marshal(w io.Writer) {
    writeLen(w, m.CryptoKeyVersion, 31)
}

func (m* CThostFtdcReqApiHandshakeField) Unmarshal(r io.Reader) {
    m.CryptoKeyVersion = readLen(r, 31)
}

type  CThostFtdcRspApiHandshakeField struct {
    FrontHandshakeDataLen int32;
    FrontHandshakeData string;
    IsApiAuthEnabled int32;
}

func (m* CThostFtdcRspApiHandshakeField) Marshal(w io.Writer) {
    binary.Write(w, binary.BigEndian, &m.FrontHandshakeDataLen)
    writeLen(w, m.FrontHandshakeData, 301)
    binary.Write(w, binary.BigEndian, &m.IsApiAuthEnabled)
}

func (m* CThostFtdcRspApiHandshakeField) Unmarshal(r io.Reader) {
    binary.Read(r, binary.BigEndian, &m.FrontHandshakeDataLen)
    m.FrontHandshakeData = readLen(r, 301)
    binary.Read(r, binary.BigEndian, &m.IsApiAuthEnabled)
}

type  CThostFtdcReqVerifyApiKeyField struct {
    ApiHandshakeDataLen int32;
    ApiHandshakeData string;
}

func (m* CThostFtdcReqVerifyApiKeyField) Marshal(w io.Writer) {
    binary.Write(w, binary.BigEndian, &m.ApiHandshakeDataLen)
    writeLen(w, m.ApiHandshakeData, 301)
}

func (m* CThostFtdcReqVerifyApiKeyField) Unmarshal(r io.Reader) {
    binary.Read(r, binary.BigEndian, &m.ApiHandshakeDataLen)
    m.ApiHandshakeData = readLen(r, 301)
}

type  CThostFtdcAppIDAuthAssignField struct {
    BrokerID string;
    AppID string;
    DRIdentityID int32;
}

func (m* CThostFtdcAppIDAuthAssignField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.AppID, 33)
    binary.Write(w, binary.BigEndian, &m.DRIdentityID)
}

func (m* CThostFtdcAppIDAuthAssignField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.AppID = readLen(r, 33)
    binary.Read(r, binary.BigEndian, &m.DRIdentityID)
}

type  CThostFtdcAppAuthenticationCodeField struct {
    BrokerID string;
    AppID string;
    AuthCode string;
    PreAuthCode string;
    AppType uint8;
}

func (m* CThostFtdcAppAuthenticationCodeField) Marshal(w io.Writer) {
    writeLen(w, m.BrokerID, 11)
    writeLen(w, m.AppID, 33)
    writeLen(w, m.AuthCode, 17)
    writeLen(w, m.PreAuthCode, 17)
    binary.Write(w, binary.BigEndian, &m.AppType)
}

func (m* CThostFtdcAppAuthenticationCodeField) Unmarshal(r io.Reader) {
    m.BrokerID = readLen(r, 11)
    m.AppID = readLen(r, 33)
    m.AuthCode = readLen(r, 17)
    m.PreAuthCode = readLen(r, 17)
    binary.Read(r, binary.BigEndian, &m.AppType)
}

type  CThostFtdcQueryFreqField struct {
    QueryFreq int32;
}

func (m* CThostFtdcQueryFreqField) Marshal(w io.Writer) {
    binary.Write(w, binary.BigEndian, &m.QueryFreq)
}

func (m* CThostFtdcQueryFreqField) Unmarshal(r io.Reader) {
    binary.Read(r, binary.BigEndian, &m.QueryFreq)
}

