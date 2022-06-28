package main

import (
	"context"
	"log"

	"github.com/adiletelf/authentication-go/internal/config"
	"github.com/adiletelf/authentication-go/internal/handler"
	"github.com/adiletelf/authentication-go/internal/middleware"
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
	defer db.Drop(ctx)
	collection := db.Collection(config.DB.CollectionName)

	ur := repository.NewUserRepo(ctx, collection, client)
	h := handler.New(ur)
	r := gin.Default()

	configureRoutes(r, h)
	r.Run(config.ListenAddress)
}

func configureRoutes(r *gin.Engine, h *handler.Handler) {
	r.POST("/register", h.Register)
	r.POST("/login", h.Login)

	protected := r.Group("/api")
	protected.Use(middleware.JwtAuthMiddleware())

	protected.GET("/", h.HandleHome)

}