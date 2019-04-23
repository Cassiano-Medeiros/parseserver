package service

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

const separator = ";"

//calculateColumns - Funcao para calcular o tamanho de cada coluna
func calculateColumns(list []string) (columns []int) {
	size := 0
	secondLine := list[1]

	for i := 0; i < len(secondLine)-1; i++ {
		size++
		if (string(secondLine[i]) == " ") && (string(secondLine[i+1]) != " ") {
			columns = append(columns, size)
		}
	}
	return
}

//isMn - Funcao de apoio para normalizacao do texto unicode
func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r)
}

//textNormalization - Funcao para normalizacao do texto unicode
func textNormalization(s string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	result, _, _ := transform.String(t, s)
	return result
}

//removeWhiteSpace - Funcao para remover espacos em branco
func removeWhiteSpace(text string) string {
	text = strings.TrimSpace(text)
	text = strings.Replace(text, " ", "_", -1)
	return text
}

//generateHeader - Funcao para gerar o conteudo do cabecalho
func generateHeader(list []string, columns []int) (headers []string) {
	headerline := textNormalization(list[0])
	header := ""
	id := 0

	for i := 0; i < len(headerline); i++ {
		header = header + string(headerline[i])

		if (id < len(columns)) && (i >= columns[id]-1) {
			headers = append(headers, removeWhiteSpace(header))
			header = ""
			id++
		}
	}
	headers = append(headers, removeWhiteSpace(header))
	return
}

// ExtractHeaderFile - Funcao para "parsear" os arquivos texto e extrair o cabecalho
func ExtractHeaderFile(fileName string) []string {
	list := FileToList(fileName)

	columns := calculateColumns(list)

	headers := generateHeader(list, columns)

	fmt.Println(headers)
	return headers
}

// FormatFile - Funcao para "parsear" os arquivos texto separados em colunas
func FormatFile(fileName string) string {
	list := FileToList(fileName)

	var re = regexp.MustCompile(" +")
	var fileFormated []string
	var newFile = "file_to_import_db.txt"

	isHeader := true

	for _, line := range list {
		if isHeader {
			isHeader = false
			continue
		}

		line := re.ReplaceAllString(line, separator)
		fileFormated = append(fileFormated, line)
	}

	SaveNewFile(fileFormated, newFile)
	return newFile
}
