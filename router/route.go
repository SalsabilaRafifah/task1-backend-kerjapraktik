// Mendefinisikan rute-rute API dengan menggunakan grup "/api" dan grup "/user". Setiap grup memiliki rute-rute yang terkait dengan pengelolaan data pengguna (user).

package router

import (
	"github.com/SalsabilaRafifah/go-fiber-postgres/handler" //untuk menangani permintaan HTTP terkait pengguna.
	"github.com/gofiber/fiber/v2"                           //  Fiber adalah framework web untuk Go yang digunakan untuk menangani permintaan HTTP dengan efisien.
)

// Untuk mendefinisikan semua rute API yang digunakan dalam aplikasi
// menerima parameter app yang merupakan instance dari objek fiber.App.
func SetupRoutes(app *fiber.App) {
	// Mengelompokkan rute-rute API untuk membantu dalam mengorganisir dan memisahkan rute-rute berdasarkan fungsionalitas atau entitas tertentu.
	api := app.Group("/api")     // api adalah grup utama dengan prefiks /api
	v1 := api.Group("/lecturer") // v1 adalah subgrup dari api dengan prefiks /user

	// mendefinisikan rute-rute HTTP terkait pengguna di dalam grup /api/user.
	// fungsi handler dari package handler akan dipanggil ketika permintaan HTTP mencocokkan salah satu rute, dan Fiber akan menangani proses HTTP secara efisien.
	v1.Get("/", handler.GetAllLecturers)
	v1.Get("/:id", handler.GetSingleLecturer)
	v1.Post("/", handler.CreateLecturer)
	v1.Put("/:id", handler.UpdateLecturer)
	v1.Delete("/:id", handler.DeleteLecturerByID)
}
