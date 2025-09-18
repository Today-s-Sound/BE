package user

import "time"

// User struct는 데이터베이스의 user 테이블과 매핑되는 엔티티입니다.
type User struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name" validate:"required,max=30"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
