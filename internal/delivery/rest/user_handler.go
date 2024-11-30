package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dhavisiregar/go-restaurant-app/internal/model"
	"github.com/dhavisiregar/go-restaurant-app/internal/tracing"
	"github.com/labstack/echo/v4"
)

func(h *handler) RegisterUser(c echo.Context) error {
    ctx, span := tracing.CreateSpan(c.Request().Context(), "RegisterUser")
	defer span.End()

	var request model.RegisterRequest
    err := json.NewDecoder(c.Request().Body).Decode(&request)
    if err!= nil {
        fmt.Printf("got error %s\n", err.Error())

        return c.JSON(http.StatusInternalServerError, map[string]interface{}{
            "error": err.Error(),
        })
    }

    userData, err := h.restoUsecase.RegisterUser(ctx, request)
	if err!= nil {
        fmt.Printf("got error %s\n", err.Error())

        return c.JSON(http.StatusInternalServerError, map[string]interface{}{
            "error": err.Error(),
        })
    }

	return c.JSON(http.StatusOK, map[string]interface{}{
        "data": userData,
    })
}

func (h *handler) Login(c echo.Context) error {
    ctx, span := tracing.CreateSpan(c.Request().Context(), "Login")
	defer span.End()

    var request model.LoginRequest
    err := json.NewDecoder(c.Request().Body).Decode(&request)
    if err!= nil {
        fmt.Printf("got error %s\n", err.Error())

        return c.JSON(http.StatusInternalServerError, map[string]interface{}{
            "error": err.Error(),
        })
    }

    sessionData, err := h.restoUsecase.Login(ctx, request)
    if err!= nil {
        fmt.Printf("got error %s\n", err.Error())

        return c.JSON(http.StatusInternalServerError, map[string]interface{}{
            "error": err.Error(),
        })
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "data": sessionData,
    })
}