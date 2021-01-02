package pegawai

import (
	"github.com/muchlist/gorm-imp/database"
)

var (
	PegawaiDao PegawaiDaoInterface = &pasienDao{}
)

type pasienDao struct{}

type PegawaiDaoInterface interface {
	Find() []Pegawai
	Create(data Pegawai) (Pegawai, error)
}

func (p *pasienDao) Create(data Pegawai) (Pegawai, error) {
	db := database.DbConn
	var pegawai = data
	result := db.Create(&pegawai)

	return pegawai, result.Error
}

func (p *pasienDao) Find() []Pegawai {
	db := database.DbConn
	var pegawais []Pegawai
	db.Find(&pegawais)

	return pegawais
}
