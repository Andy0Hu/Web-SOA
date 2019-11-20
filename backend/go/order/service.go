package order

import (
	"context"
	"errors"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Service interface {
	AllOrders(id string) (interface{}, error)
}

type OrderService struct {
}

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

type AOrders struct {
	User  string   `form:"user" json:"user"`
	Order []Orders `form:"order" json:"order"`
}

func (s OrderService) AllOrders(id string) (interface{}, error) {
	logrus.Info("user: ", id)

	if len(id) == 0 {
		return nil, errors.New("no id specified")
	}
	collection := Client.Database("express").Collection("orders")
	var findRes AOrders
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := collection.FindOne(ctx, bson.M{"user": id}).Decode(&findRes)
	if err != nil {
		logrus.Info("fetch error", err)
		return nil, errors.New("fetch data error")
	}
	logrus.Info(findRes)
	var all []Orders
	all = findRes.Order
	return all, nil
}
