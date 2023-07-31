package main

import (
	http "github.com/LayssonENS/go-FastHTTP-api/user/delivery/http"
	userRepo "github.com/LayssonENS/go-FastHTTP-api/user/repository"
	userUsecase "github.com/LayssonENS/go-FastHTTP-api/user/usecase"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func main() {

	userRepository := userRepo.NewUserRepository(nil)
	uUseCase := userUsecase.NewUserUseCase(userRepository, 10)

	router := fasthttprouter.New()
	http.NewUserHandler(router, uUseCase)

	fasthttp.ListenAndServe(":8080", router.Handler)
}
