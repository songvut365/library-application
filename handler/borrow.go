package handler

import (
	"context"
	"fmt"
	"library-app/config"
	"library-app/model"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/gofiber/fiber/v2"
)

var ctx = context.Background()

func BorrowBook(c *fiber.Ctx) error {
	db := config.DB
	rdb := config.RDB

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
	state, err := rdb.Get(ctx, bookId).Result()

	if err == redis.Nil {
		book := model.Book{}
		err = db.Model(&model.Book{}).Where("id = ?", input.BookID).First(&book).Error
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Book not found",
			})
		}
		rdb.Set(ctx, bookId, book.State, 0)
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
		rdb.Set(ctx, bookId, "borrowed", 0)

		borrow := model.Borrow{
			Date:     time.Now().String(),
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
	db := config.DB

	borrowedBooks := []model.Book{}
	db.Model(&model.Book{}).Where("state = ?", "borrowed").Find(&borrowedBooks)

	return c.Status(fiber.StatusOK).JSON(borrowedBooks)
}
