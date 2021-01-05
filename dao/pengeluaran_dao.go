package dao

import (
	"github.com/muchlist/gorm-imp/database"
	dto2 "github.com/muchlist/gorm-imp/dto"
)

var (
	PengeluaranDao PengeluaranDaoInterface = &pengeluaranDao{}
)

type pengeluaranDao struct{}

type PengeluaranDaoInterface interface {
	Find() []dto2.Pengeluaran
	Create(data dto2.Pengeluaran) (dto2.Pengeluaran, error)
}

func (p *pengeluaranDao) Create(data dto2.Pengeluaran) (dto2.Pengeluaran, error) {
	db := database.DbConn
	var pengeluaran = data
	err := db.Create(&pengeluaran).Error
	if err != nil {
		return dto2.Pengeluaran{}, err
	}

	return pengeluaran, nil
}

func (p *pengeluaranDao) Find() []dto2.Pengeluaran {
	db := database.DbConn
	var pengeluarans []dto2.Pengeluaran
	db.Find(&pengeluarans)

	return pengeluarans
}
