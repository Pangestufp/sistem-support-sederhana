package entity

import "time"

type UserRole struct {
	UserID     int
	RoleID     int
	AssignedAt time.Time
}
