package main

import (
	"fmt"
	"net/http"

	"github.com/khandev-bac/lemon/internals/router"
)

func main() {
	router := router.MainRoute()
	fmt.Println("Server started at port 3000")
	http.ListenAndServe(":3000", router)
}
