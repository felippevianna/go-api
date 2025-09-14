package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/mux"

	"go-api/config"
	"go-api/models"
	"go-api/handlers"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Olá, Go API!")
}

func main() {
	
	http.HandleFunc("/hello", helloHandler)
	
	dbConnection := config.SetupDatabase()
	
	_, err := dbConnection.Exec(models.CREATE_TABLE_QUERY)
	if err != nil {
		log.Fatalf("Erro ao criar a tabela: %v", err)
	}
	fmt.Println("Tabela Tasks criada ou já existe.");
	
	// Fechar a conexão com o banco de dados quando a aplicação for encerrada
	defer dbConnection.Close()

	// Criação das rotas da API
	router := mux.NewRouter()

	// Instanciando o TaskHandler com a conexão do banco de dados
	// e armazenando na variável taskHandler
	// que será usada para mapear as rotas
	taskHandler := handlers.NewTaskHandler(dbConnection)

	// Mapeando as rodas para os handlers
	// Criando um endpoint de exemplo para testar a API
	// Exemplo de rota: GET /hello
	router.HandleFunc("/hello", helloHandler).Methods("GET")
	// Rota para ler as tarefas : GET /tasks
	router.HandleFunc("/tasks", taskHandler.ReadTasks).Methods("GET")
	// Rota para criar uma nova tarefa : POST /tasks
	router.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	// Rota para ler uma tarefa específica : GET /tasks/{id}
	//router.HandleFunc("/tasks/{id}", taskHandler.GetTaskByID).Methods("GET")
	// Rota para atualizar uma tarefa específica : PUT /tasks/{id}
	//router.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")
	// Rota para deletar uma tarefa específica : DELETE /tasks/{id}
	//router.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")


	fmt.Println("Servidor rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}