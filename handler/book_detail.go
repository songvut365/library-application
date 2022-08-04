package handler

import (
	"library-app/database"
	"library-app/model"

	"github.com/gofiber/fiber/v2"
)

func GetAllBookDetail(c *fiber.Ctx) error {
	db := database.DB

	bookDetails := []model.BookDetail{}
	db.Order("id asc").Find(&bookDetails)

	return c.Status(fiber.StatusOK).JSON(bookDetails)
}

func GetBookDetailById(c *fiber.Ctx) error {
	db := database.DB

	bookDetailId := c.Params("book_detail_id")

	bookDetail := model.BookDetail{}
	err := db.Where("id = ?", bookDetailId).First(&bookDetail).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Book detail not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(bookDetail)
}

func UpdateBookDetail(c *fiber.Ctx) error {
	db := database.DB

	bookDetailId := c.Params("book_detail_id")

	updateBookDetail := model.UpdateBookDetail{}
	c.BodyParser(&updateBookDetail)

	bookDetail := model.BookDetail{}
	err := db.Model(&model.BookDetail{}).
		Where("id = ?", bookDetailId).
		Update("name", updateBookDetail.Name).
		First(&bookDetail).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Book detail not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":     "Update Book Detail Success",
		"book_detail": bookDetail,
	})
}
