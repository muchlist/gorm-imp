package terapi

import (
	"github.com/mashingan/smapping"
	"github.com/muchlist/gorm-imp/domains/dto"
)

func TranslateResToEntity(data dto.TerapiRequest) (*dto.Terapi, error) {
	terapiEntity := dto.Terapi{}
	err := smapping.FillStruct(&terapiEntity, smapping.MapFields(&data))
	if err != nil {
		return nil, err
	}
	return &terapiEntity, err
}

func TranslateEntityToRes(data dto.Terapi) (*dto.TerapiResponse, error) {
	terapiResponse := dto.TerapiResponse{}
	err := smapping.FillStruct(&terapiResponse, smapping.MapFields(&data))
	if err != nil {
		return nil, err
	}
	return &terapiResponse, err
}
