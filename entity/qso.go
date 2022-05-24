package entity

type Qso struct {
	Band          string `json:"band"`
	Mode          string `json:"mode"`
	Date          string `json:"date"`
	Time          string `json:"time"`
	OurCall       string `json:"ourCall"`
	OurCategory   string `json:"ourCategory"`
	OurSection    string `json:"ourSection"`
	TheirCall     string `json:"theirCall"`
	TheirCategory string `json:"theirCategory"`
	TheirSection  string `json:"theirSection"`
}
