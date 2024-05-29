package utils

import (
	"blog_server/config"
	"blog_server/global"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type MyCustomClaims struct {
	UserID     int    `json:"user_id"`
	Username   string `json:"username"`
	GrantScope string `json:"grant_scope"`
	Role       int    `json:"role"`
	jwt.RegisteredClaims
}

// pkcs1
func parsePriKeyBytes(buf []byte) (*rsa.PrivateKey, error) {
	p := &pem.Block{}
	p, buf = pem.Decode(buf)
	if p == nil {
		return nil, errors.New("parse key error")
	}
	return x509.ParsePKCS1PrivateKey(p.Bytes)
}

// GenerateTokenUsingRS256 生成token
func GenerateTokenUsingRS256(userID uint, username string) (string, error) {
	claim := MyCustomClaims{
		UserID:   int(userID),
		Username: username,
		//GrantScope: viper.GetString("jwt.grant_scope"),
		GrantScope: global.Config.Jwt.GrantScope,
		RegisteredClaims: jwt.RegisteredClaims{
			//Issuer:    viper.GetString("jwt.issuer"),                                                                  // 签发者
			Issuer: global.Config.Jwt.Issuer, // 签发者
			//Subject:   viper.GetString("jwt.subject"),                                                                 // 签发对象
			Subject:  global.Config.Jwt.Subject,                             // 签发对象
			Audience: jwt.ClaimStrings{"Android_APP", "IOS_APP", "WEB_APP"}, //签发受众
			//ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(viper.GetInt("jwt.expire_time")))), //过期时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(global.Config.Jwt.Expires))), //过期时间
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Second)),                                          //最早使用时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                                           //签发时间
			ID:        GenerateSalt(10),                                                                         // jwt ID, 类似于盐值
		},
	}
	//rsaPriKey, err := parsePriKeyBytes([]byte(viper.GetString("jwt.private_key")))
	rsaPriKey, err := parsePriKeyBytes([]byte(config.PRI_KEY))
	if err != nil {
		return "", err
	}
	//global.Log.Infof("查看rsa_pri_key：%v", rsaPriKey)
	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claim).SignedString(rsaPriKey)
	return token, err
}

// parsePubKeyBytes 解析公钥
func parsePubKeyBytes(pub_key []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pub_key)
	if block == nil {
		return nil, errors.New("block nil")
	}
	pub_ret, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, errors.New("x509.ParsePKCS1PublicKey error")
	}

	return pub_ret, nil
}

// ParseTokenRs256 解析token
func ParseTokenRs256(token_string string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(token_string, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		pub, err := parsePubKeyBytes([]byte(config.PUB_KEY))
		if err != nil {
			global.Log.Errorf("parsePubKeyBytes error: %v", err)
			return nil, err
		}
		return pub, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("claim invalid")
	}

	claims, ok := token.Claims.(*MyCustomClaims)
	if !ok {
		return nil, errors.New("invalid claim type")
	}

	return claims, nil
}
