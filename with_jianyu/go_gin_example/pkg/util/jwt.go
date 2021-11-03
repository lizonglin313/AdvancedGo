package util

import (
	"AdvancedGo/with_jianyu/go_gin_example/pkg/setting"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// GenerateToken
// @Desc: 	获取token
// @Param:	username
// @Param:	password
// @Return:	string
// @Return:	error
// @Notice:	
func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	// 最后需要转成 int64 类型
	expireTime := nowTime.Add(3 * time.Hour).Unix()

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime,
			Issuer:    "gin_blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成完整签名字符串然后再用于获取完整的、已签名的token
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	// 解析鉴权声明
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
