package handlers

import (
	"net/http"
	"voice-out-be/internal/constants"
	"voice-out-be/internal/request"
	"voice-out-be/internal/response"
	"voice-out-be/internal/service"
	"voice-out-be/internal/voerrors"

	"github.com/labstack/echo/v4"
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

	r.POST("/", h.createPost)
}

func (h *Post) createPost(c echo.Context) error {

	createPostReq := &request.CreatePostRequest{}

	if err := c.Bind(createPostReq); err != nil {
		return err
	}

	err := createPostReq.Validate()

	if err != nil {
		validationErrorMessage := err.Error()
		return h.failedPostResponse(c, response.BadRequest, err, validationErrorMessage)
	}

	code, err := h.postService.CreatePost(createPostReq)

	if err != nil {
		return h.failedPostResponse(c, code, err, "")
	}

	return c.JSON(http.StatusOK, nil)
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
