package book

import (
	"encoding/json"
	"lit-log/internal/models/books"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	bolt "go.etcd.io/bbolt"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

type createBookInput struct {
	Title      string `json:"title" binding:"required"`
	Author     string `json:"author" binding:"required"`
	TotalPages uint   `json:"totalPages" binding:"required"`
	DailyGoal  uint   `json:"dailyGoal" binding:"required"`
}

type updateBookInput struct {
	PagesRead uint `json:"pagesRead" binding:"required"`
}

func (h handler) getBook(context *gin.Context) {
	id := context.Param("id")

	var book books.Book

	err := h.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Books"))
		if b == nil {
			return bolt.ErrBucketNotFound
		}

		bookData := b.Get([]byte(id))
		if bookData == nil {
			return bolt.ErrBucketNotFound
		}

		if err := json.Unmarshal(bookData, &book); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	pagesReadToday := CountPagesReadToday(book.PagesRead)

	context.JSON(http.StatusOK, gin.H{"book": book, "pagesReadToday": pagesReadToday})
}

func (h handler) getAllBooks(context *gin.Context) {
	var allBooks []books.Book

	err := h.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Books"))
		if b == nil {
			return bolt.ErrBucketNotFound
		}

		return b.ForEach(func(k, v []byte) error {
			var book books.Book
			if err := json.Unmarshal(v, &book); err != nil {
				return err
			}
			allBooks = append(allBooks, book)
			return nil
		})
	})

	if err != nil {
		context.AbortWithError(http.StatusNotFound, err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"books": allBooks})
}

func (h handler) addBook(context *gin.Context) {
	var input createBookInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, _ := gonanoid.New()

	book := books.Book{
		ID:          id,
		Title:       input.Title,
		Author:      input.Author,
		TotalPages:  input.TotalPages,
		DailyGoal:   input.DailyGoal,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		CurrentPage: 0,
		IsActive:    true,
		IsDone:      false,
		PagesRead:   make(map[time.Time]uint),
	}

	err := h.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Books"))
		if b == nil {
			return bolt.ErrBucketNotFound
		}

		bookData, err := json.Marshal(book)
		if err != nil {
			return err
		}

		return b.Put([]byte(book.ID), bookData)

	})

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pagesReadToday := CountPagesReadToday(book.PagesRead)

	context.JSON(http.StatusOK, gin.H{"book": book, "pagesReadToday": pagesReadToday})
}

func (h handler) deleteBook(context *gin.Context) {
	bookID := context.Param("id")

	err := h.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Books"))
		if b == nil {
			return bolt.ErrBucketNotFound
		}

		bookData := b.Get([]byte(bookID))
		if bookData == nil {
			return bolt.ErrBucketNotFound
		}

		return b.Delete([]byte(bookID))
	})

	if err != nil {
		if err == bolt.ErrBucketNotFound {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить запись"})
		}
		return
	}

	context.JSON(http.StatusOK, gin.H{"success": true})
}

func (h handler) updateCurrentPage(context *gin.Context) {
	bookId := context.Param("id")

	var book books.Book

	var input updateBookInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Books"))
		if b == nil {
			return bolt.ErrBucketNotFound
		}

		bookData := b.Get([]byte(bookId))
		if bookData == nil {
			return bolt.ErrBucketNotFound
		}

		if err := json.Unmarshal(bookData, &book); err != nil {
			return err
		}

		if book.IsDone {
			return nil
		}

		book.PagesRead[time.Now()] = input.PagesRead
		book.UpdatedAt = time.Now()

		if book.CurrentPage+input.PagesRead >= book.TotalPages {
			book.IsActive = false
			book.IsDone = true
			book.FinishedAt = time.Now()
			book.CurrentPage = book.TotalPages
		} else {
			book.CurrentPage += input.PagesRead
		}

		updatedBook, err := json.Marshal(book)
		if err != nil {
			return err
		}

		return b.Put([]byte(book.ID), updatedBook)
	})

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pagesReadToday := CountPagesReadToday(book.PagesRead)

	context.JSON(http.StatusOK, gin.H{"book": book, "pagesReadToday": pagesReadToday})
}

func CountPagesReadToday(pagesRead map[time.Time]uint) uint {
	if len(pagesRead) == 0 {
		return 0
	}

	today := time.Now().Truncate(24 * time.Hour)
	var totalPages uint

	for date, pages := range pagesRead {
		if date.Truncate(24 * time.Hour).Equal(today) {
			totalPages += pages
		}
	}

	return totalPages
}
