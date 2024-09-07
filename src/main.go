package main

import (
	"log"
	"src/database"
	"src/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	database.InitDB()

	moviesRoute := app.Group("/movies")
	moviesRoute.Get("/", handlers.GetAllMovies)
	moviesRoute.Get("/:id", handlers.GetMoviebyID)
	moviesRoute.Post("/add", handlers.AddMovie)
	moviesRoute.Put("/add/:id", handlers.UpdateMovie)
	moviesRoute.Delete("/delete/:id", handlers.DeleteMovie)

	log.Fatal(app.Listen(":3000"))
}
