package cwmwadapters

import (
	"context"
	cwmiddlware "github.com/tradjick/cwmiddleware"
	"net/http"
)

type RequestIdGenerator func() ([]byte, error)

func EmbedRequestId(g RequestIdGenerator) cwmiddlware.Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqId, err := g()
			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			h.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "requestId", string(reqId))))
		})
	}
}
