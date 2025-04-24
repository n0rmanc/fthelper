package xtemplates

import (
	"strings"
	"text/template"

	"github.com/n0rmanc/fthelper/shared/datatype"
	"github.com/n0rmanc/fthelper/shared/utils"
)

func join(input ...interface{}) string {
	var str []string = make([]string, 0)
	for _, i := range input {
		str = append(str, datatype.ForceString(i))
	}

	return utils.JoinString("-", str...)
}

func joinArray(input []interface{}) string {
	var str []string = make([]string, 0)
	for _, i := range input {
		str = append(str, datatype.ForceString(i))
	}

	return utils.JoinString(",", str...)
}

var stringFuncs template.FuncMap = map[string]interface{}{
	"join":      join,
	"joinArray": joinArray,
	"toLower":   strings.ToLower,
	"toUpper":   strings.ToUpper,
	"toTitle":   strings.ToTitle,
}
