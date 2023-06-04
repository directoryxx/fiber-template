package test

import (
	"bytes"
	httppkg "clean-arch-template/internal/delivery/http"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestRoleAllRoute(t *testing.T) {
	initInfrastructure()
	migrationUp()
	// Define a structure for specifying input and output
	// data of a single test case. This structure is then used
	// to create a so called test map, which contains all test
	// cases, that should be run for testing this function
	tests := []struct {
		description string

		// Test input
		route string

		// Expected output
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			description:   "roles route",
			route:         "/api/v1/roles",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "1",
		},
		{
			description:   "roles route",
			route:         "/api/v1/roles?page=2",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "1",
		},
		{
			description:   "non existing route",
			route:         "/i-dont-exist",
			expectedError: false,
			expectedCode:  404,
			expectedBody:  "Cannot GET /i-dont-exist",
		},
	}

	// Setup the app as it is done in the main function
	app := httppkg.Run()

	// Iterate through test single test cases
	for _, test := range tests {
		// Create a new http request with the route
		// from the test case
		req, _ := http.NewRequest(
			"GET",
			test.route,
			nil,
		)

		// Perform the request plain with the app.
		// The -1 disables request latency.
		res, err := app.Test(req, -1)

		// verify that no error occured, that is not expected
		assert.Equalf(t, test.expectedError, err != nil, test.description)

		// As expected errors lead to broken responses, the next
		// test case needs to be processed
		if test.expectedError {
			continue
		}

		// Verify if the status code is as expected
		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

		// Read the response body
		//body, err := io.ReadAll(res.Body)
		//
		//var resBody *response.PaginationResponse
		//
		//json.Unmarshal([]byte(string(body)), &resBody)
		//
		//data := resBody.Data

		//fmt.Println(string(body))
		//// Reading the response body should work everytime, such that
		//// the err variable should be nil
		//assert.Nilf(t, err, test.description)
		//
		//// Verify, that the reponse body equals the expected body
		//assert.Equalf(t, test.expectedBody, string(body), test.description)
	}
}

func TestRoleGetRoute(t *testing.T) {
	initInfrastructure()
	migrationUp()
	// Define a structure for specifying input and output
	// data of a single test case. This structure is then used
	// to create a so called test map, which contains all test
	// cases, that should be run for testing this function
	tests := []struct {
		description string

		// Test input
		route string

		// Expected output
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			description:   "roles get 1",
			route:         "/api/v1/roles/1",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "1",
		},
		{
			description:   "roles get 2",
			route:         "/api/v1/roles/2",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "1",
		},
		{
			description:   "roles get 3",
			route:         "/api/v1/roles/3",
			expectedError: false,
			expectedCode:  404,
			expectedBody:  "1",
		},
		{
			description:   "non existing route",
			route:         "/i-dont-exist",
			expectedError: false,
			expectedCode:  404,
			expectedBody:  "Cannot GET /i-dont-exist",
		},
	}

	// Setup the app as it is done in the main function
	app := httppkg.Run()

	// Iterate through test single test cases
	for _, test := range tests {
		// Create a new http request with the route
		// from the test case
		req, _ := http.NewRequest(
			"GET",
			test.route,
			nil,
		)

		// Perform the request plain with the app.
		// The -1 disables request latency.
		res, err := app.Test(req, -1)

		// verify that no error occured, that is not expected
		assert.Equalf(t, test.expectedError, err != nil, test.description)

		// As expected errors lead to broken responses, the next
		// test case needs to be processed
		if test.expectedError {
			continue
		}

		// Verify if the status code is as expected
		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

		// Read the response body
		//body, err := io.ReadAll(res.Body)
		//
		//var resBody *response.PaginationResponse
		//
		//json.Unmarshal([]byte(string(body)), &resBody)
		//
		//data := resBody.Data

		//fmt.Println(string(body))
		//// Reading the response body should work everytime, such that
		//// the err variable should be nil
		//assert.Nilf(t, err, test.description)
		//
		//// Verify, that the reponse body equals the expected body
		//assert.Equalf(t, test.expectedBody, string(body), test.description)
	}
}

