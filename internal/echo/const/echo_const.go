package _const

const CmdApp = `package main

import (
	"log"
	"%s/internal/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}
	a.Run()
}
`

const InternalApp = `package app

import (
	"fmt"

	handler "%s/internal/internal"
	"github.com/labstack/echo/v4"
)

type App struct {
	h    *handler.Handler
	echo *echo.echo
}

func New() (*App, error) {
	a := &App{}
	a.h = handler.New()
	a.echo = echo.New()
	a.echo.GET("/", a.h.StartHandler)
	return a, nil
}

func (a *App) Run() {
	fmt.Println("[SERVER STARTED]")
	a.echo.Logger.Fatal(a.echo.Start(":8080"))
}`

const InternalHandlers = `package handler

import (
	"log"
	"net/http"
	
	"github.com/labstack/echo/v4"
)

func New() *Handler {
	return &Handler{}
}

type Handler struct {
}

func (h *Handler) StartHandler(c echo.Context) error {
	err := c.String(http.StatusOK, "[SERVER STARTED]")
	if err != nil {
		log.Fatal(err)
	}
	return nil
}`
