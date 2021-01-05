package services

import (
	"github.com/muchlist/erru_utils_go/rest_err"
	"github.com/muchlist/gorm-imp/dao"
	dto2 "github.com/muchlist/gorm-imp/dto"
	"math/rand"
	"time"
)

var (
	TerapiService terapiServiceInterface = &terapiService{}
)

type terapiService struct{}

type terapiServiceInterface interface {
	Create(data dto2.TerapiRequest) (*dto2.Terapi, rest_err.APIError)
	Find() ([]dto2.TerapiResponse, rest_err.APIError)
	FindByDateRange(start, end time.Time) ([]dto2.TerapiResponse, rest_err.APIError)
}

func (p *terapiService) Create(data dto2.TerapiRequest) (*dto2.Terapi, rest_err.APIError) {

	terapiData, err := data.TranslateToEntity()
	if err != nil {
		return nil, rest_err.NewInternalServerError("error mapping terapi request ke entity", err)
	}

	terapiData.Tglterapi = time.Now()

	//Pegawai ID dipilih secara random
	pegawaiList, err := dao.PegawaiDao.Find()
	if err != nil {
		return nil, rest_err.NewInternalServerError("gagal query pegawwai", err)
	}
	rand.Seed(time.Now().Unix())
	indexRandom := rand.Intn(len(pegawaiList) - 1)
	terapiData.PegawaiID = pegawaiList[indexRandom].ID

	//upah diisi dengan 25% biaya
	terapiData.Upah = terapiData.Biaya * 0.25

	terapiResponse, err := dao.TerapiDao.Create(*terapiData)
	if err != nil {
		return nil, rest_err.NewInternalServerError("error create", err)
	}

	return terapiResponse, nil
}

func (p *terapiService) Find() ([]dto2.TerapiResponse, rest_err.APIError) {
	var terapiListDisplay []dto2.TerapiResponse
	terapiList, err := dao.TerapiDao.Find()
	if err != nil {
		return nil, rest_err.NewInternalServerError("gagal query terapi", err)
	}
	for _, t := range terapiList {
		terapiDisplay, err := t.TranslateToRes()
		if err != nil {
			continue
		}
		terapiListDisplay = append(terapiListDisplay, *terapiDisplay)
	}

	return terapiListDisplay, nil
}

func (p *terapiService) FindByDateRange(start, end time.Time) ([]dto2.TerapiResponse, rest_err.APIError) {
	var terapiListDisplay []dto2.TerapiResponse

	terapiList, err := dao.TerapiDao.FindByDateRange(start, end)
	if err != nil {
		return nil, rest_err.NewInternalServerError("gagal query terapi", err)
	}

	for _, t := range terapiList {
		terapiDisplay, err := t.TranslateToRes()
		if err != nil {
			continue
		}
		terapiListDisplay = append(terapiListDisplay, *terapiDisplay)
	}

	return terapiListDisplay, nil
}
