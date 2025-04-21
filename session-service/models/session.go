package models

import "time"

type Session struct {
	UserID    string    `bson:"user_id"`
	Email     string    `bson:"email"`
	Source    string    `bson:"source"`
	IP        string    `bson:"ip"`
	Timestamp time.Time `bson:"timestamp"`
}
