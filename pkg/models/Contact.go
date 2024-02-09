package models

type Contact struct {
	Id               string `json:"id"`
	UserId           string `json:"user_id"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	PhoneNumber      string `json:"phone_number"`
	EmailAddress     string `json:"email_address"`
	IsPrimaryContact bool   `json:"is_primary_contract"`
}
