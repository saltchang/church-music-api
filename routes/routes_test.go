package routes

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func Test_GetIndex(t *testing.T) {
	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetIndex)

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
	handler := http.HandlerFunc(GetAllSongs)

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

func Test_GetSongsBySID(t *testing.T) {

	testParams := []struct {
		sid            string
		expectedPassed bool
	}{
		{"1010066", true},
		{"2007014", true},
		{"1011050", true},
		{"1000000", false},
		{"", false},
		{"-3", false},
		{"a", false},
		{";", false},
	}

	db.InitDB()

	for _, testCase := range testParams {
		routePath := fmt.Sprintf("/api/songs/sid/%s", testCase.sid)

		request, err := http.NewRequest("GET", routePath, nil)
		if err != nil {
			t.Fatal(err)
		}

		responseRecorder := httptest.NewRecorder()

		testRouter := mux.NewRouter()

		testRouter.HandleFunc("/api/songs/sid/{sid}", GetSongBySID).Methods("GET")

		testRouter.ServeHTTP(responseRecorder, request)

		resBody := responseRecorder.Body.String()

		expected := `"sid":"` + testCase.sid
		if !testCase.expectedPassed && responseRecorder.Code == http.StatusOK {
			t.Fatal("It sould fails in this case but got OK. Case:", testCase.sid)
		}
		if testCase.expectedPassed && !strings.Contains(resBody, expected) {
			t.Fatalf("Expected: %s, but wrong", testCase.sid)
		}
	}

}
