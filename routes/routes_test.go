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
		if testCase.expectedPassed && responseRecorder.Code != http.StatusOK {
			t.Fatal("It sould pass in this case but got not OK. Case:", testCase.sid)
		}
		if testCase.expectedPassed && !strings.Contains(resBody, expected) {
			t.Fatalf("Expected: %s, but wrong", testCase.sid)
		}
	}

}

func Test_GetSongsBySearch(t *testing.T) {

	testParams := []struct {
		lang   string
		c      string
		to     string
		title  string
		exCode int
	}{
		{"", "", "", "", http.StatusBadRequest},
		{"Chinese", "10", "G", "前來敬拜", http.StatusOK},
		{"!Chinese", "10", "G", "前來敬拜", http.StatusNotFound},
		{"Chinese", "10!", "G", "前來敬拜", http.StatusNotFound},
		{"Chinese", "10", "G!", "前來敬拜", http.StatusNotFound},
		{"Chinese", "10", "G", "前來敬拜!", http.StatusNotFound},
		{"Chinese", "1", "F", "前來敬拜", http.StatusNotFound},
		{"Taiwanese", "", "", "", http.StatusOK},
		{"", "2", "", "", http.StatusOK},
		{"", "", "Bb", "", http.StatusOK},
		{"", "", "", "愛", http.StatusOK},
		{"-1", "", "", "", http.StatusNotFound},
		{"", "-1", "", "", http.StatusNotFound},
		{"", "", "-1", "", http.StatusNotFound},
		{"", "", "", "-1", http.StatusNotFound},
		{"", "", "", " ", http.StatusBadRequest},
		{"", "", " ", "", http.StatusBadRequest},
		{"", " ", "", "", http.StatusBadRequest},
		{" ", "", "", "", http.StatusBadRequest},
	}

	db.InitDB()

	for caseIndex, testCase := range testParams {
		routePath := fmt.Sprintf("/api/songs/search?lang=%s&c=%s&to=%s&title=%s", testCase.lang, testCase.c, testCase.to, testCase.title)

		request, err := http.NewRequest("GET", routePath, nil)
		if err != nil {
			t.Fatal(err)
		}

		responseRecorder := httptest.NewRecorder()

		testRouter := mux.NewRouter()

		testRouter.HandleFunc("/api/songs/search", GetSongBySearch).Queries("title", "{title}", "lang", "{lang}", "c", "{c}", "to", "{to}").Methods("GET")

		testRouter.ServeHTTP(responseRecorder, request)

		resBody := responseRecorder.Body.String()

		if testCase.exCode != responseRecorder.Code {
			fmt.Println("response:\n", resBody)
			t.Fatalf("Not expected status code, want: %d got: %d in case: %d", testCase.exCode, responseRecorder.Code, caseIndex)
		}
		if testCase.title != "" && testCase.exCode == http.StatusOK && !strings.Contains(resBody, testCase.title) {
			fmt.Println("response:\n", resBody)
			t.Fatalf("Not expected title: %s", testCase.title)
		}
	}

}
