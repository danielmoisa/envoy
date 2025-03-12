package model

import (
	"time"

	"github.com/danielmoisa/envoy/src/utils/idconvertor"
	"github.com/google/uuid"
)

type RawUser struct {
	ID             string    `json:"id" gorm:"column:id;type:bigint;primary_key;index:users_ukey"`
	UID            uuid.UUID `json:"uid" gorm:"column:uid;type:uuid;not null;index:users_ukey"`
	Nickname       string    `json:"nickname" gorm:"column:nickname;type:varchar;size:15"`
	PasswordDigest string    `json:"passworddigest" gorm:"column:password_digest;type:varchar;size:60;not null"`
	Email          string    `json:"email" gorm:"column:email;type:varchar;size:255;not null"`
	Avatar         string    `json:"avatar" gorm:"column:avatar;type:varchar;size:255;not null"`
	SSOConfig      string    `json:"SSOConfig" gorm:"column:sso_config;type:jsonb"`        // for single sign-on data
	Customization  string    `json:"customization" gorm:"column:customization;type:jsonb"` // for user itself customization config, including: Language, IsSubscribed
	CreatedAt      time.Time `gorm:"column:created_at;type:timestamp"`
	UpdatedAt      time.Time `gorm:"column:updated_at;type:timestamp"`
}

type RawUsers struct {
	Users map[string]*RawUser `json:"users"`
}

type User struct {
	ID             int       `json:"id" gorm:"column:id;type:bigint;primary_key;index:users_ukey"`
	UID            uuid.UUID `json:"uid" gorm:"column:uid;type:uuid;not null;index:users_ukey"`
	Nickname       string    `json:"nickname" gorm:"column:nickname;type:varchar;size:15"`
	PasswordDigest string    `json:"passworddigest" gorm:"column:password_digest;type:varchar;size:60;not null"`
	Email          string    `json:"email" gorm:"column:email;type:varchar;size:255;not null"`
	Avatar         string    `json:"avatar" gorm:"column:avatar;type:varchar;size:255;not null"`
	CreatedAt      time.Time `gorm:"column:created_at;type:timestamp"`
	UpdatedAt      time.Time `gorm:"column:updated_at;type:timestamp"`
}

func NewUser(u *RawUser) *User {
	return &User{
		ID:             idconvertor.ConvertStringToInt(u.ID),
		UID:            u.UID,
		Nickname:       u.Nickname,
		PasswordDigest: u.PasswordDigest,
		Email:          u.Email,
		Avatar:         u.Avatar,
		CreatedAt:      u.CreatedAt,
		UpdatedAt:      u.UpdatedAt,
	}
}

func (u *User) ExportIDToString() string {
	return idconvertor.ConvertIntToString(u.ID)
}
