package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type Film struct {
	Index       int
	Title       string
	ReleaseYear string
}

func main() {
	db, err := sql.Open("sqlite3", "../db.sqlite3")
	if err != nil {
		log.Fatal("Failed to open the db: ", err)
	}
	db.SetMaxOpenConns(1)
	r := gin.New()
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		rows, _ := db.Query("SELECT title,release_year FROM film LIMIT 100")
		films := []Film{}
		index := 0
		for rows.Next() {
			film := Film{Index: index}
			rows.Scan(&film.Title, &film.ReleaseYear)
			films = append(films, film)
			index += 1
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"name":  c.Query("name"),
			"films": films,
		})
	})
	r.Run("0.0.0.0:3000")
}
