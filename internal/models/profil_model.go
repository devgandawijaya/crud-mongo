package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Profil struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Age       int                `bson:"age" json:"age"`
	Phone     string             `bson:"phone" json:"phone"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}
