package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/opensource-cloud/sycorax/internal/config"
)

func RegisterV1RoutesGroups(rg *gin.RouterGroup, c *config.Config) {
	RegisterQueueV1Routes(rg, c)
}
