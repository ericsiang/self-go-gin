package jwt_secret

import (
	"api/common/common_const"
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

var jwtSecret []byte

type Claims struct {
	UserID  uint `json:"user_id"`
	AdminID uint `json:"admin_id"`
	jwt.MapClaims
}

func SetJwtSecret(secret string) {
	jwtSecret = []byte(secret)
}

// GenerateToken 根據用戶的用戶id 生成JWT token
func GenerateToken(checkLoginRole common_const.LoginRole, loginId uint) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(1 * time.Hour)
	claims := Claims{}
	switch checkLoginRole {
	case common_const.LoginUser:
		claims.UserID = loginId
	case common_const.LoginAdmin:
		claims.AdminID = loginId
	default:
		return "", errors.New("LoginRole not allow")
	}
	claims.MapClaims = jwt.MapClaims{
		"exp": expireTime.Unix(),
		"iss": "gin-blog",
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret) //簽名字串
	return token, err

}

func ParseToken(token string) (*Claims, error) {
	//用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
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
