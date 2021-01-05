package dao

import (
	"github.com/muchlist/gorm-imp/database"
	dto2 "github.com/muchlist/gorm-imp/dto"
)

var (
	PegawaiDao PegawaiDaoInterface = &pegawaiDao{}
)

type pegawaiDao struct{}

type PegawaiDaoInterface interface {
	Find() ([]dto2.Pegawai, error)
	Create(data dto2.Pegawai) (*dto2.Pegawai, error)
}

func (p *pegawaiDao) Create(data dto2.Pegawai) (*dto2.Pegawai, error) {
	db := database.DbConn
	var pegawai = data
	err := db.Create(&pegawai).Error
	if err != nil {
		return nil, err
	}

	return &pegawai, nil
}

func (p *pegawaiDao) Find() ([]dto2.Pegawai, error) {
	db := database.DbConn
	var pegawais []dto2.Pegawai
	err := db.Find(&pegawais).Error
	if err != nil {
		return nil, err
	}

	return pegawais, nil
}
