package service

import (
	"testing"

	service "github.com/cassiano-medeiros/parseserver/src/service"
)

func TestValidateCpf(t *testing.T) {
	expected := true
	result := service.ValidateCpf("562.196.170-61")

	if result != expected {
		t.Error(
			"expected", expected,
			"got", result,
		)
	}

	expected = false
	result = service.ValidateCpf("562")

	if result != expected {
		t.Error(
			"expected", expected,
			"got", result,
		)
	}
}

func TestValidateCnpj(t *testing.T) {
	expected := true
	result := service.ValidateCnpj("49.976.547/0001-16")

	if result != expected {
		t.Error(
			"expected", expected,
			"got", result,
		)
	}

	expected = false
	result = service.ValidateCnpj("123")

	if result != expected {
		t.Error(
			"expected", expected,
			"got", result,
		)
	}
}
