package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"session-service/database"

	"go.mongodb.org/mongo-driver/bson"
)

func GetSessions(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	collection := database.GetCollection()

	cursor, err := collection.Find(context.TODO(), bson.M{"email": email})
	if err != nil {
		http.Error(w, "Failed to get sessions", http.StatusInternalServerError)
		return
	}

	defer cursor.Close(context.Background())

	var sessions []bson.M
	if err = cursor.All(context.Background(), &sessions); err != nil {
		http.Error(w, "Cursor error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(sessions)
}
