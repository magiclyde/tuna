/**
 * Created by GoLand.
 * @author: clyde
 * @date: 2021/10/4 下午11:10
 * @note:
 */

package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

// AesGcmEncrypt
func AesGcmEncrypt(plaintext, aseKey []byte) ([]byte, error) {
	block, err := aes.NewCipher(aseKey)
	if err != nil {
		return nil, err
	}

	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := aesGcm.Seal(nonce, nonce, plaintext, nil)
	return ciphertext, nil
}

// AesGcmDecrypt
func AesGcmDecrypt(ciphertext, aseKey []byte) ([]byte, error) {
	block, err := aes.NewCipher(aseKey)
	if err != nil {
		return nil, err
	}

	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := ciphertext[:12]
	plaintext, err := aesGcm.Open(nil, nonce, bytes.TrimPrefix(ciphertext, nonce), nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
