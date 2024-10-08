package main

import (
	"net/http"
	"strings"
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

func Home(w http.ResponseWriter, r *http.Request) {
	posts :=fetchBlogPost()
	if r.Header.Get("HX-Request") == "true" {
		section := strings.TrimPrefix(r.URL.Path, "/")
		if section == "" {
			section = "home"
		}

		switch section {
		case "home":
			templ.Home().Render(r.Context(), w)
		case "about":
			templ.About().Render(r.Context(), w)
		case "blog":
			templ.Blog(posts).Render(r.Context(), w)
		case "contact":
			templ.Contact().Render(r.Context(), w)
		default:
			http.NotFound(w, r)
		}
	} else {
		templ.Layout(posts).Render(r.Context(), w)
	}
}


// func Layout(s string) {
// 	panic("unimplemented")
// }

// func Blog(w http.ResponseWriter, r *http.Request) {
// 	posts := fetchBlogPost()
// 	templ.Blog(posts).Render(r.Context(), w)

// }

// func About(w http.ResponseWriter, r *http.Request) {
// 	templ.About().Render(context.Background(), w)
// }

// func Contact(w http.ResponseWriter, r *http.Request) {
// 	templ.Contact().Render(context.Background(), w)

// }