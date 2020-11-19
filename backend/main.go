package main

import (
	"net/http"

	"github.com/kagepedia/go-next-pj/backend/controllers"
)

func main() {
	http.HandleFunc("/news", controllers.NewsIndexHandler)
	http.HandleFunc("/news/detail", controllers.NewsDetailHandler)
	http.ListenAndServe(":8888", nil)
}
