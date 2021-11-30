package service

import (
	"live/internal/entity"
)

type TxService interface {
	AddTx(tx *entity.Tx) (uint32, error)
	DeleteTx(tx *entity.Tx) error
	UpdateTx(tx *entity.Tx) uint32
	QueryTxByID(ID uint32) *entity.Tx
}
