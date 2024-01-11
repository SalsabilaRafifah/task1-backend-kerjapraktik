// file berisi definisi struktur data model untuk pengguna

package model

import (
	"github.com/google/uuid" // untuk menghasilkan dan bekerja dengan UUID (Universally Unique Identifier)
	"gorm.io/gorm"           // GORM adalah ORM untuk go
)

// struktur data yang mendefinisikan entitas pengguna dengan berbagai field
type User struct {
	gorm.Model         // Menambahkan field-field standar yang biasanya digunakan oleh GORM, seperti ID, CreatedAt, UpdatedAt, dan DeletedAt.
	ID                 uuid.UUID `gorm:"type:uuid;"` // UUID yang digunakan sebagai identifikasi unik untuk setiap pengguna
	Nama               string    `json:"nama"`       // informasi dasar dari seorang dosen
	Golongan           string    `json:"golongan"`
	Jabatan            string    `json:"jabatan"`
	BidangKeahlian     string    `json:"bidang_keahlian"`
	PendidikanTerakhir string    `json:"pendidikan_terakhir"`
	Email              string    `json:"email"`
}

// Users adalah struktur yang digunakan untuk menampung sebuah array berisi kumpulan data dari pengguna. Struktur ini mencakup field Users yang merupakan slice dari struktur User.
type Users struct {
	Users []User `json:"users"`
}

// fungsi hook GORM yang akan dipanggil sebelum entitas User disimpan dalam database.
// kegunaan : menghasilkan UUID baru (versi 4) dan mengatur nilai ID pengguna sebelum entitas dibuat.
func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	return
}

// dapat menggunakan GORM untuk berinteraksi dengan database, membuat, dan mengelola entitas pengguna dengan dukungan UUID untuk identifikasi unik.
