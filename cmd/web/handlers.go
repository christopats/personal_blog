package main

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/christopats/go-blogv2/ui/templ"
)

func fetchBlogPost() []templ.BlogPost {
	return []templ.BlogPost{
		{
			Slug: "example-post",
			Title: "First Blog Post",
			Description: "This is the first blog post.",
			Date: time.Now().AddDate(0, 0, -5),
			Tags: []string{"Go", "HTMX", "Tailwind", "DaisyUI"},
			Content: "Hello, World! Check out this insightful blog post",
		},
		{
			Slug: "example-post",
			Title: "Second Blog Post",
			Description: "This is another blog post.",
			Date: time.Now().AddDate(0, 0, -5),
			Tags: []string{"Go", "HTMX", "Tailwind", "DaisyUI"},
			Content: "Hello, World! Check out this insightful blog post",
		},
		{
			Slug: "example-post",
			Title: "Third Blog Post",
			Description: "Are you tired of blog posts yet?",
			Date: time.Now().AddDate(0, 0, -5),
			Tags: []string{"Go", "HTMX", "Tailwind", "DaisyUI"},
			Content: "Hello, World! Check out this insightful blog post",
		},
	}
}

func fetchPostbySlug(slug string) *templ.BlogPost {
	posts := fetchBlogPost()

	for _, post := range posts {
		if post.Slug == slug {
			return &post
		}

	}
	return nil
}

func Index(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-type", "text/html")
	templ.Layout(templ.Home()).Render(context.Background(), w)
}

func Home(w http.ResponseWriter, r *http.Request) {

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

func BlogPost(w http.ResponseWriter, r *http.Request) {
	slug := strings.TrimPrefix(r.URL.Path, "/blog/")
	if slug == "" {
		http.NotFound(w, r)
		return
	}
	post := fetchPostbySlug(slug)
	if post == nil {
		http.NotFound(w, r)
		return
	}
	if r.Header.Get("HX-Request") == "true" {
		templ.BlogArticle(post).Render(r.Context(), w)
	} else {
		templ.Layout(templ.BlogArticle(post)).Render(r.Context(), w)
	}
} 

func Contact(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("HX-Request") == "true" {
		templ.Contact().Render(context.Background(), w)
	} else {
		templ.Layout(templ.Contact()).Render(context.Background(), w)
	}
}

