package util

import (
	"crypto/md5"
	"fmt"
	"time"
)

const SALT = "Hello-word"

func GetPass(pass string) (saltPass string) {
	data := []byte(pass + SALT)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}

func GetToken(userId uint) (token string) {
	data := []byte(string(userId) + SALT + string(time.Now().Unix()))
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}