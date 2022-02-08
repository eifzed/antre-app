package auth

import (
	"time"

	"github.com/eifzed/antre-app/lib/common/commonerr"
	"github.com/eifzed/antre-app/lib/utility/jwt"
)

type UserDetail struct {
	UserID         int64      `json:"-" xorm:"user_id pk autoincr"`
	Name           string     `json:"name" xorm:"name"`
	PhoneNumber    string     `json:"phone_number" xorm:"phone_number"`
	Email          string     `json:"email" xorm:"email"`
	Username       string     `json:"username" xorm:"username"`
	Address        string     `json:"address" xorm:"address"`
	Password       string     `json:"password" xorm:"-"`
	PasswordHashed string     `json:"-" xorm:"password"`
	Roles          []jwt.Role `json:"user_roles" xorm:"-"`
	PostalCode     int64      `json:"postal_code" xorm:"postal_code"`
	CreateTime     *time.Time `xorm:"create_time created"`
	UpdateTime     *time.Time `xorm:"update_time updated"`
	DeleteTime     time.Time  `xorm:"delete_time deleted"`
}

func (user *UserDetail) ValidateRegistrationParam() error {
	if user.Password == "" || user.Name == "" || user.PhoneNumber == "" || user.Email == "" || user.Username == "" {
		return commonerr.ErrorBadRequest("Register Account", "Name, Username, Phone Number, Email, and Password cannot be empty")
	}
	return nil
}

type MapUserRole struct {
	MapID  int64 `xorm:"id pk autoincr"`
	UserID int64 `xorm:"user_id"`
	RoleID int64 `xorm:"role_id"`
}

type LoginParams struct {
	Email          string `json:"email"`
	Password       string `json:"password"`
	PasswordHashed string
}
