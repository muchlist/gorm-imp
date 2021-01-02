package pegawai_services

import (
	"github.com/muchlist/erru_utils_go/rest_err"
	"github.com/muchlist/gorm-imp/domains/dto"
	"github.com/muchlist/gorm-imp/domains/pegawai"
	"github.com/muchlist/gorm-imp/utils/crypto"
)

var (
	PegawaiService pegawaiServiceInterface = &pegawaiService{}
)

type pegawaiService struct{}

type pegawaiServiceInterface interface {
	Find() []dto.PegawaiResponse
	Create(data dto.PegawaiRequest) (*dto.PegawaiResponse, rest_err.APIError)
}

func (p *pegawaiService) Find() []dto.PegawaiResponse {
	var pegawaiListDisplay []dto.PegawaiResponse
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

func (p *pegawaiService) Create(data dto.PegawaiRequest) (*dto.PegawaiResponse, rest_err.APIError) {
	pegawaiData, err := pegawai.TranslateReqToEntity(data)
	if err != nil {
		return nil, rest_err.NewInternalServerError("gagal mapping pegawaiReq ke pegawai", err)
	}

	// Hash password menggunakan bcrypt
	hashedpassword, _ := crypto.Obj.GenerateHash(pegawaiData.Password)
	pegawaiData.Password = hashedpassword

	pegawaiResponse, err := pegawai.PegawaiDao.Create(*pegawaiData)
	if err != nil {
		return nil, rest_err.NewInternalServerError("error create", err)
	}

	pegawaiDisplay, err := pegawai.TranslateEntityToRes(pegawaiResponse)
	if err != nil {
		return nil, rest_err.NewInternalServerError("error mapping pegawai response", err)
	}

	return pegawaiDisplay, nil
}
