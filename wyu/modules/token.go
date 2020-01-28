package modules

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
)

/**
 * 	signature, err := s.srv.Parent.Token.RsaSign([]byte("test"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	sign := s.srv.Parent.Token.HexEncodeToString(signature)
	signByte, _ := s.srv.Parent.Token.HexStringToEncode(sign)
	fmt.Println(signature)
	fmt.Println(sign)
	fmt.Println(signByte)
	fmt.Println(s.srv.Parent.Token.RsaSignVerify([]byte("test"), signByte))
	fmt.Println("-----")

	cipher, _ := s.srv.Parent.Token.RsaEncrypt([]byte("test success"))
	fmt.Println(cipher)
	cipherTxt := s.srv.Parent.Token.HexEncodeToString(cipher)
	fmt.Println(cipherTxt)
	fmt.Println(s.srv.Parent.Token.HexStringToEncode(cipherTxt))
	data, _ := s.srv.Parent.Token.RsaDecrypt(cipher)
	fmt.Println(string(data))
**/

var (
	pubKey string = `-----BEGIN RSA PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC36Xn/4fHhKqayEqQ0Ymmk/5+g
gDvgCW8TBR4JG3igUvU6tMLUe1KAGzMAQjJ6ArPmZejgMjxD6rGTWj2vtRoBBuZM
C9jPpkZwKn8Qno0detWD67UlhIoRMFt6rE+fAxLlJlsHrkHXgII8/iVKjOip3fag
ucCdbGiiFhAg1i7BWwIDAQAB
-----END RSA PUBLIC KEY-----`
	prvKey string = `-----BEGIN RSA PRIVATE KEY-----
MIICWwIBAAKBgQC36Xn/4fHhKqayEqQ0Ymmk/5+ggDvgCW8TBR4JG3igUvU6tMLU
e1KAGzMAQjJ6ArPmZejgMjxD6rGTWj2vtRoBBuZMC9jPpkZwKn8Qno0detWD67Ul
hIoRMFt6rE+fAxLlJlsHrkHXgII8/iVKjOip3fagucCdbGiiFhAg1i7BWwIDAQAB
AoGALPVpepEshSPdDkkaVSf9tXU7+4t9l54WxiqJFibeDStagZhwzGq9V03O4PBN
0J/ahKdDD5OYQe4crO3xiIOYMF/zYPJy6lY+SGp/0qlJwDx9b7eh3PLUkiDoaetX
JZshRI2AiTFH9dv2KPxvvQ8kFQoZZrr3XWzzHkoxoVnz/EECQQDTALttH/tN1GGM
kCDa8EgmMzvgZGzGBROolbyGUAdWu9YUCwJSFWpacR1jRNCYvBmELW8POp2m996O
KHz3Qg/7AkEA3yHGiTdnHD4dXnoQO1InqquT8Ygo26QWEQYLjgVUbdUG0GXOI7OC
xQKNKeTANywAcNIUwRF9YJqoL47eL1d2IQJAQ5UlcxNeS5Rt1jbHvzhc85dPY1Tn
Hhm8LTAgnSh+4UHylKLeEGp5kRRP5F7DLVh6F8LxooAUxMj5iLDhLdUEBwJAV4Yd
JXfY90gaJxQER/Ca5KR23LhHJpi/mx/e6m+GxapZCOfWK0Tf171/d95l035sEdUm
FPFyV7FypW0KFFHfYQJAHq1NyaHymQ9ijlj7t1txZxqXcO8K5Xd5iSB7PZMmDUSJ
cE4spT93Vk8V3CIa3M9yswI6MhwXRuAIjAn2BRSRDw==
-----END RSA PRIVATE KEY-----`
)

type Token struct {
	PubKey []byte
	PrvKey []byte
}

func NewToken(keys ...string) (token *Token) {
	token = &Token{}

	if len(keys) != 2 {
		token.PubKey = []byte(pubKey)
		token.PrvKey = []byte(prvKey)
	} else {
		token.PubKey = []byte(keys[0])
		token.PrvKey = []byte(keys[1])
	}

	return
}

func (token *Token) GenToken(keys ...string) (cipher string, err error) {
	if len(keys) < 1 {
		err = errors.New("empty keys")
		return
	}

	var strKey string = ""
	for _, key := range keys {
		strKey = strKey + key
	}

	random, err := UtilsRandUUID(3)
	if err != nil {
		return
	}

	strKey = strKey + random

	h := sha256.New()
	h.Write([]byte(strKey))
	cipher = hex.EncodeToString(h.Sum([]byte("")))

	return
}

func (token *Token) GenRsaKey() (prvKey, pubKey []byte, err error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return
	}

	prvKey = pem.EncodeToMemory(&pem.Block{
		Type: "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	x, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return
	}

	pubKey = pem.EncodeToMemory(&pem.Block{
		Type: "RSA PUBLIC KEY",
		Bytes: x,
	})

	return
}

func (token *Token) HexEncodeToString(signature []byte) (sign string) {
	sign = hex.EncodeToString(signature)
	return
}

func (token *Token) HexStringToEncode(sign string) (signature []byte, err error) {
	signature, err = hex.DecodeString(sign)
	if err != nil {
		return
	}

	return
}

/**
 * Todo: signature by Private Key
 */
func (token *Token) RsaSign(data []byte) (signature []byte, err error) {
	h := sha256.New()
	h.Write(data)
	hashed := h.Sum([]byte(""))

	block, _ := pem.Decode(token.PrvKey)
	if block == nil {
		err = errors.New("error private key")
		return
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return
	}

	signature, err = rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		return
	}

	return
}

/**
 * Todo: verify signature by Public Key
 */
func (token *Token) RsaSignVerify(data, signature []byte) (ok bool, err error) {
	block, _ := pem.Decode(token.PubKey)
	if block == nil {
		err = errors.New("error public key")
		return
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return
	}

	hashed := sha256.Sum256(data)
	err = rsa.VerifyPKCS1v15(publicKey.(*rsa.PublicKey), crypto.SHA256, hashed[:], signature)
	if err != nil {
		return
	}

	ok = true
	return
}

/**
 * Todo: encrypt cipher by public key
 */
func (token *Token) RsaEncrypt(data []byte) (cipher []byte, err error) {
	block, _ := pem.Decode(token.PubKey)
	if block == nil {
		err = errors.New("error public key")
		return
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return
	}

	cipher, err = rsa.EncryptPKCS1v15(rand.Reader, publicKey.(*rsa.PublicKey), data)
	if err != nil {
		return
	}

	return
}

/**
 * Todo: decrypt cipher by Private Key
 */
func (token *Token) RsaDecrypt(cipher []byte) (data []byte, err error) {
	block, _ := pem.Decode(token.PrvKey)
	if block == nil {
		err = errors.New("error private key")
		return
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return
	}

	data, err = rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipher)
	if err != nil {
		return
	}

	return
}
