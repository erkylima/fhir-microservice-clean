package formatter

import "testing"

func TestCPFFormatter(t *testing.T) {
	cpf := "22222222222"
	expected := "222.222.222-22"
	result := CPFFormatter(cpf)
	if result != expected {
		t.Fail()
	}
}

func TestRGFormatter(t *testing.T) {
	rg := "333333333"
	expected := "33.333.333-3"
	result := RGFormatter(rg)
	if result != expected {
		t.Fail()
	}
}
