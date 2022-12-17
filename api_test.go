import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	// Initialize the Echo instance
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Add the register handler as a route
	e.POST("/v1/register", registerHandler)

	// Create a test client
	client := echo.NewTestClient(e)

	// Define the request payload
	req := Request{User: "demo", Password: "pass"}
	reqBody, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	// Send a POST request to the /v1/register endpoint
	res, err := client.Post("/v1/register", "application/json", bytes.NewReader(reqBody))
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	// Check the status code
	assert.Equal(t, http.StatusOK, res.StatusCode)

	// Parse the response body as a JSON object
	var resp map[string]string
	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}

	// Check the response message
	assert.Equal(t, "Register success", resp["message"])
}