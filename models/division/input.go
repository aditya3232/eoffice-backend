package division

type CreateInput struct {
	Nama     string `json:"nama" form:"nama" binding:"required"`
	ParentID int    `json:"parent_id" form:"parent_id"`
}

type UpdateInput struct {
	Nama     string `json:"nama" form:"nama"`
	ParentID int    `json:"parent_id" form:"parent_id"`
}
