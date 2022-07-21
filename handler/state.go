package handler

import (
	"library-app/config"

	"github.com/go-redis/redis/v9"
	"github.com/gofiber/fiber/v2"
)

func GetStates(c *fiber.Ctx) error {
	rdb := config.RDB

	states := map[string]string{}

	iter := rdb.Scan(CTX, 0, "book:*", 0).Iterator()
	for iter.Next(CTX) {
		state, _ := rdb.Get(CTX, iter.Val()).Result()
		states[iter.Val()] = state
	}

	return c.JSON(states)
}

func GetStateByKey(c *fiber.Ctx) error {
	rdb := config.RDB

	bookId := c.Params("book_id")

	state, err := rdb.Get(CTX, "book:"+bookId).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "State not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"state": state,
	})
}
