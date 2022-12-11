package handlers

import (
	"strconv"

	"github.com/javlonbekNormurodov/api_gateway/api/http"
	"github.com/javlonbekNormurodov/api_gateway/config"
	"github.com/javlonbekNormurodov/api_gateway/pkg/logger"
	"github.com/javlonbekNormurodov/api_gateway/services"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	cfg      config.Config
	log      logger.LoggerI
	services services.ServiceManagerI
}

func NewHandler(cfg config.Config, log logger.LoggerI, svcs services.ServiceManagerI) Handler {
	return Handler{
		cfg:      cfg,
		log:      log,
		services: svcs,
	}
}

func (h *Handler) handleResponse(c *gin.Context, status http.Status, data interface{}) {
	switch code := status.Code; {
	case code < 300:
		h.log.Info(
			"response",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			// logger.Any("data", data),
		)
	case code < 400:
		h.log.Warn(
			"response",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	default:
		h.log.Error(
			"response",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	}

	c.JSON(status.Code, http.Response{
		Status:      status.Status,
		Description: status.Description,
		Data:        data,
	})
}

func (h *Handler) getOffsetParam(c *gin.Context) (offset int, err error) {
	offsetStr := c.DefaultQuery("offset", h.cfg.DefaultOffset)
	return strconv.Atoi(offsetStr)
}

func (h *Handler) getLimitParam(c *gin.Context) (offset int, err error) {
	offsetStr := c.DefaultQuery("limit", h.cfg.DefaultLimit)
	return strconv.Atoi(offsetStr)
}
