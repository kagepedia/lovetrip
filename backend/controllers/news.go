package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kagepedia/go-next-pj/backend/models"
)

func NewsIndexHandler(w http.ResponseWriter, r *http.Request) {
	news, _ := models.GetAll()
	e, err := json.Marshal(news)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, string(e))
}

func NewsDetailHandler(w http.ResponseWriter, r *http.Request) {
	news, _ := models.FindById(2)
	e, err := json.Marshal(news)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, string(e))
}
