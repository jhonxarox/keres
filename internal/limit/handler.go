package limit

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Repo *Repository
}

func NewHandler(repo *Repository) *Handler {
	return &Handler{Repo: repo}
}

// GetLimits returns limits for a specific customer
func (h *Handler) GetLimits(c *gin.Context) {
	customerID, err := strconv.Atoi(c.Param("customer_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	limits, err := h.Repo.GetLimitsByCustomerID(customerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, limits)
}

// CreateLimit creates a new limit for a customer
func (h *Handler) CreateLimit(c *gin.Context) {
	var newLimit Limit
	if err := c.ShouldBindJSON(&newLimit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Repo.CreateLimit(&newLimit); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newLimit)
}
