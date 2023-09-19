package auth

type LoginInput struct {
	Nip      int    `json:"nip" form:"nip" binding:"required,numeric"`
	Password string `json:"password" form:"password" binding:"required"`
}
