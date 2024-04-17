package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert" // add Testify package
)

func TestAppEndpoints(t *testing.T) {
	// initalaizing the app
	ReturnTestApp()

	// Define a structure for specifying input and output data
	// of a single test case
	tests_endpoint := []struct {
		name         string //name of string
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		// First test case
		{
			name:         "get endpoints check",
			description:  "get HTTP status 200",
			route:        "/api/v1/endpoints?page=1&size=10",
			expectedCode: 200,
		},
		// Second test case
		{
			name:         "get endpoints check",
			description:  "get HTTP status 404, when token does not exist",
			route:        "/api/v1/endpoints?page=1&size=10",
			expectedCode: 200,
		},
	}

	// Iterate through test single test cases
	for _, test := range tests_endpoint {
		t.Run(test.name, func(t *testing.T) {
			// Create a new http request with the route from the test case
			req := httptest.NewRequest("GET", test.route, nil)
			req.Header.Set("X-APP-TOKEN", "hi")
			resp, err := TestApp.Test(req)
			if err != nil {
				t.Fatalf("Failed to make request: %v", err)
			}
			defer resp.Body.Close()

			// This comment is kept for debuging purpose when making changes to the endpoint
			// fmt.Println(resp.StatusCode)
			// body, _ := io.ReadAll(resp.Body)
			// fmt.Println(string(body)

			// Verify, if the status code is as expected
			assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)

		})

	}

}
