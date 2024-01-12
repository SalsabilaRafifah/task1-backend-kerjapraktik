// file berisi fungsi-fungsi untuk menghubungkan aplikasi ke database PostgreSQL menggunakan GORM (GO Object Relational Mapper)

package database

import (
	"fmt"     //untuk mencetak dan memformat output.
	"log"     //menyediakan fungsi-fungsi untuk logging.
	"os"      //memberikan akses ke fungsi-fungsi sistem operasi.
	"strconv" //untuk konversi string ke tipe data numerik.

	"github.com/SalsabilaRafifah/go-fiber-postgres/config"
	"github.com/SalsabilaRafifah/go-fiber-postgres/model"
	"gorm.io/driver/postgres" //driver GORM untuk PostgreSQL.
	"gorm.io/gorm"            //GORM adalah ORM untuk Go.
	"gorm.io/gorm/logger"     //menyediakan logger untuk GORM.
)

// Struct untuk menyimpan instance database yang dikembalikan oleh GORM setelah terjadi koneksi ke database
type Dbinstance struct {
	Db *gorm.DB // properti Db bertipe data pointer ke objek gorm.DB.
}

// Variabel DB digunakan untuk menyimpan instance database sehingga dapat diakses dari berbagai bagian aplikasi.
var DB Dbinstance

// Connect function untuk membuat koneksi ke database PostgreSQL menggunakan konfigurasi yang diambil dari .env
func Connect() {
	// Nilai port diambil dari konfigurasi menggunakan fungsi Config dari paket config
	p := config.Config("DB_PORT")
	// because our config function returns a string, we are parsing our str to int here
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		fmt.Println("Error parsing str to int")
	}

	// mengembalikan informasi koneksi dari file .env dalam bentuk string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", config.Config("DB_HOST"), config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"), port)
	// Membuat koneksi ke database PostgreSQL menggunakan GORM dengan konfigurasi yang telah disiapkan
	// fungsi open memerlukan dua parameter: jenis database yang akan digunakan dan konfigurasi tambahan
	// db adalah variabel yang menyimpan instance dari objek gorm.DB. Instance ini akan digunakan untuk berinteraksi dengan database, seperti menjalankan query atau melakukan migrasi tabel.
	// Jika terjadi kesalahan selama pembukaan koneksi, nilai variabel err akan berisi pesan kesalahan
	// untuk deklarasi dan penanganan kesalahan sekaligus.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //Mengatur logger untuk instance gorm.DB. GORM memiliki sistem logging yang dapat dikonfigurasi. Di sini, digunakan logger default GORM dengan level log set ke logger.Info. Artinya, hanya pesan log dengan level info atau di atasnya yang akan dicetak.
	})
	// mengecek kesalahan koneksi = Jika terdapat kesalahan saat membuka koneksi, aplikasi akan mengakhiri dirinya dengan pesan kesalahan yang dicetak.
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}
	// jika berhasil melakukan koneksi
	log.Println("Connected")
	// mengatur logger untuk objek gorm.DB. Di sini, logger diatur untuk mencetak pesan log dengan level informasi (logger.Info). GORM menyediakan logger bawaan yang dapat diatur untuk mencatat aktivitas database, dan dalam hal ini, hanya pesan log dengan level informasi yang dicetak.
	db.Logger = logger.Default.LogMode(logger.Info)
	// mencetak pesan log bahwa koneksi ke database telah berhasil dan migrasi tabel sedang berlangsung.
	log.Println("running migrations")
	// melakukan migrasi otomatis untuk tabel User menggunakan GORM. Ini akan memastikan bahwa struktur tabel di database sesuai dengan definisi model di aplikasi.
	db.AutoMigrate(&model.Lecturer{})
	// Menyimpan instance database dalam variabel DB untuk digunakan di berbagai bagian aplikasi.
	DB = Dbinstance{
		Db: db,
	}
}

// terhubung ke database PostgreSQL dan melakukan migrasi tabel otomatis saat aplikasi dimulai.
