package entities

import (
	"encoding/base64"
	"errors"
	"fmt"
)

type QrCode struct {
	code string
}

func NewQrCode(code []byte) (*QrCode, error) {
	if len(code) == 0 {
		return nil, errors.New("não foi informado qr code")
	}
	base64QR := base64.StdEncoding.EncodeToString(code)
	urlImg := fmt.Sprintf("data:image/png;base64,%s", base64QR)
	return &QrCode{code: urlImg}, nil
}

func (q *QrCode) GetQrCode() string {
	return q.code
}
