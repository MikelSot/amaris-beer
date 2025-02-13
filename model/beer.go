package model

type Beer struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	ViewCount   int     `json:"view_count"`
}

type Beers []Beer

func (b Beer) HasID() bool { return b.ID > 0 }
