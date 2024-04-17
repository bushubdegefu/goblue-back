package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert" // add Testify package
)

func TestAppRoles(t *testing.T) {
	// initalaizing the app
	ReturnTestApp()

	// Define a structure for specifying input and output data
	// of a single test case
	tests_role := []struct {
		name         string //name of string
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		// First test case
		{
			name:         "get roles check",
			description:  "get HTTP status 200",
			route:        "/api/v1/roles?page=1&size=10",
			expectedCode: 200,
		},
		// Second test case
		{
			name:         "get roles check",
			description:  "get HTTP status 404, when token does not exist",
			route:        "/api/v1/roles?page=1&size=10",
			expectedCode: 200,
		},
	}

	// Iterate through test single test cases
	for _, test := range tests_role {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", test.route, nil)
			req.Header.Set("X-APP-TOKEN", "hi")
			resp, _ := TestApp.Test(req)
			assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)

		})
		// Create a new http request with the route from the test case

		// This comment is kept for debuging purpose when making changes to the endpoint
		// fmt.Println(resp.StatusCode)
		// body, _ := io.ReadAll(resp.Body)
		// fmt.Println(string(body))

		// Verify, if the status code is as expected

	}

}
