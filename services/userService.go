package services

import (
	"github.com/ctbsea/Go-Message/entry"
	"github.com/ctbsea/Go-Message/repositories"
	"github.com/ctbsea/Go-Message/util"
	"github.com/ctbsea/Go-Message/util/jwtlogin"
	"time"
)

type LoginRep struct {
	Token string
	Id    uint
}

type RegRep struct {
	Id uint
}

type UserService interface {
	Login(where map[string]string) (data LoginRep, code int)
	Register(where map[string]string) (data RegRep, code int)
}

func NewUserService(rep *repositories.Rep) UserService {
	return &userService{Rep: rep}
}

type userService struct {
	Rep *repositories.Rep
}

func (r *userService) Login(params map[string]string) (data LoginRep, code int) {
	where := make(map[string]interface{})
	where["user_name"] = params["user_name"]
	userInfo, found := r.Rep.UserRep.Select(where)
	if !found {
		return LoginRep{}, entry.NO_FOUND_USER
	}
	if userInfo.Pass != util.GetPass(params["user_pass"]) {
		return LoginRep{}, entry.ERROR_PASS
	}
	//update
	updateData := make(map[string]interface{})
	updateData["LoginAt"] = time.Now()
	updateData["LoginIp"] = util.InetAtoN(params["login_ip"])
	r.Rep.UserRep.Update(where, updateData)
	//token生成
	token := jwtlogin.Sign(&jwtlogin.ClaimsInfo{UserID: userInfo.ID, Username: userInfo.UserName})
	return LoginRep{Token: token, Id: userInfo.ID}, entry.SUCCESS
}

func (r *userService) Register(params map[string]string) (data RegRep, code int) {
	where := make(map[string]interface{})
	where["user_name"] = params["user_name"]
	_, found := r.Rep.UserRep.Select(where)
	if found {
		return RegRep{}, entry.REGISTER_FOUND_USER
	}
	uid := r.Rep.UserRep.InsertGetId(params)
	return RegRep{Id: uid}, entry.SUCCESS
}
