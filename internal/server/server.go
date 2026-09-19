package server

import (
	"html/template"
	"net/http"
)

type Server struct {
	httpServer *http.Server
	templates *template.Template
}

func New() *Server {
	templates := template.Must(template.ParseGlob("web/templates/*.html"))

	mux := http.NewServeMux();

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if(r.URL.Path != "/"){
			http.NotFound(w, r);
			return;
		}

		props := struct {
			Title string
			Name string
		}{
			Title: "Hi",
			Name: "Hi",
		}

		err := templates.ExecuteTemplate(w, "index.html", props)

		if(err != nil){
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})

	static := http.FileServer(http.Dir("web/static"))

	mux.Handle("/static/", http.StripPrefix("/static/", static))


	return &Server{
		templates: templates,
		httpServer: &http.Server{
			Addr: ":8080",
			Handler: mux,
		},
	};
}

func (s *Server) ListenAndServe() error {
	return s.httpServer.ListenAndServe()
}