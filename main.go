package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"graphqlexample/internal/db"
	graphqlserver "graphqlexample/internal/graphql"
	"graphqlexample/internal/web"
)

func main() {
	cfg := db.Config{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnvInt("DB_PORT", 5432),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", ""),
		DBName:   getEnv("DB_NAME", "gestion_documentos"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}

	conn, err := db.New(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer conn.Close()

	if getEnvBool("RUN_MIGRATIONS", true) {
		if err := db.Migrate(conn); err != nil {
			log.Fatalf("failed to migrate database: %v", err)
		}
	} else {
		log.Println("skipping database migrations (RUN_MIGRATIONS=false)")
	}

	mux := http.NewServeMux()
	mux.Handle("/", web.NewHandler())
	mux.Handle("/graphql", graphqlserver.NewHTTPHandler(conn))

	log.Printf("GraphQL server running at http://0.0.0.0:%s", getEnv("APP_PORT", "8080"))
	log.Fatal(http.ListenAndServe(":"+getEnv("APP_PORT", "8080"), mux))
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		parsed, err := strconv.Atoi(value)
		if err == nil {
			return parsed
		}
	}
	return fallback
}

func getEnvBool(key string, fallback bool) bool {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		parsed, err := strconv.ParseBool(value)
		if err == nil {
			return parsed
		}
	}
	return fallback
}
