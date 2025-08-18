package routers

import (
	"github.com/marcosfrancomarinho/short_url/internal/domain/gateway"
	"github.com/marcosfrancomarinho/short_url/internal/shared/container"
)

type Routers struct {
	app gateway.HttpServer
}

func NewRouter(app gateway.HttpServer) *Routers {
	return &Routers{app: app}
}

func (r *Routers) Register(container *container.Container) {
	dependencies := container.Dependencies()

	r.app.On("GET", "/:short_id", dependencies.FindorController)

	r.app.On("POST", "/", dependencies.SaviorController)

}
