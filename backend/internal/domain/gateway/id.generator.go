package gateway

import "github.com/marcosfrancomarinho/short_url/internal/domain/valuesobject"

type IdGenerator interface {
	Generate() (*valuesobject.ID, error)
}
