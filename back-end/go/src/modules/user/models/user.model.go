package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"index:,sort:desc,type:btree,unique"`
	Password string
	Active   sql.NullBool `gorm:"default:true"`
}
