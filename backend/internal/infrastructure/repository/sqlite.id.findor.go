package repository

import (
	"github.com/marcosfrancomarinho/short_url/internal/domain/entities"
	"github.com/marcosfrancomarinho/short_url/internal/domain/gateway"
	"github.com/marcosfrancomarinho/short_url/internal/domain/valuesobject"
	"github.com/marcosfrancomarinho/short_url/internal/infrastructure/database"
)

type SqliteURLFindor struct {
}

func NewSqliteURLFindor() gateway.URLFindor {
	return &SqliteURLFindor{}
}

type url struct {
	id       string
	short_id string
	long_url string
}

func (s *SqliteURLFindor) FindByShortId(shorId *valuesobject.ShortId) (*entities.URL, error) {
	db := database.InitDataBase()
	var url url
	command, err := db.Prepare("SELECT id, short_id, long_url FROM urls WHERE short_id = ?")
	if err != nil {
		return nil, err
	}
	if err := command.QueryRow(shorId.GetValue()).Scan(&url.id, &url.short_id, &url.long_url); err != nil {
		return nil, err
	}

	id, err := valuesobject.NewID(url.id)
	if err != nil {
		return nil, err
	}

	shortId, err := valuesobject.NewShortId(url.short_id)
	if err != nil {
		return nil, err
	}

	longUrl, err := valuesobject.NewLongUrl(url.long_url)
	if err != nil {
		return nil, err
	}

	return entities.NewURL(id, shortId, longUrl), nil
}
