package main

import (
	"fmt"

	"github.com/diegors-prog/api-go-rest/database"
	"github.com/diegors-prog/api-go-rest/models"
	"github.com/diegors-prog/api-go-rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "História 1"},
		{Id: 2, Nome: "Nome 2", Historia: "História 2"},
	}
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando servidor Rest com Go")
	routes.HandleRequest()
}
