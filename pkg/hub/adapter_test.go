package hub

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"
	"time"

	"github.com/mineralres/protos/src/go/ctp"
	"github.com/mineralres/goshare/pkg/util"
	gspb "github.com/mineralres/protos/src/go/goshare"
	hubpb "github.com/mineralres/protos/src/go/hub"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

type config struct {
	Account  string   `json:"account"`
	Password string   `json:"password"`
	BrokerID string   `json:"brokerId"`
	AppID    string   `json:"appId"`
	AuthCode string   `json:"authCode"`
	Fronts   []string `json:"fronts"`
}

func loadConfig(f string, out interface{}) error {
	data, err := ioutil.ReadFile(f)
	if err != nil {
		log.Println(err)
		return err
	}
	return json.Unmarshal(data, &out)
}

var (
	timeout = time.Second * 10
)

func Test_hub(t *testing.T) {
	opt := &DemoEnvOptions{}
	opt.OnDemoOrder = func(do *hubpb.DemoOrder) {

	}
	opt.OnDemoTrade = func(trade *gspb.Trade) {

	}
	opt.GetUID = func() string {
		return ""
	}
	NewDemoEnv(opt)
}

func Test_sync(t *testing.T) {
	return
	var c config
	err := loadConfig("config.json", &c)
	if err != nil {
		panic("需要自己在config.json里配置账号密码，前置地址等")
	}
	log.Println("c", c)
	var requestID int32
	adapter, err := NewSyncAdapter("47.100.1.102:8205", timeout, c.Fronts, func(pkt *Packet) {
		log.Println(ctp.CtpMessageType(pkt.MsgType))
	})
	if err != nil {
		panic(err)
	}
	{
		// 认证
		var req ctp.CThostFtdcReqAuthenticateField
		req.BrokerID = c.BrokerID
		req.AppID = c.AppID
		req.AuthCode = c.AuthCode
		req.UserID = c.Account
		requestID++
		ret, err := adapter.Send(int32(ctp.CtpMessageType_TD_ReqAuthenticate), &req, requestID, timeout)
		if err != nil {
			panic(err)
		}
		if len(ret) != 1 {
			panic("should be 1")
		}
		if ret[0].MsgType != int32(ctp.CtpMessageType_TD_OnRspAuthenticate) {
			panic("")
		}
		var rsp ctp.CThostFtdcRspAuthenticateField
		var rspInfo ctp.CThostFtdcRspInfoField
		if err := ret[0].Get2(&rsp, &rspInfo); err != nil {
			panic(err)
		}
		log.Println(req, rspInfo.ErrorID, util.StringFromGBK2(rspInfo.ErrorMsg))
	}
	// 登陆
	{
		var req ctp.CThostFtdcReqUserLoginField
		req.BrokerID = c.BrokerID
		req.UserID = c.Account
		req.Password = c.Password
		requestID++
		ret, err := adapter.Send(int32(ctp.CtpMessageType_TD_ReqUserLogin), &req, requestID, timeout)
		if err != nil || len(ret) == 0 {
			panic(err)
		}
		var rsp ctp.CThostFtdcRspUserLoginField
		var rspInfo ctp.CThostFtdcRspInfoField
		if err := ret[0].Get2(&rsp, &rspInfo); err == nil {
			log.Println(rsp, util.StringFromGBK2(rspInfo.ErrorMsg))
		}
	}
	// 查询资金
	{
		var req ctp.CThostFtdcQryTradingAccountField
		req.BrokerID = c.BrokerID
		req.InvestorID = c.Account
		requestID++
		ret, err := adapter.Send(int32(ctp.CtpMessageType_TD_ReqQryTradingAccount), &req, requestID, timeout)
		if err != nil || len(ret) == 0 {
			panic(err)
		}
		var rsp ctp.CThostFtdcTradingAccountField
		var rspInfo ctp.CThostFtdcRspInfoField
		if err := ret[0].Get2(&rsp, &rspInfo); err == nil {
			log.Println(rsp, util.StringFromGBK2(rspInfo.ErrorMsg))
		}
	}
	// 查询合约
	{
		var req ctp.CThostFtdcQryInstrumentField
		requestID++
		ret, err := adapter.Send(int32(ctp.CtpMessageType_TD_ReqQryInstrument), &req, requestID, timeout)
		if err != nil || len(ret) == 0 {
			panic(err)
		}
		var rsp ctp.CThostFtdcInstrumentField
		var rspInfo ctp.CThostFtdcRspInfoField
		for i := range ret {
			if err := ret[i].Get2(&rsp, &rspInfo); err == nil {
				log.Println(rsp.InstrumentID, util.StringFromGBK2(rsp.InstrumentName), ret[i].IsLast)
			}
		}
	}

	sig := make(chan bool)
	<-sig
}

