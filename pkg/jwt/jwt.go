package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Claims 自定义JWT声明结构
type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

// JWTConfig JWT配置
type JWTConfig struct {
	SecretKey     string
	ExpiresAt     time.Duration
	Issuer        string
	RefreshExpire time.Duration
}

// JWT JWT工具类
type JWT struct {
	config JWTConfig
}

// NewJWT 创建JWT实例
// secretKey: 密钥
// issuer: 签发者
// expire: 访问令牌过期时间
// refreshExpire: 刷新令牌过期时间
func NewJWT(secretKey, issuer string, expire, refreshExpire time.Duration) *JWT {
	return &JWT{
		config: JWTConfig{
			SecretKey:     secretKey,
			ExpiresAt:     expire,
			Issuer:        issuer,
			RefreshExpire: refreshExpire,
		},
	}
}

// GenerateToken 生成访问令牌
func (j *JWT) GenerateToken(userID uint, username, email string) (string, error) {
	claims := Claims{
		UserID:   userID,
		Username: username,
		Email:    email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.config.ExpiresAt)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    j.config.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.config.SecretKey))
}

// GenerateRefreshToken 生成刷新令牌
func (j *JWT) GenerateRefreshToken(userID uint, username, email string) (string, error) {
	claims := Claims{
		UserID:   userID,
		Username: username,
		Email:    email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.config.RefreshExpire)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    j.config.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.config.SecretKey))
}

// ParseToken 解析令牌
func (j *JWT) ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.config.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// RefreshToken 刷新令牌
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	claims, err := j.ParseToken(tokenString)
	if err != nil {
		return "", err
	}

	return j.GenerateToken(claims.UserID, claims.Username, claims.Email)
}

var (
	// ErrTokenExpired 令牌已过期
	ErrTokenExpired = errors.New("token has expired")
	// ErrTokenInvalid 令牌无效
	ErrTokenInvalid = errors.New("token is invalid")
	// ErrTokenMalformed 令牌格式错误
	ErrTokenMalformed = errors.New("token is malformed")
	// ErrTokenNotValidYet 令牌尚未生效
	ErrTokenNotValidYet = errors.New("token is not valid yet")
)
