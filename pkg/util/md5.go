package util

import (
	"crypto/md5"
	"encoding/hex"
	"go_server/pkg/setting"
	"log"
)

func PasswordMd5(str string) string {
	h := md5.New()
	log.Println(setting.PasswordPreHalt)
	h.Write([]byte(str + setting.PasswordPreHalt))
	return hex.EncodeToString(h.Sum(nil))
}
