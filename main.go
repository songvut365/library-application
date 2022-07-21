package main

import (
	"library-app/config"
	"library-app/handler"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	// Config
	config.InitDatabase()
	config.InitRedis()

	// Router
	bookDetail := app.Group("/book_detail")
	bookDetail.Get("/", handler.GetAllBookDetail)
	bookDetail.Get("/:book_detail_id", handler.GetBookDetailById)
	bookDetail.Put("/:book_detail_id", handler.UpdateBookDetail)

	books := app.Group("/books")
	books.Get("/", handler.GetBooks)
	books.Get("/:book_id", handler.GetBookById)
	books.Post("/", handler.CreateBook)
	books.Put("/:book_id", handler.UpdateBook)
	books.Delete("/:book_id", handler.DeleteBook)

	members := app.Group("/members")
	members.Get("/", handler.GetMembers)
	members.Get(("/:member_id"), handler.GetMemberById)
	members.Post("/", handler.CreateMember)
	members.Put("/:member_id", handler.UpdateMember)
	members.Delete("/:member_id", handler.DeleteMember)

	borrows := app.Group("/borrows")
	borrows.Get("/", handler.GetBorrowList)
	borrows.Post("/", handler.BorrowBook)

	states := app.Group("/states")
	states.Get("/", handler.GetStates)
	states.Get("/:book_id", handler.GetStateByKey)

	// Run
	log.Fatal(app.Listen(":3000"))
}
