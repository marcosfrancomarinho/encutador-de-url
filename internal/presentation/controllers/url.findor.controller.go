package controllers

import (
	"net/http"
	"github.com/marcosfrancomarinho/short_url/internal/application/dto"
	"github.com/marcosfrancomarinho/short_url/internal/application/usecase"
	"github.com/marcosfrancomarinho/short_url/internal/domain/gateway"
)

type URLFindorController struct {
	urlFindorUseCase *usecase.URLFindorUseCase
}

func NewURLFindorController(urlFindorUseCase *usecase.URLFindorUseCase) gateway.HttpController {
	return &URLFindorController{urlFindorUseCase: urlFindorUseCase}
}

type requestParamFindorURL struct {
	ShortId string `uri:"short_id"`
}

func (u *URLFindorController) Execute(ctx gateway.HttpContext) {
	var param requestParamFindorURL

	if err := ctx.GetParam(&param); err != nil {
		ctx.SendError(err)
		return
	}

	paylod := dto.RequestURLFindorDTO{ShortId: param.ShortId}

	output, err := u.urlFindorUseCase.Find(&paylod)
	if err != nil {
		ctx.SendError(err)
		return
	}
	ctx.Redirect(http.StatusMovedPermanently, output.LongUrl)
}
