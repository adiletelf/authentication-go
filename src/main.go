package main

import (
	"context"
	"log"

	"github.com/adiletelf/authentication-go/internal/config"
	"github.com/adiletelf/authentication-go/internal/handler"
	"github.com/adiletelf/authentication-go/internal/repository"
	"github.com/adiletelf/authentication-go/internal/util"

	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	client, err := util.SetupDB(ctx, config)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	db := client.Database(config.DB.DatabaseName)
	collection := db.Collection(config.DB.CollectionName)
	defer db.Drop(ctx)

	ur := repository.NewUserRepo(ctx, collection, client)
	h := handler.New(ur)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Ok")
	})

	r.POST("/register", h.Register)
	r.POST("/login", h.Login)

	r.Run(config.ListenAddress)
}
