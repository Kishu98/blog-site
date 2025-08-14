package main

import (
	db "blog-site/internals/database"
	"blog-site/internals/handlers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Connecting to the database
	err := db.ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Creating a new HTTP server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	// Routes for this app
	http.HandleFunc("/blog", handlers.HandleBlogs)
	http.HandleFunc("/blog/", handlers.HandleBlog)
	http.HandleFunc("/login", handlers.HandleAuth)
	// http.HandleFunc("/signup", handlers.HandleSignup)

	// Channel to listen for interupt or terminate signals
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

	// Starting server in a goroutine
	// go func() {
	// 	log.Println("Server is running on port 8080...")
	// 	if err := srv.ListenAndServeTLS("/etc/letsencrypt/live/kishu-jain.com/fullchain.pem", "/etc/letsencrypt/live/kishu-jain.com/privkey.pem"); err != nil && err != http.ErrServerClosed {
	// 		log.Fatalf("ListenAndServe(): %v", err)
	// 	}
	// }()

	go func() {
		log.Println("Server is running on port 8080 (HTTP)...")
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
