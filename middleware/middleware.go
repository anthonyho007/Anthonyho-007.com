package middleware

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
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
		fmt.Println("trying handler", handler)
		handler.ServeHTTP(mWriter, req)
		fmt.Println("Writen? ", mWriter.written)
		if mWriter.written {
			return
		}
	}
	fmt.Println("Page not found, redirect to 404")
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

func (writer *MiddlewareResponseWriter) WriteHeader(num int) {
	writer.written = true
	writer.ResponseWriter.WriteHeader(num)
}

func NewMiddlewareResponseWriter(writer http.ResponseWriter) *MiddlewareResponseWriter {
	rw := MiddlewareResponseWriter{ResponseWriter: writer}
	return &rw
}

func Router() *httprouter.Router {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(http.ResponseWriter, *http.Request) {})
	return router
}
