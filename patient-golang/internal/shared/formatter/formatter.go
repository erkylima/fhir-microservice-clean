package formatter

import (
	"fmt"
	"strings"
)

func RemoveFormat(text string) string {
	text = strings.ReplaceAll(text, ".", "")
	text = strings.ReplaceAll(text, "-", "")
	text = strings.ReplaceAll(text, " ", "")
	return text
}

func CPFFormatter(cpf string) string {
	return fmt.Sprintf("%s.%s.%s-%s", cpf[0:3], cpf[3:6], cpf[6:9], cpf[9:11])
}

func RGFormatter(rg string) string {
	return fmt.Sprintf("%s.%s.%s-%s", rg[0:2], rg[2:5], rg[5:7], rg[7:])
}

func CelphoneFormatter(rg string) string {
	return fmt.Sprintf("(%s) %s", rg[0:2], rg[2:])
}
