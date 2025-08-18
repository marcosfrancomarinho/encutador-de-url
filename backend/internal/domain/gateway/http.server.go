package gateway

type HttpServer interface {
	On(method string, path string, controller HttpController)
	Listen(port int) error
}
