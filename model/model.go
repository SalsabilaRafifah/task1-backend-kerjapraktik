// file berisi definisi struktur data model untuk dosen

package model

import (
	"github.com/google/uuid" // untuk menghasilkan dan bekerja dengan UUID (Universally Unique Identifier)
	"gorm.io/gorm"           // GORM adalah ORM untuk go
)

// struktur data yang mendefinisikan entitas dosen dengan berbagai field
type Lecturer struct {
	gorm.Model                   // Menambahkan field-field standar yang biasanya digunakan oleh GORM, seperti ID, CreatedAt, UpdatedAt, dan DeletedAt.
	ID                 uuid.UUID `gorm:"type:uuid;"` // UUID yang digunakan sebagai identifikasi unik untuk setiap dosen
	Nama               string    `json:"nama"`       // informasi dasar dari seorang dosen
	Golongan           string    `json:"golongan"`
	Jabatan            string    `json:"jabatan"`
	BidangKeahlian     string    `json:"bidang_keahlian"`
	PendidikanTerakhir string    `json:"pendidikan_terakhir"`
	Email              string    `json:"email"`
}

// Lecturers adalah struktur yang digunakan untuk menampung sebuah array berisi kumpulan data dari dosen. Struktur ini mencakup field Users yang merupakan slice dari struktur User.
type Lecturers struct {
	Lecturers []Lecturer `json:"lecturers"`
}

// fungsi hook GORM yang akan dipanggil sebelum entitas User disimpan dalam database.
// kegunaan : menghasilkan UUID baru (versi 4) dan mengatur nilai ID dosen sebelum entitas dibuat.
func (lecturer *Lecturer) BeforeCreate(tx *gorm.DB) (err error) {
	lecturer.ID = uuid.New()
	return
}

// dapat menggunakan GORM untuk berinteraksi dengan database, membuat, dan mengelola entitas dosen dengan dukungan UUID untuk identifikasi unik.
