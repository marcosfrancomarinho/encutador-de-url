package valuesobject

import (
	"errors"
	"strings"
)

type LongUrl struct {
	value string
}

func NewLongUrl(value string) (*LongUrl, error) {
	trimmed, err := validateLongUrl(value)
	if err != nil {
		return nil, err
	}
	return &LongUrl{value: trimmed}, nil
}

func validateLongUrl(value string) (string, error) {
	trimmed := strings.TrimSpace(value)
	if len(trimmed) == 0 {
		return "", errors.New("url longa nao foi informada")
	}
	return trimmed, nil
}

func (l *LongUrl) GetValue() string{
	return  l.value
}