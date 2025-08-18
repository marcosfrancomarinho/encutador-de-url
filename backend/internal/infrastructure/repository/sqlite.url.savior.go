package repository

import (
	"fmt"

	"github.com/marcosfrancomarinho/short_url/internal/domain/entities"
	"github.com/marcosfrancomarinho/short_url/internal/domain/gateway"
	"github.com/marcosfrancomarinho/short_url/internal/infrastructure/database"
)

type SqliteURLSavior struct {
}

func NewSqliteURLSavior() gateway.URLSavior {
	return &SqliteURLSavior{}
}

func (s *SqliteURLSavior) Save(url *entities.URL) error {
	db := database.InitDataBase()

	command, err := db.Prepare("INSERT INTO urls (id, short_id, long_url) VALUES(?, ?, ?)")
	if err != nil {
		fmt.Println(err)
		return err
	}

	if  _, err:= command.Exec(url.GetID(), url.GetShortId(), url.GetLongUrl()); err != nil{
		return  err
	}
	return  nil
}
