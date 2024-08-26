package handlers

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/danilaekimov/short_link/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestRootHandler_GetEndpoint(t *testing.T) {
	shortenerPort := ":8080"
	host := "http://localhost" + shortenerPort
	mockStorage := *storage.GenMockData()
	type want struct {
		code           int
		locationHeader string
	}
	type test struct {
		name     string
		shortURL string
		want     want
	}
	tests := []test{}
	testIndex := 1
	nameTest := "GetEndpoint test "

	//add correct mock data
	for shortURL, webResourse := range mockStorage.ShortURLs {
		tests = append(tests, test{
			name:     nameTest + strconv.Itoa(testIndex),
			shortURL: shortURL,
			want: want{
				code:           http.StatusTemporaryRedirect,
				locationHeader: webResourse.URL,
			},
		})
		testIndex += 1
	}
	// add wrong URL
	tests = append(tests, test{
		name:     nameTest + strconv.Itoa(testIndex) + " wrong data",
		shortURL: "WrongShortUrl",
		want: want{
			code:           http.StatusNotFound,
			locationHeader: "",
		},
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "/"+tt.shortURL, nil)
			w := httptest.NewRecorder()

			h := NewRootHandler(host)
			h.storage = mockStorage

			h.ServeHTTP(w, request)
			res := w.Result()
			// check status code
			assert.Equal(t, tt.want.code, res.StatusCode, "wrong answer code")

			// check Location in header
			assert.Equal(t, tt.want.locationHeader, res.Header.Get("Location"), "wrong answer code")
			defer res.Body.Close()
		})
	}
}

func TestRootHandler_PostEndpoint(t *testing.T) {
	shortenerPort := ":8080"
	host := "http://localhost" + shortenerPort
	type want struct {
		code        int
		result      string
		contentType string
	}
	type test struct {
		name string
		URL  string
		want want
	}
	tests := []test{}
	testIndex := 1
	nameTest := "PostEndpoint test "
	mockStorage := *storage.GenMockData()

	for shortURL, webResourse := range mockStorage.ShortURLs {
		tests = append(tests, test{
			name: nameTest + strconv.Itoa(testIndex),
			URL:  webResourse.URL,
			want: want{
				code:        http.StatusCreated,
				result:      fmt.Sprintf("%s/%s", host, shortURL),
				contentType: "text/plain; charset=utf-8",
			},
		})
		testIndex += 1
	}
	tests = append(tests, test{
		name: nameTest + strconv.Itoa(testIndex) + " wrong data",
		URL:  "wrong url",
		want: want{
			code:        http.StatusBadRequest,
			result:      "Wrong URL given\n",
			contentType: "text/plain; charset=utf-8",
		},
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := strings.NewReader(tt.URL)
			request := httptest.NewRequest(http.MethodPost, "/", body)
			w := httptest.NewRecorder()

			h := NewRootHandler(host)

			h.ServeHTTP(w, request)
			res := w.Result()
			// check status code
			resBody, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.want.code, res.StatusCode, "wrong answer code")

			assert.Equal(t, tt.want.result, string(resBody), "wrong result")

			assert.Equal(t, tt.want.contentType, res.Header.Get("Content-Type"), "wrong content-type")

			defer res.Body.Close()

		})
	}
}
