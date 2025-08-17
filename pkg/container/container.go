package container

import (
	"sync"
	"github.com/marcosfrancomarinho/short_url/internal/application/usecase"
	"github.com/marcosfrancomarinho/short_url/internal/domain/gateway"
	"github.com/marcosfrancomarinho/short_url/internal/infrastructure/id_generator"
	"github.com/marcosfrancomarinho/short_url/internal/infrastructure/qr_code_generator"
	"github.com/marcosfrancomarinho/short_url/internal/infrastructure/repository"
	"github.com/marcosfrancomarinho/short_url/internal/infrastructure/short_id_generator"
	"github.com/marcosfrancomarinho/short_url/internal/presentation/controllers"
)

type Container struct {
}

var (
	once     sync.Once
	instance *Container
)

func GetInstance() *Container {
	once.Do(func() {
		instance = &Container{}
	})
	return instance
}

type handlers struct {
	FindorController gateway.HttpController
	SaviorController gateway.HttpController
}

func (c *Container) Dependencies() *handlers {

	urlFindor := repository.NewSqliteURLFindor()
	urlSavior := repository.NewSqliteURLSavior()
	idGenerator := idgenerator.NewUUID()
	shortId := shortidgenerator.NewTerisShortID()
	qrCodeGenerator := qrcodegenerator.NewSkipQrCodeGenerator()

	urlFindorUseCase := usecase.NewURLFindorUseCase(urlFindor)
	urlSaviorUseCase := usecase.NewURLSaviorUseCase(idGenerator, urlSavior, shortId, qrCodeGenerator)

	urlFindorController := controllers.NewURLFindorController(urlFindorUseCase)
	urlSaviorController := controllers.NewURLSaviorController(urlSaviorUseCase)

	return &handlers{
		FindorController: urlFindorController,
		SaviorController: urlSaviorController,
	}
}
