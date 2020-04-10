package jwtlogin

import (
	"github.com/ctbsea/Go-Message/entry"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const JWT_KEY = "loginSign"
const JWT_SLAT = "HI"
const JWT_EXP = 30 * time.Minute

var jwtSalt = []byte(JWT_SLAT)

type ClaimsInfo struct {
	UserID   uint   `json:"uid"`
	Username string `json:"uname"`
}

type Claims struct {
	Login *ClaimsInfo
	jwt.StandardClaims
}

//生成JWT Token
func Sign(UserInfo *ClaimsInfo) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Login: UserInfo,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "Iris",
			IssuedAt:  time.Now().Unix(),
			Id:        "Iris-9597",
			ExpiresAt: time.Now().Add(JWT_EXP).Unix(),
		},
	})
	str, _ := token.SignedString(jwtSalt)
	return str
}

//检查JWT TOKEN
func Check(token string) (*ClaimsInfo, int) {
	m, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSalt, nil
	})
	if err != nil || m == nil || !m.Valid {
		return &ClaimsInfo{}, entry.JWT_ERR_TOKEN
	}
	if claims, ok := m.Claims.(*Claims); ok {
		return claims.Login, entry.SUCCESS
	}
	return &ClaimsInfo{}, entry.JWT_ERR_TOKEN
}

//刷新
func Refresh(token string) (string, int, time.Time) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtSalt, nil
		})
	if err != nil {
		return "", entry.JWT_ERR_TOKEN, time.Now()
	}
	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		jwt.TimeFunc = time.Now
		expireTime := jwt.TimeFunc().Add(JWT_EXP) //有效期 三个小时
		claims.StandardClaims.ExpiresAt = expireTime.Unix()
		newToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSalt)
		return newToken, 0, expireTime
	}
	return "", entry.JWT_ERR_TOKEN, time.Now()
}
