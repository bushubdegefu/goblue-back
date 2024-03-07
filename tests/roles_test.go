package responses

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert" // add Testify package
	"semay.com/manager"
)

func TestAppRoles(t *testing.T) {
	// initalaizing the app
	app, file := manager.MakeApp("test")
	defer file.Close()
	// Define a structure for specifying input and output data
	// of a single test case
	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		// First test case
		{
			description:  "get HTTP status 200",
			route:        "/api/v1/roles?page=1&size=10",
			expectedCode: 200,
		},
		// Second test case
		{
			description:  "get HTTP status 404, when token does not exist",
			route:        "/api/v1/roles?page=1&size=10",
			expectedCode: 200,
		},
	}

	// Iterate through test single test cases
	for _, test := range tests {
		// Create a new http request with the route from the test case
		req := httptest.NewRequest("GET", test.route, nil)
		req.Header.Set("X-APP-KEY", "hi")
		resp, _ := app.Test(req)

		// This comment is kept for debuging purpose when making changes to the endpoint
		// fmt.Println(resp.StatusCode)
		// body, _ := io.ReadAll(resp.Body)
		// fmt.Println(string(body))

		// Verify, if the status code is as expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)

	}
}
