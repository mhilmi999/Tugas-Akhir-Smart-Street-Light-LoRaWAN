package models

import "time"

type User struct {
	User_id      int
	Name         string
	Email        string
	Password     string
	Avatar       string
	Date_updated time.Time
	Date_created time.Time
}