package main

import (
	"net/http"
	"text/template"
)

func loginHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		tmpl, err := template.New("login.html").ParseFiles("resources/pages/login.html")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Content-Type", "text/html")
		err = tmpl.ExecuteTemplate(writer, "login.html", "")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
func main() {
	mux := http.NewServeMux()

	mux.Handle("/login", loginHandler())
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("resources/images"))))
	http.ListenAndServe("localhost:55443", mux)
}
