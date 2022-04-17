package cwmwadapters

import (
	cwmiddlware "github.com/tradjick/cwmiddleware"
	"net/http"
)

type HttpLog interface {
	Request(r *http.Request)
	Response(r *http.Request, l *LogResponseWriter)
}

func LogRequestAdapter(hl HttpLog) cwmiddlware.Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			hl.Request(r)
			m := makeLogResponseWriter(w)
			defer hl.Response(r, m)
			h.ServeHTTP(m, r)
		})
	}
}

type LogResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func makeLogResponseWriter(w http.ResponseWriter) *LogResponseWriter {
	return &LogResponseWriter{w, http.StatusOK}
}

func (m *LogResponseWriter) WriteHeader(code int) {
	m.statusCode = code
	m.ResponseWriter.WriteHeader(code)
}
