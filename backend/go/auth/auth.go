package auth

import (
	"context"
	"errors"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

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

func Sessions(c *gin.Context) {
	var u User
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  -1,
			"message": err.Error(),
		})
		return
	}
	if len(u.Id) == 0 || len(u.Password) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  -1,
			"message": "empty input",
		})
		return
	}

	name, err := checkPasswd(u.Id, u.Password)
	u.Username = name
	if err != nil {
		logrus.Info(err)
		c.JSON(200, gin.H{
			"status":  -1,
			"message": "Check Password error: ",
		})
	} else {
		genToken(c, u)
	}

	c.JSON(200, gin.H{
		"status": 0,
		"name":   name,
	})
}

func Users(c *gin.Context) {
	var u User
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  -1,
			"message": err.Error(),
		})
		return
	}

	if len(u.Id) == 0 || len(u.Username) == 0 || len(u.Password) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  -1,
			"message": "empty input",
		})
		return
	}
	err := checkExist(u.Id)
	if err != nil {
		logrus.Info("exists: ", err)
		c.JSON(200, gin.H{
			"status":  -1,
			"message": "has exists",
		})
		return
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
		c.JSON(200, gin.H{
			"status":  -1,
			"message": "register failed!",
		})
		return
	}

	logrus.Info("insert ID: ", res.InsertedID)
	c.JSON(200, gin.H{
		"status":  0,
		"message": "register succeeded!",
	})
}

func TokenRefresh(c *gin.Context, user User) {
	j := &JWT{
		[]byte("WEBSOA"),
	}
	tokenToRefresh := c.GetHeader("token")
	token, err := j.TokenRefresh(tokenToRefresh)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  -1,
			"message": err.Error(),
		})
	}
	logrus.Info(token)

	data := LoginRes{
		Token: token,
		User:  user,
	}
	c.JSON(200, gin.H{
		"status":  0,
		"message": "TokenRefresh Succeeded!",
		"data":    data,
	})

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

func genToken(c *gin.Context, user User) {
	j := &JWT{
		[]byte("WEBSOA"),
	}
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
		c.JSON(200, gin.H{
			"status":  -1,
			"message": err.Error(),
		})
	}

	logrus.Info(token)

	data := LoginRes{
		Token: token,
		User:  user,
	}
	c.JSON(200, gin.H{
		"status":  0,
		"message": "Login Succeeded!",
		"data":    data,
	})
}
