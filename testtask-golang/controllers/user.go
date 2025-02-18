package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"testtask-golang/db"
	"testtask-golang/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Hash password
	if err := user.HashPassword(); err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	// Check if username already exists
	var existingUser models.User
	err := db.GetCollection("users").FindOne(context.Background(), bson.M{"username": user.Username}).Decode(&existingUser)
	if err != mongo.ErrNoDocuments {
		http.Error(w, "Username already exists", http.StatusBadRequest)
		return
	}

	// Insert user
	_, err = db.GetCollection("users").InsertOne(context.Background(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginUser models.User
	if err := json.NewDecoder(r.Body).Decode(&loginUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var dbUser models.User
	err := db.GetCollection("users").FindOne(context.Background(), bson.M{"username": loginUser.Username}).Decode(&dbUser)
	if err == mongo.ErrNoDocuments {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if err := dbUser.ComparePassword(loginUser.Password); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Login successful"})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	cursor, err := db.GetCollection("users").Find(context.Background(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	var users []models.User
	if err = cursor.All(context.Background(), &users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Remove passwords from response
	for i := range users {
		users[i].Password = ""
	}

	json.NewEncoder(w).Encode(users)
}
