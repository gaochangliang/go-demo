package util

import (
	"crypto/md5"
	"encoding/hex"
)

//The name of the uploaded image is not usually exposed directly, so we run MD5 on the image name to achieve this effect
func EncodeMD5(imgName string) string {
	m := md5.New()
	m.Write([]byte(imgName))
	return hex.EncodeToString(m.Sum(nil))
}
