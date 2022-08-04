package handler

import (
	"library-app/database"
	"library-app/model"

	"github.com/gofiber/fiber/v2"
)

func GetMembers(c *fiber.Ctx) error {
	db := database.DB

	members := []model.Member{}
	db.Model(&model.Member{}).Order("id asc").Find(&members)

	return c.Status(fiber.StatusOK).JSON(members)
}

func GetMemberById(c *fiber.Ctx) error {
	db := database.DB

	memberId := c.Params("member_id")
	member := model.Member{}
	err := db.Model(&model.Member{}).Where("id = ?", memberId).First(&member).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Member not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(member)
}

func CreateMember(c *fiber.Ctx) error {
	db := database.DB

	member := model.Member{}
	c.BodyParser(&member)
	db.Create(&member)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Create member successfully",
	})
}

func UpdateMember(c *fiber.Ctx) error {
	db := database.DB

	updateMember := model.Member{}
	c.BodyParser(&updateMember)

	memberId := c.Params("member_id")
	db.Model(&model.Member{}).Where("id = ?", memberId).Updates(&updateMember)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Update member successfully",
	})
}

func DeleteMember(c *fiber.Ctx) error {
	db := database.DB

	memberId := c.Params("member_id")
	member := model.Member{}
	err := db.Model(&model.Member{}).Where("id = ?", memberId).First(&member).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Member not found",
		})
	}

	db.Model(&model.Member{}).Delete(&member)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Delete member successfully",
	})
}
