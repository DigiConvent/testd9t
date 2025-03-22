package sys_router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *SysRouter) SetSmallLogo(c *gin.Context) {
	c.Request.ParseMultipartForm(10 << 20)
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

	fmt.Println(err)
	c.JSON(status.Code, gin.H{
		"status": status.Message,
	})
}
