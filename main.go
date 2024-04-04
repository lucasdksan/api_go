package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"net/http"
)

func main() {
	config.LoadingEnv()
	fmt.Println(config.Port)
	r := router.Generate()

	fmt.Println("Start server!!")
	http.ListenAndServe(":8080", r)
}
