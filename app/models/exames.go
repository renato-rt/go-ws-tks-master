package models

import (
	"fmt"

	"github.com/nathanzeras/go-ws-tks/config"
)

type ProcedimentosSolicitados struct {
	Exames []Exames `json:"procedimentosSolicitados"`
}

type Exames struct {
	NumeroGuiaPrestador   string `json:"numeroGuiaPrestador"`
	CodigoProcedimento    string `json:"codigoProcedimento"`
	DescricaoProcedimento string `json:"descricaoProcedimento"`
	QuantidadeSolicitada  int    `json:"quantidadeSolicitada"`
	QuantidadeAutorizada  int    `json:"quantidadeAutorizada"`
	QuantidadeRealizada   int    `json:"quantidadeRealizada"`
	CodigoStatus          string `json:"codigoStatus"`
	DescricaoGlosa        string `json:"descricaoGlosa"`
	CodigoGlosa           string `json:"codigoGlosa"`
	CodigoBarra           string `json:"codigoBarra"`
	CodigoTabela          string `json:"codigoTabela"`
	Data                  string `json:"data"`
	Fabricante            string `json:"fabricante"`
	FatorRedAcresc        string `json:"fatorRedAcresc"`
	HoraFinal             string `json:"horaFinal"`
	HoraInicial           string `json:"horaInicial"`
	Tecnica               string `json:"tecnica"`
	UnidadeMedida         string `json:"unidadeMedida"`
	Via					  string `json:"via"`				
}

//Função que verifica se já existe um paciente cadastrado, caso contrário, realiza a inserção no Clinux
func (pedido *Pedidos) CreateExames(exames ProcedimentosSolicitados) int {
	var cdExame int
	//var err error

	db := config.ConnectDB()

	for _, exame := range exames.Exames {

		err := db.QueryRow(`insert into exames (ds_guia_prestador, cd_procedimento, ds_procedimento, qtd_solicitada,
			qtd_autorizada, cd_status, ds_glosa, cd_glosa, qtd_realizada, cd_barra, cd_tabela, ds_fabricante, 
			ds_fator_red_acresc, ds_tecnica, ds_unidade_medida, ds_via, ds_hora_inicial, ds_hora_final, dt_procedimento) 
			values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19) returning id`,
			pedido.NumeroGuiaPrestador, exame.CodigoProcedimento, exame.DescricaoProcedimento, exame.QuantidadeSolicitada,
			exame.QuantidadeAutorizada, exame.CodigoStatus, exame.DescricaoGlosa, exame.CodigoGlosa, exame.QuantidadeRealizada,
			exame.CodigoBarra, exame.CodigoTabela, exame.Fabricante, exame.FatorRedAcresc, exame.Tecnica, exame.UnidadeMedida,
			exame.Via, exame.HoraInicial, exame.HoraFinal, exame.Data).Scan(&cdExame)
		if err != nil {
			//handle error
			fmt.Printf("Error during insert exames: %v", err)
			//log.Fatalf("Error during insert exames: %v", err)
		}
	}

	defer db.Close()

	return cdExame
}
