package antre

import (
	"time"

	"github.com/eifzed/antre-app/lib/utility/jwt"
)

type MstRole struct {
	ID         int64      `xorm:"role_id pk autoincr"`
	Name       string     `xorm:"role_name"`
	CreateTime time.Time  `xorm:"create_time created"`
	UpdateTime time.Time  `xorm:"create_time updated"`
	DeleteTime *time.Time `xorm:"create_time deleted"`
}

type RegistrationResponse struct {
	Name       string     `json:"name"`
	Username   string     `json:"username"`
	Email      string     `json:"email"`
	Token      string     `json:"token"`
	Roles      []jwt.Role `json:"roles"`
	ValidUntil time.Time  `json:"valid_until"`
}
