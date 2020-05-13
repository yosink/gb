package comm

import (
	blogger "blog/logging"
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
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
