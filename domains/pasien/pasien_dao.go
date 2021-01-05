package pasien

import (
	"github.com/muchlist/gorm-imp/database"
	dto2 "github.com/muchlist/gorm-imp/dto"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

var (
	PasienDao pasienDaoInterface = &pasienDao{}
)

type pasienDao struct{}

type pasienDaoInterface interface {
	Find(gender string) []dto2.Pasien
	Create(data dto2.Pasien) (dto2.Pasien, error)
	GetPasienLastIDWithGender(gender int) (int, error)
}

func (p *pasienDao) Create(data dto2.Pasien) (dto2.Pasien, error) {
	db := database.DbConn
	var pasien = data
	result := db.Create(&pasien)

	// pasien.ID             // returns inserted data's primary key
	// result.Error        // returns error
	// result.RowsAffected // returns inserted records count

	return pasien, result.Error
}

func (p *pasienDao) GetPasienLastIDWithGender(gender int) (int, error) {
	db := database.DbConn
	var pasien dto2.Pasien
	result := db.Where("jk = ?", gender).Last(&pasien)
	if result.Error == gorm.ErrRecordNotFound {
		return 0, nil
	}

	pasienNumber, err := strconv.Atoi(pasien.NoPasien)
	if err != nil {
		return 0, err
	}

	return pasienNumber, nil
}

func (p *pasienDao) Find(gender string) []dto2.Pasien {
	db := database.DbConn
	var pasiens []dto2.Pasien

	if gender != "" {
		genderNum := 0
		if strings.ToLower(gender) == "p" {
			genderNum = 1
		}
		db.Where("jk = ?", genderNum).Preload("Terapis").Order("jumlah_terapi asc").Find(&pasiens)
	} else {
		db.Preload("Terapis").Order("jumlah_terapi asc").Find(&pasiens)
	}

	return pasiens
}
