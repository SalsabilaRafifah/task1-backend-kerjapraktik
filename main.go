// file  utama menggabungkan semua komponen seperti koneksi database, middleware, rute-rute, dan server pada satu tempat
// untuk membentuk aplikasi web sederhana menggunakan Fiber dan PostgreSQL.

package main
import (
	"github.com/SalsabilaRafifah/go-fiber-postgres/database" // berisi fungsi-fungsi terkait pengaturan dan koneksi ke database.
	"github.com/SalsabilaRafifah/go-fiber-postgres/router"   // berisi definisi rute-rute API menggunakan framework Fiber.
	"github.com/gofiber/fiber/v2"                            // paket utama dari fiber
	"github.com/gofiber/fiber/v2/middleware/cors"            // Middleware untuk menangani CORS (Cross-Origin Resource Sharing).
	"github.com/gofiber/fiber/v2/middleware/logger"          // Middleware untuk logging.
	_ "github.com/lib/pq"                                    // Driver PostgreSQL untuk Golang.
)

// fungsi utama yang melakukan konfigurasi dasar server
func main() {
	// menginisialisasi koneksi ke database PostgreSQL
	database.Connect()

	// membuat instance baru dari aplikasi fiber yang akan digunakan untuk menangani rute dan middleware.
	app := fiber.New()

	// menambahkan middleware logger untuk logging setiap request yang masuk ke server.
	app.Use(logger.New())

	// Menambahkan middleware CORS untuk menangani kebijakan Cross-Origin Resource Sharing.
	// Middleware CORS memungkinkan atau memblokir akses ke sumber daya pada server dari luar domain asal.
	app.Use(cors.New())

	// mendefinisikan rute-rute API menggunakan framework Fiber
	router.SetupRoutes(app)

	// Menambahkan middleware untuk menangani rute yang tidak tersedia dengan mengembalikan status 404 ("Not Found").
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	// Memulai server Fiber dan mendengarkan pada port 8080. Aplikasi akan berjalan secara asinkron, menangani permintaan HTTP yang masuk.
	app.Listen(":8080")
}
