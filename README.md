GoBookmarkDB is a sqlite-based Go language-based bookmark server and API
for uploading, viewing, and managing users' browser bookmarks.
It is meant to be run as a standalone executable with a separate sdlite.db file
that may be moved, copied or manipulated independently of this project.
Otherwise, it runs a backend CRUD API to interact with sqlite in combination
with a simple HTML server for the frontend in order to be browser-agnostic and
run from any appropriate device on one's network.
This is my first project in Go and is being written in order to learn the language.
