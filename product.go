package jewerly

// Products
type CreateProductInput struct {
	Titles        MultiLanguageInput `json:"titles" binding:"required"`
	Descriptions  MultiLanguageInput `json:"descriptions" binding:"required"`
	CurrentPrice  float32            `json:"current_price" binding:"required"`
	PreviousPrice float32            `json:"previous_price"`
	Code          string             `json:"code" binding:"required"`
	ImageIds      []int              `json:"image_ids" binding:"required"`
	CategoryId    int                `json:"category_id" binding:"required"`
}

type MultiLanguageInput struct {
	English   string `json:"english" binding:"required"`
	Russian   string `json:"russian" binding:"required"`
	Ukrainian string `json:"ukrainian" binding:"required"`
}

// Categories

type Category int

const (
	CategoryRings = iota + 1
	CategoryBracelets
	CategoryPendants
	CategoryEarring
	CategoryNecklaces
)