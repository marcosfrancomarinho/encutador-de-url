package http

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/marcosfrancomarinho/short_url/internal/domain/gateway"
)

type GinHttpContext struct {
	ctx *gin.Context
}

func NewGinHttpContext(ctx *gin.Context) gateway.HttpContext {
	return &GinHttpContext{ctx: ctx}
}

func (g *GinHttpContext) Send(status int, datas any) {
	g.ctx.JSON(status, datas)
}

func (g *GinHttpContext) SendError(err error) {
	g.ctx.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
}

func (g *GinHttpContext) GetBody(datas any) error {
	return g.ctx.ShouldBindBodyWithJSON(datas)
}

func (g *GinHttpContext) GetParam(datas any) error {
	return g.ctx.BindUri(datas)
}
func (g *GinHttpContext) GetHost() string {
	host := g.ctx.Request.Host
	scheme := "http"
	if g.ctx.Request.TLS != nil {
		scheme = "https"
	}
	return fmt.Sprintf("%s://%s", scheme, host)
}

func (g *GinHttpContext) Redirect(status int, url string) {
	g.ctx.Redirect(status, url)
}
