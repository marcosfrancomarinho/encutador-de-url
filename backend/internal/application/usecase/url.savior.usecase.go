package usecase

import (
	"github.com/marcosfrancomarinho/short_url/internal/application/dto"
	"github.com/marcosfrancomarinho/short_url/internal/domain/entities"
	"github.com/marcosfrancomarinho/short_url/internal/domain/gateway"
	"github.com/marcosfrancomarinho/short_url/internal/domain/valuesobject"
)

type URLSaviorUseCase struct {
	idGenerator      gateway.IdGenerator
	urlSavior        gateway.URLSavior
	shortIdGenerator gateway.ShortIdGenerator
	qrCodeGenerator  gateway.QrCodeGenerator
}

func NewURLSaviorUseCase(
	idGenerator gateway.IdGenerator,
	urlSavior gateway.URLSavior,
	shortIdGenerator gateway.ShortIdGenerator,
	qrCodeGenerator gateway.QrCodeGenerator,
) *URLSaviorUseCase {
	return &URLSaviorUseCase{idGenerator: idGenerator, urlSavior: urlSavior, shortIdGenerator: shortIdGenerator, qrCodeGenerator: qrCodeGenerator}
}

func (u *URLSaviorUseCase) Save(paylod *dto.RequestURLSaviorDTO) (*dto.ResponseURLSaviorDTO, error) {
	id, err := u.idGenerator.Generate()
	if err != nil {
		return nil, err
	}

	shortId, err := u.shortIdGenerator.Generator()
	if err != nil {
		return nil, err
	}

	longUrl, err := valuesobject.NewLongUrl(paylod.LongUrl)
	if err != nil {
		return nil, err
	}

	url := entities.NewURL(id, shortId, longUrl)

	if err := u.urlSavior.Save(url); err != nil {
		return nil, err
	}

	shortUrl := url.GetShortUrl(paylod.Host)

	qrCode, err := u.qrCodeGenerator.Generate(url)
	if err != nil {
		return nil, err
	}

	return &dto.ResponseURLSaviorDTO{
		UrlShort: shortUrl,
		ID:       url.GetID(),
		QrCode:   qrCode.GetQrCode(),
	}, nil
}
