package util

import (
	"encoding/hex"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomHex(n int) string {
	bytes := make([]byte, n)
	if _, err := r.Read(bytes); err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}
