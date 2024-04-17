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

type userpatch struct {
	Email    string `json:"email" example:"someone@domain.com"`
	Disabled bool   `json:"disabled" example:"true"`
}

type userpost struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func TestAppUser(t *testing.T) {
	// initalaizing the app
	ReturnTestApp()

	// Define a structure for specifying input and output data
	// of a single test case
	tests_user := []struct {
		name         string //name of string
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		// First test case
		{
			name:         "get users check",
			description:  "get HTTP status 200",
			route:        "/api/v1/users?page=1&size=10",
			expectedCode: 200,
		},
		// Second test case
		{
			name:         "get users check",
			description:  "get HTTP status 404, when token does not exist",
			route:        "/api/v1/users?page=1&size=10",
			expectedCode: 200,
		},
		{
			name:         "get users by id",
			description:  "get HTTP status 404, when token does not exist",
			route:        "/api/v1/users/1",
			expectedCode: 200,
		},
		{
			name:         "get non existent user ",
			description:  "get HTTP status 404, when token does not exist",
			route:        "/api/v1/users/50",
			expectedCode: 503,
		},
	}

	// Iterate through test single test cases
	for _, test := range tests_user {
		t.Run(test.name, func(t *testing.T) {
			// Create a new http request with the route from the test case
			req := httptest.NewRequest("GET", test.route, nil)
			req.Header.Set("X-APP-TOKEN", "hi")
			resp, _ := TestApp.Test(req)

			// This comment is kept for debuging purpose when making changes to the endpoint
			// fmt.Println(resp.StatusCode)
			// body, _ := io.ReadAll(resp.Body)
			// fmt.Println(string(body))

			// Verify, if the status code is as expected
			assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)

		})

	}

}

func TestAppUserOperations(t *testing.T) {
	// initalaizing the app
	ReturnTestApp()

	// Define a structure for specifying input and output data
	// of a single test case
	tests_user := []struct {
		name         string //name of string
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
		postData     userpost
	}{
		// First test case
		{
			name:         "post users check new post and unique trail one",
			description:  "get HTTP status 200",
			route:        "/api/v1/users",
			expectedCode: 200,
			postData: userpost{
				Email:    "testaddone@mail.com",
				Password: "default@123",
			},
		},
		// Second test case
		{
			name:         "post users check new post and unique trail two",
			description:  "get HTTP status 404, when token does not exist",
			route:        "/api/v1/users",
			expectedCode: 200,
			postData: userpost{
				Email:    "testaddtwo@mail.com",
				Password: "default@123",
			},
		},
		{
			name:         "post users unique user check",
			description:  "get HTTP status 404, when token does not exist",
			route:        "/api/v1/users",
			expectedCode: 200,
			postData: userpost{
				Email:    "testaddthree@mail.com",
				Password: "default@123",
			},
		},
	}

	// Iterate through test single test cases
	for _, test := range tests_user {
		t.Run(test.name, func(t *testing.T) {

			//  changing post data to json
			post_data, _ := json.Marshal(test.postData)

			// Create a new http request with the route from the test case
			req := httptest.NewRequest("POST", test.route, bytes.NewReader(post_data))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("X-APP-TOKEN", "hi")
			resp, _ := TestApp.Test(req)

			var responseMap map[string]interface{}
			body, _ := io.ReadAll(resp.Body)
			uerr := json.Unmarshal(body, &responseMap)
			if uerr != nil {
				// fmt.Printf("Error marshaling response : %v", uerr)
				fmt.Println()
			}
			t.Run("Checkinng Unique Constraint", func(t *testing.T) {
				//  checking token decode options
				req := httptest.NewRequest("POST", test.route, bytes.NewReader(post_data))
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("X-APP-TOKEN", "hi")
				// checking refresh_token options
				resp, _ := TestApp.Test(req)
				assert.Equalf(t, 500, resp.StatusCode, "Checking unique constraint")
			})

			t.Run("Checkinng Patch User Constraint", func(t *testing.T) {

				// creating patchdata
				patch_data_token := userpatch{
					Email:    test.postData.Email,
					Disabled: true,
				}

				patch_data_string, _ := json.Marshal(patch_data_token)

				// creating path
				test_route := fmt.Sprintf("%v/%v", test.route, responseMap["data"].(map[string]interface{})["id"])

				//  checking token decode options
				req := httptest.NewRequest("PATCH", test_route, bytes.NewReader(patch_data_string))
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("X-APP-TOKEN", "hi")
				// checking refresh_token options
				resp, _ := TestApp.Test(req)
				assert.Equalf(t, 200, resp.StatusCode, "Checking unique constraint")
			})

			t.Run("Checkinng Activate Deactivate User", func(t *testing.T) {

				test_route := fmt.Sprintf("%v/%v?status=true", test.route, responseMap["data"].(map[string]interface{})["id"])
				//  checking token decode options
				req := httptest.NewRequest("PUT", test_route, nil)
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("X-APP-TOKEN", "hi")
				// checking Delete object options
				resp, _ := TestApp.Test(req)
				assert.Equalf(t, 200, resp.StatusCode, "Checking Activate/Deactivate User")
			})

			t.Run("Checkinng Reset Password", func(t *testing.T) {

				// creating patchdata
				patch_data_token := userpost{
					Email:    test.postData.Email,
					Password: "default@12345",
				}

				patch_data_string, _ := json.Marshal(patch_data_token)

				test_route := fmt.Sprintf("%v/%v?reset=true", test.route, responseMap["data"].(map[string]interface{})["id"])
				//  checking token decode options
				req := httptest.NewRequest("PATCH", test_route, bytes.NewReader(patch_data_string))
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("X-APP-TOKEN", "hi")
				// checking Delete object options
				resp, _ := TestApp.Test(req)
				assert.Equalf(t, 200, resp.StatusCode, "Checking Delete Object")
			})

			t.Run("Checkinng Change Password", func(t *testing.T) {

				// creating patchdata
				patch_data_token := userpost{
					Email:    test.postData.Email,
					Password: "default@12345",
				}

				patch_data_string, _ := json.Marshal(patch_data_token)
				test_route := fmt.Sprintf("%v/%v?reset=false", test.route, responseMap["data"].(map[string]interface{})["id"])
				//  checking token decode options
				req := httptest.NewRequest("PUT", test_route, bytes.NewReader(patch_data_string))
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("X-APP-TOKEN", "hi")
				// checking Delete object options
				resp, _ := TestApp.Test(req)
				assert.Equalf(t, 200, resp.StatusCode, "Checking Delete Object")
			})

			t.Run("Checkinng Delete User", func(t *testing.T) {

				test_route := fmt.Sprintf("%v/%v", test.route, responseMap["data"].(map[string]interface{})["id"])
				//  checking token decode options
				req := httptest.NewRequest("DELETE", test_route, nil)
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("X-APP-TOKEN", "hi")
				// checking Delete object options
				resp, _ := TestApp.Test(req)
				assert.Equalf(t, 200, resp.StatusCode, "Checking Delete Object")
			})

			// This comment is kept for debuging purpose when making changes to the endpoint
			// fmt.Println(resp.StatusCode)
			// body, _ := io.ReadAll(resp.Body)
			// fmt.Println(string(body))

			// Verify, if the status code is as expected
			assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)

		})

	}

}
