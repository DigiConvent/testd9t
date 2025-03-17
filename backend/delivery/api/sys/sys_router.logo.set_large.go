package sys_router

import (
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (r *SysRouter) SetLargeLogo(c *gin.Context) {
	bytes, err := getLogoBytes(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file: " + err.Error()})
		return
	}
	status := r.sysService.SetLargeLogo(bytes)

	if status.Err() {
		c.JSON(status.Code, gin.H{"error": status.Message})
		return
	}

	c.JSON(status.Code, gin.H{
		"status": status.Message,
	})
}

func getLogoBytes(c *gin.Context) ([]byte, error) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		return nil, err
	}

	if !strings.HasSuffix(header.Filename, ".jpg") {
		return nil, err
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
