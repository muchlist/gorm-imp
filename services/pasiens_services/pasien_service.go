package pasiens_services

import (
	"github.com/muchlist/erru_utils_go/rest_err"
	"github.com/muchlist/gorm-imp/domains/pasien"
	"strconv"
)

var (
	PasienService pasienServiceInterface = &pasienService{}
)

type pasienService struct{}

type pasienServiceInterface interface {
	Find() []pasien.Pasien
	Create(data pasien.Pasien) (*pasien.Pasien, rest_err.APIError)
}

func (p *pasienService) Find() []pasien.Pasien {
	return pasien.PasienDao.Find()
}

func (p *pasienService) Create(data pasien.Pasien) (*pasien.Pasien, rest_err.APIError) {

	pasienData := data
	id, err := pasien.PasienDao.GetPasienLastIDWithGender(data.Jk)
	if err != nil {
		return nil, rest_err.NewInternalServerError("error create", err)
	}
	pasienData.NoPasien = strconv.Itoa(id + 1)
	pasienResponse, err := pasien.PasienDao.Create(pasienData)
	if err != nil {
		return nil, rest_err.NewInternalServerError("error create", err)
	}

	return &pasienResponse, nil
}
