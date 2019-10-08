package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

func IndexPost(context echo.Context) error {
	return context.String(http.StatusCreated, "Hello, World POST")
}

func IndexGet(context echo.Context) error {
	return context.String(http.StatusOK, "Hello, World GET")
}

func IndexPut(context echo.Context) error {
	return context.String(http.StatusAccepted, "Hello, World PUT")
}

func IndexDelete(context echo.Context) error {
	return context.String(http.StatusAccepted, "Hello, World DELETE")

}
