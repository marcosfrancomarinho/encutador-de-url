package gateway

import (
	"github.com/marcosfrancomarinho/short_url/internal/domain/entities"
	"github.com/marcosfrancomarinho/short_url/internal/domain/valuesobject"
)

type URLFindor interface {
	FindByShortId(shortId *valuesobject.ShortId) (*entities.URL, error)
}
