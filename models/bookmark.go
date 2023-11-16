package models

import (
  "fmt"
  "net/http"
)

type Bookmark struct {
  Rowid int `json:"rowid"`
  URL string `json:"url"`
  Desc string `json:"description"`
  CreatedAt string `json:"created_at"`
  UpdatedAt string `json:"updated_at"`
}

type BookmarkList struct {
  Bookmark []Bookmark `json:"bookmarks"`
}

func (b *Bookmark) Bind(r *http.Request) error {
  if b.URL == "" {
    return fmt.Errorf("URL is required")
  }
  return nil
}

func (*BookmarkList) Render(w http.ResponseWriter, r *http.Request) error {
  return nil
}

func (*Bookmark) Render(w http.ResponseWriter, r *http.Request) error {
  return nil
}


