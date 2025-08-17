package gateway

import "github.com/marcosfrancomarinho/short_url/internal/domain/valuesobject"

type ShortIdGenerator interface {
	Generator() (*valuesobject.ShortId, error)
}
