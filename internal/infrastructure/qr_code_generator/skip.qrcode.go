package qrcodegenerator

import (
	"github.com/marcosfrancomarinho/short_url/internal/domain/entities"
	"github.com/marcosfrancomarinho/short_url/internal/domain/gateway"
	"github.com/skip2/go-qrcode"
)

type SkipQrCodeGenerator struct {
}

func NewSkipQrCodeGenerator() gateway.QrCodeGenerator {
	return &SkipQrCodeGenerator{}
}

func (s *SkipQrCodeGenerator) Generate(url *entities.URL) (*entities.QrCode, error) {
	img, err := qrcode.Encode(url.GetLongUrl(), qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}
	return entities.NewQrCode(img)
}
