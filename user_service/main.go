package main

import "github.com/gofiber/fiber/v2"

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{
	{ID: 1, Name: "СонПон", Age: 24},
	{ID: 2, Name: "Данек", Age: 21},
	{ID: 3, Name: "Гевин", Age: 24},
	{ID: 4, Name: "Степик", Age: 21},
	{ID: 5, Name: "Максим", Age: 24},
}

func main() {
	app := fiber.New()

	app.Get("/users", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(users)
	})

	app.Listen(":8080")
}
