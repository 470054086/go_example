package model

import (
	"mall/dao"
)
// user的数据类型
type User struct {
	ID        uint `gorm:"primary_key"`
	Mobile string
	Password string
	Sex int
	//RegisterTime     *time.Time
}

/**
增加数据的操作
 */
func (u *User) Create(user *User)(uint,error) {
	create := dao.G_db.Create(user)
	if err := create.Error; err!= nil {
		return 0,err
	}
	if create.RowsAffected > 0  {
		return user.ID,nil
	}
	return 0,create.Error
}