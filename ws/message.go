package ws

import (
	"fmt"
	"net/http"
	"strings"
)

func sendMessage(r *http.Request, message string) {
	hub, ok := r.Context().Value(WS_HUB_CTX_KEY).(*Hub)
	if !ok {
		return
	}
	hub.broadcast <- message
}

func sendCrudMsg(r *http.Request, action string, model string, id int64) {
	sendMessage(r, fmt.Sprintf("%s %s %d", action, strings.ToLower(model), id))
}

func SendCreatedMsg(r *http.Request, model string, id int64) {
	sendCrudMsg(r, "created", model, id)
}

func SendUpdatedMsg(r *http.Request, model string, id int64) {
	sendCrudMsg(r, "updated", model, id)
}

func SendDeletedMsg(r *http.Request, model string, id int64) {
	sendCrudMsg(r, "deleted", model, id)
}
