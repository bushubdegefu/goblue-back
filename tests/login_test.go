package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert" // add Testify package
)

type PostData struct {
	GrantType string `json:"grant_type"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Token     string `json:"token"`
}

func TestAppsLogin(t *testing.T) {
	// initalaizing the app
	ReturnTestApp()

	// Define a structure for specifying input and output data
	// of a single test case
	tests := []struct {
		name         string   //name of string
		description  string   // description of the test case
		route        string   // route path to test
		expectedCode int      // expected HTTP status code
		postData     PostData // expects post data to the uri
	}{
		// First test case
		{
			name:         "invalid credential login check",
			description:  "get HTTP status 200",
			route:        "/api/v1/login",
			expectedCode: 202,
			postData: PostData{
				GrantType: "authorization_code",
				Email:     "superuser@mail.com",
				Password:  "default@123",
				Token:     "token1",
			},
		},
		// Second test case
		{
			name:         "invalid credential login check",
			description:  "get HTTP status 404, when token does not exist",
			route:        "/api/v1/login",
			expectedCode: 401,
			postData: PostData{
				GrantType: "authorization_code",
				Email:     "superuser@mail.com",
				Password:  "default@12345",
				Token:     "token1",
			},
		},
		{
			name:         "invalid post data login check",
			description:  "get HTTP status 404, when token does not exist",
			route:        "/api/v1/login",
			expectedCode: 400,
			postData: PostData{
				GrantType: "password",
				Email:     "superuser@mail.com",
				Password:  "default@123",
				Token:     "token1",
			},
		},
	}

	// Iterate through test single test cases
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//  changing post data to json
			post_data, _ := json.Marshal(test.postData)

			// Create a new http request with the route from the test case
			req := httptest.NewRequest("POST", test.route, bytes.NewReader(post_data))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("X-APP-TOKEN", "hi")
			resp, _ := TestApp.Test(req)

			if resp.StatusCode == 202 {
				var responseMap map[string]map[string]string
				body, _ := io.ReadAll(resp.Body)
				uerr := json.Unmarshal(body, &responseMap)
				if uerr != nil {
					// fmt.Printf("Error marshaling response : %v", uerr)
					fmt.Println()
				}

				t.Run("Checking Token Decode", func(t *testing.T) {
					post_data_token := PostData{
						GrantType: "token_decode",
						Email:     "superuser@mail.com",
						Password:  "default@123",
						Token:     responseMap["data"]["access_token"],
					}

					post_data_string, _ := json.Marshal(post_data_token)
					//  checking token decode options
					req := httptest.NewRequest("POST", test.route, bytes.NewReader(post_data_string))
					req.Header.Set("Content-Type", "application/json")
					req.Header.Set("X-APP-TOKEN", "hi")
					// checking refresh_token options
					resp, _ := TestApp.Test(req)
					assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
				})

				t.Run("Checking Refresh Token", func(t *testing.T) {
					post_data_token := PostData{
						GrantType: "refresh_token",
						Email:     "superuser@mail.com",
						Password:  "default@123",
						Token:     responseMap["data"]["access_token"],
					}

					post_data_string, _ := json.Marshal(post_data_token)
					//  checking token decode options
					req := httptest.NewRequest("POST", test.route, bytes.NewReader(post_data_string))
					req.Header.Set("Content-Type", "application/json")
					req.Header.Set("X-APP-TOKEN", "hi")
					// checking refresh_token options
					resp, _ := TestApp.Test(req)
					assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
				})
			}
			// This comment is kept for debuging purpose when making changes to the endpoint
			// fmt.Println(resp.StatusCode)
			// fmt.Println(string(body))

			// Verify, if the status code is as expected
			assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)

		})

	}

}

func TestAppsCheckLogin(t *testing.T) {
	// initalaizing the app
	ReturnTestApp()

	// Define a structure for specifying input and output data
	// of a single test case
	tests := []struct {
		name         string   //name of string
		description  string   // description of the test case
		route        string   // route path to test
		expectedCode int      // expected HTTP status code
		postData     PostData // expects post data to the uri
	}{
		// First test case
		{
			name:         "credential login check",
			description:  "get HTTP status 200",
			route:        "/api/v1/login",
			expectedCode: 202,
			postData: PostData{
				GrantType: "authorization_code",
				Email:     "superuser@mail.com",
				Password:  "default@123",
				Token:     "token1",
			},
		},
	}

	// Iterate through test single test cases
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//  changing post data to json
			post_data, _ := json.Marshal(test.postData)

			// Create a new http request with the route from the test case
			req := httptest.NewRequest("POST", test.route, bytes.NewReader(post_data))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("X-APP-TOKEN", "hi")
			resp, _ := TestApp.Test(req)

			if resp.StatusCode == 202 {
				var responseMap map[string]interface{}
				body, _ := io.ReadAll(resp.Body)
				uerr := json.Unmarshal(body, &responseMap)
				if uerr != nil {
					// fmt.Printf("Error marshaling response : %v", uerr)
					fmt.Println()
				}

				t.Run("Checking Login Status", func(t *testing.T) {
					token := responseMap["data"].(map[string]interface{})["access_token"]
					//  checking token decode options
					req := httptest.NewRequest("GET", "/api/v1/checklogin", nil)
					req.Header.Set("Content-Type", "application/json")
					req.Header.Set("X-APP-TOKEN", token.(string))
					// checking refresh_token options
					resp, _ := TestApp.Test(req)
					assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
				})

			}
			// This comment is kept for debuging purpose when making changes to the endpoint
			// fmt.Println(resp.StatusCode)
			// fmt.Println(string(body))

			// Verify, if the status code is as expected
			assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)

		})

	}

}
