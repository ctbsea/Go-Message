package repositories

import (
	"github.com/ctbsea/Go-Message/config/db"
	"github.com/ctbsea/Go-Message/datamodels"
)

type Rep struct {
	UserRep UserRep
}

func InitRep(db *db.Db, model *datamodels.Models) *Rep {
	return &Rep{
		UserRep: NewUserRep(db, model.UserModel),
	}
}
