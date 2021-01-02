package pengeluaran_services

import (
	"github.com/muchlist/erru_utils_go/rest_err"
	"github.com/muchlist/gorm-imp/domains/dto"
	"github.com/muchlist/gorm-imp/domains/pengeluaran"
	"time"
)

var (
	PengeluaranService pengeluaranServiceInterface = &pengeluaranService{}
)

type pengeluaranService struct{}

type pengeluaranServiceInterface interface {
	Find() []dto.Pengeluaran
	Create(data dto.PengeluaranRequest) (*dto.Pengeluaran, rest_err.APIError)
}

func (p *pengeluaranService) Find() []dto.Pengeluaran {
	pengeluaranList := pengeluaran.PengeluaranDao.Find()
	return pengeluaranList
}

func (p *pengeluaranService) Create(data dto.PengeluaranRequest) (*dto.Pengeluaran, rest_err.APIError) {
	pengeluaranData, err := pengeluaran.TranslateReqToEntity(data)
	if err != nil {
		return nil, rest_err.NewInternalServerError("gagal mapping pengeluaranRequest ke pengeluaran", err)
	}
	pengeluaranData.Tanggal = time.Now()
	pengeluaranResponse, err := pengeluaran.PengeluaranDao.Create(*pengeluaranData)
	if err != nil {
		return nil, rest_err.NewInternalServerError("error create", err)
	}

	return &pengeluaranResponse, nil
}
