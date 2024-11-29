package books

import "time"

type User struct {
	Id string `json:"id" gorm:"primary_key"`
}

type Book struct {
	Id         string    `json:"id" gorm:"primary_key"`
	Title      string    `json:"title"`
	Author     string    `json:"author"`
	TotalPages int       `json:"totalPages"`
	CreatedAt  time.Time `json:"createdAt"`
}

type ReadingProgress struct {
	Id          string    `json:"id" gorm:"primary_key"`
	User_id     string    `json:"userId"`
	Book_id     string    `json:"bookId"`
	CurrentPage int       `json:"currentPage"`
	DailyGoal   int       `json:"dailyGoal"`
	StartedAt   time.Time `json:"startedAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
