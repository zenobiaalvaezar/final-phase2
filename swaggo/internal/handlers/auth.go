package handlers

import (
	"car-rental/internal/models"
	"car-rental/internal/services"
	"car-rental/pkg/database"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// Register godoc
// @Summary Register a new user
// @Description Create a new user account
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body handlers.RegisterRequest true "User registration data"
// @Success 201 {object} map[string]string "message: Registration successful"
// @Failure 400 {object} map[string]string "message: Error"
// @Router /api/v1/register [post]
func Register(c echo.Context) error {
	var req RegisterRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to hash password")
	}

	// Create user
	user := models.User{
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Email already exists")
	}

	// Send welcome email
	emailService := services.NewEmailService()
	go emailService.SendEmail(
		user.Email,
		"Welcome to Car Rental System",
		"Thank you for registering with our car rental service!",
	)

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "Registration successful",
	})
}

// Login godoc
// @Summary Login a user
// @Description Authenticate a user and return a JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body handlers.LoginRequest true "User login data"
// @Success 200 {object} map[string]interface{} "JWT token and user info"
// @Failure 401 {object} map[string]string "message: Invalid credentials"
// @Router /api/v1/login [post]
func Login(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Find user
	var user models.User
	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
	}

	// Generate JWT token
	token, err := services.GenerateJWT(user.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate token")
	}

	// Send login notification email
	emailService := services.NewEmailService()
	go emailService.SendEmail(
		user.Email,
		"New Login Detected",
		"A new login was detected on your account.",
	)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
		"user": map[string]interface{}{
			"id":             user.ID,
			"email":          user.Email,
			"deposit_amount": user.DepositAmount,
		},
	})
}
