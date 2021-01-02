package pegawai_services

import (
	"github.com/muchlist/erru_utils_go/rest_err"
	"github.com/muchlist/gorm-imp/domains/pegawai"
	"github.com/muchlist/gorm-imp/utils/crypto"
)

var (
	PegawaiService pegawaiServiceInterface = &pegawaiService{}
)

type pegawaiService struct{}

type pegawaiServiceInterface interface {
	Find() []pegawai.PegawaiResponse
	Create(data pegawai.Pegawai) (*pegawai.PegawaiResponse, rest_err.APIError)
}

func (p *pegawaiService) Find() []pegawai.PegawaiResponse {
	var pegawaiListDisplay []pegawai.PegawaiResponse
	pegawaiList := pegawai.PegawaiDao.Find()
	for _, p := range pegawaiList {
		pegawaiDisplay, err := pegawai.TranslateEntityToRes(p)
		if err != nil {
			continue
		}
		pegawaiListDisplay = append(pegawaiListDisplay, *pegawaiDisplay)
	}

	return pegawaiListDisplay
}

func (p *pegawaiService) Create(data pegawai.Pegawai) (*pegawai.PegawaiResponse, rest_err.APIError) {
	pegawaiData := data

	// Hash password menggunakan bcrypt
	hashedpassword, _ := crypto.Obj.GenerateHash(pegawaiData.Password)
	pegawaiData.Password = hashedpassword

	pegawaiResponse, err := pegawai.PegawaiDao.Create(pegawaiData)
	if err != nil {
		return nil, rest_err.NewInternalServerError("error create", err)
	}

	pegawaiDisplay, err := pegawai.TranslateEntityToRes(pegawaiResponse)
	if err != nil {
		return nil, rest_err.NewInternalServerError("error mapping pegawai response", err)
	}

	return pegawaiDisplay, nil
}
