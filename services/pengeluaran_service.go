package services

import (
	"github.com/muchlist/erru_utils_go/rest_err"
	"github.com/muchlist/gorm-imp/dao"
	dto2 "github.com/muchlist/gorm-imp/dto"
	"time"
)

var (
	PengeluaranService pengeluaranServiceInterface = &pengeluaranService{}
)

type pengeluaranService struct{}

type pengeluaranServiceInterface interface {
	Find() ([]dto2.Pengeluaran, rest_err.APIError)
	Create(data dto2.PengeluaranRequest) (*dto2.Pengeluaran, rest_err.APIError)
}

func (p *pengeluaranService) Find() ([]dto2.Pengeluaran, rest_err.APIError) {
	pengeluaranList, err := dao.PengeluaranDao.Find()
	if err != nil {
		return nil, rest_err.NewInternalServerError("gagal query pengeluaran", err)
	}
	return pengeluaranList, nil
}

func (p *pengeluaranService) Create(data dto2.PengeluaranRequest) (*dto2.Pengeluaran, rest_err.APIError) {
	pengeluaranData, err := data.TranslateToEntity()
	if err != nil {
		return nil, rest_err.NewInternalServerError("gagal mapping pengeluaranRequest ke pengeluaran", err)
	}
	pengeluaranData.Tanggal = time.Now()
	pengeluaranResponse, err := dao.PengeluaranDao.Create(*pengeluaranData)
	if err != nil {
		return nil, rest_err.NewInternalServerError("error create", err)
	}

	return pengeluaranResponse, nil
}
