package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// JSONResponse is the default API response type
type JSONResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type scanURL struct {
	Number uint `uri:"number" binding:"required,min=1,max=999679368229"`
}

// ValidateScanURL validates scan URLs
func ValidateScanURL(c *gin.Context) {
	var v scanURL

	if err := c.ShouldBindUri(&v); err != nil {
		errorHandling(c, err.Error())
	}
}

func errorHandling(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, JSONResponse{Success: false, Error: msg})
	c.Abort()
}
