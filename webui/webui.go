package webui

import (
  "fmt"
  "html/template"
  "net/http"
  "github.com/go-chi/chi/v5"
)

func WebsiteRouter(router chi.Router) {
  router.Get("/", getRoot)
  router.Get("/favicon.ico", faviconHandler)
  router.Get("/index.js", javascriptHandler)

}

func getRoot(w http.ResponseWriter, r *http.Request) {

  tmpl, err := template.ParseFiles("webui/Public/index.html")
  if err != nil {
    fmt.Print(err)
    return
  }
  tmpl.Execute(w, "")
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w,r, "webui/Public/favicon.ico")
}

func javascriptHandler(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w,r, "webui/Public/index.js")
}
