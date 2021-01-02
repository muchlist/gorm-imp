package pegawai

import (
	"github.com/muchlist/gorm-imp/database"
	"github.com/muchlist/gorm-imp/domains/dto"
)

var (
	PegawaiDao PegawaiDaoInterface = &pegawaiDao{}
)

type pegawaiDao struct{}

type PegawaiDaoInterface interface {
	Find() []dto.Pegawai
	Create(data dto.Pegawai) (dto.Pegawai, error)
}

func (p *pegawaiDao) Create(data dto.Pegawai) (dto.Pegawai, error) {
	db := database.DbConn
	var pegawai = data
	result := db.Create(&pegawai)

	return pegawai, result.Error
}

func (p *pegawaiDao) Find() []dto.Pegawai {
	db := database.DbConn
	var pegawais []dto.Pegawai
	db.Find(&pegawais)

	return pegawais
}
