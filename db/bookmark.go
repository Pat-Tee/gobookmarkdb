package db

import (
  "database/sql"
  "github.com/pat-tee/gobookmarkdb/models"
)

func (db Database) GetAllBookmarks() (*models.BookmarkList, error) {
  list := &models.BookmarkList{}

  rows, err := db.Conn.Query(
    "SELECT rowid, url, description, created_at, updated_at FROM bookmark")// ORDER BY rowid DESC")
  if err != nil {
    return list, err
  }

  for rows.Next() {
    var bookmark models.Bookmark
    err := rows.Scan(
      &bookmark.Rowid, 
      &bookmark.URL,
      &bookmark.Desc,
      &bookmark.CreatedAt, 
      &bookmark.UpdatedAt, )
    if err != nil {
      return list, err
    }
    list.Bookmark = append(list.Bookmark, bookmark)
  }
  return list, nil
}

func (db Database) AddBookmark(bookmark *models.Bookmark) error {
  var rowid int
  var createdAt string
  query := `INSERT INTO bookmark (url, description) VALUES ($1, $2) RETURNING rowid, created_at`
  err := db.Conn.QueryRow(query, bookmark.URL, bookmark.Desc).Scan(&rowid, &createdAt)
  if err != nil {
    return err
  }
  bookmark.Rowid = rowid
  bookmark.CreatedAt = createdAt
  return nil
}

func (db Database) DeleteBookmark(bookmarkId int) error {
  query := `DELETE FROM bookmark WHERE rowid = $1`
  _, err := db.Conn.Exec(query, bookmarkId)
  switch err {
  case sql.ErrNoRows:
    return ErrNoMatch
  default:
    return err
  }
}

func (db Database) GetBookmark(bookmarkId int) (*models.Bookmark, error) {
  bookmark := models.Bookmark{}
  query := `SELECT rowid, url, description, created_at, updated_at FROM bookmark WHERE rowid = $1`
  err := db.Conn.QueryRow(query, bookmarkId).Scan(
    &bookmark.Rowid, &bookmark.URL, &bookmark.Desc, &bookmark.CreatedAt, &bookmark.UpdatedAt)
  switch err {
  case sql.ErrNoRows:
    return nil, ErrNoMatch
  default:
    return &bookmark, err
  }
}

func (db Database) UpdateBookmark(bookmarkId int, bookmarkData models.Bookmark) (*models.Bookmark, error) {
  bookmark := models.Bookmark{}
  query := 
  `UPDATE bookmark SET url=$1, description=$2 
  WHERE rowid=$3 
  RETURNING rowid, url, description, created_at, updated_at`
  err:= db.Conn.QueryRow(query, bookmarkData.URL, bookmarkData.Desc, bookmarkId).Scan(
    &bookmark.Rowid, &bookmark.URL, &bookmark.Desc, &bookmark.CreatedAt, &bookmark.UpdatedAt)
  if err != nil {
    if err == sql.ErrNoRows {
      return &bookmark, ErrNoMatch
    }
    return &bookmark, err
  }
  return &bookmark, nil
}
