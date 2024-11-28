package data

import (
	"context"
	"time"

	"github.com/christopats/go-blogv2/validator"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lib/pq"
)

type BlogPost struct {
	ID 			int64 
	CreatedAt 	time.Time
	Title       string 
	Description string 
	Content     string 
	Tags        []string 
	Slug        string 
}

type BlogModel struct {
	DB *pgxpool.Pool
}

func ValidateBlog(v *validator.Validator, blog *BlogPost) {

	v.Check(blog.Title != "", "title", "must be provided")
	v.Check(len(blog.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(blog.Tags != nil, "genres", "must be provided")
	v.Check(len(blog.Tags) >= 1, "genres", "must contain at least 1 genre") 
	v.Check(len(blog.Tags) <= 5, "genres", "must not contain more than 5 genres") 
	v.Check(validator.Unique(blog.Tags), "genres", "must not contain duplicate values")
	}



func (b BlogModel) Insert(blogPost *BlogPost) error {
	// DEFINE THE QUERY FOR INSERTING NEW BLOG
	query := `
		INSERT INTO blogPost (title, description, content, tags, slug)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at`

	// DECALRE A SLICE CONTAINING THE VALUES FOR THE PLACE HOLDERS
	args := []interface{}{blogPost.Title, blogPost.Description, blogPost.Content, pq.Array(blogPost.Tags), blogPost.Slug}

	var ctx context.Context
	return b.DB.QueryRow(ctx, query, args...).Scan(&blogPost.ID, &blogPost.CreatedAt)

	
}

func (b BlogModel) Get(id int64) (*BlogPost, error) {
	return nil, nil
}

func (b BlogModel) Update(blogPost *BlogPost) error {
	return nil
}

func (b BlogModel) Delete(id int64) error {
	return nil
}
