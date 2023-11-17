package webui

import (
  "fmt"
  "html/template"
  "net/http"
  "github.com/go-chi/chi/v5"
)

func Website(router chi.Router) {
  router.Get("/", getRoot)
}

func getRoot(w http.ResponseWriter, r *http.Request) {

  tmpl, err := template.ParseFiles("webui/Public/index.html")
  if err != nil {
    fmt.Print(err)
    return
  }
  tmpl.Execute(w, "")
}