func Test_async(t *testing.T) {
	return
	var c config
	err := loadConfig("config.json", &c)
	if err != nil {
		panic("需要自己在config.json里配置账号密码，前置地址等")
	}
	log.Println("c", c)
	password := c.Password
	userid := c.Account
	brokerid := c.BrokerID
	appid := c.AppID
	authcode := c.AuthCode
	sig := make(chan interface{})
	var requestID int32
	var adapter *Adapter
	adapter, err = NewAdapter("47.100.1.102:8205", timeout, func(pkt *Packet) {
		switch ctp.CtpMessageType(pkt.MsgType) {
		case ctp.CtpMessageType_TD_OnFrontConnected:
			var req ctp.CThostFtdcReqAuthenticateField
			req.BrokerID = brokerid
			req.AppID = appid
			req.AuthCode = authcode
			req.UserID = userid
			requestID++
			adapter.Post(int32(ctp.CtpMessageType_TD_ReqAuthenticate), &req, requestID)
		case ctp.CtpMessageType_TD_OnRspAuthenticate:
			var rsp ctp.CThostFtdcRspAuthenticateField
			var rspInfo ctp.CThostFtdcRspInfoField
			if err := pkt.Get2(&rsp, &rspInfo); err == nil {
				log.Println(rsp, rspInfo, util.StringFromGBK2(rspInfo.ErrorMsg))
				var req ctp.CThostFtdcReqUserLoginField
				req.BrokerID = brokerid
				req.UserID = userid
				req.Password = password
				requestID++
				adapter.Post(int32(ctp.CtpMessageType_TD_ReqUserLogin), &req, requestID)
			}
		case ctp.CtpMessageType_TD_OnRspUserLogin:
			var rsp ctp.CThostFtdcRspUserLoginField
			var rspInfo ctp.CThostFtdcRspInfoField
			if err := pkt.Get2(&rsp, &rspInfo); err == nil {
				log.Println(rsp, util.StringFromGBK2(rspInfo.ErrorMsg))
			}
			sig <- true
		case ctp.CtpMessageType_TD_OnRtnOrder:
			var rtn ctp.CThostFtdcOrderField
			if err := pkt.Get1(&rtn); err == nil {
				log.Println(rtn, util.StringFromGBK2(rtn.StatusMsg))
			}
		case ctp.CtpMessageType_TD_OnRtnTrade:
			var rtn ctp.CThostFtdcTradeField
			if err := pkt.Get1(&rtn); err == nil {
				log.Println(rtn)
			}
		case ctp.CtpMessageType_TD_OnRtnInstrumentStatus:
			var rtn ctp.CThostFtdcInstrumentStatusField
			if err := pkt.Get1(&rtn); err == nil {
				// log.Println(rtn)
			}
		default:
			log.Println(ctp.CtpMessageType(pkt.MsgType), len(pkt.BodyList))
		}
	})
	// trade
	var req ctp.CThostFtdcReqRegisterFrontField
	req.Fronts = c.Fronts
	requestID++
	adapter.Post(int32(ctp.CtpMessageType_TD_RegisterFront), &req, requestID)
	requestID++
	adapter.Post(int32(ctp.CtpMessageType_TD_Init), nil, requestID)
	<-sig
}

func Test_md(t *testing.T) {
	var fronts []string
	fronts = append(fronts, "tcp://182.131.17.103:41168")
	s, err := NewSubscriber("47.100.1.102:8213", "test", "", fronts, time.Second*3, func(rtn *ctp.CThostFtdcDepthMarketDataField) {
		log.Println(rtn)
	})
	if err != nil {
		panic(err)
	}
	s.Subscribe("SHFE", "ru1909")
	s.Subscribe("CFFEX", "IF1908")
	<-time.After(time.Second * 60)
}
