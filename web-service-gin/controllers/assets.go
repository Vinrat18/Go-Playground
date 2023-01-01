package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	"web-service-gin/schema"
	"web-service-gin/service"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var assetCollection *mongo.Collection = service.GetCollection(service.DB, "assets")

func errorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetAssets(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	assetSource := c.Param("source")
	fmt.Println(assetSource)
	var assets []schema.Asset
	defer cancel()

	cur, err := assetCollection.Find(ctx, bson.M{"source": assetSource})
	for cur.Next(ctx) {
		var asset schema.Asset
		errorHandler(cur.Decode(&asset))
		assets = append(assets, asset)
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, assets)
}

func GetAsset(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	assetId := c.Param("id")
	assetSource := c.Param("source")
	var asset schema.Asset
	defer cancel()

	err := assetCollection.FindOne(ctx, bson.M{"id": assetId, "source": assetSource}).Decode(&asset)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, asset)
}

func CreateUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var asset schema.Asset
	defer cancel()

	//validate the request body
	if err := c.BindJSON(&asset); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	result, err := assetCollection.InsertOne(ctx, asset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)

	}

	c.JSON(http.StatusOK, result)
}
