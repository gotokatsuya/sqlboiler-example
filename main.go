package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/gotokatsuya/sqlboiler-example/domain/entity"
	"github.com/gotokatsuya/sqlboiler-example/domain/repository"
	"github.com/gotokatsuya/sqlboiler-example/infra/datastore"
)

func main() {
	rand.Seed(time.Now().Unix())

	ctx := context.Background()

	db := datastore.MustNew()
	userRepo := repository.NewUser(db)

	user := entity.NewUser("gotokatsuya", time.Date(1992, 11, 10, 0, 0, 0, 0, time.Local), "M", "000000000")
	entity.UserEncryptPhoneNumber(user)

	if err := userRepo.Create(ctx, user); err != nil {
		panic(err)
	}
	users, err := userRepo.Find(ctx)
	if err != nil {
		panic(err)
	}
	for _, v := range users {
		fmt.Printf("%+v\n", *v)
	}
}
