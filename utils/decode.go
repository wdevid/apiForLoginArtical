package utils
import (
	"encoding/base64"
	"crypto/md5"
	"io"
	"fmt"
)

const (
	base64Table = "123QRSTUabcdVWXYZHijKLAWDCABDstEFGuvwxyzGHIJklmnopqr234560178912"
)

var coder = base64.NewEncoding(base64Table)

func Base64Encode(src []byte) []byte {
	return []byte(coder.EncodeToString(src))
}

func Base64Decode(src []byte) ([]byte, error) {
	return coder.DecodeString(string(src))
}

func GetMD5(lurl string) string {
	h := md5.New()
	salt1 := "salt4shorturl"
	io.WriteString(h, lurl+salt1)
	urlmd5 := fmt.Sprintf("%x", h.Sum(nil))
	return urlmd5
}
