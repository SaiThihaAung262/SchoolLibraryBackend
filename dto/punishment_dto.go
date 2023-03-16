package dto

type CreatePunishmentDto struct {
	PackageName string `json:"package_name" form:"package_name" binding:"required"`
	Duration    uint64 `json:"duration" form:"duration" binding:"required"`
}

type UpdatePunishmentDto struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	PackageName string `json:"package_name" form:"package_name"`
	Duration    uint64 `json:"duration" form:"duration"`
}
