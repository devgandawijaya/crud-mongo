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

func companyCol() *mongo.Collection {
	return config.MongoDB.Collection("companies")
}

// CREATE
func CreateCompany(c *gin.Context) {
	var company models.Company

	if err := c.ShouldBindJSON(&company); err != nil {
		views.Error(c, err.Error())
		return
	}

	company.ID = primitive.NewObjectID()
	company.CreatedAt = time.Now()

	_, err := companyCol().InsertOne(context.Background(), company)
	if err != nil {
		views.Error(c, err.Error())
		return
	}

	views.Success(c, company)
}

// READ ALL
func GetCompanies(c *gin.Context) {
	cursor, err := companyCol().Find(context.Background(), bson.M{})
	if err != nil {
		views.Error(c, err.Error())
		return
	}
	defer cursor.Close(context.Background())

	var companies []models.Company
	if err := cursor.All(context.Background(), &companies); err != nil {
		views.Error(c, err.Error())
		return
	}

	views.Success(c, companies)
}

// READ BY ID
func GetCompanyByID(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		views.Error(c, "Invalid ID")
		return
	}

	var company models.Company
	err = companyCol().FindOne(context.Background(), bson.M{"_id": id}).Decode(&company)
	if err != nil {
		views.Error(c, "Company not found")
		return
	}

	views.Success(c, company)
}

// UPDATE
func UpdateCompany(c *gin.Context) {
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

	_, err = companyCol().UpdateByID(context.Background(), id, update)
	if err != nil {
		views.Error(c, err.Error())
		return
	}

	views.Success(c, gin.H{"updated": true})
}

// DELETE
func DeleteCompany(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		views.Error(c, "Invalid ID")
		return
	}

	_, err = companyCol().DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		views.Error(c, err.Error())
		return
	}

	views.Success(c, gin.H{"deleted": true})
}
