package models

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// CheckAuth
// @Desc: 	验证对应的用户名密码是否存在
// @Param:	username
// @Param:	password
// @Return:	bool
// @Notice:
func CheckAuth(username, password string) bool {
	var auth Auth
	db.Select("id").Where(Auth{
		Username: username,
		Password: password,
	}).First(&auth)

	if auth.ID > 0 {
		return true
	}
	return false
}
