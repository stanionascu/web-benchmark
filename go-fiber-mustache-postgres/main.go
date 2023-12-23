package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/template/mustache/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Film struct {
	Index       int
	Title       string
	ReleaseYear string
}

func main() {
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

	views := mustache.New("./templates", ".mustache")

	app := fiber.New(fiber.Config{
		Views: views,
	})
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))

	stmt, err := db.Prepare("SELECT title,release_year FROM film LIMIT 100")
	if err != nil {
		fmt.Println("Failed to prepare statement:", err)
		os.Exit(1)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		var rows *sql.Rows
		var films []Film
		rows, err = stmt.Query()
		if err != nil {
			fmt.Println("Failed to query films from database:", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Error occurred while querying from DB")
		} else {
			index := 1
			for rows.Next() {
				film := Film{}
				rows.Scan(&film.Title, &film.ReleaseYear)
				film.Index = index
				films = append(films, film)
				index++
			}
		}
		return c.Render("index", fiber.Map{
			"name":  c.Query("name"),
			"films": films,
		})
	})

	err = app.Listen("0.0.0.0:3000")
	if err != nil {
		fmt.Println("Failed to bind to port:", err)
		os.Exit(1)
	}
}
