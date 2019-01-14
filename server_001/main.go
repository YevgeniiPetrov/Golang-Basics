package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
)

type (
	Bio struct {
		Name string
		Age  uint
	}

	User struct {
		Login    string
		Password string
		Bio
	}

	Cache map[string]User
)

var cacheByLogin Cache

func init() {
	adminUser := User{
		Login:    "admin",
		Password: "admin",
		Bio: Bio{
			Name: "admin",
			Age:  100,
		},
	}

	cacheByLogin = Cache{
		adminUser.Login: adminUser,
	}
}

func loginCheck(login, password string) (err error) {
	if user, ok := cacheByLogin[login]; ok {
		if user.Password == password {
			return
		}
	}

	return errors.New("invalid pass or login")
}

func register(login, password string) (err error) {
	if _, ok := cacheByLogin[login]; ok {
		return errors.New("user exist")
	}

	cacheByLogin[login] = User{
		Login:    login,
		Password: password,
	}

	return
}

func mainRouterHandler(w http.ResponseWriter, r *http.Request) {
	authorized := "unauthorized"

	t, err := template.ParseFiles(
		"templates/index.html",
		"templates/"+authorized+"/header.html",
		"templates/"+authorized+"/main.html",
		"templates/footer.html",
	)

	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	t.ExecuteTemplate(w, "index", nil)
}

func loginRouterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var login, pass string
	if _, ok := r.Form["login"]; ok {
		login = r.Form["login"][0]
	}
	if _, ok := r.Form["pass"]; ok {
		pass = r.Form["pass"][0]

	}
	err := loginCheck(login, pass)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	http.Redirect(w, r, "/?login="+login+"&pass="+pass, 303)
}

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/", mainRouterHandler)
	http.HandleFunc("/login", loginRouterHandler)
	//http.HandleFunc("/register", registerRouterHandler)
	http.ListenAndServe(":8080", nil)
}
