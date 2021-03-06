package views

import (
	"net/http"
	"focus-gin/app/models"

	"github.com/gin-gonic/gin"
)

// GetTickets is ...
func GetTickets(c *gin.Context) {
	// Connection handling of the database
	db := models.InitDb()

	var tickets []models.Ticket
	// SELECT * FROM tickets
	// db.Find(&tickets, []int{4, 6})
	db.Find(&tickets, 1)

	// Display JSON result
	out := gin.H{"tickets": tickets}
	c.JSON(http.StatusOK, out)
}
