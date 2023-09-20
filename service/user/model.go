package user

// Param 接口绑定参数
type Param struct {
	Email    string `json:"Email"  binding:"required"`
	Password string `json:"Password"  binding:"required"`
}
