package service

import (
	"strconv"
	"strings"
)

//ValidateCnpj - Funcao para validar o CNPJ
//TODO: Funcao copiada - Usar algum servico pronto
func ValidateCnpj(documento string) bool {
	documento = strings.Replace(documento, ".", "", -1)
	documento = strings.Replace(documento, "-", "", -1)
	documento = strings.Replace(documento, "/", "", -1)
	if len(documento) != 14 {
		return false
	}
	algs := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	var algProdCpfDig1 = make([]int, 12, 12)
	for key, val := range algs {
		intParsed, _ := strconv.Atoi(string(documento[key]))
		sumTmp := val * intParsed
		algProdCpfDig1[key] = sumTmp
	}
	sum := 0
	for _, val := range algProdCpfDig1 {
		sum += val
	}
	digit1 := sum % 11
	if digit1 < 2 {
		digit1 = 0
	} else {
		digit1 = 11 - digit1
	}
	char12, _ := strconv.Atoi(string(documento[12]))
	if char12 != digit1 {
		return false
	}
	algs = append([]int{6}, algs...)
	var algProdCpfDig2 = make([]int, 13, 13)
	for key, val := range algs {
		intParsed, _ := strconv.Atoi(string(documento[key]))
		sumTmp := val * intParsed
		algProdCpfDig2[key] = sumTmp
	}
	sum = 0
	for _, val := range algProdCpfDig2 {
		sum += val
	}
	digit2 := sum % 11
	if digit2 < 2 {
		digit2 = 0
	} else {
		digit2 = 11 - digit2
	}
	char13, _ := strconv.Atoi(string(documento[13]))
	if char13 != digit2 {
		return false
	}
	return true
}

//ValidateCpf - Funcao para validar o documento
//TODO: Funcao copiada - Usar algum servico pronto
func ValidateCpf(documento string) bool {
	documento = strings.Replace(documento, ".", "", -1)
	documento = strings.Replace(documento, "-", "", -1)
	if len(documento) != 11 {
		return false
	}
	var eq bool
	var dig string
	for _, val := range documento {
		if len(dig) == 0 {
			dig = string(val)
		}
		if string(val) == dig {
			eq = true
			continue
		}
		eq = false
		break
	}
	if eq {
		return false
	}
	i := 10
	sum := 0
	for index := 0; index < len(documento)-2; index++ {
		pos, _ := strconv.Atoi(string(documento[index]))
		sum += pos * i
		i--
	}
	prod := sum * 10
	mod := prod % 11
	if mod == 10 {
		mod = 0
	}
	digit1, _ := strconv.Atoi(string(documento[9]))
	if mod != digit1 {
		return false
	}
	i = 11
	sum = 0
	for index := 0; index < len(documento)-1; index++ {
		pos, _ := strconv.Atoi(string(documento[index]))
		sum += pos * i
		i--
	}
	prod = sum * 10
	mod = prod % 11
	if mod == 10 {
		mod = 0
	}
	digit2, _ := strconv.Atoi(string(documento[10]))
	if mod != digit2 {
		return false
	}
	return true
}
