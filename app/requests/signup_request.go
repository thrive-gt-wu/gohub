package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SignupPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
}

func ValidateSignupPhoneExist(data interface{}, c *gin.Context) map[string][]string {
	// 自定义验证规则
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}

	// 自定义演示出错时候的提示
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称phone",
			"digits:手机号长度必须为11为的数字",
		},
	}

	// 配置初始化
	opts := govalidator.Options{
		Data: data,
		Rules: rules,
		TagIdentifier: "valid",	// 模型中的 struct 标签标识符
		Messages: messages,

	}

	// 验证
	return govalidator.New(opts).ValidateStruct()
}