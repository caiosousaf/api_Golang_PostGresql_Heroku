package models

import "time"

type TasksbyTeam struct {
	ID_Task        uint       `json:"id_task"`
	Descricao_Task string     `json:"descricao_task"`
	Status         string     `json:"status" enums:"Em Andamento, Concluido"`
	Projeto_ID     uint       `json:"projeto_id"`
	Data_Criacao   *time.Time `json:"data_criacao"`
	Data_Conclusao *time.Time `json:"data_conclusao"`
	Prazo_Entrega  *time.Time `json:"prazo_entrega"`
	Prioridade     int        `json:"prioridade"`
}
