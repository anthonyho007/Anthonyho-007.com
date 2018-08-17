package middleware

import (
	"log"
	"net/http"
)

func App(port string) {
	mux := Router()
	mux.HandlerFunc("GET", "/", HandleMainPage)
	mux.HandlerFunc("GET", "/about", HandleAboutPage)
	mux.HandlerFunc("GET", "/projects", HandleProjectsPage)
	mux.HandlerFunc("GET", "/experiences", HandleExperiencePage)
	mux.ServeFiles(
		"/assets/*filepath",
		http.Dir("assets/"),
	)
	mw := Middleware{}
	mw.Add(mux)

	log.Fatal(http.ListenAndServe(":"+port, mw))
}
