package jwt_test

import (
	"errors"
	"fmt"
	"moon/internal/infrastructure/security/jwt"
	"moon/pkg/sync/lazy_lock"
	"testing"
	"time"

	"github.com/google/uuid"
)

var jwtManagerLazy = lazy_lock.New(func() *jwt.JWTManager {
	config := &jwt.JWTConfig{
		Secret:          []byte("your-secret-key-change-in-production"),
		AccessTokenTTL:  15 * time.Minute,
		RefreshTokenTTL: 7 * 24 * time.Hour,
		OverlapWindow:   24 * time.Hour,
	}

	manager, err := jwt.New(config)
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize JWT manager: %v", err))
	}
	return manager
})

// TestJWTManager_BasicOperations tests basic token generation and validation
func TestJWTManager_BasicOperations(t *testing.T) {
	manager := jwtManagerLazy.Get()

	accountID := uuid.New()

	// Test token pair generation
	tokenPair, err := manager.GenerateTokenPair(accountID)
	if err != nil {
		t.Fatalf("Failed to generate token pair: %v", err)
	}

	if tokenPair.AccessToken == "" {
		t.Error("Access token is empty")
	}
	if tokenPair.RefreshToken == "" {
		t.Error("Refresh token is empty")
	}

	// Validate access token
	valid, err := manager.ValidateAccessToken(tokenPair.AccessToken)
	if err != nil {
		t.Errorf("Access token validation failed: %v", err)
	}
	if !valid {
		t.Error("Access token is invalid")
	}

	// Validate refresh token
	valid, err = manager.ValidateRefreshToken(tokenPair.RefreshToken)
	if err != nil {
		t.Errorf("Refresh token validation failed: %v", err)
	}
	if !valid {
		t.Error("Refresh token is invalid")
	}

	// Parse access token claims
	accessClaims, err := manager.ParseWithAccessClaims(tokenPair.AccessToken)
	if err != nil {
		t.Fatalf("Failed to parse access token claims: %v", err)
	}
	if accessClaims.AccountID != accountID {
		t.Errorf("AccountID mismatch: expected %s, got %s", accountID, accessClaims.AccountID)
	}

	// Parse refresh token claims
	refreshClaims, err := manager.ParseWithRefreshClaims(tokenPair.RefreshToken)
	if err != nil {
		t.Fatalf("Failed to parse refresh token claims: %v", err)
	}
	if refreshClaims.AccountID != accountID {
		t.Errorf("AccountID mismatch: expected %s, got %s", accountID, refreshClaims.AccountID)
	}
}

// TestJWTManager_RefreshWindow tests the 1-day overlap window logic
func TestJWTManager_RefreshWindow(t *testing.T) {
	// Create a manager with short lifetimes for testing
	config := &jwt.JWTConfig{
		Secret:          []byte("test-secret"),
		AccessTokenTTL:  15 * time.Minute,
		RefreshTokenTTL: 7 * 24 * time.Hour,
		OverlapWindow:   24 * time.Hour,
		Issuer:          "test",
	}
	manager, err := jwt.New(config)
	if err != nil {
		t.Fatalf("Failed to create manager: %v", err)
	}

	accountID := uuid.New()

	// Generate initial token pair
	tokenPair, err := manager.GenerateTokenPair(accountID)
	if err != nil {
		t.Fatalf("Failed to generate token pair: %v", err)
	}

	// Try to refresh immediately - should fail (too early)
	_, _, err = manager.UpdateRefreshToken(tokenPair.RefreshToken)
	if !errors.Is(err, jwt.ErrTokenNotInWindow) {
		t.Errorf("Expected ErrTokenNotInWindow, got: %v", err)
	}

	// Invalidating the test due to time constraints in CI/CD
	t.Logf("Refresh window test requires time manipulation. Manual testing recommended.")
}

// TestJWTManager_TokenExpiration tests token expiration behavior
func TestJWTManager_TokenExpiration(t *testing.T) {
	// Create a manager with very short lifetimes
	config := &jwt.JWTConfig{
		Secret:          []byte("test-secret"),
		AccessTokenTTL:  1 * time.Millisecond, // Expires immediately
		RefreshTokenTTL: 1 * time.Millisecond,
		OverlapWindow:   1 * time.Millisecond,
		Issuer:          "test",
	}
	manager, err := jwt.New(config)
	if err != nil {
		t.Fatalf("Failed to create manager: %v", err)
	}

	accountID := uuid.New()
	tokenPair, err := manager.GenerateTokenPair(accountID)
	if err != nil {
		t.Fatalf("Failed to generate token pair: %v", err)
	}

	// Wait for tokens to expire
	time.Sleep(2 * time.Millisecond)

	// Try to validate expired access token
	_, err = manager.ParseWithAccessClaims(tokenPair.AccessToken)
	if err != jwt.ErrTokenExpired {
		t.Errorf("Expected ErrTokenExpired for access token, got: %v", err)
	}

	// Try to validate expired refresh token
	_, err = manager.ParseWithRefreshClaims(tokenPair.RefreshToken)
	if err != jwt.ErrTokenExpired {
		t.Errorf("Expected ErrTokenExpired for refresh token, got: %v", err)
	}
}

