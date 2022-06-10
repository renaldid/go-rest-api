package domain

import "time"

type Products struct {
	Id         int
	Name       string
	Price      int
	CategoryId int
	CreatedAt  time.Time
	UploadedAt time.Time
}
