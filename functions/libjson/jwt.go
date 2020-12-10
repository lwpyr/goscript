package libjson

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	jsoniter "github.com/json-iterator/go"
	"github.com/lwpyr/goscript/common"
	"io/ioutil"
	"net/http"
)

func JwtParse(_ *common.Memory, stk *common.Stack) {
	token := (stk.Top()).(string)
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return getTokenKey(token)
	})

	if err != nil {
		panic(err)
	}

	m := make(map[string]interface{})

	headerMap := make(map[interface{}]interface{})

	for k, v := range parsedToken.Header {
		headerMap[k] = v
	}

	m["header"] = headerMap
	m["payload"] = jsonObjectify(parsedToken.Claims)
	m["signature"] = parsedToken.Signature

	stk.ReturnValue(m)
}

func getTokenKey(token *jwt.Token) (interface{}, error) {
	resp, err := http.Get(token.Header["x5u"].(string))
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()
	key, err := ioutil.ReadAll(resp.Body)

	parsedKey, err := parseRSAPublicKeyFromPEM(key)
	if err != nil {
		return false, err
	}

	return parsedKey, nil
}

func parseRSAPublicKeyFromPEM(key []byte) (*rsa.PublicKey, error) {
	var err error

	// Parse PEM block
	var block *pem.Block
	if block, _ = pem.Decode(key); block == nil {
		return nil, fmt.Errorf("invalid KeyType: KeyType must be a PEM encoded PKCS1 or PKCS8 key")
	}

	// Parse the key
	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKIXPublicKey(block.Bytes); err != nil {
		if cert, err := x509.ParseCertificate(block.Bytes); err == nil {
			parsedKey = cert.PublicKey
		} else {
			return nil, err
		}
	}

	var pkey *rsa.PublicKey
	var ok bool
	if pkey, ok = parsedKey.(*rsa.PublicKey); !ok {
		return nil, fmt.Errorf("key is not a valid RSA public key")
	}

	return pkey, nil
}

func jsonObjectify(s interface{}) map[string]interface{} {
	var payload map[string]interface{}
	bytes, _ := jsoniter.Marshal(s)
	_ = jsoniter.Unmarshal(bytes, &payload)
	return payload
}
