package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ErrInvalidToken     = errors.New("jwt: invalid token")
	ErrTokenExpired     = errors.New("jwt: token expired")
	ErrWrongType        = errors.New("jwt: wrong token type")
	ErrBadSign          = errors.New("jwt: unexpected signing method")
	ErrTokenNotInWindow = errors.New("jwt: token not in refresh window")
)

const (
	TypeAccessToken  = "AccessToken"
	TypeRefreshToken = "RefreshToken"
)

type JWTManager struct {
	secret []byte

	// Token lifetimes
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
	overlapWindow   time.Duration
	issuer          string
}

type JWTConfig struct {
	Secret []byte

	// Access token lifetime (default: 15 minutes)
	AccessTokenTTL time.Duration

	// Refresh token lifetime (default: 7 days)
	RefreshTokenTTL time.Duration

	// Overlap window for refresh token (default: 1 day)
	OverlapWindow time.Duration

	// Issuer (optional)
	Issuer string
}

type TokenPair struct {
	AccessToken           string
	AccessTokenExpiresAt  time.Time
	RefreshToken          string
	RefreshTokenExpiresAt time.Time
}

// AccessTokenClaims - kept exactly as provided
type AccessTokenClaims struct {
	jwt.RegisteredClaims
	AccountID uuid.UUID
}

// RefreshTokenClaims - kept exactly as provided
type RefreshTokenClaims struct {
	jwt.RegisteredClaims
	AccountID uuid.UUID
}

// New creates a new JWTManager instance with the given config
func New(config *JWTConfig) (*JWTManager, error) {
	if config == nil {
		return nil, errors.New("config cannot be nil")
	}

	// Set defaults if not provided
	accessTTL := config.AccessTokenTTL
	if accessTTL == 0 {
		accessTTL = 15 * time.Minute
	}

	refreshTTL := config.RefreshTokenTTL
	if refreshTTL == 0 {
		refreshTTL = 7 * 24 * time.Hour
	}

	overlap := config.OverlapWindow
	if overlap == 0 {
		overlap = 24 * time.Hour
	}

	issuer := config.Issuer
	if issuer == "" {
		issuer = "app"
	}

	if len(config.Secret) == 0 {
		return nil, errors.New("secret key cannot be empty")
	}

	return &JWTManager{
		secret:          config.Secret,
		accessTokenTTL:  accessTTL,
		refreshTokenTTL: refreshTTL,
		overlapWindow:   overlap,
		issuer:          issuer,
	}, nil
}

// GenerateTokenPair generates both access and refresh tokens for an account
func (jwtm *JWTManager) GenerateTokenPair(accountID uuid.UUID) (*TokenPair, error) {
	// Generate access token
	accessToken, accessExp, err := jwtm.GenerateAccessToken(accountID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	// Generate refresh token
	refreshToken, refreshExp, err := jwtm.GenerateRefreshToken(accountID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}

	return &TokenPair{
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessExp,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshExp,
	}, nil
}

// GenerateAccessToken generates an access token
func (jwtm *JWTManager) GenerateAccessToken(accountID uuid.UUID) (string, time.Time, error) {
	now := time.Now()
	exp := now.Add(jwtm.accessTokenTTL)

	claims := &AccessTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   accountID.String(),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(exp),
			Issuer:    jwtm.issuer,
			ID:        uuid.New().String(),
		},
		AccountID: accountID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtm.secret)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("failed to sign access token: %w", err)
	}

	return tokenString, exp, nil
}

// GenerateRefreshToken generates a refresh token
func (jwtm *JWTManager) GenerateRefreshToken(accountID uuid.UUID) (string, time.Time, error) {
	now := time.Now()
	exp := now.Add(jwtm.refreshTokenTTL)

	claims := &RefreshTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   accountID.String(),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(exp),
			Issuer:    jwtm.issuer,
			ID:        uuid.New().String(),
		},
		AccountID: accountID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtm.secret)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("failed to sign refresh token: %w", err)
	}

	return tokenString, exp, nil
}

// UpdateRefreshToken validates an old refresh token and generates a new one if within the overlap window
func (jwtm *JWTManager) UpdateRefreshToken(oldTokenString string) (string, time.Time, error) {
	// Parse and validate the old refresh token
	oldClaims, err := jwtm.ParseWithRefreshClaims(oldTokenString)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("invalid refresh token: %w", err)
	}

	// Check if we're in the refresh window (overlap period before expiration)
	now := time.Now()
	refreshWindowStart := oldClaims.ExpiresAt.Time.Add(-jwtm.overlapWindow)

	if now.Before(refreshWindowStart) {
		return "", time.Time{}, fmt.Errorf("%w: token can only be refreshed within %v of expiration",
			ErrTokenNotInWindow, jwtm.overlapWindow)
	}

	// Generate new refresh token
	newToken, newExp, err := jwtm.GenerateRefreshToken(oldClaims.AccountID)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("failed to generate new refresh token: %w", err)
	}

	return newToken, newExp, nil
}

// ValidateAccessToken checks if an access token is valid
func (jwtm *JWTManager) ValidateAccessToken(tokenString string) (bool, error) {
	_, err := jwtm.ParseWithAccessClaims(tokenString)
	if err != nil {
		return false, err
	}
	return true, nil
}

// ValidateRefreshToken checks if a refresh token is valid
func (jwtm *JWTManager) ValidateRefreshToken(tokenString string) (bool, error) {
	_, err := jwtm.ParseWithRefreshClaims(tokenString)
	if err != nil {
		return false, err
	}
	return true, nil
}

// ParseWithAccessClaims parses and validates an access token, returning its claims
func (jwtm *JWTManager) ParseWithAccessClaims(tokenString string) (*AccessTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Verify signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("%w: %v", ErrBadSign, token.Header["alg"])
		}
		return jwtm.secret, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrTokenExpired
		}
		return nil, fmt.Errorf("%w: %v", ErrInvalidToken, err)
	}

	claims, ok := token.Claims.(*AccessTokenClaims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}

// ParseWithRefreshClaims parses and validates a refresh token, returning its claims
func (jwtm *JWTManager) ParseWithRefreshClaims(tokenString string) (*RefreshTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &RefreshTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Verify signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("%w: %v", ErrBadSign, token.Header["alg"])
		}
		return jwtm.secret, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrTokenExpired
		}
		return nil, fmt.Errorf("%w: %v", ErrInvalidToken, err)
	}

	claims, ok := token.Claims.(*RefreshTokenClaims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}
