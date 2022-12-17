package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Request struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type Response struct {
	Message string `json:"message"`
}

func registerHandler(c echo.Context) error {
	// Parse the request body as a JSON object
	var req Request
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
	}

	// Validate the request parameters
	if req.User == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Missing required parameters"})
	}

	// Perform the registration
	err = doRegister(req.User, req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal server error"})
	}

	// Return success response
	return c.JSON(http.StatusOK, map[string]string{"message": "Register success"})
}

func main() {
	e := echo.New()
	// Set up middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.POST("/v1/register", registerHandler)
	e.Logger.Fatal(e.Start(":8080"))
}