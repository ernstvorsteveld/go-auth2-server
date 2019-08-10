package oauth2

import (
    "net/http"
    "log"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Println(r.URL.RequestURI(), r.URL.Query().Encode())
        f(w, r)
    }
}

func InitServer() {
    http.HandleFunc("/v1/authorize", logging(v1Authorize))
    http.HandleFunc("/v1/token", logging(v1Token))

    fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.ListenAndServe(":8080", nil)
}