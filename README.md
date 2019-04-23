# Requisitos

O Docker deve estar instalado. 


# Utiliza��o

## Passo 1: Conex�o com o banco de dados postresql

Editar o arquivo "desenv_config.json" informando Host, User, Password e DBName.

## Passo 2: Criar imagem docker

$ docker build -t parseserver:latest .

## Passo 3: Rodar o container com a imagem criada

$ docker run -d parseserver:latest .

## Passo 4: Acesso ao servi�o

Acessar o servi�o na porta :8080\upload

Selecionar o arquivo .txt

Enviar

## Passo 5: Confer�ncia dos dados
Acessar a tabela Import e verificar se as colunas e dados foram criados conforme o arquivo.


# Notas
Voc� pode executar o servi�o localmente se tiver o Golang instalado.

No diretorio src

$ go run main.go


