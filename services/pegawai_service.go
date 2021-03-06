package services

import (
	"github.com/muchlist/erru_utils_go/rest_err"
	"github.com/muchlist/gorm-imp/dao"
	dto2 "github.com/muchlist/gorm-imp/dto"
	"github.com/muchlist/gorm-imp/utils/crypto"
)

var (
	PegawaiService pegawaiServiceInterface = &pegawaiService{}
)

type pegawaiService struct{}

type pegawaiServiceInterface interface {
	Find() ([]dto2.PegawaiResponse, rest_err.APIError)
	Create(data dto2.PegawaiRequest) (*dto2.PegawaiResponse, rest_err.APIError)
}

func (p *pegawaiService) Find() ([]dto2.PegawaiResponse, rest_err.APIError) {
	var pegawaiListDisplay []dto2.PegawaiResponse
	pegawaiList, err := dao.PegawaiDao.Find()
	if err != nil {
		return nil, rest_err.NewInternalServerError("gagal query pegawai", err)
	}
	for _, p := range pegawaiList {
		pegawaiDisplay, err := p.TranslateToResponse()
		if err != nil {
			continue
		}
		pegawaiListDisplay = append(pegawaiListDisplay, *pegawaiDisplay)
	}

	return pegawaiListDisplay, nil
}

func (p *pegawaiService) Create(data dto2.PegawaiRequest) (*dto2.PegawaiResponse, rest_err.APIError) {
	pegawaiData, err := data.TranslateToEntity()
	if err != nil {
		return nil, rest_err.NewInternalServerError("gagal mapping pegawaiReq ke pegawai", err)
	}

	// Hash password menggunakan bcrypt
	hashedpassword, _ := crypto.Obj.GenerateHash(pegawaiData.Password)
	pegawaiData.Password = hashedpassword

	pegawaiResponse, err := dao.PegawaiDao.Create(*pegawaiData)
	if err != nil {
		return nil, rest_err.NewInternalServerError("error create", err)
	}

	pegawaiDisplay, err := pegawaiResponse.TranslateToResponse()
	if err != nil {
		return nil, rest_err.NewInternalServerError("error mapping pegawai response", err)
	}

	return pegawaiDisplay, nil
}
