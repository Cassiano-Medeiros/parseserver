package service

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/cassiano-medeiros/parseserver/src/general"
)

//MsgSuccess - Mensagem padrao
const MsgSuccess = "Arquivo %s importado"

//DefaultFolder - Diretorio padrao dos arquivos
const DefaultFolder = "./files/"

//GetAbsolutePath - Retorna o caminho completo do arquivo
func GetAbsolutePath(fileName string) string {
	absolute, err := filepath.Abs(DefaultFolder + fileName)
	general.CheckError(err, "")
	return absolute
}

// SaveNewFile - Funcao que escreve um texto local
func SaveNewFile(lines []string, fileName string) {
	newFile, err := os.Create(DefaultFolder + fileName)

	general.CheckError(err, "")

	defer newFile.Close()

	writer := bufio.NewWriter(newFile)
	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}

	writer.Flush()
}

//FileToList - Funcao para abrir o arquivo local e gerar uma lista "string" de saida
func FileToList(fileName string) (list []string) {
	file, err := os.Open(DefaultFolder + fileName)

	general.CheckError(err, "Erro ao abrir o arquivo "+fileName)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		list = append(list, scanner.Text())
	}

	file.Close()
	return
}

//SaveFormFile - funcao que salva o arquivo em disco
func SaveFormFile(w http.ResponseWriter, r *http.Request) (fileName string) {
	formFile, handler, err := r.FormFile("uploadfile")

	if general.CheckError(err, "") {
		return
	}

	fileName = handler.Filename

	fmt.Printf(MsgSuccess, fileName)

	defer formFile.Close()

	localfile, err := os.OpenFile(DefaultFolder+fileName, os.O_WRONLY|os.O_CREATE, 0666)

	if general.CheckError(err, "") {
		return
	}

	defer localfile.Close()

	io.Copy(localfile, formFile)

	return fileName
}
