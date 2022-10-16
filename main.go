package main

import (
	"github.com/poboisvert/GO_shortener/model"
	"github.com/poboisvert/GO_shortener/server"
)

func main() {
	model.SetupDB()
	server.InitializeApp()
}
