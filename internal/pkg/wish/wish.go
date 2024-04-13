package wish

import "time"

type Wish struct {
	ID          int64      `gorm:"primaryKey" json:"id"`
	Name        string     `gorm:"not null" json:"name"`
	Description *string    `json:"description"`
	Price       float64    `json:"price"`
	Link        *string    `json:"link"`
	Image       *string    `json:"image"`
	ExpireDate  *time.Time `json:"expire_date"`
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updated_at"`

	// references
	UserID int64 `gorm:"not null" json:"user_id"`
}
