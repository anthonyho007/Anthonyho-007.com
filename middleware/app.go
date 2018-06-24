package middleware

import (
	"log"
	"net/http"
)

func App(port string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HandleMainPage)
	mux.Handle(
		"/assets/",
		http.StripPrefix("/assets/",
			http.FileServer(http.Dir("assets/")),
		),
	)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
