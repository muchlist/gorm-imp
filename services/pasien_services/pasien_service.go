package pasien_services

import (
	"github.com/muchlist/erru_utils_go/rest_err"
	"github.com/muchlist/gorm-imp/dao"
	dto2 "github.com/muchlist/gorm-imp/dto"
	"strconv"
)

var (
	PasienService pasienServiceInterface = &pasienService{}
)

type pasienService struct{}

type pasienServiceInterface interface {
	Find(gender string) []dto2.Pasien
	Create(data dto2.PasienRequest) (*dto2.Pasien, rest_err.APIError)
}

func (p *pasienService) Find(gender string) []dto2.Pasien {
	return dao.PasienDao.Find(gender)
}

func (p *pasienService) Create(data dto2.PasienRequest) (*dto2.Pasien, rest_err.APIError) {

	pasienData, err := data.TranslateReqToEntity()
	if err != nil {
		return nil, rest_err.NewInternalServerError("error mapping pegawai response", err)
	}

	// NoPasien membutuhkan penomoran yang berbeda antara pasien laki-laki dan perempuan
	id, err := dao.PasienDao.GetPasienLastIDWithGender(data.Jk)
	if err != nil {
		return nil, rest_err.NewInternalServerError("error create", err)
	}
	pasienData.NoPasien = strconv.Itoa(id + 1)

	pasienResponse, err := dao.PasienDao.Create(*pasienData)
	if err != nil {
		return nil, rest_err.NewInternalServerError("error create", err)
	}

	return &pasienResponse, nil
}
