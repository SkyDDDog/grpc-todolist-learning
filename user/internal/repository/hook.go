package repository

import "gorm.io/gorm"

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UserId = uint(GenID())
	return nil
}
