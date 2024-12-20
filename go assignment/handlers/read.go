// handlers/read.go
package handlers

import (
	"encoding/json"
	"net/http"
	"project/db"
	"project/models"

	"go.mongodb.org/mongo-driver/bson"
)

// ReadUser handler
func ReadUser(client *db.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Получаем email из запроса
		email := r.URL.Query().Get("email")
		if email == "" {
			http.Error(w, "Email is required", http.StatusBadRequest)
			return
		}

		// Ищем пользователя по email
		collection := db.GetUserCollection(client)
		var user models.User
		err := collection.FindOne(r.Context(), bson.M{"email": email}).Decode(&user)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		// Отправляем ответ
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}
