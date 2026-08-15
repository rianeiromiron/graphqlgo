package main

import (
	"log"
	"net/http"

	"graphqlexample/internal/db"
	graphqlserver "graphqlexample/internal/graphql"
	"graphqlexample/internal/web"
)

func main() {
	cfg := db.Config{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "norimorienair4614",
		DBName:   "gestion_documentos",
		SSLMode:  "disable",
	}

	conn, err := db.New(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer conn.Close()

	if err := db.Migrate(conn); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", web.NewHandler())
	mux.Handle("/graphql", graphqlserver.NewHTTPHandler(conn))

	log.Println("GraphQL server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
