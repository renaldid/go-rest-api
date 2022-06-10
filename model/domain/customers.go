package domain

import "time"

type Customers struct {
	Id          int
	Name        string
	Address     string
	Email       string
	PhoneNumber string
	CreatedAt   time.Time
	UploadedAt  time.Time
}
