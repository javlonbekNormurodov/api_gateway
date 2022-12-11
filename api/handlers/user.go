package handlers

import (
	"context"

	"github.com/javlonbekNormurodov/api_gateway/api/http"
	"github.com/javlonbekNormurodov/api_gateway/genproto/user_service"

	// "github.com/javlonbekNormurodov/api_gateway/config"
	// "github.com/javlonbekNormurodov/api_gateway/pkg/logger"

	"github.com/gin-gonic/gin"
)

// CreateUser godoc
// @Security ApiKeyAuth
// @ID create_user
// @Router /v1/user [POST]
// @Summary Create user
// @Description Create user
// @Tags user
// @Accept json
// @Produce json
// @Param app body user_service.CreateUserRequest true "CreateAppRequestBody"
// @Success 201 {object} http.Response{data=string} "user data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateUser(c *gin.Context) {
	var app user_service.CreateUserRequest

	err := c.ShouldBindJSON(&app)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.UserService().CreateUser(
		context.Background(),
		&app,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// CreateUser godoc
// @ID create_user
// @Router /v1/user [GET]
// @Summary Create user
// @Description Create user
// @Tags user
// @Accept json
// @Produce json
// @Param app body user_service.GetUsersRequest true "CreateUserRequestBody"
// @Success 201 {object} http.Response{data=string} "user data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetUsers(c *gin.Context) {
	var app user_service.GetUsersRequest

	err := c.ShouldBindJSON(&app)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.UserService().GetAllUsers(
		context.Background(),
		&app,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// CreateApp godoc
// @Security ApiKeyAuth
// @ID create_app
// @Router /v1/app [GET]
// @Summary Create app
// @Description Create app
// @Tags App
// @Accept json
// @Produce json
// @Param app body user_service.GetUserByIDRequest true "CreateAppRequestBody"
// @Success 201 {object} http.Response{data=string} "App data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetUserById(c *gin.Context) {
	var app user_service.GetUserByIDRequest

	err := c.ShouldBindJSON(&app)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.UserService().GetUserByID(
		context.Background(),
		&app,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// CreateApp godoc
// @Security ApiKeyAuth
// @ID create_app
// @Router /v1/app [PUT]
// @Summary Create app
// @Description Create app
// @Tags App
// @Accept json
// @Produce json
// @Param app body user_service.UpdateUserRequest true "CreateAppRequestBody"
// @Success 201 {object} http.Response{data=string} "App data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateUser(c *gin.Context) {
	var app user_service.UpdateUserRequest

	err := c.ShouldBindJSON(&app)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.UserService().UpdateUser(
		context.Background(),
		&app,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// CreateApp godoc
// @Security ApiKeyAuth
// @ID create_app
// @Router /v1/app [DELETE]
// @Summary Create app
// @Description Create app
// @Tags App
// @Accept json
// @Produce json
// @Param app body user_service.DeleteUserRequest true "CreateAppRequestBody"
// @Success 201 {object} http.Response{data=string} "App data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteUser(c *gin.Context) {
	var app user_service.DeleteUserRequest

	err := c.ShouldBindJSON(&app)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.UserService().DeleteUser(
		context.Background(),
		&app,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}
