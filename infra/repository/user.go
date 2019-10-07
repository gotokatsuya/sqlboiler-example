package repository

import (
	"context"

	"github.com/volatiletech/sqlboiler/boil"

	"github.com/gotokatsuya/sqlboiler-example/domain/sqlboiler"
	"github.com/gotokatsuya/sqlboiler-example/infra/datastore"
)

type User struct {
	cli *datastore.Client
}

func NewUser(cli *datastore.Client) *User {
	return &User{cli}
}

func (u *User) Create(ctx context.Context, ent *sqlboiler.User) error {
	return ent.Insert(ctx, u.cli, boil.Infer())
}

func (u *User) Find(ctx context.Context) (sqlboiler.UserSlice, error) {
	list, err := sqlboiler.Users().All(ctx, u.cli)
	if err != nil {
		return nil, err
	}
	return list, nil
}
