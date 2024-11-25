package middleware

import (
	"context"
	"cruiseapp/database"
	fct "cruiseapp/repository/factory"
	"net/http"
)

func PgRepoFactoryMiddleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db := database.GetDb(r)
		var factory fct.RepoFactory = fct.PgRepoFactory{Conn: db}

		ctx := context.WithValue(r.Context(), fct.MiddlewareCtxKey, factory)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
