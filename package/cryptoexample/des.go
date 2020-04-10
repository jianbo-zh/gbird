package cryptoexample

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

// DesEncrypte example
func DesEncrypte() {
	key := []byte("12345678")

	rawStr := "Hello World"
	rawBytes := PKCS7Padding([]byte(rawStr), des.BlockSize)

	// des加密算法，密钥必须8位
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}

	cipherText := make([]byte, block.BlockSize()+len(rawBytes))

	// iv初始化向量
	iv := cipherText[:block.BlockSize()]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	// 使用CBC加密模式
	cbcMode := cipher.NewCBCEncrypter(block, iv)

	cbcMode.CryptBlocks(cipherText[block.BlockSize():], rawBytes)

	fmt.Printf("DES Encrypte: %s -> %x\n", rawStr, cipherText)
}

// DesDecrypte example
func DesDecrypte() {
	key := []byte("12345678")
	cipherText := "0abaf77078453689a4d5e0c5f3b461b03942eb78ef2112fd"
	cipherBytes, _ := hex.DecodeString(cipherText)

	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	iv := cipherBytes[:block.BlockSize()]
	cbcMode := cipher.NewCBCDecrypter(block, iv)
	decrypteText := make([]byte, len(cipherBytes)-block.BlockSize())
	cbcMode.CryptBlocks(decrypteText, cipherBytes[block.BlockSize():])

	rawStr := PKCS7UnPadding(decrypteText)
	fmt.Printf("DES Decrypte: %s -> %s\n", cipherText, rawStr)
}

// PKCS7Padding test 填充到blockSize长度的整数倍, PKCS5Padding固定块大小为8位
func PKCS7Padding(cipherText []byte, blockSize int) []byte {
	num := blockSize - len(cipherText)%blockSize
	padding := bytes.Repeat([]byte{byte(num)}, num)
	return append(cipherText, padding...)
}

// PKCS7UnPadding 去掉填充数据
func PKCS7UnPadding(paddingText []byte) []byte {
	length := len(paddingText)
	num := length - int(paddingText[length-1])
	return paddingText[:num]
}
