package permission

type CreateInput struct {
	Nama     string `json:"nama" form:"nama" binding:"required"`
	ParentID int    `json:"parent_id" form:"parent_id"`
	Url      string `json:"url" form:"url" binding:"required"`
	Position int    `json:"position" form:"position"`
}

type UpdateInput struct {
	Nama     string `json:"nama" form:"nama"`
	ParentID int    `json:"parent_id" form:"parent_id"`
	Url      string `json:"url" form:"url"`
	Position int    `json:"position" form:"position"`
}
