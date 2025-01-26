package main

import (
	"flag"
	"log"
	"net/http"
)


type Config struct {
  Addr string
}


func home(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, "static/index.html")
}

func main() {
  cfg := new(Config)
  flag.StringVar(&cfg.Addr, "addr", ":8000", "HTTP network address")
  flag.Parse()

  router := http.NewServeMux()
  router.HandleFunc("/", home)

  fileServer := http.FileServer(http.Dir("./static"))
  router.Handle("/static/", http.StripPrefix("/static", fileServer))



  server := http.Server {
    Addr: cfg.Addr,
    Handler: router,
  }
  log.Printf("Starting server on port %s", cfg.Addr)
  server.ListenAndServe()

}
