package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const DB_CONNECTION_CTX_KEY = "foobar"

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	SslMode  string
}

type DbHandler struct {
	Config DbConfig
}

func (dh *DbHandler) Open() *sql.DB {
	cfg := dh.Config
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DbName, cfg.SslMode)

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}

	return conn
}

func DbMiddleware(next http.Handler) http.Handler {
	cfg := DbConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		DbName:   "cruisedb",
		SslMode:  "disable",
	}

	dbHandler := DbHandler{
		Config: cfg,
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db := dbHandler.Open()
		defer db.Close()

		ctx := context.WithValue(r.Context(), DB_CONNECTION_CTX_KEY, db)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetDb(r *http.Request) *sql.DB {

	return r.Context().Value(DB_CONNECTION_CTX_KEY).(*sql.DB)
}
