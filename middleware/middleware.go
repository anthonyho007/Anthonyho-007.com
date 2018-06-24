package middleware

import (
	"net/http"
)

type Middleware []http.Handler

func (mw *Middleware) Add(handler http.Handler) {
	*mw = append(*mw, handler)
}

func (mw Middleware) ServeHTTP(
	writer http.ResponseWriter,
	req *http.Request,
) {
	mWriter := NewMiddlewareResponseWriter(writer)

	for _, handler := range mw {
		handler.ServeHTTP(mWriter, req)
		if mWriter.written {
			return
		}
	}
	Handle404(mWriter, req)
}

type MiddlewareResponseWriter struct {
	http.ResponseWriter
	written bool
}

func (writer *MiddlewareResponseWriter) Write(bytes []byte) (int, error) {
	writer.written = true
	return writer.ResponseWriter.Write(bytes)
}

func (writer * MiddlewareResponseWriter) WriteHeader(num int) {
	writer.written = true
	writer.ResponseWriter.WriteHeader(num)
}

func NewMiddlewareResponseWriter(writer http.ResponseWriter) *MiddlewareResponseWriter {
	rw := MiddlewareResponseWriter{ResponseWriter: writer}
	return &rw
}
