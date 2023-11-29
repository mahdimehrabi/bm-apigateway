package entity

type Article struct {
	ID       int64   `json:"ID"`
	Title    string  `json:"title" binding:"required"`
	Body     string  `json:"body" binding:"required"`
	Tags     string  `json:"tags" binding:"required"`
	Price    float64 `json:"price" binding:"required"`
	BuyCount int64   `json:"buyCount" `
	UserID   int64   `json:"userID" binding:"required"`
}
