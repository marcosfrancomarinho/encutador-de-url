package gateway

import "github.com/marcosfrancomarinho/short_url/internal/domain/entities"

type URLSavior interface {
	Save(url *entities.URL) error
}
