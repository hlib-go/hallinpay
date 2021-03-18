package hallinpay

import (
	"crypto"
	cryptorand "crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"sort"
	"strings"
)

// 签名
func RsaSign(conf *Config, bm map[string]string) (sign string, err error) {
	value := sortMap(bm, false)

	priBytes, err := base64.StdEncoding.DecodeString(conf.PriKey)
	if err != nil {
		return
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(priBytes)
	if err != nil {
		return
	}

	hash := sha1.New()
	hash.Write([]byte(value))
	shaBytes := hash.Sum(nil)
	b, err := rsa.SignPKCS1v15(cryptorand.Reader, privateKey.(*rsa.PrivateKey), crypto.SHA1, shaBytes)
	if err != nil {
		return
	}
	sign = base64.StdEncoding.EncodeToString(b)
	return
}

// 验签
func RsaVerify(conf *Config, bm map[string]string) (err error) {
	sign := bm["sign"]
	value := sortMap(bm, false)

	signBytes, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		err = errors.New("验签错误，Base64解码签名出错 " + err.Error())
		return
	}
	pubBytes, err := base64.StdEncoding.DecodeString(conf.PubKey)
	if err != nil {
		return
	}
	publicKey, err := x509.ParsePKIXPublicKey(pubBytes)
	if err != nil {
		err = errors.New("验签错误，ParsePKIXPublicKey " + err.Error())
		return
	}
	hash := sha1.New()
	hash.Write([]byte(value))
	shaBytes := hash.Sum(nil)
	err = rsa.VerifyPKCS1v15(publicKey.(*rsa.PublicKey), crypto.SHA1, shaBytes, signBytes)
	if err != nil {
		err = errors.New("SIGN_ERROR " + err.Error())
		return
	}
	return
}

// map排序
// @params containNilVal true空字段参与签名 false空字段不参与签名
func sortMap(m map[string]string, containNilVal bool) string {
	if m == nil {
		return ""
	}
	var (
		buf     strings.Builder
		keyList []string
	)
	for k := range m {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	for _, k := range keyList {
		if "sign" == k {
			continue
		}
		// 不包含value为空的字段
		if !containNilVal && m[k] == "" {
			continue
		}
		buf.WriteString(k)
		buf.WriteByte('=')
		buf.WriteString(m[k])
		buf.WriteByte('&')
	}
	s := buf.String()
	s = s[0 : len(s)-1]
	return s
}
