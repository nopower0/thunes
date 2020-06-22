package tools

import (
	"crypto/md5"
	"fmt"
)

const (
	HashSalt = "k@wlxR6V4C"
)

func PasswordHash(password string) string {
	saltedPassword := fmt.Sprintf("%s%s", password, HashSalt)
	hash := md5.Sum([]byte(saltedPassword))
	return fmt.Sprintf("%x", hash)
}

