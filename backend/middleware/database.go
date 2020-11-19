package middleware

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func DbConnect() *sql.DB {
	err := godotenv.Load("/go/src/github.com/kagepedia/go-next-pj/backend/config/.env")
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_DATABASE")+"?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
