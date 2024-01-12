// Mendefinisikan rute-rute API dengan menggunakan grup "/api" dan grup "/user"

package router

import (
	"github.com/SalsabilaRafifah/go-fiber-postgres/handler" // untuk menangani permintaan HTTP terkait dosen.
	"github.com/gofiber/fiber/v2"                           // fiber adalah framework Go untuk menangani permintaan HTTP
)

// Untuk mendefinisikan semua rute API yang digunakan dalam aplikasi
func SetupRoutes(app *fiber.App) {
	// Mengelompokkan rute-rute API untuk membantu dalam mengorganisir dan memisahkan rute-rute berdasarkan fungsionalitas atau entitas tertentu.
	api := app.Group("/api")
	v1 := api.Group("/lecturer")

	// mendefinisikan rute-rute HTTP terkait dosen di dalam grup /api/user.
	v1.Get("/", handler.GetAllLecturers)
	v1.Get("/:id", handler.GetSingleLecturer)
	v1.Post("/", handler.CreateLecturer)
	v1.Put("/:id", handler.UpdateLecturer)
	v1.Delete("/:id", handler.DeleteLecturerByID)
}
