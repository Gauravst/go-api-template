package dto

import (
	"time"
)

type UserResponse struct {
	Id         string    `json:"id"`
	Name       *string   `json:"name,omitempty"`
	Username   *string   `json:"username,omitempty"`
	Email      string    `json:"email"`
	Role       string    `json:"role"`
	ProfilePic *string   `json:"profile_pic"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
