package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Question struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Question string             `json:"question,omitempty" validate:"required"`
	Answer   string             `json:"answer,omitempty" validate:"required"`
}
