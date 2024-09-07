package handlers

import (
	"src/database"
	"src/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllMovies(c *fiber.Ctx) error {
	var movies []models.Movie
	database.DB.Find(&movies)
	return c.JSON(movies)
}

func GetMoviebyID(c *fiber.Ctx) error {
	id := c.Params("id")
	var movie models.Movie
	result := database.DB.First(&movie, id)
	if result.Error != nil {
		return c.Status(404).SendString("Movie not found")
	}
	return c.JSON(movie)
}

func AddMovie(c *fiber.Ctx) error {
	var movie models.Movie
	if err := c.BodyParser((&movie)); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}
	database.DB.Create(&movie)
	return c.JSON(movie)
}

func UpdateMovie(c *fiber.Ctx) error {
	id := c.Params(("id"))
	var movie models.Movie
	if err := database.DB.First(&movie, id).Error; err != nil {
		return c.Status(404).SendString("Movie not found")
	}
	if err := c.BodyParser(&movie); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}
	database.DB.Save(&movie)
	return c.JSON(movie)
}

func DeleteMovie(c *fiber.Ctx) error {
	id := c.Params("id")
	var movie models.Movie
	if err := database.DB.First(&movie, id).Error; err != nil {
		return c.Status(404).SendString("Movie not found")
	}
	database.DB.Delete(&movie)
	return c.Status(204).SendString("Movie deleted")
}
