package models

import "time"

type Person struct {
	ID          uint
	Name        string
	Age         uint8
	Communities Communities
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Persons []*Person
