package terapi

import (
	"github.com/muchlist/gorm-imp/database"
	"github.com/muchlist/gorm-imp/domains/pasien"
	"time"
)

var (
	TerapiDao terapiDaoInterface = &terapiDao{}
)

type terapiDao struct{}

type terapiDaoInterface interface {
	Create(data Terapi) (Terapi, error)
	Find() []Terapi
	FindByDateRange(start, end time.Time) []Terapi
}

func (p *terapiDao) Create(data Terapi) (Terapi, error) {
	db := database.DbConn

	tx := db.Begin()
	var terapiData = data
	err := tx.Create(&terapiData).Error
	if err != nil {
		tx.Rollback()
		return Terapi{}, err
	}

	var pasienToUpdate pasien.Pasien
	err = tx.First(&pasienToUpdate, terapiData.PasienID).Error
	if err != nil {
		tx.Rollback()
		return Terapi{}, err
	}

	pasienToUpdate.JumlahTerapi = pasienToUpdate.JumlahTerapi + 1

	err = tx.Save(&pasienToUpdate).Error
	if err != nil {
		tx.Rollback()
		return Terapi{}, err
	}

	tx.Commit()

	return terapiData, nil
}

func (p *terapiDao) Find() []Terapi {
	db := database.DbConn
	var terapis []Terapi
	db.Find(&terapis)

	return terapis
}

func (p *terapiDao) FindByDateRange(start, end time.Time) []Terapi {
	db := database.DbConn
	var terapis []Terapi
	db.Where("tglterapi >= ? AND tglterapi <= ?", start, end).Find(&terapis)
	// SELECT * FROM terapis WHERE tglterapi >= '2000-01-01 00:00:00' AND tglterapi <= '2000-01-01 00:00:00';

	return terapis
}
