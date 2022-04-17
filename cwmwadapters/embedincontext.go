package cwmwadapters

import (
	"context"
	cwmiddlware "github.com/tradjick/cwmiddleware"
	"net/http"
)

func EmbedInContext(k string, v interface{}) cwmiddlware.Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(context.WithValue(r.Context(), k, v))
			h.ServeHTTP(w, r)
		})
	}
}
