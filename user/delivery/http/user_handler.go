package http

import (
	"encoding/json"
	"fmt"
	"github.com/LayssonENS/go-FastHTTP-api/domain"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"strconv"
)

type UserHandler struct {
	UserUCase domain.UserUseCase
}

func NewUserHandler(router *fasthttprouter.Router, us domain.UserUseCase) {
	handler := &UserHandler{
		UserUCase: us,
	}

	router.GET("/articles/:id", handler.GetByID)
}

func (a *UserHandler) GetByID(ctx *fasthttp.RequestCtx) {
	idP, err := strconv.Atoi(ctx.UserValue("id").(string))
	if err != nil {
		responseError(ctx, "Invalid parameter", fasthttp.StatusBadRequest)
		return
	}

	id := int64(idP)
	art, err := a.UserUCase.GetByID(id)
	if err != nil {
		responseError(ctx, err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	responseJSON(ctx, fasthttp.StatusOK, art)
}

func responseError(ctx *fasthttp.RequestCtx, message string, status int) {
	ctx.Response.SetStatusCode(status)
	fmt.Fprintf(ctx, `{"message": "%s"}`, message)
}

func responseJSON(ctx *fasthttp.RequestCtx, status int, data interface{}) {
	body, _ := json.Marshal(data)
	ctx.Response.SetStatusCode(status)
	ctx.Response.SetBody(body)
	ctx.Response.Header.Set("Content-Type", "application/json")
}
