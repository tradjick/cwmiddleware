package limittomethods

import (
	"Middleware"
	"net/http"
)

func LimitToMethods(ms ...string) Middleware.Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			in := false
			for _, m := range ms {
				if m == r.Method || m == "*" {
					in = true
					break
				}
			}
			if !in {
				http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}
