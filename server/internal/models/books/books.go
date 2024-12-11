package books

import (
	"time"
)

type User struct {
	ID int64 `json:"id"`
}

type Book struct {
	ID          string
	IsActive    bool               `json:"isActive"`
	IsDone      bool               `json:"isDone"`
	CreatedAt   time.Time          `json:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt"`
	FinishedAt  time.Time          `json:"finishedAt"`
	Title       string             `json:"title"`
	Author      string             `json:"author"`
	TotalPages  uint               `json:"totalPages"`
	CurrentPage uint               `json:"currentPage"`
	DailyGoal   uint               `json:"dailyGoal"`
	PagesRead   map[time.Time]uint `json:"pagesRead"`
	UserId      User               `json:"userId"`
}