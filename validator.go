package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
	"reflect"
)

var trans ut.Translator

func init() {
	// 创建翻译器
	uni := ut.New(zh.New())
	trans, _ = uni.GetTranslator("zh")
	// 注册翻译器
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		_ = zh_translations.RegisterDefaultTranslations(v, trans)
		_ = v.RegisterValidation("demo", func(fl validator.FieldLevel) bool {
			return false
		})

		_ = v.RegisterTranslation("demo", trans, func(ut ut.Translator) error {
			err := ut.Add("demo", "{0}xxxx", false)
			if err != nil {
				return err
			}
			return nil
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, err := ut.T(fe.Tag(), fe.Field())
			if err != nil {
				return ""
			}
			return t
		})

		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			label := field.Tag.Get("label")
			if len(label) > 0 {
				return label
			} else {
				return field.Name
			}
		})
	}
}

type User struct {
	Name  string `form:"name" binding:"demo" label:"名字"`
	Email string `form:"email" binding:"required,email"`
}

func main() {
	r := gin.Default()
	// 注册路由
	r.POST("/user", func(c *gin.Context) {

		var user User
		if err := c.ShouldBindWith(&user, binding.Form); err != nil {
			// 参数验证失败
			var errs validator.ValidationErrors
			ok := errors.As(err, &errs)
			if !ok {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			// 获取错误信息
			errMsg := make(map[string]string)

			for _, e := range errs {
				errMsg[e.StructField()] = e.Translate(trans)
			}

			c.JSON(http.StatusBadRequest, gin.H{
				"error": errMsg,
			})
			return
		}

		// 参数验证成功
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Hello, %s! Your email is %s.", user.Name, user.Email),
		})
	})

	// 启动HTTP服务器
	r.Run(":8899")
}
