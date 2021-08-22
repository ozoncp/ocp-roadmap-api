package repo

//go:generate mockgen -destination=../mocks/repo_mock.go -package=mocks github.com/ozoncp/ocp-roadmap-api/internal/repo Repo

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/ozoncp/ocp-roadmap-api/internal/entity"
	"github.com/rs/zerolog/log"
	"time"
)

type Repo interface {
	AddEntities(ctx context.Context, entities []entity.Roadmap) error
	ListEntities(ctx context.Context, limit, offset uint64) ([]entity.Roadmap, error)
	DescribeEntity(ctx context.Context, entityId uint64) (*entity.Roadmap, error)
	RemoveEntity(ctx context.Context, entityId uint64) error
}
type Repository struct {
	connection *pgx.Conn
}

func NewRepository(connection *pgx.Conn) Repo {
	return &Repository{
		connection: connection,
	}
}

func (r Repository) AddEntities(ctx context.Context, entities []entity.Roadmap) error {
	command := "INSERT INTO roadmap(user_id, link, created_at) VALUES($1, $2, $3)"
	batch := &pgx.Batch{}

	for _, v := range entities {
		batch.Queue(command, v.UserId, v.Link, time.Now())
	}

	br := r.connection.SendBatch(ctx, batch)
	if _, err := br.Exec(); err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	return nil
}

func (r Repository) ListEntities(ctx context.Context, limit, offset uint64) ([]entity.Roadmap, error) {
	var listData []entity.Roadmap
	command := "SELECT id, user_id, link FROM roadmap LIMIT $1 OFFSET $2"
	rows, err := r.connection.Query(ctx, command, limit, offset)
	if err != nil {
		log.Fatal().Msg(err.Error())
		return listData, err
	}
	defer rows.Close()

	for rows.Next() {
		var id, userId uint64
		var link string

		if err := rows.Scan(&id, &userId, &link); err != nil {
			log.Error().Msg(err.Error())
			return listData, err
		}
		listData = append(listData, *entity.NewRoadMap(id, userId, link))
	}

	return listData, nil
}

func (r Repository) DescribeEntity(ctx context.Context, entityId uint64) (*entity.Roadmap, error) {
	var roadmap *entity.Roadmap
	var id, userId uint64
	var link string

	command := "SELECT id, user_id, link FROM roadmap WHERE id=$1"
	row := r.connection.QueryRow(ctx, command, entityId)

	if err := row.Scan(&id, &userId, &link); err != nil {
		log.Error().Msg(err.Error())
		return roadmap, err
	}
	roadmap = entity.NewRoadMap(id, userId, link)

	return roadmap, nil
}

func (r Repository) RemoveEntity(ctx context.Context, entityId uint64) error {
	command := "DELETE FROM roadmap WHERE id=$1"
	commandTag, err := r.connection.Exec(ctx, command, entityId)
	if err != nil {
		log.Error().Msg(err.Error())
		return err
	}

	if commandTag.RowsAffected() != 1 {
		err := errors.New("no row found to delete")
		log.Error().Msg(err.Error())
		return err
	}

	return nil
}
