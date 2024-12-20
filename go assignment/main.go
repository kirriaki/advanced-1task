package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var collection *mongo.Collection

func main() {
	// Подключаемся к MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Подключаемся к базе данных med_portal и коллекции patients
	collection = client.Database("med_portal").Collection("patients")

	// Обработчик для добавления пациента
	http.HandleFunc("/add-patient", addPatient)

	// Запуск сервера
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func addPatient(w http.ResponseWriter, r *http.Request) {
	// Только POST-запросы
	if r.Method == http.MethodPost {
		var patient Patient
		// Чтение данных из тела запроса
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&patient); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Логируем полученные данные
		log.Println("Received patient data:", patient)

		// Вставка данных в MongoDB
		_, err := collection.InsertOne(context.TODO(), patient)
		if err != nil {
			http.Error(w, "Error saving patient data", http.StatusInternalServerError)
			return
		}

		// Успешный ответ
		w.Write([]byte("Patient added successfully"))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

type Patient struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}
