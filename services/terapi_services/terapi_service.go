package terapi_services

import (
	"github.com/muchlist/erru_utils_go/rest_err"
	"github.com/muchlist/gorm-imp/domains/dto"
	"github.com/muchlist/gorm-imp/domains/pegawai"
	"github.com/muchlist/gorm-imp/domains/terapi"
	"math/rand"
	"time"
)

var (
	TerapiService terapiServiceInterface = &terapiService{}
)

type terapiService struct{}

type terapiServiceInterface interface {
	Create(data dto.TerapiRequest) (*dto.Terapi, rest_err.APIError)
	Find() []dto.TerapiResponse
	FindByDateRange(start, end time.Time) []dto.TerapiResponse
}

func (p *terapiService) Create(data dto.TerapiRequest) (*dto.Terapi, rest_err.APIError) {

	terapiData, err := terapi.TranslateResToEntity(data)
	if err != nil {
		return nil, rest_err.NewInternalServerError("error mapping terapi request ke entity", err)
	}

	terapiData.Tglterapi = time.Now()

	//Pegawai ID dipilih secara random
	pegawaiList := pegawai.PegawaiDao.Find()
	rand.Seed(time.Now().Unix())
	indexRandom := rand.Intn(len(pegawaiList) - 1)
	terapiData.PegawaiID = pegawaiList[indexRandom].ID

	//upah diisi dengan 25% biaya
	terapiData.Upah = terapiData.Biaya * 0.25

	terapiResponse, err := terapi.TerapiDao.Create(*terapiData)
	if err != nil {
		return nil, rest_err.NewInternalServerError("error create", err)
	}

	return &terapiResponse, nil
}

func (p *terapiService) Find() []dto.TerapiResponse {
	var terapiListDisplay []dto.TerapiResponse
	terapiList := terapi.TerapiDao.Find()
	for _, t := range terapiList {
		terapiDisplay, err := terapi.TranslateEntityToRes(t)
		if err != nil {
			continue
		}
		terapiListDisplay = append(terapiListDisplay, *terapiDisplay)
	}

	return terapiListDisplay
}

func (p *terapiService) FindByDateRange(start, end time.Time) []dto.TerapiResponse {
	var terapiListDisplay []dto.TerapiResponse

	terapiList := terapi.TerapiDao.FindByDateRange(start, end)

	for _, t := range terapiList {
		terapiDisplay, err := terapi.TranslateEntityToRes(t)
		if err != nil {
			continue
		}
		terapiListDisplay = append(terapiListDisplay, *terapiDisplay)
	}

	return terapiListDisplay
}
