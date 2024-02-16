package postgres

import (
	"context"
	"errors"
	"fmt"
	"mospol/internal/entity"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
)

var ErrAuthorDoesNotExists error = errors.New("this author doesn`t exists")
var ErrArticleDoesNotExists error = errors.New("this article doesn`t exists")

type Article struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type Comment struct {
	Commentator string `json:"commentator"`
	Comment     string `json:"comment"`
}

type ArticleStr struct {
	Article  Article   `json:"article"`
	Comments []Comment `json:"comments"`
}

type PostgresDB struct {
	Config
	url    string
	ctx    context.Context
	client *pgx.Conn
}

func NewPostgres() (PostgresDB, error) {
	config := NewConfig()
	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", config.User, config.Password, config.Host, config.Port, config.DbName)

	return PostgresDB{
		Config: config,
		url:    url,
		ctx:    context.Background(),
	}, nil
}

func (p *PostgresDB) Connection() error {
	db, err := pgx.Connect(p.ctx, p.url)
	if err != nil {
		return err
	}

	p.client = db
	return nil
}

func (p PostgresDB) CheckAuthor(author_id int) error {
	var author int
	err := p.client.QueryRow(p.ctx, "SELECT author_id FROM authors WHERE author_id=$1", author_id).Scan(&author)
	if err != nil {
		return err
	}

	return nil
}

func (p PostgresDB) CheckArticle(article_id int) error {
	var article int

	err := p.client.QueryRow(p.ctx, "SELECT article_id FROM articles WHERE article_id=$1", article_id).Scan(&article)
	if err != nil {
		return err
	}

	return nil
}

func (p PostgresDB) WriteAuthor(author entity.CreateAuthorRequest) error {
	query := "INSERT INTO authors (email, author_name, password, is_moderator) VALUES (@email, @name, @password, @moder)"

	args := pgx.NamedArgs{
		"email":    author.Email,
		"name":     author.Name,
		"password": author.Password,
		"moder":    author.IsModerator,
	}

	if _, err := p.client.Exec(p.ctx, query, args); err != nil {
		return err
	}

	return nil
}

func (p PostgresDB) WriteAtricle(article entity.CreateAtricleRequest) error {
	if err := p.CheckAuthor(article.Author); err != nil {
		return ErrAuthorDoesNotExists
	}

	query := "INSERT INTO articles (title, content, author) VALUES (@title, @content, @author)"

	args := pgx.NamedArgs{
		"title":   article.Title,
		"content": article.Content,
		"author":  article.Author,
	}

	if _, err := p.client.Exec(p.ctx, query, args); err != nil {
		return err
	}

	return nil
}

func (p PostgresDB) WriteComment(comment entity.CreateCommentRequest) error {
	if err := p.CheckAuthor(comment.Commentator); err != nil {
		return ErrArticleDoesNotExists
	}

	if err := p.CheckArticle(comment.Article); err != nil {
		return ErrAuthorDoesNotExists
	}

	query := "INSERT INTO comments (comment, commentator, article) VALUES (@comment, @commentator, @article)"

	args := pgx.NamedArgs{
		"article":     comment.Article,
		"commentator": comment.Commentator,
		"comment":     comment.Comment,
	}

	if _, err := p.client.Exec(p.ctx, query, args); err != nil {
		return err
	}

	return nil
}

func (p PostgresDB) ReadArticles(page int) (articles []Article, err error) {
	var art Article
	offset := page * 7
	if page == 1 {
		offset = 0
	}
	query := fmt.Sprintf("SELECT header, content, author_name FROM articles JOIN authors on articles.author = authors.author_id LIMIT %d OFFSET %d", 7, offset)

	rows, err := p.client.Query(p.ctx, query)
	if err != nil {
		return
	}

	for rows.Next() {
		err = rows.Scan(&art.Title, &art.Content, &art.Author)
		if err != nil {
			return
		}

		articles = append(articles, art)
	}

	return
}

func (p *PostgresDB) LastPage() (lp int, err error) {
	query := "SELECT COUNT(header) FROM articles"

	err = p.client.QueryRow(p.ctx, query).Scan(&lp)
	if err != nil {
		return
	}

	if lp < 7 {
		lp = 1
		return
	}

	lp = lp / 7

	return
}

func (p PostgresDB) ReadArticle(article_id int) (article ArticleStr, err error) {
	query := fmt.Sprintf("SELECT header, content, author_name FROM articles JOIN authors ON articles.author = authors.author_id  WHERE article_id = %d", article_id)
	rows, err := p.client.Query(p.ctx, query)
	if err != nil {
		return
	}

	for rows.Next() {
		err = rows.Scan(&article.Article.Title, &article.Article.Content, &article.Article.Author)
		if err != nil {
			return
		}
	}

	comments, err := p.ReadComments(article_id)
	if err != nil {
		return
	}

	article.Comments = comments

	return
}

func (p PostgresDB) ReadComments(article int) (comments []Comment, err error) {
	var comment Comment
	query := fmt.Sprintf("SELECT comment, author_name FROM comments JOIN authors on comments.commentator = authors.author_id WHERE article = %d", article)
	rows, err := p.client.Query(p.ctx, query)
	if err != nil {
		return
	}

	for rows.Next() {
		err = rows.Scan(&comment.Comment, &comment.Commentator)
		if err != nil {
			return
		}

		comments = append(comments, comment)
	}

	return
}

func (p PostgresDB) Close() error {
	err := p.client.Close(p.ctx)
	if err != nil {
		return err
	}

	return nil
}

func (p *PostgresDB) MigrationsUp() error {
	sourceURL := "file://database/migrations"

	m, err := migrate.New(sourceURL, p.url)
	if err != nil {
		return err
	}

	if err = m.Up(); err != nil {
		return err
	}

	return nil
}
