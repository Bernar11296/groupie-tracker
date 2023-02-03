package internal

import (
	"html/template"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		Errors(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		Errors(w, http.StatusMethodNotAllowed)
		return
	}
	tmp, err := template.ParseFiles("./ui/html/index.html")
	if err != nil {
		Errors(w, http.StatusInternalServerError)
		return
	}
	// var test []Artists

	test := TakeCards()

	err = tmp.Execute(w, test)
	if err != nil {
		Errors(w, http.StatusInternalServerError)
		return
	}
}
