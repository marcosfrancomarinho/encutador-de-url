package idgenerator

import (
	"github.com/google/uuid"
	"github.com/marcosfrancomarinho/short_url/internal/domain/gateway"
	"github.com/marcosfrancomarinho/short_url/internal/domain/valuesobject"
)

type UUID struct{}

func NewUUID() gateway.IdGenerator {
	return &UUID{}
}

func (u *UUID) Generate() (*valuesobject.ID, error) {
	id := uuid.New().String()
	return valuesobject.NewID(id)
}
