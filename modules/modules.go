package modules

import (
	"context"
	"fmt"
	"log"
	"time"

	"mongoDB-Golang/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var client *mongo.Client

func ConnectDatabase() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}

func GetCollection(database string, collection string) *mongo.Collection {
	return client.Database(database).Collection(collection)
}

func CreateInitialAdmin(userCollection *mongo.Collection) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"username": "admin"}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
		admin := models.User{
			ID:       primitive.NewObjectID(),
			Username: "admin",
			Password: string(hashedPassword),
			Role:     "admin",
		}
		userCollection.InsertOne(ctx, admin)
		fmt.Println("Initial admin user created")
	} else {
		fmt.Println("Admin user already exists")
	}
}
