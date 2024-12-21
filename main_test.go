package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/html"
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
	body := w.Body.String()
	title, err := extractTitle(body)
	if err != nil {
		t.Fatalf("Ошибка извлечения <title>: %v", err)
	}
	expectedTitle := "Test" // Укажите ожидаемое значение
	if title != expectedTitle {
		t.Errorf("Ожидался <title> %q, но получен %q", expectedTitle, title)
	}
}

// Функция для извлечения <title> из HTML
func extractTitle(htmlContent string) (string, error) {
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return "", err
	}

	// Рекурсивный поиск <title>
	var findTitle func(*html.Node) string
	findTitle = func(n *html.Node) string {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			return n.FirstChild.Data
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if title := findTitle(c); title != "" {
				return title
			}
		}
		return ""
	}

	return findTitle(doc), nil
}
