package handler

import (
	"context"
	"library-app/database"
	"library-app/model"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	db := database.DB

	books := []model.BookResponse{}
	db.Model(&model.Book{}).
		Order("id asc").
		Select("books.id, book_details.name, books.date, books.state").
		Joins("left join book_details on book_details.id = books.book_id").
		Scan(&books)

	return c.Status(fiber.StatusOK).JSON(books)
}

func GetBookById(c *fiber.Ctx) error {
	db := database.DB

	bookId := c.Params("book_id")

	book := model.BookResponse{}
	err := db.Model(&model.Book{}).
		Select("books.id, book_details.name, books.date, books.state").
		Joins("left join book_details on book_details.id = books.book_id").
		Where("books.id = ?", bookId).
		First(&book).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Book not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(book)
}

func CreateBook(c *fiber.Ctx) error {
	db := database.DB

	var input model.NewBook
	c.BodyParser(&input)

	// Find book detail
	bookDetail := model.BookDetail{}
	err := db.Model(&model.BookDetail{}).Where("name = ?", input.Name).First(&bookDetail).Error
	if err != nil {
		bookDetail.Name = input.Name
		bookDetail.Amount = 0
		db.Model(&model.BookDetail{}).Create(&bookDetail)
	}

	// Add book
	newBook := model.Book{
		BookID: *bookDetail.ID,
		Date:   time.Now().Format("02-01-2006 15:04:05"),
		State:  "available",
	}
	db.Model(&model.Book{}).Create(&newBook)

	// Update book amount
	bookDetail.Amount = bookDetail.Amount + 1
	db.Updates(&bookDetail)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Create book successfully",
		"book":    newBook,
	})
}

func UpdateBook(c *fiber.Ctx) error {
	db := database.DB
	rdb := database.RDB

	updateBook := model.UpdateBook{}
	c.BodyParser(&updateBook)

	// Find book
	bookId := c.Params("book_id")
	book := model.Book{}
	err := db.Model(&model.Book{}).Where("id = ?", bookId).First(&book).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Book not found",
		})
	}

	// Update book
	book.State = updateBook.State
	db.Model(&model.Book{}).Where("id = ?", bookId).Updates(&book)

	// Update state on Redis
	rdb.Set(context.Background(), "book:"+bookId, updateBook.State, 0)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Update book successfully",
		"book":    book,
	})
}

func DeleteBook(c *fiber.Ctx) error {
	db := database.DB
	rdb := database.RDB

	updateBook := model.UpdateBook{}
	c.BodyParser(&updateBook)

	// Find book
	bookId := c.Params("book_id")
	book := model.Book{}
	err := db.Model(&model.Book{}).Where("id = ?", bookId).First(&book).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Book not found",
		})
	}

	// Delete book
	db.Model(&model.Book{}).Delete(&book)

	// Decrease book amount
	bookDetail := model.BookDetail{}
	db.Model(&model.BookDetail{}).Where("id = ?", book.BookID).First(&bookDetail)

	if bookDetail.Amount == 1 {
		bookDetail.Amount = 0
		db.Model(&model.BookDetail{}).Where("id = ?", book.BookID).Select("amount").Update("amount", 0)
	} else {
		bookDetail.Amount = bookDetail.Amount - 1
		db.Model(&model.BookDetail{}).Where("id = ?", book.BookID).Updates(&bookDetail)
	}

	// Delete from Redis
	rdb.Del(context.Background(), "book:"+bookId)

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "Delete book successfully",
		"book":    book,
	})
}
