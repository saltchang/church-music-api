package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/saltchang/church-music-api/routes"
)

func Test_GetIndex(t *testing.T) {
	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.GetIndex)

	handler.ServeHTTP(responseRecorder, request)

	statusCode := responseRecorder.Code
	if statusCode != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", statusCode, http.StatusOK)
	}

	expected := `"index"` + "\n"
	res := responseRecorder.Body.String()

	if res != expected {
		fmt.Println(len(res), ", ", len(expected))
		t.Errorf("Handler returned unexpected body: got %v want %v", res, expected)
	}
}

func Test_GetAllSongs(t *testing.T) {

	db.InitDB()

	request, err := http.NewRequest("GET", "/api/songs", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.GetAllSongs)

	handler.ServeHTTP(responseRecorder, request)

	statusCode := responseRecorder.Code
	if statusCode != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", statusCode, http.StatusOK)
	}

	resBody := responseRecorder.Body.String()

	expected1 := `"sid":"1001010"`
	expected2 := `"sid":"2007030"`
	if !strings.Contains(resBody, expected1) {
		t.Fatal("Expected 1 wrong.")
	} else if !strings.Contains(resBody, expected2) {
		t.Fatal("Expected 2 wrong.")
	}
}
