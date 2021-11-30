package service

import (
	"errors"
	"live/internal/dao"
	"live/internal/entity"
)

type RoomServiceImpl struct {
	roomDao dao.RoomDao
	txDao   dao.TxDao
}

func NewRoomServiceImpl() *RoomServiceImpl {
	return &RoomServiceImpl{roomDao: dao.NewRoomDaoImpl(), txDao: dao.NewTxDaoImpl()}
}

func (u *RoomServiceImpl) AddRoom(room *entity.Room) uint32 {
	return u.roomDao.AddRoom(room)
}

func (u *RoomServiceImpl) DeleteRoom(room *entity.Room) uint32 {
	return u.roomDao.DeleteRoom(room)
}

func (u *RoomServiceImpl) UpdateRoom(room *entity.Room) uint32 {
	return u.roomDao.UpdateRoom(room)
}

func (u *RoomServiceImpl) QueryRoomByID(ID uint32) *entity.Room {
	return u.roomDao.QueryRoomByID(ID)
}

func (u *RoomServiceImpl) QueryTurnoverByID(ID uint32) (turnover uint32, err error) {
	room := u.roomDao.QueryRoomByID(ID)
	if room != nil {
		txs := u.txDao.QueryTxsByRoomID(ID)
		for _, tx := range *txs {
			turnover += tx.Turnover
		}
	} else {
		err = errors.New("直播间ID不存在")
	}
	return
}
