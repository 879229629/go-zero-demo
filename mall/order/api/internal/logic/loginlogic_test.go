package logic

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	s := "key"

	h := sha256.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)

	fmt.Printf("origin: %s, %d sha256 hash: %x \n", s, len(bs), bs)
}

func TestHmac(t *testing.T) {
	key := "kuteng"
	data := "key"
	hmac := hmac.New(md5.New, []byte(key))
	hmac.Write([]byte(data))
	bs := hmac.Sum(nil)
	fmt.Printf("origin: %s, %d sha256 hash: %x \n", data, len(bs), bs)

}
