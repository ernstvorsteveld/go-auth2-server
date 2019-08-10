package oauth2

import (
    "fmt"
    "net/http"
)

func v1Authorize(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "So you want to authorize huh?\n")
}
