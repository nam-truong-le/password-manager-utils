package pkg

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Password struct {
	ID    primitive.ObjectID `bson:"_id" json:"id"`
	Name  string             `bson:"name" json:"name"`
	Value string             `bson:"value" json:"value"`
}
