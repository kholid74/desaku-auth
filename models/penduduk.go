package models

import (
	"gorm.io/gorm"
)

type Penduduk struct {
	gorm.Model
	Nik           string `json:"name"`
	No_kk         string `json:"email" gorm:"unique"`
	Nama          string `json:"password"`
	Jenis_kelamin string `json:"jenis_kelamin"`
	Tempat_lahir  string `json:"tempat_lahir"`
	Tgl_lahir     string `json:"tgl_lahir"`
	Agama         string `json:"agama"`
	Pendidikan    string `json:"pendidikan"`
	Pekerjaan     string `json:"pekerjaan"`
	Alamat        string `json:"alamat"`
	Rt            int    `json:"rt"`
	Rw            int    `json:"rw"`
	Kelurahan     string `json:"kelurahan"`
	Kecamatan     string `json:"kecamatan"`
	Kota          string `json:"kota"`
	Provinsi      string `json:"provinsi"`
	Kodepos       int    `json:"kodepos"`
	Negara        string `json:"negara"`
	Status_nikah  string `json:"status_nikah"`
	Qr_code       string `json:"qr_code"`
}
