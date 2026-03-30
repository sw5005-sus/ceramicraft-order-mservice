package router

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	auditclient "github.com/sw5005-sus/ceramicraft-audit-client"
	"github.com/sw5005-sus/ceramicraft-order-mservice/server/config"
	_ "github.com/sw5005-sus/ceramicraft-order-mservice/server/docs"
	"github.com/sw5005-sus/ceramicraft-order-mservice/server/http/api"
	"github.com/sw5005-sus/ceramicraft-order-mservice/server/metrics"
	"github.com/sw5005-sus/ceramicraft-user-mservice/common/middleware"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

const (
	serviceURIPrefix = "/order-ms/v1"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	auditMiddleware := auditclient.AuditMiddleware(
		"order-ms",
		config.Config.AuditClient.Host,
		config.Config.AuditClient.Port)
	basicGroup := r.Group(serviceURIPrefix)
	{
		basicGroup.Use(metrics.MetricsMiddleware())
		basicGroup.GET("/metrics", gin.WrapH(promhttp.Handler()))

		basicGroup.GET("/swagger/*any", gs.WrapHandler(
			swaggerFiles.Handler,
			gs.URL("/order-ms/v1/swagger/doc.json"),
		))
		basicGroup.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		merchantGroup := basicGroup.Group("/merchant")
		{
			merchantGroup.Use(middleware.AuthMiddleware())
			merchantGroup.POST("/orders/list", api.ListOrders)
			merchantGroup.GET("/orders/:order_no", api.GetOrderDetail)
			merchantGroup.GET("/orders/:order_no/receive-info", middleware.RequireRoles("merchant_admin"), auditMiddleware, api.GetOrderReceiverInfo)
			merchantGroup.PATCH("/orders/:order_no/ship", middleware.RequireRoles("merchant_admin"), auditMiddleware, api.ShipOrder) // ship order
			merchantGroup.GET("/order-stats", api.GetOrderStats)                                                                     // get order stats
		}

		customerGroup := basicGroup.Group("/customer")
		{
			customerGroup.Use(middleware.AuthMiddleware())
			customerGroup.POST("/orders", api.CreateOrder) // create order
			customerGroup.POST("/orders/list", api.CustomerListOrders)
			customerGroup.GET("/orders/:order_no", api.CustomerGetOrderDetail) // get order detail
			customerGroup.PATCH("/orders/:order_no/confirm", api.ConfirmOrder) // confirm order
		}
	}
	return r
}
