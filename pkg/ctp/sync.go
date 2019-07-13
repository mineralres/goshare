package ctp

import (
	"errors"
	"time"

	proto "github.com/golang/protobuf/proto"
	"github.com/mineralres/goshare/pkg/pb/ctp"
)

// SyncAdapter  sync
type SyncAdapter struct {
	adapter *Adapter
	chIn    chan *Packet
}

// NewSyncAdapter create new sync adapter
func NewSyncAdapter(host string, timeout time.Duration, fronts []string) (*SyncAdapter, error) {
	var err error
	ret := &SyncAdapter{}
	ret.adapter, err = NewAdapter(host, timeout, func(pkt *Packet) {
		if pkt.MsgType == int32(ctp.CtpMessageType_HEARTBEAT) {
			return
		}
		ret.chIn <- pkt
	})
	if err != nil {
		return nil, err
	}
	ret.chIn = make(chan *Packet, 1000)
	var req ctp.CThostFtdcReqRegisterFrontField
	req.Fronts = fronts
	ret.adapter.Post(int32(ctp.CtpMessageType_TD_RegisterFront), &req, 0)
	ret.adapter.Post(int32(ctp.CtpMessageType_TD_Init), &req, 0)
	for {
		select {
		case <-time.After(timeout):
			return ret, errors.New("timeout")
		case pkt := <-ret.chIn:
			if pkt.MsgType == int32(ctp.CtpMessageType_TD_OnFrontConnected) {
				return ret, nil
			}
		}
	}
	return ret, nil
}

// Send send msg
func (sa *SyncAdapter) Send(msgType int32, req proto.Message, requestID int32, timeout time.Duration) ([]*Packet, error) {
	sa.adapter.Post(msgType, req, requestID)
	var ret []*Packet
	for {
		select {
		case <-time.After(timeout):
			return ret, errors.New("timeout")
		case pkt := <-sa.chIn:
			if pkt.RequestID == requestID && pkt.IsLast > 0 {
				ret = append(ret, pkt)
				return ret, nil
			}
		}
	}
}
