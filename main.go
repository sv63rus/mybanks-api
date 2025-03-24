package main

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"mybanks-api/ent"
	"mybanks-api/ent/ogent"
	gen "mybanks-api/ent/ogent"
	"net/http"
)

func main() {

	dsn := "postgres://postgres:@localhost:5432/postgres?sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer db.Close()

	client := ent.NewClient(ent.Driver(db))

	// Миграции (если нужно)
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	h := &handler{
		OgentHandler: ogent.NewOgentHandler(client),
		client:       client,
	}

	srv, err := gen.NewServer(h)
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", withCORS(srv)))

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK")
	})

	log.Println("🚀 Server started on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("server error: %v", err)
	}
}

func withCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Разрешаем запросы с локального фронта
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// OPTIONS-запросы — это preflight, отвечаем сразу
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Передаём управление дальше
		h.ServeHTTP(w, r)
	})
}
