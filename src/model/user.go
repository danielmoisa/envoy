package model

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/danielmoisa/envoy/src/utils/idconvertor"
	"github.com/google/uuid"
)

const AnonymousUserId = -1

type RawUser struct {
	ID             string    `json:"id" gorm:"column:id;type:bigint;primary_key;index:users_ukey"`
	UID            uuid.UUID `json:"uid" gorm:"column:uid;type:uuid;not null;index:users_ukey"`
	Nickname       string    `json:"nickname" gorm:"column:nickname;type:varchar;size:15"`
	PasswordDigest string    `json:"passworddigest" gorm:"column:password_digest;type:varchar;size:60;not null"`
	Email          string    `json:"email" gorm:"column:email;type:varchar;size:255;not null"`
	Avatar         string    `json:"avatar" gorm:"column:avatar;type:varchar;size:255;not null"`
	SSOConfig      string    `json:"SSOConfig" gorm:"column:sso_config;type:jsonb"`        // for single sign-on data
	Customization  string    `json:"customization" gorm:"column:customization;type:jsonb"` // for user itself customization config, including: Language, IsSubscribed
	TeamID         int       `json:"teamID" gorm:"column:team_id;type:bigint"`
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

type UserForEditedBy struct {
	ID       string    `json:"userID"`
	Nickname string    `json:"nickname"`
	Email    string    `json:"email"`
	Avatar   string    `json:"avatar"`
	EditedAt time.Time `json:"editedAt"`
}

type UserForModifiedBy struct {
	ID       string `json:"userID"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
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

func NewInvalidatedUser() *User {
	return &User{
		ID:       AnonymousUserId,
		Nickname: "invalidated",
		Email:    "invalidated",
		Avatar:   "invalidated",
	}
}

func NewUserForEditedBy(user *User, editedAt time.Time) *UserForEditedBy {
	return &UserForEditedBy{
		ID:       idconvertor.ConvertIntToString(user.ID),
		Nickname: user.Nickname,
		Email:    user.Email,
		Avatar:   user.Avatar,
		EditedAt: editedAt,
	}
}

func NewUserForModifiedBy(user *User) *UserForModifiedBy {
	return &UserForModifiedBy{
		ID:       idconvertor.ConvertIntToString(user.ID),
		Nickname: user.Nickname,
		Email:    user.Email,
		Avatar:   user.Avatar,
	}
}

func NewUserByDataControlRawData(rawUserString string) (*User, error) {
	rawUser := RawUser{}
	errInUnmarshal := json.Unmarshal([]byte(rawUserString), &rawUser)
	if errInUnmarshal != nil {
		return nil, errInUnmarshal
	}
	return NewUser(&rawUser), nil
}

func NewUsersByDataControlRawData(rawUsersString string) (map[int]*User, error) {
	RawUsers := RawUsers{}
	errInUnmarshal := json.Unmarshal([]byte(rawUsersString), &RawUsers)
	if errInUnmarshal != nil {
		return nil, errInUnmarshal
	}
	usersRet := make(map[int]*User, len(RawUsers.Users))
	for userIDString, rawUser := range RawUsers.Users {
		userIDInt, _ := strconv.Atoi(userIDString)
		usersRet[userIDInt] = NewUser(rawUser)
	}
	return usersRet, nil
}

func (u *User) ExportIDToString() string {
	return idconvertor.ConvertIntToString(u.ID)
}
