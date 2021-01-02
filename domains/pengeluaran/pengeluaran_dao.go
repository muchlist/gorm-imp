package pengeluaran

import (
	"github.com/muchlist/gorm-imp/database"
	"github.com/muchlist/gorm-imp/domains/dto"
)

var (
	PengeluaranDao PengeluaranDaoInterface = &pengeluaranDao{}
)

type pengeluaranDao struct{}

type PengeluaranDaoInterface interface {
	Find() []dto.Pengeluaran
	Create(data dto.Pengeluaran) (dto.Pengeluaran, error)
}

func (p *pengeluaranDao) Create(data dto.Pengeluaran) (dto.Pengeluaran, error) {
	db := database.DbConn
	var pengeluaran = data
	err := db.Create(&pengeluaran).Error
	if err != nil {
		return dto.Pengeluaran{}, err
	}

	return pengeluaran, nil
}

func (p *pengeluaranDao) Find() []dto.Pengeluaran {
	db := database.DbConn
	var pengeluarans []dto.Pengeluaran
	db.Find(&pengeluarans)

	return pengeluarans
}
