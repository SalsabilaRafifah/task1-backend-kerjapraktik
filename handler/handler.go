// file berisi fungsi-fungsi yang menangani permintaan HTTP dan berinteraksi dengan database

package handler

import (
	"github.com/SalsabilaRafifah/go-fiber-postgres/database"
	"github.com/SalsabilaRafifah/go-fiber-postgres/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// c adalah objek konteks Fiber yang digunakan untuk memproses informasi permintaan dan mengirimkan respons balik ke klien.
// CreateUser: Fungsi untuk membuat pengguna baru dalam database
func CreateUser(c *fiber.Ctx) error {
	db := database.DB.Db
	user := new(model.User)
	// Membaca body request dan menyimpannya dalam objek user dan mengembalikan error if encountered
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	// Membuat entitas user baru dalam database
	err = db.Create(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err})
	}
	// Mengembalikan response berhasil dengan data user yang baru dibuat
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "User has created", "data": user})
}

// Get All Users: Fungsi untuk mendapatkan semua pengguna dari database
func GetAllUsers(c *fiber.Ctx) error {
	db := database.DB.Db
	var users []model.User
	// find all users in the database
	db.Find(&users)
	// If no user found, return an error
	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Users not found", "data": nil})
	}
	// return users yang ditemukan
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Users Found", "data": users})
}

// GetSingleUser: Fungsi untuk mendapatkan satu pengguna berdasarkan ID dari database
func GetSingleUser(c *fiber.Ctx) error {
	db := database.DB.Db
	// get id dari parameter
	id := c.Params("id")
	var user model.User
	// find satu pengguna in the database by id
	db.Find(&user, "id = ?", id)
	// Jika pengguna tidak ditemukan, mengembalikan error
	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}
	// Mengembalikan data pengguna yang ditemukan
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User Found", "data": user})
}

// UpdateUser: Fungsi untuk memperbarui informasi pengguna dalam database berdasarkan ID
func UpdateUser(c *fiber.Ctx) error {
	type updateUser struct {
		Nama               string    `json:"nama"`       // informasi dasar dari seorang dosen
		Golongan           string    `json:"golongan"`
		Jabatan            string    `json:"jabatan"`
		BidangKeahlian     string    `json:"bidang_keahlian"`
		PendidikanTerakhir string    `json:"pendidikan_terakhir"`
		Email              string    `json:"email"`
	}
	db := database.DB.Db
	var user model.User
	// get id params
	id := c.Params("id")
	// find satu pengguna in the database by id
	db.Find(&user, "id = ?", id)
	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}
	var updateUserData updateUser
	// Membaca body request untuk mendapatkan data yang akan diperbarui
	err := c.BodyParser(&updateUserData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// Update fields
	if updateUserData.Nama != "" {
		user.Nama = updateUserData.Nama
	}
	if updateUserData.Golongan != "" {
		user.Golongan = updateUserData.Golongan
	}
	if updateUserData.Jabatan != "" {
		user.Jabatan = updateUserData.Jabatan
	}
	if updateUserData.BidangKeahlian != "" {
		user.BidangKeahlian = updateUserData.BidangKeahlian
	}
	if updateUserData.PendidikanTerakhir != "" {
		user.PendidikanTerakhir= updateUserData.PendidikanTerakhir
	}
	if updateUserData.Email != "" {
		user.Email = updateUserData.Email
	}
	// Save the Changes
	db.Save(&user)
	// Return the updated user
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "users Found", "data": user})
}

// DeleteUserByID: Fungsi untuk menghapus pengguna berdasarkan ID dari database
func DeleteUserByID(c *fiber.Ctx) error {
	db := database.DB.Db
	var user model.User
	// get id params
	id := c.Params("id")
	// find single user in the database by id
	db.Find(&user, "id = ?", id)
	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}
	err := db.Delete(&user, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete user", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User deleted"})
}
