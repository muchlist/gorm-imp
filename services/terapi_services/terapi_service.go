package terapi_services

import (
	"github.com/muchlist/erru_utils_go/rest_err"
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
	Create(data terapi.TerapiRequest) (*terapi.Terapi, rest_err.APIError)
}

func (p *terapiService) Create(data terapi.TerapiRequest) (*terapi.Terapi, rest_err.APIError) {

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
