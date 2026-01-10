package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	secret                []byte
	DefaultAccessExpire   = time.Duration(15) * time.Minute
	DefaultRefreshExpire  = time.Duration(10080) * time.Minute
	DefaultRefreshOverlap = time.Duration(1440) * time.Minute

	ErrInvalidToken = errors.New("jwt: invalid token")
	ErrWrongType    = errors.New("jwt: wrong token type")
	ErrBadSign      = errors.New("jwt: unexpected signing method")
)

const (
	AccessTokenType  = "AccessToken"
	RefreshTokenType = "RefreshToken"
)

type TokenPair struct {
	AccessToken           string
	AccessTokenExpiresAt  time.Time
	RefreshToken          string
	RefreshTokenExpiresAt time.Time
}

type AccessTokenClaims struct {
	jwt.RegisteredClaims
	AccountID uuid.UUID
}

type RefreshTokenClaims struct {
	jwt.RegisteredClaims
	AccountID uuid.UUID
}

func GenerateAccessToken(accountID uuid.UUID) (string, time.Time, error) {
	now := time.Now()
	exp := now.Add(DefaultAccessExpire)

	claims := AccessTokenClaims{
		AccountID: accountID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   AccessTokenType,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS512, claims).SignedString(secret)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("sign access: %w", err)
	}
	return token, exp, nil
}

func GenerateRefreshToken(accountID uuid.UUID) (string, time.Time, error) {
	now := time.Now()
	exp := now.Add(DefaultRefreshExpire)

	claims := RefreshTokenClaims{
		AccountID: accountID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   RefreshTokenType,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS512, claims).SignedString(secret)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("sign refresh: %w", err)
	}
	return token, exp, nil
}

func GenerateTokenPair(accountID uuid.UUID) (*TokenPair, error) {
	acc, accExp, err := GenerateAccessToken(accountID)
	if err != nil {
		return nil, err
	}
	ref, refExp, err := GenerateRefreshToken(accountID)
	if err != nil {
		return nil, err
	}
	return &TokenPair{
		AccessToken:           acc,
		AccessTokenExpiresAt:  accExp,
		RefreshToken:          ref,
		RefreshTokenExpiresAt: refExp,
	}, nil
}

func ValidateAccessToken(tokenString string) (bool, error) {
	claims, err := ParseWithAccessClaims(tokenString)
	if err != nil {
		return false, err
	}
	if claims.Subject != AccessTokenType {
		return false, ErrWrongType
	}
	return true, nil
}

func ValidateRefreshToken(tokenString string) (bool, error) {
	claims, err := ParseWithRefreshClaims(tokenString)
	if err != nil {
		return false, err
	}
	if claims.Subject != RefreshTokenType {
		return false, ErrWrongType
	}
	return true, nil
}

func ParseWithAccessClaims(tokenString string) (*AccessTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS512 {
			return nil, ErrBadSign
		}
		return secret, nil
	})
	if err != nil {
		return nil, fmt.Errorf("jwt parse: %w", err)
	}

	claims, ok := token.Claims.(*AccessTokenClaims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}
	return claims, nil
}

func ParseWithRefreshClaims(tokenString string) (*RefreshTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &RefreshTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS512 {
			return nil, ErrBadSign
		}
		return secret, nil
	})
	if err != nil {
		return nil, fmt.Errorf("jwt parse: %w", err)
	}

	claims, ok := token.Claims.(*RefreshTokenClaims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}
	return claims, nil
}
