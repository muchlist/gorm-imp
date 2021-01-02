package terapi

import "github.com/mashingan/smapping"

func TranslateResToEntity(data TerapiRequest) (*Terapi, error) {
	terapiEntity := Terapi{}
	err := smapping.FillStruct(&terapiEntity, smapping.MapFields(&data))
	if err != nil {
		return nil, err
	}
	return &terapiEntity, err
}

func TranslateEntityToRes(data Terapi) (*TerapiResponse, error) {
	terapiResponse := TerapiResponse{}
	err := smapping.FillStruct(&terapiResponse, smapping.MapFields(&data))
	if err != nil {
		return nil, err
	}
	return &terapiResponse, err
}
