package encrypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// 公钥和私钥可以从文件中读取
var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDDrcoyDSZYNWgQENKCyu00aWAqu5MHe0kyS59VnrlVXR4Y8l1t
OLxM7yoybAaOdP/y1rZenbjyga92n+eCufeU7+YebABZ0O2AbMbWO9sn9tzlv3hY
7OF/cw99Zxt5wEghQ1n1/VXCUMrAWCobWi+KFTkNHEgDfZjGEOQo3cuucQIDAQAB
AoGBAK9pUlEt4oq+TWvheKRQvvT15YRJI9NYHFSe39WD9MXmNH3OfhvT+VDKMLyE
hBgeH/cTrOYCY3HY+W7Qh1tz09C40+jK1TVnSbdyxDOtcJo93+7iEZmztnfTupqv
ZFuzRmqNgSx2zt/UF6mQbawic7Dx32Q2jVpqG4koHp6FBYsVAkEA6g8i9+LY7+iV
BH3KlOrVrPqO3s3lJeVSX13WwRczKxV5XZGNOfIVM2CQ9taruAu/xD95LkCaxqv1
cOWwt9kZ1wJBANYFnjraYi2WzCT3ry7GSA4ZIuEZdyhIjWdHLXCRMXRsqNkOQ8kP
tnZA+AI+aJhteoX1/cG9sWi4iV9OvsJSQPcCQDpsstbbqjkgfmoTmEjZ4aJ/HMCi
9osiFhC2FNA4IU6k2pmvpmgLdJ1Rgn4LEewsCp9LFM2l1Ly42dhnjVgm+hsCQQDG
T5JcSjqqr44duvuyRbxCg/wTw/rrcr7Dsepi4caHcJ/L8DHTPiH91Rl5Ssa0Zs0f
D97ABLs8o7F2hIqxmHHlAkAeWFUrTMqzc3HVV3RWJVhsk68iV4HE8uxuA/vkeiPm
JPHvpv1NYCShlvbWa0Om5eAPQJohtlWumRq3yDpkGk8i
-----END RSA PRIVATE KEY-----
`)

var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDDrcoyDSZYNWgQENKCyu00aWAq
u5MHe0kyS59VnrlVXR4Y8l1tOLxM7yoybAaOdP/y1rZenbjyga92n+eCufeU7+Ye
bABZ0O2AbMbWO9sn9tzlv3hY7OF/cw99Zxt5wEghQ1n1/VXCUMrAWCobWi+KFTkN
HEgDfZjGEOQo3cuucQIDAQAB
-----END PUBLIC KEY-----
`)

// 加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
