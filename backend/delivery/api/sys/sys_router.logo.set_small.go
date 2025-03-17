package sys_router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *SysRouter) SetSmallLogo(c *gin.Context) {

	bytes, err := getLogoBytes(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file: " + err.Error()})
		return
	}

	status := r.sysService.SetSmallLogo(bytes)

	if status.Err() {
		c.JSON(status.Code, gin.H{"error": status.Message})
		return
	}

	c.JSON(status.Code, gin.H{
		"status": status.Message,
	})
}
