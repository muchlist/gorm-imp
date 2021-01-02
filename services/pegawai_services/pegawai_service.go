package pegawai_services

import (
	"github.com/muchlist/erru_utils_go/rest_err"
	"github.com/muchlist/gorm-imp/domains/pegawai"
)

var (
	PegawaiService pegawaiServiceInterface = &pegawaiService{}
)

type pegawaiService struct{}

type pegawaiServiceInterface interface {
	Find() []pegawai.Pegawai
	Create(data pegawai.Pegawai) (*pegawai.Pegawai, rest_err.APIError)
}

func (p *pegawaiService) Find() []pegawai.Pegawai {
	return pegawai.PegawaiDao.Find()
}

func (p *pegawaiService) Create(data pegawai.Pegawai) (*pegawai.Pegawai, rest_err.APIError) {

	pegawaiData := data
	pegawaiResponse, err := pegawai.PegawaiDao.Create(pegawaiData)
	if err != nil {
		return nil, rest_err.NewInternalServerError("error create", err)
	}

	return &pegawaiResponse, nil
}
