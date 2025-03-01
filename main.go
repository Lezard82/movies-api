package main

import (
	"github.com/Lezard82/movies-api/config"
	"github.com/Lezard82/movies-api/infrastructure/api"
)

func main() {
	config.LoadEnv()
	api.StartServer()
}
