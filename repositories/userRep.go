package repositories

import (
	"fmt"
	"github.com/ctbsea/Go-Message/config/db"
	"github.com/ctbsea/Go-Message/datamodels"
)

type UserRep interface {
	Select(where map[string]interface{}) (data datamodels.User, found bool)
	Update(where map[string]interface{}, update map[string]interface{}) bool
	InsertGetId(data map[string]string) uint
}

type userRep struct {
	Db        *db.Db
	UserModel *datamodels.User
}

func NewUserRep(db *db.Db, userModel *datamodels.User) UserRep {
	return &userRep{Db: db, UserModel: userModel}
}

func (r *userRep) Select(where map[string]interface{}) (data datamodels.User, found bool) {
	db := r.Db.Mysql.Model(r.UserModel).Where(where).Find(&data)
	if db.RecordNotFound() {
		return datamodels.User{} , false
	}
	return data, true
}

func (r *userRep) Update(where map[string]interface{}, update map[string]interface{}) bool {
	db := r.Db.Mysql.Model(r.UserModel).Omit("Id", "UserName").Where(where).Update(update)
	if db.RowsAffected >= 0 {
		return true
	}
	return false
}

func (r *userRep) InsertGetId(data map[string]string) uint {
	params := &datamodels.User{
		UserName: data["user_name"] ,
		Pass: data["user_pass"],
	}
	db := r.Db.Mysql.Create(params)
	if db.Error != nil {
		fmt.Println(db.Error.Error())
	}
	return params.ID
}
