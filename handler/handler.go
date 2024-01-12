// file berisi fungsi-fungsi yang menangani permintaan HTTP dan berinteraksi dengan database

package handler

import (
	"github.com/SalsabilaRafifah/go-fiber-postgres/database"
	"github.com/SalsabilaRafifah/go-fiber-postgres/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// c sbagai objek konteks Fiber untuk memproses informasi permintaan dan mengirimkan respons balik ke klien.
// Fungsi untuk membuat dosen baru dalam database
func CreateLecturer(c *fiber.Ctx) error {
	db := database.DB.Db
	lecturer := new(model.Lecturer)
	err := c.BodyParser(lecturer)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// Membuat entitas dosen baru dalam database menggunakan objek model.Lecturer
	err = db.Create(&lecturer).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create lecturer", "data": err})
	}
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Lecturer has created", "data": lecturer})
}

// Fungsi untuk mendapatkan semua dosen dari database
func GetAllLecturers(c *fiber.Ctx) error {
	db := database.DB.Db
	var lecturers []model.Lecturer
	db.Find(&lecturers)
	if len(lecturers) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Lecturers not found", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Lecturers Found", "data": lecturers})
}

// Fungsi untuk mendapatkan satu dosen berdasarkan ID dari database
func GetSingleLecturer(c *fiber.Ctx) error {
	db := database.DB.Db
	// Mengambil ID dosen dari parameter permintaan.
	id := c.Params("id")
	var lecturer model.Lecturer
	// Mencari satu dosen dalam database berdasarkan ID.
	db.Find(&lecturer, "id = ?", id)
	if lecturer.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Lecturer not found", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Lecturer Found", "data": lecturer})
}

// Fungsi untuk memperbarui informasi dosen dalam database berdasarkan ID
func UpdateLecturer(c *fiber.Ctx) error {
	type updateLecturer struct {
		Nama               string `json:"nama"`
		Golongan           string `json:"golongan"`
		Jabatan            string `json:"jabatan"`
		BidangKeahlian     string `json:"bidang_keahlian"`
		PendidikanTerakhir string `json:"pendidikan_terakhir"`
		Email              string `json:"email"`
	}
	db := database.DB.Db
	var lecturer model.Lecturer
	id := c.Params("id")
	// Mencari satu dosen dalam database berdasarkan ID.
	db.Find(&lecturer, "id = ?", id)
	if lecturer.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Lecturer not found", "data": nil})
	}
	// variabel untuk membaca data yang akan diperbarui
	var updateLecturerData updateLecturer
	// Membaca body request untuk mendapatkan data yang akan diperbarui
	err := c.BodyParser(&updateLecturerData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// Memperbarui field dosen sesuai dengan data yang diterima.
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
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Lecturer Found", "data": lecturer})
}

// Fungsi untuk menghapus dosen berdasarkan ID dari database
func DeleteLecturerByID(c *fiber.Ctx) error {
	db := database.DB.Db
	var lecturer model.Lecturer
	id := c.Params("id")
	db.Find(&lecturer, "id = ?", id)
	if lecturer.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Lecturer not found", "data": nil})
	}
	// Menghapus dosen dari database berdasarkan ID.
	err := db.Delete(&lecturer, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete lecturer", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Lecturer deleted"})
}
