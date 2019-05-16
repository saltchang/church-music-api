package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Token struct
type Token struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Token      string             `json:"token" bson:"token"`
	Autho      string             `json:"autho" bson:"autho"`
	CreateDate int                `json:"create_date" bson:"create_date"`
}
