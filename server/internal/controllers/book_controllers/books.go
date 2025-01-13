package book_controllers

import (
	"encoding/json"
	"errors"
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

func (h handler) init(context *gin.Context) {
	userId := context.Param("userId")
	var userData books.UserData

	err := h.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("Books"))
		if bucket == nil {
			return bolt.ErrBucketNotFound
		}

		userDataDb := bucket.Get([]byte(userId))
		if userDataDb == nil {
			return errors.New("user's data is not found")
		}
		if err := json.Unmarshal(userDataDb, &userData); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": true})
		return
	}

	context.JSON(http.StatusOK, gin.H{"bookId": userData.CurrentBookId})

}

func (h handler) getCurrentBook(context *gin.Context) {
	userId := context.Param("userId")
	var userData books.UserData
	var currentBook books.Book

	err := h.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("Books"))
		if bucket == nil {
			return bolt.ErrBucketNotFound
		}

		userDataDb := bucket.Get([]byte(userId))
		if userDataDb == nil {
			return errors.New("user's data is not found")
		}
		if err := json.Unmarshal(userDataDb, &userData); err != nil {
			return err
		}

		return nil
	})

	var isFound bool
	for _, book := range userData.Books {
		if userData.CurrentBookId == book.ID {
			currentBook = book
			isFound = true
			break
		}
	}

	if !isFound {
		context.JSON(http.StatusOK, gin.H{"error": true})
		return
	}

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": true})
		return
	}

	pagesReadToday := countPagesReadToday(currentBook.PagesRead)
	context.JSON(http.StatusOK, gin.H{"currentBook": currentBook, "pagesReadToday": pagesReadToday})

}

func (h handler) getBook(context *gin.Context) {
	userId := context.Param("userId")
	bookId := context.Param("bookId")

	var userData books.UserData

	err := h.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("Books"))
		if bucket == nil {
			return bolt.ErrBucketNotFound
		}

		userDataDb := bucket.Get([]byte(userId))
		if userDataDb == nil {
			return bolt.ErrBucketNotFound
		}

		if err := json.Unmarshal(userDataDb, &userData); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	var book books.Book
	var foundBook bool

	for _, val := range userData.Books {
		if bookId == val.ID {
			book = val
			foundBook = true
			break
		}
	}

	if !foundBook {
		context.JSON(http.StatusBadRequest, gin.H{"error": "запись не найдена"})
		return
	}

	pagesReadToday := countPagesReadToday(book.PagesRead)

	context.JSON(http.StatusOK, gin.H{"book": book, "pagesReadToday": pagesReadToday})
}

func (h handler) getAllBooks(context *gin.Context) {
	userId := context.Param("userId")
	var userData books.UserData

	err := h.DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("Books"))
		if bucket == nil {
			return bolt.ErrBucketNotFound
		}

		userDataDb := bucket.Get([]byte(userId))
		if userDataDb == nil {
			return bolt.ErrBucketNotFound
		}

		if err := json.Unmarshal(userDataDb, &userData); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		context.AbortWithError(http.StatusNotFound, err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"books": userData.Books})
}

func (h handler) addBook(context *gin.Context) {
	userId := context.Param("userId")
	var userData books.UserData

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
		IsDone:      false,
		PagesRead:   make(map[time.Time]uint),
	}

	err := h.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("Books"))
		if bucket == nil {
			return bolt.ErrBucketNotFound
		}

		userDataDb := bucket.Get([]byte(userId))
		if userDataDb == nil {
			return bolt.ErrBucketNotFound
		}

		if err := json.Unmarshal(userDataDb, &userData); err != nil {
			return err
		}

		userData.Books = append(userData.Books, book)
		userData.CurrentBookId = book.ID

		userDataJson, err := json.Marshal(userData)
		if err != nil {
			return err
		}

		return bucket.Put([]byte(userId), userDataJson)

	})

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pagesReadToday := countPagesReadToday(book.PagesRead)

	context.JSON(http.StatusOK, gin.H{"book": book, "pagesReadToday": pagesReadToday})
}

func (h handler) deleteBook(context *gin.Context) {
	bookId := context.Param("bookId")
	userId := context.Param("userId")
	var userData books.UserData

	err := h.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("Books"))
		if bucket == nil {
			return bolt.ErrBucketNotFound
		}

		userDataDb := bucket.Get([]byte(userId))
		if userDataDb == nil {
			return bolt.ErrBucketNotFound
		}

		if err := json.Unmarshal(userDataDb, &userData); err != nil {
			return err
		}

		newBooks := []books.Book{}

		for _, item := range userData.Books {
			if item.ID != bookId {
				newBooks = append(newBooks, item)
			}
		}

		userData.Books = newBooks

		if userData.CurrentBookId == bookId {
			if len(userData.Books) > 0 {
				userData.CurrentBookId = userData.Books[len(userData.Books)-1].ID
			} else {
				userData.CurrentBookId = ""
			}
		}

		userDataJson, err := json.Marshal(userData)
		if err != nil {
			return err
		}

		return bucket.Put([]byte(userId), userDataJson)
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
	bookId := context.Param("bookId")
	userId := context.Param("userId")

	var currentBook books.Book
	var userData books.UserData

	var input updateBookInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("Books"))
		if bucket == nil {
			return bolt.ErrBucketNotFound
		}

		userDataDb := bucket.Get([]byte(userId))
		if userDataDb == nil {
			return bolt.ErrBucketNotFound
		}

		if err := json.Unmarshal(userDataDb, &userData); err != nil {
			return err
		}

		for _, book := range userData.Books {
			if book.ID == bookId {
				currentBook = book
				break
			}
		}

		if currentBook.IsDone {
			return nil
		}

		currentBook.PagesRead[time.Now()] = input.PagesRead
		currentBook.UpdatedAt = time.Now()

		if currentBook.CurrentPage+input.PagesRead >= currentBook.TotalPages {
			currentBook.IsDone = true
			currentBook.FinishedAt = time.Now()
			currentBook.CurrentPage = currentBook.TotalPages
		} else {
			currentBook.CurrentPage += input.PagesRead
		}

		for i, book := range userData.Books {
			if book.ID == currentBook.ID {
				userData.Books[i] = currentBook
				break
			}
		}

		userDataJson, err := json.Marshal(userData)
		if err != nil {
			return err
		}

		return bucket.Put([]byte(userId), userDataJson)
	})

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pagesReadToday := countPagesReadToday(currentBook.PagesRead)

	context.JSON(http.StatusOK, gin.H{"currentBook": currentBook, "pagesReadToday": pagesReadToday})
}

func (h handler) setCurrentBook(context *gin.Context) {
	bookId := context.Param("bookId")
	userId := context.Param("userId")

	var currentBook books.Book
	var userData books.UserData

	err := h.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("Books"))
		if bucket == nil {
			return bolt.ErrBucketNotFound
		}

		userDataDb := bucket.Get([]byte(userId))
		if userDataDb == nil {
			return bolt.ErrBucketNotFound
		}

		if err := json.Unmarshal(userDataDb, &userData); err != nil {
			return err
		}

		userData.CurrentBookId = bookId

		for _, val := range userData.Books {
			if bookId == val.ID {
				currentBook = val
				break
			}
		}

		userDataJson, err := json.Marshal(userData)
		if err != nil {
			return err
		}

		return bucket.Put([]byte(userId), userDataJson)
	})

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pagesReadToday := countPagesReadToday(currentBook.PagesRead)
	context.JSON(http.StatusOK, gin.H{"currentBook": currentBook, "pagesReadToday": pagesReadToday})
}

func countPagesReadToday(pagesRead map[time.Time]uint) uint {
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
