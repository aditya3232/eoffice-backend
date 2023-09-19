package employee

type CreateInput struct {
	Nama           string `json:"nama" form:"nama" binding:"required" validate:"required"`
	Nip            int    `json:"nip" form:"nip" binding:"required" validate:"required"`
	TempatLahir    string `json:"tempat_lahir" form:"tempat_lahir" binding:"required" validate:"required"`
	TanggalLahir   string `json:"tanggal_lahir" form:"tanggal_lahir" binding:"required" validate:"required"`
	Alamat         string `json:"alamat" form:"alamat" binding:"required" validate:"required"`
	NoHp           string `json:"no_hp" form:"no_hp" binding:"required" validate:"required"`
	EmailPersonal  string `json:"email_personal" form:"email_personal" binding:"required" validate:"required"`
	EmailCorporate string `json:"email_corporate" form:"email_corporate"`
	DivisionID     int    `json:"division_id" form:"division_id" binding:"required" validate:"required"`
	PositionID     int    `json:"position_id" form:"position_id" binding:"required" validate:"required"`
	StartDate      string `json:"start_date" form:"start_date" binding:"required" validate:"required"`
	EndDate        string `json:"end_date" form:"end_date"`
	Avatar         string `json:"avatar" form:"avatar"`
}

type UpdateInput struct {
	Nama           string `json:"nama" form:"nama"`
	Nip            int    `json:"nip" form:"nip"`
	TempatLahir    string `json:"tempat_lahir" form:"tempat_lahir"`
	TanggalLahir   string `json:"tanggal_lahir" form:"tanggal_lahir"`
	Alamat         string `json:"alamat" form:"alamat"`
	NoHp           string `json:"no_hp" form:"no_hp"`
	EmailPersonal  string `json:"email_personal" form:"email_personal"`
	EmailCorporate string `json:"email_corporate" form:"email_corporate"`
	DivisionID     int    `json:"division_id" form:"division_id"`
	PositionID     int    `json:"position_id" form:"position_id"`
	StartDate      string `json:"start_date" form:"start_date"`
	EndDate        string `json:"end_date" form:"end_date"`
	Avatar         string `json:"avatar" form:"avatar"`
}
