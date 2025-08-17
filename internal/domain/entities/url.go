package entities

import (
	"fmt"
	"github.com/marcosfrancomarinho/short_url/internal/domain/valuesobject"
)

type URL struct {
	id      *valuesobject.ID
	shortId *valuesobject.ShortId
	longUrl *valuesobject.LongUrl
}


func NewURL(
	id *valuesobject.ID,
	shortId *valuesobject.ShortId,
	longUrl *valuesobject.LongUrl,
) *URL {
	return &URL{id: id, shortId: shortId, longUrl: longUrl}
}

func (u *URL) GetID() string {
	return u.id.GetValue()
}
func (u *URL) GetShortId() string {
	return u.shortId.GetValue()
}
func (u *URL) GetLongUrl() string {
	return u.longUrl.GetValue()
}
func (u *URL) GetShortUrl(host string) string {
	return fmt.Sprintf("%s/%s", host, u.shortId.GetValue())
}

