package handlers

import(
	"database/sql"
	"net/http"
	"go-api/models"
	"encoding/json"
)

type TaskHandler struct {
	// Aqui você pode adicionar dependências, como o banco de dados
	DB *sql.DB
}

// Construtor para TaskHandler 
// Função que retorna um ponteiro para TaskHandler
// Recebe como parâmetro a conexão com o banco de dados
// e retorna uma instância de TaskHandler
// Isso facilita a injeção de dependências e a testabilidade do código
func NewTaskHandler(db *sql.DB) *TaskHandler {
	return &TaskHandler{DB: db}
}

func (ReadTaskHandler *TaskHandler) ReadTasks(writer http.ResponseWriter, request *http.Request) {
	// Lógica para ler tarefas do banco de dados
	rows, err := ReadTaskHandler.DB.Query(models.GET_TASKS_QUERY)
	
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	
	// Criacao de um array para armazenar as tarefas
	var tasks []models.Task

	// Iterar sobre as linhas retornadas pela consulta
	// rows.Next() é uma função que avança para a próxima linha do resultado
	// é nativa do pacote database/sql
	// enquanto houver linhas, o loop continua
	for rows.Next() {
		var task models.Task

		// Scanear os dados da linha para a struct Task
		// rows.Scan é uma função que mapeia os dados da linha para as variáveis fornecidas
		// é nativa do pacote database/sql
		// recebe os endereços das variáveis onde os dados serão armazenados
		// retorna um erro caso ocorra algum problema durante o mapeamento
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status);

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}	

		// Adicionar a tarefa ao array de tarefas
		tasks = append(tasks, task)
	}

	// writer.header().Set é uma função que define o cabeçalho da resposta HTTP
	// writer.WriteHeader é uma função que define o status da resposta HTTP
	// json.NewEncoder(writer).Encode é uma função que codifica os dados em JSON e os escreve na resposta HTTP
	// todas essas funções são nativas do pacote net/http e encoding/json
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(tasks)

}

func (CreateTaskHandler *TaskHandler) CreateTask(writer http.ResponseWriter, request *http.Request) {
	// Lógica para criar uma nova tarefa no banco de dados
	var task models.Task	

	// Decodificar o corpo da requisição JSON para a struct Task
	// json.NewDecoder é uma função que cria um decodificador JSON
	// Decode é uma função que decodifica os dados JSON para a struct fornecida
	// ambas são nativas do pacote encoding/json
	err := json.NewDecoder(request.Body).Decode(&task)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}	

	// Inserir a nova tarefa no banco de dados
	// Pode ser usado Exec ou QueryRow dependendo se você espera um retorno
	// Aqui usamos Exec pois não precisamos do ID retornado
	_, err = CreateTaskHandler.DB.Exec(models.INSERT_TASK_QUERY, task.Title, task.Description, task.Status)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(task)

	// JSON de exemplo para teste
	// {
	// 	"title": "Nova Tarefa",
	// 	"description": "Descrição da nova tarefa",
	// 	"status": false
	// }
}	

// func (GetTaskByIDHandler *TaskHandler) GetTaskByID(writer http.ResponseWriter, request *http.Request) {
// 	// var taskID string
// 	// Extrair o ID da tarefa dos parâmetros da URL
// 	// Pode ser usado mux.Vars(request) se estiver usando gorilla/mux
// 	// taskID = mux.Vars(request)["id"]
// 	// Lógica para ler uma tarefa específica do banco de dados
// 	// TODO: Pesquisar melhor sobre como funciona a extração de parâmetros com gorilla/mux
// }