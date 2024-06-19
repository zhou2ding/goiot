package encrypt

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/gogf/gf/v2/crypto/gaes"
	"goiot/internal/pkg/logger"
)

var aesKey = []byte("Linkview2024")

func EncStringAES(in string) string {
	hash := sha256.Sum256(aesKey)
	encByte, err := gaes.Encrypt([]byte(in), hash[:])
	if err != nil {
		logger.Logger.Errorf("enc err %s", err)
	}
	return hex.EncodeToString(encByte)
}

func DecStringAES(in string) string {
	cipherText, err := hex.DecodeString(in)
	if err != nil {
		logger.Logger.Errorf("decode hex err %s", err)
	}
	hash := sha256.Sum256(aesKey)
	decByte, err := gaes.Decrypt(cipherText, hash[:])
	if err != nil {
		logger.Logger.Errorf("decode aes err %s", err)
	}
	return string(decByte)
}
