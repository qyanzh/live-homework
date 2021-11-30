package dao

import (
	"gorm.io/gorm"
	"live/internal/entity"
	"strconv"
)

type TxDaoImpl struct {
	db *gorm.DB
}

func NewTxDaoImpl() TxDao {
	return &TxDaoImpl{db: DbClient}
}

func (u *TxDaoImpl) AddTx(tx *entity.Tx) uint32 {
	result := u.db.Create(&tx)
	return uint32(result.RowsAffected)
}

func (u *TxDaoImpl) DeleteTx(tx *entity.Tx) uint32 {
	result := u.db.Where("tx_id = ?", strconv.Itoa(int(tx.TxID))).Delete(&tx)
	return uint32(result.RowsAffected)
}

func (u *TxDaoImpl) UpdateTx(tx *entity.Tx) uint32 {
	result := u.db.Model(&tx).Updates(tx)
	return uint32(result.RowsAffected)
}

func (u *TxDaoImpl) QueryTxByID(ID uint32) *entity.Tx {
	tx := entity.Tx{TxID: ID}
	result := u.db.First(&tx)
	if result.Error != nil {
		return nil
	}
	return &tx
}

func (u *TxDaoImpl) QueryTxsByRoomID(roomID uint32) *[]entity.Tx {
	var txs []entity.Tx
	u.db.Where("room_id = ?", strconv.Itoa(int(roomID))).Find(&txs)
	return &txs
}
