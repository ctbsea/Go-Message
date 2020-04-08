package util

import (
	"crypto/md5"
	"fmt"
)

const SALT = "Hello-word"

func GetPass(pass string) (saltPass string) {
	data := []byte(pass + SALT)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}