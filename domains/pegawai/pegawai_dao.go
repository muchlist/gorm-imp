package pegawai

import (
	"github.com/muchlist/gorm-imp/database"
	"github.com/muchlist/gorm-imp/domains/dto"
)

var (
	PegawaiDao PegawaiDaoInterface = &pasienDao{}
)

type pasienDao struct{}

type PegawaiDaoInterface interface {
	Find() []dto.Pegawai
	Create(data dto.Pegawai) (dto.Pegawai, error)
}

func (p *pasienDao) Create(data dto.Pegawai) (dto.Pegawai, error) {
	db := database.DbConn
	var pegawai = data
	result := db.Create(&pegawai)

	return pegawai, result.Error
}

func (p *pasienDao) Find() []dto.Pegawai {
	db := database.DbConn
	var pegawais []dto.Pegawai
	db.Find(&pegawais)

	return pegawais
}
