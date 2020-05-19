package tool

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
	"hash"
)

var h hash.Hash
var p hash.Hash

func BuildSha1(password string) string {
	if h == nil {
		h = sha1.New()
	}
	h.Write([]byte(password))
	defer h.Reset()
	return hex.EncodeToString(h.Sum(nil))
}

func BuildMd5(password string) string {
	if p == nil {
		p = md5.New()
	}
	p.Write([]byte(password))
	defer p.Reset()
	return hex.EncodeToString(p.Sum(nil))
}

func PasswordHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

func PasswordVerify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
