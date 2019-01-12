package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func unauthorizedRouterHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"templates/index.html",
		"templates/unauthorized/header.html",
		"templates/unauthorized/main.html",
		"templates/footer.html",
	)

	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)
}

func authorizedRouterHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"templates/index.html",
		"templates/authorized/header.html",
		"templates/authorized/main.html",
		"templates/footer.html",
	)

	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)
}

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/", authorizedRouterHandler)
	http.ListenAndServe(":8080", nil)
}
