package config

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/lib/pq"
	"log"
)

// SetupDatabase configura a conexão com o banco de dados
// Nesse caso, o * indica que a função retorna um ponteiro para sql.DB -> Singleton
// Com isso, podemos garantir que a conexão com o banco de dados seja reutilizada em toda a aplicação
// evitando a criação de múltiplas conexões desnecessárias
func SetupDatabase() *sql.DB {
	// Lógica para configurar a conexão com o banco de dados

	// Carregar variáveis de ambiente do arquivo .env
	// se o arquivo .env não existir, a aplicação pode falhar
	// log.fatal encerra a aplicação em caso de erro

	// Usando docker essa funcionalidade não é mais necessária
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	// }	

	// Exemplo de string de conexão (ajuste conforme necessário)
	// fmt é importado para formatar a string de conexão
	// os é importado para acessar variáveis de ambiente
	connectionSts := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	// Abre a conexão com o banco de dados
	dbConnection, err := sql.Open("postgres", connectionSts)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	// Verifica se a conexão está ativa
	// ping é um método que verifica a conexão com o banco de dados
	// se a conexão falhar, a aplicação é encerrada
	err = dbConnection.Ping()	

	if err != nil {
		log.Fatalf("Erro ao verificar a conexão com o banco de dados: %v", err)
	}

	fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")
	return dbConnection

}