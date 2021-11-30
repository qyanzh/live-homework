package main

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"live/kitex_gen/api"
	"live/kitex_gen/api/live"
	"live/kitex_gen/google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"testing"
	"time"
)

var c live.Client

func init() {
	var err error
	c, err = live.NewClient("live", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
}

func handleResp(resp interface{}, err error) {
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}

func TestAdd(t *testing.T) {
	req := &api.AddReq{TxID: uint32(time.Now().UnixMicro()), RoomID: 100, GoodID: 100, TxTurnover: 100, TxTime: timestamppb.Now(), UserID: 100}
	resp, err := c.AddTxByID(context.Background(), req /*, callopt.WithRPCTimeout(3600*time.Second)*/)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}

func TestDelete(t *testing.T) {
	req := &api.DeleteReq{TxID: 2074760185}
	resp, err := c.DeleteTxByID(context.Background(), req)
	handleResp(resp, err)
}

func TestQueryTurnoverByRoomID(t *testing.T) {
	req := &api.QueryTurnoverByRoomIDReq{RoomID: 100}
	resp, err := c.QueryTurnoverByRoomID(context.Background(), req)
	handleResp(resp, err)
}
