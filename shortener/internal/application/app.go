package application

import "net/http"

type App struct {
	httpRoutes []http.Handler
}

func NewApp() *App {
	return &App{}
}

func (a *App) Init() {

}

func (a *App) Run() {

}
