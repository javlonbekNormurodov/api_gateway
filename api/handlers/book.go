package handlers

import (
	"context"

	"github.com/javlonbekNormurodov/api_gateway/api/http"
	"github.com/javlonbekNormurodov/api_gateway/genproto/book_service"

	// "github.com/javlonbekNormurodov/api_gateway/config"
	// "github.com/javlonbekNormurodov/api_gateway/pkg/logger"

	"github.com/gin-gonic/gin"
)

// CreateApp godoc
// @Security ApiKeyAuth
// @ID create_app
// @Router /v1/app [POST]
// @Summary Create app
// @Description Create app
// @Tags App
// @Accept json
// @Produce json
// @Param app body book_service.CreateBookRequest true "CreateBook"
// @Success 201 {object} http.Response{data=string} "App data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateBook(c *gin.Context) {
	var app book_service.CreateBookRequest

	err := c.ShouldBindJSON(&app)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.BookService().Create(
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
// @Router /v1/app [POST]
// @Summary Create app
// @Description Create app
// @Tags App
// @Accept json
// @Produce json
// @Param app body book_service.GetBooksRequest true "CreateAppRequestBody"
// @Success 201 {object} http.Response{data=string} "App data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetAllBooks(c *gin.Context) {
	var app book_service.GetBooksRequest

	err := c.ShouldBindJSON(&app)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.BookService().GetAll(
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
// @Router /v1/app [POST]
// @Summary Create app
// @Description Create app
// @Tags App
// @Accept json
// @Produce json
// @Param app body book_service.GetBookByIDRequest true "CreateAppRequestBody"
// @Success 201 {object} http.Response{data=string} "App data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetBookByID(c *gin.Context) {
	var app book_service.GetBookByIDRequest

	err := c.ShouldBindJSON(&app)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.BookService().GetByID(
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
// @Router /v1/app [POST]
// @Summary Create app
// @Description Create app
// @Tags App
// @Accept json
// @Produce json
// @Param app body book_service.UpdateBookRequest true "CreateAppRequestBody"
// @Success 201 {object} http.Response{data=string} "App data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateBook(c *gin.Context) {
	var app book_service.UpdateBookRequest

	err := c.ShouldBindJSON(&app)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.BookService().Update(
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
// @Router /v1/app [POST]
// @Summary Create app
// @Description Create app
// @Tags App
// @Accept json
// @Produce json
// @Param app body book_service.DeleteBookRequest true "CreateAppRequestBody"
// @Success 201 {object} http.Response{data=string} "App data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteBook(c *gin.Context) {
	var app book_service.DeleteBookRequest

	err := c.ShouldBindJSON(&app)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.BookService().Delete(
		context.Background(),
		&app,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}
