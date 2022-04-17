package embedrequestid

import (
	"Middleware"
	"context"
	"log"
	"net/http"
)

type RequestIdGenerator func() ([]byte, error)

func EmbedRequestId(g RequestIdGenerator, l *log.Logger) Middleware.Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqId, err := g()
			if err != nil {
				l.Printf("error creating http request id : %s", err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			h.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "requestId", string(reqId))))
		})
	}
}
