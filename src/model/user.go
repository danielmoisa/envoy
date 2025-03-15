package model

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	Username  string    `json:"username" gorm:"column:username;type:varchar;size:15;uniqueIndex"`
	Password  string    `json:"-" gorm:"column:password;type:varchar;size:60;not null;default:null"`
	Email     string    `json:"email" gorm:"column:email;type:varchar;size:255;not null;default:null;uniqueIndex"`
	Avatar    string    `json:"avatar" gorm:"column:avatar;type:varchar;size:255"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp"`
}

// HashPassword hashes a plain text password and returns it as a string
func (u *User) HashPassword(plainPassword string) string {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return ""
	}
	return string(hashedBytes)
}

// CheckPassword verifies if the provided password matches the stored hash
func (u *User) CheckPassword(plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainPassword))
	return err == nil
}
