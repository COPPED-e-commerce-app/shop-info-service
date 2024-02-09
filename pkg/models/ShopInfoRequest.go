package models

type ShopInfoRequest struct {
	UserId  string    `json:"user_id"`
	Name    string    `json:"name"`
	Address Address   `json:"address"`
	Links   []string  `json:"links"`
	Contact []Contact `json:"contact"`
}
