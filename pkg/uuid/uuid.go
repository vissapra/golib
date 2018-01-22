package uuid

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

//Returns a random string that confirms to UUID format
// 8-4-4-4-12 for a total of 36 characters (32 alphanumeric characters and four hyphens)
// Ex: 32c1b9d3-3436-5927-5528-a03b9cfdd24e
func UUID() (string, error) {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	hx := hex.EncodeToString(bytes)
	return fmt.Sprintf("%s-%s-%s-%s-%s", hx[0:8], hx[8:12], hx[12:16], hx[16:20], hx[20:]), nil
}
