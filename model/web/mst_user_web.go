package web

type LoginRequest struct {
	Username string `json:"username" xml:"username" form:"username" validate:"required"`
	Password string `json:"password" xml:"password" form:"password" validate:"required"`
}
