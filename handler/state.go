package handler

import (
	"context"
	"library-app/database"

	"github.com/go-redis/redis/v9"
	"github.com/gofiber/fiber/v2"
)

func GetStates(c *fiber.Ctx) error {
	rdb := database.RDB

	states := map[string]string{}

	iter := rdb.Scan(context.Background(), 0, "book:*", 0).Iterator()
	for iter.Next(context.Background()) {
		state, _ := rdb.Get(context.Background(), iter.Val()).Result()
		states[iter.Val()] = state
	}

	return c.JSON(states)
}

func GetStateByKey(c *fiber.Ctx) error {
	rdb := database.RDB

	bookId := c.Params("book_id")

	state, err := rdb.Get(context.Background(), "book:"+bookId).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "State not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"state": state,
	})
}
