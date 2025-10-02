package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// ConnectMongoDB establishes a connection to MongoDB
func ConnectMongoDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:example@localhost:27017"))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Ping the database to verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	fmt.Println("Connected to MongoDB!")
}

// AddUser inserts a new user into the database
func AddUser(username, password string) {
	collection := client.Database("testdb").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, map[string]interface{}{
		"username": username,
		"password": password,
	})
	if err != nil {
		log.Fatalf("Failed to insert user: %v", err)
	}

	fmt.Println("User added successfully")
}

// Login validates user credentials
func Login(username, password string) bool {
	collection := client.Database("testdb").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var result map[string]interface{}
	err := collection.FindOne(ctx, map[string]interface{}{
		"username": username,
		"password": password,
	}).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("Invalid credentials")
			return false
		}
		log.Fatalf("Error finding user: %v", err)
	}

	fmt.Println("Login successful")
	return true
}
