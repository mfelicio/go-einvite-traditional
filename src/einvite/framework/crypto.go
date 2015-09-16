package framework

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"time"
	//"encoding/hex"
	//"fmt"
	mrand "math/rand"
)

type security struct {
}

var Security = &security{}

func (this *security) Sign(value string, key []byte) string {

	raw := this.SignRaw([]byte(value), key)

	//sig := hex.EncodeToString(raw)
	sig := base64.URLEncoding.EncodeToString(raw)

	return sig
}

func (this *security) Encrypt(value string, key []byte) string {

	raw := this.EncryptRaw([]byte(value), key)

	//out := hex.EncodeToString(raw)
	out := base64.URLEncoding.EncodeToString(raw)
	return out
}

func (this *security) Decrypt(value string, key []byte) string {

	//rawValue, _ := hex.DecodeString(value)
	rawValue, _ := base64.URLEncoding.DecodeString(value)
	raw := this.DecryptRaw(rawValue, key)
	out := string(raw)
	return out
}

func (this *security) EncryptAndSign(value string, signKey []byte, cipherKey []byte) (signed string, encrypted string) {

	rawSign, rawEncrypted := this.EncryptAndSignRaw([]byte(value), signKey, cipherKey)

	//signed = hex.EncodeToString(rawSign)
	//encrypted = hex.EncodeToString(rawEncrypted)

	signed = base64.URLEncoding.EncodeToString(rawSign)
	encrypted = base64.URLEncoding.EncodeToString(rawEncrypted)

	return
}

func (this *security) VerifySignature(signature string, value string, signKey []byte) bool {

	//signatureRaw, _ := hex.DecodeString(signature)
	//valueRaw, _ := hex.DecodeString(value)
	signatureRaw, _ := base64.URLEncoding.DecodeString(signature)
	valueRaw, _ := base64.URLEncoding.DecodeString(value)

	validSignature := this.SignRaw(valueRaw, signKey)

	return hmac.Equal(signatureRaw, validSignature)

	//validSignature := this.Sign(value, signKey)
	//return validSignature == signature
}

func (this *security) EncryptAndSignRaw(value []byte, signKey []byte, cipherKey []byte) (signed []byte, encrypted []byte) {

	encrypted = this.EncryptRaw(value, cipherKey)
	signed = this.SignRaw(encrypted, signKey)
	return
}

func (this *security) SignRaw(value []byte, key []byte) []byte {

	hashFunc := sha1.New
	mac := hmac.New(hashFunc, key)

	mac.Write(value)

	out := mac.Sum(nil)

	mac.Reset()
	return out
}

func (this *security) EncryptRaw(value []byte, key []byte) []byte {

	block, _ := aes.NewCipher(key)

	iv := this.generateRandomKey(block.BlockSize())

	// Encrypt it.
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(value, value)

	// Return iv + ciphertext.
	out := append(iv, value...)

	return out
}

func (this *security) DecryptRaw(value []byte, key []byte) []byte {

	block, _ := aes.NewCipher(key)

	size := block.BlockSize()

	// Extract iv.
	iv := value[:size]
	// Extract ciphertext.
	value = value[size:]
	// Decrypt it.
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(value, value)
	out := value

	return out
}

var mathRandomzier = mrand.New(mrand.NewSource(time.Now().UnixNano()))

func (this *security) generateRandomKey(strength int) []byte {

	k := make([]byte, strength)

	if bytesRead, err := rand.Read(k); err != nil {

		//assuming it may read some bytes.. continue from there
		for i := bytesRead; i < strength; i++ {
			k[i] = byte(mathRandomzier.Int())
		}
	}

	return k
}
