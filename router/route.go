package router

import (
	"library-app/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Book Detail
	bookDetail := app.Group("/book_detail")
	bookDetail.Get("/", handler.GetAllBookDetail)
	bookDetail.Get("/:book_detail_id", handler.GetBookDetailById)
	bookDetail.Put("/:book_detail_id", handler.UpdateBookDetail)

	// Book
	books := app.Group("/books")
	books.Get("/", handler.GetBooks)
	books.Get("/:book_id", handler.GetBookById)
	books.Post("/", handler.CreateBook)
	books.Put("/:book_id", handler.UpdateBook)
	books.Delete("/:book_id", handler.DeleteBook)

	// Member
	members := app.Group("/members")
	members.Get("/", handler.GetMembers)
	members.Get(("/:member_id"), handler.GetMemberById)
	members.Post("/", handler.CreateMember)
	members.Put("/:member_id", handler.UpdateMember)
	members.Delete("/:member_id", handler.DeleteMember)

	// Borrow
	borrows := app.Group("/borrows")
	borrows.Get("/", handler.GetBorrowList)
	borrows.Post("/", handler.BorrowBook)

	// State
	states := app.Group("/states")
	states.Get("/", handler.GetStates)
	states.Get("/:book_id", handler.GetStateByKey)
}
