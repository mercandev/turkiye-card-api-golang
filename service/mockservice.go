package service

import (
	"net/http"
	"turkey-card/domain"

	"github.com/gin-gonic/gin"
)

func GetMockData(c *gin.Context) {

	mockResponse := domain.MockData()

	c.JSON(http.StatusOK, gin.H{"result": mockResponse})
}
