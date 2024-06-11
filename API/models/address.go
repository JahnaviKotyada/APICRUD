package models

type Address struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	UserID  int    `json:"user_id"`
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
}
