package webui

import (
  "fmt"
  "html/template"
  "net/http"
  "io"
  "os"
  "github.com/go-chi/chi/v5"
)

func WebsiteRouter(router chi.Router) {
  router.Get("/", getRoot)
  router.Get("/favicon.ico", faviconHandler)
  router.Get("/index.js", javascriptHandler)
  router.Post("/upload", uploadHandler)
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

func uploadHandler(w http.ResponseWriter, r *http.Request) {

  r.ParseMultipartForm(10 << 20)

  file, handler, err := r.FormFile("form-files")
  if err != nil && handler != nil {
    fmt.Printf("Uploaded file: %+v\n", handler.Filename)
    fmt.Printf("MIME header: %v\n", handler.Header)
  }

  tempFile, err := os.CreateTemp("temp-html", "upload-*.html")
  if err != nil {
    fmt.Println(err)
  }

  defer tempFile.Close()

  fileBytes, err := io.ReadAll(file)

  if err != nil {
    fmt.Println(err)
  }

  tempFile.Write(fileBytes)

  fmt.Fprintf(w, "Successfully uploaded files\n")

}
