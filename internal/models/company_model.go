package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Company struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Address   string             `bson:"address" json:"address"`
	Stores    []Store            `bson:"stores,omitempty" json:"stores,omitempty"`
	Employees []Employee         `bson:"employees,omitempty" json:"employees,omitempty"`
	Customers []Customer         `bson:"customers,omitempty" json:"customers,omitempty"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

type Store struct {
	Name string `bson:"name" json:"name"`
}

type Employee struct {
	Name string `bson:"name" json:"name"`
}

type Customer struct {
	Name string `bson:"name" json:"name"`
}
