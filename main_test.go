package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestWebServer(t *testing.T) {
	// Создаем тестовый сервер Gin
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.LoadHTMLGlob("HTML/*.html")
	r.GET("/", handlerIndex)

	// Создаем тестовый HTTP-запрос
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Не удалось создать запрос: %v", err)
	}

	// Создаем для получения ответа
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Проверяем код ответа
	if w.Code != http.StatusOK {
		t.Errorf("Ожидался код %d, но получен %d", http.StatusOK, w.Code)
	}

	// Проверяем содержимое ответа (HTML-код)
	expected := "" // Ожидаемое содержимое
	if w.Body.String() != expected {
		t.Errorf("Ожидался ответ %q, но получен %q", expected, w.Body.String())
	}
}
