package service

import (
	"live/internal/dao"
	"live/internal/entity"
)

type GoodServiceImpl struct {
	goodDao dao.GoodDao
}

func NewGoodServiceImpl() *GoodServiceImpl {
	return &GoodServiceImpl{goodDao: dao.NewGoodDaoImpl()}
}

func (u *GoodServiceImpl) AddGood(good *entity.Good) uint32 {
	return u.goodDao.AddGood(good)
}

func (u *GoodServiceImpl) DeleteGood(good *entity.Good) uint32 {
	return u.goodDao.DeleteGood(good)
}

func (u *GoodServiceImpl) UpdateGood(good *entity.Good) uint32 {
	return u.goodDao.UpdateGood(good)
}

func (u *GoodServiceImpl) QueryGoodByID(ID uint32) *entity.Good {
	return u.goodDao.QueryGoodByID(ID)
}
