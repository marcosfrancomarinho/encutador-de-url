package dto

type ResponseURLSaviorDTO struct{
	ID string `json:"id"`
	UrlShort string `json:"url_short"`
	QrCode string `json:"qr_code"`
}