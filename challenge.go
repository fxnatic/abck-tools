package abcktools

import (
	"crypto/sha256"
	"encoding/hex"
)

func MN_S(byteArr []byte) []byte {
	h := sha256.New()
	h.Write(byteArr)
	return h.Sum(nil)
}

func ATS(byteArr []byte) string {
	return hex.EncodeToString(byteArr)
}
