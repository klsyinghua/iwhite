package handlers

import (
	"github.com/labstack/echo/v4"
	"iwhite/models"
	"net/http"
)

func (h *ServerHandler) SearchHandler(c echo.Context) error {
	// Implement search logic and return suggestions
	suggestions, err := (&models.Server{}).QueryAllServers(h.db)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to query servers")
	}

	return c.JSON(http.StatusOK, suggestions)
}

func (h *ServerHandler) QueryHandler(c echo.Context) error {
	identifier := c.Param("identifier")
	responseData := &models.Server{}
	err := responseData.QueryServerByHostnameOrIP(h.db, identifier)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to query server")
	}
	return c.JSON(http.StatusOK, responseData)
}

func (h *ServerHandler) HelloHandler(c echo.Context) error {
	// Implement annotations logic and return data
	// ...
	// annotationsData := ...
	return c.JSON(http.StatusOK, "")
}

func (h *ServerHandler) AnnotationsHandler(c echo.Context) error {
	// Get the start and end timestamp from the query parameters
	start := c.QueryParam("from")
	end := c.QueryParam("to")

	// Implement your logic to retrieve annotations from the database
	annotations, err := h.getAnnotationsFromDatabase(start, end)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to retrieve annotations")
	}

	// Construct the JSON response
	return c.JSON(http.StatusOK, annotations)
}

func (h *ServerHandler) getAnnotationsFromDatabase(start, end string) ([]map[string]interface{}, error) {
	// Implement your logic to query the database for annotations
	// Use the provided start and end timestamps to filter the annotations

	// Sample annotations data structure
	annotations := []map[string]interface{}{
		{
			"time":     1630944000000, // Unix timestamp in milliseconds
			"title":    "Event 1",
			"text":     "Something important happened",
			"tags":     []string{"event", "important"},
			"color":    "#FF5733",
			"priority": 1,
		},
		// Add more annotations as needed
	}

	return annotations, nil
}
