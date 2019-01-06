package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func homeRouterHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		"templates/index.html",
		"templates/header.html",
		"templates/main.html",
		"templates/footer.html",
	)

	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)
}

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/", homeRouterHandler)
	http.ListenAndServe(":9000", nil)
}
