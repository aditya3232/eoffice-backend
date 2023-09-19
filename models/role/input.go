package role

type CreateInput struct {
	Nama string `json:"nama" form:"nama" binding:"required"`
}

type UpdateInput struct {
	Nama string `json:"nama" form:"nama"`
}
