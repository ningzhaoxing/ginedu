package request

type Login struct {
	Name     string `json:"name" form:"name" binding:"required,demo=xxx"`
	Password string `json:"password" form:"password" binding:"required"`
}
