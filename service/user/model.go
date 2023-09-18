package user

// Param 接口绑定参数
type Param struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"Password" json:"Password" binding:"required,min=4,max=64"`
}
