package models

type ShopInfo struct {
	Id      string    `json:"id"`
	UserId  string    `json:"user_id"`
	Name    string    `json:"name"`
	Address Address   `json:"address"`
	Links   []string  `json:"links"`
	Contact []Contact `json:"contact"`
}
