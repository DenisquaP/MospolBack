package postgres

import (
	"context"
	"fmt"
	"mospol/internal/entity"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
)

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

func (p PostgresDB) WriteAuthor(author entity.CreateAuthorRequest) error {
	query := "INSERT INTO authors (author_name, is_moderator) VALUES (@name, @moder)"

	args := pgx.NamedArgs{
		"name":  author.Name,
		"moder": author.IsModerator,
	}

	if _, err := p.client.Exec(p.ctx, query, args); err != nil {
		return err
	}

	return nil
}

func (p PostgresDB) WriteAtricle(article entity.CreateAtricleRequest) error {
	query := "INSERT INTO articles (header, content, author) VALUES (@header, @content, @author)"

	args := pgx.NamedArgs{
		"header":  article.Title,
		"content": article.Content,
		"author":  article.Author,
	}

	if _, err := p.client.Exec(p.ctx, query, args); err != nil {
		return err
	}

	return nil
}

func (p PostgresDB) Close() error {
	err := p.client.Close(p.ctx)
	if err != nil {
		return err
	}

	return nil
}

func (p *PostgresDB) MigrationsUp() error {
	sourceURL := "file://database/migrations/up"

	m, err := migrate.New(sourceURL, p.url)
	if err != nil {
		return err
	}

	if err = m.Up(); err != nil {
		return err
	}

	return nil
}
