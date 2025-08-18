package valuesobject

import (
	"errors"
	"strings"
)

type ID struct {
	value string
}

func NewID(value string) (*ID, error) {
	trimmed, err := validateId(value)
	if err != nil {
		return nil, err
	}
	return &ID{value: trimmed}, nil
}

func validateId(value string) (string, error) {
	trimmed := strings.TrimSpace(value)
	if len(trimmed) == 0 {
		return "", errors.New("id não informado")
	}
	return trimmed, nil
}


func (i * ID) GetValue() string{
	return  i.value
}