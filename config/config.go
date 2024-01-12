// file berisi fungsi untuk membaca variabel lingkungan dari file .env menggunakan library godotenv

package config
import (
	"fmt"//menyediakan fungsi-fungsi dasar untuk mencetak dan memformat output.
	"os" //memberikan akses ke fungsi-fungsi sistem operasi, termasuk untuk mengakses variabel lingkungan.
	"github.com/joho/godotenv" //untuk membaca nilai variabel lingkungan dari file .env.
)
// key sebagai kunci variabel lingkungan yang ingin diambil nilainya.
func Config(key string) string {
	// membaca variabel lingkungan dari file .env
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}

	// Mengembalikan nilai dari variabel lingkungan sesuai dengan kuncinya (key)
	return os.Getenv(key)
}