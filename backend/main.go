package main

import (
	"github.com/marcosfrancomarinho/short_url/internal/infrastructure/http"
	"github.com/marcosfrancomarinho/short_url/internal/presentation/routers"
	"github.com/marcosfrancomarinho/short_url/internal/shared/container"
)

func main() {
	server := http.NewGinHttpSever()
	routers := routers.NewRouter(server)
	container := container.GetInstance()
	routers.Register(container)
	server.Listen(3030)

}
