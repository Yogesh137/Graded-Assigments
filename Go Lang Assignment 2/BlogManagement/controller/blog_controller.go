package controller

import (
	"crud/model"
	"crud/service"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type BlogController struct {
	BlogService *service.BlogService
}

func NewBlogController(blogService *service.BlogService) *BlogController {
	return &BlogController{BlogService: blogService}
}

func (controller *BlogController) CreateBlog(w http.ResponseWriter, r *http.Request) {
	var blog model.Blog
	if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdBlog, err := controller.BlogService.CreateBlog(&blog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdBlog)
}

func (controller *BlogController) GetBlog(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/blog/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	blog, err := controller.BlogService.GetBlog(id)
	if err != nil {
		http.Error(w, "Blog not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blog)
}

func (controller *BlogController) GetAllBlogs(w http.ResponseWriter, r *http.Request) {
	blogs, err := controller.BlogService.GetAllBlogs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blogs)
}

func (controller *BlogController) UpdateBlog(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/blog/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var blog model.Blog
	if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	blog.ID = id
	updatedBlog, err := controller.BlogService.UpdateBlog(&blog)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedBlog)
}

func (controller *BlogController) DeleteBlog(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/blog/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = controller.BlogService.DeleteBlog(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Blog deleted successfully"}`))
}
