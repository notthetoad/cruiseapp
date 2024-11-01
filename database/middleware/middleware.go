package middleware

import (
	"context"
	"cruiseapp/database"
	"net/http"
)

func DbMiddleware(next http.Handler) http.HandlerFunc {
	cfg := database.LoadConfig()

	var dbHandler database.Databaser

	dbHandler = &database.PgHandler{
		Config: cfg,
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db := dbHandler.Open()
		defer db.Close()

		ctx := context.WithValue(r.Context(), database.DB_CONNECTION_CTX_KEY, db)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}