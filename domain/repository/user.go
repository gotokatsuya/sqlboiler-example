package repository

import (
	"context"

	"github.com/gotokatsuya/sqlboiler-example/domain/sqlboiler"
	"github.com/gotokatsuya/sqlboiler-example/infra/datastore"
	"github.com/gotokatsuya/sqlboiler-example/infra/repository"
)

type User interface {
	Create(ctx context.Context, ent *sqlboiler.User) error
	Find(ctx context.Context) (sqlboiler.UserSlice, error)
}

func NewUser(cli *datastore.Client) User {
	return repository.NewUser(cli)
}
