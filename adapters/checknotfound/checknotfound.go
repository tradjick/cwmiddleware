package checknotfound

import (
	"Middleware"
	"net/http"
)

func CheckNotFound(p string) Middleware.Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != p {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}
