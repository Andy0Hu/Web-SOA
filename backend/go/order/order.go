package order

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func init() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	Client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
}

type Orders struct {
	Id       string `form:"id" json:"id"`
	Content  string `form:"content" json:"content"`
	Position string `form:"position" json:"position"`
	Signed   bool   `form:"signed" json:"signed"`
}

type User struct {
	Id string `form:"id" json:"id"`
}

type AOrders struct {
	User  string   `form:"user" json:"user"`
	Order []Orders `form:"order" json:"order"`
}

func AllOrders(c *gin.Context) {
	logrus.Info("ALL_ORDERS")
	var u User
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  -1,
			"message": err.Error(),
		})
		return
	}
	logrus.Info("user: ", u)

	if len(u.Id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  -1,
			"message": "empty input",
		})
		return
	}
	collection := Client.Database("express").Collection("orders")
	var findRes AOrders
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := collection.FindOne(ctx, bson.M{"user": u.Id}).Decode(&findRes)
	if err != nil {
		logrus.Info("fetch error", err)
		c.JSON(200, gin.H{
			"status":  -1,
			"message": "fetch errors",
		})
		return
	}
	logrus.Info(findRes)
	var all []Orders
	all = findRes.Order
	c.JSON(200, gin.H{
		"status":  0,
		"message": all,
	})
}
