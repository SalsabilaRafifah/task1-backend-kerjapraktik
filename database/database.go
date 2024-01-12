// file berisi fungsi-fungsi untuk menghubungkan aplikasi ke database PostgreSQL menggunakan GORM (GO Object Relational Mapper)

package database

import (
	"fmt"     //untuk mencetak dan memformat output.
	"log"     //menyediakan fungsi-fungsi untuk logging.
	"os"      //untuk mengakses variabel lingkungan.
	"strconv" //untuk konversi string ke tipe data numerik.

	"github.com/SalsabilaRafifah/go-fiber-postgres/config"
	"github.com/SalsabilaRafifah/go-fiber-postgres/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

// Variabel DB digunakan untuk menyimpan instance database sehingga dapat diakses dari berbagai bagian aplikasi.
var DB Dbinstance

// Membuat koneksi ke database PostgreSQL menggunakan konfigurasi yang diambil dari .env
func Connect() {
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		fmt.Println("Error parsing str to int")
	}

	// mengembalikan informasi koneksi dari file .env dalam bentuk string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", config.Config("DB_HOST"), config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"), port)

	// db sebagai variabel yang menyimpan instance dari objek gorm.DB untuk menjalankan query atau melakukan migrasi tabel.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}
	log.Println("Connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	// melakukan migrasi otomatis untuk tabel User menggunakan GORM
	db.AutoMigrate(&model.Lecturer{})
	// Menyimpan instance database dalam variabel DB untuk digunakan di berbagai bagian aplikasi.
	DB = Dbinstance{
		Db: db,
	}
}

// terhubung ke database PostgreSQL dan melakukan migrasi tabel otomatis saat aplikasi dimulai.
