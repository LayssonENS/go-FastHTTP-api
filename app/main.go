package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/LayssonENS/go-FastHTTP-api/config"
	"github.com/LayssonENS/go-FastHTTP-api/config/database"
	"github.com/LayssonENS/go-FastHTTP-api/user/delivery/http"
	userRepo "github.com/LayssonENS/go-FastHTTP-api/user/repository"
	userUsecase "github.com/LayssonENS/go-FastHTTP-api/user/usecase"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func main() {
	var db *sql.DB
	var err error

	switch config.GetEnv().DbConfig.DbType {
	case "postgres":
		db, err = database.NewPostgresConnection()
	case "sqlite":
		db, err = database.NewSqliteConnection()
	default:
		log.Fatalf("Tipo de banco de dados desconhecido: %s", config.GetEnv().DbConfig.DbType)
	}

	if err != nil {
		log.Fatalf("Falha ao se conectar ao banco de dados: %v", err)
	}

	userRepository := userRepo.NewUserRepository(db)
	userUseCase := userUsecase.NewUserUseCase(userRepository, 10)

	router := fasthttprouter.New()
	http.NewUserHandler(router, userUseCase)

	err = fasthttp.ListenAndServe(fmt.Sprintf(":%s", config.GetEnv().Port), router.Handler)
	if err != nil {
		log.Fatalln(err)
		return
	}
}
