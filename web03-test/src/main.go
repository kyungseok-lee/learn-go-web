package main

import (
	"myapp/myapp"
	"myapp/utils"
	"net/http"
)

func main() {
	utils.Hello()
	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
