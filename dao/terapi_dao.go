package dao

import (
	"github.com/muchlist/gorm-imp/database"
	dto2 "github.com/muchlist/gorm-imp/dto"
	"time"
)

var (
	TerapiDao terapiDaoInterface = &terapiDao{}
)

type terapiDao struct{}

type terapiDaoInterface interface {
	Create(data dto2.Terapi) (dto2.Terapi, error)
	Find() []dto2.Terapi
	FindByDateRange(start, end time.Time) []dto2.Terapi
}

func (p *terapiDao) Create(data dto2.Terapi) (dto2.Terapi, error) {
	db := database.DbConn

	tx := db.Begin()
	var terapiData = data
	err := tx.Create(&terapiData).Error
	if err != nil {
		tx.Rollback()
		return dto2.Terapi{}, err
	}

	var pasienToUpdate dto2.Pasien
	err = tx.First(&pasienToUpdate, terapiData.PasienID).Error
	if err != nil {
		tx.Rollback()
		return dto2.Terapi{}, err
	}

	pasienToUpdate.JumlahTerapi = pasienToUpdate.JumlahTerapi + 1

	err = tx.Save(&pasienToUpdate).Error
	if err != nil {
		tx.Rollback()
		return dto2.Terapi{}, err
	}

	tx.Commit()

	return terapiData, nil
}

func (p *terapiDao) Find() []dto2.Terapi {
	db := database.DbConn
	var terapis []dto2.Terapi
	db.Find(&terapis)

	return terapis
}

func (p *terapiDao) FindByDateRange(start, end time.Time) []dto2.Terapi {
	db := database.DbConn
	var terapis []dto2.Terapi
	db.Where("tglterapi >= ? AND tglterapi <= ?", start, end).Find(&terapis)
	// SELECT * FROM terapis WHERE tglterapi >= '2000-01-01 00:00:00' AND tglterapi <= '2000-01-01 00:00:00';

	return terapis
}
