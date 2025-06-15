package utils

import (
	"text/template"

	"github.com/gin-gonic/gin"
)

func SetUpFuncs(r *gin.Engine) {
	r.SetFuncMap(template.FuncMap{
		"mul": func(a float64, b float64) float64 {
			return a * b
		},
	})
}
