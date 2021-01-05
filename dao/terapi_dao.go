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
	Create(data dto2.Terapi) (*dto2.Terapi, error)
	Find() ([]dto2.Terapi, error)
	FindByDateRange(start, end time.Time) ([]dto2.Terapi, error)
}

func (p *terapiDao) Create(data dto2.Terapi) (*dto2.Terapi, error) {
	db := database.DbConn

	tx := db.Begin()
	var terapiData = data
	err := tx.Create(&terapiData).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	var pasienToUpdate dto2.Pasien
	err = tx.First(&pasienToUpdate, terapiData.PasienID).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	pasienToUpdate.JumlahTerapi = pasienToUpdate.JumlahTerapi + 1

	err = tx.Save(&pasienToUpdate).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &terapiData, nil
}

func (p *terapiDao) Find() ([]dto2.Terapi, error) {
	db := database.DbConn
	var terapis []dto2.Terapi
	err := db.Find(&terapis).Error
	if err != nil {
		return nil, err
	}

	return terapis, nil
}

func (p *terapiDao) FindByDateRange(start, end time.Time) ([]dto2.Terapi, error) {
	db := database.DbConn
	var terapis []dto2.Terapi
	err := db.Where("tglterapi >= ? AND tglterapi <= ?", start, end).Find(&terapis).Error
	// SELECT * FROM terapis WHERE tglterapi >= '2000-01-01 00:00:00' AND tglterapi <= '2000-01-01 00:00:00';
	if err != nil {
		return nil, err
	}

	return terapis, nil
}
