// file berisi fungsi-fungsi yang menangani permintaan HTTP dan berinteraksi dengan database

package handler

import (
	"github.com/SalsabilaRafifah/go-fiber-postgres/database" // menyediakan akses ke database
	"github.com/SalsabilaRafifah/go-fiber-postgres/model"    // definisi struktur data
	"github.com/gofiber/fiber/v2"                            // framework web
	"github.com/google/uuid"                                 // menangani UUID
)

// c adalah objek konteks Fiber yang digunakan untuk memproses informasi permintaan dan mengirimkan respons balik ke klien.
// CreateLecturer: Fungsi untuk membuat pengguna baru dalam database
func CreateLecturer(c *fiber.Ctx) error {
	// Mengambil instance database dari package database
	db := database.DB.Db
	// Membuat instance baru dari struktur data model.Lecturer yang akan digunakan untuk menyimpan data pengguna dari permintaan.
	lecturer := new(model.Lecturer)
	// Membaca body request yang dikirim oleh klien dan mengonversinya menjadi objek Lecturer
	err := c.BodyParser(lecturer)
	// Jika parsing body request gagal, mengembalikan respons error dengan status 500.
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// Membuat entitas pengguna baru dalam database menggunakan objek model.Lecturer
	err = db.Create(&lecturer).Error
	// Jika operasi penciptaan pengguna gagal, mengembalikan respons error dengan status 500.
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create lecturer", "data": err})
	}
	// Jike berhasil, mengembalikan response sukses dengan status 201 dan data pengguna yang baru dibuat
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Lecturer has created", "data": lecturer})
}

// Get All Lecturers: Fungsi untuk mendapatkan semua pengguna dari database
func GetAllLecturers(c *fiber.Ctx) error {
	// Mendapatkan akses ke instance database dan membuat variabel untuk menyimpan daftar pengguna.
	db := database.DB.Db
	var lecturers []model.Lecturer
	// Mengambil semua pengguna dari database menggunakan db.Find.
	db.Find(&lecturers)
	// Jika tidak ada pengguna yang ditemukan, mengembalikan respons error dengan status 404.
	if len(lecturers) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Lecturers not found", "data": nil})
	}
	// Jika berhasil, mengembalikan respons sukses dengan status 200 dan data pengguna yang ditemukan.
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Lecturers Found", "data": lecturers})
}

// GetSingleLecturer: Fungsi untuk mendapatkan satu pengguna berdasarkan ID dari database
func GetSingleLecturer(c *fiber.Ctx) error {
	db := database.DB.Db
	// Mengambil ID pengguna dari parameter permintaan.
	id := c.Params("id")
	var lecturer model.Lecturer
	// Mencari satu pengguna dalam database berdasarkan ID.
	db.Find(&lecturer, "id = ?", id)
	// Jika pengguna tidak ditemukan, mengembalikan respons error dengan status 404.
	if lecturer.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Lecturer not found", "data": nil})
	}
	// Jika berhasil, mengembalikan respons sukses dengan status 200 dan data pengguna yang ditemukan.
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Lecturer Found", "data": lecturer})
}

// UpdateLecturer: Fungsi untuk memperbarui informasi pengguna dalam database berdasarkan ID
func UpdateLecturer(c *fiber.Ctx) error {
	type updateLecturer struct {
		Nama               string `json:"nama"` // informasi dasar dari seorang dosen
		Golongan           string `json:"golongan"`
		Jabatan            string `json:"jabatan"`
		BidangKeahlian     string `json:"bidang_keahlian"`
		PendidikanTerakhir string `json:"pendidikan_terakhir"`
		Email              string `json:"email"`
	}
	db := database.DB.Db
	var lecturer model.Lecturer
	// Mengambil ID pengguna dari parameter permintaan.
	id := c.Params("id")
	// Mencari satu pengguna dalam database berdasarkan ID.
	db.Find(&lecturer, "id = ?", id)
	// Jika pengguna tidak ditemukan, mengembalikan respons error dengan status 404.
	if lecturer.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Lecturer not found", "data": nil})
	}
	// updateLecturer untuk membaca data yang akan diperbarui
	var updateLecturerData updateLecturer
	// Membaca body request untuk mendapatkan data yang akan diperbarui
	err := c.BodyParser(&updateLecturerData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// Memperbarui field pengguna sesuai dengan data yang diterima.
	if updateLecturerData.Nama != "" {
		lecturer.Nama = updateLecturerData.Nama
	}
	if updateLecturerData.Golongan != "" {
		lecturer.Golongan = updateLecturerData.Golongan
	}
	if updateLecturerData.Jabatan != "" {
		lecturer.Jabatan = updateLecturerData.Jabatan
	}
	if updateLecturerData.BidangKeahlian != "" {
		lecturer.BidangKeahlian = updateLecturerData.BidangKeahlian
	}
	if updateLecturerData.PendidikanTerakhir != "" {
		lecturer.PendidikanTerakhir = updateLecturerData.PendidikanTerakhir
	}
	if updateLecturerData.Email != "" {
		lecturer.Email = updateLecturerData.Email
	}
	// Menyimpan perubahan ke dalam database.
	db.Save(&lecturer)
	// Jika berhasil, mengembalikan respons sukses dengan status 200 dan data pengguna yang diperbarui.
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Lecturer Found", "data": lecturer})
}

// DeleteLecturerByID: Fungsi untuk menghapus pengguna berdasarkan ID dari database
func DeleteLecturerByID(c *fiber.Ctx) error {
	db := database.DB.Db
	var lecturer model.Lecturer
	// Mengambil ID pengguna dari parameter permintaan.
	id := c.Params("id")
	// Mencari satu pengguna dalam database berdasarkan ID.
	db.Find(&lecturer, "id = ?", id)
	// Jika pengguna tidak ditemukan, mengembalikan respons error dengan status 404.
	if lecturer.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Lecturer not found", "data": nil})
	}
	// Menghapus pengguna dari database berdasarkan ID.
	err := db.Delete(&lecturer, "id = ?", id).Error
	// Jika operasi penghapusan gagal, mengembalikan respons error dengan status 404.
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete lecturer", "data": nil})
	}
	// Jika berhasil, mengembalikan respons sukses dengan status 200.
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Lecturer deleted"})
}
