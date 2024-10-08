package main

import (
	"context"
	"net/http"
	"time"

	"github.com/christopats/go-blogv2/ui/templ"
)

func fetchBlogPost() []templ.BlogPost {
	return []templ.BlogPost{
		{
			Title: "First Blog Post",
			Description: "This is the first blog post.",
			Date: time.Now().AddDate(0, 0, -5),
			Tags: []string{"Go", "HTMX", "Tailwind", "DaisyUI"},
		},
		{
			Title: "Second Blog Post",
			Description: "This is another blog post.",
			Date: time.Now().AddDate(0, 0, -5),
			Tags: []string{"Go", "HTMX", "Tailwind", "DaisyUI"},
		},
		{
			Title: "Third Blog Post",
			Description: "Are you tired of blog posts yet?",
			Date: time.Now().AddDate(0, 0, -5),
			Tags: []string{"Go", "HTMX", "Tailwind", "DaisyUI"},
		},
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	// posts :=fetchBlogPost()

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-type", "text/html")
	templ.Layout(templ.Home()).Render(context.Background(), w)
}

func Home(w http.ResponseWriter, r *http.Request) {
	// posts :=fetchBlogPost()

	if r.URL.Path != "/home" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-type", "text/html")

	if r.Header.Get("HX-Request") == "true" {
		templ.Home().Render(context.Background(), w)
	} else {
		templ.Layout(templ.Home()).Render(context.Background(), w)
	}

	
}

func About(w http.ResponseWriter, r *http.Request) {
	// posts :=fetchBlogPost()

	if r.URL.Path != "/about" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-type", "text/html")
	
	if r.Header.Get("HX-Request") == "true" {
		templ.About().Render(context.Background(), w)
	} else {
		templ.Layout(templ.About()).Render(context.Background(), w)
	}
}

func Blog(w http.ResponseWriter, r *http.Request) {
	posts := fetchBlogPost()

	if r.Header.Get("HX-Request") == "true" {
		templ.Blog(posts).Render(r.Context(), w)
	} else {
		templ.Layout(templ.Blog(posts)).Render(r.Context(), w)
	}
}

func Contact(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("HX-Request") == "true" {
		templ.Contact().Render(context.Background(), w)
	} else {
		templ.Layout(templ.Contact()).Render(context.Background(), w)
	}
}