func TestRolePostRoute(t *testing.T) {
	initInfrastructure()
	migrationUp()
	// Define a structure for specifying input and output
	// data of a single test case. This structure is then used
	// to create a so called test map, which contains all test
	// cases, that should be run for testing this function
	tests := []struct {
		description string

		// Test input
		route string

		// Expected output
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			description:   "create roles",
			route:         "/api/v1/roles",
			expectedError: false,
			expectedCode:  201,
			expectedBody:  "1",
		},
	}

	// Setup the app as it is done in the main function
	app := httppkg.Run()

	// Iterate through test single test cases
	for _, test := range tests {
		var jsonStr = []byte(`{"name":"coba"}`)
		// Create a new http request with the route
		// from the test case
		req, _ := http.NewRequest(
			"POST",
			test.route,
			bytes.NewBuffer(jsonStr),
		)
		req.Header.Set("Content-Type", "application/json")

		// Perform the request plain with the app.
		// The -1 disables request latency.
		res, err := app.Test(req, -1)

		// verify that no error occured, that is not expected
		assert.Equalf(t, test.expectedError, err != nil, test.description)

		// As expected errors lead to broken responses, the next
		// test case needs to be processed
		if test.expectedError {
			continue
		}

		// Verify if the status code is as expected
		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

		// Read the response body
		//body, err := io.ReadAll(res.Body)

		//// Reading the response body should work everytime, such that
		//// the err variable should be nil
		//assert.Nilf(t, err, test.description)
		//
		//// Verify, that the reponse body equals the expected body
		//assert.Equalf(t, test.expectedBody, string(body), test.description)
	}
}

func TestRoleDeleteRoute(t *testing.T) {
	initInfrastructure()
	migrationUp()
	// Define a structure for specifying input and output
	// data of a single test case. This structure is then used
	// to create a so called test map, which contains all test
	// cases, that should be run for testing this function
	tests := []struct {
		description string

		// Test input
		route string

		// Expected output
		expectedError bool
		expectedCode  int
		expectedBody  string

		method string
	}{
		{
			description:   "create roles",
			route:         "/api/v1/roles",
			expectedError: false,
			expectedCode:  201,
			expectedBody:  "1",
			method:        "POST",
		},
		{
			description:   "delete roles",
			route:         "/api/v1/roles/3",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "1",
			method:        "DELETE",
		},
		{
			description:   "delete roles",
			route:         "/api/v1/roles/4",
			expectedError: false,
			expectedCode:  404,
			expectedBody:  "1",
			method:        "DELETE",
		},
	}

	// Setup the app as it is done in the main function
	app := httppkg.Run()

	// Iterate through test single test cases
	for _, test := range tests {
		var jsonStr = []byte(`{"name":"coba"}`)
		// Create a new http request with the route
		// from the test case
		req, _ := http.NewRequest(
			test.method,
			test.route,
			bytes.NewBuffer(jsonStr),
		)
		req.Header.Set("Content-Type", "application/json")

		// Perform the request plain with the app.
		// The -1 disables request latency.
		res, err := app.Test(req, -1)

		// verify that no error occured, that is not expected
		assert.Equalf(t, test.expectedError, err != nil, test.description)

		// As expected errors lead to broken responses, the next
		// test case needs to be processed
		if test.expectedError {
			continue
		}

		// Verify if the status code is as expected
		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

		// Read the response body
		//body, err := io.ReadAll(res.Body)

		//// Reading the response body should work everytime, such that
		//// the err variable should be nil
		//assert.Nilf(t, err, test.description)
		//
		//// Verify, that the reponse body equals the expected body
		//assert.Equalf(t, test.expectedBody, string(body), test.description)
	}
}
