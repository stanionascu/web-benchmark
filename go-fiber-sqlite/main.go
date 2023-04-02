package main

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/template/handlebars"
	_ "github.com/mattn/go-sqlite3"
)

type Film struct {
	Title       string
	ReleaseYear string
}

func main() {
	db, _ := sql.Open("sqlite3", "../db.sqlite3")
	db.SetMaxOpenConns(1)

	views := handlebars.New("./templates", ".html")

	app := fiber.New(fiber.Config{
		Views: views,
	})
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))

	stmt, _ := db.Prepare("SELECT title,release_year FROM film LIMIT 100")
	app.Get("/", func(c *fiber.Ctx) error {
		rows, _ := stmt.Query()
		films := []Film{}
		for rows.Next() {
			film := Film{}
			rows.Scan(&film.Title, &film.ReleaseYear)
			films = append(films, film)
		}
		return c.Render("index", fiber.Map{
			"name":  c.Query("name"),
			"films": films,
		})
	})

	app.Listen("0.0.0.0:3000")
}
