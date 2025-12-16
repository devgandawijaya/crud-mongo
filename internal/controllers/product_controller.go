package controllers

import (
	"context"
	"time"

	"crud-mongo/internal/config"
	"crud-mongo/internal/helper"
	"crud-mongo/internal/models"
	"crud-mongo/internal/views"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func productCol() *mongo.Collection {
	return config.MongoDB.Collection("product")
}

func CreateProduct(c *gin.Context) {

	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		var errs []string
		for _, e := range err.(validator.ValidationErrors) {
			errs = append(errs, e.Field()+" is required")
		}
		views.Error(c, errs)
		return
	}

	product.ID = primitive.NewObjectID()
	product.CreatedAt = time.Now()

	_, err := productCol().InsertOne(context.Background(), product)
	if err != nil {
		views.Error(c, err.Error())
		return
	}

	views.Success(c, product)
}

func GetProducts(c *gin.Context) {
	// Ambil parameter dengan default value
	page := helper.Max(1, helper.AtoiOrDefault(c.Query("page"), 1))
	limit := helper.AtoiOrDefault(c.Query("limit"), 10)
	if limit < 1 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	offset := (page - 1) * limit
	ctx := c.Request.Context()

	opts := options.Find().
		SetSkip(int64(offset)).
		SetLimit(int64(limit)).
		SetSort(bson.M{"created_at": -1})

	cursor, err := productCol().Find(ctx, bson.M{}, opts)
	if err != nil {
		views.Error(c, err.Error())
		return
	}
	defer cursor.Close(ctx)

	var products []models.Product
	if err := cursor.All(ctx, &products); err != nil {
		views.Error(c, err.Error())
		return
	}

	total, err := productCol().CountDocuments(ctx, bson.M{})
	if err != nil {
		views.Error(c, err.Error())
		return
	}

	views.Success(c, gin.H{
		"data": products,
		"pagination": gin.H{
			"page":       page,
			"limit":      limit,
			"total":      total,
			"total_page": helper.CalculateTotalPages(total, limit),
		},
	})
}

func GetProductByID(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		views.Error(c, "Invalid ID")
		return
	}

	var product models.Product
	err = productCol().FindOne(context.Background(), bson.M{"_id": id}).Decode(&product)
	if err != nil {
		views.Error(c, "Product not found")
		return
	}

	views.Success(c, product)
}

func UpdateProduct(c *gin.Context) {

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

	_, err = productCol().UpdateByID(context.Background(), id, update)
	if err != nil {
		views.Error(c, err.Error())
		return
	}

	views.Success(c, gin.H{"updated": true})
}

func DeleteProduct(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		views.Error(c, "Invalid ID")
		return
	}

	_, err = productCol().DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		views.Error(c, err.Error())
		return
	}

	views.Success(c, gin.H{"deleted": true})

}
