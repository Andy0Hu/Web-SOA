package auth

import (
	"context"
	"errors"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Service Define a service interface
type Service interface {
	Login(id, pwd string) (string, error)
	Register(id, pwd, username string) error
}

type AuthService struct {
}

var Client *mongo.Client

func init() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	Client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
}

type User struct {
	Id       string `form:"id" json:"id"`
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

type LoginRes struct {
	Token string `json:"token"`
	User
}

func (s AuthService) Login(id, pwd string) (string, error) {
	u := User{
		Id:       id,
		Password: pwd,
	}
	if len(id) == 0 || len(pwd) == 0 {
		return "", errors.New("empty id or password")
	}

	var token string
	name, err := checkPasswd(id, pwd)
	if err != nil {
		logrus.Info(err)
		return "", errors.New("password not match")
	} else {
		u.Username = name
		token = genToken(u)
	}

	return token, nil
}

func (s AuthService) Register(id, pwd, username string) error {
	u := User{
		Id:       id,
		Password: pwd,
		Username: username,
	}
	if len(u.Id) == 0 || len(u.Username) == 0 || len(u.Password) == 0 {
		return errors.New("empty input")
	}
	err := checkExist(u.Id)
	if err != nil {
		logrus.Info("exists: ", err)
		return errors.New("the user has existed")
	}

	collection := Client.Database("express").Collection("user")
	newUser := bson.M{
		"id":       u.Id,
		"password": u.Password,
		"username": u.Username,
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertOne(ctx, newUser)

	if err != nil {
		logrus.Info("insert new user error\n", err)
		return errors.New("register failed!")
	}

	logrus.Info("insert ID: ", res.InsertedID)
	return nil
}

func checkPasswd(id string, passwd string) (string, error) {
	collection := Client.Database("express").Collection("user")
	var user User

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := collection.FindOne(ctx, bson.M{"id": id}).Decode(&user)
	if err != nil {
		return "", err
	}
	if err != nil {
		return "", err
	}
	if passwd == user.Password {
		return user.Username, nil
	}
	return "", errors.New("password error")
}

func checkExist(id string) error {
	collection := Client.Database("express").Collection("user")
	var user User
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := collection.FindOne(ctx, bson.M{"id": id}).Decode(&user)
	if err == nil {
		logrus.Info("exists: ", err)
		return errors.New("exists")
	}
	return nil
}

func genToken(user User) string {
	j := NewJWT()
	claims := CustomClaims{
		user.Id,
		user.Username,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),
			ExpiresAt: int64(time.Now().Unix() + 3600),
			Issuer:    "WEBSOA",
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
	}

	logrus.Info(token)

	return token
}
