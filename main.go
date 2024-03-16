package main

// 2 routes, GET and POST

import (
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Coverage struct {
	Coverage float64 `json:"coverage"`
}

func generateAuthKey() string {
	// generate a random key
	var newKey = ""
	key := uuid.New().String()

	// remove the hyphens
	newKey = strings.ReplaceAll(key, "-", "")
	// return the key
	return newKey
}

func main() {

	var coverage = new(Coverage)
	var key = generateAuthKey()

	log.Default().Println("Auth Key: ", key)

	app := fiber.New()

	app.Get("/badge", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("Coverage: %f", coverage.Coverage))
	})

	app.Post("/badge", func(c *fiber.Ctx) error {

		// check if the key is valid
		if c.Get("Authorization") != key {
			return c.SendStatus(401)
		}

		if err := c.BodyParser(coverage); err != nil {
			return err
		}

		return c.SendString("Coverage updated")
	})

	log.Fatal(app.Listen(":3000"))

}
