package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/Jack-Music-Streaming/server/src/config"
	"github.com/Jack-Music-Streaming/server/src/database"
	"github.com/Jack-Music-Streaming/server/src/models"
)

func InitializeTests() {
	config.GetConfig()

	if err := database.InitDB(); err != nil {
		log.Panic(err)
		return
	}
}

func TestSignUp(t *testing.T) {

	actualUser := &models.User{}

	tests := []Test{
		{
			Name: "SignUp - Should Respond With An User",
			Body: map[string]string{
				"name":     "Rodrigo",
				"password": "123123",
				"email":    "email@gmail.com",
			},
			ExpectedResponse: &models.User{
				Name:  "Rodrigo",
				Email: "email@gmail.com",
			},
			ExpectedStatus: 203,
		},
	}

	for _, test := range tests {

		if actualUser.Name != "" {
			database.DB.Model(&models.User{}).Delete(actualUser)
		}

		t.Run(test.Name, func(t *testing.T) {
			body, err := json.Marshal(&test.Body)

			if err != nil {
				t.Error("body marshal error")
			}

			req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/api/auth/signUp", bytes.NewReader(body))

			if err != nil {
				t.Error("client could not create request")
			}

			req.Header.Set("Content-Type", "application/json")

			client := http.Client{
				Timeout: 15 * time.Second,
			}

			res, err := client.Do(req)

			if err != nil {
				t.Error("error making http request")
			}

			defer res.Body.Close()

			var respBody models.User

			err = json.NewDecoder(res.Body).Decode(&respBody)

			if err != nil {
				fmt.Println(err)
				t.Error("error decoding response")
			}

			if !reflect.DeepEqual(respBody, test.ExpectedResponse) {
				t.Errorf("FAILED: expected %v, got %v\n", test.ExpectedResponse, body)
			}

			if !reflect.DeepEqual(res.StatusCode, test.ExpectedStatus) {
				t.Errorf("FAILED: expected %d, got %d\n", test.ExpectedStatus, res.StatusCode)
			}

		})

	}
}
