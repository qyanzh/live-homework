package dao

import (
	"live/internal/entity"
)

type TxDao interface {
	AddTx(tx *entity.Tx) uint32
	DeleteTx(tx *entity.Tx) uint32
	UpdateTx(tx *entity.Tx) uint32
	QueryTxByID(ID uint32) *entity.Tx
	QueryTxsByRoomID(roomID uint32) *[]entity.Tx
}
