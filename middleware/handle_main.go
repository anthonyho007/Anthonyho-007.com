package middleware

import (
	"net/http"
)

func HandleMainPage(writer http.ResponseWriter, req *http.Request) {
	RenderTemplate(writer, req, "index/home", nil)
}

func HandleAboutPage(writer http.ResponseWriter, req *http.Request) {
	RenderTemplate(writer, req, "index/about", nil)
}
