package database

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func GetConfig() DbConfig {
	return DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
		SslMode:  os.Getenv("DB_SSL_MODE"),
	}
}

const DB_CONNECTION_CTX_KEY = "DB_CONN_CTX_KEY"

type Databaser interface {
	Open() *sql.DB
}

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	SslMode  string
}

type DbConfiger interface {
	LoadConfig() DbConfig
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
	err = conn.Ping()
	if err != nil {
		panic(err)
	}

	return conn
}

type FooDbHandler struct {
	Cfg DbConfig
}

func (f *FooDbHandler) Open() *sql.DB {
	cfg := f.LoadConfig()
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DbName, cfg.SslMode)
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println(err)
	}
	err = conn.Ping()
	if err != nil {
		panic(err)
	}
	return conn
}

func (f *FooDbHandler) LoadConfig() DbConfig {
	fmt.Println("Loading config")
	return DbConfig{}
}

func GetDb(r *http.Request) *sql.DB {
	return r.Context().Value(DB_CONNECTION_CTX_KEY).(*sql.DB)
}
