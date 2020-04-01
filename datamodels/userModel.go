package datamodels

import (
	"github.com/ctbsea/Go-Message/util"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	UserName  string     `gorm:"column:user_name;size:20;not null;unique_index"` // string默认长度为255。
	Pass      string     `gorm:"column:user_pass;size:50;not null"`
	LoginIp   string     `gorm:"column:login_ip`
	LoginAt   time.Time  `gorm:"column:login_at"`
}

func (user *User) TableName() string {
	return "user_info"
}


func (user *User) BeforeSave(scope *gorm.Scope) {
	scope.SetColumn("Pass", util.GetPass(user.Pass))
	scope.SetColumn("LoginAt", user.LoginAt.Format("2006-01-02 15:04:05"))
	scope.SetColumn("LoginIp", util.InetAtoN(user.LoginIp))
}


