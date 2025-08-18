package shortidgenerator

import (
	"github.com/marcosfrancomarinho/short_url/internal/domain/gateway"
	"github.com/marcosfrancomarinho/short_url/internal/domain/valuesobject"
	"github.com/teris-io/shortid"
)

type TerisShortID struct{

}

func NewTerisShortID() gateway.ShortIdGenerator{
	return  &TerisShortID{}
}

func (t *TerisShortID) Generator() (*valuesobject.ShortId, error){
		id, err:= shortid.Generate()
		if err != nil{
			return  nil, err
		}
		return  valuesobject.NewShortId(id)
}