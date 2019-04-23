package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/cassiano-medeiros/parseserver/src/general"
	"github.com/cassiano-medeiros/parseserver/src/repository/postgres"
	"github.com/cassiano-medeiros/parseserver/src/service"
)

const defaultPort = "8080"

func openWebPage(w http.ResponseWriter) {
	webpage, err := template.ParseFiles("template.html")

	if general.CheckError(err, "") {
		return
	}

	webpage.Execute(w, nil)
}

// upload - Funcao principal para upload do arquivo
func upload(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		openWebPage(w)
	case "POST":
		{
			db := postgres.Connect()

			headers := service.ExtractHeaderFile("Import.txt")
			postgres.CreateTable(headers, db)

			newFileName := service.FormatFile("Import.txt")
			postgres.CopyTableFromFile(headers, newFileName, db)

			//TODO: Falta parsear o CPF e CNPJ - Criar uma coluna nova?

			defer db.Close()
		}
	default:
		fmt.Println("Metodo desconhecido")
	}
}

func main() {
	http.HandleFunc("/upload", upload)

	fmt.Println("Servidor iniciado na porta " + defaultPort)
	fmt.Println("Acesse /upload")

	http.ListenAndServe(":"+defaultPort, nil)
}
