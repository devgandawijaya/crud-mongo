package controllers

import (
	"context"
	"time"

	"crud-mongo/internal/config"
	"crud-mongo/internal/models"
	"crud-mongo/internal/views"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func profilCol() *mongo.Collection {
	return config.MongoDB.Collection("profil")
}

// CREATE
func Createprofil(c *gin.Context) {
	var profil models.Profil

	if err := c.ShouldBindJSON(&profil); err != nil {
		views.Error(c, err.Error())
		return
	}

	profil.ID = primitive.NewObjectID()
	profil.CreatedAt = time.Now()

	_, err := profilCol().InsertOne(context.Background(), profil)
	if err != nil {
		views.Error(c, err.Error())
		return
	}

	views.Success(c, profil)
}

// READ ALL
func GetProfil(c *gin.Context) {
	cursor, err := profilCol().Find(context.Background(), bson.M{})
	if err != nil {
		views.Error(c, err.Error())
		return
	}
	defer cursor.Close(context.Background())

	var profil []models.Profil
	if err := cursor.All(context.Background(), &profil); err != nil {
		views.Error(c, err.Error())
		return
	}

	views.Success(c, profil)
}

// READ BY ID
func GetProfilByID(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		views.Error(c, "Invalid ID")
		return
	}

	var profil models.Profil
	err = profilCol().FindOne(context.Background(), bson.M{"_id": id}).Decode(&profil)
	if err != nil {
		views.Error(c, "profil not found")
		return
	}

	views.Success(c, profil)
}

// UPDATE
func UpdateProfil(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		views.Error(c, "Invalid ID")
		return
	}

	var payload bson.M
	if err := c.ShouldBindJSON(&payload); err != nil {
		views.Error(c, err.Error())
		return
	}

	update := bson.M{
		"$set": payload,
	}

	_, err = profilCol().UpdateByID(context.Background(), id, update)
	if err != nil {
		views.Error(c, err.Error())
		return
	}

	views.Success(c, gin.H{"updated": true})
}

// DELETE
func DeleteProfil(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		views.Error(c, "Invalid ID")
		return
	}

	_, err = profilCol().DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		views.Error(c, err.Error())
		return
	}

	views.Success(c, gin.H{"deleted": true})
}
