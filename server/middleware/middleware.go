package middleware

import (
	"context"
	"cruiseapp/database"
	fct "cruiseapp/repository/factory"
	"net/http"
)

type Middleware func(http.Handler) http.HandlerFunc

var WrapMiddleware = ChainMiddleware(
	DbMiddleware,
	PgRepoFactoryMiddleware,
)

func ChainMiddleware(middleware ...Middleware) Middleware {
	return func(next http.Handler) http.HandlerFunc {
		for i := len(middleware) - 1; i >= 0; i-- {
			next = middleware[i](next)
		}
		return next.ServeHTTP
	}
}

func DbMiddleware(next http.Handler) http.HandlerFunc {
	cfg := database.DbConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		DbName:   "cruisedb",
		SslMode:  "disable",
	}

	dbHandler := database.DbHandler{
		Config: cfg,
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db := dbHandler.Open()
		defer db.Close()

		ctx := context.WithValue(r.Context(), database.DB_CONNECTION_CTX_KEY, db)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func PgRepoFactoryMiddleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db := database.GetDb(r)
		var factory fct.RepoFactory = fct.PgRepoFactory{Conn: db}

		ctx := context.WithValue(r.Context(), fct.MIDDLEWARE_CTX_KEY, factory)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
