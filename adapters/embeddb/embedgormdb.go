package embeddb

import (
	"Middleware"
	"context"
	"gorm.io/gorm"
	"net/http"
)

func EmbedGormDb(db *gorm.DB) Middleware.Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(context.WithValue(r.Context(), "db", db))
			h.ServeHTTP(w, r)
		})
	}
}
