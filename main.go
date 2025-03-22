package main

import (
	"context"
	"entgo.io/ent/dialect"
	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"mybanks-api/ent"
	"os"
)

func main() {
	// Инициализация контекста
	ctx := context.Background()

	// Подключение к БД (можешь заменить на Postgres, MySQL и т.д.)
	client, err := ent.Open(dialect.SQLite, "file:myapp.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("❌ Failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	// Автоматическая миграция схемы
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("❌ Failed creating schema resources: %v", err)
	}

	// Запуск Fiber-приложения
	app := fiber.New()

	// Middleware
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Content-Type", "application/json")
		return c.Next()
	})

	// Базовый healthcheck
	app.Get("/api/v1/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	// Подключение роутов (позже)
	// handler.RegisterBankRoutes(app, client)

	// Запуск
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("🚀 Server listening on http://localhost:%s", port)
	log.Fatal(app.Listen(":" + port))
}
