package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Contact struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Number      uint               `json:"number,omitempty" bson:"number,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
}
