package main

import (
	"crypto/rand"
	"crypto/rsa"
	"time"

	"github.com/dhavisiregar/go-restaurant-app/internal/database"
	"github.com/dhavisiregar/go-restaurant-app/internal/delivery/rest"
	"github.com/dhavisiregar/go-restaurant-app/internal/logger"
	mRepo "github.com/dhavisiregar/go-restaurant-app/internal/repository/menu"
	oRepo "github.com/dhavisiregar/go-restaurant-app/internal/repository/order"
	uRepo "github.com/dhavisiregar/go-restaurant-app/internal/repository/user"
	"github.com/dhavisiregar/go-restaurant-app/internal/tracing"
	rUsecase "github.com/dhavisiregar/go-restaurant-app/internal/usecase/resto"
	"github.com/labstack/echo/v4"
)

const (
	dbAddress = "host=localhost port=5432 user=postgres password=redline2 dbname=go_resto_app sslmode=disable"
)

func main() {
	logger.Init()
	tracing.Init("http://localhost:14268/api/trace")
	e := echo.New()

	db := database.GetDB(dbAddress)
	secret := "AES256Key-32Characters1234567890"
	signKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err!= nil {
        panic(err)
    }

	menuRepo := mRepo.GetRepository(db)
	orderRepo := oRepo.GetRepository(db)
	userRepo, err := uRepo.GetRepository(db, secret, 1, 64*1024, 4, 32, signKey, 60 * time.Second)
	if err != nil {
		panic(err)
	}

	restoUsecase := rUsecase.GetUsecase(menuRepo, orderRepo, userRepo)

	h := rest.NewHandler(restoUsecase)

	rest.LoadMiddleware(e)
	rest.LoadRoutes(e, h)

	e.Logger.Fatal(e.Start(":14045"))
}


