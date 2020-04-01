package datamodels

type Models struct {
	UserModel *User
}

func InitModels() *Models {
	return &Models{
		UserModel: &User{},
	}
}
