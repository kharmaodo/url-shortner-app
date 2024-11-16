package url

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func Shorten(originalUrl string) (string, string) {
	h := sha256.New()
	h.Write([]byte(originalUrl))
	fmt.Println(h.Sum(nil))
	hash := hex.EncodeToString(h.Sum(nil))
	shortUrl := hash[:8]
	return hash, shortUrl
}
