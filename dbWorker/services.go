package dbWorker

import (
	"github.com/jinzhu/gorm"
)

func DbListener(db *gorm.DB) {
	user := User{}
	transaction := db.Begin()
	transaction.First(&user, 1)
	transaction.Model(&user).Update("counter", user.Counter+1)
	transaction.Commit()
}

type User struct {
	Id      int `gorm:"primary_key"`
	Rating  int
	Counter int
}

func (User) TableName() string {
	return "Users"
}
