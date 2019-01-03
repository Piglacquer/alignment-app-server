package server

import (
	"fmt"
	"log"
	"net/http"
)

type App struct {
	Name string
}

func (a App) Run() {
	fmt.Println("im logging")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
