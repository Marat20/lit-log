package books

import (
	"database/sql"
	"time"
)

type User struct {
	ID string `json:"id" gorm:"primarykey"`
}

type Book struct {
	ID         string       `gorm:"primarykey"`
	CreatedAt  time.Time    `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt  time.Time    `json:"updatedAt" gorm:"column:updated_at"`
	DeletedAt  sql.NullTime `json:"deletedAt" gorm:"index;column:deleted_at"`
	Title      string       `json:"title" gorm:"column:title"`
	Author     string       `json:"author" gorm:"column:author"`
	TotalPages uint         `json:"totalPages" gorm:"column:total_pages"`
}

type ReadingProgress struct {
	ID          string       `gorm:"primarykey"`
	CreatedAt   time.Time    `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt   time.Time    `json:"updatedAt" gorm:"column:updated_at"`
	DeletedAt   sql.NullTime `json:"deletedAt" gorm:"index;column:deleted_at"`
	UserId      string       `json:"userId"  gorm:"column:user_id"`
	BookId      string       `json:"bookId"  gorm:"column:book_id"`
	CurrentPage uint         `json:"currentPage" gorm:"column:current_page"`
	DailyGoal   uint         `json:"dailyGoal" gorm:"column:daily_goal"`
	User        User         `gorm:"foreignKey:UserId;references:ID"`
	Book        Book         `gorm:"foreignKey:BookId;references:ID"`
}
