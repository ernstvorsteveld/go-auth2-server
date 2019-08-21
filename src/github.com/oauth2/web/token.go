package web

import (
    "fmt"
    "net/http"
)

func v1Token(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
  
  for k, v := range r.Header {
    fmt.Fprintf(w, "Header field %q, Value %q\n", k, v)
  }

  r.ParseForm()
  
  kvs := r.Form
  
  for key, value := range kvs {
    fmt.Fprintf(w, "%s == %s\n", key, value)
  }
}

