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
	Find() []dto2.Pegawai
	Create(data dto2.Pegawai) (dto2.Pegawai, error)
}

func (p *pegawaiDao) Create(data dto2.Pegawai) (dto2.Pegawai, error) {
	db := database.DbConn
	var pegawai = data
	result := db.Create(&pegawai)

	return pegawai, result.Error
}

func (p *pegawaiDao) Find() []dto2.Pegawai {
	db := database.DbConn
	var pegawais []dto2.Pegawai
	db.Find(&pegawais)

	return pegawais
}
