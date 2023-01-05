package controller

import (
	"github.com/gin-gonic/gin"
	orderApi "github.com/h-varmazyar/snappfood/services/order/api/proto"
	"github.com/h-varmazyar/snappfood/services/order/internal/app/manager/service"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Controller struct {
	managerService *service.Service
	logger         *log.Logger
	configs        *Configs
}

func NewController(logger *log.Logger, configs *Configs, managerService *service.Service) *Controller {
	return &Controller{
		logger:         logger,
		configs:        configs,
		managerService: managerService,
	}
}

func (c *Controller) RegisterRoutes(router *gin.Engine) {
	orderRoutes := router.Group("/order")
	orderRoutes.POST("/", c.createOrder)
}

func (c *Controller) createOrder(ctx *gin.Context) {
	req := new(orderApi.ManagerCreateOrderReq)
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if voucher, err := c.managerService.CreateOrder(ctx, req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusCreated, voucher)
	}
}
