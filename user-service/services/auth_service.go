package services

import (
	"context"
	"errors"
	"time"
	"user-service/config"
	"user-service/models"
	"user-service/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var userCollection = config.GetCollection("users")

func RegisterUser(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	var existing models.User

	err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&existing)
	if err == nil {
		return errors.New("user already exists")
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	user.ID = primitive.NewObjectID()

	_, err = userCollection.InsertOne(ctx, user)
	return err
}

func LoginUser(user *models.User) (string, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	var stored models.User
	err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&stored)
	if err != nil {
		return "", errors.New("user not found")
	}

	if !utils.CheckPasswordHash(user.Password, stored.Password) {
		return "", errors.New("invalid password")
	}

	return "dummy-token", nil
}
