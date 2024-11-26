package main

import (
	"context"
	"fmt"
	"net/http"
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
		{
			Slug: "example-post",
			Title: "Third Blog Post",
			Description: "Are you tired of blog posts yet?",
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
		{
			Slug: "example-post",
			Title: "Third Blog Post",
			Description: "Are you tired of blog posts yet?",
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

func fetchProjects() []templ.ProjectPost {
	return []templ.ProjectPost{
		{
			Slug: "personal-site",
			Title: "Personal Site",
			Description: "My personal portfolio site project using Golang, Templ and HTMX",
			ImagePath: "gopher.png",
			Content: "I built this site using go and htmx. inspired by all the trendy youtubers, here are the steps i took",
		},
		{
			Slug: "personal-site",
			Title: "Personal Site",
			Description: "My personal portfolio site project using Golang, Templ and HTMX",
			ImagePath: "gopher.png",
			Content: "I built this site using go and htmx. inspired by all the trendy youtubers, here are the steps i took",
		},
	} 
}

func fetchProjectBySlug(slug string) *templ.ProjectPost {
	projects := fetchProjects()

	for _, project := range projects {
		if project.Slug == slug {
			return &project
		}
	}
	return nil
}

func (app *application) Index(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-type", "text/html")
	templ.Layout(templ.Home()).Render(context.Background(), w)
}

func (app *application) createBlogHandler (w http.ResponseWriter, r *http.Request) {
	// RETURN PLAIN TEXT FOR TIME BEING
	fmt.Fprintln(w, "create a new blog")
}

func (app *application) showBlogHandler (w http.ResponseWriter, r *http.Request) {
	
	// USING READIDPARAM() HELPER METHOD, READ ID FROM REQUEST CONTEXT
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// INTERPOLATE ID INTO RESPONSE
	fmt.Fprintf(w, "show blog %d\n", id)
}

// func (app *application) Home(w http.ResponseWriter, r *http.Request) {

// 	if r.URL.Path != "/home" {
// 		http.NotFound(w, r)
// 		return
// 	}
// 	w.Header().Set("Content-type", "text/html")

// 	if r.Header.Get("HX-Request") == "true" {
// 		templ.Home().Render(context.Background(), w)
// 	} else {
// 		templ.Layout(templ.Home()).Render(context.Background(), w)
// 	}

	
// }

// func (app *application) About(w http.ResponseWriter, r *http.Request) {

// 	if r.URL.Path != "/about" {
// 		http.NotFound(w, r)
// 		return
// 	}

// 	w.Header().Set("Content-type", "text/html")
	
// 	if r.Header.Get("HX-Request") == "true" {
// 		templ.About().Render(context.Background(), w)
// 	} else {
// 		templ.Layout(templ.About()).Render(context.Background(), w)
// 	}
// }

// func (app *application) Blog(w http.ResponseWriter, r *http.Request) {
// 	posts := fetchBlogPost()

// 	if r.Header.Get("HX-Request") == "true" {
// 		templ.Blog(posts).Render(r.Context(), w)
// 	} else {
// 		templ.Layout(templ.Blog(posts)).Render(r.Context(), w)
// 	}
// }

// func (app *application) BlogPost(w http.ResponseWriter, r *http.Request) {
// 	slug := strings.TrimPrefix(r.URL.Path, "/blog/")
// 	if slug == "" {
// 		http.NotFound(w, r)
// 		return
// 	}
// 	post := fetchPostbySlug(slug)
// 	if post == nil {
// 		http.NotFound(w, r)
// 		return
// 	}
// 	if r.Header.Get("HX-Request") == "true" {
// 		templ.BlogArticle(post).Render(r.Context(), w)
// 	} else {
// 		templ.Layout(templ.BlogArticle(post)).Render(r.Context(), w)
// 	}
// } 

// func (app *application) Projects(w http.ResponseWriter, r *http.Request) {
// 	projects := fetchProjects()

// 	if r.Header.Get("HX-Request") == "true" {
// 		templ.Projects(projects).Render(r.Context(), w)
// 	} else {
// 		templ.Layout(templ.Projects(projects)).Render(r.Context(), w)
// 	}
// }

// func (app *application) ProjectPostPage(w http.ResponseWriter, r *http.Request) {
// 	slug := strings.TrimPrefix(r.URL.Path, "/projects/")
// 	if slug == "" {
// 		http.NotFound(w, r)
// 		return
// 	}

// 	project := fetchProjectBySlug(slug)
// 	if project == nil {
// 		http.NotFound(w, r)
// 		return
// 	}
	

// 	if r.Header.Get("HX-Request") == "true" {
// 		templ.ProjectPage(project).Render(r.Context(), w)
// 	} else {
// 		templ.Layout(templ.ProjectPage(project)).Render(r.Context(), w)
// 	}
// }


// func (app *application) Contact(w http.ResponseWriter, r *http.Request) {
// 	if r.Header.Get("HX-Request") == "true" {
// 		templ.Contact().Render(context.Background(), w)
// 	} else {
// 		templ.Layout(templ.Contact()).Render(context.Background(), w)
// 	}
// }

