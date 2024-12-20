package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestData struct {
	Message string `json:"message"`
}

type ResponseData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/api", handler)
	fmt.Println("Сервер запущен на порту 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodPost {
		var requestData RequestData

		// Декодируем JSON из тела запроса
		if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
			respondWithError(w, "Некорректное JSON-сообщение")
			return
		}

		// Проверяем наличие поля "message" и его тип
		if requestData.Message == "" {
			respondWithError(w, "Некорректное JSON-сообщение")
			return
		}

		// Выводим сообщение в консоль
		fmt.Println("Получено сообщение:", requestData.Message)

		// Отправляем успешный ответ
		respondWithSuccess(w, "Data successfully received")
		return
	}

	if r.Method == http.MethodGet {
		respondWithSuccess(w, "GET запрос успешен")
		return
	}

	respondWithError(w, "Invalid JSON message")
}

func respondWithSuccess(w http.ResponseWriter, message string) {
	response := ResponseData{
		Status:  "success",
		Message: message,
	}
	json.NewEncoder(w).Encode(response)
}

func respondWithError(w http.ResponseWriter, message string) {
	response := ResponseData{
		Status:  "fail",
		Message: message,
	}
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(response)
}
