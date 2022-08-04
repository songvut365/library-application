package handler

import (
	"context"
	"fmt"
	"library-app/database"
	"library-app/model"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/gofiber/fiber/v2"
)

func BorrowBook(c *fiber.Ctx) error {
	db := database.DB
	rdb := database.RDB

	input := model.BorrowInput{}
	c.BodyParser(&input)

	// Check member
	member := model.Member{}
	err := db.Model(&model.Member{}).Where("national_id = ?", input.NationalID).First(&member).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Member not found",
		})
	}

	// Get book state from Redis
	bookId := fmt.Sprintf("%v", input.BookID)
	state, err := rdb.Get(context.Background(), "book:"+bookId).Result()

	if err == redis.Nil {
		book := model.Book{}
		err = db.Model(&model.Book{}).Where("id = ?", input.BookID).First(&book).Error
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Book not found",
			})
		}
		rdb.Set(context.Background(), "book:"+bookId, book.State, 0)
		state = book.State
	}

	// Check state
	var message string

	switch state {
	case "damaged":
		message = "Damaged book"
	case "not available":
		message = "Book not available"
	case "borrowed":
		message = "Book has been borrowed"
	case "available":
		db.Model(&model.Book{}).Where("id = ?", input.BookID).Update("state", "borrowed")
		rdb.Set(context.Background(), "book:"+bookId, "borrowed", 0)

		borrow := model.Borrow{
			Date:     time.Now().Format("02-01-2006 15:04:05"),
			BookID:   input.BookID,
			MemberID: *member.ID,
		}
		db.Model(&model.Borrow{}).Create(&borrow)

		message = "Borrow successfully"
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": message,
	})
}

func GetBorrowList(c *fiber.Ctx) error {
	db := database.DB

	borrowedBooks := []model.BorrowResponse{}
	db.Model(&model.Borrow{}).
		Select("borrows.id, borrows.book_id, book_details.name as book_name, borrows.member_id, borrows.date").
		Joins("join books on books.id = borrows.book_id").
		Joins("join book_details on book_details.id = books.book_id").
		Scan(&borrowedBooks)

	return c.Status(fiber.StatusOK).JSON(borrowedBooks)
}
