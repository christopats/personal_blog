CREATE TABLE blogs (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    title text NOT NULL,
    description text NOT NULL,
    content text NOT NULL, 
    tags text[] NOT NULL, 
    slug text NOT NULL
);