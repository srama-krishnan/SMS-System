package models

import "time"

type Message struct {
	ID          string    `json:"id" bson:"id"`
	UserID      string    `json:"userId" bson:"userId"`
	PhoneNumber string    `json:"phoneNumber" bson:"phoneNumber"`
	Text        string    `json:"text" bson:"text"`
	Status      string    `json:"status" bson:"status"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
}
