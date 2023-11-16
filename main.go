package main

import (
  "context"
  "fmt"
  "github.com/pat-tee/gobookmarkdb/db"
  "github.com/pat-tee/gobookmarkdb/handler"
  "log"
  "net"
  "net/http"
  "os"
  "os/signal"
  "syscall"
  "time"
  "github.com/joho/godotenv"
)

func main() {

  err := godotenv.Load()
  if err != nil {
    log.Fatal("Could not load env credentials.")
  }

  port := os.Getenv("PORT")

  listener, err := net.Listen("tcp", port)
  if err != nil {
    log.Fatalf("Error occurred: %s", err.Error())
  }
  
  dbFilename := os.Getenv("DBFILENAME")
  if dbFilename == "" {
    log.Fatal("Couldn't get database filename from env")
  }
  
  log.Printf("Using file: %s", dbFilename)
  database, err := db.Initialize(dbFilename)
  if err != nil {
    log.Fatalf("Database failed initialization: %v", err)
  }
 
  defer database.Conn.Close()

  httpHandler := handler.NewHandler(database)
  server := &http.Server{
    Handler: httpHandler,
  }

  go func() {
    server.Serve(listener)
  }()
  defer Stop(server)
  log.Printf("Started server on %s", port)
  ch := make(chan os.Signal, 1)
  signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
  log.Println(fmt.Sprint(<-ch))
  log.Println("Stopping API server.")
}

func Stop(server *http.Server) {
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()
  if err := server.Shutdown(ctx); err != nil {
    log.Printf("Could not shut down server correctly: %v\n", err)
    os.Exit(1)
  }
}

