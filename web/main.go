package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	srv := &http.Server{
		Addr:    ":8081",
		Handler: nil,
	}

	fs := http.FileServer(http.Dir("./dist"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := http.Dir("./dist").Open(r.URL.Path); err != nil {
			http.ServeFile(w, r, "./dist/index.html")
		} else {
			fs.ServeHTTP(w, r)
		}
	})

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// // Redirect to HTTPS
	// go func() {
	// 	log.Println("Redirecting HTTP to HTTPS...")
	// 	if err := http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		http.Redirect(w, r, "https://"+r.Host+r.URL.String(), http.StatusMovedPermanently)
	// 	})); err != nil && err != http.ErrServerClosed {
	// 		log.Fatalf("HTTP server failed: %v", err)
	// 	}
	// }()

	// // Starting server in a goroutine
	// go func() {
	// 	log.Println("Server is running on port 443 (HTTPS)...")
	// 	if err := srv.ListenAndServeTLS("/etc/letsencrypt/live/kishu-jain.com/fullchain.pem", "/etc/letsencrypt/live/kishu-jain.com/privkey.pem"); err != nil && err != http.ErrServerClosed {
	// 		log.Fatalf("ListenAndServeTLS(): %v", err)
	// 	}
	// }()
	go func() {
		log.Println("Frontend is running on port 8081 (HTTP)...")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	<-stop

	log.Println("Shutting down server...")

	// for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed: %v", err)
	}

	log.Println("Server exited properly")
}
