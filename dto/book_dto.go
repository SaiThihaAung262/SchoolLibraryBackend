package dto

type CreateBookDTO struct {
	Title         string `json:"title" form:"title" binding:"required"`
	SubCategoryID uint64 `json:"sub_category_id" form:"sub_category_id" binding:"required"`
	Status        uint64 `json:"status" form:"status" binding:"required"`
	Author        string `json:"author" form:"author" binding:"required"`
	Summary       string `json:"summary" from:"summary" binding:"required"`
	BookImage     string `json:"book_image" form:"book_image"`
	AvailableQty  uint64 `json:"available_qty" form:"available_qty" binding:"required"`
	PublishDate   string `json:"publish_date" form:"publish_date"`
	DownloadLink  string `json:"download_link" form:"download_link"`
}

type BookGetRequest struct {
	ID           uint64 `json:"id" form:"id"`
	UUID         string `json:"uuid" form:"uuid"`
	Title        string `json:"title" form:"title"`
	SubCategorID uint64 `json:"sub_category_id" form:"category_id"`
	Status       uint64 `json:"status" form:"status"`
	Author       string `json:"author" form:"author"`
	Page         uint64 `json:"page" form:"page"`
	PageSize     uint64 `json:"page_size" form:"page_size"`
}

type UpdateBookDTO struct {
	ID           uint64 `json:"id" form:"id" binding:"required"`
	Title        string `json:"title" form:"title"`
	SubCategorID uint64 `json:"sub_category_id" form:"sub_category_id"`
	Status       uint64 `json:"status" form:"status"`
	Author       string `json:"author" form:"author"`
	Summary      string `json:"summary" from:"summary"`
	BookImage    string `json:"book_image" form:"book_image"`
	AvailableQty uint64 `json:"available_qty" form:"available_qty"`
	BorrowQty    uint64 `json:"borrow_qty" form:"borrow_qty"`
	PublishDate  string `json:"publish_date" form:"publish_date"`
	DownloadLink string `json:"download_link" form:"download_link"`
}

type GetBookByUUIDDto struct {
	UUID string `json:"uuid" form:"uuid" binding:"required"`
}
