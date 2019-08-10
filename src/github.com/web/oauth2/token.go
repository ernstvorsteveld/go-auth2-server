package oauth2

import (
    "fmt"
    "net/http"
)

func v1Token(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "So you want to token huh?\n")
}

