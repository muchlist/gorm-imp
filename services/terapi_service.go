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
	Find() []dto2.TerapiResponse
	FindByDateRange(start, end time.Time) []dto2.TerapiResponse
}

func (p *terapiService) Create(data dto2.TerapiRequest) (*dto2.Terapi, rest_err.APIError) {

	terapiData, err := data.TranslateToEntity()
	if err != nil {
		return nil, rest_err.NewInternalServerError("error mapping terapi request ke entity", err)
	}

	terapiData.Tglterapi = time.Now()

	//Pegawai ID dipilih secara random
	pegawaiList := dao.PegawaiDao.Find()
	rand.Seed(time.Now().Unix())
	indexRandom := rand.Intn(len(pegawaiList) - 1)
	terapiData.PegawaiID = pegawaiList[indexRandom].ID

	//upah diisi dengan 25% biaya
	terapiData.Upah = terapiData.Biaya * 0.25

	terapiResponse, err := dao.TerapiDao.Create(*terapiData)
	if err != nil {
		return nil, rest_err.NewInternalServerError("error create", err)
	}

	return &terapiResponse, nil
}

func (p *terapiService) Find() []dto2.TerapiResponse {
	var terapiListDisplay []dto2.TerapiResponse
	terapiList := dao.TerapiDao.Find()
	for _, t := range terapiList {
		terapiDisplay, err := t.TranslateToRes()
		if err != nil {
			continue
		}
		terapiListDisplay = append(terapiListDisplay, *terapiDisplay)
	}

	return terapiListDisplay
}

func (p *terapiService) FindByDateRange(start, end time.Time) []dto2.TerapiResponse {
	var terapiListDisplay []dto2.TerapiResponse

	terapiList := dao.TerapiDao.FindByDateRange(start, end)

	for _, t := range terapiList {
		terapiDisplay, err := t.TranslateToRes()
		if err != nil {
			continue
		}
		terapiListDisplay = append(terapiListDisplay, *terapiDisplay)
	}

	return terapiListDisplay
}