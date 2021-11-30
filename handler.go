package main

import (
	"context"
	"live/internal/entity"
	"live/internal/service"
	"live/kitex_gen/api"
	"strconv"
)

// LiveImpl implements the last service interface defined in the IDL.
type LiveImpl struct {
	txService   service.TxService
	roomService service.RoomService
}

func NewLiveImpl() *LiveImpl {
	return &LiveImpl{txService: service.NewTxServiceImpl(), roomService: service.NewRoomServiceImpl()}
}

// AddTxByID implements the LiveImpl interface.
func (s *LiveImpl) AddTxByID(ctx context.Context, req *api.AddReq) (resp *api.BaseResp, err error) {
	tx := entity.Tx{TxID: req.TxID, RoomID: req.RoomID, GoodID: req.GoodID, Turnover: req.TxTurnover, TxTime: req.TxTime.AsTime(), UserID: req.UserID}
	ID, err := s.txService.AddTx(&tx)
	if err == nil {
		resp = &api.BaseResp{Msg: "新增交易记录成功，交易ID：" + strconv.Itoa(int(ID))}
	}
	return
}

// DeleteTxByID implements the LiveImpl interface.
func (s *LiveImpl) DeleteTxByID(ctx context.Context, req *api.DeleteReq) (resp *api.BaseResp, err error) {
	err = s.txService.DeleteTx(&entity.Tx{TxID: req.TxID})
	if err == nil {
		resp = &api.BaseResp{Msg: "删除交易成功"}
	}
	return
}

// QueryTurnoverByRoomID implements the LiveImpl interface.
func (s *LiveImpl) QueryTurnoverByRoomID(ctx context.Context, req *api.QueryTurnoverByRoomIDReq) (resp *api.QueryTurnoverByRoomIDResp, err error) {
	turnover, err := s.roomService.QueryTurnoverByID(req.RoomID)
	if err == nil {
		resp = &api.QueryTurnoverByRoomIDResp{Turnover: turnover}
	}
	return
}
