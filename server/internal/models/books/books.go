package books

import (
	"time"
)

type Book struct {
	ID          string             `json:"id"`
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
}

type UserData struct {
	Books         []Book `json:"books"`
	CurrentBookId string `json:"currentBookId"`
}
