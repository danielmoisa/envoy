package model

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserRole string

const (
	RoleCandidate UserRole = "candidate"
	RoleCompany   UserRole = "company"
	RoleAdmin     UserRole = "admin"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	Username  string    `json:"username" gorm:"column:username;type:varchar;size:15;uniqueIndex"`
	Password  string    `json:"-" gorm:"column:password;type:varchar(60);not null"`
	Email     string    `json:"email" gorm:"column:email;type:varchar;size:255;not null;uniqueIndex"`
	Avatar    string    `json:"avatar" gorm:"column:avatar;type:varchar;size:255"`
	Role      UserRole  `json:"role" gorm:"column:role;type:varchar;size:20;not null"`
	IsActive  bool      `json:"is_active" gorm:"column:is_active;type:boolean;default:true"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:timestamp"`
	UpdatedAt time.Time `json:"update_at" gorm:"column:updated_at;type:timestamp"`

	// Associations
	Candidate *Candidate `json:"candidate,omitempty" gorm:"foreignKey:UserID"`
	Company   *Company   `json:"company,omitempty" gorm:"foreignKey:UserID"`
}

func (user *User) HashPassword(plainPassword string) (string, error) {
	if plainPassword == "" {
		return "", errors.New("password cannot be empty")
	}

	// Trim any whitespace to prevent accidental spaces
	plainPassword = strings.TrimSpace(plainPassword)

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(plainPassword), 14)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	return string(hashedBytes), nil
}

func (user *User) CheckPassword(plainPassword string) bool {
	plainPassword = strings.TrimSpace(plainPassword)

	if plainPassword == "" {
		log.Printf("Error: Plain password is empty")
		return false
	}
	if user.Password == "" {
		log.Printf("Error: Stored hash is empty")
		return false
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(plainPassword))
	if err != nil {
		log.Printf("3. Bcrypt comparison error: %v", err)

		return false
	}

	return true
}
