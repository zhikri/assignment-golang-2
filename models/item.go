package models

type Item struct {
	ID          uint   `gorm:"primaryKey" json:"-"`
	OrderID     uint   `json:"-"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
