package framework

import (
	"fmt"
	"testing"
)

var sign_key = []byte("signKeyW/16Chars")
var cipher_key = []byte("ciphKeyW/16Chars")

func Test_Sign(t *testing.T) {

	s1 := "mydata"
	h1 := Security.Sign(s1, sign_key)

	s2 := "myData"
	h2 := Security.Sign(s2, sign_key)

	s3 := "mydata"
	h3 := Security.Sign(s3, sign_key)

	if h1 == h2 {
		t.Errorf("Same signature for %s and %s", s1, s2)
	}

	if h1 != h3 {
		t.Errorf("%s and %s have different signatures: %s and %s)", s1, s3, h1, h3)
	}

}

func Test_EncryptDecrypt(t *testing.T) {

	fmt.Print("") // just to use it and avoid having to import and remove

	s1 := "mydata"
	h1 := Security.Encrypt(s1, cipher_key)

	d1 := Security.Decrypt(h1, cipher_key)

	if s1 != d1 {
		t.Errorf("s1 %s and d1 %s", s1, d1)
	}
}

func Test_EncryptAndSign(t *testing.T) {
	s1 := "abcdefghijklmnopqrstuvwxyz0123456789"
	h1, e1 := Security.EncryptAndSign(s1, sign_key, cipher_key)

	ok := Security.VerifySignature(h1, e1, sign_key)

	if !ok {
		t.Errorf("bad signature %s for string %s", h1, s1)
		return
	}

	d1 := Security.Decrypt(e1, cipher_key)

	if s1 != d1 {
		t.Errorf("Decrypted string %s should be %s", d1, s1)
	}

	fmt.Printf("Passed with s1 %s e1 %s h1 %s", s1, e1, h1)
	fmt.Println()
	fmt.Printf("SignEncrypt length %d for string with length %d", len(h1)+len(e1), len(s1))
	fmt.Println()
}
