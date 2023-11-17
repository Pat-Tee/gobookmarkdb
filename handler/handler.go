package handler
import (
  "net/http"
  "github.com/go-chi/chi/v5"
  "github.com/go-chi/render"
//  "github.com/go-chi/cors"
  "github.com/pat-tee/gobookmarkdb/db"
  "github.com/pat-tee/gobookmarkdb/webui"
)

func NewWebuiHandler() http.Handler {
  router := chi.NewRouter()

  router.NotFound(notFoundHandler)
  router.Route("/", webui.Website)

  return router
}

var dbInstance db.Database
func NewDBHandler(db db.Database) http.Handler {
  router := chi.NewRouter()
/*  router.Use(cors.Handler(cors.Options{
    AllowedOrigins: []string{"https://*", "http://*", "file://*", "null"},
    AllowedMethods: []string{"GET","POST","PUT","DELETE","OPTIONS"},
    AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CRSF-Token"},
    ExposedHeaders: []string{"Link"},
    AllowCredentials: false,
    MaxAge: 300,
  }))
*/
  dbInstance = db
  router.MethodNotAllowed(methodNotAllowedHandler)
  router.NotFound(notFoundHandler)
  router.Route("/bookmarks", bookmarks)
  return router
}

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application.json")
  w.WriteHeader(405)
  render.Render(w,r,ErrMethodNotAllowed)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(400)
  render.Render(w,r,ErrNotFound)
}
