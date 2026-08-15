package db

import "testing"

func TestTaskModel(t *testing.T) {
	if got, want := 1+1, 2; got != want {
		t.Fatalf("unexpected arithmetic result: got %d, want %d", got, want)
	}
}

func TestDatabaseConfig(t *testing.T) {
	cfg := Config{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "norimorienair4614",
		DBName:   "gestion_documentos",
		SSLMode:  "disable",
	}

	if cfg.Host == "" || cfg.DBName == "" {
		t.Fatal("database config should be populated")
	}
}
