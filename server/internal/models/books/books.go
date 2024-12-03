package books

import (
	"time"
)

// type User struct {
// 	ID string `json:"id" gorm:"primarykey"`
// }

type Book struct {
	ID          string    `gorm:"primarykey"`
	IsActive    bool      `json:"isActive" gorm:"column:is_active"`
	IsDone      bool      `json:"isDone" gorm:"column:is_done"`
	CreatedAt   time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"column:updated_at"`
	FinishedAt  time.Time `json:"finishedAt" gorm:"column:finished_at"`
	Title       string    `json:"title" gorm:"column:title"`
	Author      string    `json:"author" gorm:"column:author"`
	TotalPages  uint      `json:"totalPages" gorm:"column:total_pages"`
	CurrentPage uint      `json:"currentPage" gorm:"column:current_page"`
	DailyGoal   uint      `json:"dailyGoal" gorm:"column:daily_goal"`
	PagesRead   uint      `json:"pagesRead" gorm:"column:pages_read"`
	// UserId          string    `json:"userId"  gorm:"column:user_id"`
	// User            User      `gorm:"foreignKey:UserId;references:ID"`
}
