package gateway

type HttpContext interface {
	Send(status int, datas any)
	SendError(err error)
	Redirect(status int, url string)
	GetBody(datas any) error
	GetParam(datas any) error
	GetHost()string
}
