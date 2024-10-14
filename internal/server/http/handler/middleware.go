package handler

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func (ir *inventoryRoutes) userIdentity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("authorization")

		if header == "" {
			log.Printf("InventoryRoutes - userIdentity: empty auth header")
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			log.Printf("InventoryRoutes - userIdentity: invalid auth header")
			return
		}

		userId, err := ir.is.ParseToken(headerParts[1])

		if err != nil {
			log.Printf("InventoryRoutes - userIdentity:%v", err)
			return
		}

		if err != nil {
			http.Error(w, fmt.Sprintf("JSON encode error: %v", err), http.StatusInternalServerError)
			return
		}

		if userId == 0 {
			log.Printf("InventoryRoutes - userIdentity: invalid userid")
			return
		}

		next.ServeHTTP(w, r)
	})
}
