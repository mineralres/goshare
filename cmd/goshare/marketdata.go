package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mineralres/goshare/pkg/api"
	pb "github.com/mineralres/goshare/pkg/pb/goshare"
	// "google.golang.org/grpc"
)

func (g *Gateway) handleStream(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to set websocket upgrade:", err)
		return
	}
	defer conn.Close()
	type message struct {
		Type string      `json:"type"`
		Data interface{} `json:"data"`
	}
	chOut := make(chan *message, 100)
	chSub := make(chan *pb.MarketDataSnapshot, 9999)
	index := atomic.AddInt64(&g.wsIndex, 1)
	go func() {
		// 写函数
		for {
			select {
			case msg := <-chOut:
				d, err := json.Marshal(msg)
				if err != nil {
					continue
				}
				err = conn.WriteMessage(websocket.TextMessage, d)
				if err != nil {
					return
				}
			case md := <-chSub:
				msg := new(message)
				msg.Type = pb.MessageType_RTN_MARKET_DATA_UPDATE.String()
				msg.Data = md
				chOut <- msg
			}
		}
	}()
	// 先接收订阅
	for {
		// 读取
		t, p, err := conn.ReadMessage()
		if err != nil {
			return
		}
		if t != websocket.TextMessage {
			continue
		}
		msg := new(message)
		err = json.Unmarshal(p, msg)
		if err != nil {
			continue
		}
		typ, ok := pb.MessageType_value[msg.Type]
		if !ok {
			continue
		}
		switch pb.MessageType(typ) {
		case pb.MessageType_REQ_UNSUBSCRIBE_MARKET_DATA:
			var req pb.ReqUnSubscribe
			msg.Data = &req
			if err = json.Unmarshal(p, msg); err != nil {
				log.Println("req", req, err)
				continue
			}
		case pb.MessageType_REQ_SUBSCRIBE_MARKET_DATA:
			// 订阅行情
			log.Printf("front[%d] received msg[%v] len[%d]", index, msg.Type, len(p))
			var req pb.ReqSubscribe
			msg.Data = &req
			if err = json.Unmarshal(p, msg); err != nil {
				log.Println("req", req, err)
				continue
			}
		case pb.MessageType_HEATBEAT:
			// 心跳
			log.Printf("front[%d] received msg[%v]", index, msg.Type)
			chOut <- msg
		default:
			log.Printf("front[%d] received msg[%v]", index, msg.Type)
		}
	}
}

var (
	errUnsported = errors.New("unsported")
)

func (g *Gateway) instrumentList(r *http.Request) (interface{}, error) {
	var req pb.ReqGetTradingInstrumentList
	var err error
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	var ctx api.Context
	for i := range g.dsList {
		if resp, err := g.dsList[i].TradingInstrumentList(&ctx, &req); err == nil {
			return resp, err
		}
	}
	return nil, errUnsported
}

func (g *Gateway) mainContract(r *http.Request) (interface{}, error) {
	// day := getDay()
	// key := []byte(fmt.Sprintf("-main-contract-%d", day))
	// var l []*pb.TradingInstrument
	// d, err := g.backend.Get(key)
	// if err != nil {
	// 	l, _ = g.getMainContract()
	// 	out, _ := json.Marshal(&l)
	// 	g.backend.Set(key, out)
	// } else {
	// 	json.Unmarshal(d, &l)
	// }
	// var ctx api.Context
	// var ret []*pb.MarketDataSnapshot
	// for _, ti := range l {
	// 	for _, ds := range g.dsList {
	// 		if md, err := ds.GetLastTick(&ctx, ti.Symbol); err == nil && md != nil && md.Symbol != nil {
	// 			ret = append(ret, md)
	// 		}
	// 	}
	// }
	// return ret, nil
	return nil, nil
}

func (g *Gateway) getMainContract() ([]*pb.Instrument, error) {
	// var ctx api.Context
	// var resp []*pb.TradingInstrument
	// var err error
	// for _, ds := range c.dsList {
	// 	if resp, err = ds.TradingInstrumentList(&ctx, &pb.ReqGetTradingInstrumentList{}); err == nil {
	// 		break
	// 	}
	// }
	// m := make(map[string]*pb.TradingInstrument)
	// for _, ti := range resp {
	// 	v, ok := m[ti.ProductInfo.ProductId.Code]
	// 	if ok {
	// 		if ti.InstrumentInfo.PrePosition >= v.InstrumentInfo.PrePosition {
	// 			m[ti.ProductInfo.ProductId.Code] = ti
	// 		}
	// 	} else {
	// 		m[ti.ProductInfo.ProductId.Code] = ti
	// 	}
	// }
	// var ret []*pb.TradingInstrument
	// for _, ti := range m {
	// 	ret = append(ret, ti)
	// }
	// return ret, nil
	return nil, nil
}

func getDay() int32 {
	str := time.Now().Format("20060102")
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return int32(i)
}
