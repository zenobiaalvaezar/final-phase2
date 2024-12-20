package main

import (
	_ "car-rental/docs"
	"car-rental/internal/handlers"
	"car-rental/pkg/database"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	swag "github.com/swaggo/echo-swagger"
	"log"
)

// @title Car Rental API
// @version 1.0
// @description API documentation for Car Rental Service.
// @contact.name API Support
// @contact.email support@car-rental.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /api/v1

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	database.InitDB()

	e := echo.New()

	// Swagger documentation route
	e.GET("/swagger/*", swag.WrapHandler)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", handlers.GetCars)

	// @Summary Register a new user
	// @Description Create a new user account
	// @Tags Authentication
	// @Accept json
	// @Produce json
	// @Param request body handlers.RegisterRequest true "User registration data"
	// @Success 201 {object} map[string]string "{\"message\": \"Registration successful\"}"
	// @Failure 400 {object} map[string]string "{\"message\": \"Error\"}"
	// @Router /api/v1/register [post]
	e.POST("/api/v1/register", handlers.Register)

	// @Summary Login a user
	// @Description Authenticate a user and return a JWT token
	// @Tags Authentication
	// @Accept json
	// @Produce json
	// @Param request body handlers.LoginRequest true "User login data"
	// @Success 200 {object} map[string]interface{} "OK"
	// @Failure 401 {object} map[string]string "{\"message\": \"Invalid credentials\"}"
	// @Router /api/v1/login [post]
	e.POST("/api/v1/login", handlers.Login)

	api := e.Group("/api/v1")

	// @Summary Get user profile
	// @Description Retrieve details of the logged-in user
	// @Tags User
	// @Produce json
	// @Success 200 {object} map[string]interface{} "OK"
	// @Failure 401 {object} map[string]string "Unauthorized"
	// @Router /api/v1/profile [get]
	api.GET("/profile", handlers.GetProfile)

	// @Summary Top up user deposit
	// @Description Add funds to the user account
	// @Tags User
	// @Accept json
	// @Produce json
	// @Param request body handlers.TopUpRequest true "Top-up request data"
	// @Success 200 {object} map[string]interface{} "OK"
	// @Failure 400 {object} map[string]string "Bad request"
	// @Router /api/v1/topup [post]
	api.POST("/topup", handlers.TopUp)

	// @Summary Get car details
	// @Description Retrieve details of a specific car
	// @Tags Cars
	// @Produce json
	// @Param id path int true "Car ID"
	// @Success 200 {object} map[string]interface{} "OK"
	// @Failure 404 {object} map[string]string "Car not found"
	// @Router /api/v1/cars/{id} [get]
	api.GET("/cars/:id", handlers.GetCarDetail)

	// @Summary Create a car rental
	// @Description Create a new rental for a car
	// @Tags Rentals
	// @Accept json
	// @Produce json
	// @Param request body handlers.CreateRentalRequest true "Rental request data"
	// @Success 201 {object} map[string]interface{} "Rental created"
	// @Failure 400 {object} map[string]string "Bad request"
	// @Router /api/v1/rentals [post]
	api.POST("/rentals", handlers.CreateRental)

	// @Summary Get user rentals
	// @Description Retrieve all rentals made by the logged-in user
	// @Tags Rentals
	// @Produce json
	// @Success 200 {array} map[string]interface{} "OK"
	// @Failure 500 {object} map[string]string "Internal server error"
	// @Router /api/v1/rentals [get]
	api.GET("/rentals", handlers.GetUserRentals)

	// @Summary Return a rented car
	// @Description Complete a car rental by returning it
	// @Tags Rentals
	// @Produce json
	// @Param id path int true "Rental ID"
	// @Success 200 {object} map[string]string "Rental completed"
	// @Failure 400 {object} map[string]string "Bad request"
	// @Router /api/v1/rentals/{id}/return [post]
	api.POST("/rentals/:id/return", handlers.ReturnCar)

	// @Summary Get payment history
	// @Description Retrieve payment history for the logged-in user
	// @Tags Payments
	// @Produce json
	// @Success 200 {array} map[string]interface{} "Payment history"
	// @Failure 500 {object} map[string]string "Internal server error"
	// @Router /api/v1/payments [get]
	api.GET("/payments", handlers.GetPaymentHistory)

	// @Summary Get payment details
	// @Description Retrieve details of a specific payment
	// @Tags Payments
	// @Produce json
	// @Param id path int true "Payment ID"
	// @Success 200 {object} map[string]interface{} "Payment details"
	// @Failure 404 {object} map[string]string "Payment not found"
	// @Router /api/v1/payments/{id} [get]
	api.GET("/payments/:id", handlers.GetPaymentDetail)

	// @Summary Handle payment webhook
	// @Description Process payment webhook notifications
	// @Tags Payments
	// @Accept json
	// @Produce json
	// @Param request body handlers.WebhookPayload true "Webhook payload"
	// @Success 200 {object} map[string]string "Webhook processed"
	// @Failure 401 {object} map[string]string "Unauthorized"
	// @Router /api/v1/payments/webhook [post]
	e.POST("/api/v1/payments/webhook", handlers.WebhookHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
