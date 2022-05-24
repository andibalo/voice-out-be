package handlers

import (
	"fmt"
	"net/http"
	"voice-out-be/internal/constants"
	"voice-out-be/internal/request"
	"voice-out-be/internal/response"
	"voice-out-be/internal/service"
	"voice-out-be/internal/voerrors"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

type Post struct {
	postService service.PostService
}

func NewPostHandler(postService service.PostService) *Post {
	return &Post{
		postService: postService,
	}
}

func (h *Post) AddRoutes(e *echo.Echo) {
	r := e.Group(constants.V1BasePath + constants.PostAPIPath)

	r.Use(middleware.JWT([]byte(viper.GetString("JWT_SECRET"))))
	r.POST("/", h.createPost)
	r.GET("/", h.getAllPosts)
	r.GET(constants.FetchPostsByUserIDAPIPath, h.getAllPostsByUserID)
}

func (h *Post) createPost(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	userId := fmt.Sprintf("%s", claims["userId"])

	createPostReq := &request.CreatePostRequest{}

	if err := c.Bind(createPostReq); err != nil {
		return err
	}

	err := createPostReq.Validate()

	if err != nil {
		validationErrorMessage := err.Error()
		return h.failedPostResponse(c, response.BadRequest, err, validationErrorMessage)
	}

	code, err := h.postService.CreatePost(createPostReq, userId)

	if err != nil {
		return h.failedPostResponse(c, code, err, "")
	}

	resp := response.NewResponse(code, nil)

	return c.JSON(http.StatusOK, resp)
}

func (h *Post) getAllPosts(c echo.Context) error {

	code, posts, err := h.postService.FetchAllPosts()

	if err != nil {
		return h.failedPostResponse(c, code, err, "")
	}

	resp := response.NewResponse(code, posts)

	return c.JSON(http.StatusOK, resp)
}

func (h *Post) getAllPostsByUserID(c echo.Context) error {

	userId := c.Param("id")

	code, posts, err := h.postService.FetchAllPostsByUserID(userId)

	if err != nil {
		return h.failedPostResponse(c, code, err, "")
	}

	resp := response.NewResponse(code, posts)

	return c.JSON(http.StatusOK, resp)
}

func (h *Post) failedPostResponse(c echo.Context, code response.Code, err error, errorMsg string) error {
	if code == "" {
		code = voerrors.MapErrorsToCode(err)
	}

	resp := response.Wrapper{
		ResponseCode: code,
		Status:       code.GetStatus(),
		Message:      code.GetMessage(),
	}

	if errorMsg != "" {
		resp.SetResponseMessage(errorMsg)
	}

	return c.JSON(voerrors.MapErrorsToStatusCode(err), resp)
}
