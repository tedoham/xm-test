package domain

import (
	"context"
	"time"
)

type User struct {
	ID         string    `json:"id" db:"id"`
	Email      string    `json:"email" validate:"required" db:"email"`
	Password   string    `json:"password" validate:"required" db:"password"`
	Username   string    `json:"username" db:"username"`
	TokenHash  string    `json:"tokenhash" db:"tokenhash"`
	IsVerified bool      `json:"isverified" db:"isverified"`
	CreatedAt  time.Time `json:"createdat" db:"createdat"`
	UpdatedAt  time.Time `json:"updatedat" db:"updatedat"`
}

type VerificationDataType int

const (
	MailConfirmation VerificationDataType = iota + 1
	PassReset
)

// VerificationData represents the type for the data stored for verification.
type VerificationData struct {
	Email     string               `json:"email" validate:"required" db:"email"`
	Code      string               `json:"code" validate:"required" db:"code"`
	ExpiresAt time.Time            `json:"expiresat" db:"expiresat"`
	Type      VerificationDataType `json:"type" db:"type"`
}

// Authentication interface lists the methods that our authentication service should implement
type Authentication interface {
	Authenticate(reqUser *User, user *User) bool
	GenerateAccessToken(user *User) (string, error)
	GenerateRefreshToken(user *User) (string, error)
	GenerateCustomKey(userID string, password string) string
	ValidateAccessToken(token string) (string, error)
	ValidateRefreshToken(token string) (string, string, error)
}

// Repository is an interface for the storage implementation of the auth service
type Repository interface {
	Create(ctx context.Context, user *User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByID(ctx context.Context, userID string) (*User, error)
	UpdateUsername(ctx context.Context, user *User) error
	StoreVerificationData(ctx context.Context, verificationData *VerificationData) error
	GetVerificationData(ctx context.Context, email string, verificationDataType VerificationDataType) (*VerificationData, error)
	UpdateUserVerificationStatus(ctx context.Context, email string, status bool) error
	DeleteVerificationData(ctx context.Context, email string, verificationDataType VerificationDataType) error
	UpdatePassword(ctx context.Context, userID string, password string, tokenHash string) error
}