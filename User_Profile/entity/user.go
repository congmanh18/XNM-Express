package entity

import "github.com/congmanh18/lucas-core/record"

type User struct {
	record.BaseEntity
	Username     *string
	PasswordHash *string
	Email        *string
	Phone        *string
	Role         *string
}

func (u User) TableName() string {
	return "user_profile_service.users"
}
