package ctp

import (
	"errors"
	"log"
	"testing"

	proto "github.com/golang/protobuf/proto"
	"github.com/mineralres/goshare/pkg/pb/ctp"
	"github.com/mineralres/goshare/pkg/util"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func makeData(p proto.Message) []byte {
	d, err := proto.Marshal(p)
	if err != nil {
		return nil
	}
	return d
}

func parse1(pkt *packet, p1 proto.Message) error {
	if len(pkt.BodyList) < 1 {
		log.Println("len(pkt.BodyList) < 1")
		return errors.New("len(pkt.BodyList) < 1")
	}
	if err := proto.Unmarshal(pkt.BodyList[0], p1); err != nil {
		return err
	}
	return nil
}

func parse2(pkt *packet, p1 proto.Message, p2 proto.Message) error {
	if len(pkt.BodyList) < 2 {
		log.Println("len(pkt.BodyList) < 2")
		return errors.New("len(pkt.BodyList) < 2")
	}
	if err := proto.Unmarshal(pkt.BodyList[0], p1); err != nil {
		return err
	}
	if err := proto.Unmarshal(pkt.BodyList[1], p2); err != nil {
		return err
	}
	return nil
}

func Test_i(t *testing.T) {
	const password = "dyc123"
	const userid = "146257"
	const brokerid = "9999"
	const appid = "8510788746"
	const authcode = "0000000000000000"
	sig := make(chan interface{})
	var requestID int32
	var adapter *Adapter
	adapter = NewAdapter("47.100.1.102:8205", func(pkt *packet) {
		switch ctp.CtpMessageType(pkt.MsgType) {
		case ctp.CtpMessageType_TD_OnFrontConnected:
			var req ctp.CThostFtdcReqAuthenticateField
			req.BrokerID = brokerid
			req.AppID = appid
			req.AuthCode = authcode
			req.UserID = userid
			requestID++
			adapter.Send(int32(ctp.CtpMessageType_TD_ReqAuthenticate), makeData(&req), requestID)
		case ctp.CtpMessageType_TD_OnRspAuthenticate:
			var rsp ctp.CThostFtdcRspAuthenticateField
			var rspInfo ctp.CThostFtdcRspInfoField
			if err := parse2(pkt, &rsp, &rspInfo); err == nil {
				log.Println(rsp, rspInfo, util.StringFromGBK2(rspInfo.ErrorMsg))
				var req ctp.CThostFtdcReqUserLoginField
				req.BrokerID = brokerid
				req.UserID = userid
				req.Password = password
				requestID++
				adapter.Send(int32(ctp.CtpMessageType_TD_ReqUserLogin), makeData(&req), requestID)
			}
		case ctp.CtpMessageType_TD_OnRspUserLogin:
			var rsp ctp.CThostFtdcRspUserLoginField
			var rspInfo ctp.CThostFtdcRspInfoField
			if err := parse2(pkt, &rsp, &rspInfo); err == nil {
				log.Println(rsp, rspInfo, util.StringFromGBK2(rspInfo.ErrorMsg))
			}
		case ctp.CtpMessageType_TD_OnRtnOrder:
			var rtn ctp.CThostFtdcOrderField
			if err := parse1(pkt, &rtn); err == nil {
				log.Println(rtn, util.StringFromGBK2(rtn.StatusMsg))
			}
		case ctp.CtpMessageType_TD_OnRtnTrade:
			var rtn ctp.CThostFtdcTradeField
			if err := parse1(pkt, &rtn); err == nil {
				log.Println(rtn)
			}
		default:
			log.Println(ctp.CtpMessageType(pkt.MsgType), len(pkt.BodyList))
		}
	})
	// trade
	var req ctp.CThostFtdcReqRegisterFrontField
	req.Front = "tcp://180.168.146.187:10001"
	requestID++
	adapter.Send(int32(ctp.CtpMessageType_TD_RegisterFront), makeData(&req), requestID)
	requestID++
	adapter.Send(int32(ctp.CtpMessageType_TD_Init), nil, requestID)
	<-sig
}

func Test_md(t *testing.T) {
	sig := make(chan interface{})
	var requestID int32
	var adapter *Adapter
	adapter = NewAdapter("47.100.1.102:8213", func(pkt *packet) {
		switch ctp.CtpMessageType(pkt.MsgType) {
		case ctp.CtpMessageType_MD_OnRspSubMarketData:
		case ctp.CtpMessageType_MD_OnRtnDepthMarketData:
			var rtn ctp.CThostFtdcDepthMarketDataField
			if err := parse1(pkt, &rtn); err == nil {
				log.Println(rtn)
			}
		case ctp.CtpMessageType_MD_OnFrontConnected:
			var req ctp.CThostFtdcReqUserLoginField
			req.UserID = "test"
			requestID++
			adapter.Send(int32(ctp.CtpMessageType_MD_ReqUserLogin), makeData(&req), requestID)
		case ctp.CtpMessageType_MD_OnRspUserLogin:
			var rsp ctp.CThostFtdcRspUserLoginField
			var rspInfo ctp.CThostFtdcRspInfoField
			if err := parse2(pkt, &rsp, &rspInfo); err == nil {
				// log.Println(rspInfo.ErrorID, util.StringFromGBK2(rspInfo.ErrorMsg))
			}
			var req ctp.CThostFtdcReqSubscribeMarketData
			req.Instruments = append(req.Instruments, "ru1909")
			req.Instruments = append(req.Instruments, "IF1907")
			req.Instruments = append(req.Instruments, "SR909")
			requestID++
			adapter.Send(int32(ctp.CtpMessageType_MD_SubscribeMarketData), makeData(&req), requestID)
		default:
			log.Println(ctp.CtpMessageType(pkt.MsgType), len(pkt.BodyList))
		}
	})
	// md
	var req ctp.CThostFtdcReqRegisterFrontField
	req.Front = "tcp://182.131.17.103:41168"
	requestID++
	adapter.Send(int32(ctp.CtpMessageType_MD_RegisterFront), makeData(&req), requestID)
	requestID++
	adapter.Send(int32(ctp.CtpMessageType_MD_Init), nil, requestID)
	<-sig
}
