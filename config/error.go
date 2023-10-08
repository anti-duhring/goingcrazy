package config

import "errors"

var ErrNicknameAlreadyExists = errors.New("nickname already exists")
var ErrRedisGetPerson = errors.New("failed to get person from redis")
