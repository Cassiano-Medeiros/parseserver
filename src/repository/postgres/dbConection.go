package postgres

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"

	"github.com/cassiano-medeiros/parseserver/src/general"
)

//DataBaseConfiguration - Mapeamento do arquivo de configuracao
type DataBaseConfiguration struct {
	Host     string `json: "host"`
	Port     string `json: "port"`
	User     string `json: "user"`
	DbName   string `json: "dbname"`
	Password string `json: "password"`
}

const testConnectionString = "host=127.0.0.1 port=5432 user=postgres dbname=parsedb password=teste123 sslmode=disable"

//loadConfig - Le o arquivo de configuracao
func loadConfig() string {
	file, _ := ioutil.ReadFile("./desenv_config.json")
	config := DataBaseConfiguration{}

	_ = json.Unmarshal([]byte(file), &config)

	connectionString := "host=" + config.Host + " port=" + config.Port + " user=" + config.User + " dbname=" + config.DbName + " password=" + config.Password + " sslmode=disable"

	return connectionString
}

//Connect - Conexao com o banco de dados
func Connect() (db *sql.DB) {
	db, err := sql.Open("postgres", loadConfig())
	general.CheckError(err, "Erro ao conectar no banco de dados. Verifique o arquivo de configuracao")
	return
}
