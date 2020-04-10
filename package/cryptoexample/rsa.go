package cryptoexample

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

var privateKeyPem = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDY79PLCxLf1SQgpEimdl9sCF0A3ouA/sbD3fjwRVaiC7RFmCKY
NV/mFTo2EDVBFvP7b7C6EXbC0fM9qDPWH9xuh5ArnEnIPd3F+RpY5zXAhkd2QhCv
pzfvTIfyKzlbg4lBWqYzrDowCDTI0NiclLRJv0Y2v5WNUq3EIx6ErwLEDwIDAQAB
AoGAQMrwoPY/vua0EaO/pyg3u9aLoJTXacGusBV+IpUzGNcSEq8rtfZLHDc+2aLh
pP0sBe8IA6rvo6R9V+8C/HMrrLKTg9FDqDV8FGRq1SCpfuwwxFW3pHXDUcSV6WJk
QhZ+jPrMETgorsWjnIv61ecOPDZ/CNVj1FHHZaiuig1ZqyECQQD3hPGDz0imXzqg
RidQnWaQkdLabeJVWES+/wsfxXb482FR6Q4SYHIDsnQZuZh17PbpfLMC44nDx8Yw
7t/9WyGJAkEA4F6jDdbkcFVSgEIB4l093wJZsAoKvpyhr9S7g+k6OOobp6E8qyr5
Y9+AECRai3xklhQE190tT2R+y4ymB4hK1wJBAMyaioIosKdGhNHD6+/JjOToheGl
f7iItJslfG6Q7l2v4byx573tF5JSy4IQVyTz8s7jE57JtDGwS/ZbH7DwyfECQQC0
mXiF20NlaEhQFNGPc54ps6qdmHetlkZPUdzeAQ1sYoSAGbjLznuPeIeMdrarMCWG
5/Y+Czo9fZbSIRRDW+FhAkBf0Igkn4ygVW1K3KL5AbdkxQ26OUlDa7eroZ/LdH6x
TbJGzx5rpDqUzbVVecokt1iwBuU+GfUZBJydiKo+zkUK
-----END RSA PRIVATE KEY-----
`)

var publicKeyPem = []byte(`
-----BEGIN RSA PUBLIC KEY-----
MIGJAoGBANjv08sLEt/VJCCkSKZ2X2wIXQDei4D+xsPd+PBFVqILtEWYIpg1X+YV
OjYQNUEW8/tvsLoRdsLR8z2oM9Yf3G6HkCucScg93cX5GljnNcCGR3ZCEK+nN+9M
h/IrOVuDiUFapjOsOjAINMjQ2JyUtEm/Rja/lY1SrcQjHoSvAsQPAgMBAAE=
-----END RSA PUBLIC KEY-----
`)

// RsaEncryptWithPKCS1v15 example
func RsaEncryptWithPKCS1v15(msg []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKeyPem)
	publicKey, _ := x509.ParsePKCS1PublicKey(block.Bytes)

	return rsa.EncryptPKCS1v15(rand.Reader, publicKey, msg)
}

// RsaDecryptWithPKCS1v15 example
func RsaDecryptWithPKCS1v15(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKeyPem)
	privateKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)

	return rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
}

// RsaEncryptWithOAEP example
func RsaEncryptWithOAEP(msg []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKeyPem)
	publicKey, _ := x509.ParsePKCS1PublicKey(block.Bytes)
	return rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, msg, []byte("orders"))
}

// RsaDecryptWithOAEP example
func RsaDecryptWithOAEP(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKeyPem)
	privateKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)

	return rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, ciphertext, []byte("orders"))
}

// RsaSignWithPKCS1v15 example
func RsaSignWithPKCS1v15(msg []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKeyPem)
	privateKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)

	hashed := sha256.Sum256(msg)

	return rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
}

// RsaVerifyPKCS1v15 example
func RsaVerifyPKCS1v15(msg []byte, signature []byte) error {
	block, _ := pem.Decode(publicKeyPem)
	publicKey, _ := x509.ParsePKCS1PublicKey(block.Bytes)

	hashed := sha256.Sum256(msg)
	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], signature)

	if err != nil {
		return err
	}

	return nil
}

// GenRsaKey example
func GenRsaKey() {
	bits := 1024

	privateKey, _ := rsa.GenerateKey(rand.Reader, bits)
	privDerForm := x509.MarshalPKCS1PrivateKey(privateKey)
	privBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privDerForm,
	}
	fmt.Printf("%s", pem.EncodeToMemory(privBlock))

	pubDerForm := x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)
	pubBlock := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubDerForm,
	}
	fmt.Printf("%s", string(pem.EncodeToMemory(pubBlock)))

}
