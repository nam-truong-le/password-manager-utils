package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/nam-truong-le/lambda-utils-go/v4/pkg/aws/secretsmanager"
	"github.com/nam-truong-le/lambda-utils-go/v4/pkg/mongodb"
)

const colPasswordsName = "passwords"

func CollectionPasswords(ctx context.Context) (*mongo.Collection, error) {
	database, err := secretsmanager.GetParameter(ctx, "/mongo/db")
	if err != nil {
		return nil, err
	}
	return mongodb.Collection(ctx, database, colPasswordsName)
}

type Password struct {
	ID                primitive.ObjectID `bson:"_id" json:"id"`
	Name              string             `bson:"name" json:"name"`
	Username          string             `bson:"username" json:"username"`
	Password          string             `bson:"password" json:"password"`
	PasswordUpdatedAt time.Time          `bson:"passwordUpdatedAt" json:"passwordUpdatedAt"`
	Note              string             `bson:"note" json:"note"`
	CreatedAt         time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt         time.Time          `bson:"updatedAt" json:"updatedAt"`
}
