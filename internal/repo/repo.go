package repo

//go:generate mockgen -destination=../mocks/repo_mock.go -package=mocks github.com/ozoncp/ocp-roadmap-api/internal/repo Repo

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/ozoncp/ocp-roadmap-api/internal/entity"
	"github.com/rs/zerolog/log"
	"time"

	"github.com/Masterminds/squirrel"
)

type Repo interface {
	AddEntities(ctx context.Context, entities []entity.Roadmap) error
	CreateEntity(ctx context.Context, entity entity.Roadmap) error
	ListEntities(ctx context.Context, limit, offset uint64) ([]entity.Roadmap, error)
	DescribeEntity(ctx context.Context, entityId uint64) (*entity.Roadmap, error)
	RemoveEntity(ctx context.Context, entityId uint64) error
}
type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateEntity(ctx context.Context, entity entity.Roadmap) error {
	query := squirrel.
		Insert("roadmap").
		Columns("user_id", "link", "created_at").
		Values(entity.UserId, entity.Link, entity.CreatedAt).
		Suffix("RETURNING \"id\"").
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	err := query.QueryRowContext(ctx).Scan(&entity.Id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	return nil
}

func (r *Repository) AddEntities(ctx context.Context, entities []entity.Roadmap) error {
	query := squirrel.
		Insert("roadmap").
		Columns("user_id", "link", "created_at").
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	for _, v := range entities {
		query = query.Values(v.UserId, v.Link, v.CreatedAt)
	}

	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) ListEntities(ctx context.Context, limit, offset uint64) ([]entity.Roadmap, error) {
	var listData []entity.Roadmap
	query := squirrel.Select("id", "user_id", "link", "created_at").
		From("roadmap").
		RunWith(r.db).
		Limit(limit).
		Offset(offset).
		PlaceholderFormat(squirrel.Dollar)

	rows, err := query.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id, userId uint64
		var link string
		var createdAt time.Time

		if err := rows.Scan(&id, &userId, &link, &createdAt); err != nil {
			log.Error().Msg(err.Error())
			return listData, err
		}
		listData = append(listData, *entity.NewRoadMap(id, userId, link, createdAt))
	}

	return listData, nil
}

func (r *Repository) DescribeEntity(ctx context.Context, entityId uint64) (*entity.Roadmap, error) {
	var roadmap *entity.Roadmap
	var id, userId uint64
	var link string
	var createdAt time.Time

	query := squirrel.Select("id", "user_id", "link", "created_at").
		From("roadmap").
		Where(squirrel.Eq{"id": entityId}).
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	if err := query.QueryRowContext(ctx).Scan(&id, &userId, &link, &createdAt); err != nil {
		return nil, err
	}
	roadmap = entity.NewRoadMap(id, userId, link, createdAt)

	return roadmap, nil
}

func (r *Repository) RemoveEntity(ctx context.Context, entityId uint64) error {
	query := squirrel.Delete("roadmap").
		Where(squirrel.Eq{"id": entityId}).
		RunWith(r.db).
		PlaceholderFormat(squirrel.Dollar)

	res, err := query.ExecContext(ctx)
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}
	row, err := res.RowsAffected()
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}
	if row <= 0 {
		err := errors.New("no row found to delete")
		log.Error().Msg(err.Error())
		return err
	}

	return nil
}
