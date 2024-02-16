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

  apiPort := os.Getenv("API_PORT")
  htmlPort := os.Getenv("HTML_PORT")

  apiListener, err := net.Listen("tcp", apiPort)
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

  httpHandler := handler.NewDBHandler(database)
  apiServer := &http.Server{
    Handler: httpHandler,
  }

  htmlListener, err := net.Listen("tcp", ":8080")
  if err != nil {
    log.Fatalf("Error occurred: %s", err.Error())
  }
  htmlHandler := handler.NewWebuiHandler() 
  htmlServer := &http.Server{
    Handler: htmlHandler,
  }

  go func() {
    apiServer.Serve(apiListener)
  }()
  go func() {
    htmlServer.Serve(htmlListener)
  }()  
  
  defer Stop(apiServer)
  defer Stop(htmlServer)
  
  log.Printf("Started API server on %s", apiPort)
  log.Printf("Started HTML server on %s", htmlPort)
  ch := make(chan os.Signal, 1)
  signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
  log.Println(fmt.Sprint(<-ch))
  log.Println("Stopping API server.")
  log.Println("Stopping HTML server.")
}

func Stop(server *http.Server) {
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()
  if err := server.Shutdown(ctx); err != nil {
    log.Printf("Could not shut down server correctly: %v\n", err)
    os.Exit(1)
  }
}
