package model

import (
	"mall/dao"
)

// user的数据类型
type User struct {
	ID       uint `gorm:"primary_key"`
	Mobile   string
	Password string `gorm:"column:Password"`
	Sex      int
	//RegisterTime     *time.Time
}

/**
增加数据的操作
*/
func (u *User) Create(user *User) (uint, error) {
	create := dao.G_db.Create(user)
	if err := create.Error; err != nil {
		return 0, err
	}
	if create.RowsAffected > 0 {
		return user.ID, nil
	}
	return 0, create.Error
}

func (u *User) GetFirst(mobile string) *User {
	var user User
	found := dao.G_db.Where("mobile = ?", mobile).First(&user).RecordNotFound()
	if found {
		return nil
	}
	return &user
}

func (u *User) GetList(mobile string, sex int) []*User {
	var db = dao.G_db
	if mobile != "" {
		db = db.Where("mobile=?", mobile)
	}
	if sex != 0 {
		db = db.Where("sex = ?", sex)
	}
	var user []*User
	db.Find(&user)
	return user
}

func (u *User) GetListCunt(mobile string, sex int) int {
	var db = dao.G_db
	db = db.Model(&User{})
	if mobile != "" {
		db = db.Where("mobile=?", mobile)
	}
	if sex != 0 {
		db = db.Where("sex = ?", sex)
	}
	var count int
	db.Count(&count)
	return count
}
