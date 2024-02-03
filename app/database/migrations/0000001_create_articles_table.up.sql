CREATE TABLE IF NOT EXISTS "authors" (
    author_id serial PRIMARY KEY,
    author_name varchar(120),
    is_moderator boolean
);

CREATE TABLE IF NOT EXISTS "articles" (
    article_id serial PRIMARY KEY,
    header TEXT,
    content TEXT,
    author INT,
    FOREIGN KEY (author) REFERENCES authors(author_id)
); 

CREATE TABLE IF NOT EXISTS "comments" (
    comment_id serial PRIMARY KEY,
    comment TEXT,
    commentator INT UNIQUE,
    article INT UNIQUE,
    FOREIGN KEY (commentator) REFERENCES authors(author_id),
    FOREIGN KEY (article) REFERENCES articles(article_id) 
)