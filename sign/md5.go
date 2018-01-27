package sign

import (
	"crypto/md5"
)

type MD5Client struct{}

var MD5 = MD5Client{}

// Encrypt is
func (*MD5Client) Encrypt(p []byte) []byte {
	result := md5.Sum(p)
	return result[:]
}

// EncryptWithSalt is给要加密的信息加把盐
func (*MD5Client) EncryptWithSalt(plantext []byte, salt []byte) []byte {
	hash := md5.New()
	hash.Write(plantext)
	hash.Write(salt)
	return hash.Sum(nil)
}
