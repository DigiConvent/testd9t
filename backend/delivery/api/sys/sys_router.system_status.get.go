package sys_router

import (
	"github.com/gin-gonic/gin"
)

func (r *SysRouter) SystemStatusGet(c *gin.Context) {
	systemStatus, status := r.sysService.GetSystemStatus()

	if status.Code != 200 {
		c.JSON(status.Code, status)
		return
	} else {
		c.JSON(status.Code, systemStatus)
	}
}
