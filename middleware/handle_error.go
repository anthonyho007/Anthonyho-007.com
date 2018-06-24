package middleware

import (
	"net/http"
)

func Handle404(writer http.ResponseWriter, req *http.Request) {
	RenderTemplate(
		writer,
		req,
		"errors/404",
		nil,
	)
}

func Handle500(writer http.ResponseWriter, req *http.Request, err error) {
	writer.WriteHeader(http.StatusInternalServerError)
	RenderTemplate(
		writer,
		req,
		"errors/500",
		map[string]interface{}{
			"Error": err.Error(),
		})
}
