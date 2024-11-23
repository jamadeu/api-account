package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/jamadeu/api-account/models"
	"gorm.io/gorm"
)

var today = time.Now()
var userTest = models.User{
	Model: gorm.Model{
		ID:        1,
		CreatedAt: today,
		UpdatedAt: today,
	},
	Name:      "User Test",
	Document:  "1234567890",
	Email:     "test@test.com",
	AccountID: 1,
}
var accountTest = models.Account{
	Model: gorm.Model{
		ID:        1,
		CreatedAt: today,
		UpdatedAt: today,
	},
	User:    userTest,
	Balance: 1000,
}

func TestAccountHandler(t *testing.T) {
	ar := &MockAccountRepository{}
	h := NewAccountHandler(ar)
	r := gin.Default()
	h.RegisterRoutes(r, "/api/v1")

	t.Run("account handler should get account by ID", func(t *testing.T) {
		w := httptest.NewRecorder()
		accountId := strconv.Itoa(int(accountTest.ID))
		req, err := http.NewRequest("GET", "/api/v1/account?id="+accountId, nil)
		if err != nil {
			t.Fatal(err)
		}
		r.ServeHTTP(w, req)

		expectedBody := "{" +
			"\"data\":" + jsonToString(accountTest) + "," +
			"\"message\":\"operation from handler: find-account-by-id successfull\"" +
			"}"

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, expectedBody, w.Body.String())
	})

	t.Run("handle find should return 400 when accountId is nil", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/api/v1/account?id=", nil)
		if err != nil {
			t.Fatal(err)
		}
		r.ServeHTTP(w, req)

		expectedResponseBody := "{\"errorCode\":400,\"message\":\"id queryParameter is required\"}"
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, expectedResponseBody, w.Body.String())
	})

	t.Run("handle find should return 404 when account is not found", func(t *testing.T) {
		w := httptest.NewRecorder()
		accountId := "2"
		req, err := http.NewRequest("GET", "/api/v1/account?id="+accountId, nil)
		if err != nil {
			t.Fatal(err)
		}
		r.ServeHTTP(w, req)

		expectedResponseBody := "{\"errorCode\":404,\"message\":\"account with id: " + accountId + " not found\"}"
		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Equal(t, expectedResponseBody, w.Body.String())
	})

}

type MockAccountRepository struct{}

func (m *MockAccountRepository) FindById(id string) (*models.Account, error) {
	if id == "1" {
		return &accountTest, nil
	} else {
		return nil, errors.New("user not found")
	}
}

func jsonToString(s interface{}) string {
	b, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return string(b)
}
