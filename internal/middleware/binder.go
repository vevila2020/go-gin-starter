package middleware

import (
	"go_basic_framework_template/internal/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func BindJSON(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		response.BadRequest(c, "invalid request body: "+err.Error())
		return err
	}
	if err := validate.Struct(obj); err != nil {
		response.BadRequest(c, "validation failed: "+err.Error())
		return err
	}
	return nil
}

func BindQuery(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindQuery(obj); err != nil {
		response.BadRequest(c, "invalid query parameters: "+err.Error())
		return err
	}
	if err := validate.Struct(obj); err != nil {
		response.BadRequest(c, "validation failed: "+err.Error())
		return err
	}
	return nil
}

func BindURI(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindUri(obj); err != nil {
		response.BadRequest(c, "invalid uri parameters: "+err.Error())
		return err
	}
	if err := validate.Struct(obj); err != nil {
		response.BadRequest(c, "validation failed: "+err.Error())
		return err
	}
	return nil
}
