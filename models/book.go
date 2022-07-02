package models

type (
	Book struct {
		ID     int    `gorm:"column:id"`
		Title  string `gorm:"column:title"`
		Author string `gorm:"column:author"`
		Year   string `gorm:"column:year"`
	}

	CreateBook struct {
		Title  string `json:"title" validate:"required"`
		Author string `json:"author" validate:"required"`
		Year   string `json:"year" validate:"required"`
	}

	UpdateBook struct {
		ID     int
		Author string `json:"author" validate:"required"`
		Year   string `json:"year"`
	}
)
