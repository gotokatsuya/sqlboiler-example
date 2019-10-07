package entity

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gotokatsuya/sqlboiler-example/domain/sqlboiler"
)

func NewUser(
	username string,
	birthday time.Time,
	gender,
	phoneNumber string,
) *sqlboiler.User {
	return &sqlboiler.User{
		Username:    username,
		Birthday:    birthday,
		Gender:      gender,
		PhoneNumber: phoneNumber,
	}
}

func UserEncryptPhoneNumber(u *sqlboiler.User) {
	u.PhoneNumber = fmt.Sprintf("%d", rand.Int())
}
