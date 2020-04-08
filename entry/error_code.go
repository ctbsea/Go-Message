package entry

const (
	SUCCESS       = 0  //成功
	INVAILD_PARAM = -1 //参数不合法

	NO_FOUND_USER = 10001 //未找到该用户
	ERROR_PASS    = 10002
	REGISTER_FOUND_USER    = 10003

	JWT_ERR_TOKEN = 20001 //无效的jwt token
	JWT_EXP_TOKEN = 20002 //token 过期
)
