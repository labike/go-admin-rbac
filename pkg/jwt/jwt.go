package jwt

import (
	"errors"
	"fmt"
	"go-admin/api/entity"
	"go-admin/common/constant"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type userStdClaims struct {
	entity.JwtAdmin
	jwt.StandardClaims
}

const TokenExpiresDuration = time.Hour * 24

var Secret = []byte("jp3j9e-03r;$Tety3,cv;.nn[ugdgy2e-zk6")
var (
	ErrAbsent  = "token absent"  // 令牌不存在
	ErrInvalid = "token invalid" // 令牌无效
)

func GenerateTokenByAdmin(admin entity.SysAdmin) (string, error) {
	var jwtAdmin = entity.JwtAdmin{
		ID:       admin.ID,
		Username: admin.Username,
		Nickname: admin.Nickname,
		Icon:     admin.Icon,
		Email:    admin.Email,
		Phone:    admin.Phone,
		Note:     admin.Note,
	}

	c := userStdClaims{
		jwtAdmin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpiresDuration).Unix(),
			Issuer:    "admin",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(Secret)
}

// 解析jwt
func ValidateToken(tokenString string) (*entity.JwtAdmin, error) {
	if tokenString == "" {
		return nil, errors.New(ErrAbsent)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})

	if token == nil {
		return nil, errors.New(ErrInvalid)
	}

	claims := userStdClaims{}
	_, err = jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return Secret, nil
	})

	if err != nil {
		return nil, err
	}

	return &claims.JwtAdmin, err
}

// 返回用户信息
func GetAdminId(c *gin.Context) (uint, error) {
	u, exists := c.Get(constant.ContextKeyUserObj)
	if !exists {
		return 0, errors.New("can't get user id")
	}
	admin, ok := u.(entity.JwtAdmin)
	if ok {
		return admin.ID, nil
	}
	return 0, errors.New("can't convert to id struct")
}

func GetAdminName(c *gin.Context) (string, error) {
	u, exists := c.Get(constant.ContextKeyUserObj)
	if !exists {
		return string(string(0)), errors.New("can't get user name")
	}
	admin, ok := u.(*entity.JwtAdmin)
	if ok {
		return admin.Username, nil
	}
	return string(string(0)), errors.New("can't convert to username struct")
}

func GetAdmin(c *gin.Context) (*entity.JwtAdmin, error) {
	u, exists := c.Get(constant.ContextKeyUserObj)
	if !exists {
		return nil, errors.New("can't get user info")
	}
	admin, ok := u.(entity.JwtAdmin)
	if ok {
		return &admin, nil
	}
	return nil, errors.New("can't convert to user info")
}
