package dao

import (
	"github.com/muchlist/gorm-imp/database"
	"github.com/muchlist/gorm-imp/dto"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

var (
	PasienDao pasienDaoInterface = &pasienDao{}
)

type pasienDao struct{}

type pasienDaoInterface interface {
	Find(gender string) ([]dto.Pasien, error)
	Create(data dto.Pasien) (*dto.Pasien, error)
	GetPasienLastIDWithGender(gender int) (int, error)
}

func (p *pasienDao) Create(data dto.Pasien) (*dto.Pasien, error) {
	db := database.DbConn
	var pasien = data
	err := db.Create(&pasien).Error
	if err != nil {
		return nil, err
	}

	return &pasien, nil
}

func (p *pasienDao) GetPasienLastIDWithGender(gender int) (int, error) {
	db := database.DbConn
	var pasien dto.Pasien
	result := db.Where("jk = ?", gender).Last(&pasien)
	if result.Error == gorm.ErrRecordNotFound {
		return 0, nil
	}
	if result.Error != nil {
		return 0, result.Error
	}

	pasienNumber, err := strconv.Atoi(pasien.NoPasien)
	if err != nil {
		return 0, err
	}

	return pasienNumber, nil
}

func (p *pasienDao) Find(gender string) ([]dto.Pasien, error) {
	db := database.DbConn
	var pasiens []dto.Pasien

	if gender != "" {
		genderNum := 0
		if strings.ToLower(gender) == "p" {
			genderNum = 1
		}
		err := db.Where("jk = ?", genderNum).Preload("Terapis").Order("jumlah_terapi asc").Find(&pasiens).Error
		if err != nil {
			return nil, err
		}
	} else {
		err := db.Preload("Terapis").Order("jumlah_terapi asc").Find(&pasiens).Error
		if err != nil {
			return nil, err
		}
	}

	return pasiens, nil
}
