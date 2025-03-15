package model

import (
	"log"
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
	Password  string    `json:"-" gorm:"column:password;type:varchar;size:60;not null"`
	Email     string    `json:"email" gorm:"column:email;type:varchar;size:255;not null;uniqueIndex"`
	Avatar    string    `json:"avatar" gorm:"column:avatar;type:varchar;size:255"`
	Role      UserRole  `json:"role" gorm:"column:role;type:varchar;size:20;not null"`
	IsActive  bool      `json:"is_active" gorm:"column:is_active;type:boolean;default:true"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp"`

	// Associations
	Candidate *Candidate `json:"candidate,omitempty" gorm:"foreignKey:UserID"`
	Company   *Company   `json:"company,omitempty" gorm:"foreignKey:UserID"`
}

func (user *User) HashPassword(plainPassword string) string {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return ""
	}
	return string(hashedBytes)
}

func (user *User) CheckPassword(plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(plainPassword))
	return err == nil
}
