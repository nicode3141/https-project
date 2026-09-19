package server

import (
	"html/template"
	"net/http"
)

type Server struct {
	httpServer *http.Server
	template *template.Template
}

func New() *Server {
	templates := template.Must(template.ParseGlob("web/templates/*.html"))

	mux := http.NewServeMux();

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if(r.URL.Path != "/"){
			http.NotFound(w, r);
			return;
		}
	})

	return nil;
}