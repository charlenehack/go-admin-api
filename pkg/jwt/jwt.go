package jwt

import (
	"admin-api/api/entity"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

type userStdClaims struct {
	JwtUser entity.JwtUser
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 2 // token过期时间
var Secret = []byte("akdfjaldkfj")

// 根据用户信息生成token
func GenerateTokenByUser(user entity.User) (string, error) {
	var jwtUser = entity.JwtUser{
		ID:          user.ID,
		Username:    user.Username,
		Nickname:    user.Nickname,
		Avatar:      user.Avatar,
		Email:       user.Email,
		Phone:       user.Phone,
		Description: user.Description,
	}
	claims := userStdClaims{
		JwtUser: jwtUser,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "admin",                                    // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(Secret)
}

// 解析token
func ValidateToken(tokenString string) (*entity.JwtUser, error) {
	if tokenString == "" {
		return nil, errors.New("token不存在")
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if token == nil {
		return nil, errors.New("token无效")
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
	return &claims.JwtUser, err
}

// 解析token后返回id
func GetUserId(c *gin.Context) (int, error) {
	u, exist := c.Get("authedUserObj")
	if !exist {
		return 0, errors.New("无法获取用户id")
	}
	user, ok := u.(*entity.JwtUser)
	if ok {
		return user.ID, nil
	}
	return 0, errors.New("获取用户id失败")
}

// 解析token后返回用户名
func GetUserName(c *gin.Context) (string, error) {
	u, exist := c.Get("authedUserObj")
	if !exist {
		return "0", errors.New("无法获取用户名")
	}
	user, ok := u.(*entity.JwtUser)
	if ok {
		return user.Username, nil
	}
	return "0", errors.New("获取用户名失败")
}

// 解析token后返回用户信息
func GetUser(c *gin.Context) (*entity.JwtUser, error) {
	u, exist := c.Get("authedUserObj")
	if !exist {
		return nil, errors.New("无法获取用户名")
	}
	user, ok := u.(*entity.JwtUser)
	if ok {
		return user, nil
	}
	return nil, errors.New("获取用户名失败")
}
