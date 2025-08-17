package controllers

import (
	"net/http"
	"github.com/marcosfrancomarinho/short_url/internal/application/dto"
	"github.com/marcosfrancomarinho/short_url/internal/application/usecase"
	"github.com/marcosfrancomarinho/short_url/internal/domain/gateway"
)

type URLSaviorController struct {
	urlSaviorUseCase *usecase.URLSaviorUseCase
}

func NewURLSaviorController(urlSaviorUseCase *usecase.URLSaviorUseCase) gateway.HttpController {
	return &URLSaviorController{urlSaviorUseCase: urlSaviorUseCase}
}

type requestBodySaviorURL struct {
	LongURL string `json:"long_url"`
}

func (u *URLSaviorController) Execute(ctx gateway.HttpContext) {
	var body requestBodySaviorURL
	host := ctx.GetHost()

	if err := ctx.GetBody(&body); err != nil {
		ctx.SendError(err)
		return
	}

	payload := dto.RequestURLSaviorDTO{LongUrl: body.LongURL, Host: host}

	output, err := u.urlSaviorUseCase.Save(&payload)
	if err != nil {
		ctx.SendError(err)
		return
	}

	ctx.Send(http.StatusOK, output)
}
