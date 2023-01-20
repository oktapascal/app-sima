package domain

import "time"

type Session struct {
	AuthToken string     `firestore:"auth_token"`
	CreatedAt time.Time  `firestore:"created_at"`
	DeletedAt *time.Time `firestore:"deleted_at"`
}
