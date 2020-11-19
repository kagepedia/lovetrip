package models

import (
	"log"
	"time"

	"github.com/kagepedia/go-next-pj/backend/middleware"
)

type T_news struct {
	Pk        int64
	Title     string
	Category  int8
	Body      string
	CreatedAt time.Time
	UpdateAt  time.Time
}

type News []T_news

func GetAll() ([]T_news, error) {
	db := middleware.DbConnect()
	defer db.Close()

	// 構造体マッピング
	var news T_news

	// 配列宣言
	var res News

	rows, err := db.Query("SELECT * FROM t_news")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&news.Pk, &news.Title, &news.Category, &news.Body, &news.CreatedAt, &news.UpdateAt)
		if err != nil {
			log.Fatal(err)
		} else {
			res = append(res, news)
		}
	}

	return res, nil
}

func FindById(pk int) (T_news, error) {
	db := middleware.DbConnect()
	defer db.Close()

	var news T_news
	err := db.QueryRow("SELECT * FROM t_news WHERE pk = ?", pk).Scan(&news.Pk, &news.Title, &news.Category, &news.Body, &news.CreatedAt, &news.UpdateAt)
	if err != nil {
		log.Fatal(err)
	}
	return news, nil
}
