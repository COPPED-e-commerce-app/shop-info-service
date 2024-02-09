package models

type Address struct {
	Id            string `json:"id"`
	UserId        string `json:"user_id"`
	StreetAddress string `json:"street_address"`
	UnitNumber    string `json:"unit_number`
	ZipCode       string `json:"zip_code"`
	State         string `json:"state"`
}
