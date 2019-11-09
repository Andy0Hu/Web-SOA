package auth

import (
	"errors"
	"net/http"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type User struct {
	Id     int    `form:"id" json:"id"`
	Name   string `form:"name" json:"name"`
	Passwd string `form:"passwd" json:"passwd"`
}

type LoginRes struct {
	Token string `json:"token"`
	User
}

func Sessions(c *gin.Context) {
	var u User
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"error":  err.Error(),
		})
		return
	}
	name, err := checkPasswd(u.Id, u.Passwd)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  -1,
			"message": err.Error,
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
			"status": -1,
			"error":  err.Error(),
		})
		return
	}
	err := checkExist(u.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  -1,
			"message": "has exists",
		})
	}

	c.JSON(200, gin.H{
		"status":  0,
		"message": "register succeeded!",
	})
}

func checkPasswd(id int, passwd string) (string, error) {
	if id == 111 && passwd == "222" {
		return "333", nil
	}
	return "", errors.New("Not pass")
}

func checkExist(id int) error {
	return nil
}

func genToken(c *gin.Context, user User) {
	j := &JWT{
		[]byte("WEBSOA"),
	}
	claims := CustomClaims{
		user.Id,
		user.Name,
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
