package main

import (
	"context"
	"log"
	"fmt"

	"github.com/adiletelf/authentication-go/internal/config"
	"github.com/adiletelf/authentication-go/internal/model"
	"github.com/adiletelf/authentication-go/internal/repository"
	"github.com/adiletelf/authentication-go/internal/util"
	"github.com/google/uuid"
	// "github.com/gin-gonic/gin"
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
	db := client.Database(config.DatabaseName)
	collection := db.Collection(config.CollectionName)
	defer db.Drop(ctx)

	ur := repository.NewUserRepo(ctx, collection, client)

	user := model.User{
		ID:       uuid.New(),
		Username: "root",
		Password: "root",
	}
	user2 := model.User{
		ID:       uuid.New(),
		Username: "second",
		Password: "second",
	}
	fmt.Println(ur.Save(&user))
	fmt.Println(ur.Save(&user2))

	// h := handler.New(ur)
	// println(h)

	// r := gin.Default()
	// r.GET("/", func(c *gin.Context) {
	// 	c.String(200, "Ok")
	// })
	// r.Run(config.ListenAddress)
}