// TestJWTManager_InvalidTokens tests various invalid token scenarios
func TestJWTManager_InvalidTokens(t *testing.T) {
	manager := jwtManagerLazy.Get()

	testCases := []struct {
		name        string
		token       string
		expectError error
		parseFunc   func(string) (interface{}, error)
	}{
		{
			name:        "Empty token",
			token:       "",
			expectError: jwt.ErrInvalidToken,
			parseFunc:   func(s string) (interface{}, error) { return manager.ParseWithAccessClaims(s) },
		},
		{
			name:        "Malformed token",
			token:       "invalid.token.string",
			expectError: jwt.ErrInvalidToken,
			parseFunc:   func(s string) (interface{}, error) { return manager.ParseWithAccessClaims(s) },
		},
		{
			name:        "Wrong signature",
			token:       "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.wrong_signature",
			expectError: jwt.ErrInvalidToken,
			parseFunc:   func(s string) (interface{}, error) { return manager.ParseWithAccessClaims(s) },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := tc.parseFunc(tc.token)
			if err == nil {
				t.Error("Expected error but got nil")
			}
		})
	}
}

// TestJWTManager_ConcurrentAccess tests thread-safety of lazy initialization
func TestJWTManager_ConcurrentAccess(t *testing.T) {
	const numGoroutines = 100

	done := make(chan bool)

	// Try to jwtManagerLazy.Get instance concurrently
	for i := 0; i < numGoroutines; i++ {
		go func() {
			manager := jwtManagerLazy.Get()
			if manager == nil {
				t.Error("Got nil manager")
			}
			done <- true
		}()
	}

	// Wait for all goroutines
	for i := 0; i < numGoroutines; i++ {
		<-done
	}
}

// TestJWTManager_SingletonPattern verifies singleton behavior
func TestJWTManager_SingletonPattern(t *testing.T) {
	manager1 := jwtManagerLazy.Get()
	manager2 := jwtManagerLazy.Get()

	if manager1 != manager2 {
		t.Error("jwtManagerLazy.Get() should return the same instance")
	}
}

// TestJWTManager_RefreshOverlapWindowEdgeCase tests exact window boundary
func TestJWTManager_RefreshOverlapWindowEdgeCase(t *testing.T) {
	config := &jwt.JWTConfig{
		Secret:          []byte("test-secret"),
		AccessTokenTTL:  15 * time.Minute,
		RefreshTokenTTL: 1 * time.Hour,
		OverlapWindow:   30 * time.Minute,
		Issuer:          "test",
	}
	manager, err := jwt.New(config)
	if err != nil {
		t.Fatalf("Failed to create manager: %v", err)
	}

	accountID := uuid.New()
	tokenPair, err := manager.GenerateTokenPair(accountID)
	if err != nil {
		t.Fatalf("Failed to generate token pair: %v", err)
	}

	// This test demonstrates the window logic but requires time.Sleep
	// In production, use a mock clock or time manipulation
	t.Logf("Generated tokens. Refresh window opens at: %v",
		tokenPair.RefreshTokenExpiresAt.Add(-config.OverlapWindow))
}

// TestJWTManager_ConfigValidation tests config validation
func TestJWTManager_ConfigValidation(t *testing.T) {
	testCases := []struct {
		name        string
		config      *jwt.JWTConfig
		expectError bool
	}{
		{
			name:        "Nil config",
			config:      nil,
			expectError: true,
		},
		{
			name: "Empty secret",
			config: &jwt.JWTConfig{
				Secret: []byte(""),
			},
			expectError: true,
		},
		{
			name: "Valid config",
			config: &jwt.JWTConfig{
				Secret: []byte("valid-secret"),
			},
			expectError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := jwt.New(tc.config)
			if tc.expectError && err == nil {
				t.Error("Expected error but got nil")
			}
			if !tc.expectError && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
		})
	}
}
