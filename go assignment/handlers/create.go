package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"project/models"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Логирование входящего запроса
		log.Println("Received request to create a patient")

		// Убедитесь, что запрос POST
		if r.Method != http.MethodPost {
			log.Println("Invalid request method:", r.Method)
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var patient models.Patient
		// Декодируем JSON тело запроса в структуру patient
		err := json.NewDecoder(r.Body).Decode(&patient)
		if err != nil {
			log.Println("Failed to parse request body:", err) // Логирование ошибки при парсинге
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
			return
		}

		// Подключение к коллекции MongoDB
		collection := client.Database("healthcare").Collection("patients")

		// Вставка данных пациента в MongoDB
		_, err = collection.InsertOne(r.Context(), patient)
		if err != nil {
			log.Println("Failed to insert patient data:", err) // Логирование ошибки при вставке
			http.Error(w, "Failed to insert patient data", http.StatusInternalServerError)
			return
		}

		// Логирование успешного добавления
		log.Println("Patient data successfully inserted:", patient)

		// Ответ клиенту
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(patient)
	}
}
