package middleware

import (
	dm "cruiseapp/database/middleware"
	fm "cruiseapp/repository/factory/middleware"
	"net/http"
)

type Middleware func(http.Handler) http.HandlerFunc

var WrapMiddleware = ChainMiddleware(
	dm.DbMiddleware,
	fm.PgRepoFactoryMiddleware,
)

func ChainMiddleware(middleware ...Middleware) Middleware {
	return func(next http.Handler) http.HandlerFunc {
		for i := len(middleware) - 1; i >= 0; i-- {
			next = middleware[i](next)
		}
		return next.ServeHTTP
	}
}
