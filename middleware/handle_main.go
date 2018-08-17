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

func HandleProjectsPage(writer http.ResponseWriter, req *http.Request) {
	RenderTemplate(writer, req, "index/projects", nil)
}

func HandleExperiencePage(writer http.ResponseWriter, req *http.Request) {
	RenderTemplate(writer, req, "index/experiences", nil)
}
