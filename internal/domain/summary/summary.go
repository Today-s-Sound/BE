package summary

import "time"

// Summary struct는 데이터베이스의 summary 테이블과 매핑되는 엔티티입니다.
type Summary struct {
	ID             int64      `json:"id"`
	SubscriptionID int64      `json:"subscription_id" validate:"required,gt=0"`
	Hash           string     `json:"hash" validate:"required"`
	Summary        string     `json:"summary" validate:"required"`
	IsRead         bool       `json:"is_read"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`
}
