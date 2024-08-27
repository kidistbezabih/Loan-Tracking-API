package repositories

import "errors"

var ErrNoUesrWitThisId = errors.New("no user with this id")
var ErrFailToCreateUser = errors.New("fail to create user")
var ErrNoUesrWitThisEmail = errors.New("no user with this id")
var ErrNoUesrWitThisUsername = errors.New("no user with this username")
