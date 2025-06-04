package models

type MemberInput struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Handphone    string
	Fullname     string
	Email        string
	Umur         string
	Alamat       string
	Kota         string
	Provinsi     string
	Role         string
	Kelurahan    string
	Kecamatan    string
	Kodepos      string
	kecamatan    string
	Lastjob      []string
	Skill        []string
	IsLocked     string
	Kodeverif    string
	Birthplace   string
	Birthdate    string
	Nationality  string
	Rating       int
	Emailverifie string
	Profileimage string
	Imei         string
	Limitbalance string
	VAnumber     string
}
