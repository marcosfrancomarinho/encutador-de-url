package valuesobject

import (
	"errors"
	"strings"
)

type ShortId struct {
	value string
}

func NewShortId(value string) (*ShortId, error) {
	trimmed, err := validateShortId(value)
	if err != nil {
		return nil, err
	}
	return &ShortId{value: trimmed}, nil
}

func validateShortId(value string) (string, error) {
	trimmed := strings.TrimSpace(value)
	if len(trimmed) == 0 {
		return "", errors.New("short id não foi informado")
	}
	return trimmed, nil
}

func (s * ShortId) GetValue() string{
	return  s.value
}
