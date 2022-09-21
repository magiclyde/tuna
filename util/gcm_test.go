/**
 * Created by GoLand.
 * @author: clyde
 * @date: 2022/9/21 下午3:36
 * @note:
 */

package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCrypto(t *testing.T) {
	originalText := `{"ai":"test-accountId","name":"某某某","idNum":"371321199012310912"}`
	secretKey := "2836e95fcd10e04b0069bb1ee659955b"

	cipherText, err := GCMEncrypt(originalText, secretKey)
	if err != nil {
		t.Error(err)
		return
	}
	//t.Log(cipherText)

	dec, err := GCMDecrypt(cipherText, secretKey)
	if err != nil {
		t.Error(err)
		return
	}
	//t.Log(dec)
	assert.Equal(t, originalText, dec)
}
