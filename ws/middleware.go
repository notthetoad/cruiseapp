package ws

import (
	"context"
	"cruiseapp/server/middleware"
	"net/http"
)

const WsHubCtxKey = "WS_HUB_CTX_KEY"

func WsHubMiddleware(hub *Hub) middleware.Middleware {
	return func(next http.Handler) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), WsHubCtxKey, hub)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
