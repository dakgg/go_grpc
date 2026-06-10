package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"os"
)

// HashPublicKey hashes a public key with HMAC-SHA256 using a pepper from env.
func HashPublicKey(publicKey string) string {
	pepper := os.Getenv("HASH_PEPPER")
	mac := hmac.New(sha256.New, []byte(pepper))
	mac.Write([]byte(publicKey))
	return hex.EncodeToString(mac.Sum(nil))
}
