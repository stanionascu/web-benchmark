package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Film struct {
	Title       string
	ReleaseYear string
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	// export DATABASE="user=postgres password=<password> dbname=sakila sslmode=disable"
	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE"))
	db.DB.SetConnMaxLifetime(0)
	db.DB.SetMaxOpenConns(48)
	db.DB.SetMaxIdleConns(0)
	if err != nil {
		fmt.Println("Failed to connect to the db:", err)
		os.Exit(1)
	}
	defer db.Close()

	r := gin.New()
	r.Use(gzip.Gzip(gzip.BestSpeed))

	stmt, err := db.Prepare("SELECT title,release_year FROM film LIMIT 100")
	if err != nil {
		fmt.Println("Failed to prepare statement:", err)
		os.Exit(1)
	}

	r.GET("/", func(c *gin.Context) {
		rows, err := stmt.Query()
		if err != nil {
			fmt.Println("Failed to query films from database:", err)
			c.Status(http.StatusInternalServerError)
		} else {
			films := []Film{}
			for rows.Next() {
				film := Film{}
				rows.Scan(&film.Title, &film.ReleaseYear)
				films = append(films, film)
			}
			c.Header("Content-Type", "text/html; charset=utf-8")
			index(c.Query("name"), films).Render(context.Background(), c.Writer)
		}
	})
	r.Run("0.0.0.0:3000")
}
