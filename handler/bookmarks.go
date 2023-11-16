package handler

import (
  "context"
  "fmt"
  "net/http"
  "strconv"
  "github.com/go-chi/chi/v5"
  "github.com/go-chi/render"
  "github.com/pat-tee/gobookmarkdb/db"
  "github.com/pat-tee/gobookmarkdb/models"
)
var bookmarkIdKey = "bookmarkId"

func bookmarks(router chi.Router) {
  router.Get("/", getAllBookmarks)
  router.Post("/", createBookmark)
  router.Route("/{bookmarkId}", func (router chi.Router) {
    router.Use(BookmarkContext)
    router.Delete("/", deleteBookmark)
    router.Get("/", getBookmark)
    router.Put("/", updateBookmark)
  })
}

func BookmarkContext(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    bookmarkId := chi.URLParam(r, "bookmarkId")
    if bookmarkId == "" {
      render.Render(w,r,ErrorRenderer(fmt.Errorf("bookmark id is required")))
      return
    }
    id, err := strconv.Atoi(bookmarkId)
    if err != nil {
      render.Render(w,r,ErrorRenderer(fmt.Errorf("invalid bookmark id")))
    }
    ctx := context.WithValue(r.Context(), bookmarkIdKey, id)
    next.ServeHTTP(w,r.WithContext(ctx))
  })
}

func createBookmark(w http.ResponseWriter, r *http.Request) {
  bookmark := &models.Bookmark{}
  if err := render.Bind(r, bookmark); err != nil {
    render.Render(w,r,ErrBadRequest)
    return
  }
  if err := dbInstance.AddBookmark(bookmark); err != nil {
    render.Render(w,r,ErrorRenderer(err))
    return
  }
  if err := render.Render(w,r,bookmark); err != nil {
    render.Render(w,r,ServerErrorRenderer(err))
    return
  }
}

func getAllBookmarks(w http.ResponseWriter, r *http.Request) {
  bookmarks, err := dbInstance.GetAllBookmarks()
  if err != nil {
    render.Render(w,r,ServerErrorRenderer(err))
    return
  }
  if err := render.Render(w,r,bookmarks); err != nil {
    render.Render(w,r,ErrorRenderer(err))
  }
}

func getBookmark(w http.ResponseWriter, r*http.Request) {
  bookmarkId := r.Context().Value(bookmarkIdKey).(int)
  bookmark, err := dbInstance.GetBookmark(bookmarkId)
  if err != nil {
    if err == db.ErrNoMatch {
        render.Render(w,r,ErrNotFound)
      } else {
        render.Render(w,r,ErrorRenderer(err))
      }
      return
    }
  if err := render.Render(w,r,bookmark); err != nil {
    render.Render(w,r,ServerErrorRenderer(err))
    return
  }
}

func deleteBookmark(w http.ResponseWriter, r *http.Request) {
  bookmarkId := r.Context().Value(bookmarkIdKey).(int)
  err := dbInstance.DeleteBookmark(bookmarkId)
  if err != nil {
    if err == db.ErrNoMatch {
      render.Render(w,r,ErrNotFound)
    } else {
      render.Render(w,r,ServerErrorRenderer(err))
    }
    return
  }
}

func updateBookmark(w http.ResponseWriter, r *http.Request) {

  bookmarkId := r.Context().Value(bookmarkIdKey).(int)
  bookmarkData := models.Bookmark{}
  if err := render.Bind(r, &bookmarkData); err != nil {
    render.Render(w,r,ErrBadRequest)
    return
  }

  bookmark, err := dbInstance.UpdateBookmark(bookmarkId, bookmarkData)
  if err != nil {
    if err == db.ErrNoMatch {
      render.Render(w,r,ErrNotFound)
    } else {
      render.Render(w,r,ServerErrorRenderer(err))
    }
    return
  }
  if err := render.Render(w, r, bookmark); err != nil {
    render.Render(w,r,ServerErrorRenderer(err))
    return
  }
}
