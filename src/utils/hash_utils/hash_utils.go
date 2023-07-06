package hash_utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/google/uuid"
)

func GenerateTransactionId() string {
	id := uuid.New()
	return id.String()
}

func TimeStamp() string {
	now := time.Now()
	return now.Format("20060102150405")
}

func EncryptSHA256(val []byte) string {
	hash := sha256.Sum256(val)
	return hex.EncodeToString(hash[:])
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
