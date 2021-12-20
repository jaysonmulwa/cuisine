package main

import (
	"fmt"

	handler "github.com/jaysonmulwa/cuisine/internal/http"
)

type App struct {
}

func (a *App) Run() error {

	handler.SetupRoutes()
	return nil

}

func main() {
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Err")
		fmt.Println(err)
	}

}
