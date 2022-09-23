package handlers

import (
	"context"
	"net/http"

	hash "urls/pkg/hash"
	redis "urls/pkg/redis"

	"github.com/gin-gonic/gin"
)

type URL struct {
	Url    string `json:"url"`
}

var contx = context.Background()

func GetHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := redis.InitRedisConnect()
		token := ctx.Param("token")
	
		url, err := db.Get(contx, token).Result()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "This url doens't exist",
			})
		} else {
		ctx.Redirect(http.StatusFound, url)
		}
	}
}


func PostHandler() gin.HandlerFunc {
	body := new(URL)
	return func(ctx *gin.Context) {
		ctx.BindJSON(&body)

		urlPayload := URL{
			Url: body.Url,
		}
		data := hash.HashFunction(urlPayload.Url)
		
		isValid := IsValidUrl(urlPayload.Url)

		if !isValid {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Url isn't valid",
			})
			return
		}

		db := redis.InitRedisConnect()
		err := db.Set(contx, data, urlPayload.Url, 0).Err()
		if err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{
				"error": "Cannot write to the database",
			})
			return
		}
		ctx.JSON(http.StatusCreated, "localhost:8080/api/v1/" + data)
	}
}