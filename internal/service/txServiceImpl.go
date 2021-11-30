package service

import (
	"errors"
	"live/internal/dao"
	"live/internal/entity"
)

type TxServiceImpl struct {
	txDao dao.TxDao
}

func NewTxServiceImpl() *TxServiceImpl {
	return &TxServiceImpl{txDao: dao.NewTxDaoImpl()}
}

func (u *TxServiceImpl) AddTx(tx *entity.Tx) (ID uint32, err error) {
	if u.txDao.AddTx(tx) != 0 {
		ID = tx.TxID
	} else {
		err = errors.New("新增交易失败")
	}
	return
}

func (u *TxServiceImpl) DeleteTx(tx *entity.Tx) error {
	if u.txDao.DeleteTx(tx) != 0 {
		return nil
	}
	return errors.New("交易ID不存在")
}

func (u *TxServiceImpl) UpdateTx(tx *entity.Tx) uint32 {
	return u.txDao.UpdateTx(tx)
}

func (u *TxServiceImpl) QueryTxByID(ID uint32) *entity.Tx {
	return u.txDao.QueryTxByID(ID)
}
