package main

import (
	"fmt"
	"log"

	"github.com/1garo/daedalus/bootstrap"
	"github.com/1garo/daedalus/pkg/env"
)

func main() {
	app := bootstrap.NewApplication()

	host := env.GetEnv("APP_HOST", "localhost")
	port := env.GetEnv("APP_PORT", "3000")
	listen := fmt.Sprintf("%s:%s", host, port)

	log.Fatal(app.Listen(listen))
}
