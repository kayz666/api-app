package utils

import (
	"crypto/md5"
	"encoding/hex"
	"time"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)
const (
	AES_128_KEY="KAY443622@!.1234"
)
//加密函数
//parms word  string  待加密数据
//parms al	string	  加密算法
func Encrypt(word string,al string)(string){
	switch al {
	case "AES_128":
		result,_ :=AesEncrypt([]byte(word),[]byte(AES_128_KEY))
		return fmt.Sprintf("AES_128:|%s",base64.StdEncoding.EncodeToString(result))

	}
	return ""
}

//解密函数
func Decryption(cryptogram string)(word string){



	return ""
}

func StrToMD5v1(s string) string{
	md5ctx := md5.New()
	md5ctx.Write([]byte(s))
	cipher :=md5ctx.Sum([]byte("wanglewo"))
	return hex.EncodeToString(cipher)
}

func StrToMD5v2(s string) string{
	md5ctx := md5.New()
	md5ctx.Write([]byte(s))
	cipher :=md5ctx.Sum([]byte("A"))
	return hex.EncodeToString(cipher)
}

func StrToMD5v3(s string) string{
	md5ctx := md5.New()
	md5ctx.Write([]byte(s))
	cipher :=md5ctx.Sum([]byte("Device"))
	return hex.EncodeToString(cipher)
}


func GetDateTime() string{

	return time.Now().String()
}



func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}