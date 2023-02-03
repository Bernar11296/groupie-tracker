package internal

import (
	"html/template"
	"net/http"
	"strconv"
)

func result(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		Errors(w, http.StatusMethodNotAllowed)
		return
	}
	tmp, err := template.ParseFiles("./ui/html/artist.html")
	if err != nil {
		Errors(w, http.StatusInternalServerError)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 || id > 52 {
		Errors(w, http.StatusNotFound)
		return
	}

	test := TakeArtist(id)
	concert := TakeConcert(id)
	if err != nil {
		Errors(w, http.StatusInternalServerError)
		return
	}
	test.Relations = concert
	err = tmp.Execute(w, test)
	if err != nil {
		Errors(w, http.StatusInternalServerError)
		return
	}
}
