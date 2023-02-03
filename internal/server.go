package internal

import (
	"log"
	"net/http"
)

func Server() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/artist/", result)
	mux.HandleFunc("/search", search)
	fileServer := http.FileServer(http.Dir("./ui/style"))
	mux.Handle("/style", http.NotFoundHandler())
	mux.Handle("/style/", http.StripPrefix("/style", fileServer))
	log.Println("Запуск веб-сервера на http://127.0.0.1:8000")
	http.ListenAndServe(":8000", mux)
}
