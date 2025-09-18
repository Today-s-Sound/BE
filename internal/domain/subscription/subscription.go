package subscription

import "time"

type Subscription struct {
	ID            int64      `json:"id"`
	UserID        int64      `json:"user_id" validate:"required,gt=0"`
	URL           string     `json:"url" validate:"required,url"`
	Alias         string     `json:"alias,omitempty" validate:"omitempty,max=30"`
	KeywordFilter string     `json:"keyword_filter,omitempty" validate:"omitempty,max=100"`
	IsUrgent      bool       `json:"is_urgent"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
}
