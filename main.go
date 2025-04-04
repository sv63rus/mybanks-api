package main

import (
	"bytes"
	"context"
	"entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"
	"github.com/vektah/gqlparser/v2/formatter"
	"log"
	"mybanks-api/ent"
	"mybanks-api/graph"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8080"

// corsMiddleware устанавливает CORS-заголовки для разрешения запросов с http://localhost:3000
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Установите нужный origin или "*" если хотите разрешить все домены
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		// Обрабатываем preflight-запросы
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {

	dsn := "postgres://postgres:@localhost:5432/postgres?sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed connecting to postgres: %v", err)
	}
	defer db.Close()

	client := ent.NewClient(ent.Driver(db))
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	schema := graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{client}})

	srv := handler.New(schema)

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	// Добавляем отдачу схемы по HTTP:
	var buf bytes.Buffer
	f := formatter.NewFormatter(&buf)
	f.FormatSchema(schema.Schema())

	http.HandleFunc("/schema.graphql", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write(buf.Bytes())
	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	// Оборачиваем DefaultServeMux в corsMiddleware
	log.Fatal(http.ListenAndServe(":"+port, corsMiddleware(http.DefaultServeMux)))
}
