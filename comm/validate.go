package comm

import (
	blogger "blog/logging"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

var valid *validator.Validate

func init() {
	valid = validator.New()
}

func BindAndValid(c *gin.Context, form interface{}) (bool, error) {
	err := c.Bind(form)
	if err != nil {
		blogger.Error("binding error:", err)
		return false, errors.New("参数错误")
	}
	checked, err := govalidator.ValidateStruct(form)
	if err != nil {
		blogger.Error("validate error:", err)
		errs := err.(govalidator.Errors).Errors()
		return false, errs[0]
	}
	if !checked {
		blogger.Error("校验失败")
		return false, errors.New("校验失败")
	}
	return true, nil
}

func ValidateBind(c *gin.Context, form interface{}) (bool, error) {
	err := c.Bind(form)
	if err != nil {
		blogger.Error("binding error:", err)
		errs := err.(validator.ValidationErrors)
		return false, fmt.Errorf("%s should be %s", errs[0].Field(), errs[0].Tag())
	}
	return true, nil
}
