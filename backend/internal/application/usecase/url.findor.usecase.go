package usecase

import (
	"github.com/marcosfrancomarinho/short_url/internal/application/dto"
	"github.com/marcosfrancomarinho/short_url/internal/domain/gateway"
	"github.com/marcosfrancomarinho/short_url/internal/domain/valuesobject"
)

type URLFindorUseCase struct {
	urlFindor gateway.URLFindor
}

func NewURLFindorUseCase(urlFindor gateway.URLFindor) *URLFindorUseCase {
	return &URLFindorUseCase{urlFindor: urlFindor}
}

func (u * URLFindorUseCase) Find(payload *dto.RequestURLFindorDTO) (*dto.ResponseURLFindorDTO, error){
	shortId, err := valuesobject.NewShortId(payload.ShortId)
	if err != nil {
		return nil, err
	}
	url, err:= u.urlFindor.FindByShortId(shortId)
	if err != nil {
		return nil, err
	}
	
	return  &dto.ResponseURLFindorDTO{ID: url.GetID(), LongUrl: url.GetLongUrl()}, nil
}