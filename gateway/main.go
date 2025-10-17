package main

import (
	"log"

	"github.com/carlmjohnson/requests"
	"github.com/gofiber/fiber/v2"
)

type HealthStatus struct {
	Status string `json:"status"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type CombinedResponse struct {
	Health HealthStatus `json:"health"`
	Users  []User       `json:"users"`
}

func main() {
	app := fiber.New()

	app.Get("/aggregate", func(c *fiber.Ctx) error {
		var healthResp HealthStatus
		var usersResp []User

		err := requests.
			URL("http://health-service:8080/health").
			ToJSON(&healthResp).
			Fetch(c.Context())
		if err != nil {
			log.Println("Error fetching health:", err)
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": "health-service unavailable"})
		}

		err = requests.
			URL("http://user-service:8080/users").
			ToJSON(&usersResp).
			Fetch(c.Context())
		if err != nil {
			log.Println("Error fetching users:", err)
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": "user-service unavailable"})
		}

		response := CombinedResponse{
			Health: healthResp,
			Users:  usersResp,
		}

		return c.JSON(response)
	})

	log.Println("API Gateway running on :8080")
	app.Listen(":8080")
}
