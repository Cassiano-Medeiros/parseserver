package postgres

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/cassiano-medeiros/parseserver/src/general"
	"github.com/cassiano-medeiros/parseserver/src/service"
)

//CreateTable - Cria a tabela no banco de dados baseado na estrutura do arquivo a ser importado
func CreateTable(columns []string, db *sql.DB) {
	sql := "CREATE TABLE IF NOT EXISTS Import (id SERIAL PRIMARY KEY"

	for i := 0; i < len(columns); i++ {
		sql = sql + ", " + columns[i] + " TEXT"
	}

	sql = sql + ")"

	_, err := db.Exec(sql)

	general.CheckError(err, "Erro ao criar a tabela no banco de dados")
}

//CopyTableFromFile - Importa o arquivo local para o banco de dados
func CopyTableFromFile(colunms []string, fileName string, db *sql.DB) {
	sql := "COPY Import ("

	for i := 0; i < len(colunms); i++ {
		sql = sql + colunms[i] + ", "
	}

	sql = strings.TrimSuffix(sql, ", ")

	sql = sql + ") FROM '" + service.GetAbsolutePath(fileName) + "' (DELIMITER(';'))"

	fmt.Println(sql)

	_, err := db.Exec(sql)

	general.CheckError(err, "Erro ao inserir os dados na tabela Import")
}
