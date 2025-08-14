package handlers

import (
	db "blog-site/internals/database"
	"blog-site/internals/middleware"
	"blog-site/internals/models"
	"encoding/json"
	"net/http"
	"strconv"
)

// Main handler function
func HandleBlogs(w http.ResponseWriter, r *http.Request) {
	checkEnableCORS(w, r)
	switch r.Method {
	case http.MethodPost:
		middleware.JWTMiddleware(http.HandlerFunc(createBlogHandler)).ServeHTTP(w, r)
	case http.MethodGet:
		getAllBlogsHandler(w)
	}
}

func HandleBlog(w http.ResponseWriter, r *http.Request) {
	checkEnableCORS(w, r)

	id, err := getBlogIDFromURL(r)
	if err != nil {
		http.Error(w, "Invalid blog id", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodPut:
		middleware.JWTMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			updateBlogHandler(w, r, id)
		})).ServeHTTP(w, r)
	case http.MethodGet:
		getBlogHandler(w, id)
	case http.MethodDelete:
		middleware.JWTMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			deleteBlogHandler(w, id)
		})).ServeHTTP(w, r)
	}
}

// handler function to create a blog post
func createBlogHandler(w http.ResponseWriter, r *http.Request) {
	var blog models.Blog

	if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	blog, err := models.CreateBlog(db.GetDB(), blog)
	if err != nil {
		http.Error(w, "Error creating blog post. Please try again.", http.StatusInternalServerError)
		return
	}

	// Sending back the new Blog as re sponse
	jsonResponse(w, http.StatusCreated, blog)

}

// handler function to get all blogs
func getAllBlogsHandler(w http.ResponseWriter) {
	blogs, err := models.GetAllBlogs(db.GetDB())
	if err != nil {
		http.Error(w, "Error retrieving blog posts. Please try again later.", http.StatusNotFound)
		return
	}

	jsonResponse(w, http.StatusOK, blogs)
}

// handler function to update blog
func updateBlogHandler(w http.ResponseWriter, r *http.Request, id int) {
	var blog models.Blog

	if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	blog, err := models.UpdateBlog(db.GetDB(), id, blog)
	if err != nil {
		http.Error(w, "Error updating blog post. Please try again later.", http.StatusNotFound)
		return
	}

	jsonResponse(w, http.StatusOK, blog)
}

// handler function to get blog by id
func getBlogHandler(w http.ResponseWriter, id int) {

	blog, err := models.GetBlog(db.GetDB(), id)
	if err != nil {
		http.Error(w, "Error retrieving blog. Please try again later.", http.StatusNotFound)
		return
	}

	jsonResponse(w, http.StatusOK, blog)
}

func deleteBlogHandler(w http.ResponseWriter, id int) {

	blog, err := models.GetBlog(db.GetDB(), id)
	if err != nil {
		http.Error(w, "Error retrieving blog to delete. Please try again later.", http.StatusNotFound)
		return
	}

	err = models.DeleteBlog(db.GetDB(), id)
	if err != nil {
		http.Error(w, "Error deleting blog. Please try again later.", http.StatusInternalServerError)
		return
	}

	jsonResponse(w, http.StatusOK, blog)
}

func jsonResponse(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func checkEnableCORS(w http.ResponseWriter, r *http.Request) {
	// react_url := os.Getenv("REACT_URL")
	react_url := "http://localhost:8080"
	// react_url := "https://kishu-jain.com"
	w.Header().Set("Access-Control-Allow-Origin", react_url)
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}
}

func getBlogIDFromURL(r *http.Request) (int, error) {
	return strconv.Atoi(r.URL.Path[len("/blog/"):])
}
