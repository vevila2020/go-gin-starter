package response

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type PageData struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
}

const (
	CodeSuccess      = 0
	CodeBadRequest   = 400
	CodeUnauthorized = 401
	CodeForbidden    = 403
	CodeNotFound     = 404
	CodeServerError  = 500
)

var codeMessages = map[int]string{
	CodeSuccess:      "success",
	CodeBadRequest:   "bad request",
	CodeUnauthorized: "unauthorized",
	CodeForbidden:    "forbidden",
	CodeNotFound:     "not found",
	CodeServerError:  "internal server error",
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    CodeSuccess,
		Message: codeMessages[CodeSuccess],
		Data:    data,
	})
}

func SuccessWithMessage(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    CodeSuccess,
		Message: message,
		Data:    data,
	})
}

func SuccessPage(c *gin.Context, list interface{}, total int64, page, pageSize int) {
	c.JSON(http.StatusOK, Response{
		Code:    CodeSuccess,
		Message: codeMessages[CodeSuccess],
		Data: PageData{
			List:     list,
			Total:    total,
			Page:     page,
			PageSize: pageSize,
		},
	})
}

func Error(c *gin.Context, code int, message string) {
	if message == "" {
		message = codeMessages[code]
	}
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
	})
}

func BadRequest(c *gin.Context, message string) {
	Error(c, CodeBadRequest, message)
}

func Unauthorized(c *gin.Context, message string) {
	Error(c, CodeUnauthorized, message)
}

func Forbidden(c *gin.Context, message string) {
	Error(c, CodeForbidden, message)
}

func NotFound(c *gin.Context, message string) {
	Error(c, CodeNotFound, message)
}

func ServerError(c *gin.Context, message string) {
	Error(c, CodeServerError, message)
}

func File(c *gin.Context, filepath, filename string) {
	c.FileAttachment(filepath, filename)
}

func FileFromBytes(c *gin.Context, data []byte, filename, contentType string) {
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Type", contentType)
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Expires", "0")
	c.Header("Cache-Control", "must-revalidate")
	c.Header("Pragma", "public")
	c.Data(http.StatusOK, contentType, data)
}

func Stream(c *gin.Context, reader io.Reader, filename, contentType string) {
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Type", contentType)
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Expires", "0")
	c.Header("Cache-Control", "must-revalidate")
	c.Header("Pragma", "public")
	c.DataFromReader(http.StatusOK, -1, contentType, reader, nil)
}
