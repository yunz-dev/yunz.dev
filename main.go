package main

import (
	"flag"
	"log"
	"net/http"
)


type Config struct {
  Addr string
}


func main() {
  router := http.NewServeMux()
  cfg := new(Config)
  flag.StringVar(&cfg.Addr, "addr", ":8000", "HTTP network address")
  flag.Parse()


  server := http.Server {
    Addr: cfg.Addr,
    Handler: router,
  }
  log.Printf("Starting server on port %s", cfg.Addr)
  server.ListenAndServe()

}
