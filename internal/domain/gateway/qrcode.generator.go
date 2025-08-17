package gateway

import "github.com/marcosfrancomarinho/short_url/internal/domain/entities"

type QrCodeGenerator interface{
	Generate(url *entities.URL) (*entities.QrCode, error)
}