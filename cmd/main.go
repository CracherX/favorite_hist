package main

import (
	"github.com/CracherX/favorite_hist/pkg/app"
	"log"
)

func main() {
	App, err := app.New()
	if err != nil {
		log.Fatalf("Ошибка создания экземпляра приложения")
	}
	App.Run()
}
