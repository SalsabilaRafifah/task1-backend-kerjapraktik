// file berisi fungsi untuk membaca variabel lingkungan dari file .env menggunakan pustaka godotenv

package config
import (
	"fmt"//menyediakan fungsi-fungsi dasar untuk mencetak dan memformat output.
	"os" //memberikan akses ke fungsi-fungsi sistem operasi, termasuk untuk mengakses variabel lingkungan.
	"github.com/joho/godotenv" //untuk membaca nilai variabel lingkungan dari file .env.
)
// menerima satu parameter key yang merupakan kunci untuk variabel lingkungan yang ingin diambil nilainya.
func Config(key string) string {
	// Fungsi ini mencoba membaca variabel lingkungan dari file .env di direktori yang sama dengan file sumber kode. Jika file .env tidak ditemukan atau terjadi kesalahan saat membacanya, maka sebuah pesan kesalahan dicetak menggunakan fmt.Print.
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}

	// Mengembalikan nilai dari variabel lingkungan sesuai dengan kuncinya (key)
	return os.Getenv(key)
}