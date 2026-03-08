package gin

import (
	"bcc-workshop-2/config"

	ginfwk "github.com/gin-gonic/gin"
)

func New(cfg *config.Config) *ginfwk.Engine {
	if cfg.App.Env == "production" {
		ginfwk.SetMode(ginfwk.ReleaseMode)
	}

	router := ginfwk.Default()

	router.GET("/health", func(c *ginfwk.Context) {
		c.JSON(200, ginfwk.H{
			"status":  "ok",
			"message": "Server is running",
		})
	})

	return router
}
