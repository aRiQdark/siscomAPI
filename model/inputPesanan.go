package models

type Inputpesanan struct {
	MemberID        string `json:"memberid"`
	Nopesanan       string
	Membername      string
	Type            string
	DestinationID   string
	DestinationName string
	SkillID         string
	Skillname       string
	Deskripsi       string
	Timestart       string
	Timeend         string
	Price           string
	Recordstatus    string
	Kebutuhan       string
	Sesipengerjaan  string
	Image           string
	Alamat          string
}
