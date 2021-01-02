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
