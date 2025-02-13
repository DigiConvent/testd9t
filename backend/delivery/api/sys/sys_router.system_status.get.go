package sys_router

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (r *SysRouter) SystemStatusGet(c *gin.Context) {
	id := uuid.MustParse("00000000-0000-0000-0000-000000000000")
	r.postService.SendEmail(&id, os.Getenv("EMAIL"), "Test", "Hello world")
	systemStatus, status := r.sysService.GetSystemStatus()

	if status.Code != 200 {
		c.JSON(status.Code, status)
		return
	} else {
		c.JSON(status.Code, systemStatus)
	}
}
