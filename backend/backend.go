package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func run(addr string, c chan os.Signal) error {
	mux := http.NewServeMux()
	mux.Handle("/",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			clientAddr := r.Header.Get("X-Forwarded-For")
			log.Printf("%s -> %s -> %s", clientAddr, r.RemoteAddr, r.URL)
			_, _ = w.Write(index)
		}),
	)
	srv := &http.Server{
		Addr:              addr,
		Handler:           mux,
		IdleTimeout:       time.Minute,
		ReadHeaderTimeout: 30 * time.Second,
	}
	go func() {
		for {
			if <-c == os.Interrupt {
				_ = srv.Close()
				return
			}
		}
	}()
	fmt.Printf("Listening on %s ...\n", srv.Addr)
	err := srv.ListenAndServe()

	if err == http.ErrServerClosed {
		err = nil
	}
	return err
}

func main() {
	addr := flag.String("listen", "localhost:8080", "listen address")
	flag.Parse()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	err := run(*addr, c)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server stopped")
}

var index = []byte(`<!DOCTYPE html>
<html lang="en">
<head>
 <meta charset="UTF-8">
 <title>Caddy Backend Test</title>
</head>
<body>
 <h1>Hello from Caddy!</h1>
</body>
</html>`)
