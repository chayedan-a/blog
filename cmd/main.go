package main

import (
	"blog/internal/router"
	"blog/internal/service"
)

func main() {
	s := service.New()
	router.InitRouter(s)
}
