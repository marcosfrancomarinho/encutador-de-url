package gateway

type HttpController interface{
	Execute(ctx HttpContext)
}