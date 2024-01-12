// file berisi definisi struktur data model untuk dosen

package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// mendefinisikan entitas dosen dengan berbagai field
type Lecturer struct {
	gorm.Model
	ID                 uuid.UUID `gorm:"type:uuid;"`
	Nama               string    `json:"nama"`
	Golongan           string    `json:"golongan"`
	Jabatan            string    `json:"jabatan"`
	BidangKeahlian     string    `json:"bidang_keahlian"`
	PendidikanTerakhir string    `json:"pendidikan_terakhir"`
	Email              string    `json:"email"`
}

// array berisi kumpulan data dari dosen
type Lecturers struct {
	Lecturers []Lecturer `json:"lecturers"`
}

// menghasilkan UUID baru (versi 4) sebagai ID dosen sebelum entitas dibuat.
func (lecturer *Lecturer) BeforeCreate(tx *gorm.DB) (err error) {
	lecturer.ID = uuid.New()
	return
}