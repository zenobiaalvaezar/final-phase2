package handlers

import (
	"car-rental/internal/models"
	"car-rental/pkg/database"
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetCars handler

// GetCars godoc
// @Summary Get all cars
// @Description Retrieve a list of all available cars
// @Tags Cars
// @Produce json
// @Success 200 {array} map[string]interface{} "List of cars"
// @Router /api/v1/cars [get]
func GetCars(c echo.Context) error {
	var cars []models.Car

	// Get query parameters
	category := c.QueryParam("category")
	available := c.QueryParam("available")

	// Build query
	query := database.DB
	if category != "" {
		query = query.Where("category = ?", category)
	}
	if available == "true" {
		query = query.Where("stock_availability > ?", 0)
	}

	// Execute query
	if err := query.Find(&cars).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch cars")
	}

	formattedCars := []map[string]interface{}{}
	for _, car := range cars {
		formattedCars = append(formattedCars, map[string]interface{}{
			"id":                 car.ID,
			"name":               car.Name,
			"stock_availability": car.StockAvailability,
			"rental_costs":       car.RentalCosts,
			"category":           car.Category,
			"created_at":         car.CreatedAt.Format("2006-01-02 15:04:05"), // Format baru
			"updated_at":         car.UpdatedAt.Format("2006-01-02 15:04:05"), // Format baru
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": formattedCars,
	})
}

// GetCarDetail godoc
// @Summary Get car details
// @Description Retrieve details of a specific car
// @Tags Cars
// @Produce json
// @Param id path int true "Car ID"
// @Success 200 {object} map[string]interface{} "Car details"
// @Failure 404 {object} map[string]string "message: Car not found"
// @Router /api/v1/cars/{id} [get]
func GetCarDetail(c echo.Context) error {
	carID := c.Param("id")

	var car models.Car
	if err := database.DB.First(&car, carID).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Car not found")
	}

	formattedCar := map[string]interface{}{
		"id":                 car.ID,
		"name":               car.Name,
		"stock_availability": car.StockAvailability,
		"rental_costs":       car.RentalCosts,
		"category":           car.Category,
		"created_at":         car.CreatedAt.Format("2006-01-02 15:04:05"), // Format baru
		"updated_at":         car.UpdatedAt.Format("2006-01-02 15:04:05"), // Format baru
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": formattedCar,
	})
}
