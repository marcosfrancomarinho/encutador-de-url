package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/marcosfrancomarinho/short_url/internal/domain/gateway"
)

type GinHttpSever struct {
	app *gin.Engine
}

func NewGinHttpSever() gateway.HttpServer {
	app := gin.Default()
	return &GinHttpSever{app: app}
}

func (g *GinHttpSever) On(method string, path string, controller gateway.HttpController) {
	g.app.Handle(method, path, func(ctx *gin.Context) {
		httpContext := NewGinHttpContext(ctx)
		controller.Execute(httpContext)
	})

}

func (g *GinHttpSever) Listen(port int) error {
	return g.app.Run(fmt.Sprintf(":%d", port))
}
