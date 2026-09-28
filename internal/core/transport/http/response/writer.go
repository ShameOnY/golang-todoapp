package core_http_response

import "net/http"

const (
	StatusCodeUninitialization = -1
)

type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: w,
		statusCode:     StatusCodeUninitialization,
	}
}

func (rw *ResponseWriter) WriteHeader(statusCode int) {
	if rw.statusCode != StatusCodeUninitialization {
		return
	}

	rw.ResponseWriter.WriteHeader(statusCode)
	rw.statusCode = statusCode
}

func (rw *ResponseWriter) Write(b []byte) (int, error) {
	if rw.statusCode == StatusCodeUninitialization {
		rw.statusCode = http.StatusOK
	}

	return rw.ResponseWriter.Write(b)
}

func (rw *ResponseWriter) GetStatusCodeOrPanic() int {
	if rw.statusCode == StatusCodeUninitialization {
		panic("no status code set")
	}

	return rw.statusCode
}
