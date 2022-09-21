/**
 * Created by GoLand.
 * @author: clyde
 * @date: 2022/9/21 下午3:20
 * @note: aes-128-gcm 用于防沉迷加密
 */

package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
)

func GCMEncrypt(originalText, secretKey string) (string, error) {
	key, err := hex.DecodeString(secretKey)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// iv 需放在密文开头，所以 dst 初始设定为 nonce，随后密文会 append 到 dst
	enc := aesGcm.Seal(nonce, nonce, S2B(originalText), nil)

	return base64.StdEncoding.EncodeToString(enc), nil
}

func GCMDecrypt(cipherText, secretKey string) (string, error) {
	enc, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	key, err := hex.DecodeString(secretKey)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGcm.NonceSize()

	// iv 放在密文头部，解密前需去掉
	plain, err := aesGcm.Open(nil, enc[:nonceSize], enc[nonceSize:], nil)
	if err != nil {
		return "", err
	}

	return B2S(plain), nil
}
